# 🧠 northwind-mcp-layer
A lightweight Go-based Model Context Protocol (MCP) server, designed to run locally in VS Code, that bridges GitHub Copilot Chat with the live `northwind-api`. 
It exposes all of the Northwind business endpoints as structured MCP tools, enabling natural-language querying of real-time data.

## Usage
1. **Open in VS Code**  
   Open the `northwind-mcp-layer` project in VS Code.

2. **Start the local MCP server**  
   Launch the MCP server using the configuration in `.vscode/mcp.json`.  
   This automatically starts the Gin HTTP server and exposes the MCP tools.

3. **Configure GitHub Copilot Chat**  
   In GitHub Copilot Chat, load the MCP tools by pointing to the local server defined in `mcp.json`.

4. **Query Northwind data**  
   Use natural-language prompts in Copilot Chat to interact with live Northwind endpoints, for example:  
   - `Provide details on customers from Germany.`  
   - `What's the inventory status of products that need reordering?`

## Notes
- The `northwind-api` backend is hosted as a free-tier Render web service. It may take a few seconds to start up when initially accessed.
- In my experience, Claude models tend to work best when interacting with this MCP layer.

## Architecture Overview
```text
northwind-mcp-layer/
├── cmd/
│   └── mcp/               # Application entrypoint (starts the MCP Gin HTTP server)
├── internal/
│   └── mcp/               # Core MCP logic (request handling, tool routing)
├── go.mod / go.sum        # Go module dependencies
├── .vscode/
│   └── mcp.json           # MCP server configuration for Github Copilot (auto-loads server, tools, and API base)
└── README.md              # Project documentation
```

## Example Tools
| Tool Name                  | Description                                    | Example Parameters                              |
| -------------------------- | ---------------------------------------------- | ----------------------------------------------- |
| `get_customers`            | Retrieve customers from the Northwind API      | `country=Germany`, `city=London`                |
| `get_orders`               | Retrieve orders by customer, employee, or year | `year=1998`, `customer_id=ALFKI`                |
| `get_sales_by_country`     | View total sales by country                    | `year=1998`                                     |
| `get_top_customers`        | View top customers by revenue                  | `limit=10`, `country=USA`                       |

Each MCP tool corresponds to a live endpoint on the `northwind-api` and is dynamically available to Github Copilot via the MCP protocol.

## Example Prompts
- `Provide some details on our customers from Germany.`
- `What are our shipping costs for the year 1997?`
- `What's the inventory status of our products that need to be reordered?`
- `What's the supplier performance of Pavlova, Ltd.?`

## Tech Stack
- Language: Go 1.23+
- Framework: Gin
- Protocol: Model Context Protocol (MCP)
- Integration: GitHub Copilot Chat

## Companion Project
https://github.com/nicholasraynes/northwind-api
- The REST API backend that provides the live data for this MCP layer. Implements the Northwind dataset with analytical extensions (orders, products, suppliers, KPIs).

