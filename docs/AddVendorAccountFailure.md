# AddVendorAccountFailure

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | [***AddVendorAccountsErrorReason**](add_vendor_accounts_error_reason.md) |  | [default to null]
**ReasonDescription** | **string** | A human-readable message describing the reason for the failure. | [default to null]
**VendorAccountId** | **string** | The vendor account ID for the account that failed. For Plaid, this is an &#x60;account_id&#x60;.  | [default to null]
**VendorErrorMessage** | **string** | The display_message returned by the vendor. Only returned if reason is set to &#x60;PROVIDER_ERROR&#x60;. For Plaid, this is the &#x60;display_message&#x60;.  | [optional] [default to null]
**VendorRequestId** | **string** | A unique identifier for the request from the vendor, which can be used for troubleshooting. Only returned if reason is set to &#x60;PROVIDER_ERROR&#x60;.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

