# TemplateFieldsRevolvingCreditSecured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountProgramId** | Pointer to **string** |  | [optional] 
**AccountType** | [**AccountType**](AccountType.md) |  | 
**BankAccountId** | Pointer to **string** | The bank account ID for this account. This is a unique identifier for the bank side account that this Synctera account belongs to. This field can be auto filled if only one bank account of the appropriate type exist for the tenant of concern.  | [optional] 
**BankCountry** | **string** | Bank country of the account. ISO 3166-1 Alpha-2 or Alpha-3 country code. | 
**Currency** | **string** | Account currency. ISO 4217 alphabetic currency code | 
**ApplicationWorkflowId** | **string** | Taktile workflow ID for credit application processing | 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 
**VendorInfo** | [**TemplateVendorInfo**](TemplateVendorInfo.md) |  | 

## Methods

### NewTemplateFieldsRevolvingCreditSecured

`func NewTemplateFieldsRevolvingCreditSecured(accountType AccountType, bankCountry string, currency string, applicationWorkflowId string, vendorInfo TemplateVendorInfo, ) *TemplateFieldsRevolvingCreditSecured`

NewTemplateFieldsRevolvingCreditSecured instantiates a new TemplateFieldsRevolvingCreditSecured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldsRevolvingCreditSecuredWithDefaults

`func NewTemplateFieldsRevolvingCreditSecuredWithDefaults() *TemplateFieldsRevolvingCreditSecured`

NewTemplateFieldsRevolvingCreditSecuredWithDefaults instantiates a new TemplateFieldsRevolvingCreditSecured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountProgramId

`func (o *TemplateFieldsRevolvingCreditSecured) GetAccountProgramId() string`

GetAccountProgramId returns the AccountProgramId field if non-nil, zero value otherwise.

### GetAccountProgramIdOk

`func (o *TemplateFieldsRevolvingCreditSecured) GetAccountProgramIdOk() (*string, bool)`

GetAccountProgramIdOk returns a tuple with the AccountProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProgramId

`func (o *TemplateFieldsRevolvingCreditSecured) SetAccountProgramId(v string)`

SetAccountProgramId sets AccountProgramId field to given value.

### HasAccountProgramId

`func (o *TemplateFieldsRevolvingCreditSecured) HasAccountProgramId() bool`

HasAccountProgramId returns a boolean if a field has been set.

### GetAccountType

`func (o *TemplateFieldsRevolvingCreditSecured) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *TemplateFieldsRevolvingCreditSecured) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *TemplateFieldsRevolvingCreditSecured) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetBankAccountId

`func (o *TemplateFieldsRevolvingCreditSecured) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *TemplateFieldsRevolvingCreditSecured) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *TemplateFieldsRevolvingCreditSecured) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *TemplateFieldsRevolvingCreditSecured) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### GetBankCountry

`func (o *TemplateFieldsRevolvingCreditSecured) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *TemplateFieldsRevolvingCreditSecured) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *TemplateFieldsRevolvingCreditSecured) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.


### GetCurrency

`func (o *TemplateFieldsRevolvingCreditSecured) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TemplateFieldsRevolvingCreditSecured) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TemplateFieldsRevolvingCreditSecured) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetApplicationWorkflowId

`func (o *TemplateFieldsRevolvingCreditSecured) GetApplicationWorkflowId() string`

GetApplicationWorkflowId returns the ApplicationWorkflowId field if non-nil, zero value otherwise.

### GetApplicationWorkflowIdOk

`func (o *TemplateFieldsRevolvingCreditSecured) GetApplicationWorkflowIdOk() (*string, bool)`

GetApplicationWorkflowIdOk returns a tuple with the ApplicationWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationWorkflowId

`func (o *TemplateFieldsRevolvingCreditSecured) SetApplicationWorkflowId(v string)`

SetApplicationWorkflowId sets ApplicationWorkflowId field to given value.


### GetSpendControlIds

`func (o *TemplateFieldsRevolvingCreditSecured) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *TemplateFieldsRevolvingCreditSecured) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *TemplateFieldsRevolvingCreditSecured) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *TemplateFieldsRevolvingCreditSecured) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.

### GetVendorInfo

`func (o *TemplateFieldsRevolvingCreditSecured) GetVendorInfo() TemplateVendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *TemplateFieldsRevolvingCreditSecured) GetVendorInfoOk() (*TemplateVendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *TemplateFieldsRevolvingCreditSecured) SetVendorInfo(v TemplateVendorInfo)`

SetVendorInfo sets VendorInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


