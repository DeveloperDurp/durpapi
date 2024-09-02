package dadjoke

import (
	"gitlab.com/DeveloperDurp/DurpAPI/pkg/shared"
	"net/http"

	"gitlab.com/developerdurp/stdmodels"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

type DadJoke struct {
	JOKE string `json:"joke"`
}

func NewHandler(db *gorm.DB) (*Handler, error) {
	err := db.AutoMigrate(&DadJoke{})
	if err != nil {
		return nil, err
	}
	return &Handler{db: db}, nil
}

// GetDadJoke godoc
//
//	@Summary		Get dadjoke
//	@Description	get a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Success		200	{object}	DadJoke	"response"
//	@failure		500	{object}	stdmodels.StandardError"error"
//
//	@Security		Authorization
//
//	@Router			/jokes/dadjoke [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) (*stdmodels.StandardMessage, error) {
	joke, err := h.GetRandomDadJoke()

	if err != nil {
		resp := stdmodels.NewFailureResponse("Failed to get Joke",
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
		return nil, resp
	}

	resp := stdmodels.NewMessageResponse(joke, http.StatusOK)
	return resp, nil
}

// PostDadJoke godoc
//
//	@Summary		Generate dadjoke
//	@Description	create a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Param			joke	query		string						true	"Dad Joke you wish to enter into database"
//	@Success		200		{object}	stdmodels.StandardMessage	"response"
//	@failure		500		{object}	stdmodels.StandardError"error"
//
//	@Security		Authorization
//
//	@Router			/jokes/dadjoke [post]
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) (*stdmodels.StandardMessage, error) {

	request, err := shared.GetParams(r, &DadJoke{})
	if err != nil {
		resp := stdmodels.NewFailureResponse(
			"Failed to add Joke",
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
		return nil, resp
	}
	req := *request.(*DadJoke)

	err = h.PostDadJoke(req)
	if err != nil {
		resp := stdmodels.NewFailureResponse(
			"Failed to add Joke",
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
		return nil, resp
	}

	resp := stdmodels.NewBasicResponse()
	return resp, nil
}

// DeleteDadJoke godoc
//
//	@Summary		Generate dadjoke
//	@Description	create a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Param			joke	query		string						true	"Dad joke you wish to delete from the database"
//	@Success		200		{object}	stdmodels.StandardMessage	"response"
//	@failure		500		{object}	stdmodels.StandardError"error"
//
//	@Security		Authorization
//
//	@Router			/jokes/dadjoke [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) (*stdmodels.StandardMessage, error) {

	request, err := shared.GetParams(r, &DadJoke{})
	if err != nil {
		resp := stdmodels.NewFailureResponse(
			"Failed to delete Joke",
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
		return nil, resp
	}
	req := *request.(*DadJoke)

	err = h.DeleteDadJoke(req)
	if err != nil {
		resp := stdmodels.NewFailureResponse(
			"Failed to delete Joke",
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
		return nil, resp
	}

	resp := stdmodels.NewBasicResponse()
	return resp, nil
}