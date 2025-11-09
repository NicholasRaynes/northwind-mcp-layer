package mcp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type SchemaResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tools       []MCPTool `json:"tools"`
}

// Returns the MCP schema including available tools.
func GetSchema(c *gin.Context) {
	c.JSON(http.StatusOK, SchemaResponse{
		Name:        "northwind-mcp-layer",
		Description: "MCP layer bridging GitHub Copilot to the live Azure-hosted Northwind API. Allows natural-language queries to execute real-time REST API calls.",
		Tools:       GetAvailableTools(),
	})
}

// Handles an MCP request and routes it to the correct API endpoint.
func RunTool(c *gin.Context) {
	loadEnvironment()

	apiBaseURL := os.Getenv("NORTHWIND_API")
	if apiBaseURL == "" {
		sendError(c, http.StatusInternalServerError, "NORTHWIND_API not configured")
		return
	}

	requestData, err := parseRequestBody(c)
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	selectedTool := findTool(requestData.Tool)
	if selectedTool == nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("Unknown tool: %s", requestData.Tool))
		return
	}

	requestURL := buildRequestURL(apiBaseURL, selectedTool.Endpoint, requestData.Params)
	responseData, err := callAPI(requestURL)
	if err != nil {
		sendError(c, http.StatusBadGateway, fmt.Sprintf("REST API call failed: %v", err))
		return
	}

	c.JSON(http.StatusOK, responseData)
}

// Loads .env if present.
// Will remove once validating Github Copilot can use mcp.json successfully.
func loadEnvironment() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: No .env file found")
	}
}

// Extracts tool name and params from JSON.
func parseRequestBody(c *gin.Context) (*struct {
	Tool   string         `json:"tool"`
	Params map[string]any `json:"params"`
}, error) {
	var req struct {
		Tool   string         `json:"tool"`
		Params map[string]any `json:"params"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, fmt.Errorf("invalid request body")
	}
	if req.Tool == "" {
		return nil, fmt.Errorf("missing tool name")
	}
	return &req, nil
}

// Returns the MCPTool matching the requested name.
func findTool(toolName string) *MCPTool {
	for _, tool := range GetAvailableTools() {
		if tool.Name == toolName {
			return &tool
		}
	}
	return nil
}

// Builds a full API request URL with query parameters (if any).
func buildRequestURL(baseURL, endpoint string, params map[string]any) string {
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
		queryString += fmt.Sprintf("%s=%v", paramName, paramValue)
	}
	return requestURL + queryString
}

// Executes the HTTP GET and parses JSON response.
func callAPI(url string) (any, error) {
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

// Sends consistent JSON error responses.
func sendError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"Error": message})
}
