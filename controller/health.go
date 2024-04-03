package controller

import (
	"encoding/json"
	"net/http"

	"gitlab.com/DeveloperDurp/DurpAPI/model"
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
	message := model.Message{
		Message: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
