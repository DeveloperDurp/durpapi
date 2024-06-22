package health

import (
	"net/http"

	"gitlab.com/developerdurp/stdmodels"
)

type Handler struct{}

func NewHandler() (*Handler, error) {
	return &Handler{}, nil
}

// getHealth godoc
//
//	@Summary		Generate Health status
//	@Description	Get the health of the API
//	@Tags			health
//	@Accept			json
//	@Produce		application/json
//	@Success		200		{object}	stdmodels.StandardMessage	"response"
//	@failure		500	{object}	stdmodels.StandardError"error"
//
// @Security Authorization
//
//	@Router			/health/gethealth [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	stdmodels.SuccessResponse("OK", w, http.StatusOK)
}
