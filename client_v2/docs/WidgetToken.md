# WidgetToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**WidgetToken** | **string** | A short-lived, one-time token for use with Synctera widgets | 

## Methods

### NewWidgetToken

`func NewWidgetToken(tenant string, widgetToken string, ) *WidgetToken`

NewWidgetToken instantiates a new WidgetToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetTokenWithDefaults

`func NewWidgetTokenWithDefaults() *WidgetToken`

NewWidgetTokenWithDefaults instantiates a new WidgetToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *WidgetToken) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WidgetToken) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WidgetToken) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetWidgetToken

`func (o *WidgetToken) GetWidgetToken() string`

GetWidgetToken returns the WidgetToken field if non-nil, zero value otherwise.

### GetWidgetTokenOk

`func (o *WidgetToken) GetWidgetTokenOk() (*string, bool)`

GetWidgetTokenOk returns a tuple with the WidgetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetToken

`func (o *WidgetToken) SetWidgetToken(v string)`

SetWidgetToken sets WidgetToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


