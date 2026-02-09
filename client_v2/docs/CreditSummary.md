# CreditSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOverLimit** | Pointer to **int64** | The portion of the posted account balance that exceeds the account&#39;s credit limit, in ISO 4217 minor currency units. | [optional] 
**AmountPastDue** | Pointer to **int64** | Any outstanding balance from a previous statement for this account in ISO 4217 minor currency units.  | [optional] 
**Apr** | Pointer to **int64** | The annual percentage rate of the credit account, in basis points. | [optional] 
**BalanceForInterest** | Pointer to **int64** | Balance on which interest is computed during the time interval covered by the statement. | [optional] 
**CreditLimit** | Pointer to **int64** | The credit limit of the credit account, in ISO 4217 minor currency units. Not applicable for charge secured accounts. | [optional] 
**Disclosure** | Pointer to **string** | A suggested regulatory disclosure to display on the statement. | [optional] 
**DisputedBalance** | Pointer to **int64** | The disputed balance on the credit account in ISO 4217 minor currency units. | [optional] 
**Fees** | Pointer to **int64** | The total fees charged on the credit account in ISO 4217 minor currency units during the time interval covered by the statement. | [optional] 
**FeesYtd** | Pointer to **int64** | The total fees charged on the credit account in ISO 4217 minor currency units for the year to date. | [optional] 
**Interest** | Pointer to **int64** | The total interest charged on the credit account in ISO 4217 minor currency units during the time interval covered by the statement. | [optional] 
**InterestYtd** | Pointer to **int64** | The total interest charged on the credit account in ISO 4217 minor currency units for the year to date. | [optional] 
**IsPastDue** | Pointer to **bool** | Whether or not the credit account is past due on payment. | [optional] 
**LastPaymentDate** | Pointer to **string** | The date of the last payment received. | [optional] 
**MinimumPaymentDue** | Pointer to **int64** | The minimum payment amount due by the payment due date in ISO 4217 minor currency units. | [optional] 
**PaymentDue** | Pointer to **int64** | The total amount due by the payment due date in ISO 4217 minor currency units. | [optional] 
**PaymentDueDate** | Pointer to **string** | The date the next payment is due. | [optional] 
**PaymentsReceived** | Pointer to **int64** | The sum of all payments received in ISO 4217 minor currency units. | [optional] 

## Methods

### NewCreditSummary

`func NewCreditSummary() *CreditSummary`

NewCreditSummary instantiates a new CreditSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditSummaryWithDefaults

`func NewCreditSummaryWithDefaults() *CreditSummary`

NewCreditSummaryWithDefaults instantiates a new CreditSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOverLimit

`func (o *CreditSummary) GetAmountOverLimit() int64`

GetAmountOverLimit returns the AmountOverLimit field if non-nil, zero value otherwise.

### GetAmountOverLimitOk

`func (o *CreditSummary) GetAmountOverLimitOk() (*int64, bool)`

GetAmountOverLimitOk returns a tuple with the AmountOverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOverLimit

`func (o *CreditSummary) SetAmountOverLimit(v int64)`

SetAmountOverLimit sets AmountOverLimit field to given value.

### HasAmountOverLimit

`func (o *CreditSummary) HasAmountOverLimit() bool`

HasAmountOverLimit returns a boolean if a field has been set.

### GetAmountPastDue

`func (o *CreditSummary) GetAmountPastDue() int64`

GetAmountPastDue returns the AmountPastDue field if non-nil, zero value otherwise.

### GetAmountPastDueOk

`func (o *CreditSummary) GetAmountPastDueOk() (*int64, bool)`

GetAmountPastDueOk returns a tuple with the AmountPastDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPastDue

`func (o *CreditSummary) SetAmountPastDue(v int64)`

SetAmountPastDue sets AmountPastDue field to given value.

### HasAmountPastDue

`func (o *CreditSummary) HasAmountPastDue() bool`

HasAmountPastDue returns a boolean if a field has been set.

### GetApr

`func (o *CreditSummary) GetApr() int64`

GetApr returns the Apr field if non-nil, zero value otherwise.

### GetAprOk

`func (o *CreditSummary) GetAprOk() (*int64, bool)`

GetAprOk returns a tuple with the Apr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApr

`func (o *CreditSummary) SetApr(v int64)`

SetApr sets Apr field to given value.

### HasApr

`func (o *CreditSummary) HasApr() bool`

HasApr returns a boolean if a field has been set.

### GetBalanceForInterest

`func (o *CreditSummary) GetBalanceForInterest() int64`

GetBalanceForInterest returns the BalanceForInterest field if non-nil, zero value otherwise.

### GetBalanceForInterestOk

`func (o *CreditSummary) GetBalanceForInterestOk() (*int64, bool)`

GetBalanceForInterestOk returns a tuple with the BalanceForInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceForInterest

`func (o *CreditSummary) SetBalanceForInterest(v int64)`

SetBalanceForInterest sets BalanceForInterest field to given value.

### HasBalanceForInterest

`func (o *CreditSummary) HasBalanceForInterest() bool`

HasBalanceForInterest returns a boolean if a field has been set.

### GetCreditLimit

`func (o *CreditSummary) GetCreditLimit() int64`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *CreditSummary) GetCreditLimitOk() (*int64, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *CreditSummary) SetCreditLimit(v int64)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *CreditSummary) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetDisclosure

