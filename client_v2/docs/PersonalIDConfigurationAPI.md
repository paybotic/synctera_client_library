# \PersonalIDConfigurationAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePersonalIdConfiguration**](PersonalIDConfigurationAPI.md#CreatePersonalIdConfiguration) | **Post** /personal_id_configurations | Create a personal ID configuration
[**DeletePersonalIdConfiguration**](PersonalIDConfigurationAPI.md#DeletePersonalIdConfiguration) | **Delete** /personal_id_configurations/{id} | Delete Personal ID Configuration
[**GetPersonalIdConfiguration**](PersonalIDConfigurationAPI.md#GetPersonalIdConfiguration) | **Get** /personal_id_configurations/{id} | Get Personal ID Configuration
[**ListPersonalIdConfigurations**](PersonalIDConfigurationAPI.md#ListPersonalIdConfigurations) | **Get** /personal_id_configurations | List personal ID configurations



## CreatePersonalIdConfiguration

> PersonalIdConfigurationResponse CreatePersonalIdConfiguration(ctx).PersonalIdConfigurationCreate(personalIdConfigurationCreate).Execute()

Create a personal ID configuration



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
	personalIdConfigurationCreate := *openapiclient.NewPersonalIdConfigurationCreate("-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...") // PersonalIdConfigurationCreate | Personal ID configuration to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalIDConfigurationAPI.CreatePersonalIdConfiguration(context.Background()).PersonalIdConfigurationCreate(personalIdConfigurationCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalIDConfigurationAPI.CreatePersonalIdConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePersonalIdConfiguration`: PersonalIdConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `PersonalIDConfigurationAPI.CreatePersonalIdConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePersonalIdConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personalIdConfigurationCreate** | [**PersonalIdConfigurationCreate**](PersonalIdConfigurationCreate.md) | Personal ID configuration to create | 

### Return type

[**PersonalIdConfigurationResponse**](PersonalIdConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePersonalIdConfiguration

> DeleteResponse DeletePersonalIdConfiguration(ctx, id).Execute()

Delete Personal ID Configuration



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
	id := "64438afd-fa20-4010-a573-2bbdca77cdb6" // string | The unique identifier of a personal ID configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalIDConfigurationAPI.DeletePersonalIdConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalIDConfigurationAPI.DeletePersonalIdConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePersonalIdConfiguration`: DeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `PersonalIDConfigurationAPI.DeletePersonalIdConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of a personal ID configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonalIdConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonalIdConfiguration

> PersonalIdConfigurationResponse GetPersonalIdConfiguration(ctx, id).Execute()

Get Personal ID Configuration



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
	id := "64438afd-fa20-4010-a573-2bbdca77cdb6" // string | The unique identifier of a personal ID configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalIDConfigurationAPI.GetPersonalIdConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalIDConfigurationAPI.GetPersonalIdConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersonalIdConfiguration`: PersonalIdConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `PersonalIDConfigurationAPI.GetPersonalIdConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of a personal ID configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalIdConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PersonalIdConfigurationResponse**](PersonalIdConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPersonalIdConfigurations

> PersonalIdConfigurationList ListPersonalIdConfigurations(ctx).BankId(bankId).PartnerId(partnerId).Limit(limit).PageToken(pageToken).Execute()

List personal ID configurations



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
	bankId := int32(6) // int32 | Filter by bank ID (optional)
	partnerId := int32(6) // int32 | Filter by partner ID (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalIDConfigurationAPI.ListPersonalIdConfigurations(context.Background()).BankId(bankId).PartnerId(partnerId).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalIDConfigurationAPI.ListPersonalIdConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPersonalIdConfigurations`: PersonalIdConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `PersonalIDConfigurationAPI.ListPersonalIdConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPersonalIdConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankId** | **int32** | Filter by bank ID | 
 **partnerId** | **int32** | Filter by partner ID | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**PersonalIdConfigurationList**](PersonalIdConfigurationList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

