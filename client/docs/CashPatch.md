# CashPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationAccountId** | Pointer to **string** | The UUID of the Synctera account resource that is the destination of the transfer for incoming transfers. This can only be updated if the transfer is suspended.  | [optional] 
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCashPatch

`func NewCashPatch() *CashPatch`

NewCashPatch instantiates a new CashPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashPatchWithDefaults

`func NewCashPatchWithDefaults() *CashPatch`

NewCashPatchWithDefaults instantiates a new CashPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationAccountId

`func (o *CashPatch) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *CashPatch) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *CashPatch) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.

### HasDestinationAccountId

`func (o *CashPatch) HasDestinationAccountId() bool`

HasDestinationAccountId returns a boolean if a field has been set.

### GetSourceData

`func (o *CashPatch) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *CashPatch) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *CashPatch) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *CashPatch) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetStatus

`func (o *CashPatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CashPatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CashPatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CashPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


