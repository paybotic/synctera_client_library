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

type WatchlistSubscription struct {
	// Whether this subscription should automatically renew when the subscription period is over (default: vendor-dependent). 
	AutoRenew bool `json:"auto_renew,omitempty"`
	// When this subscription was created
	Created time.Time `json:"created,omitempty"`
	// Whether this customer has consented to being enrolled for watchlist monitoring 
	CustomerConsent bool `json:"customer_consent"`
	// Unique identifier for this subscription
	Id string `json:"id,omitempty"`
	// The date when monitoring of this individual should end.
	PeriodEnd string `json:"period_end,omitempty"`
	// The date when monitoring of this individual should begin (default: today).
	PeriodStart string `json:"period_start,omitempty"`
	// External provider subscription id
	ProviderSubscriptionId string `json:"provider_subscription_id,omitempty"`
	Status string `json:"status,omitempty"`
}
