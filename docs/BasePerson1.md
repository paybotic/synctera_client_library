# BasePerson1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | [**time.Time**](time.Time.md) | The date and time the resource was created. | [optional] [default to null]
**Dob** | **string** | Person&#x27;s date of birth in RFC 3339 full-date format (YYYY-MM-DD). | [optional] [default to null]
**Email** | **string** | Person&#x27;s email. | [optional] [default to null]
**FirstName** | **string** | Person&#x27;s first name. | [optional] [default to null]
**Id** | **string** | Person&#x27;s unique identifier. | [optional] [default to null]
**IsCustomer** | **bool** |  | [optional] [default to null]
**LastName** | **string** | Person&#x27;s last name. | [optional] [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) | The date and time the resource was last updated. | [optional] [default to null]
**LegalAddress** | [***Address**](address.md) |  | [optional] [default to null]
**Metadata** | [***Metadata**](metadata.md) |  | [optional] [default to null]
**MiddleName** | **string** | Person&#x27;s middle name. | [optional] [default to null]
**PhoneNumber** | **string** | Person&#x27;s mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated | [optional] [default to null]
**ShippingAddress** | [***Address**](address.md) |  | [optional] [default to null]
**Ssn** | **string** | Person&#x27;s full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC. The response contains the last 4 digits only (e.g. 6789). | [optional] [default to null]
**SsnSource** | [***SsnSource**](ssn_source.md) |  | [optional] [default to null]
**Status** | [***Status1**](status1.md) |  | [optional] [default to null]
**VerificationLastRun** | [**time.Time**](time.Time.md) | Date and time KYC verification was last run on the person. | [optional] [default to null]
**VerificationStatus** | [***VerificationStatus**](verification_status.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

