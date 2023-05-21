package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getHealth godoc
//
//	@Summary		Generate Health status
//	@Description	Get the health of the API
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json	"response"
//	@Router			/health/getHealth [get]
func (c *Controller) GetHealth(ctx *gin.Context) {
	// Return the health in the response body
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
