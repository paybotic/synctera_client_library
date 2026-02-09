# DisputeRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCustomerReported** | **time.Time** | The timestamp representing when the customer reported the dispute. | 
**DisputedAmount** | **int32** | The amount to be disputed in cents. Disputes cannot exceed the full amount of the posted transaction. | 
**Memo** | **string** | Brief written message describing the reason for disputing the transaction. | 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**TransactionId** | **string** | The ID of the posted transaction to be disputed. | 

## Methods

### NewDisputeRequestDetails

`func NewDisputeRequestDetails(dateCustomerReported time.Time, disputedAmount int32, memo string, paymentRail PaymentRail, transactionId string, ) *DisputeRequestDetails`

NewDisputeRequestDetails instantiates a new DisputeRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeRequestDetailsWithDefaults

`func NewDisputeRequestDetailsWithDefaults() *DisputeRequestDetails`

NewDisputeRequestDetailsWithDefaults instantiates a new DisputeRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCustomerReported

`func (o *DisputeRequestDetails) GetDateCustomerReported() time.Time`

GetDateCustomerReported returns the DateCustomerReported field if non-nil, zero value otherwise.

### GetDateCustomerReportedOk

`func (o *DisputeRequestDetails) GetDateCustomerReportedOk() (*time.Time, bool)`

GetDateCustomerReportedOk returns a tuple with the DateCustomerReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCustomerReported

`func (o *DisputeRequestDetails) SetDateCustomerReported(v time.Time)`

SetDateCustomerReported sets DateCustomerReported field to given value.


### GetDisputedAmount

`func (o *DisputeRequestDetails) GetDisputedAmount() int32`

GetDisputedAmount returns the DisputedAmount field if non-nil, zero value otherwise.

### GetDisputedAmountOk

`func (o *DisputeRequestDetails) GetDisputedAmountOk() (*int32, bool)`

GetDisputedAmountOk returns a tuple with the DisputedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmount

`func (o *DisputeRequestDetails) SetDisputedAmount(v int32)`

SetDisputedAmount sets DisputedAmount field to given value.


### GetMemo

`func (o *DisputeRequestDetails) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *DisputeRequestDetails) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *DisputeRequestDetails) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetPaymentRail

`func (o *DisputeRequestDetails) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *DisputeRequestDetails) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *DisputeRequestDetails) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetTransactionId

`func (o *DisputeRequestDetails) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DisputeRequestDetails) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DisputeRequestDetails) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


