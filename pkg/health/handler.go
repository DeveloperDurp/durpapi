package health

import (
	"gitlab.com/developerdurp/durpify/handlers"
	"net/http"
)

type Handler struct{}

func RegisterOpenAIHandler(
	router *http.ServeMux,
	handler *Handler,
) error {

	router.HandleFunc(
		"GET /health/gethealth",
		handlers.Make(handler.Get),
	)

	return nil
}

func NewHandler() *Handler {
	return &Handler{}
}

// getHealth godoc
//
//	@Summary		Generate Health status
//	@Description	Get the health of the API
//	@Tags			health
//	@Accept			json
//	@Produce		application/jso
//	@Success		200		{object}	handlers.StandardMessage	"response"
//	@failure		500	{object}	handlers.StandardError"error"
//
// @Security Authorization
//
//	@Router			/health/gethealth [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) (*handlers.StandardMessage, error) {

	// err := handlers.StandardError{
	//	Message: "This is a test",
	//	Status:  500,
	//	Description: []string{
	//		"Test",
	//		"test2",
	//	},
	// }
	// return nil, err

	resp := handlers.NewBasicResponse()
	return resp, nil
}
