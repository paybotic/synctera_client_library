# AddAccountsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifiers** | [***AddAccountsRequestAccountIdentifiers**](add_accounts_request_account_identifiers.md) |  | [default to null]
**AccountOwnerNames** | **[]string** | The names of the account owners. | [default to null]
**BusinessId** | **string** | The identifier for the business customer associated with this external account. Exactly one of &#x60;business_id&#x60; or &#x60;customer_id&#x60; must be specified.  | [optional] [default to null]
**CustomerId** | **string** | The identifier for the personal customer associated with this external account. Exactly one of &#x60;customer_id&#x60; or &#x60;business_id&#x60; must be specified.  | [optional] [default to null]
**CustomerType** | [***ExtAccountCustomerType**](ext_account_customer_type.md) |  | [default to null]
**Metadata** | [***interface{}**](interface{}.md) | User-supplied metadata | [optional] [default to null]
**Nickname** | **string** | A user-meaningful name for the account | [optional] [default to null]
**RoutingIdentifiers** | [***AddAccountsRequestRoutingIdentifiers**](add_accounts_request_routing_identifiers.md) |  | [default to null]
**Type_** | **string** | The type of the account | [default to null]
**VendorAccountId** | **string** | The ID of the vendor account, will be empty for MANUAL vendor | [optional] [default to null]
**Verification** | [***AccountVerification**](account_verification.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

