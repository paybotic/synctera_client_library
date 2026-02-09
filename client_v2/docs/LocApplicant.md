# LocApplicant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to [**LocBusinessApplicant**](LocBusinessApplicant.md) |  | [optional] 
**Personal** | Pointer to [**LocPersonalApplicant**](LocPersonalApplicant.md) |  | [optional] 
**Type** | [**LocApplicantType**](LocApplicantType.md) |  | 

## Methods

### NewLocApplicant

`func NewLocApplicant(type_ LocApplicantType, ) *LocApplicant`

NewLocApplicant instantiates a new LocApplicant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocApplicantWithDefaults

`func NewLocApplicantWithDefaults() *LocApplicant`

NewLocApplicantWithDefaults instantiates a new LocApplicant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *LocApplicant) GetBusiness() LocBusinessApplicant`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *LocApplicant) GetBusinessOk() (*LocBusinessApplicant, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *LocApplicant) SetBusiness(v LocBusinessApplicant)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *LocApplicant) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetPersonal

`func (o *LocApplicant) GetPersonal() LocPersonalApplicant`

GetPersonal returns the Personal field if non-nil, zero value otherwise.

### GetPersonalOk

`func (o *LocApplicant) GetPersonalOk() (*LocPersonalApplicant, bool)`

GetPersonalOk returns a tuple with the Personal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonal

`func (o *LocApplicant) SetPersonal(v LocPersonalApplicant)`

SetPersonal sets Personal field to given value.

### HasPersonal

`func (o *LocApplicant) HasPersonal() bool`

HasPersonal returns a boolean if a field has been set.

### GetType

`func (o *LocApplicant) GetType() LocApplicantType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocApplicant) GetTypeOk() (*LocApplicantType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocApplicant) SetType(v LocApplicantType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


