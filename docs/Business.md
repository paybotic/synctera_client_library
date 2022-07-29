# Business

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | [**time.Time**](time.Time.md) | The date and time the resource was created. | [optional] [default to null]
**Ein** | **string** | U.S. Employer Identification Number (EIN) for this business, in the format xx-xxxxxxx. | [optional] [default to null]
**Email** | **string** | Business&#x27;s email. | [optional] [default to null]
**EntityName** | **string** | Business&#x27;s legal name. | [optional] [default to null]
**FormationDate** | **string** | Date the business was legally registered in RFC 3339 full-date format (YYYY-MM-DD). | [optional] [default to null]
**FormationState** | **string** | U.S. state where the business is legally registered (2-letter abbreviation). | [optional] [default to null]
**Id** | **string** | Business&#x27;s unique identifier. | [optional] [default to null]
**IsCustomer** | **bool** |  | [optional] [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) | The date and time the resource was last updated. | [optional] [default to null]
**LegalAddress** | [***Address**](address.md) |  | [optional] [default to null]
**Metadata** | [***Metadata**](metadata.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | Business&#x27;s phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated. | [optional] [default to null]
**Status** | **string** | Status of the business. One of the following: * &#x60;PROSPECT&#x60; – a potential customer, used for information-gathering and disclosures. * &#x60;ACTIVE&#x60; –  is an integrator defined status.  Integrators should set a business to active if they believe the person to be qualified for conducting business.  Synctera will combine this status with other statuses such a verification to determine if the business is eligible for specific actions such as initiating transactions or issuing a card. * &#x60;FROZEN&#x60; – business&#x27;s actions are blocked for security, legal, or other reasons. * &#x60;SANCTION&#x60; – business is on a sanctions list and should be carefully monitored. * &#x60;DISSOLVED&#x60; – an inactive status indicating a business entity has filed articles of dissolution or a certificate of termination to terminate its existence. * &#x60;CANCELLED&#x60; – an inactive status indicating that a business entity has filed a cancellation or has failed to file its periodic report after notice of forfeiture of its rights to do business. * &#x60;SUSPENDED&#x60; – an inactive status indicating that the business entity has lost the right to operate in it&#x27;s registered jurisdiction. * &#x60;MERGED&#x60; – an inactive status indicating that the business entity has terminated existence by merging into another entity. * &#x60;INACTIVE&#x60; – an inactive status indicating that the business entity is no longer active. * &#x60;CONVERTED&#x60; – An inactive status indicating that the business entity has been converted to another type of business entity in the same jurisdiction.  | [optional] [default to null]
**Structure** | **string** | Business&#x27;s legal structure. | [optional] [default to null]
**TradeNames** | **[]string** | Other names by which this business is known. | [optional] [default to null]
**VerificationLastRun** | [**time.Time**](time.Time.md) | Date and time KYB verification was last run on the business. | [optional] [default to null]
**VerificationStatus** | [***VerificationStatus**](verification_status.md) |  | [optional] [default to null]
**Website** | **string** | Business&#x27;s website. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

