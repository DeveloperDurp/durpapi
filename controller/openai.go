package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

var client = openai.NewClient("")

// GeneralOpenAI godoc
//
//	@Summary		ping example
//	@Description	do ping
//	@Tags			openai
//	@Accept			json
//	@Produce		plain
//	@Param			message	query		string		true	"Ask ChatGPT a general question"
//	@Success		200	{string}	string	"response"
//	@Failure		400	{string}	string	"ok"
//	@Failure		404	{string}	string	"ok"
//	@Failure		500	{string}	string	"ok"
//	@Router			/openai/GeneralOpenAI [get]
func (c *Controller) GeneralOpenAI(ctx *gin.Context) {
	message := ctx.Query("message")

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
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.String(http.StatusOK, resp.Choices[0].Message.Content)
}
