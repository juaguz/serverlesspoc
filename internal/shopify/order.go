package shopify

import "encoding/json"

type ShopifyDiscountApplications struct {
	Type             string `json:"type"`
	Value            string `json:"value"`
	ValueType        string `json:"value_type"`
	AllocationMethod string `json:"allocation_method"`
	TargetSelection  string `json:"target_selection"`
	TargetType       string `json:"target_type"`
	Description      string `json:"description"`
	Title            string `json:"title"`
}

type DiscountCode struct {
	Code   string      `json:"code"`
	Amount json.Number `json:"amount"`
	Type   string      `json:"string"`
}

type ShopifyPriceSet struct {
	ShopMoney        Money `json:"shop_money"`
	PresentmentMoney Money `json:"presentment_money"`
}

type Money struct {
	Amount       json.Number `json:"amount"`
	CurrencyCode string      `json:"currency_code"`
}

type WebHookLineItem struct {
	ID                         int64           `json:"id"`
	VariantID                  int64           `json:"variant_id"`
	Title                      string          `json:"title"`
	Quantity                   int             `json:"quantity"`
	Sku                        string          `json:"sku"`
	VariantTitle               string          `json:"variant_title"`
	Vendor                     string          `json:"vendor"`
	FulfillmentService         string          `json:"fulfillment_service"`
	ProductID                  int64           `json:"product_id"`
	RequiresShipping           bool            `json:"requires_shipping"`
	Taxable                    bool            `json:"taxable"`
	GiftCard                   bool            `json:"gift_card"`
	Name                       string          `json:"name"`
	VariantInventoryManagement string          `json:"variant_inventory_management"`
	Properties                 []interface{}   `json:"properties"`
	ProductExists              bool            `json:"product_exists"`
	FulfillableQuantity        int             `json:"fulfillable_quantity"`
	Grams                      int             `json:"grams"`
	Price                      json.Number     `json:"price"`
	TotalDiscount              string          `json:"total_discount"`
	FulfillmentStatus          string          `json:"fulfillment_status"`
	PriceSet                   ShopifyPriceSet `json:"price_set"`
	TotalDiscountSet           ShopifyPriceSet `json:"total_discount_set"`
	DiscountAllocations        []interface{}   `json:"discount_allocations"`
	Duties                     []interface{}   `json:"duties"`
	AdminGraphqlAPIID          string          `json:"admin_graphql_api_id"`
	TaxLines                   []interface{}   `json:"tax_lines"`
}

type ShopifyShippingLines struct {
	ID                            int64           `json:"id"`
	Title                         string          `json:"title"`
	Price                         string          `json:"price"`
	Code                          interface{}     `json:"code"`
	Source                        string          `json:"source"`
	Phone                         interface{}     `json:"phone"`
	RequestedFulfillmentServiceID interface{}     `json:"requested_fulfillment_service_id"`
	DeliveryCategory              interface{}     `json:"delivery_category"`
	CarrierIdentifier             interface{}     `json:"carrier_identifier"`
	DiscountedPrice               string          `json:"discounted_price"`
	PriceSet                      ShopifyPriceSet `json:"price_set"`
	DiscountedPriceSet            ShopifyPriceSet `json:"discounted_price_set"`
	DiscountAllocations           []interface{}   `json:"discount_allocations"`
	TaxLines                      []interface{}   `json:"tax_lines"`
}

type ShopifyAddress struct {
	FirstName    string  `json:"first_name"`
	Address1     string  `json:"address1"`
	Phone        string  `json:"phone"`
	City         string  `json:"city"`
	Zip          string  `json:"zip"`
	Province     string  `json:"province"`
	Country      string  `json:"country"`
	LastName     string  `json:"last_name"`
	Address2     string  `json:"address2"`
	Company      string  `json:"company"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name"`
	CountryCode  string  `json:"country_code"`
	ProvinceCode string  `json:"province_code"`
}

type ShopifyCustomer struct {
	ID                        int64                 `json:"id"`
	Email                     string                `json:"email"`
	AcceptsMarketing          bool                  `json:"accepts_marketing"`
	CreatedAt                 interface{}           `json:"created_at"`
	UpdatedAt                 interface{}           `json:"updated_at"`
	FirstName                 string                `json:"first_name"`
	LastName                  string                `json:"last_name"`
	OrdersCount               int                   `json:"orders_count"`
	State                     string                `json:"state"`
	TotalSpent                string                `json:"total_spent"`
	LastOrderID               interface{}           `json:"last_order_id"`
	Note                      interface{}           `json:"note"`
	VerifiedEmail             bool                  `json:"verified_email"`
	MultipassIdentifier       interface{}           `json:"multipass_identifier"`
	TaxExempt                 bool                  `json:"tax_exempt"`
	Phone                     interface{}           `json:"phone"`
	Tags                      string                `json:"tags"`
	LastOrderName             interface{}           `json:"last_order_name"`
	Currency                  string                `json:"currency"`
	AcceptsMarketingUpdatedAt interface{}           `json:"accepts_marketing_updated_at"`
	MarketingOptInLevel       interface{}           `json:"marketing_opt_in_level"`
	AdminGraphqlAPIID         string                `json:"admin_graphql_api_id"`
	DefaultAddress            ShopifyDefaultAddress `json:"default_address"`
}

type ShopifyDefaultAddress struct {
	ID           int64  `json:"id"`
	CustomerID   int64  `json:"customer_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Company      string `json:"company"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Zip          string `json:"zip"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	ProvinceCode string `json:"province_code"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	Default      bool   `json:"default"`
}
