# GatewayConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Time when Gateway Config object was created | [optional] 
**CustomHeaders** | Pointer to **map[string][]string** | Optional parameter that allows to configure custom http headers for the Auth request to Gateway URL if needed | [optional] 
**Disabled** | Pointer to **bool** | Indicates if the Gateway Config is active for Auth Flow for the current Fintech (Tenant) | [optional] [default to false]
**Id** | Pointer to **string** | Identifier of the Gateway Config object | [optional] 
**MaxWaitMs** | Pointer to **int32** | Shows maximum amount of time in milliseconds that we will wait for the response from Gateway URL Auth request | [optional] [default to 1500]
**Tenant** | **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | 
**Updated** | Pointer to **time.Time** | Time when Gateway Config object was updated | [optional] 
**Url** | Pointer to **string** | The URL address which will be used for the ACH in Auth Flow requests to get authorization from the fintech to process the transaction | [optional] 

## Methods

### NewGatewayConfig

`func NewGatewayConfig(tenant string, ) *GatewayConfig`

NewGatewayConfig instantiates a new GatewayConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayConfigWithDefaults

`func NewGatewayConfigWithDefaults() *GatewayConfig`

NewGatewayConfigWithDefaults instantiates a new GatewayConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GatewayConfig) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GatewayConfig) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GatewayConfig) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GatewayConfig) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *GatewayConfig) GetCustomHeaders() map[string][]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *GatewayConfig) GetCustomHeadersOk() (*map[string][]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *GatewayConfig) SetCustomHeaders(v map[string][]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *GatewayConfig) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetDisabled

`func (o *GatewayConfig) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GatewayConfig) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GatewayConfig) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GatewayConfig) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetId

`func (o *GatewayConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxWaitMs

`func (o *GatewayConfig) GetMaxWaitMs() int32`

GetMaxWaitMs returns the MaxWaitMs field if non-nil, zero value otherwise.

### GetMaxWaitMsOk

`func (o *GatewayConfig) GetMaxWaitMsOk() (*int32, bool)`

GetMaxWaitMsOk returns a tuple with the MaxWaitMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitMs

`func (o *GatewayConfig) SetMaxWaitMs(v int32)`

SetMaxWaitMs sets MaxWaitMs field to given value.

### HasMaxWaitMs

`func (o *GatewayConfig) HasMaxWaitMs() bool`

HasMaxWaitMs returns a boolean if a field has been set.

### GetTenant

`func (o *GatewayConfig) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *GatewayConfig) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *GatewayConfig) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetUpdated

`func (o *GatewayConfig) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GatewayConfig) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GatewayConfig) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GatewayConfig) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *GatewayConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GatewayConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GatewayConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GatewayConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


