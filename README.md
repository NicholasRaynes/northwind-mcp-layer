# 🧠 northwind-mcp-layer
A lightweight Go-based Model Context Protocol (MCP) server that bridges GitHub Copilot Chat with the live `northwind-api`.
It exposes all of the Northwind business endpoints as structured MCP tools, enabling natural-language querying of real-time data.

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

