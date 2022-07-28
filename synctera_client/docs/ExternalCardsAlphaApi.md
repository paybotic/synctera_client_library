# \ExternalCardsAlphaApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalCardFromToken**](ExternalCardsAlphaApi.md#CreateExternalCardFromToken) | **Post** /external_cards/tokens | Create external card from token
[**CreateExternalCardTransfer**](ExternalCardsAlphaApi.md#CreateExternalCardTransfer) | **Post** /external_cards/transfers | Create external card transfer
[**GetExternalCard**](ExternalCardsAlphaApi.md#GetExternalCard) | **Get** /external_cards/{external_card_id} | Get a external card
[**GetExternalCardTransfer**](ExternalCardsAlphaApi.md#GetExternalCardTransfer) | **Get** /external_cards/transfers/{transfer_id} | Get an external card transfer
[**ListExternalCardTransfers**](ExternalCardsAlphaApi.md#ListExternalCardTransfers) | **Get** /external_cards/transfers | List external transfers
[**ListExternalCards**](ExternalCardsAlphaApi.md#ListExternalCards) | **Get** /external_cards | List external cards



## CreateExternalCardFromToken

> ExternalCardResponse CreateExternalCardFromToken(ctx).ExternalCardRequest(externalCardRequest).Execute()

Create external card from token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    externalCardRequest := *openapiclient.NewExternalCardRequest("37a054b8-fbdb-44c6-ae20-091f94ada475", "Token_example") // ExternalCardRequest | Details of the card to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalCardsAlphaApi.CreateExternalCardFromToken(context.Background()).ExternalCardRequest(externalCardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalCardsAlphaApi.CreateExternalCardFromToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExternalCardFromToken`: ExternalCardResponse
    fmt.Fprintf(os.Stdout, "Response from `ExternalCardsAlphaApi.CreateExternalCardFromToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalCardFromTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalCardRequest** | [**ExternalCardRequest**](ExternalCardRequest.md) | Details of the card to create | 

### Return type

[**ExternalCardResponse**](ExternalCardResponse.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## CreateExternalCardTransfer

> TransferResponse CreateExternalCardTransfer(ctx).TransferRequest(transferRequest).Execute()

Create external card transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferRequest := *openapiclient.NewTransferRequest(int32(123), "USD", "b542b637-ce49-4be6-a26d-973f050985b9", "d3b3de70-661d-4c72-be44-3ca8d696d969", openapiclient.transfer_type("PULL")) // TransferRequest | Details of the transfer to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalCardsAlphaApi.CreateExternalCardTransfer(context.Background()).TransferRequest(transferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalCardsAlphaApi.CreateExternalCardTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExternalCardTransfer`: TransferResponse
    fmt.Fprintf(os.Stdout, "Response from `ExternalCardsAlphaApi.CreateExternalCardTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalCardTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferRequest** | [**TransferRequest**](TransferRequest.md) | Details of the transfer to create | 

### Return type

[**TransferResponse**](TransferResponse.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## GetExternalCard

> ExternalCardResponse GetExternalCard(ctx, externalCardId).Execute()

Get a external card



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    externalCardId := "4675ebf0-0691-4a2b-b1db-7ca6f4ff9ec5" // string | The unique identifier of an external card

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalCardsAlphaApi.GetExternalCard(context.Background(), externalCardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalCardsAlphaApi.GetExternalCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalCard`: ExternalCardResponse
    fmt.Fprintf(os.Stdout, "Response from `ExternalCardsAlphaApi.GetExternalCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalCardId** | **string** | The unique identifier of an external card | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalCardResponse**](ExternalCardResponse.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## GetExternalCardTransfer

> TransferResponse GetExternalCardTransfer(ctx, transferId).Execute()

Get an external card transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferId := "ddcacaa4-e0e4-4652-ae9f-5ef7f1b7d7e1" // string | The unique identifier of a transfer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalCardsAlphaApi.GetExternalCardTransfer(context.Background(), transferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalCardsAlphaApi.GetExternalCardTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalCardTransfer`: TransferResponse
    fmt.Fprintf(os.Stdout, "Response from `ExternalCardsAlphaApi.GetExternalCardTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | The unique identifier of a transfer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalCardTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransferResponse**](TransferResponse.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## ListExternalCardTransfers

> TransferListResponse ListExternalCardTransfers(ctx).Limit(limit).PageToken(pageToken).CustomerId(customerId).ExternalCardId(externalCardId).OriginatingAccountId(originatingAccountId).Type_(type_).Execute()

List external transfers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "xwsfu1mkaq" // string |  (optional)
    customerId := []string{"37a054b8-fbdb-44c6-ae20-091f94ada475"} // []string |  (optional)
    externalCardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    originatingAccountId := []string{"e85e7aeb-712d-4e9c-b2eb-d65fce19507b"} // []string |  (optional)
    type_ := openapiclient.transfer_type("PULL") // TransferType | The type of an external transfer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalCardsAlphaApi.ListExternalCardTransfers(context.Background()).Limit(limit).PageToken(pageToken).CustomerId(customerId).ExternalCardId(externalCardId).OriginatingAccountId(originatingAccountId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalCardsAlphaApi.ListExternalCardTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExternalCardTransfers`: TransferListResponse
    fmt.Fprintf(os.Stdout, "Response from `ExternalCardsAlphaApi.ListExternalCardTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExternalCardTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **customerId** | **[]string** |  | 
 **externalCardId** | **string** |  | 
 **originatingAccountId** | **[]string** |  | 
 **type_** | [**TransferType**](TransferType.md) | The type of an external transfer | 

### Return type

[**TransferListResponse**](TransferListResponse.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)


## ListExternalCards

> ExternalCardListResponse ListExternalCards(ctx).Limit(limit).PageToken(pageToken).CustomerId(customerId).Execute()

List external cards



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "xwsfu1mkaq" // string |  (optional)
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalCardsAlphaApi.ListExternalCards(context.Background()).Limit(limit).PageToken(pageToken).CustomerId(customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalCardsAlphaApi.ListExternalCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExternalCards`: ExternalCardListResponse
    fmt.Fprintf(os.Stdout, "Response from `ExternalCardsAlphaApi.ListExternalCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExternalCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **customerId** | **string** |  | 

### Return type

[**ExternalCardListResponse**](ExternalCardListResponse.md)

### Authorization

[bearerAuth](../../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../../README.md#documentation-for-models)
[[Back to README]](../../README.md)

