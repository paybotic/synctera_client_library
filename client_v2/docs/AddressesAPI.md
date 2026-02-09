# \AddressesAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddress**](AddressesAPI.md#CreateAddress) | **Post** /addresses | Create a address
[**GetAddress**](AddressesAPI.md#GetAddress) | **Get** /addresses/{address_id} | Get address information by id
[**ListAddresses**](AddressesAPI.md#ListAddresses) | **Get** /addresses | List Addresses
[**UpdateAddress**](AddressesAPI.md#UpdateAddress) | **Patch** /addresses/{address_id} | Update address information by id



## CreateAddress

> AddressResponse CreateAddress(ctx).AddressPost(addressPost).Execute()

Create a address



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
	addressPost := *openapiclient.NewAddressPost("100 Main St.", "New York", "US", "NY", "SHIPPING") // AddressPost | Address to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.CreateAddress(context.Background()).AddressPost(addressPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.CreateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAddress`: AddressResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.CreateAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressPost** | [**AddressPost**](AddressPost.md) | Address to create | 

### Return type

[**AddressResponse**](AddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddress

> AddressResponse GetAddress(ctx, addressId).Execute()

Get address information by id



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
	addressId := "1a582c51-e4ff-4e57-9558-08cab6b963aa" // string | ID of the address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddress(context.Background(), addressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddress`: AddressResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **string** | ID of the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressResponse**](AddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddresses

> AddressesList ListAddresses(ctx).CustomerId(customerId).BusinessId(businessId).AddressType(addressType).IsRegisteredAgent(isRegisteredAgent).IsActive(isActive).Country(country).Limit(limit).PageToken(pageToken).Execute()

List Addresses

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
	customerId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of a customer. (optional)
	businessId := "edebece0-4a06-4806-b062-7de25a848be8" // string | ID of the business that the risk score applies to (optional)
	addressType := []string{"AddressType_example"} // []string | Specifies the address type.  (optional)
	isRegisteredAgent := true // bool |  (optional)
	isActive := true // bool |  (optional)
	country := "U" // string | Country code.  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.ListAddresses(context.Background()).CustomerId(customerId).BusinessId(businessId).AddressType(addressType).IsRegisteredAgent(isRegisteredAgent).IsActive(isActive).Country(country).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.ListAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAddresses`: AddressesList
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.ListAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | The unique identifier of a customer. | 
 **businessId** | **string** | ID of the business that the risk score applies to | 
 **addressType** | **[]string** | Specifies the address type.  | 
 **isRegisteredAgent** | **bool** |  | 
 **isActive** | **bool** |  | 
 **country** | **string** | Country code.  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**AddressesList**](AddressesList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddress

> AddressResponse UpdateAddress(ctx, addressId).Body(body).Execute()

Update address information by id



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
	addressId := "1a582c51-e4ff-4e57-9558-08cab6b963aa" // string | ID of the address
	body := AddressBase(987) // AddressBase | Address patch details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.UpdateAddress(context.Background(), addressId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.UpdateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAddress`: AddressResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.UpdateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | **string** | ID of the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **AddressBase** | Address patch details | 

### Return type

[**AddressResponse**](AddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

