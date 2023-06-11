package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/DeveloperDurp/DurpAPI/model"
)

var (
	unraidAPIKey = model.UnraidAPIKey
	UnraidURI    = model.UnraidURI
)

// UnraidPowerUsage godoc
//
//	@Summary		Unraid PSU Stats
//	@Description	Gets the PSU Data from unraid
//	@Tags			unraid
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.PowerSupply	"response"
//	@failure		412 {object}	model.Message	"error"
//	@Router			/unraid/powerusage [get]
func (c *Controller) UnraidPowerUsage(ctx *gin.Context) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Jar: jar,
	}

	form := url.Values{
		"username": {"root"},
		"password": {unraidAPIKey},
	}

	req, err := http.NewRequest("POST", "https://"+UnraidURI+"/login", strings.NewReader(form.Encode()))
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

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Login failed!")
		return
	}

	req, err = http.NewRequest("GET", "https://"+UnraidURI+"/plugins/corsairpsu/status.php", nil)
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

	var responseJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"message": "Bad Response from Unraid"})
		return
	}

	ctx.JSON(http.StatusOK, responseJSON)
}
