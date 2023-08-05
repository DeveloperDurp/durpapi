package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

type ChatRequest struct {
	Message string `json:"message"`
}

// GeneralOpenAI godoc
//
//	@Summary		Gerneral ChatGPT
//	@Description	Ask ChatGPT a general question
//	@Tags			openai
//	@Accept			json
//	@Produce		application/json
//	@Param			message	query		string			true	"Ask ChatGPT a general question"
//	@Success		200		{object}	model.Message	"response"
//
//	@failure		400		{object}	model.Message	"error"
//
//	@Router			/openai/general [get]
func (c *Controller) GeneralOpenAI(ctx *gin.Context) {
	var req ChatRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		req.Message = ctx.Query("message")
	}

	result, err := c.createChatCompletion(req.Message)
	if err != nil {
		err := ctx.AbortWithError(http.StatusInternalServerError, err)
		if err != nil {
			fmt.Println("Failed to send message")
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result})
}

// TravelAgentOpenAI godoc
//
//	@Summary		Travel Agent ChatGPT
//	@Description	Ask ChatGPT for suggestions as if it was a travel agent
//	@Tags			openai
//	@Accept			json
//	@Produce		application/json
//	@Param			message	query		string			true	"Ask ChatGPT for suggestions as a travel agent"
//	@Success		200		{object}	model.Message	"response"
//	@failure		400		{object}	model.Message	"error"
//	@Router			/openai/travelagent [get]
func (c *Controller) TravelAgentOpenAI(ctx *gin.Context) {
	var req ChatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		req.Message = ctx.Query("message")
	}

	req.Message = "I want you to act as a travel guide. I will give you my location and you will give me suggestions. " + req.Message

	result, err := c.createChatCompletion(req.Message)
	if err != nil {
		err := ctx.AbortWithError(http.StatusInternalServerError, err)
		if err != nil {
			fmt.Println("Failed to send message")
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result})
}

func (c *Controller) createChatCompletion(message string) (string, error) {
	client := c.Cfg.OpenaiClient
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
