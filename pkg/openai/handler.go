package openai

import (
	"gitlab.com/DeveloperDurp/DurpAPI/pkg/shared"
	"gitlab.com/developerdurp/durpify/handlers"
	"net/http"
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
//	@Param			message	query		string						true	"Ask ChatGPT a general question"
//	@Success		200		{object}	handlers.StandardMessage	"response"
//	@failure		500		{object}	handlers.StandardError"error"
//
//	@Security		Authorization
//
//	@Router			/openai/general [get]
func (h *Handler) GeneralOpenAI(w http.ResponseWriter, r *http.Request) (*handlers.StandardMessage, error) {

	request, err := shared.GetParams(r, &ChatRequest{})
	if err != nil {
		resp := handlers.NewFailureResponse(
			"Failed to send message",
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
		return nil, resp
	}
	req := *request.(*ChatRequest)

	result, err := h.createChatCompletion(req.Message, "mistral:instruct")
	if err != nil {
		resp := handlers.NewFailureResponse(
			"Failed to send message",
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
		return nil, resp
	}

	resp := handlers.NewMessageResponse(result, http.StatusOK)
	return resp, nil
}

// TravelAgentOpenAI godoc
//
//	@Summary		Travel Agent ChatGPT
//	@Description	Ask ChatGPT for suggestions as if it was a travel agent
//	@Tags			openai
//	@Accept			json
//	@Produce		application/json
//	@Param			message	query		string						true	"Ask ChatGPT for suggestions as a travel agent"
//	@Success		200		{object}	handlers.StandardMessage	"response"
//	@failure		500		{object}	handlers.StandardError"error"
//
//	@Security		Authorization
//
//	@Router			/openai/travelagent [get]
func (h *Handler) TravelAgentOpenAI(w http.ResponseWriter, r *http.Request) (*handlers.StandardMessage, error) {

	request, err := shared.GetParams(r, &ChatRequest{})
	if err != nil {
		resp := handlers.NewFailureResponse(
			"Failed to send message",
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
		return nil, resp
	}
	req := *request.(*ChatRequest)

	req.Message = "I want you to act as a travel guide. I will give you my location and you will give me suggestions. " + req.Message

	result, err := h.createChatCompletion(req.Message, "mistral:instruct")
	if err != nil {
		resp := handlers.NewFailureResponse(
			"Failed to send message",
			http.StatusInternalServerError,
			[]string{err.Error()},
		)
		return nil, resp
	}

	resp := handlers.NewMessageResponse(result, http.StatusOK)
	return resp, nil
}
