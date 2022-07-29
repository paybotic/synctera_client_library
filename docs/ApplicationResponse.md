# ApplicationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDetails** | [***interface{}**](interface{}.md) | Details about the applicant. The exact schema is to be determined with your bank. | [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | Application creation timestamp in RFC3339 format | [default to null]
**CustomerId** | **string** | Customer ID for the application | [default to null]
**Id** | **string** | Generated ID for the application | [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) | Timestamp of the last application modification in RFC3339 format | [default to null]
**Status** | [***ApplicationStatus**](application_status.md) |  | [default to null]
**Type_** | [***ApplicationType**](application_type.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

