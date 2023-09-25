# InternalAccountPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | Generated internal account number | [optional] [readonly] 
**AccountType** | Pointer to [**InternalAccountType**](InternalAccountType.md) |  | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for internal account based on different type | [optional] [readonly] 
**BankRouting** | Pointer to **string** | Bank routing number | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Currency** | **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. | 
**Description** | Pointer to **string** | A user provided description for the current account | [optional] 
**GlType** | Pointer to **string** | What type of general ledger account this internal account represents.  | [optional] 
**Id** | Pointer to **string** | Generated ID for internal account | [optional] [readonly] 
**IsSystemAcc** | Pointer to **bool** | Is a system-controlled internal account. When this field is true, this internal account will be reserved exclusively for internal use by the Synctera platform and any internal transfers to or from this internal account will be declined. | [optional] [default to false]
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**Purpose** | Pointer to [**InternalAccountPurpose**](InternalAccountPurpose.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewInternalAccountPost

`func NewInternalAccountPost(currency string, status string, ) *InternalAccountPost`

NewInternalAccountPost instantiates a new InternalAccountPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalAccountPostWithDefaults

`func NewInternalAccountPostWithDefaults() *InternalAccountPost`

NewInternalAccountPostWithDefaults instantiates a new InternalAccountPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *InternalAccountPost) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *InternalAccountPost) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *InternalAccountPost) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *InternalAccountPost) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountType

`func (o *InternalAccountPost) GetAccountType() InternalAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *InternalAccountPost) GetAccountTypeOk() (*InternalAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *InternalAccountPost) SetAccountType(v InternalAccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *InternalAccountPost) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBalances

`func (o *InternalAccountPost) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *InternalAccountPost) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *InternalAccountPost) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *InternalAccountPost) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankRouting

`func (o *InternalAccountPost) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *InternalAccountPost) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *InternalAccountPost) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *InternalAccountPost) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetCreationTime

`func (o *InternalAccountPost) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *InternalAccountPost) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *InternalAccountPost) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *InternalAccountPost) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *InternalAccountPost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalAccountPost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalAccountPost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *InternalAccountPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalAccountPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalAccountPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalAccountPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGlType

`func (o *InternalAccountPost) GetGlType() string`

GetGlType returns the GlType field if non-nil, zero value otherwise.

### GetGlTypeOk

`func (o *InternalAccountPost) GetGlTypeOk() (*string, bool)`

GetGlTypeOk returns a tuple with the GlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlType

`func (o *InternalAccountPost) SetGlType(v string)`

SetGlType sets GlType field to given value.

### HasGlType

`func (o *InternalAccountPost) HasGlType() bool`

HasGlType returns a boolean if a field has been set.

### GetId

`func (o *InternalAccountPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalAccountPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalAccountPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InternalAccountPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSystemAcc

`func (o *InternalAccountPost) GetIsSystemAcc() bool`

GetIsSystemAcc returns the IsSystemAcc field if non-nil, zero value otherwise.

### GetIsSystemAccOk

`func (o *InternalAccountPost) GetIsSystemAccOk() (*bool, bool)`

GetIsSystemAccOk returns a tuple with the IsSystemAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAcc

`func (o *InternalAccountPost) SetIsSystemAcc(v bool)`

SetIsSystemAcc sets IsSystemAcc field to given value.

### HasIsSystemAcc

`func (o *InternalAccountPost) HasIsSystemAcc() bool`

HasIsSystemAcc returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *InternalAccountPost) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *InternalAccountPost) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *InternalAccountPost) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *InternalAccountPost) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetPurpose

`func (o *InternalAccountPost) GetPurpose() InternalAccountPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *InternalAccountPost) GetPurposeOk() (*InternalAccountPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *InternalAccountPost) SetPurpose(v InternalAccountPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *InternalAccountPost) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetStatus

`func (o *InternalAccountPost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalAccountPost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalAccountPost) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


