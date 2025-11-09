package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nicholasraynes/northwind-mcp-layer/internal/mcp"
)

func main() {
	// Will remove once validating Github Copilot can use mcp.json successfully.
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
		log.Println("ℹWarning: No configured port found, defaulting to port 8082")
	}

	r := gin.Default()
	r.GET("/mcp/schema", mcp.GetSchema)
	r.POST("/mcp/run", mcp.RunTool)

	log.Printf("Success: Northwind MCP Layer running at http://localhost:%s\n", port)

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Error: Failed to start server: %v", err)
	}
}
