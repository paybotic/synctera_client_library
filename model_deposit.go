/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// Deposit using remote deposit capture
type Deposit struct {
	// The ID of the account
	AccountId string `json:"account_id"`
	// ID of the uploaded image of the back of the check
	BackImageId string `json:"back_image_id"`
	BusinessId string `json:"business_id,omitempty"`
	// Amount on check in ISO 4217 minor currency units
	CheckAmount int32 `json:"check_amount"`
	CreationTime time.Time `json:"creation_time,omitempty"`
	// Date the deposit was captured, in RFC 3339 format
	DateCaptured time.Time `json:"date_captured"`
	// Date the deposit was processed, in RFC 3339 format
	DateProcessed time.Time `json:"date_processed"`
	// Amount deposited in ISO 4217 minor currency units
	DepositAmount int32 `json:"deposit_amount"`
	// ISO 4217 currency code for the deposit amount
	DepositCurrency string `json:"deposit_currency"`
	// ID of the uploaded image of the front of the check
	FrontImageId string `json:"front_image_id"`
	// Remote Check Deposit ID
	Id string `json:"id"`
	LastUpdatedTime time.Time `json:"last_updated_time,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	PersonId string `json:"person_id,omitempty"`
	// The status of the deposit
	Status string `json:"status"`
	// The ID of the transaction associated with this deposit
	TransactionId string `json:"transaction_id"`
	VendorInfo *VendorInfo1 `json:"vendor_info"`
}
