# BasePerson

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | [**time.Time**](time.Time.md) | The date and time the resource was created. | [optional] [default to null]
**Email** | **string** | Customer&#x27;s email | [optional] [default to null]
**Id** | **string** | Customer unique identifier | [optional] [default to null]
**KycExempt** | **bool** | Customer&#x27;s KYC exemption | [optional] [default to null]
**KycLastRun** | [**time.Time**](time.Time.md) | Date and time KYC was last run on the customer | [optional] [default to null]
**KycStatus** | [***CustomerKycStatus**](customer_kyc_status.md) |  | [optional] [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) | The date and time the resource was last updated. | [optional] [default to null]
**LegalAddress** | [***Address**](address.md) |  | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | User-supplied metadata. Do not use to store PII. | [optional] [default to null]
**MiddleName** | **string** | Customer&#x27;s middle name | [optional] [default to null]
**PhoneNumber** | **string** | Customer&#x27;s mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated. | [optional] [default to null]
**RelatedCustomers** | [**[]Relationship1**](relationship1.md) | Customer&#x27;s relationships with other accounts eg. guardian | [optional] [default to null]
**ShippingAddress** | [***Address**](address.md) |  | [optional] [default to null]
**Ssn** | **string** | Customer&#x27;s full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC on a customer. Input must match the pattern ^\\d{3}-\\d{2}-\\d{4}$. The response contains the last 4 digits only (e.g. 6789). | [optional] [default to null]
**SsnSource** | [***SsnSource**](ssn_source.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

