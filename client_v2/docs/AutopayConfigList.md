# AutopayConfigList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**AutopayConfigs** | [**[]AutopayConfig**](AutopayConfig.md) | Array of autopay configurations | 

## Methods

### NewAutopayConfigList

`func NewAutopayConfigList(autopayConfigs []AutopayConfig, ) *AutopayConfigList`

NewAutopayConfigList instantiates a new AutopayConfigList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayConfigListWithDefaults

`func NewAutopayConfigListWithDefaults() *AutopayConfigList`

NewAutopayConfigListWithDefaults instantiates a new AutopayConfigList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *AutopayConfigList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AutopayConfigList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AutopayConfigList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AutopayConfigList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetAutopayConfigs

`func (o *AutopayConfigList) GetAutopayConfigs() []AutopayConfig`

GetAutopayConfigs returns the AutopayConfigs field if non-nil, zero value otherwise.

### GetAutopayConfigsOk

`func (o *AutopayConfigList) GetAutopayConfigsOk() (*[]AutopayConfig, bool)`

GetAutopayConfigsOk returns a tuple with the AutopayConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopayConfigs

`func (o *AutopayConfigList) SetAutopayConfigs(v []AutopayConfig)`

SetAutopayConfigs sets AutopayConfigs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


