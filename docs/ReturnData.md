# ReturnData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalId** | Pointer to **string** | Wire UUID of the original wire that was returned | [optional] 
**OriginalTransactionId** | Pointer to **string** | Transaction UUID of the original wire that was returned | [optional] 
**PreviousMessageId** | **string** | Wire reference ID of the original wire that was returned | 
**Reason** | Pointer to **string** | The cause of the return | [optional] 

## Methods

### NewReturnData

`func NewReturnData(previousMessageId string, ) *ReturnData`

NewReturnData instantiates a new ReturnData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnDataWithDefaults

`func NewReturnDataWithDefaults() *ReturnData`

NewReturnDataWithDefaults instantiates a new ReturnData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalId

`func (o *ReturnData) GetOriginalId() string`

GetOriginalId returns the OriginalId field if non-nil, zero value otherwise.

### GetOriginalIdOk

`func (o *ReturnData) GetOriginalIdOk() (*string, bool)`

GetOriginalIdOk returns a tuple with the OriginalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalId

`func (o *ReturnData) SetOriginalId(v string)`

SetOriginalId sets OriginalId field to given value.

### HasOriginalId

`func (o *ReturnData) HasOriginalId() bool`

HasOriginalId returns a boolean if a field has been set.

### GetOriginalTransactionId

`func (o *ReturnData) GetOriginalTransactionId() string`

GetOriginalTransactionId returns the OriginalTransactionId field if non-nil, zero value otherwise.

### GetOriginalTransactionIdOk

`func (o *ReturnData) GetOriginalTransactionIdOk() (*string, bool)`

GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTransactionId

`func (o *ReturnData) SetOriginalTransactionId(v string)`

SetOriginalTransactionId sets OriginalTransactionId field to given value.

### HasOriginalTransactionId

`func (o *ReturnData) HasOriginalTransactionId() bool`

HasOriginalTransactionId returns a boolean if a field has been set.

### GetPreviousMessageId

`func (o *ReturnData) GetPreviousMessageId() string`

GetPreviousMessageId returns the PreviousMessageId field if non-nil, zero value otherwise.

### GetPreviousMessageIdOk

`func (o *ReturnData) GetPreviousMessageIdOk() (*string, bool)`

GetPreviousMessageIdOk returns a tuple with the PreviousMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousMessageId

`func (o *ReturnData) SetPreviousMessageId(v string)`

SetPreviousMessageId sets PreviousMessageId field to given value.


### GetReason

`func (o *ReturnData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ReturnData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ReturnData) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ReturnData) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


