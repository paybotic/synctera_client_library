# StopPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisputeId** | Pointer to **string** | ID of the dispute that created the stop payment | [optional] 
**ExpiresOn** | Pointer to **time.Time** | The date when this stop payment is no longer valid. This is only for business accounts. | [optional] 
**OriginatorName** | **string** | Name of the originator | 
**StopPaymentId** | **string** |  | 
**TransactionId** | Pointer to **string** | If this stop payment was created from a disputed transaction, transaction_id references the posted transaction. | [optional] 

## Methods

### NewStopPayment

`func NewStopPayment(originatorName string, stopPaymentId string, ) *StopPayment`

NewStopPayment instantiates a new StopPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopPaymentWithDefaults

`func NewStopPaymentWithDefaults() *StopPayment`

NewStopPaymentWithDefaults instantiates a new StopPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisputeId

`func (o *StopPayment) GetDisputeId() string`

GetDisputeId returns the DisputeId field if non-nil, zero value otherwise.

### GetDisputeIdOk

`func (o *StopPayment) GetDisputeIdOk() (*string, bool)`

GetDisputeIdOk returns a tuple with the DisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeId

`func (o *StopPayment) SetDisputeId(v string)`

SetDisputeId sets DisputeId field to given value.

### HasDisputeId

`func (o *StopPayment) HasDisputeId() bool`

HasDisputeId returns a boolean if a field has been set.

### GetExpiresOn

`func (o *StopPayment) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *StopPayment) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *StopPayment) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *StopPayment) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetOriginatorName

`func (o *StopPayment) GetOriginatorName() string`

GetOriginatorName returns the OriginatorName field if non-nil, zero value otherwise.

### GetOriginatorNameOk

`func (o *StopPayment) GetOriginatorNameOk() (*string, bool)`

GetOriginatorNameOk returns a tuple with the OriginatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorName

`func (o *StopPayment) SetOriginatorName(v string)`

SetOriginatorName sets OriginatorName field to given value.


### GetStopPaymentId

`func (o *StopPayment) GetStopPaymentId() string`

GetStopPaymentId returns the StopPaymentId field if non-nil, zero value otherwise.

### GetStopPaymentIdOk

`func (o *StopPayment) GetStopPaymentIdOk() (*string, bool)`

GetStopPaymentIdOk returns a tuple with the StopPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPaymentId

`func (o *StopPayment) SetStopPaymentId(v string)`

SetStopPaymentId sets StopPaymentId field to given value.


### GetTransactionId

`func (o *StopPayment) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *StopPayment) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *StopPayment) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *StopPayment) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


