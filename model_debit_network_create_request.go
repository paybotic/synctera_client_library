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

type DebitNetworkCreateRequest struct {
	// indicates whether debit network is active
	Active bool `json:"active,omitempty"`
	// The timestamp representing when the debit network was created
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The time when debit network became inactive
	EndDate time.Time `json:"end_date,omitempty"`
	// Debit Network ID
	Id string `json:"id,omitempty"`
	// The timestamp representing when the debit network was last modified
	LastModifiedTime time.Time `json:"last_modified_time,omitempty"`
	// The name describing the debit network
	Name string `json:"name"`
	// The time when debit network goes live
	StartDate time.Time `json:"start_date,omitempty"`
}
