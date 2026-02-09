# \InStoreCashDepositSimulationsBetaAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBarcodeDepositSimulation**](InStoreCashDepositSimulationsBetaAPI.md#CreateBarcodeDepositSimulation) | **Post** /cash/transaction_simulations/barcodes/deposits | Create a cash deposit transaction simulation for a barcode
[**GetBarcodeSimulationStore**](InStoreCashDepositSimulationsBetaAPI.md#GetBarcodeSimulationStore) | **Get** /cash/transaction_simulations/barcodes/stores | Retrieve Store Information for Barcode Simulation



## CreateBarcodeDepositSimulation

> CreateBarcodeDepositSimulation(ctx).BarcodeDepositSimulationRequest(barcodeDepositSimulationRequest).Execute()

Create a cash deposit transaction simulation for a barcode



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	barcodeDepositSimulationRequest := *openapiclient.NewBarcodeDepositSimulationRequest(float32(100), "550e8400-e29b-41d4-a716-446655440000", "Status_example", "StoreId_example", "Type_example") // BarcodeDepositSimulationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InStoreCashDepositSimulationsBetaAPI.CreateBarcodeDepositSimulation(context.Background()).BarcodeDepositSimulationRequest(barcodeDepositSimulationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InStoreCashDepositSimulationsBetaAPI.CreateBarcodeDepositSimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBarcodeDepositSimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **barcodeDepositSimulationRequest** | [**BarcodeDepositSimulationRequest**](BarcodeDepositSimulationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBarcodeSimulationStore

> BarcodeSimulationStoreResponse GetBarcodeSimulationStore(ctx).Lat(lat).Lng(lng).Radius(radius).Limit(limit).Execute()

Retrieve Store Information for Barcode Simulation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	lat := float32(37.7749) // float32 | 
	lng := float32(-122.4194) // float32 | 
	radius := float32(3.4) // float32 |  (optional) (default to 5)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InStoreCashDepositSimulationsBetaAPI.GetBarcodeSimulationStore(context.Background()).Lat(lat).Lng(lng).Radius(radius).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InStoreCashDepositSimulationsBetaAPI.GetBarcodeSimulationStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBarcodeSimulationStore`: BarcodeSimulationStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `InStoreCashDepositSimulationsBetaAPI.GetBarcodeSimulationStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBarcodeSimulationStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lat** | **float32** |  | 
 **lng** | **float32** |  | 
 **radius** | **float32** |  | [default to 5]
 **limit** | **int32** |  | [default to 50]

### Return type

[**BarcodeSimulationStoreResponse**](BarcodeSimulationStoreResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

