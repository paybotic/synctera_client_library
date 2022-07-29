/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InternalTransfer struct {
	// The amount (in cents) to transfer from originating account to receiving account.
	Amount int32 `json:"amount"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	// A short note to the recipient
	Memo string `json:"memo,omitempty"`
	// Arbitrary key-value metadata to associate with the transaction
	Metadata *interface{} `json:"metadata,omitempty"`
	// An alias representing a GL account to debit. This is alternative to specifying by account id
	OriginatingAccountAlias string `json:"originating_account_alias,omitempty"`
	// The customer id of the owner of the originating account.
	OriginatingAccountCustomerId string `json:"originating_account_customer_id,omitempty"`
	// The UUID of the account being debited
	OriginatingAccountId string `json:"originating_account_id,omitempty"`
	// An alias representing a GL account to credit. This is an alternative to specifying by account id
	ReceivingAccountAlias string `json:"receiving_account_alias,omitempty"`
	// The customer id of the owner of the receiving account. Only required when type is \"outgoing_remittance\"
	ReceivingAccountCustomerId string `json:"receiving_account_customer_id,omitempty"`
	// The UUID of the account being credited
	ReceivingAccountId string `json:"receiving_account_id,omitempty"`
	// The desired transaction type to use for this transfer
	Type_ string `json:"type"`
}
