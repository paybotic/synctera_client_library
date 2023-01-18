# \ExternalCardsAlphaApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalCardFromToken**](ExternalCardsAlphaApi.md#CreateExternalCardFromToken) | **Post** /external_cards/tokens | Create External Card from token
[**CreateExternalCardTransfer**](ExternalCardsAlphaApi.md#CreateExternalCardTransfer) | **Post** /external_cards/transfers | Create External Card Transfer
[**CreateExternalCardTransferReversal**](ExternalCardsAlphaApi.md#CreateExternalCardTransferReversal) | **Post** /external_cards/transfers/{transfer_id}/reversals | Create External Card Transfer Reversal
[**GetExternalCard**](ExternalCardsAlphaApi.md#GetExternalCard) | **Get** /external_cards/{external_card_id} | Get External Card
[**GetExternalCardTransfer**](ExternalCardsAlphaApi.md#GetExternalCardTransfer) | **Get** /external_cards/transfers/{transfer_id} | Get External Card Transfer
[**ListExternalCardTransfers**](ExternalCardsAlphaApi.md#ListExternalCardTransfers) | **Get** /external_cards/transfers | List External Card Transfers
[**ListExternalCards**](ExternalCardsAlphaApi.md#ListExternalCards) | **Get** /external_cards | List External Cards



## CreateExternalCardFromToken

> ExternalCardResponse CreateExternalCardFromToken(ctx).ExternalCardRequest(externalCardRequest).Execute()

Create External Card from token



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
    externalCardRequest := *openapiclient.NewExternalCardRequest("7d943c51-e4ff-4e57-9558-08cab6b963c7", "Jean Valjean", "Token_example") // ExternalCardRequest | Details of the External Card to create

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
 **externalCardRequest** | [**ExternalCardRequest**](ExternalCardRequest.md) | Details of the External Card to create | 

### Return type

[**ExternalCardResponse**](ExternalCardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalCardTransfer

> TransferResponse CreateExternalCardTransfer(ctx).TransferRequest(transferRequest).Execute()

Create External Card Transfer



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
    transferRequest := *openapiclient.NewTransferRequest(int32(123), "USD", "7d943c51-e4ff-4e57-9558-08cab6b963c7", "7d943c51-e4ff-4e57-9558-08cab6b963c7", openapiclient.transfer_type_request("PULL")) // TransferRequest | Details of the External Card Transfer to create

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
 **transferRequest** | [**TransferRequest**](TransferRequest.md) | Details of the External Card Transfer to create | 

### Return type

[**TransferResponse**](TransferResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalCardTransferReversal

> TransferResponse CreateExternalCardTransferReversal(ctx, transferId).TransferReversalRequest(transferReversalRequest).Execute()

Create External Card Transfer Reversal



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
    transferId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique identifier of a transfer
    transferReversalRequest := *openapiclient.NewTransferReversalRequest(int32(123), "USD") // TransferReversalRequest | Details of the External Card Transfer Reversal to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalCardsAlphaApi.CreateExternalCardTransferReversal(context.Background(), transferId).TransferReversalRequest(transferReversalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalCardsAlphaApi.CreateExternalCardTransferReversal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExternalCardTransferReversal`: TransferResponse
    fmt.Fprintf(os.Stdout, "Response from `ExternalCardsAlphaApi.CreateExternalCardTransferReversal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | The unique identifier of a transfer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalCardTransferReversalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferReversalRequest** | [**TransferReversalRequest**](TransferReversalRequest.md) | Details of the External Card Transfer Reversal to create | 

### Return type

[**TransferResponse**](TransferResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalCard

> ExternalCardResponse GetExternalCard(ctx, externalCardId).Execute()

Get External Card



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
    externalCardId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique identifier of an external card

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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalCardTransfer

> TransferResponse GetExternalCardTransfer(ctx, transferId).Execute()

Get External Card Transfer



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
    transferId := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | The unique identifier of a transfer

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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalCardTransfers

> TransferListResponse ListExternalCardTransfers(ctx).Limit(limit).PageToken(pageToken).CustomerId(customerId).ExternalCardId(externalCardId).OriginatingAccountId(originatingAccountId).Type_(type_).Execute()

List External Card Transfers



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
    pageToken := "a8937a0d" // string |  (optional)
    customerId := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string |  (optional)
    externalCardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    originatingAccountId := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string |  (optional)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalCards

> ExternalCardListResponse ListExternalCards(ctx).Limit(limit).PageToken(pageToken).CustomerId(customerId).BusinessId(businessId).Execute()

List External Cards



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
    pageToken := "a8937a0d" // string |  (optional)
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    businessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalCardsAlphaApi.ListExternalCards(context.Background()).Limit(limit).PageToken(pageToken).CustomerId(customerId).BusinessId(businessId).Execute()
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
 **businessId** | **string** |  | 

### Return type

[**ExternalCardListResponse**](ExternalCardListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

