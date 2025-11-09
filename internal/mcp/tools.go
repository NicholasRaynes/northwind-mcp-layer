package mcp

type MCPParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type MCPTool struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Endpoint    string     `json:"endpoint"`
	Method      string     `json:"method"`
	Params      []MCPParam `json:"params,omitempty"`
}

// Returns the list of available MCP tools based on northwind-api endpoints.
func GetAvailableTools() []MCPTool {
	return []MCPTool{
		{
			Name:        "get_customers",
			Description: "Retrieve all customers, including company names, cities, and contact details.",
			Endpoint:    "/customers",
			Method:      "GET",
			Params: []MCPParam{
				{"country", "string", "Filter by country"},
				{"city", "string", "Filter by city"},
				{"id", "string", "Filter by customer ID"},
			},
		},
		{
			Name:        "get_orders",
			Description: "Retrieve all orders with customer, employee, shipper, and date information.",
			Endpoint:    "/orders",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by order year"},
				{"customer_id", "string", "Filter by customer ID"},
				{"employee", "string", "Filter by employee"},
				{"country", "string", "Filter by customer country"},
			},
		},
		{
			Name:        "get_order_details",
			Description: "Retrieve product-level order details including quantity, price, and discounts.",
			Endpoint:    "/orders/details",
			Method:      "GET",
			Params: []MCPParam{
				{"order_id", "integer", "Filter by order ID"},
				{"product", "string", "Filter by product name"},
				{"customer_id", "string", "Filter by customer ID"},
			},
		},
		{
			Name:        "get_products",
			Description: "Retrieve products with supplier, category, price, and stock information.",
			Endpoint:    "/products",
			Method:      "GET",
			Params: []MCPParam{
				{"category", "string", "Filter by category"},
				{"supplier", "string", "Filter by supplier"},
				{"name", "string", "Filter by product name"},
			},
		},
		{
			Name:        "get_suppliers",
			Description: "Retrieve all supplier information, including contact details and country.",
			Endpoint:    "/suppliers",
			Method:      "GET",
			Params: []MCPParam{
				{"country", "string", "Filter by supplier country"},
				{"company", "string", "Filter by company name"},
			},
		},
		{
			Name:        "get_sales_by_country",
			Description: "Aggregate total sales grouped by customer country.",
			Endpoint:    "/summary/sales-by-country",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_sales_by_category",
			Description: "Retrieve total revenue aggregated by product category.",
			Endpoint:    "/summary/sales-by-category",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_sales_by_employee",
			Description: "Retrieve total sales and order counts grouped by employee.",
			Endpoint:    "/summary/sales-by-employee",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_sales_by_year",
			Description: "Retrieve annual sales totals across all countries and categories.",
			Endpoint:    "/summary/sales-by-year",
			Method:      "GET",
		},
		{
			Name:        "get_sales_by_shipper",
			Description: "Retrieve total sales grouped by freight company.",
			Endpoint:    "/summary/sales-by-shipper",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_top_customers",
			Description: "Retrieve top customers ranked by total revenue.",
			Endpoint:    "/analytics/top-customers",
			Method:      "GET",
			Params: []MCPParam{
				{"limit", "integer", "Limit number of top results"},
				{"country", "string", "Filter by country"},
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_customer_orders",
			Description: "Retrieve detailed order history per customer.",
			Endpoint:    "/analytics/customer-orders",
			Method:      "GET",
			Params: []MCPParam{
				{"customer_id", "string", "Filter by customer ID"},
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_customer_ltv",
			Description: "Retrieve total lifetime revenue per customer.",
			Endpoint:    "/analytics/customer-ltv",
			Method:      "GET",
			Params: []MCPParam{
				{"limit", "integer", "Limit number of top customers"},
				{"country", "string", "Filter by country"},
			},
		},
		{
			Name:        "get_customer_retention",
			Description: "Retrieve customer retention metrics such as repeat rates and retention percentage.",
			Endpoint:    "/analytics/customer-retention",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_top_products",
			Description: "Retrieve top-selling products by total revenue and units sold.",
			Endpoint:    "/analytics/top-products",
			Method:      "GET",
			Params: []MCPParam{
				{"limit", "integer", "Limit number of top products"},
				{"category", "string", "Filter by category"},
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_supplier_performance",
			Description: "Retrieve supplier contribution by revenue and efficiency.",
			Endpoint:    "/analytics/supplier-performance",
			Method:      "GET",
			Params: []MCPParam{
				{"limit", "integer", "Limit number of suppliers"},
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_inventory_status",
			Description: "Retrieve product stock levels, reorder needs, and discontinued items.",
			Endpoint:    "/analytics/inventory-status",
			Method:      "GET",
			Params: []MCPParam{
				{"supplier", "string", "Filter by supplier"},
				{"category", "string", "Filter by category"},
			},
		},
		{
			Name:        "get_employee_performance",
			Description: "Retrieve employees ranked by sales totals and order count.",
			Endpoint:    "/analytics/employee-performance",
			Method:      "GET",
			Params: []MCPParam{
				{"limit", "integer", "Limit number of employees"},
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_shipping_costs",
			Description: "Retrieve freight cost and order volume per shipping company.",
			Endpoint:    "/analytics/shipping-costs",
			Method:      "GET",
			Params: []MCPParam{
				{"limit", "integer", "Limit number of shipping records"},
				{"year", "integer", "Filter by year"},
			},
		},
		{
			Name:        "get_delivery_times",
			Description: "Retrieve average delivery days, late shipments, and on-time rate by shipper or employee.",
			Endpoint:    "/analytics/delivery-times",
			Method:      "GET",
			Params: []MCPParam{
				{"group", "string", "Filter or group results by shipper or employee"},
				{"year", "integer", "Filter by year"},
			},
		},
	}
}
