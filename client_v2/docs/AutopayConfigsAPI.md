# \AutopayConfigsAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutopayConfig**](AutopayConfigsAPI.md#CreateAutopayConfig) | **Post** /autopay_configs | Create autopay configuration
[**DeleteAutopayConfig**](AutopayConfigsAPI.md#DeleteAutopayConfig) | **Delete** /autopay_configs/{autopay_config_id} | Delete autopay configuration
[**GetAutopayConfig**](AutopayConfigsAPI.md#GetAutopayConfig) | **Get** /autopay_configs/{autopay_config_id} | Get autopay configuration
[**ListAutopayConfigs**](AutopayConfigsAPI.md#ListAutopayConfigs) | **Get** /autopay_configs | List autopay configurations
[**UpdateAutopayConfig**](AutopayConfigsAPI.md#UpdateAutopayConfig) | **Patch** /autopay_configs/{autopay_config_id} | Update autopay configuration



## CreateAutopayConfig

> AutopayConfig CreateAutopayConfig(ctx).AutopayConfigCreateRequest(autopayConfigCreateRequest).Execute()

Create autopay configuration



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
	autopayConfigCreateRequest := *openapiclient.NewAutopayConfigCreateRequest(*openapiclient.NewAutopayConfigData(openapiclient.autopay_amount_rule("MINIMUM_DUE"), openapiclient.autopay_failure_policy("NO_RETRY"), openapiclient.autopay_payment_method("ACH"), openapiclient.autopay_timing_rule("ON_DUE_DATE")), "7d943c51-e4ff-4e57-9558-08cab6b963c7") // AutopayConfigCreateRequest | Autopay configuration to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutopayConfigsAPI.CreateAutopayConfig(context.Background()).AutopayConfigCreateRequest(autopayConfigCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutopayConfigsAPI.CreateAutopayConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAutopayConfig`: AutopayConfig
	fmt.Fprintf(os.Stdout, "Response from `AutopayConfigsAPI.CreateAutopayConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutopayConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autopayConfigCreateRequest** | [**AutopayConfigCreateRequest**](AutopayConfigCreateRequest.md) | Autopay configuration to create | 

### Return type

[**AutopayConfig**](AutopayConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutopayConfig

> DeleteAutopayConfig(ctx, autopayConfigId).Execute()

Delete autopay configuration



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
	autopayConfigId := "2d0f7601-ecc6-48b4-b82f-5d64fa84628b" // string | Autopay configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutopayConfigsAPI.DeleteAutopayConfig(context.Background(), autopayConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutopayConfigsAPI.DeleteAutopayConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autopayConfigId** | **string** | Autopay configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutopayConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutopayConfig

> AutopayConfig GetAutopayConfig(ctx, autopayConfigId).Execute()

Get autopay configuration



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
	autopayConfigId := "2d0f7601-ecc6-48b4-b82f-5d64fa84628b" // string | Autopay configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutopayConfigsAPI.GetAutopayConfig(context.Background(), autopayConfigId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutopayConfigsAPI.GetAutopayConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutopayConfig`: AutopayConfig
	fmt.Fprintf(os.Stdout, "Response from `AutopayConfigsAPI.GetAutopayConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autopayConfigId** | **string** | Autopay configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutopayConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutopayConfig**](AutopayConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutopayConfigs

> AutopayConfigList ListAutopayConfigs(ctx).Limit(limit).PageToken(pageToken).BusinessId(businessId).ExternalAccountId(externalAccountId).FromCreationTime(fromCreationTime).FromLastUpdatedTime(fromLastUpdatedTime).Id(id).LendingAccountId(lendingAccountId).PersonId(personId).SortBy(sortBy).SourceAccountId(sourceAccountId).Status(status).Tenant(tenant).ToCreationTime(toCreationTime).ToLastUpdatedTime(toLastUpdatedTime).Execute()

List autopay configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)
	businessId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	externalAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	fromCreationTime := time.Now() // time.Time |  (optional)
	fromLastUpdatedTime := time.Now() // time.Time |  (optional)
	id := []string{"Inner_example"} // []string |  (optional)
	lendingAccountId := []string{"Inner_example"} // []string |  (optional)
	personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	sortBy := []string{"SortBy_example"} // []string |  (optional)
	sourceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	status := openapiclient.autopay_config_status("ACTIVE") // AutopayConfigStatus |  (optional)
	tenant := []string{"abcdef_ghijkl"} // []string |  (optional)
	toCreationTime := time.Now() // time.Time |  (optional)
	toLastUpdatedTime := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutopayConfigsAPI.ListAutopayConfigs(context.Background()).Limit(limit).PageToken(pageToken).BusinessId(businessId).ExternalAccountId(externalAccountId).FromCreationTime(fromCreationTime).FromLastUpdatedTime(fromLastUpdatedTime).Id(id).LendingAccountId(lendingAccountId).PersonId(personId).SortBy(sortBy).SourceAccountId(sourceAccountId).Status(status).Tenant(tenant).ToCreationTime(toCreationTime).ToLastUpdatedTime(toLastUpdatedTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutopayConfigsAPI.ListAutopayConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAutopayConfigs`: AutopayConfigList
	fmt.Fprintf(os.Stdout, "Response from `AutopayConfigsAPI.ListAutopayConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAutopayConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **businessId** | **string** |  | 
 **externalAccountId** | **string** |  | 
 **fromCreationTime** | **time.Time** |  | 
 **fromLastUpdatedTime** | **time.Time** |  | 
 **id** | **[]string** |  | 
 **lendingAccountId** | **[]string** |  | 
 **personId** | **string** |  | 
 **sortBy** | **[]string** |  | 
 **sourceAccountId** | **string** |  | 
 **status** | [**AutopayConfigStatus**](AutopayConfigStatus.md) |  | 
 **tenant** | **[]string** |  | 
 **toCreationTime** | **time.Time** |  | 
 **toLastUpdatedTime** | **time.Time** |  | 

### Return type

[**AutopayConfigList**](AutopayConfigList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutopayConfig

> AutopayConfig UpdateAutopayConfig(ctx, autopayConfigId).AutopayConfigUpdateRequest(autopayConfigUpdateRequest).Execute()

Update autopay configuration



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
	autopayConfigId := "2d0f7601-ecc6-48b4-b82f-5d64fa84628b" // string | Autopay configuration ID
	autopayConfigUpdateRequest := *openapiclient.NewAutopayConfigUpdateRequest() // AutopayConfigUpdateRequest | Fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutopayConfigsAPI.UpdateAutopayConfig(context.Background(), autopayConfigId).AutopayConfigUpdateRequest(autopayConfigUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutopayConfigsAPI.UpdateAutopayConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutopayConfig`: AutopayConfig
	fmt.Fprintf(os.Stdout, "Response from `AutopayConfigsAPI.UpdateAutopayConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**autopayConfigId** | **string** | Autopay configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutopayConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autopayConfigUpdateRequest** | [**AutopayConfigUpdateRequest**](AutopayConfigUpdateRequest.md) | Fields to update | 

### Return type

[**AutopayConfig**](AutopayConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

