package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
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

//	@BasePath	/api/v1

func main() {
	r := gin.Default()
	c := controller.NewController()

	docs.SwaggerInfo.Host = c.Cfg.Host
	docs.SwaggerInfo.Version = c.Cfg.Version

	v1 := r.Group("/api/v1")
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

		if groupsenv != "" {
			groups = strings.Split(groupsenv, ",")
		} else {
			groupsHeader := c.GetHeader("X-authentik-groups")

			fmt.Println(c.GetHeader("X-authentik-name"))
			fmt.Println(groupsHeader)
			groups = strings.Split(groupsHeader, "|")
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

		// Call the next handler
		c.Next()
	}
}
