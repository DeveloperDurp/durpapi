package dadjoke

import (
	"encoding/json"
	"net/http"

	"gitlab.com/developerdurp/logger"
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
//	@Success		200	{object}	stdmodels.StandardMessage	"response"
//	@failure		500	{object}	stdmodels.StandardError"error"
//
// @Security Authorization
//
//	@Router			/jokes/dadjoke [get]
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	joke, err := h.GetRandomDadJoke()

	if err != nil {
		stdmodels.FailureReponse("Failed to get Joke", w, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	message := stdmodels.StandardMessage{
		Message: joke,
	}

	json.NewEncoder(w).Encode(message)
}

// PostDadJoke godoc
//
//	@Summary		Generate dadjoke
//	@Description	create a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Param			joke	query		string			true	"Dad Joke you wish to enter into database"
//	@Success		200		{object}	stdmodels.StandardMessage	"response"
//	@failure		500	{object}	stdmodels.StandardError"error"
//
// @Security Authorization
//
//	@Router			/jokes/dadjoke [post]
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var req DadJoke

	if contentType == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.LogError("Failed to decode json file")
			return
		}
	} else {
		queryParams := r.URL.Query()
		req.JOKE = queryParams.Get("joke")
	}

	err := h.PostDadJoke(req)
	if err != nil {
		stdmodels.FailureReponse("Failed to add joke", w, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	stdmodels.SuccessResponse("OK", w, http.StatusOK)
}

// DeleteDadJoke godoc
//
//	@Summary		Generate dadjoke
//	@Description	create a dad joke
//	@Tags			DadJoke
//	@Accept			json
//	@Produce		application/json
//	@Param			joke	query		string			true	"Dad joke you wish to delete from the database"
//	@Success		200		{object}	stdmodels.StandardMessage	"response"
//	@failure		500	{object}	stdmodels.StandardError"error"
//
// @Security Authorization
//
//	@Router			/jokes/dadjoke [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var req DadJoke

	if contentType == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
	} else {
		queryParams := r.URL.Query()
		req.JOKE = queryParams.Get("joke")
	}

	err := h.DeleteDadJoke(req)
	if err != nil {
		stdmodels.FailureReponse("Failed to delete joke", w, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	stdmodels.SuccessResponse("OK", w, http.StatusOK)
}
