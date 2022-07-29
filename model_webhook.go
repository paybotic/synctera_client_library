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

// Webhook object
type Webhook struct {
	// A description of what the webhook is used for
	Description string `json:"description,omitempty"`
	// A list of the events that will trigger the webhook
	EnabledEvents []EventType1 `json:"enabled_events"`
	// The unique ID of the webhook
	Id string `json:"id,omitempty"`
	// Set the webhook to be enabled or disabled
	IsEnabled bool `json:"is_enabled"`
	// Timestamp that this webhook was created or the last time any field was changed
	LastUpdated time.Time `json:"last_updated,omitempty"`
	// Additional information stored to the webhook
	Metadata string `json:"metadata,omitempty"`
	// URL that the webhook will send request to
	Url string `json:"url"`
}
