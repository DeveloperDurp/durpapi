package openai

import (
	"encoding/json"
	"net/http"

	"gitlab.com/developerdurp/stdmodels"
)

type Handler struct {
	LlamaURL string
}

func NewHandler(LlamaURL string) (*Handler, error) {
	return &Handler{LlamaURL: LlamaURL}, nil
}

type ChatRequest struct {
	Message string `json:"message"`
}

// Response struct to unmarshal the JSON response
type Response struct {
	Response string `json:"response"`
}

// GeneralOpenAI godoc
//
//	@Summary		Gerneral ChatGPT
//	@Description	Ask ChatGPT a general question
//	@Tags			openai
//	@Accept			json
//	@Produce		application/json
//	@Param			message	query		string			true	"Ask ChatGPT a general question"
//	@Success		200	{object}	stdmodels.StandardMessage	"response"
//	@failure		500	{object}	stdmodels.StandardError"error"
//
// @Security Authorization
//
//	@Router			/openai/general [get]
func (h *Handler) GeneralOpenAI(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var req ChatRequest

	if contentType == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			stdmodels.FailureReponse("Failed To decode content", w, http.StatusInternalServerError, []string{err.Error()})
			return
		}
	} else {
		queryParams := r.URL.Query()
		req.Message = queryParams.Get("message")
	}

	result, err := h.createChatCompletion(req.Message, "mistral:instruct")
	if err != nil {
		stdmodels.FailureReponse("Failed to Send Message", w, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	stdmodels.SuccessResponse(result, w, http.StatusOK)
}

// TravelAgentOpenAI godoc
//
//	@Summary		Travel Agent ChatGPT
//	@Description	Ask ChatGPT for suggestions as if it was a travel agent
//	@Tags			openai
//	@Accept			json
//	@Produce		application/json
//	@Param			message	query		string			true	"Ask ChatGPT for suggestions as a travel agent"
//	@Success		200	{object}	stdmodels.StandardMessage	"response"
//	@failure		500	{object}	stdmodels.StandardError"error"
//
// @Security Authorization
//
//	@Router			/openai/travelagent [get]
func (h *Handler) TravelAgentOpenAI(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	var req ChatRequest

	if contentType == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			stdmodels.FailureReponse("Failed To decode content", w, http.StatusInternalServerError, []string{err.Error()})
			return
		}
	} else {
		queryParams := r.URL.Query()
		req.Message = queryParams.Get("message")
	}

	req.Message = "I want you to act as a travel guide. I will give you my location and you will give me suggestions. " + req.Message

	result, err := h.createChatCompletion(req.Message, "mistral:instruct")
	if err != nil {
		stdmodels.FailureReponse("Failed to Send Message", w, http.StatusInternalServerError, []string{err.Error()})
		return
	}

	stdmodels.SuccessResponse(result, w, http.StatusOK)
}
