/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// DisclosureType : Describes the regulatory requirement that triggered the disclosure. One of the following: * `REG_DD` –     US regulation that implements the Truth in Savings Act,     to inform customers about the terms and rules for a deposit account. * `KYC_DATA_COLLECTION` –     Document advising the customer that you will collect their personal information     (name, date of birth, tax ID, etc.)     and will be validating their information against external data/documentation. * `REG_E` –     US regulation that implements the Electronic Funds Transfer Act:     covers liability for electronic transactions,     disputes for fraudulent or unrecognized electronic transactions,     and consent for electronic debits from a consumer's account. * `REG_CC` –     US regulation that implements the Expedited Funds Availability Act:     describes standards for when a financial institution     makes funds available in a deposit account. * `E_SIGN` –     US law with rules around electronic agreements/documents/disclosures:     used to obtain consent from consumers to receive electronic communications     (agreements, disclosures, statements, etc) about their accounts. * `PRIVACY_NOTICE` –     Document that tells customers what is done with their non-public information,     who it is shared with, how is is secured,     and how they can opt out of it being shared beyond Synctera. * `TERMS_AND_CONDITIONS` –     A detailed agreement between you and the consumer for the     structure, terms, fees, charges, rates of the product or service,     and the rules for the relationship between you and the consumer. 
type DisclosureType string

// List of disclosure_type
const (
	REG_DD_DisclosureType DisclosureType = "REG_DD"
	KYC_DATA_COLLECTION_DisclosureType DisclosureType = "KYC_DATA_COLLECTION"
	REG_E_DisclosureType DisclosureType = "REG_E"
	REG_CC_DisclosureType DisclosureType = "REG_CC"
	E_SIGN_DisclosureType DisclosureType = "E_SIGN"
	PRIVACY_NOTICE_DisclosureType DisclosureType = "PRIVACY_NOTICE"
	TERMS_AND_CONDITIONS_DisclosureType DisclosureType = "TERMS_AND_CONDITIONS"
)
