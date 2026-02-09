# AccountClosureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | Pointer to **NullableString** | internal/external/customer account ID, depending on payment_method | [optional] 
**PaymentMethod** | Pointer to [**NullableAccountClosurePaymentMethod**](AccountClosurePaymentMethod.md) |  | [optional] 
**Reason** | [**AccountClosureReason**](AccountClosureReason.md) |  | 
**ReasonDetails** | **string** | Additional details about the reason for closing the account | 

## Methods

### NewAccountClosureRequest

`func NewAccountClosureRequest(reason AccountClosureReason, reasonDetails string, ) *AccountClosureRequest`

NewAccountClosureRequest instantiates a new AccountClosureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountClosureRequestWithDefaults

`func NewAccountClosureRequestWithDefaults() *AccountClosureRequest`

NewAccountClosureRequestWithDefaults instantiates a new AccountClosureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *AccountClosureRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *AccountClosureRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *AccountClosureRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *AccountClosureRequest) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### SetDestinationIdNil

`func (o *AccountClosureRequest) SetDestinationIdNil(b bool)`

 SetDestinationIdNil sets the value for DestinationId to be an explicit nil

### UnsetDestinationId
`func (o *AccountClosureRequest) UnsetDestinationId()`

UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
### GetPaymentMethod

`func (o *AccountClosureRequest) GetPaymentMethod() AccountClosurePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *AccountClosureRequest) GetPaymentMethodOk() (*AccountClosurePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *AccountClosureRequest) SetPaymentMethod(v AccountClosurePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *AccountClosureRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *AccountClosureRequest) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *AccountClosureRequest) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetReason

`func (o *AccountClosureRequest) GetReason() AccountClosureReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccountClosureRequest) GetReasonOk() (*AccountClosureReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccountClosureRequest) SetReason(v AccountClosureReason)`

SetReason sets Reason field to given value.


### GetReasonDetails

`func (o *AccountClosureRequest) GetReasonDetails() string`

GetReasonDetails returns the ReasonDetails field if non-nil, zero value otherwise.

### GetReasonDetailsOk

`func (o *AccountClosureRequest) GetReasonDetailsOk() (*string, bool)`

GetReasonDetailsOk returns a tuple with the ReasonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDetails

`func (o *AccountClosureRequest) SetReasonDetails(v string)`

SetReasonDetails sets ReasonDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


