# AutopayAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptedAt** | **time.Time** | Timestamp when the attempt was made | 
**PaymentAttributes** | Pointer to [**AutopayPaymentAttributes**](AutopayPaymentAttributes.md) |  | [optional] 
**Reason** | Pointer to **string** | Reason for failure or skip (if applicable) | [optional] 
**Status** | [**AutopayAttemptStatus**](AutopayAttemptStatus.md) |  | 

## Methods

### NewAutopayAttempt

`func NewAutopayAttempt(attemptedAt time.Time, status AutopayAttemptStatus, ) *AutopayAttempt`

NewAutopayAttempt instantiates a new AutopayAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayAttemptWithDefaults

`func NewAutopayAttemptWithDefaults() *AutopayAttempt`

NewAutopayAttemptWithDefaults instantiates a new AutopayAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptedAt

`func (o *AutopayAttempt) GetAttemptedAt() time.Time`

GetAttemptedAt returns the AttemptedAt field if non-nil, zero value otherwise.

### GetAttemptedAtOk

`func (o *AutopayAttempt) GetAttemptedAtOk() (*time.Time, bool)`

GetAttemptedAtOk returns a tuple with the AttemptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedAt

`func (o *AutopayAttempt) SetAttemptedAt(v time.Time)`

SetAttemptedAt sets AttemptedAt field to given value.


### GetPaymentAttributes

`func (o *AutopayAttempt) GetPaymentAttributes() AutopayPaymentAttributes`

GetPaymentAttributes returns the PaymentAttributes field if non-nil, zero value otherwise.

### GetPaymentAttributesOk

`func (o *AutopayAttempt) GetPaymentAttributesOk() (*AutopayPaymentAttributes, bool)`

GetPaymentAttributesOk returns a tuple with the PaymentAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAttributes

`func (o *AutopayAttempt) SetPaymentAttributes(v AutopayPaymentAttributes)`

SetPaymentAttributes sets PaymentAttributes field to given value.

### HasPaymentAttributes

`func (o *AutopayAttempt) HasPaymentAttributes() bool`

HasPaymentAttributes returns a boolean if a field has been set.

### GetReason

`func (o *AutopayAttempt) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AutopayAttempt) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AutopayAttempt) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AutopayAttempt) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *AutopayAttempt) GetStatus() AutopayAttemptStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutopayAttempt) GetStatusOk() (*AutopayAttemptStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutopayAttempt) SetStatus(v AutopayAttemptStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


