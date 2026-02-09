# CreditCardPersonalApplicant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employment** | Pointer to [**CreditCardEmployment**](CreditCardEmployment.md) |  | [optional] 
**PersonId** | **string** | Person ID for the applicant | 

## Methods

### NewCreditCardPersonalApplicant

`func NewCreditCardPersonalApplicant(personId string, ) *CreditCardPersonalApplicant`

NewCreditCardPersonalApplicant instantiates a new CreditCardPersonalApplicant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardPersonalApplicantWithDefaults

`func NewCreditCardPersonalApplicantWithDefaults() *CreditCardPersonalApplicant`

NewCreditCardPersonalApplicantWithDefaults instantiates a new CreditCardPersonalApplicant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployment

`func (o *CreditCardPersonalApplicant) GetEmployment() CreditCardEmployment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *CreditCardPersonalApplicant) GetEmploymentOk() (*CreditCardEmployment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *CreditCardPersonalApplicant) SetEmployment(v CreditCardEmployment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *CreditCardPersonalApplicant) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### GetPersonId

`func (o *CreditCardPersonalApplicant) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *CreditCardPersonalApplicant) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *CreditCardPersonalApplicant) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


