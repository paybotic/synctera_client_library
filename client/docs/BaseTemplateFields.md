# BaseTemplateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**BankAccountId** | Pointer to **string** | The bank account ID for this account. This is a unique identifier for the bank side account that this Synctera account belongs to. This field can be auto filled if only one bank account of the appropriate type exist for the tenant of concern.  | [optional] 
**BankCountry** | **string** | Bank country of the account. ISO 3166-1 Alpha-2 or Alpha-3 country code. | 
**Currency** | **string** | Account currency. ISO 4217 alphabetic currency code | 

## Methods

### NewBaseTemplateFields

`func NewBaseTemplateFields(accountType AccountType, bankCountry string, currency string, ) *BaseTemplateFields`

NewBaseTemplateFields instantiates a new BaseTemplateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTemplateFieldsWithDefaults

`func NewBaseTemplateFieldsWithDefaults() *BaseTemplateFields`

NewBaseTemplateFieldsWithDefaults instantiates a new BaseTemplateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *BaseTemplateFields) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BaseTemplateFields) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BaseTemplateFields) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetBankAccountId

`func (o *BaseTemplateFields) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BaseTemplateFields) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BaseTemplateFields) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BaseTemplateFields) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### GetBankCountry

`func (o *BaseTemplateFields) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *BaseTemplateFields) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *BaseTemplateFields) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.


### GetCurrency

`func (o *BaseTemplateFields) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BaseTemplateFields) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BaseTemplateFields) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


