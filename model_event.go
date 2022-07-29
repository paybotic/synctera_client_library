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

// Webhook event object
type Event struct {
	// Json string of object associated with the event. For example, if your event is ACCOUNT.CREATED, You can refer to Acccount to parse the account event to obtain the ID, status etc. 
	EventResource string `json:"event_resource,omitempty"`
	// Json string of object associated with the event related to a resource change. This only contains those fields that have value changed on the event, and the field values are prior to the resource change event. 
	EventResourceChangedFields string `json:"event_resource_changed_fields,omitempty"`
	// Timestamp of the current event raised
	EventTime time.Time `json:"event_time,omitempty"`
	// Unique event ID of the webhook request. Use event endpoints to get more event summary data
	Id string `json:"id,omitempty"`
	// Metadata that stored in the webhook subscription
	Metadata string `json:"metadata,omitempty"`
	// Response history of the webhook request
	ResponseHistory []ResponseHistoryItem `json:"response_history,omitempty"`
	// Current event status. Failing event will keep retry until it is purged.
	Status string `json:"status,omitempty"`
	Type_ *EventTypeExplicit `json:"type,omitempty"`
	// URL that the current event will be sent to
	Url string `json:"url,omitempty"`
	// Webhook the current event belongs to
	WebhookId string `json:"webhook_id,omitempty"`
}
