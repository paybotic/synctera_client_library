/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AddVendorAccountsRequest struct {
	// The identifier for the business customer associated with this external account. Exactly one of `business_id` or `customer_id` must be specified. 
	BusinessId string `json:"business_id,omitempty"`
	// The identifier for the personal customer associated with this external account. Exactly one of `customer_id` or `business_id` must be specified. 
	CustomerId string `json:"customer_id,omitempty"`
	CustomerType *ExtAccountCustomerType `json:"customer_type"`
	Vendor *ExternalAccountVendorValues `json:"vendor"`
	// The token provided to link external accounts. For Plaid, this is their `access_token`. 
	VendorAccessToken string `json:"vendor_access_token,omitempty"`
	// The list of vendor account IDs that the customer chose to link. For Plaid, these are `account_id`s. 
	VendorAccountIds []string `json:"vendor_account_ids,omitempty"`
	// The identifier provided by the vendor for the customer associated with this external account. 
	VendorCustomerId string `json:"vendor_customer_id,omitempty"`
	// If true, Synctera will attempt to verify that the external account owner is the same as the customer by comparing external account data to customer data. At least 2 of the following fields must match: name, phone number, email, address. Verification is disabled by default. 
	VerifyOwner bool `json:"verify_owner,omitempty"`
}
