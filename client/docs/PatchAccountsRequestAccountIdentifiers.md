# PatchAccountsRequestAccountIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | Pointer to **string** | The IBAN of the account. On write, Synctera will store the entire IBAN number; on read, we only return the last 4 characters.  | [optional] 
**Number** | Pointer to **string** | The account number.  | [optional] 

## Methods

### NewPatchAccountsRequestAccountIdentifiers

`func NewPatchAccountsRequestAccountIdentifiers() *PatchAccountsRequestAccountIdentifiers`

NewPatchAccountsRequestAccountIdentifiers instantiates a new PatchAccountsRequestAccountIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchAccountsRequestAccountIdentifiersWithDefaults

`func NewPatchAccountsRequestAccountIdentifiersWithDefaults() *PatchAccountsRequestAccountIdentifiers`

NewPatchAccountsRequestAccountIdentifiersWithDefaults instantiates a new PatchAccountsRequestAccountIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *PatchAccountsRequestAccountIdentifiers) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PatchAccountsRequestAccountIdentifiers) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PatchAccountsRequestAccountIdentifiers) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PatchAccountsRequestAccountIdentifiers) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetNumber

`func (o *PatchAccountsRequestAccountIdentifiers) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PatchAccountsRequestAccountIdentifiers) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PatchAccountsRequestAccountIdentifiers) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PatchAccountsRequestAccountIdentifiers) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


