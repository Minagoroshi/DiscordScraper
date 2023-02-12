package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// The checkTag function checks if a username and tag is available on Discord.
// It returns a bool and an error.
// The bool is true if the username and tag is available, false if not.
func checkTag(username string, tag int, auth string) (bool, error) {
	url := "https://discord.com/api/v9/users/@me"
	method := "PATCH"

	payload := map[string]string{
		"username":      username,
		"password":      "",
		"discriminator": strconv.Itoa(tag),
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return false, errors.New("error marshalling payload")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))

	if err != nil {
		return false, errors.New("error creating request")
	}

	for k, v := range discordCommonHeaders(auth) {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return false, errors.New("error sending request")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, errors.New("error reading response body")
	}

	// 400 = good
	// 401 = unauthorized
	// 429 = too many requests

	if res.StatusCode == 401 {
		return false, errors.New("unauthorized")
	}

	if res.StatusCode == 429 {
		log.Println(string(body))
		// Get the retry after time
		var ratelimitResponse RatelimitResponse
		err = json.Unmarshal(body, &ratelimitResponse)
		if err != nil {
			return false, err
		}
		log.Println("Rate limited for: "+strconv.Itoa(int(ratelimitResponse.RetryAfter))+" seconds", "Trying again...")
		// Wait for the retry after time
		time.Sleep(10 * time.Second)
		// Try again
		return checkTag(username, tag, auth)
	}

	if strings.Contains(string(body), "This username and tag are already taken") {
		return false, nil
	}

	return true, nil
}

// The discordCommonHeaders function returns a map of common headers used in Discord requests.
func discordCommonHeaders(authorization string) map[string]string {
	return map[string]string{
		"accept":             "*/*",
		"accept-language":    "en-US",
		"authorization":      authorization,
		"content-type":       "application/json",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC45MDEwIiwib3NfdmVyc2lvbiI6IjEwLjAuMjI2MjEiLCJvc19hcmNoIjoieDY0Iiwic3lzdGVtX2xvY2FsZSI6ImVuLVVTIiwiY2xpZW50X2J1aWxkX251bWJlciI6MTczNTAxLCJuYXRpdmVfYnVpbGRfbnVtYmVyIjoyOTEyOCwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=",
	}
}

type RatelimitResponse struct {
	Global     bool    `json:"global"`
	Message    string  `json:"message"`
	RetryAfter float64 `json:"retry_after"`
}
