# {{classname}}

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WipeWorkspace**](SandboxWipeAlphaApi.md#WipeWorkspace) | **Post** /wipe | Delete data

# **WipeWorkspace**
> WipeWorkspace(ctx, )
Delete data

Delete the customer and account related data, leaving other configuration data intact. This enables use cases such as bulk data deletion between tests. Data associated with below resources will be deleted:   - Accounts   - Account applications   - ACH   - Businesses   - Cards   - Card images   - Cases   - Customers   - Disclosures   - External Accounts   - Internal Accounts   - Payment schedules and history   - Persons   - RDC   - Relationships   - Transactions (including for Internal Accounts)   - Verifications  Data associated with below resources will be retained:   - Account Templates   - API Keys   - Bank/Partner data   - Card product   - Disclosure document records   - Egress config   - Groups   - PII contract with vault   - Roles   - Users   - Webhooks 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

