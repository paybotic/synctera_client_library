/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VerificationVendorJson struct {
	// Describes the content-type encoding received from the vendor.
	ContentType string `json:"content_type"`
	// Array of vendor specific information.
	Details []VerificationVendorInfoDetail `json:"details,omitempty"`
	// Data representation in JSON.
	Json *interface{} `json:"json"`
	// Name of the vendor used.
	Vendor string `json:"vendor"`
}
