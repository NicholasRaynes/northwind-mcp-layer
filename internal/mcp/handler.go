package mcp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Builds a full API request URL with query parameters (if any)
func BuildRequestURL(baseURL, endpoint string, params map[string]any) string {
	requestURL := fmt.Sprintf("%s%s", baseURL, endpoint)
	if len(params) == 0 {
		return requestURL
	}

	queryString := "?"
	firstParam := true
	for paramName, paramValue := range params {
		if !firstParam {
			queryString += "&"
		}
		firstParam = false

		queryString += fmt.Sprintf(
			"%s=%s",
			url.QueryEscape(paramName),
			url.QueryEscape(fmt.Sprintf("%v", paramValue)),
		)
	}
	return requestURL + queryString
}

// Executes the HTTP GET and parses JSON response
func CallAPI(url string) (any, error) {
	log.Printf("Calling API: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response")
	}

	var parsed any
	if err := json.Unmarshal(body, &parsed); err != nil {
		parsed = string(body) // fallback to raw text
	}
	return parsed, nil
}
