# GatewayPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomHeaders** | Pointer to **map[string][]string** | Custom http headers for the Auth request to Gateway URL | [optional] 
**Disabled** | Pointer to **bool** | Allows to disable/enable Gateway Config for the Fintech (Tenant) | [optional] 
**MaxWaitMs** | Pointer to **int32** | Maximum amount of time in milliseconds that we will wait for the response from Gateway URL request | [optional] 
**Url** | Pointer to **string** | The URL address which will be used for the ACH in Auth Flow requests to get authorization from the fintech to process the transaction | [optional] 

## Methods

### NewGatewayPatch

`func NewGatewayPatch() *GatewayPatch`

NewGatewayPatch instantiates a new GatewayPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPatchWithDefaults

`func NewGatewayPatchWithDefaults() *GatewayPatch`

NewGatewayPatchWithDefaults instantiates a new GatewayPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomHeaders

`func (o *GatewayPatch) GetCustomHeaders() map[string][]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *GatewayPatch) GetCustomHeadersOk() (*map[string][]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *GatewayPatch) SetCustomHeaders(v map[string][]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *GatewayPatch) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### SetCustomHeadersNil

`func (o *GatewayPatch) SetCustomHeadersNil(b bool)`

 SetCustomHeadersNil sets the value for CustomHeaders to be an explicit nil

### UnsetCustomHeaders
`func (o *GatewayPatch) UnsetCustomHeaders()`

UnsetCustomHeaders ensures that no value is present for CustomHeaders, not even an explicit nil
### GetDisabled

`func (o *GatewayPatch) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GatewayPatch) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GatewayPatch) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GatewayPatch) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMaxWaitMs

`func (o *GatewayPatch) GetMaxWaitMs() int32`

GetMaxWaitMs returns the MaxWaitMs field if non-nil, zero value otherwise.

### GetMaxWaitMsOk

`func (o *GatewayPatch) GetMaxWaitMsOk() (*int32, bool)`

GetMaxWaitMsOk returns a tuple with the MaxWaitMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitMs

`func (o *GatewayPatch) SetMaxWaitMs(v int32)`

SetMaxWaitMs sets MaxWaitMs field to given value.

### HasMaxWaitMs

`func (o *GatewayPatch) HasMaxWaitMs() bool`

HasMaxWaitMs returns a boolean if a field has been set.

### GetUrl

`func (o *GatewayPatch) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GatewayPatch) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GatewayPatch) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GatewayPatch) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


