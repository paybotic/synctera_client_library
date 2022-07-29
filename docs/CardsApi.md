# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateCard**](CardsApi.md#ActivateCard) | **Post** /cards/activate | Activate a card
[**CreateCardImage**](CardsApi.md#CreateCardImage) | **Post** /cards/images | Create Card Image
[**GetCard**](CardsApi.md#GetCard) | **Get** /cards/{card_id} | Get Card
[**GetCardBarcode**](CardsApi.md#GetCardBarcode) | **Get** /cards/{card_id}/barcodes | Get Card Barcode
[**GetCardImageData**](CardsApi.md#GetCardImageData) | **Get** /cards/images/{card_image_id}/data | Get Card Image Data
[**GetCardImageDetails**](CardsApi.md#GetCardImageDetails) | **Get** /cards/images/{card_image_id} | Get Card Image Details
[**GetCardWidgetURL**](CardsApi.md#GetCardWidgetURL) | **Get** /cards/card_widget_url | Get card widget URL
[**GetClientAccessToken**](CardsApi.md#GetClientAccessToken) | **Post** /cards/{card_id}/client_token | Get a client token
[**GetClientSingleUseToken**](CardsApi.md#GetClientSingleUseToken) | **Post** /cards/single_use_token | Get single-use token
[**IssueCard**](CardsApi.md#IssueCard) | **Post** /cards | Issue a Card
[**ListCardImageDetails**](CardsApi.md#ListCardImageDetails) | **Get** /cards/images | List Card Image Details
[**ListCardProducts**](CardsApi.md#ListCardProducts) | **Get** /cards/products | List Card Products
[**ListCards**](CardsApi.md#ListCards) | **Get** /cards | List Cards
[**ListChanges**](CardsApi.md#ListChanges) | **Get** /cards/{card_id}/changes | List Card Changes
[**UpdateCard**](CardsApi.md#UpdateCard) | **Patch** /cards/{card_id} | Update Card
[**UpdateCardImageDetails**](CardsApi.md#UpdateCardImageDetails) | **Patch** /cards/images/{card_image_id} | Update Card Image Details
[**UploadCardImageData**](CardsApi.md#UploadCardImageData) | **Post** /cards/images/{card_image_id}/data | Upload Card Image

# **ActivateCard**
> CardResponse ActivateCard(ctx, body)
Activate a card

Activate a card 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardActivationRequest**](CardActivationRequest.md)| Card activation code | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCardImage**
> CardImageDetails CreateCardImage(ctx, body)
Create Card Image

Create a card image entity. Note that this does not include the image data itself. You can upload the image data via a subsequent uploadCardImageData request using the ID created here. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCardImageRequest**](CreateCardImageRequest.md)| Details of the image to create | 

### Return type

[**CardImageDetails**](card_image_details.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCard**
> CardResponse GetCard(ctx, cardId)
Get Card

Get the details about a card that has been issued 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | [**string**](.md)|  | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardBarcode**
> InlineResponse200 GetCardBarcode(ctx, cardId)
Get Card Barcode

This endpoint is for testing environment only to provide access to barcode of a test card 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardImageData**
> *os.File GetCardImageData(ctx, cardImageId)
Get Card Image Data

Get card image data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardImageId** | [**string**](.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/jpeg, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardImageDetails**
> CardImageDetails GetCardImageDetails(ctx, cardImageId)
Get Card Image Details

Get card image details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardImageId** | [**string**](.md)|  | 

### Return type

[**CardImageDetails**](card_image_details.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCardWidgetURL**
> CardWidgetUrlResponse GetCardWidgetURL(ctx, widgetType, customerId, accountId, optional)
Get card widget URL

This endpoint returns a URL address of the specified widget for a given card 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **widgetType** | [**WidgetType**](.md)| The type of widget for which to construct the URL | 
  **customerId** | [**string**](.md)|  | 
  **accountId** | [**string**](.md)|  | 
 **optional** | ***CardsApiGetCardWidgetURLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiGetCardWidgetURLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cardId** | [**optional.Interface of string**](.md)| The ID of the card (required for set PIN widget) | 

### Return type

[**CardWidgetUrlResponse**](card_widget_url_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientAccessToken**
> ClientToken GetClientAccessToken(ctx, cardId)
Get a client token

Create a client access token for interacting with a card.  This token will be used on the client to identify the card for flows like viewing Full PAN or setting the PIN in a PCI compliant manner. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | [**string**](.md)|  | 

### Return type

[**ClientToken**](client_token.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientSingleUseToken**
> SingleUseTokenResponse GetClientSingleUseToken(ctx, body)
Get single-use token

This endpoint returns a single-use access token. This type of token authorizes a single request to access API endpoints and data associated with a particular user 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SingleUseTokenRequest**](SingleUseTokenRequest.md)| User token details | 

### Return type

[**SingleUseTokenResponse**](single_use_token_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueCard**
> CardResponse IssueCard(ctx, body)
Issue a Card

Issue or reissue a new card for a customer 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardIssuanceRequest**](CardIssuanceRequest.md)| Card to issue | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCardImageDetails**
> CardImageDetailsList ListCardImageDetails(ctx, customerId)
List Card Image Details

List all card image details 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | [**string**](.md)|  | 

### Return type

[**CardImageDetailsList**](card_image_details_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCardProducts**
> CardProductListResponse ListCardProducts(ctx, optional)
List Card Products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardsApiListCardProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiListCardProductsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 

### Return type

[**CardProductListResponse**](card_product_list_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCards**
> CardListResponse ListCards(ctx, optional)
List Cards

List of cards matching query parameters 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardsApiListCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiListCardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | [**optional.Interface of string**](.md)|  | 
 **accountId** | [**optional.Interface of string**](.md)|  | 
 **embossName** | **optional.String**| emboss name | 
 **lastFour** | **optional.String**| The last 4 digits of the card PAN | 
 **expirationDate** | **optional.String**| The date representing when the card would expire at | 
 **cardType** | **optional.String**| Indicates the type of card | 
 **cardBrand** | [**optional.Interface of CardBrand**](.md)| The brand of a card product | 
 **form** | [**optional.Interface of Form**](.md)| The format of the card | 
 **cardProductId** | [**optional.Interface of string**](.md)|  | 
 **cardStatus** | [**optional.Interface of CardStatus**](.md)| The status of a card | 
 **postalCode** | **optional.String**| The postal code of a card user | 
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Specifies the sort order for the returned cards.  | 

### Return type

[**CardListResponse**](card_list_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListChanges**
> CardChangesList ListChanges(ctx, cardId)
List Card Changes

List card change history 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | [**string**](.md)|  | 

### Return type

[**CardChangesList**](card_changes_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCard**
> CardResponse UpdateCard(ctx, body, cardId)
Update Card

Integrators can update the card resource to change status, update shipping (if the card hasn't been shipped) or edit metadata. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CardEditRequest**](CardEditRequest.md)| Card edits | 
  **cardId** | [**string**](.md)|  | 

### Return type

[**CardResponse**](card_response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCardImageDetails**
> CardImageDetails UpdateCardImageDetails(ctx, body, cardImageId)
Update Card Image Details

Update card image details. The only detail that can be updated is the card status as APPROVED or REJECTED.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCardImageRequest**](UpdateCardImageRequest.md)| Details of the image to create | 
  **cardImageId** | [**string**](.md)|  | 

### Return type

[**CardImageDetails**](card_image_details.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadCardImageData**
> CardImageDetails UploadCardImageData(ctx, body, cardImageId)
Upload Card Image

Upload card image data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Object**](Object.md)| Binary image data | 
  **cardImageId** | [**string**](.md)|  | 

### Return type

[**CardImageDetails**](card_image_details.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: image/jpeg
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

