package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/DeveloperDurp/DurpAPI/controller"
	_ "gitlab.com/DeveloperDurp/DurpAPI/docs"
)

//	@title			DurpAPI
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		durpapi.durp.info
//	@BasePath	/api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	r := gin.Default()
	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		token := v1.Group("/token")
		{
			token.GET("generateTokenHandler", c.GenerateTokenHandler)
		}
		openai := v1.Group("/openai")
		{
			openai.Use(authMiddleware())
			openai.GET("general", c.GeneralOpenAI)
			openai.GET("travelagent", c.TravelAgentOpenAI)
		}
		unraid := v1.Group("/unraid")
		{
			unraid.Use(authMiddleware())
			unraid.GET("powerusage", c.UnraidPowerUsage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the authorization header from the request
		authHeader := c.GetHeader("Authorization")

		// Check if the authorization header is missing or doesn't start with "Bearer"
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
			return
		}

		// Extract the token from the authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token and validate its signature
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtToken")), nil
		})

		// Check if there was an error parsing the token or if it is not valid
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
			return
		}

		// Add the token to the request context
		c.Set("token", token)

		// Call the next handler
		c.Next()
	}
}
