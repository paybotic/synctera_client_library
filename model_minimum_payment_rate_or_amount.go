/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MinimumPaymentRateOrAmount struct {
	// The maximum amount to charge as a minimum payment, in cents. For example, to set the maximum to $30, set this value to 3000. 
	MinAmount int64 `json:"min_amount"`
	// The percentage of the balance to use, in basis points. For example, to set 12.5% of the balance, set this value to 1250. 
	Rate int32 `json:"rate"`
	Type_ *MinimumPaymentType `json:"type"`
}
