# ExternalAccountLinkToken

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | **string** | The identifier for the business customer associated with this external account. Exactly one of &#x60;business_id&#x60; or &#x60;customer_id&#x60; must be specified.  | [optional] [default to null]
**ClientName** | **string** | The name of your application, as it should be displayed in Link. Maximum length of 30 characters. | [default to null]
**CountryCodes** | **[]string** | Country codes in the ISO-3166-1 alpha-2 country code standard. | [default to null]
**CustomerId** | **string** | The identifier for the personal customer associated with this external account. Exactly one of &#x60;customer_id&#x60; or &#x60;business_id&#x60; must be specified.  | [optional] [default to null]
**Expiration** | [**time.Time**](time.Time.md) | The expiration date for the link_token. Expires in 4 hours. | [optional] [default to null]
**Language** | **string** | The language that corresponds to the link token. For Plaid, see their [documentation](https://plaid.com/docs/api/tokens/#link-token-create-request-language) for a list of allowed values.  | [default to null]
**LinkCustomizationName** | **string** | The name of the Link customization from the Plaid Dashboard to be applied to Link. If not specified, the default customization will be used. When using a Link customization, the language in the customization must match the language selected via the language parameter, and the countries in the customization should match the country codes selected via country_codes.  | [optional] [default to null]
**LinkToken** | **string** | A link_token, which can be supplied to Link in order to initialize it and receive a public_token, which can be exchanged for an access_token.  | [optional] [default to null]
**RedirectUri** | **string** | A URI indicating the destination where a user should be forwarded after completing the Link flow; used to support OAuth authentication flows when launching Link in the browser or via a webview.  | [optional] [default to null]
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. | [optional] [default to null]
**SdkType** | **string** | Describes the environment of the client code running a vendor-supplied SDK | [optional] [default to SDK_TYPE.WEB]
**Type_** | **string** | The type of the link token. DEPOSITORY for checking and savings accounts, CREDIT for credit card type accounts. | [default to null]
**VendorAccessToken** | **string** | The access token associated with the Item data is being requested for. | [optional] [default to null]
**VendorInstitutionId** | **string** | The ID of the institution the access token is requested for. If present the link token will be created in an update mode.  | [optional] [default to null]
**VerifyOwner** | **bool** | If true, Synctera will attempt to verify that the external account owner is the same as the customer by comparing external account data to customer data. At least 2 of the following fields must match: name, phone number, email, address. Verification is disabled by default.  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

