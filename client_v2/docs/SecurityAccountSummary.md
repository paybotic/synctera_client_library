# SecurityAccountSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The unique identifier of the backing account. | [optional] 
**AccountNumber** | Pointer to **string** | Account number of the backing account. | [optional] 
**AccountType** | Pointer to **string** | The type of the account. It refers to the backing account.  | [optional] 
**Apy** | Pointer to **int64** | The annual percentage yield (APY) for the security account for this statement period, in basis points. For example, an APY of 5.5% will display as 550.  | [optional] 
**ClosingBalance** | Pointer to **int64** | The security account balance at the end of the statement period, in ISO 4217 minor currency units. For example, $1,000 USD will be displayed as 100000. | [optional] 
**Disclosure** | Pointer to **string** |  | [optional] 
**Interest** | Pointer to **int64** | The total interest earned by the security account for this statement period in ISO 4217 minor currency units. For example, $1.50 USD of interest will be displayed as 150.  | [optional] 
**InterestPreviousMonth** | Pointer to **int64** | The total interest earned by the security account in the previous statement period in ISO 4217 minor currency units. For example, $1.50 USD of interest will be displayed as 150.  | [optional] 
**InterestPreviousYear** | Pointer to **int64** | The total interest earned by the security account in the previous year in ISO 4217 minor currency units. For example, $100 USD of interest will be displayed as 10000.  | [optional] 
**InterestYtd** | Pointer to **int64** | The total interest earned by the security account for this year to date in ISO 4217 minor currency units. For example, $100 USD of interest will be displayed as 10000.  | [optional] 
**OpeningBalance** | Pointer to **int64** | The security account balance at the start of the statement period, in ISO 4217 minor currency units. For example, $1,000 USD will be displayed as 100000. | [optional] 

## Methods

### NewSecurityAccountSummary

`func NewSecurityAccountSummary() *SecurityAccountSummary`

NewSecurityAccountSummary instantiates a new SecurityAccountSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccountSummaryWithDefaults

`func NewSecurityAccountSummaryWithDefaults() *SecurityAccountSummary`

NewSecurityAccountSummaryWithDefaults instantiates a new SecurityAccountSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SecurityAccountSummary) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SecurityAccountSummary) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SecurityAccountSummary) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SecurityAccountSummary) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *SecurityAccountSummary) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *SecurityAccountSummary) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *SecurityAccountSummary) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *SecurityAccountSummary) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountType

`func (o *SecurityAccountSummary) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *SecurityAccountSummary) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *SecurityAccountSummary) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *SecurityAccountSummary) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetApy

`func (o *SecurityAccountSummary) GetApy() int64`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *SecurityAccountSummary) GetApyOk() (*int64, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *SecurityAccountSummary) SetApy(v int64)`

SetApy sets Apy field to given value.

### HasApy

`func (o *SecurityAccountSummary) HasApy() bool`

HasApy returns a boolean if a field has been set.

### GetClosingBalance

`func (o *SecurityAccountSummary) GetClosingBalance() int64`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *SecurityAccountSummary) GetClosingBalanceOk() (*int64, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *SecurityAccountSummary) SetClosingBalance(v int64)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *SecurityAccountSummary) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetDisclosure

`func (o *SecurityAccountSummary) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *SecurityAccountSummary) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *SecurityAccountSummary) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *SecurityAccountSummary) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetInterest

`func (o *SecurityAccountSummary) GetInterest() int64`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *SecurityAccountSummary) GetInterestOk() (*int64, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *SecurityAccountSummary) SetInterest(v int64)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *SecurityAccountSummary) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetInterestPreviousMonth

`func (o *SecurityAccountSummary) GetInterestPreviousMonth() int64`

GetInterestPreviousMonth returns the InterestPreviousMonth field if non-nil, zero value otherwise.

### GetInterestPreviousMonthOk

`func (o *SecurityAccountSummary) GetInterestPreviousMonthOk() (*int64, bool)`

GetInterestPreviousMonthOk returns a tuple with the InterestPreviousMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPreviousMonth

`func (o *SecurityAccountSummary) SetInterestPreviousMonth(v int64)`

SetInterestPreviousMonth sets InterestPreviousMonth field to given value.

### HasInterestPreviousMonth

`func (o *SecurityAccountSummary) HasInterestPreviousMonth() bool`

HasInterestPreviousMonth returns a boolean if a field has been set.

### GetInterestPreviousYear

`func (o *SecurityAccountSummary) GetInterestPreviousYear() int64`

GetInterestPreviousYear returns the InterestPreviousYear field if non-nil, zero value otherwise.

### GetInterestPreviousYearOk

`func (o *SecurityAccountSummary) GetInterestPreviousYearOk() (*int64, bool)`

GetInterestPreviousYearOk returns a tuple with the InterestPreviousYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPreviousYear

`func (o *SecurityAccountSummary) SetInterestPreviousYear(v int64)`

SetInterestPreviousYear sets InterestPreviousYear field to given value.

### HasInterestPreviousYear

`func (o *SecurityAccountSummary) HasInterestPreviousYear() bool`

HasInterestPreviousYear returns a boolean if a field has been set.

### GetInterestYtd

`func (o *SecurityAccountSummary) GetInterestYtd() int64`

GetInterestYtd returns the InterestYtd field if non-nil, zero value otherwise.

### GetInterestYtdOk

`func (o *SecurityAccountSummary) GetInterestYtdOk() (*int64, bool)`

GetInterestYtdOk returns a tuple with the InterestYtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestYtd

`func (o *SecurityAccountSummary) SetInterestYtd(v int64)`

SetInterestYtd sets InterestYtd field to given value.

### HasInterestYtd

`func (o *SecurityAccountSummary) HasInterestYtd() bool`

HasInterestYtd returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *SecurityAccountSummary) GetOpeningBalance() int64`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *SecurityAccountSummary) GetOpeningBalanceOk() (*int64, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *SecurityAccountSummary) SetOpeningBalance(v int64)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *SecurityAccountSummary) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


