# GatewayPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomHeaders** | Pointer to **map[string][]string** | Optional parameter that allows to configure custom http headers for the Auth request to Gateway URL if needed | [optional] 
**Disabled** | Pointer to **bool** | Setting this parameter to &#39;true&#39; allows create Gateway Config as inactive ( can be useful as a preparation step) | [optional] [default to false]
**MaxWaitMs** | Pointer to **int32** | Optional parameter that configures the maximum amount of time in milliseconds that we will wait for the response from Gateway URL request. Default value is used if empty | [optional] [default to 1500]
**Url** | **string** | The URL address which will be used for the ACH in Auth Flow requests to get authorization from the fintech to process the transaction | 

## Methods

### NewGatewayPost

`func NewGatewayPost(url string, ) *GatewayPost`

NewGatewayPost instantiates a new GatewayPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayPostWithDefaults

`func NewGatewayPostWithDefaults() *GatewayPost`

NewGatewayPostWithDefaults instantiates a new GatewayPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomHeaders

`func (o *GatewayPost) GetCustomHeaders() map[string][]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *GatewayPost) GetCustomHeadersOk() (*map[string][]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *GatewayPost) SetCustomHeaders(v map[string][]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *GatewayPost) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetDisabled

`func (o *GatewayPost) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GatewayPost) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GatewayPost) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GatewayPost) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetMaxWaitMs

`func (o *GatewayPost) GetMaxWaitMs() int32`

GetMaxWaitMs returns the MaxWaitMs field if non-nil, zero value otherwise.

### GetMaxWaitMsOk

`func (o *GatewayPost) GetMaxWaitMsOk() (*int32, bool)`

GetMaxWaitMsOk returns a tuple with the MaxWaitMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitMs

`func (o *GatewayPost) SetMaxWaitMs(v int32)`

SetMaxWaitMs sets MaxWaitMs field to given value.

### HasMaxWaitMs

`func (o *GatewayPost) HasMaxWaitMs() bool`

HasMaxWaitMs returns a boolean if a field has been set.

### GetUrl

`func (o *GatewayPost) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GatewayPost) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GatewayPost) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


