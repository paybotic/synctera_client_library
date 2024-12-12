# AccountClosure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cases** | Pointer to [**[]ModelCase**](ModelCase.md) | Cases associated with the account closure | [optional] [readonly] 
**DestinationId** | **string** | For an ACH payment, this is the external account UUID. For an internal payment, this is the account UUID. | 
**PaymentMethod** | **string** | Payment method for the final payment if the account being closed carries a balance. A BANK_DRAFT payment method is issued by the sponsor bank to the account holder.  | 
**Reason** | [**AccountClosureReason**](AccountClosureReason.md) |  | 
**ReasonDetails** | **string** | Additional details about the reason for closing the account | 
**ValidationResponses** | Pointer to [**[]AccountClosureValidationResponse**](AccountClosureValidationResponse.md) |  | [optional] [readonly] 

## Methods

### NewAccountClosure

`func NewAccountClosure(destinationId string, paymentMethod string, reason AccountClosureReason, reasonDetails string, ) *AccountClosure`

NewAccountClosure instantiates a new AccountClosure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountClosureWithDefaults

`func NewAccountClosureWithDefaults() *AccountClosure`

NewAccountClosureWithDefaults instantiates a new AccountClosure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCases

`func (o *AccountClosure) GetCases() []ModelCase`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *AccountClosure) GetCasesOk() (*[]ModelCase, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *AccountClosure) SetCases(v []ModelCase)`

SetCases sets Cases field to given value.

### HasCases

`func (o *AccountClosure) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetDestinationId

`func (o *AccountClosure) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *AccountClosure) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *AccountClosure) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetPaymentMethod

`func (o *AccountClosure) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *AccountClosure) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *AccountClosure) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetReason

`func (o *AccountClosure) GetReason() AccountClosureReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccountClosure) GetReasonOk() (*AccountClosureReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccountClosure) SetReason(v AccountClosureReason)`

SetReason sets Reason field to given value.


### GetReasonDetails

`func (o *AccountClosure) GetReasonDetails() string`

GetReasonDetails returns the ReasonDetails field if non-nil, zero value otherwise.

### GetReasonDetailsOk

`func (o *AccountClosure) GetReasonDetailsOk() (*string, bool)`

GetReasonDetailsOk returns a tuple with the ReasonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDetails

`func (o *AccountClosure) SetReasonDetails(v string)`

SetReasonDetails sets ReasonDetails field to given value.


### GetValidationResponses

`func (o *AccountClosure) GetValidationResponses() []AccountClosureValidationResponse`

GetValidationResponses returns the ValidationResponses field if non-nil, zero value otherwise.

### GetValidationResponsesOk

`func (o *AccountClosure) GetValidationResponsesOk() (*[]AccountClosureValidationResponse, bool)`

GetValidationResponsesOk returns a tuple with the ValidationResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResponses

`func (o *AccountClosure) SetValidationResponses(v []AccountClosureValidationResponse)`

SetValidationResponses sets ValidationResponses field to given value.

### HasValidationResponses

`func (o *AccountClosure) HasValidationResponses() bool`

HasValidationResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


