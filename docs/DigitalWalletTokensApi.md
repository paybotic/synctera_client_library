# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDigitalWalletApple**](DigitalWalletTokensApi.md#CreateDigitalWalletApple) | **Post** /cards/{card_id}/digital_wallet_tokens/applepay | Create digital wallet token provision request for Apple Pay
[**CreateDigitalWalletGoogle**](DigitalWalletTokensApi.md#CreateDigitalWalletGoogle) | **Post** /cards/{card_id}/digital_wallet_tokens/googlepay | Create digital wallet token provision request for Google Pay
[**GetDigitalWalletToken**](DigitalWalletTokensApi.md#GetDigitalWalletToken) | **Get** /cards/digital_wallet_tokens/{digital_wallet_token_id} | Get Digital Wallet Token
[**ListDigitalWalletTokens**](DigitalWalletTokensApi.md#ListDigitalWalletTokens) | **Get** /cards/digital_wallet_tokens | List Digital Wallet Tokens
[**UpdateDigitalWalletTokenStatus**](DigitalWalletTokensApi.md#UpdateDigitalWalletTokenStatus) | **Patch** /cards/digital_wallet_tokens/{digital_wallet_token_id} | Update Digital Wallet Token&#x27;s life cycle status

# **CreateDigitalWalletApple**
> AppleDigitalWalletProvisionResponse CreateDigitalWalletApple(ctx, body, cardId)
Create digital wallet token provision request for Apple Pay

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AppleDigitalWalletProvisionRequest**](AppleDigitalWalletProvisionRequest.md)| Request to provision digital wallet card data to pass to Apple Pay digital wallet | 
  **cardId** | [**string**](.md)|  | 

### Return type

[**AppleDigitalWalletProvisionResponse**](apple_digital_wallet_provision_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDigitalWalletGoogle**
> GoogleDigitalWalletProvisionResponse CreateDigitalWalletGoogle(ctx, body, cardId)
Create digital wallet token provision request for Google Pay

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GoogleDigitalWalletProvisionRequest**](GoogleDigitalWalletProvisionRequest.md)| Request to provision digital wallet card data to pass to Google Pay digital wallet | 
  **cardId** | [**string**](.md)|  | 

### Return type

[**GoogleDigitalWalletProvisionResponse**](google_digital_wallet_provision_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDigitalWalletToken**
> DigitalWalletTokenResponse GetDigitalWalletToken(ctx, digitalWalletTokenId)
Get Digital Wallet Token

Get the details about the digital wallet token of a card 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **digitalWalletTokenId** | [**string**](.md)|  | 

### Return type

[**DigitalWalletTokenResponse**](digital_wallet_token_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDigitalWalletTokens**
> TokenListResponse ListDigitalWalletTokens(ctx, optional)
List Digital Wallet Tokens

List Digital Wallet Tokens 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DigitalWalletTokensApiListDigitalWalletTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DigitalWalletTokensApiListDigitalWalletTokensOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardId** | [**optional.Interface of string**](.md)|  | 
 **tokenState** | [**optional.Interface of DigitalWalletTokenState**](.md)| The status of the Digital Wallet Token | 
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 

### Return type

[**TokenListResponse**](token_list_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDigitalWalletTokenStatus**
> DigitalWalletTokenResponse UpdateDigitalWalletTokenStatus(ctx, body, digitalWalletTokenId)
Update Digital Wallet Token's life cycle status

The status of a digital wallet token can be updated as, ACTIVE to SUSPENDED, SUSPENDED to ACTIVE, ACTIVE to TERMINATED or SUSPENDED to TERMINATED. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DigitalWalletTokenEditRequest**](DigitalWalletTokenEditRequest.md)| Update Digital wallet token status | 
  **digitalWalletTokenId** | [**string**](.md)|  | 

### Return type

[**DigitalWalletTokenResponse**](digital_wallet_token_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

