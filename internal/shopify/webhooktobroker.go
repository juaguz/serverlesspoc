package shopify

import (
	"encoding/json"
	"log"
)

type WebHookOrder struct {
	ID                     int64                         `json:"id"`
	Email                  string                        `json:"email"`
	ClosedAt               interface{}                   `json:"closed_at"`
	CreatedAt              string                        `json:"created_at"`
	UpdatedAt              string                        `json:"updated_at"`
	Number                 int                           `json:"number"`
	Note                   string                        `json:"note"`
	Token                  string                        `json:"token"`
	Gateway                string                        `json:"gateway"`
	Test                   bool                          `json:"test"`
	TotalPrice             json.Number                   `json:"total_price"`
	SubtotalPrice          json.Number                   `json:"subtotal_price"`
	TotalWeight            int                           `json:"total_weight"`
	TotalTax               string                        `json:"total_tax"`
	TaxesIncluded          bool                          `json:"taxes_included"`
	Currency               string                        `json:"currency"`
	FinancialStatus        string                        `json:"financial_status"`
	Confirmed              bool                          `json:"confirmed"`
	TotalDiscounts         json.Number                   `json:"total_discounts"`
	TotalLineItemsPrice    json.Number                   `json:"total_line_items_price"`
	CartToken              string                        `json:"cart_token"`
	BuyerAcceptsMarketing  bool                          `json:"buyer_accepts_marketing"`
	Name                   string                        `json:"name"`
	ReferringSite          string                        `json:"referring_site"`
	LandingSite            string                        `json:"landing_site"`
	CancelledAt            string                        `json:"cancelled_at"`
	CancelReason           string                        `json:"cancel_reason"`
	TotalPriceUsd          json.Number                   `json:"total_price_usd"`
	CheckoutToken          string                        `json:"checkout_token"`
	Reference              interface{}                   `json:"reference"`
	UserID                 interface{}                   `json:"user_id"`
	LocationID             interface{}                   `json:"location_id"`
	SourceIdentifier       interface{}                   `json:"source_identifier"`
	SourceURL              string                        `json:"source_url"`
	ProcessedAt            string                        `json:"processed_at"`
	DeviceID               string                        `json:"device_id"`
	Phone                  string                        `json:"phone"`
	CustomerLocale         string                        `json:"customer_locale"`
	AppID                  json.Number                   `json:"app_id"`
	BrowserIP              string                        `json:"browser_ip"`
	LandingSiteRef         string                        `json:"landing_site_ref"`
	OrderNumber            int64                         `json:"order_number"`
	DiscountApplications   []ShopifyDiscountApplications `json:"discount_applications"`
	DiscountCodes          []DiscountCode                `json:"discount_codes"`
	NoteAttributes         []interface{}                 `json:"note_attributes"`
	PaymentGatewayNames    []string                      `json:"payment_gateway_names"`
	ProcessingMethod       string                        `json:"processing_method"`
	CheckoutID             interface{}                   `json:"checkout_id"`
	SourceName             string                        `json:"source_name"`
	FulfillmentStatus      string                        `json:"fulfillment_status"`
	TaxLines               []interface{}                 `json:"tax_lines"`
	Tags                   string                        `json:"tags"`
	ContactEmail           string                        `json:"contact_email"`
	OrderStatusURL         string                        `json:"order_status_url"`
	PresentmentCurrency    string                        `json:"presentment_currency"`
	TotalLineItemsPriceSet ShopifyPriceSet               `json:"total_line_items_price_set"`
	TotalDiscountsSet      ShopifyPriceSet               `json:"total_discounts_set"`
	TotalShippingPriceSet  ShopifyPriceSet               `json:"total_shipping_price_set"`
	SubtotalPriceSet       ShopifyPriceSet               `json:"subtotal_price_set"`
	TotalPriceSet          ShopifyPriceSet               `json:"total_price_set"`
	TotalTaxSet            ShopifyPriceSet               `json:"total_tax_set"`
	LineItems              []WebHookLineItem             `json:"line_items"`
	Fulfillments           []interface{}                 `json:"fulfillments"`
	Refunds                []interface{}                 `json:"refunds"`
	TotalTipReceived       string                        `json:"total_tip_received"`
	OriginalTotalDutiesSet interface{}                   `json:"original_total_duties_set"`
	CurrentTotalDutiesSet  interface{}                   `json:"current_total_duties_set"`
	AdminGraphqlAPIID      string                        `json:"admin_graphql_api_id"`
	ShippingLines          []ShopifyShippingLines        `json:"shipping_lines"`
	BillingAddress         ShopifyAddress                `json:"billing_address"`
	ShippingAddress        ShopifyAddress                `json:"shipping_address"`
	Customer               ShopifyCustomer               `json:"customer"`
}

type Broker interface {
	Publish(string, []byte) error
}

type WebHookToBroker struct {
	ShopifyKey string
	Broker     Broker
}

func NewWebHookToBroker(shopifyKey string, broker Broker) *WebHookToBroker {
	return &WebHookToBroker{
		ShopifyKey: shopifyKey,
		Broker:     broker,
	}
}

func (w *WebHookToBroker) Handler(order *WebHookOrder) error {
	b, err := json.Marshal(order)
	if err != nil {
		return err
	}

	t := "arn:aws:sns:us-east-1:955897122252:shopify"

	log.Println(t)

	return w.Broker.Publish(t, b)
}
