# ExternalAccountAccessToken

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | **string** | The identifier for the business customer associated with this external account. Exactly one of &#x60;business_id&#x60; or &#x60;customer_id&#x60; must be specified.  | [optional] [default to null]
**CustomerId** | **string** | The identifier for the personal customer associated with this external account. Exactly one of &#x60;customer_id&#x60; or &#x60;business_id&#x60; must be specified.  | [optional] [default to null]
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting | [optional] [default to null]
**VendorAccessToken** | **string** | The access token associated with the Item data is being requested for. | [optional] [default to null]
**VendorCustomerId** | **string** | An alias for &#x60;customer_id&#x60; (deprecated). | [optional] [default to null]
**VendorInstitutionId** | **string** | The ID of the institution the access token is requested for  | [default to null]
**VendorPublicToken** | **string** | The user&#x27;s public token obtained from successful link login.  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

