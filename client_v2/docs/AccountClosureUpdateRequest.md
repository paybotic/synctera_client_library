# AccountClosureUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | Pointer to **NullableString** | internal/external/customer account ID, depending on payment_method | [optional] 
**PaymentMethod** | Pointer to [**NullableAccountClosurePaymentMethod**](AccountClosurePaymentMethod.md) |  | [optional] 

## Methods

### NewAccountClosureUpdateRequest

`func NewAccountClosureUpdateRequest() *AccountClosureUpdateRequest`

NewAccountClosureUpdateRequest instantiates a new AccountClosureUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountClosureUpdateRequestWithDefaults

`func NewAccountClosureUpdateRequestWithDefaults() *AccountClosureUpdateRequest`

NewAccountClosureUpdateRequestWithDefaults instantiates a new AccountClosureUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *AccountClosureUpdateRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *AccountClosureUpdateRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *AccountClosureUpdateRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *AccountClosureUpdateRequest) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### SetDestinationIdNil

`func (o *AccountClosureUpdateRequest) SetDestinationIdNil(b bool)`

 SetDestinationIdNil sets the value for DestinationId to be an explicit nil

### UnsetDestinationId
`func (o *AccountClosureUpdateRequest) UnsetDestinationId()`

UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
### GetPaymentMethod

`func (o *AccountClosureUpdateRequest) GetPaymentMethod() AccountClosurePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *AccountClosureUpdateRequest) GetPaymentMethodOk() (*AccountClosurePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *AccountClosureUpdateRequest) SetPaymentMethod(v AccountClosurePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *AccountClosureUpdateRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *AccountClosureUpdateRequest) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *AccountClosureUpdateRequest) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


