package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.com/DeveloperDurp/DurpAPI/model"
	"gitlab.com/DeveloperDurp/DurpAPI/service"
)

// GetDadJoke godoc
//
//	@Summary		Get dadjoke
//	@Description	get a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Success		200	{object}	model.Message	"response"
//	@failure		500	{object}	model.Message	"error"
//	@Router			/jokes/dadjoke [get]
func (c *Controller) GetDadJoke(ctx *gin.Context) {
	joke, err := service.GetRandomDadJoke(c.Db.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": joke})
}

// PostDadJoke godoc
//
//	@Summary		Generate dadjoke
//	@Description	create a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Param			joke	query		string			true	"Dad Joke you wish to enter into database"
//	@Success		200		{object}	model.Message	"response"
//	@failure		500		{object}	model.Message	"error"
//	@Router			/jokes/dadjoke [post]
func (c *Controller) PostDadJoke(ctx *gin.Context) {
	var req model.DadJoke

	if err := ctx.ShouldBindJSON(&req); err != nil {
		req.JOKE = ctx.Query("joke")
	}

	err := service.PostDadJoke(c.Db.DB, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// DeleteDadJoke godoc
//
//	@Summary		Generate dadjoke
//	@Description	create a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Param			joke	query		string			true	"Dad joke you wish to delete from the database"
//	@Success		200		{object}	model.Message	"response"
//	@failure		500		{object}	model.Message	"error"
//	@Router			/jokes/dadjoke [delete]
func (c *Controller) DeleteDadJoke(ctx *gin.Context) {
	var req model.DadJoke

	if err := ctx.ShouldBindJSON(&req); err != nil {
		req.JOKE = ctx.Query("joke")
	}

	err := service.DeleteDadJoke(c.Db.DB, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
