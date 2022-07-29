# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalCardFromToken**](ExternalCardsAlphaApi.md#CreateExternalCardFromToken) | **Post** /external_cards/tokens | Create external card from token
[**CreateExternalCardTransfer**](ExternalCardsAlphaApi.md#CreateExternalCardTransfer) | **Post** /external_cards/transfers | Create external card transfer
[**GetExternalCard**](ExternalCardsAlphaApi.md#GetExternalCard) | **Get** /external_cards/{external_card_id} | Get a external card
[**GetExternalCardTransfer**](ExternalCardsAlphaApi.md#GetExternalCardTransfer) | **Get** /external_cards/transfers/{transfer_id} | Get an external card transfer
[**ListExternalCardTransfers**](ExternalCardsAlphaApi.md#ListExternalCardTransfers) | **Get** /external_cards/transfers | List external transfers
[**ListExternalCards**](ExternalCardsAlphaApi.md#ListExternalCards) | **Get** /external_cards | List external cards

# **CreateExternalCardFromToken**
> ExternalCardResponse CreateExternalCardFromToken(ctx, body)
Create external card from token

Create an external card from token - You must first tokenize the external card using the external card iframe. You will receive a token for the external card upon successful completion. This endpoint will persist the external card and associate it to a customer. The customer's name and address should match that of the external card as entered in the iframe. <br>NB: Tokens should be associated right away. Tokens not associated within 30 mins of creation will be invalidated.</br> NB: If a valid business_id is provided, the address verification for the external card will be done against the address of the business. Otherwise, the address of the customer will be used. In either case, the name of the customer will be used to match the name of the cardholder. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExternalCardRequest**](ExternalCardRequest.md)| Details of the card to create | 

### Return type

[**ExternalCardResponse**](external_card_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateExternalCardTransfer**
> TransferResponse CreateExternalCardTransfer(ctx, body)
Create external card transfer

Create a external card transfer 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransferRequest**](TransferRequest.md)| Details of the transfer to create | 

### Return type

[**TransferResponse**](transfer_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalCard**
> ExternalCardResponse GetExternalCard(ctx, externalCardId)
Get a external card

Get an external card 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalCardId** | [**string**](.md)| The unique identifier of an external card | 

### Return type

[**ExternalCardResponse**](external_card_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalCardTransfer**
> TransferResponse GetExternalCardTransfer(ctx, transferId)
Get an external card transfer

Get an external transfer 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transferId** | [**string**](.md)| The unique identifier of a transfer | 

### Return type

[**TransferResponse**](transfer_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExternalCardTransfers**
> TransferListResponse ListExternalCardTransfers(ctx, optional)
List external transfers

List external card transfer 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExternalCardsAlphaApiListExternalCardTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalCardsAlphaApiListExternalCardTransfersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 
 **customerId** | [**optional.Interface of []string**](string.md)|  | 
 **externalCardId** | [**optional.Interface of string**](.md)|  | 
 **originatingAccountId** | [**optional.Interface of []string**](string.md)|  | 
 **type_** | [**optional.Interface of TransferType**](.md)| The type of an external transfer | 

### Return type

[**TransferListResponse**](transfer_list_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExternalCards**
> ExternalCardListResponse ListExternalCards(ctx, optional)
List external cards

List external cards 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExternalCardsAlphaApiListExternalCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExternalCardsAlphaApiListExternalCardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 
 **customerId** | [**optional.Interface of string**](.md)|  | 

### Return type

[**ExternalCardListResponse**](external_card_list_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

