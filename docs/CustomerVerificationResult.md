# CustomerVerificationResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for this verification result. | [optional] [default to null]
**Issues** | **[]string** | List of potential problems found. These are subject to change.  | [optional] [default to null]
**RawResponse** | [***RawResponse**](raw_response.md) |  | [optional] [default to null]
**Result** | **string** | The determination of this verification. | [default to null]
**VendorInfo** | [***VerificationVendorInfo**](verification_vendor_info.md) |  | [optional] [default to null]
**VerificationTime** | [**time.Time**](time.Time.md) | The date and time the verification was completed. | [default to null]
**VerificationType** | [***VerificationType**](verification_type.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

