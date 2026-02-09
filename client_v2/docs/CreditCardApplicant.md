# CreditCardApplicant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to [**CreditCardBusinessApplicant**](CreditCardBusinessApplicant.md) |  | [optional] 
**Personal** | Pointer to [**CreditCardPersonalApplicant**](CreditCardPersonalApplicant.md) |  | [optional] 
**Type** | [**CreditCardApplicantType**](CreditCardApplicantType.md) |  | 

## Methods

### NewCreditCardApplicant

`func NewCreditCardApplicant(type_ CreditCardApplicantType, ) *CreditCardApplicant`

NewCreditCardApplicant instantiates a new CreditCardApplicant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardApplicantWithDefaults

`func NewCreditCardApplicantWithDefaults() *CreditCardApplicant`

NewCreditCardApplicantWithDefaults instantiates a new CreditCardApplicant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *CreditCardApplicant) GetBusiness() CreditCardBusinessApplicant`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *CreditCardApplicant) GetBusinessOk() (*CreditCardBusinessApplicant, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *CreditCardApplicant) SetBusiness(v CreditCardBusinessApplicant)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *CreditCardApplicant) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetPersonal

`func (o *CreditCardApplicant) GetPersonal() CreditCardPersonalApplicant`

GetPersonal returns the Personal field if non-nil, zero value otherwise.

### GetPersonalOk

`func (o *CreditCardApplicant) GetPersonalOk() (*CreditCardPersonalApplicant, bool)`

GetPersonalOk returns a tuple with the Personal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonal

`func (o *CreditCardApplicant) SetPersonal(v CreditCardPersonalApplicant)`

SetPersonal sets Personal field to given value.

### HasPersonal

`func (o *CreditCardApplicant) HasPersonal() bool`

HasPersonal returns a boolean if a field has been set.

### GetType

`func (o *CreditCardApplicant) GetType() CreditCardApplicantType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditCardApplicant) GetTypeOk() (*CreditCardApplicantType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditCardApplicant) SetType(v CreditCardApplicantType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


