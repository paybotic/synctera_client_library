# ExternalAccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **NullableInt64** | The amount of funds available to be withdrawn from the account, as determined by the financial institution.  This is an integer in the minor currency unit (e.g. cents): 1025 means $10.25.  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Currency** | **string** | ISO 4217 alphabetic currency code | 
**Current** | Pointer to **NullableInt64** | For a &#x60;DEPOSITORY&#x60; account, this is the total amount of funds in the account.  For a &#x60;CREDIT&#x60; account, this is the amount owing. If negative, the lender owes the account holder.  This is an integer in the minor currency unit (e.g. cents): -2500 means $25.00 owed to the account holder.  | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The last time Synctera has fetched this balance from a vendor | [optional] 
**LastUpdatedTimeVendor** | Pointer to **time.Time** | Vendor timestamp of when the balance was last updated on the vendor side | [optional] 
**Limit** | Pointer to **NullableInt64** | For &#x60;DEPOSITORY&#x60; accounts, this is the pre-arranged overdraft limit, commonly used in Europe. In North America this is typically not set for depository accounts.  For &#x60;CREDIT&#x60; accounts, this is the credit limit on the account.  This is an integer in the minor currency unit (e.g. cents): 10000 means $100.00  | [optional] 
**TransactionsLastUpdatedTime** | Pointer to **time.Time** | The last time Synctera has fetched transactions from a vendor | [optional] 

## Methods

### NewExternalAccountBalance

`func NewExternalAccountBalance(currency string, ) *ExternalAccountBalance`

NewExternalAccountBalance instantiates a new ExternalAccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountBalanceWithDefaults

`func NewExternalAccountBalanceWithDefaults() *ExternalAccountBalance`

NewExternalAccountBalanceWithDefaults instantiates a new ExternalAccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ExternalAccountBalance) GetAvailable() int64`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ExternalAccountBalance) GetAvailableOk() (*int64, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ExternalAccountBalance) SetAvailable(v int64)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ExternalAccountBalance) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *ExternalAccountBalance) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *ExternalAccountBalance) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCreationTime

`func (o *ExternalAccountBalance) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ExternalAccountBalance) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ExternalAccountBalance) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ExternalAccountBalance) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *ExternalAccountBalance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExternalAccountBalance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExternalAccountBalance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrent

`func (o *ExternalAccountBalance) GetCurrent() int64`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ExternalAccountBalance) GetCurrentOk() (*int64, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ExternalAccountBalance) SetCurrent(v int64)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *ExternalAccountBalance) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *ExternalAccountBalance) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *ExternalAccountBalance) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetLastUpdatedTime

`func (o *ExternalAccountBalance) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ExternalAccountBalance) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ExternalAccountBalance) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *ExternalAccountBalance) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLastUpdatedTimeVendor

`func (o *ExternalAccountBalance) GetLastUpdatedTimeVendor() time.Time`

GetLastUpdatedTimeVendor returns the LastUpdatedTimeVendor field if non-nil, zero value otherwise.

### GetLastUpdatedTimeVendorOk

`func (o *ExternalAccountBalance) GetLastUpdatedTimeVendorOk() (*time.Time, bool)`

GetLastUpdatedTimeVendorOk returns a tuple with the LastUpdatedTimeVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeVendor

`func (o *ExternalAccountBalance) SetLastUpdatedTimeVendor(v time.Time)`

SetLastUpdatedTimeVendor sets LastUpdatedTimeVendor field to given value.

### HasLastUpdatedTimeVendor

`func (o *ExternalAccountBalance) HasLastUpdatedTimeVendor() bool`

HasLastUpdatedTimeVendor returns a boolean if a field has been set.

### GetLimit

`func (o *ExternalAccountBalance) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ExternalAccountBalance) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ExternalAccountBalance) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ExternalAccountBalance) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ExternalAccountBalance) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ExternalAccountBalance) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetTransactionsLastUpdatedTime

`func (o *ExternalAccountBalance) GetTransactionsLastUpdatedTime() time.Time`

GetTransactionsLastUpdatedTime returns the TransactionsLastUpdatedTime field if non-nil, zero value otherwise.

### GetTransactionsLastUpdatedTimeOk

`func (o *ExternalAccountBalance) GetTransactionsLastUpdatedTimeOk() (*time.Time, bool)`

GetTransactionsLastUpdatedTimeOk returns a tuple with the TransactionsLastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionsLastUpdatedTime

`func (o *ExternalAccountBalance) SetTransactionsLastUpdatedTime(v time.Time)`

SetTransactionsLastUpdatedTime sets TransactionsLastUpdatedTime field to given value.

### HasTransactionsLastUpdatedTime

`func (o *ExternalAccountBalance) HasTransactionsLastUpdatedTime() bool`

HasTransactionsLastUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


