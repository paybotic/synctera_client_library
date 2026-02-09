# DisputeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCustomerReported** | **time.Time** | The timestamp representing when the customer reported the dispute. | 
**DisputedAmount** | **int32** | The amount to be disputed in cents. Disputes cannot exceed the full amount of the posted transaction. | 
**Memo** | **string** | Brief written message describing the reason for disputing the transaction. | 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**TransactionId** | **string** | The ID of the posted transaction to be disputed. | 
**ReasonCode** | [**ExternalCardDisputeReasonCode**](ExternalCardDisputeReasonCode.md) |  | 
**ExternalReferenceId** | **string** | Reference ID associated with the dispute on the external network. | 

## Methods

### NewDisputeRequest

`func NewDisputeRequest(dateCustomerReported time.Time, disputedAmount int32, memo string, paymentRail PaymentRail, transactionId string, reasonCode ExternalCardDisputeReasonCode, externalReferenceId string, ) *DisputeRequest`

NewDisputeRequest instantiates a new DisputeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeRequestWithDefaults

`func NewDisputeRequestWithDefaults() *DisputeRequest`

NewDisputeRequestWithDefaults instantiates a new DisputeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCustomerReported

`func (o *DisputeRequest) GetDateCustomerReported() time.Time`

GetDateCustomerReported returns the DateCustomerReported field if non-nil, zero value otherwise.

### GetDateCustomerReportedOk

`func (o *DisputeRequest) GetDateCustomerReportedOk() (*time.Time, bool)`

GetDateCustomerReportedOk returns a tuple with the DateCustomerReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCustomerReported

`func (o *DisputeRequest) SetDateCustomerReported(v time.Time)`

SetDateCustomerReported sets DateCustomerReported field to given value.


### GetDisputedAmount

`func (o *DisputeRequest) GetDisputedAmount() int32`

GetDisputedAmount returns the DisputedAmount field if non-nil, zero value otherwise.

### GetDisputedAmountOk

`func (o *DisputeRequest) GetDisputedAmountOk() (*int32, bool)`

GetDisputedAmountOk returns a tuple with the DisputedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmount

`func (o *DisputeRequest) SetDisputedAmount(v int32)`

SetDisputedAmount sets DisputedAmount field to given value.


### GetMemo

`func (o *DisputeRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *DisputeRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *DisputeRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetPaymentRail

`func (o *DisputeRequest) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *DisputeRequest) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *DisputeRequest) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetTransactionId

`func (o *DisputeRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DisputeRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DisputeRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetReasonCode

`func (o *DisputeRequest) GetReasonCode() ExternalCardDisputeReasonCode`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DisputeRequest) GetReasonCodeOk() (*ExternalCardDisputeReasonCode, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DisputeRequest) SetReasonCode(v ExternalCardDisputeReasonCode)`

SetReasonCode sets ReasonCode field to given value.


### GetExternalReferenceId

`func (o *DisputeRequest) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *DisputeRequest) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *DisputeRequest) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


