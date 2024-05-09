package controller

import (
	"net/http"

	"gitlab.com/developerdurp/logger"
	"gitlab.com/developerdurp/stdmodels"
)

// getHealth godoc
//
//	@Summary		Generate Health status
//	@Description	Get the health of the API
//	@Tags			health
//	@Accept			json
//	@Produce		application/json
//	@Success		200	{object}	model.Message	"response"
//	@failure		500	{object}	model.Message	"error"
//
// @Security Authorization
//
//	@Router			/health/gethealth [get]
func (c *Controller) GetHealth(w http.ResponseWriter, r *http.Request) {
	logger.LogInfo("Health Check")
	stdmodels.SuccessResponse("OK", w, http.StatusOK)
}
