package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// GenerateToken godoc
//
//	@Summary		Generate Health status
//	@Description	Get the health of the API
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json	"response"
//	@Router			/token/GenerateToken [get]
func (c *Controller) GenerateToken(conf *oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Redirect user to the authorization URL
		authURL := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		c.Redirect(http.StatusTemporaryRedirect, authURL)
	}
}
