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
			openai.Use(authMiddleware("test"))
			openai.GET("general", c.GeneralOpenAI)
			openai.GET("travelagent", c.TravelAgentOpenAI)
		}
		unraid := v1.Group("/unraid")
		{
			unraid.Use(authMiddleware("grafana"))
			unraid.GET("powerusage", c.UnraidPowerUsage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func authMiddleware(allowedGroups []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user groups from the request headers
		groupsHeader := c.GetHeader("X-Forwarded-Groups")

		// Split the groups header value into individual groups
		groups := strings.Split(groupsHeader, ",")

		// Check if the user belongs to any of the allowed groups
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

		// If the user is not in any of the allowed groups, respond with unauthorized access
		if !isAllowed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
			return
		}

		// Call the next handler
		c.Next()
	}
}
