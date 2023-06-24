package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"

	"gitlab.com/DeveloperDurp/DurpAPI/model"
	"gitlab.com/DeveloperDurp/DurpAPI/storage"
)

type Controller struct {
	Cfg   model.Config
	Dbcfg model.DBConfig
	Db    model.Repository
}

func NewController() *Controller {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load file: %e", err)
	}

	controller := &Controller{
		Cfg:   model.Config{},
		Dbcfg: model.DBConfig{},
	}

	err = env.Parse(&controller.Cfg)
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}
	err = env.Parse(&controller.Dbcfg)
	if err != nil {
		log.Fatalf("unable to parse database variables: %e", err)
	}

	controller.Cfg.OpenaiClient = *openai.NewClient(controller.Cfg.OpenaiApiKey)

	Db, err := storage.Connect(controller.Dbcfg)

	if err != nil {
		panic("Failed to connect to database")
	}
	controller.Db = *Db

	return controller
}

func (c *Controller) AuthMiddleware(
	allowedGroups []string,
	currentGroups string,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var groups []string

		if currentGroups != "" {
			groups = strings.Split(currentGroups, ",")
		} else {
			// Get the user groups from the request headers
			groupsHeader := c.GetHeader("X-authentik-groups")

			// Split the groups header value into individual groups
			groups = strings.Split(groupsHeader, "|")
		}

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
