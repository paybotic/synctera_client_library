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

type Business1 struct {
	CreationTime time.Time `json:"creation_time,omitempty"`
	Ein string `json:"ein,omitempty"`
	Email string `json:"email,omitempty"`
	EntityName string `json:"entity_name,omitempty"`
	FormationDate string `json:"formation_date,omitempty"`
	FormationState string `json:"formation_state,omitempty"`
	Id string `json:"id,omitempty"`
	LastUpdatedTime time.Time `json:"last_updated_time,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Status string `json:"status,omitempty"`
	Structure string `json:"structure,omitempty"`
	TradeNames []string `json:"trade_names,omitempty"`
	VerificationLastRun time.Time `json:"verification_last_run,omitempty"`
	VerificationStatus string `json:"verification_status,omitempty"`
}
