package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/DeveloperDurp/DurpAPI/controller"
	"gitlab.com/DeveloperDurp/DurpAPI/docs"
	"gitlab.com/DeveloperDurp/DurpAPI/model"
	"golang.org/x/oauth2"
)

var (
	host    = model.Host
	version = model.Version
	Conf    = &oauth2.Config{
		ClientID:     model.ClientID,
		ClientSecret: model.ClientSecret,
		RedirectURL:  model.RedirectURL,
		Scopes: []string{
			"email",
			"groups",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  model.AuthURL,
			TokenURL: model.TokenURL,
		},
	}
)

//	@title			DurpAPI
//	@description	API for Durp's needs
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	r := gin.Default()
	c := controller.NewController()
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Version = version

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		{
			health.GET("getHealth", c.GetHealth)
		}
		token := v1.Group("/token")
		{
			token.GET("GenerateToken", c.GenerateToken(Conf))
		}
		openai := v1.Group("/openai")
		{
			//groups = []string{"openai"}
			//openai.Use(authMiddleware())
			openai.GET("general", c.GeneralOpenAI)
			openai.GET("travelagent", c.TravelAgentOpenAI)
		}
		unraid := v1.Group("/unraid")
		{
			//groups = []string{"unraid"}
			//unraid.Use(authMiddleware())
			unraid.GET("powerusage", c.UnraidPowerUsage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/callback", CallbackHandler(Conf))

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server")
	}
}

// CallbackHandler receives the authorization code and exchanges it for a token
func CallbackHandler(conf *oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the authorization code from the query parameters
		code := c.Query("code")

		// Exchange the authorization code for a token
		token, err := conf.Exchange(context.Background(), code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange authorization code"})
			return
		}

		// Create a response JSON
		response := gin.H{
			"access_token":  token.AccessToken,
			"token_type":    token.TokenType,
			"refresh_token": token.RefreshToken,
			"expiry":        token.Expiry,
		}

		c.JSON(http.StatusOK, response)
	}
}
