# EvaluationContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | **int32** | bank ID of the customer being evaluated | 
**Email** | Pointer to **string** | Evaluate rules matching this email address.  | [optional] 
**PartnerId** | **int32** | partner ID of the customer being evaluated | 
**PhoneNumber** | Pointer to **string** | Evaluate rules matching this phone number. Use E.164 format, with a leading \&quot;+\&quot; and a country code.  | [optional] 
**ResourceType** | [**ResourceType**](ResourceType.md) |  | 
**SsnToken** | Pointer to **string** | Evaluate rules matching for the SSN represented by this token.  | [optional] 

## Methods

### NewEvaluationContext

`func NewEvaluationContext(bankId int32, partnerId int32, resourceType ResourceType, ) *EvaluationContext`

NewEvaluationContext instantiates a new EvaluationContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationContextWithDefaults

`func NewEvaluationContextWithDefaults() *EvaluationContext`

NewEvaluationContextWithDefaults instantiates a new EvaluationContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *EvaluationContext) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *EvaluationContext) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *EvaluationContext) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetEmail

`func (o *EvaluationContext) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EvaluationContext) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EvaluationContext) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EvaluationContext) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPartnerId

`func (o *EvaluationContext) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *EvaluationContext) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *EvaluationContext) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPhoneNumber

`func (o *EvaluationContext) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *EvaluationContext) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *EvaluationContext) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *EvaluationContext) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetResourceType

`func (o *EvaluationContext) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *EvaluationContext) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *EvaluationContext) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.


### GetSsnToken

`func (o *EvaluationContext) GetSsnToken() string`

GetSsnToken returns the SsnToken field if non-nil, zero value otherwise.

### GetSsnTokenOk

`func (o *EvaluationContext) GetSsnTokenOk() (*string, bool)`

GetSsnTokenOk returns a tuple with the SsnToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnToken

`func (o *EvaluationContext) SetSsnToken(v string)`

SetSsnToken sets SsnToken field to given value.

### HasSsnToken

`func (o *EvaluationContext) HasSsnToken() bool`

HasSsnToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


