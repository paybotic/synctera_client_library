/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AccountIdentifiers struct {
	// The IBAN of the account. Value may be masked, in which case only the last four digits are returned. 
	Iban string `json:"iban,omitempty"`
	// The account number. Value may be masked, in which case only the last four digits are returned. 
	Number string `json:"number"`
}
