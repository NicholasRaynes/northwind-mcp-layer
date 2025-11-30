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
			Description: "Retrieve all customer records. Filterable by: customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, or fax.",
			Endpoint:    "/customers",
			Method:      "GET",
			Params: []MCPParam{
				{"country", "string", "Filter customers by country."},
				{"city", "string", "Filter customers by city."},
				{"customer_id", "string", "Filter by customer ID."},
				{"company_name", "string", "Filter by customer company name."},
				{"contact_name", "string", "Filter by customer contact name."},
				{"contact_title", "string", "Filter by customer contact title."},
				{"address", "string", "Filter by customer address."},
				{"region", "string", "Filter by customer region."},
				{"postal_code", "string", "Filter by customer postal code."},
				{"phone", "string", "Filter by customer phone number."},
				{"fax", "string", "Filter by customer fax number."},
			},
		},
		{
			Name:        "get_orders",
			Description: "Retrieve all orders, filterable by: order_id, customer_id, customer_name, employee, employee_name, year, order_date, required_date, shipped_date, ship_via, shipper_name, ship_name, ship_address, ship_city, ship_region, ship_postal_code, or country (ship country).",
			Endpoint:    "/orders",
			Method:      "GET",
			Params: []MCPParam{
				{"customer_id", "string", "Filter orders by customer ID."},
				{"employee", "string", "Filter orders by employee ID."},
				{"year", "integer", "Filter orders by order year."},
				{"country", "string", "Filter orders by ship country."},
				{"order_id", "string", "Filter orders by order ID."},
				{"customer_name", "string", "Filter orders by customer name."},
				{"employee_name", "string", "Filter orders by employee name."},
				{"order_date", "string", "Filter orders by exact order date."},
				{"required_date", "string", "Filter orders by required date."},
				{"shipped_date", "string", "Filter orders by shipped date."},
				{"ship_via", "string", "Filter orders by shipper ID."},
				{"shipper_name", "string", "Filter orders by shipper name."},
				{"ship_name", "string", "Filter orders by ship name."},
				{"ship_address", "string", "Filter orders by ship address."},
				{"ship_city", "string", "Filter orders by ship city."},
				{"ship_region", "string", "Filter orders by ship region."},
				{"ship_postal_code", "string", "Filter orders by ship postal code."},
			},
		},
		{
			Name:        "get_order_details",
			Description: "Retrieve product-level order details (price, quantity, discount). Filterable by: order_id, customer_id, product_id, product_name, category_name, or supplier_name.",
			Endpoint:    "/orders/details",
			Method:      "GET",
			Params: []MCPParam{
				{"order_id", "integer", "Filter by the exact order ID."},
				{"customer_id", "string", "Filter by customer ID."},
				{"product_id", "string", "Filter by product ID."},
				{"product_name", "string", "Filter by product name."},
				{"category_name", "string", "Filter by category name."},
				{"supplier_name", "string", "Filter by supplier name."},
			},
		},
		{
			Name:        "get_products",
			Description: "Retrieve product details (price, stock, supplier, category). Filterable by: product_id, product_name, supplier_id, supplier_name, category_id, category_name, or discontinued status.",
			Endpoint:    "/products",
			Method:      "GET",
			Params: []MCPParam{
				{"product_id", "string", "Filter by product ID."},
				{"product_name", "string", "Filter by product name."},
				{"supplier_id", "string", "Filter by supplier ID."},
				{"supplier_name", "string", "Filter by supplier name."},
				{"category_id", "string", "Filter by category ID."},
				{"category_name", "string", "Filter by category name."},
				{"discontinued", "string", "Filter by discontinued status ('true' or 'false')."},
			},
		},
		{
			Name:        "get_suppliers",
			Description: "Retrieve all supplier information. Filterable by: supplier_id, company_name, contact_name, contact_title, city, country, phone, or fax.",
			Endpoint:    "/suppliers",
			Method:      "GET",
			Params: []MCPParam{
				{"country", "string", "Filter by supplier country."},
				{"supplier_id", "string", "Filter by supplier ID."},
				{"company_name", "string", "Filter by company name."},
				{"contact_name", "string", "Filter by contact name."},
				{"contact_title", "string", "Filter by contact title."},
				{"city", "string", "Filter by city."},
				{"phone", "string", "Filter by phone number."},
				{"fax", "string", "Filter by fax number."},
			},
		},
		{
			Name:        "get_sales_by_country",
			Description: "Aggregate total sales grouped by customer country. Filterable by: year or country.",
			Endpoint:    "/summary/sales-by-country",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by order year."},
				{"country", "string", "Filter by a specific country."},
			},
		},
		{
			Name:        "get_sales_by_category",
			Description: "Total revenue aggregated by product category. Filterable by: year or category_name.",
			Endpoint:    "/summary/sales-by-category",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by order year."},
				{"category_name", "string", "Filter by category name."},
			},
		},
		{
			Name:        "get_sales_by_employee",
			Description: "Total sales and order counts grouped by employee. Filterable by: year or employee_name.",
			Endpoint:    "/summary/sales-by-employee",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by order year."},
				{"employee_name", "string", "Filter by employee name."},
			},
		},
		{
			Name:        "get_sales_by_year",
			Description: "Annual sales totals across all countries and categories. Not filterable.",
			Endpoint:    "/summary/sales-by-year",
			Method:      "GET",
		},
		{
			Name:        "get_sales_by_shipper",
			Description: "Total sales grouped by freight company. Filterable by: year or company_name (shipper name).",
			Endpoint:    "/summary/sales-by-shipper",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by order year."},
				{"company_name", "string", "Filter by shipper company name."},
			},
		},
		{
			Name:        "get_top_customers",
			Description: "Retrieve top customers ranked by total revenue. Filterable by: year, country, customer_id, or company_name.",
			Endpoint:    "/analytics/top-customers",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter results by order year."},
				{"country", "string", "Filter by customer country."},
				{"customer_id", "string", "Filter by customer ID."},
				{"company_name", "string", "Filter by customer company name."},
			},
		},
		{
			Name:        "get_customer_orders",
			Description: "Detailed order history per customer. Filterable by: customer_id, year, order_id, company_name, order_date, shipped_date, or country.",
			Endpoint:    "/analytics/customer-orders",
			Method:      "GET",
			Params: []MCPParam{
				{"customer_id", "string", "Filter by customer ID."},
				{"year", "integer", "Filter by order year."},
				{"order_id", "string", "Filter by order ID."},
				{"company_name", "string", "Filter by customer company name."},
				{"order_date", "string", "Filter by order date."},
				{"shipped_date", "string", "Filter by shipped date."},
				{"country", "string", "Filter by ship country."},
			},
		},
		{
			Name:        "get_customer_ltv",
			Description: "Total lifetime revenue (LTV) per customer. Filterable by: country, customer_id, or company_name.",
			Endpoint:    "/analytics/customer-ltv",
			Method:      "GET",
			Params: []MCPParam{
				{"country", "string", "Filter by customer country."},
				{"customer_id", "string", "Filter by customer ID."},
				{"company_name", "string", "Filter by customer company name."},
			},
		},
		{
			Name:        "get_customer_retention",
			Description: "Measures repeat customers and retention rates. Filterable by: year, customer_id, company_name, country, or repeat_customer status ('true'/'false').",
			Endpoint:    "/analytics/customer-retention",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by year."},
				{"customer_id", "string", "Filter by customer ID."},
				{"company_name", "string", "Filter by customer company name."},
				{"country", "string", "Filter by country."},
				{"repeat_customer", "string", "Filter by repeat customer status ('true' or 'false')."},
			},
		},
		{
			Name:        "get_top_products",
			Description: "Top-selling products by total revenue and units sold. Filterable by: year, product_id, product_name, category_name, or supplier_name.",
			Endpoint:    "/analytics/top-products",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter results by order year."},
				{"product_id", "string", "Filter by product ID."},
				{"product_name", "string", "Filter by product name."},
				{"category_name", "string", "Filter by product category name."},
				{"supplier_name", "string", "Filter by supplier name."},
			},
		},
		{
			Name:        "get_supplier_performance",
			Description: "Supplier contribution by revenue and efficiency. Filterable by: year, supplier_id, supplier_name, country, or top_category.",
			Endpoint:    "/analytics/supplier-performance",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by order year."},
				{"supplier_id", "string", "Filter by supplier ID."},
				{"supplier_name", "string", "Filter by supplier name."},
				{"country", "string", "Filter by supplier country."},
				{"top_category", "string", "Filter by top category."},
			},
		},
		{
			Name:        "get_inventory_status",
			Description: "Product stock levels, reorder needs, and discontinued items. Filterable by: product_id, product_name, supplier_name, category_name, discontinued status, or needs_reorder status.",
			Endpoint:    "/analytics/inventory-status",
			Method:      "GET",
			Params: []MCPParam{
				{"product_id", "string", "Filter by product ID."},
				{"product_name", "string", "Filter by product name."},
				{"supplier_name", "string", "Filter by supplier name."},
				{"category_name", "string", "Filter by category name."},
				{"discontinued", "string", "Filter by discontinued status ('true' or 'false')."},
				{"needs_reorder", "string", "Filter by reorder status ('true' or 'false')."},
			},
		},
		{
			Name:        "get_employee_performance",
			Description: "Rank employees by sales totals and order count. Filterable by: year, employee_id, full_name, title, or country.",
			Endpoint:    "/analytics/employee-performance",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by order year."},
				{"employee_id", "string", "Filter by employee ID."},
				{"full_name", "string", "Filter by employee full name."},
				{"title", "string", "Filter by employee job title."},
				{"country", "string", "Filter by employee country."},
			},
		},
		{
			Name:        "get_shipping_costs",
			Description: "Freight cost and order volume per shipping company. Filterable by: year, shipper_id, or company_name.",
			Endpoint:    "/analytics/shipping-costs",
			Method:      "GET",
			Params: []MCPParam{
				{"year", "integer", "Filter by order year."},
				{"shipper_id", "string", "Filter by shipper ID."},
				{"company_name", "string", "Filter by shipper company name."},
			},
		},
	}
}
