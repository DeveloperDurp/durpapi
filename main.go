package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gitlab.com/DeveloperDurp/DurpAPI/controller"
	"gitlab.com/DeveloperDurp/DurpAPI/docs"
)

var groupsenv = os.Getenv("groups")

//	@title			DurpAPI
//	@description	API for Durp's needs
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	https://durp.info
//	@contact.email	developerdurp@durp.info

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath					/api
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	r := gin.Default()
	c := controller.NewController()

	docs.SwaggerInfo.Host = c.Cfg.Host
	docs.SwaggerInfo.Version = c.Cfg.Version

	v1 := r.Group("/api")
	{
		health := v1.Group("/health")
		{
			health.GET("getHealth", c.GetHealth)
		}
		jokes := v1.Group("/jokes")
		{
			jokes.GET("dadjoke", c.GetDadJoke)

			jokes.Use(authMiddleware([]string{"rw-jokes"}))
			jokes.POST("dadjoke", c.PostDadJoke)
			jokes.DELETE("dadjoke", c.DeleteDadJoke)
		}
		openai := v1.Group("/openai")
		{
			openai.Use(authMiddleware([]string{"openai"}))
			openai.GET("general", c.GeneralOpenAI)
			openai.GET("travelagent", c.TravelAgentOpenAI)
		}
		unraid := v1.Group("/unraid")
		{
			unraid.Use(authMiddleware([]string{"unraid"}))
			unraid.GET("powerusage", c.UnraidPowerUsage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server")
	}
}

func authMiddleware(allowedGroups []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var groups []string
		JwksURL := os.Getenv("jwksurl")
		tokenString := c.GetHeader("Authorization")
		if tokenString != "" {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		} else {
			tokenString = c.GetHeader("X-authentik-jwt")
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "No Token in header",
			})
			return
		}

		ctx, cancel := context.WithCancel(context.Background())

		options := keyfunc.Options{
			Ctx: ctx,
			RefreshErrorHandler: func(err error) {
				log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
			},
			RefreshInterval:   time.Hour,
			RefreshRateLimit:  time.Minute * 5,
			RefreshTimeout:    time.Second * 10,
			RefreshUnknownKID: true,
		}

		jwks, err := keyfunc.Get(JwksURL, options)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Failed to create JWKS: " + err.Error(),
			})
			cancel()
			jwks.EndBackground()
			return
		}

		token, err := jwt.Parse(tokenString, jwks.Keyfunc)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			cancel()
			jwks.EndBackground()
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token: " + err.Error(),
			})
			cancel()
			jwks.EndBackground()
			return
		}

		cancel()
		jwks.EndBackground()

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization token claims",
			})
			return
		}

		groupsClaim, ok := claims["groups"].([]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Missing or invalid groups claim in the authorization token",
			})
			return
		}

		for _, group := range groupsClaim {
			if groupName, ok := group.(string); ok {
				groups = append(groups, groupName)
			}
		}

		isAllowed := false
		for _, allowedGroup := range allowedGroups {
			for _, group := range groups {
				if group == allowedGroup {
					isAllowed = true
					break
				}
			}
			if isAllowed {
				break
			}
		}

		if !isAllowed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized access",
				"groups":  groups,
			})
			return
		}

		c.Next()
	}
}
