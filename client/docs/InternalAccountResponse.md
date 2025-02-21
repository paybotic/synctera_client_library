# InternalAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | Generated internal account number | [optional] [readonly] 
**AccountType** | Pointer to [**InternalAccountType**](InternalAccountType.md) |  | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for internal account based on different type | [optional] [readonly] 
**BankAccountId** | Pointer to **string** | The ID of the bank account associated with this internal account. It will be auto-filled if the account type has only one matching bank account in the system.  | [optional] 
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
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 

## Methods

### NewInternalAccountResponse

`func NewInternalAccountResponse(currency string, status string, ) *InternalAccountResponse`

NewInternalAccountResponse instantiates a new InternalAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalAccountResponseWithDefaults

`func NewInternalAccountResponseWithDefaults() *InternalAccountResponse`

NewInternalAccountResponseWithDefaults instantiates a new InternalAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *InternalAccountResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *InternalAccountResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *InternalAccountResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *InternalAccountResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountType

`func (o *InternalAccountResponse) GetAccountType() InternalAccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *InternalAccountResponse) GetAccountTypeOk() (*InternalAccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *InternalAccountResponse) SetAccountType(v InternalAccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *InternalAccountResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBalances

`func (o *InternalAccountResponse) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *InternalAccountResponse) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *InternalAccountResponse) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *InternalAccountResponse) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankAccountId

`func (o *InternalAccountResponse) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *InternalAccountResponse) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *InternalAccountResponse) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *InternalAccountResponse) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### GetBankRouting

`func (o *InternalAccountResponse) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *InternalAccountResponse) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *InternalAccountResponse) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *InternalAccountResponse) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetCreationTime

`func (o *InternalAccountResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *InternalAccountResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *InternalAccountResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *InternalAccountResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *InternalAccountResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalAccountResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalAccountResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *InternalAccountResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalAccountResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalAccountResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalAccountResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGlType

`func (o *InternalAccountResponse) GetGlType() string`

GetGlType returns the GlType field if non-nil, zero value otherwise.

### GetGlTypeOk

`func (o *InternalAccountResponse) GetGlTypeOk() (*string, bool)`

GetGlTypeOk returns a tuple with the GlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlType

`func (o *InternalAccountResponse) SetGlType(v string)`

SetGlType sets GlType field to given value.

### HasGlType

`func (o *InternalAccountResponse) HasGlType() bool`

HasGlType returns a boolean if a field has been set.

### GetId

`func (o *InternalAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalAccountResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InternalAccountResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSystemAcc

`func (o *InternalAccountResponse) GetIsSystemAcc() bool`

GetIsSystemAcc returns the IsSystemAcc field if non-nil, zero value otherwise.

### GetIsSystemAccOk

`func (o *InternalAccountResponse) GetIsSystemAccOk() (*bool, bool)`

GetIsSystemAccOk returns a tuple with the IsSystemAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAcc

`func (o *InternalAccountResponse) SetIsSystemAcc(v bool)`

SetIsSystemAcc sets IsSystemAcc field to given value.

### HasIsSystemAcc

`func (o *InternalAccountResponse) HasIsSystemAcc() bool`

HasIsSystemAcc returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *InternalAccountResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *InternalAccountResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *InternalAccountResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *InternalAccountResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetPurpose

`func (o *InternalAccountResponse) GetPurpose() InternalAccountPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *InternalAccountResponse) GetPurposeOk() (*InternalAccountPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *InternalAccountResponse) SetPurpose(v InternalAccountPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *InternalAccountResponse) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetStatus

`func (o *InternalAccountResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalAccountResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalAccountResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *InternalAccountResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *InternalAccountResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *InternalAccountResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *InternalAccountResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


