package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.com/DeveloperDurp/DurpAPI/model"
)

// GetDadJoke godoc
//
//	@Summary		Generate dadjoke
//	@Description get a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Success		200	{object}	model.Message	"response"
//	@failure		400 {object}	model.Message	"error"
//	@Router			/jokes/dadjoke [get]
func (c *Controller) GetDadJoke(ctx *gin.Context) {
	var req model.DadJoke

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	message := "test"
	//	message, err := service.GetDadJoke(req)
	//	if err != nil {
	//		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
	//	}
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
