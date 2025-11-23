package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
	northwind "github.com/nicholasraynes/northwind-mcp-layer/internal/mcp"
)

func main() {
	// Create a context for the server
	ctx := context.Background()

	// Initialize the MCP server with a name and version
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "northwind",
		Version: "1.0.0",
	}, nil)

	// Iterate through all available tools and register them with the MCP server
	for _, toolInfo := range northwind.GetAvailableTools() {

		currentTool := toolInfo

		// Add the tool to the MCP server
		mcp.AddTool(server, &mcp.Tool{
			Name:        currentTool.Name,
			Description: currentTool.Description,
		}, func(ctx context.Context, request *mcp.CallToolRequest, args map[string]any) (*mcp.CallToolResult, any, error) {

			// Get the Northwind API base URL from environment variable
			apiBaseURL := os.Getenv("NORTHWIND_API")
			if apiBaseURL == "" {
				return mcpResponse("Environment variable NORTHWIND_API is not configured"), nil, nil
			}

			// Build the full API request URL using the tool's endpoint and the arguments
			requestURL := northwind.BuildRequestURL(apiBaseURL, currentTool.Endpoint, args)

			// Call the Northwind API
			apiResponseData, apiErr := northwind.CallAPI(requestURL)
			if apiErr != nil {
				// If the API call fails, return the error message to the client
				return mcpResponse(fmt.Sprintf("API call failed: %v", apiErr)), nil, nil
			}

			// Convert the API response to a JSON string (pretty-printed)
			jsonOutput, jsonErr := json.MarshalIndent(apiResponseData, "", "  ")
			if jsonErr != nil {
				return mcpResponse(fmt.Sprintf("Failed to convert API response to JSON: %v", jsonErr)), nil, nil
			}

			// Return the JSON string as a text response to the MCP client
			return mcpResponse(string(jsonOutput)), nil, nil
		})
	}

	log.Println("Northwind MCP server is running...")

	// Start the MCP server using standard input/output transport
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// Helper function that formats a string message into a CallToolResult that MCP can return to the client
func mcpResponse(message string) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: message},
		},
	}
}
