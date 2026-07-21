package goshopify

import (
	"fmt"
	"time"
)

const payoutsBasePath = "shopify_payments/payouts"
const payoutTransactionsBasePath = "shopify_payments/balance/transactions"

// PayoutService is an interface for interfacing with the Shopify Payments Payouts API.
// See: https://shopify.dev/docs/api/admin-rest/2025-01/resources/payouts
type PayoutService interface {
	List(interface{}) ([]Payout, error)
	Get(int64, interface{}) (*Payout, error)
	ListTransactions(interface{}) ([]PayoutTransaction, error)
}

// PayoutServiceOp handles communication with the payout related methods of the Shopify API.
type PayoutServiceOp struct {
	client *Client
}

// Payout represents a Shopify Payments payout
type Payout struct {
	ID       int64      `json:"id"`
	Date     string     `json:"date"`
	Currency string     `json:"currency"`
	Amount   string     `json:"amount"`
	Status   string     `json:"status"` // "scheduled", "in_transit", "paid", "failed", "cancelled"
	Summary  *PayoutSummary `json:"summary,omitempty"`
}

// PayoutSummary contains the breakdown of a payout
type PayoutSummary struct {
	AdjustmentsFeeAmount      string `json:"adjustments_fee_amount"`
	AdjustmentsGrossAmount    string `json:"adjustments_gross_amount"`
	ChargesFeeAmount          string `json:"charges_fee_amount"`
	ChargesGrossAmount        string `json:"charges_gross_amount"`
	RefundsFeeAmount          string `json:"refunds_fee_amount"`
	RefundsGrossAmount        string `json:"refunds_gross_amount"`
	ReservedFundsFeeAmount    string `json:"reserved_funds_fee_amount"`
	ReservedFundsGrossAmount  string `json:"reserved_funds_gross_amount"`
	RetriedPayoutsFeeAmount   string `json:"retried_payouts_fee_amount"`
	RetriedPayoutsGrossAmount string `json:"retried_payouts_gross_amount"`
}

// PayoutTransaction represents a balance transaction in a payout
type PayoutTransaction struct {
	ID            int64      `json:"id"`
	Type          string     `json:"type"` // "charge", "refund", "dispute", "reserve", "adjustment", "payout"
	Test          bool       `json:"test"`
	PayoutID      int64      `json:"payout_id"`
	PayoutStatus  string     `json:"payout_status"`
	Currency      string     `json:"currency"`
	Amount        string     `json:"amount"`
	Fee           string     `json:"fee"`
	Net           string     `json:"net"`
	SourceID      int64      `json:"source_id"`
	SourceType    string     `json:"source_type"`
	SourceOrderID int64      `json:"source_order_id"`
	ProcessedAt   *time.Time `json:"processed_at"`
}

// PayoutResource represents the result from the payouts/X.json endpoint
type PayoutResource struct {
	Payout *Payout `json:"payout"`
}

// PayoutsResource represents the result from the payouts.json endpoint
type PayoutsResource struct {
	Payouts []Payout `json:"payouts"`
}

// PayoutTransactionsResource represents the result from the balance/transactions.json endpoint
type PayoutTransactionsResource struct {
	Transactions []PayoutTransaction `json:"transactions"`
}

// List payouts
func (s *PayoutServiceOp) List(options interface{}) ([]Payout, error) {
	path := fmt.Sprintf("%s.json", payoutsBasePath)
	resource := new(PayoutsResource)
	err := s.client.Get(path, resource, options)
	return resource.Payouts, err
}

// Get a single payout by ID
func (s *PayoutServiceOp) Get(payoutID int64, options interface{}) (*Payout, error) {
	path := fmt.Sprintf("%s/%d.json", payoutsBasePath, payoutID)
	resource := new(PayoutResource)
	err := s.client.Get(path, resource, options)
	return resource.Payout, err
}

// ListTransactions fetches balance transactions (fee breakdown per payout)
func (s *PayoutServiceOp) ListTransactions(options interface{}) ([]PayoutTransaction, error) {
	path := fmt.Sprintf("%s.json", payoutTransactionsBasePath)
	resource := new(PayoutTransactionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Transactions, err
}
