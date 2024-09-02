package health

import (
	"gitlab.com/developerdurp/stdmodels"
	"net/http"
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
//	@Produce		application/jso
//	@Success		200		{object}	stdmodels.StandardMessage	"response"
//	@failure		500	{object}	stdmodels.StandardError"error"
//
// @Security Authorization
//
//	@Router			/health/gethealth [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) (*stdmodels.StandardMessage, error) {

	// err := stdmodels.StandardError{
	//	Message: "This is a test",
	//	Status:  500,
	//	Description: []string{
	//		"Test",
	//		"test2",
	//	},
	// }
	// return nil, err

	resp := stdmodels.NewBasicResponse()
	return resp, nil
}
