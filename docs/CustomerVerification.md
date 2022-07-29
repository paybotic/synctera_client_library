# CustomerVerification

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerConsent** | **bool** | Whether this customer has consented to a KYC check.  | [default to null]
**CustomerIpAddress** | **string** | IP address of the customer being verified. | [optional] [default to null]
**DocumentId** | **string** | The ID of the uploaded government-issued identification document provided by the Socure SDK.  | [optional] [default to null]
**VerificationType** | [**[]VerificationType**](verification_type.md) | List of possible checks to run on a customer. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

