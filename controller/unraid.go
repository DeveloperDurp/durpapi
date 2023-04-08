package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// UnraidPowerUsage godoc
//
//	@Summary		Unraid PSU Stats
//	@Description	Gets the PSU Data from unraid
//	@Tags			unraid
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"response"
//	@Router			/unraid/powerusage [get]
func (c *Controller) UnraidPowerUsage(ctx *gin.Context) {

	// Create a cookie jar to hold cookies for the session
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create an HTTP client with the cookie jar
	client := &http.Client{
		Jar: jar,
	}

	form := url.Values{
		"username": {"root"},
		"password": {c.unraidAPIKey},
	}

	// Login to unraid
	req, err := http.NewRequest("POST", "https://"+c.unraidURI+"/login", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Check if the login was successful by inspecting the response body or headers
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Login failed!")
		return
	}

	// Now you can use the client to send authenticated requests to other endpoints
	req, err = http.NewRequest("GET", "https://"+c.unraidURI+"/plugins/corsairpsu/status.php", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	// Convert the returned data to JSON
	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, responseJSON)
}
