package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/DeveloperDurp/DurpAPI/controller"
	_ "gitlab.com/DeveloperDurp/DurpAPI/docs"
	"golang.org/x/oauth2"
)

//	@title			DurpAPI
//	@version		1.0
//	@description	API for Durp's needs
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
	var groups []string
	conf := &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Scopes: []string{
			"email",
			"groups",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.AuthURL,
			TokenURL: c.TokenURL,
		},
	}

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		{
			health.GET("getHealth", c.GetHealth)
		}
		token := v1.Group("/token")
		{
			token.GET("GenerateToken", c.GenerateToken(conf))
		}
		openai := v1.Group("/openai")
		{
			groups = []string{"openai"}
			openai.Use(authMiddleware(groups))
			openai.GET("general", c.GeneralOpenAI)
			openai.GET("travelagent", c.TravelAgentOpenAI)
		}
		unraid := v1.Group("/unraid")
		{
			groups = []string{"unraid"}
			unraid.Use(authMiddleware(groups))
			unraid.GET("powerusage", c.UnraidPowerUsage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/callback", CallbackHandler(conf))

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server")
	}
}

func authMiddleware(groups []string) gin.HandlerFunc {

	return func(c *gin.Context) {
		// Get the access token from the request header or query parameters
		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			accessToken = c.Query("access_token")
		}

		// Create an OAuth2 token from the access token
		token := &oauth2.Token{AccessToken: accessToken}

		// Validate the token
		if !token.Valid() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Add the token to the request context for later use
		//ctx := context.WithValue(c.Request.Context(), "token", token)
		//c.Request = c.Request.WithContext(ctx)

		// Call the next handler
		c.Next()
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
