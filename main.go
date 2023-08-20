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
			jokes.POST("dadjoke", c.PostDadJoke)
			jokes.DELETE("dadjoke", c.DeleteDadJoke)
		}
		openai := v1.Group("/openai")
		{
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