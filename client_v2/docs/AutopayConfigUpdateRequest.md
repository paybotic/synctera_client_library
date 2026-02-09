# AutopayConfigUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**AutopayConfigData**](AutopayConfigData.md) |  | [optional] 
**Status** | Pointer to [**AutopayConfigStatus**](AutopayConfigStatus.md) |  | [optional] 

## Methods

### NewAutopayConfigUpdateRequest

`func NewAutopayConfigUpdateRequest() *AutopayConfigUpdateRequest`

NewAutopayConfigUpdateRequest instantiates a new AutopayConfigUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayConfigUpdateRequestWithDefaults

`func NewAutopayConfigUpdateRequestWithDefaults() *AutopayConfigUpdateRequest`

NewAutopayConfigUpdateRequestWithDefaults instantiates a new AutopayConfigUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AutopayConfigUpdateRequest) GetConfig() AutopayConfigData`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AutopayConfigUpdateRequest) GetConfigOk() (*AutopayConfigData, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AutopayConfigUpdateRequest) SetConfig(v AutopayConfigData)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AutopayConfigUpdateRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *AutopayConfigUpdateRequest) GetStatus() AutopayConfigStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutopayConfigUpdateRequest) GetStatusOk() (*AutopayConfigStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutopayConfigUpdateRequest) SetStatus(v AutopayConfigStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutopayConfigUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


