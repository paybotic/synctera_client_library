# ExternalCardDisputeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCustomerReported** | **time.Time** | The timestamp representing when the customer reported the dispute. | 
**DisputedAmount** | **int32** | The amount to be disputed in cents. Disputes cannot exceed the full amount of the posted transaction. | 
**Memo** | **string** | Brief written message describing the reason for disputing the transaction. | 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**TransactionId** | **string** | The ID of the posted transaction to be disputed. | 
**ExternalReferenceId** | **string** | Reference ID associated with the dispute on the external network. | 
**ReasonCode** | [**ExternalCardDisputeReasonCode**](ExternalCardDisputeReasonCode.md) |  | 

## Methods

### NewExternalCardDisputeRequest

`func NewExternalCardDisputeRequest(dateCustomerReported time.Time, disputedAmount int32, memo string, paymentRail PaymentRail, transactionId string, externalReferenceId string, reasonCode ExternalCardDisputeReasonCode, ) *ExternalCardDisputeRequest`

NewExternalCardDisputeRequest instantiates a new ExternalCardDisputeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardDisputeRequestWithDefaults

`func NewExternalCardDisputeRequestWithDefaults() *ExternalCardDisputeRequest`

NewExternalCardDisputeRequestWithDefaults instantiates a new ExternalCardDisputeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCustomerReported

`func (o *ExternalCardDisputeRequest) GetDateCustomerReported() time.Time`

GetDateCustomerReported returns the DateCustomerReported field if non-nil, zero value otherwise.

### GetDateCustomerReportedOk

`func (o *ExternalCardDisputeRequest) GetDateCustomerReportedOk() (*time.Time, bool)`

GetDateCustomerReportedOk returns a tuple with the DateCustomerReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCustomerReported

`func (o *ExternalCardDisputeRequest) SetDateCustomerReported(v time.Time)`

SetDateCustomerReported sets DateCustomerReported field to given value.


### GetDisputedAmount

`func (o *ExternalCardDisputeRequest) GetDisputedAmount() int32`

GetDisputedAmount returns the DisputedAmount field if non-nil, zero value otherwise.

### GetDisputedAmountOk

`func (o *ExternalCardDisputeRequest) GetDisputedAmountOk() (*int32, bool)`

GetDisputedAmountOk returns a tuple with the DisputedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmount

`func (o *ExternalCardDisputeRequest) SetDisputedAmount(v int32)`

SetDisputedAmount sets DisputedAmount field to given value.


### GetMemo

`func (o *ExternalCardDisputeRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ExternalCardDisputeRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ExternalCardDisputeRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetPaymentRail

`func (o *ExternalCardDisputeRequest) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *ExternalCardDisputeRequest) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *ExternalCardDisputeRequest) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetTransactionId

`func (o *ExternalCardDisputeRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ExternalCardDisputeRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ExternalCardDisputeRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetExternalReferenceId

`func (o *ExternalCardDisputeRequest) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *ExternalCardDisputeRequest) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *ExternalCardDisputeRequest) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.


### GetReasonCode

`func (o *ExternalCardDisputeRequest) GetReasonCode() ExternalCardDisputeReasonCode`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ExternalCardDisputeRequest) GetReasonCodeOk() (*ExternalCardDisputeReasonCode, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ExternalCardDisputeRequest) SetReasonCode(v ExternalCardDisputeReasonCode)`

SetReasonCode sets ReasonCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


