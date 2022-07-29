# AddVendorAccountsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | **string** | The identifier for the business customer associated with this external account. Exactly one of &#x60;business_id&#x60; or &#x60;customer_id&#x60; must be specified.  | [optional] [default to null]
**CustomerId** | **string** | The identifier for the personal customer associated with this external account. Exactly one of &#x60;customer_id&#x60; or &#x60;business_id&#x60; must be specified.  | [optional] [default to null]
**CustomerType** | [***ExtAccountCustomerType**](ext_account_customer_type.md) |  | [default to null]
**Vendor** | [***ExternalAccountVendorValues**](external_account_vendor_values.md) |  | [default to null]
**VendorAccessToken** | **string** | The token provided to link external accounts. For Plaid, this is their &#x60;access_token&#x60;.  | [optional] [default to null]
**VendorAccountIds** | **[]string** | The list of vendor account IDs that the customer chose to link. For Plaid, these are &#x60;account_id&#x60;s.  | [optional] [default to null]
**VendorCustomerId** | **string** | The identifier provided by the vendor for the customer associated with this external account.  | [optional] [default to null]
**VerifyOwner** | **bool** | If true, Synctera will attempt to verify that the external account owner is the same as the customer by comparing external account data to customer data. At least 2 of the following fields must match: name, phone number, email, address. Verification is disabled by default.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

