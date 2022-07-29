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

type GatewayResponse struct {
	// Current status of the Authorization gateway
	Active bool `json:"active"`
	// The timestamp representing when the gateway config request was made
	CreationTime time.Time `json:"creation_time"`
	CustomHeaders *map[string]string `json:"custom_headers,omitempty"`
	// Gateway ID
	Id string `json:"id"`
	// The timestamp representing when the gateway config was last modified at
	LastModifiedTime time.Time `json:"last_modified_time"`
	// URL of the Authorization gateway
	Url string `json:"url"`
}
