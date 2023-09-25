# TemplateFieldsChargeSecured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**BankCountry** | **string** | Bank country of the account. ISO 3166-1 Alpha-2 or Alpha-3 country code. | 
**Currency** | **string** | Account currency. ISO 4217 alphabetic currency code | 
**AutoPaymentPeriod** | Pointer to **int32** | The number of days past the billing period to initiate an auto payment.  Only applicable for accounts with type &#x60;CHARGE_SECURED&#x60;, where the account holder has opted in for auto payment functionality. This value must be lower than or equal the &#x60;grace_period&#x60; setting on the account. If this value is 0, the auto payment will happen on the same day as the statement is generated. Auto payment only occurs if regular payments are not received on time.  | [optional] 
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment. The default will be set to 21 days.  | [optional] [default to 21]
**MinimumPayment** | [**MinimumPaymentFull**](MinimumPaymentFull.md) |  | 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 

## Methods

### NewTemplateFieldsChargeSecured

`func NewTemplateFieldsChargeSecured(accountType AccountType, bankCountry string, currency string, minimumPayment MinimumPaymentFull, ) *TemplateFieldsChargeSecured`

NewTemplateFieldsChargeSecured instantiates a new TemplateFieldsChargeSecured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldsChargeSecuredWithDefaults

`func NewTemplateFieldsChargeSecuredWithDefaults() *TemplateFieldsChargeSecured`

NewTemplateFieldsChargeSecuredWithDefaults instantiates a new TemplateFieldsChargeSecured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *TemplateFieldsChargeSecured) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *TemplateFieldsChargeSecured) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *TemplateFieldsChargeSecured) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetBankCountry

`func (o *TemplateFieldsChargeSecured) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *TemplateFieldsChargeSecured) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *TemplateFieldsChargeSecured) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.


### GetCurrency

`func (o *TemplateFieldsChargeSecured) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TemplateFieldsChargeSecured) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TemplateFieldsChargeSecured) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAutoPaymentPeriod

`func (o *TemplateFieldsChargeSecured) GetAutoPaymentPeriod() int32`

GetAutoPaymentPeriod returns the AutoPaymentPeriod field if non-nil, zero value otherwise.

### GetAutoPaymentPeriodOk

`func (o *TemplateFieldsChargeSecured) GetAutoPaymentPeriodOk() (*int32, bool)`

GetAutoPaymentPeriodOk returns a tuple with the AutoPaymentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPaymentPeriod

`func (o *TemplateFieldsChargeSecured) SetAutoPaymentPeriod(v int32)`

SetAutoPaymentPeriod sets AutoPaymentPeriod field to given value.

### HasAutoPaymentPeriod

`func (o *TemplateFieldsChargeSecured) HasAutoPaymentPeriod() bool`

HasAutoPaymentPeriod returns a boolean if a field has been set.

### GetGracePeriod

`func (o *TemplateFieldsChargeSecured) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *TemplateFieldsChargeSecured) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *TemplateFieldsChargeSecured) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *TemplateFieldsChargeSecured) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *TemplateFieldsChargeSecured) GetMinimumPayment() MinimumPaymentFull`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *TemplateFieldsChargeSecured) GetMinimumPaymentOk() (*MinimumPaymentFull, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *TemplateFieldsChargeSecured) SetMinimumPayment(v MinimumPaymentFull)`

SetMinimumPayment sets MinimumPayment field to given value.


### GetSpendControlIds

`func (o *TemplateFieldsChargeSecured) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *TemplateFieldsChargeSecured) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *TemplateFieldsChargeSecured) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *TemplateFieldsChargeSecured) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


