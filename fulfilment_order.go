package goshopify

import (
	"fmt"
	"time"
)

type FulfillmentOrderService interface {
	List(int64, interface{}) ([]*FulfillmentOrder, error)
}
type FulfillmentOrderServiceOp struct {
	client     *Client
	resource   string
	resourceID int64
}
type FulfillmentOrdersService interface {
	GetListFulfillmentOrder(int64, interface{}) ([]*FulfillmentOrder, error)
}

// Get fulfilment orders
func (s *FulfillmentOrderServiceOp) List(options interface{}) ([]*FulfillmentOrder, error) {
	path := fmt.Sprintf("%s/%d/fulfillment_orders.json", ordersBasePath, s.resourceID)
	resource := new(FulfillmentOrderResouce)
	err := s.client.Get(path, resource, options)
	return resource.FulfillmentOrders, err
}

type FulfillmentOrderResouce struct {
	FulfillmentOrders []*FulfillmentOrder `json:"fulfillment_orders"`
}
type FulfillmentOrder struct {
	ID                 int64                  `json:"id"`
	ShopID             int64                  `json:"shop_id"`
	OrderID            int64                  `json:"order_id"`
	AssignedLocationID int64                  `json:"assigned_location_id"`
	RequestStatus      string                 `json:"request_status"`
	Status             string                 `json:"status"`
	SupportedActions   []string               `json:"supported_actions"`
	Destination        *Address               `json:"destination"`
	LineItems          []*FulfillmentLineItem `json:"line_items"`
	FulfillAt          *time.Time             `json:"fulfill_at"`
	FulfillBy          *string                `json:"fulfill_by"`
	DeliveryMethod     DeliveryMethod         `json:"delivery_method"`
	CreatedAt          string                 `json:"created_at"`
	UpdatedAt          string                 `json:"updated_at"`
	AssignedLocation   AssignedLocation       `json:"assigned_location"`
}
type AssignedLocation struct {
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	CountryCode string `json:"country_code"`
	LocationID  int64  `json:"location_id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Province    string `json:"province"`
	Zip         string `json:"zip"`
}
type DeliveryMethod struct {
	ID                  int64   `json:"id"`
	MethodType          string  `json:"method_type"`
	MinDeliveryDateTime *string `json:"min_delivery_date_time"`
	MaxDeliveryDateTime *string `json:"max_delivery_date_time"`
}
type FulfillmentLineItem struct {
	ID                  int64 `json:"id"`
	ShopID              int64 `json:"shop_id"`
	FulfillmentOrderID  int64 `json:"fulfillment_order_id"`
	Quantity            int   `json:"quantity"`
	LineItemID          int64 `json:"line_item_id"`
	InventoryItemID     int64 `json:"inventory_item_id"`
	FulfillableQuantity int   `json:"fulfillable_quantity"`
	VariantID           int64 `json:"variant_id"`
}
