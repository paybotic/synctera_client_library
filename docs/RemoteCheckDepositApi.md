# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRdcDeposit**](RemoteCheckDepositApi.md#CreateRdcDeposit) | **Post** /rdc/deposits | Create a Remote Check Deposit
[**GetRdcDeposit**](RemoteCheckDepositApi.md#GetRdcDeposit) | **Get** /rdc/deposits/{deposit_id} | Get Remote Check Deposit
[**ListRdcDeposits**](RemoteCheckDepositApi.md#ListRdcDeposits) | **Get** /rdc/deposits | List Remote Check Deposits

# **CreateRdcDeposit**
> Deposit CreateRdcDeposit(ctx, optional)
Create a Remote Check Deposit

Create a new deposit using remote deposit capture to an account 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemoteCheckDepositApiCreateRdcDepositOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoteCheckDepositApiCreateRdcDepositOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Deposit**](Deposit.md)| Attributes of the Remote Check Deposit to create | 

### Return type

[**Deposit**](deposit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRdcDeposit**
> Deposit GetRdcDeposit(ctx, depositId)
Get Remote Check Deposit

Retrieves one deposit made using remote deposit capture associated with an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **depositId** | [**string**](.md)| ID of a deposit for a remote deposit capture | 

### Return type

[**Deposit**](deposit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRdcDeposits**
> DepositList ListRdcDeposits(ctx, optional)
List Remote Check Deposits

Retrieves a paginated list of the deposits made using remote deposit capture associated with an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RemoteCheckDepositApiListRdcDepositsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoteCheckDepositApiListRdcDepositsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 100]
 **pageToken** | **optional.String**|  | 

### Return type

[**DepositList**](deposit_list.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

