# AccountClosureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cases** | Pointer to [**[]ModelCase**](ModelCase.md) | Cases associated with the account closure | [optional] 
**DestinationId** | **NullableString** | internal/external/customer account ID, depending on payment_method | 
**PaymentMethod** | [**NullableAccountClosurePaymentMethod**](AccountClosurePaymentMethod.md) |  | 
**Reason** | [**AccountClosureReason**](AccountClosureReason.md) |  | 
**ReasonDetails** | **string** | Additional details about the reason for closing the account | 
**ValidationResponses** | Pointer to [**[]AccountClosureValidationResponse**](AccountClosureValidationResponse.md) |  | [optional] 

## Methods

### NewAccountClosureResponse

`func NewAccountClosureResponse(destinationId NullableString, paymentMethod NullableAccountClosurePaymentMethod, reason AccountClosureReason, reasonDetails string, ) *AccountClosureResponse`

NewAccountClosureResponse instantiates a new AccountClosureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountClosureResponseWithDefaults

`func NewAccountClosureResponseWithDefaults() *AccountClosureResponse`

NewAccountClosureResponseWithDefaults instantiates a new AccountClosureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCases

`func (o *AccountClosureResponse) GetCases() []ModelCase`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *AccountClosureResponse) GetCasesOk() (*[]ModelCase, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *AccountClosureResponse) SetCases(v []ModelCase)`

SetCases sets Cases field to given value.

### HasCases

`func (o *AccountClosureResponse) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetDestinationId

`func (o *AccountClosureResponse) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *AccountClosureResponse) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *AccountClosureResponse) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### SetDestinationIdNil

`func (o *AccountClosureResponse) SetDestinationIdNil(b bool)`

 SetDestinationIdNil sets the value for DestinationId to be an explicit nil

### UnsetDestinationId
`func (o *AccountClosureResponse) UnsetDestinationId()`

UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
### GetPaymentMethod

`func (o *AccountClosureResponse) GetPaymentMethod() AccountClosurePaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *AccountClosureResponse) GetPaymentMethodOk() (*AccountClosurePaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *AccountClosureResponse) SetPaymentMethod(v AccountClosurePaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### SetPaymentMethodNil

`func (o *AccountClosureResponse) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *AccountClosureResponse) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetReason

`func (o *AccountClosureResponse) GetReason() AccountClosureReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccountClosureResponse) GetReasonOk() (*AccountClosureReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccountClosureResponse) SetReason(v AccountClosureReason)`

SetReason sets Reason field to given value.


### GetReasonDetails

`func (o *AccountClosureResponse) GetReasonDetails() string`

GetReasonDetails returns the ReasonDetails field if non-nil, zero value otherwise.

### GetReasonDetailsOk

`func (o *AccountClosureResponse) GetReasonDetailsOk() (*string, bool)`

GetReasonDetailsOk returns a tuple with the ReasonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDetails

`func (o *AccountClosureResponse) SetReasonDetails(v string)`

SetReasonDetails sets ReasonDetails field to given value.


### GetValidationResponses

`func (o *AccountClosureResponse) GetValidationResponses() []AccountClosureValidationResponse`

GetValidationResponses returns the ValidationResponses field if non-nil, zero value otherwise.

### GetValidationResponsesOk

`func (o *AccountClosureResponse) GetValidationResponsesOk() (*[]AccountClosureValidationResponse, bool)`

GetValidationResponsesOk returns a tuple with the ValidationResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResponses

`func (o *AccountClosureResponse) SetValidationResponses(v []AccountClosureValidationResponse)`

SetValidationResponses sets ValidationResponses field to given value.

### HasValidationResponses

`func (o *AccountClosureResponse) HasValidationResponses() bool`

HasValidationResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


