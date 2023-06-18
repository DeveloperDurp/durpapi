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
//	@Produce		application/json
//	@Success		200	{object}	model.Message	"response"
//	@failure		400 {object}	model.Message	"error"
//	@Router			/health/getHealth [get]
func (c *Controller) GetHealth(ctx *gin.Context) {
	// Return the health in the response body
  ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}

