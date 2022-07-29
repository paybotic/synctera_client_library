/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type StatementSummary struct {
	// The unique identifier of the account the statement belongs to
	AccountId string `json:"account_id,omitempty"`
	// The limit date when the due amount indicated on the statement should be paid
	DueDate string `json:"due_date,omitempty"`
	// The date indicating the ending of the time interval covered by the statement
	EndDate string `json:"end_date,omitempty"`
	// statement ID
	Id string `json:"id,omitempty"`
	// The date when the statement has been issued
	IssueDate string `json:"issue_date,omitempty"`
	// The date indicating the begining of the time interval covered by the statement
	StartDate string `json:"start_date,omitempty"`
}
