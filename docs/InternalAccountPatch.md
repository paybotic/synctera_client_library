# InternalAccountPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to [**InternalAccountType**](InternalAccountType.md) |  | [optional] 
**Description** | Pointer to **string** | A user provided description for the current account | [optional] 
**Purpose** | Pointer to [**InternalAccountPurpose**](InternalAccountPurpose.md) |  | [optional] 

## Methods

### NewInternalAccountPatch

`func NewInternalAccountPatch() *InternalAccountPatch`

NewInternalAccountPatch instantiates a new InternalAccountPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalAccountPatchWithDefaults

`func NewInternalAccountPatchWithDefaults() *InternalAccountPatch`

NewInternalAccountPatchWithDefaults instantiates a new InternalAccountPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *InternalAccountPatch) GetAccountType() InternalAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *InternalAccountPatch) GetAccountTypeOk() (*InternalAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *InternalAccountPatch) SetAccountType(v InternalAccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *InternalAccountPatch) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetDescription

`func (o *InternalAccountPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalAccountPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalAccountPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalAccountPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPurpose

`func (o *InternalAccountPatch) GetPurpose() InternalAccountPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *InternalAccountPatch) GetPurposeOk() (*InternalAccountPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *InternalAccountPatch) SetPurpose(v InternalAccountPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *InternalAccountPatch) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


