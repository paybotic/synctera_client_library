# ReturnData1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalId** | Pointer to **string** | Wire UUID of the original wire that was returned | [optional] 
**OriginalTransactionId** | Pointer to **string** | Transaction UUID of the original wire that was returned | [optional] 
**PreviousMessageId** | **string** | Wire reference ID of the original wire that was returned | 
**Reason** | Pointer to **string** | The cause of the return | [optional] 

## Methods

### NewReturnData1

`func NewReturnData1(previousMessageId string, ) *ReturnData1`

NewReturnData1 instantiates a new ReturnData1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnData1WithDefaults

`func NewReturnData1WithDefaults() *ReturnData1`

NewReturnData1WithDefaults instantiates a new ReturnData1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalId

`func (o *ReturnData1) GetOriginalId() string`

GetOriginalId returns the OriginalId field if non-nil, zero value otherwise.

### GetOriginalIdOk

`func (o *ReturnData1) GetOriginalIdOk() (*string, bool)`

GetOriginalIdOk returns a tuple with the OriginalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalId

`func (o *ReturnData1) SetOriginalId(v string)`

SetOriginalId sets OriginalId field to given value.

### HasOriginalId

`func (o *ReturnData1) HasOriginalId() bool`

HasOriginalId returns a boolean if a field has been set.

### GetOriginalTransactionId

`func (o *ReturnData1) GetOriginalTransactionId() string`

GetOriginalTransactionId returns the OriginalTransactionId field if non-nil, zero value otherwise.

### GetOriginalTransactionIdOk

`func (o *ReturnData1) GetOriginalTransactionIdOk() (*string, bool)`

GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTransactionId

`func (o *ReturnData1) SetOriginalTransactionId(v string)`

SetOriginalTransactionId sets OriginalTransactionId field to given value.

### HasOriginalTransactionId

`func (o *ReturnData1) HasOriginalTransactionId() bool`

HasOriginalTransactionId returns a boolean if a field has been set.

### GetPreviousMessageId

`func (o *ReturnData1) GetPreviousMessageId() string`

GetPreviousMessageId returns the PreviousMessageId field if non-nil, zero value otherwise.

### GetPreviousMessageIdOk

`func (o *ReturnData1) GetPreviousMessageIdOk() (*string, bool)`

GetPreviousMessageIdOk returns a tuple with the PreviousMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousMessageId

`func (o *ReturnData1) SetPreviousMessageId(v string)`

SetPreviousMessageId sets PreviousMessageId field to given value.


### GetReason

`func (o *ReturnData1) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ReturnData1) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ReturnData1) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ReturnData1) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


