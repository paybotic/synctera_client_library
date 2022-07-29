# PatchCustomer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dob** | **string** | Customer&#x27;s date of birth in RFC 3339 full-date format (YYYY-MM-DD) | [optional] [default to null]
**Email** | **string** | Customer&#x27;s email | [optional] [default to null]
**FirstName** | **string** | Customer&#x27;s first name | [optional] [default to null]
**LastName** | **string** | Customer&#x27;s last name | [optional] [default to null]
**LegalAddress** | [***Address**](address.md) |  | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | User-supplied JSON format metadata. Do not use to store PII. | [optional] [default to null]
**MiddleName** | **string** | Customer&#x27;s middle name | [optional] [default to null]
**PhoneNumber** | **string** | Customer&#x27;s mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated. | [optional] [default to null]
**ShippingAddress** | [***Address**](address.md) |  | [optional] [default to null]
**Ssn** | **string** | Customer&#x27;s full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC on a customer. Must be compiled with ^\\d{3}-\\d{2}-\\d{4}$. Response contains the last 4 digits only (e.g. 6789). | [optional] [default to null]
**Status** | **string** | Customer&#x27;s status | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

