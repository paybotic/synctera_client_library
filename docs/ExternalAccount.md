# ExternalAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifiers** | [***AccountIdentifiers**](account_identifiers.md) |  | [default to null]
**AccountOwnerNames** | **[]string** | The names of the account owners. Values may be masked, in which case the array will be empty.  | [default to null]
**BusinessId** | **string** | The identifier for the business customer associated with this external account. Exactly one of &#x60;business_id&#x60; or &#x60;customer_id&#x60; will be returned.  | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) |  | [default to null]
**CustomerId** | **string** | The identifier for the personal customer associated with this external account. Exactly one of &#x60;customer_id&#x60; or &#x60;business_id&#x60; will be returned.  | [optional] [default to null]
**Id** | **string** | External account unique identifier | [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) |  | [default to null]
**Metadata** | [***interface{}**](interface{}.md) | User-supplied JSON format metadata. | [optional] [default to null]
**Name** | **string** | The official name of the account | [optional] [default to null]
**Nickname** | **string** | A user-meaningful name for the account | [optional] [default to null]
**RoutingIdentifiers** | [***AccountRouting**](account_routing.md) |  | [default to null]
**Status** | **string** | The current state of the account | [default to null]
**Type_** | **string** | The type of the account | [default to null]
**VendorData** | [***ExternalAccountVendorData**](external_account_vendor_data.md) |  | [optional] [default to null]
**VendorInfo** | [***VendorInfo**](vendor_info.md) |  | [optional] [default to null]
**Verification** | [***AccountVerification**](account_verification.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

