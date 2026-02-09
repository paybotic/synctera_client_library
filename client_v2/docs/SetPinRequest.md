# SetPinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | **string** | The new PIN for the card | 

## Methods

### NewSetPinRequest

`func NewSetPinRequest(pin string, ) *SetPinRequest`

NewSetPinRequest instantiates a new SetPinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPinRequestWithDefaults

`func NewSetPinRequestWithDefaults() *SetPinRequest`

NewSetPinRequestWithDefaults instantiates a new SetPinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *SetPinRequest) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *SetPinRequest) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *SetPinRequest) SetPin(v string)`

SetPin sets Pin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