`func (o *CreditSummary) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *CreditSummary) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *CreditSummary) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *CreditSummary) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetDisputedBalance

`func (o *CreditSummary) GetDisputedBalance() int64`

GetDisputedBalance returns the DisputedBalance field if non-nil, zero value otherwise.

### GetDisputedBalanceOk

`func (o *CreditSummary) GetDisputedBalanceOk() (*int64, bool)`

GetDisputedBalanceOk returns a tuple with the DisputedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedBalance

`func (o *CreditSummary) SetDisputedBalance(v int64)`

SetDisputedBalance sets DisputedBalance field to given value.

### HasDisputedBalance

`func (o *CreditSummary) HasDisputedBalance() bool`

HasDisputedBalance returns a boolean if a field has been set.

### GetFees

`func (o *CreditSummary) GetFees() int64`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *CreditSummary) GetFeesOk() (*int64, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *CreditSummary) SetFees(v int64)`

SetFees sets Fees field to given value.

### HasFees

`func (o *CreditSummary) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetFeesYtd

`func (o *CreditSummary) GetFeesYtd() int64`

GetFeesYtd returns the FeesYtd field if non-nil, zero value otherwise.

### GetFeesYtdOk

`func (o *CreditSummary) GetFeesYtdOk() (*int64, bool)`

GetFeesYtdOk returns a tuple with the FeesYtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesYtd

`func (o *CreditSummary) SetFeesYtd(v int64)`

SetFeesYtd sets FeesYtd field to given value.

### HasFeesYtd

`func (o *CreditSummary) HasFeesYtd() bool`

HasFeesYtd returns a boolean if a field has been set.

### GetInterest

`func (o *CreditSummary) GetInterest() int64`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *CreditSummary) GetInterestOk() (*int64, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *CreditSummary) SetInterest(v int64)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *CreditSummary) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetInterestYtd

`func (o *CreditSummary) GetInterestYtd() int64`

GetInterestYtd returns the InterestYtd field if non-nil, zero value otherwise.

### GetInterestYtdOk

`func (o *CreditSummary) GetInterestYtdOk() (*int64, bool)`

GetInterestYtdOk returns a tuple with the InterestYtd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestYtd

`func (o *CreditSummary) SetInterestYtd(v int64)`

SetInterestYtd sets InterestYtd field to given value.

### HasInterestYtd

`func (o *CreditSummary) HasInterestYtd() bool`

HasInterestYtd returns a boolean if a field has been set.

### GetIsPastDue

`func (o *CreditSummary) GetIsPastDue() bool`

GetIsPastDue returns the IsPastDue field if non-nil, zero value otherwise.

### GetIsPastDueOk

`func (o *CreditSummary) GetIsPastDueOk() (*bool, bool)`

GetIsPastDueOk returns a tuple with the IsPastDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPastDue

`func (o *CreditSummary) SetIsPastDue(v bool)`

SetIsPastDue sets IsPastDue field to given value.

### HasIsPastDue

`func (o *CreditSummary) HasIsPastDue() bool`

HasIsPastDue returns a boolean if a field has been set.

### GetLastPaymentDate

`func (o *CreditSummary) GetLastPaymentDate() string`

GetLastPaymentDate returns the LastPaymentDate field if non-nil, zero value otherwise.

### GetLastPaymentDateOk

`func (o *CreditSummary) GetLastPaymentDateOk() (*string, bool)`

GetLastPaymentDateOk returns a tuple with the LastPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentDate

`func (o *CreditSummary) SetLastPaymentDate(v string)`

SetLastPaymentDate sets LastPaymentDate field to given value.

### HasLastPaymentDate

`func (o *CreditSummary) HasLastPaymentDate() bool`

HasLastPaymentDate returns a boolean if a field has been set.

### GetMinimumPaymentDue

`func (o *CreditSummary) GetMinimumPaymentDue() int64`

GetMinimumPaymentDue returns the MinimumPaymentDue field if non-nil, zero value otherwise.

### GetMinimumPaymentDueOk

`func (o *CreditSummary) GetMinimumPaymentDueOk() (*int64, bool)`

GetMinimumPaymentDueOk returns a tuple with the MinimumPaymentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPaymentDue

`func (o *CreditSummary) SetMinimumPaymentDue(v int64)`

SetMinimumPaymentDue sets MinimumPaymentDue field to given value.

### HasMinimumPaymentDue

`func (o *CreditSummary) HasMinimumPaymentDue() bool`

HasMinimumPaymentDue returns a boolean if a field has been set.

### GetPaymentDue

`func (o *CreditSummary) GetPaymentDue() int64`

GetPaymentDue returns the PaymentDue field if non-nil, zero value otherwise.

### GetPaymentDueOk

`func (o *CreditSummary) GetPaymentDueOk() (*int64, bool)`

GetPaymentDueOk returns a tuple with the PaymentDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDue

`func (o *CreditSummary) SetPaymentDue(v int64)`

SetPaymentDue sets PaymentDue field to given value.

### HasPaymentDue

`func (o *CreditSummary) HasPaymentDue() bool`

HasPaymentDue returns a boolean if a field has been set.

### GetPaymentDueDate

`func (o *CreditSummary) GetPaymentDueDate() string`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *CreditSummary) GetPaymentDueDateOk() (*string, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *CreditSummary) SetPaymentDueDate(v string)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *CreditSummary) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### GetPaymentsReceived

`func (o *CreditSummary) GetPaymentsReceived() int64`

GetPaymentsReceived returns the PaymentsReceived field if non-nil, zero value otherwise.

### GetPaymentsReceivedOk

`func (o *CreditSummary) GetPaymentsReceivedOk() (*int64, bool)`

GetPaymentsReceivedOk returns a tuple with the PaymentsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsReceived

`func (o *CreditSummary) SetPaymentsReceived(v int64)`

SetPaymentsReceived sets PaymentsReceived field to given value.

### HasPaymentsReceived

`func (o *CreditSummary) HasPaymentsReceived() bool`

HasPaymentsReceived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


