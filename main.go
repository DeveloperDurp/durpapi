package main

import (
	"net/http"

	"github.com/DeveloperDurp/DurpAPI/controller"
	_ "github.com/DeveloperDurp/DurpAPI/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

//	@securityDefinitions.basic	BasicAuth

func main() {

	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
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
		// Get the username and password from the request header
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.Header("WWW-Authenticate", "Basic realm=Restricted")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if the username and password are valid
		if username != "user" || password != "password" {
			c.Header("WWW-Authenticate", "Basic realm=Restricted")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set the user ID in the context for later use
		userID := "user"
		c.Set("userID", userID)

		// Call the next middleware or handler function
		c.Next()
	}
}
