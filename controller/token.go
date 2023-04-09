package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// generateTokenHandler godoc
//
//	@Summary		Generate JWT Token
//	@Description	Gets the PSU Data from unraid
//	@Tags			token
//	@Accept			json
//	@Produce		plain
//	@Param			token	query		string		true	"Secret Token"
//	@Success		200	{string}	string	"response"
//	@Router			/token/generateTokenHandler [get]
func (c *Controller) GenerateTokenHandler(ctx *gin.Context) {
	// Define the token claims
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
		// Add any additional claims here...
	}

	// Generate a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	// TODO: Replace "my-secret-key" with your own secret key
	tokenString, err := token.SignedString([]byte(ctx.Query("token")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	// Return the token in the response body
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
