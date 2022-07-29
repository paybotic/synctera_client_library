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

// A prospect has a unique identifier. It can be upgrade to a customer with required information
type Prospect struct {
	// The date and time the resource was created.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// Customer's email
	Email string `json:"email,omitempty"`
	// Customer unique identifier
	Id string `json:"id,omitempty"`
	// Customer's KYC exemption
	KycExempt bool `json:"kyc_exempt,omitempty"`
	// Date and time KYC was last run on the customer
	KycLastRun time.Time `json:"kyc_last_run,omitempty"`
	KycStatus *CustomerKycStatus `json:"kyc_status,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime time.Time `json:"last_updated_time,omitempty"`
	LegalAddress *Address `json:"legal_address,omitempty"`
	// User-supplied metadata. Do not use to store PII.
	Metadata *interface{} `json:"metadata,omitempty"`
	// Customer's middle name
	MiddleName string `json:"middle_name,omitempty"`
	// Customer's mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Customer's relationships with other accounts eg. guardian
	RelatedCustomers []Relationship1 `json:"related_customers,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
	// Customer's full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC on a customer. Input must match the pattern ^\\d{3}-\\d{2}-\\d{4}$. The response contains the last 4 digits only (e.g. 6789).
	Ssn string `json:"ssn,omitempty"`
	SsnSource *SsnSource `json:"ssn_source,omitempty"`
}
