# EftCaPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 
**Status** | Pointer to [**EftCaStatus**](EftCaStatus.md) |  | [optional] 

## Methods

### NewEftCaPatch

`func NewEftCaPatch() *EftCaPatch`

NewEftCaPatch instantiates a new EftCaPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEftCaPatchWithDefaults

`func NewEftCaPatchWithDefaults() *EftCaPatch`

NewEftCaPatchWithDefaults instantiates a new EftCaPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceData

`func (o *EftCaPatch) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *EftCaPatch) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *EftCaPatch) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *EftCaPatch) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetStatus

`func (o *EftCaPatch) GetStatus() EftCaStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EftCaPatch) GetStatusOk() (*EftCaStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EftCaPatch) SetStatus(v EftCaStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EftCaPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


