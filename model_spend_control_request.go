package openapi

// SpendControlRequest defines the request structure for spend controls
type SpendControlRequest struct {
	// Add fields based on your API requirements
	AccountId  string `json:"account_id,omitempty"`
	CustomerId string `json:"customer_id,omitempty"`
	BusinessId string `json:"business_id,omitempty"`
	// Add other fields as needed
}
