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

// Person object for patch purpose.
type PatchPerson struct {
	// The date and time the resource was created.
	CreationTime time.Time `json:"creation_time,omitempty"`
	// Person's date of birth in RFC 3339 full-date format (YYYY-MM-DD).
	Dob string `json:"dob,omitempty"`
	// Person's email.
	Email string `json:"email,omitempty"`
	// Person's first name.
	FirstName string `json:"first_name,omitempty"`
	// Person's unique identifier.
	Id string `json:"id,omitempty"`
	IsCustomer bool `json:"is_customer,omitempty"`
	// Person's last name.
	LastName string `json:"last_name,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime time.Time `json:"last_updated_time,omitempty"`
	LegalAddress *Address `json:"legal_address,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	// Person's middle name.
	MiddleName string `json:"middle_name,omitempty"`
	// Person's mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated
	PhoneNumber string `json:"phone_number,omitempty"`
	ShippingAddress *Address `json:"shipping_address,omitempty"`
	// Person's full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC. The response contains the last 4 digits only (e.g. 6789).
	Ssn string `json:"ssn,omitempty"`
	SsnSource *SsnSource `json:"ssn_source,omitempty"`
	Status *Status1 `json:"status,omitempty"`
	// Date and time KYC verification was last run on the person.
	VerificationLastRun time.Time `json:"verification_last_run,omitempty"`
	VerificationStatus *VerificationStatus `json:"verification_status,omitempty"`
}
