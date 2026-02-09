# EnhancedRawDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | The monetary amount of the transaction | [optional] 
**CategorizedBy** | Pointer to **int32** | The method used to detrmine the category | [optional] 
**Category** | Pointer to **string** | The category of the transaction | [optional] 
**CategoryGuid** | Pointer to **string** | The unique identifier for the category | [optional] 
**DescribedBy** | Pointer to **int32** | The method used to describe the transaction | [optional] 
**Description** | Pointer to **string** | A human-readable version of &#x60;original_description&#x60; | [optional] 
**ExtendedTransactionType** | Pointer to **string** | The transaction type assigned by the partner | [optional] 
**Id** | Pointer to **int32** | The unique partner-defined identifier for the transaction | [optional] 
**IsBillPay** | Pointer to **bool** | Whether the transaction represents a bill payment | [optional] 
**IsDirectDeposit** | Pointer to **bool** | Whether or not the transaction represents a direct deposit | [optional] 
**IsExpense** | Pointer to **bool** | Whether or not the transaction represents an expense | [optional] 
**IsFee** | Pointer to **bool** | Whether or not the transaction represents a fee | [optional] 
**IsIncome** | Pointer to **bool** | Whether or not the transaction represents income | [optional] 
**IsInternational** | Pointer to **bool** | Whether or not the transaction is international | [optional] 
**IsOverdraftFee** | Pointer to **bool** | Whether or not the transaction is an overdraft fee | [optional] 
**IsPayrollAdvance** | Pointer to **bool** | Whether or not the transaction is a payroll advance | [optional] 
**IsSubscription** | Pointer to **bool** | Whether or not the transaction is a subscription | [optional] 
**Memo** | Pointer to **string** | Additional descriptiive information about the transaction | [optional] 
**MerchantCategoryCode** | Pointer to **int32** | The ISO 18245 category code for the transaction | [optional] 
**MerchantGuid** | Pointer to **string** | The unique identifier for the merchant | [optional] 
**MerchantLocationGuid** | Pointer to **string** | The unique identifier for the merchant location | [optional] 
**OriginalDescription** | Pointer to **string** | The original description for the transaction | [optional] 
**Type** | Pointer to **string** | The type of the transsaction. This will be either &#x60;CREDIT&#x60; or &#x60;DEBIT&#x60; | [optional] 

## Methods

### NewEnhancedRawDetails

`func NewEnhancedRawDetails() *EnhancedRawDetails`

NewEnhancedRawDetails instantiates a new EnhancedRawDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedRawDetailsWithDefaults

`func NewEnhancedRawDetailsWithDefaults() *EnhancedRawDetails`

NewEnhancedRawDetailsWithDefaults instantiates a new EnhancedRawDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EnhancedRawDetails) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnhancedRawDetails) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnhancedRawDetails) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EnhancedRawDetails) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCategorizedBy

`func (o *EnhancedRawDetails) GetCategorizedBy() int32`

GetCategorizedBy returns the CategorizedBy field if non-nil, zero value otherwise.

### GetCategorizedByOk

`func (o *EnhancedRawDetails) GetCategorizedByOk() (*int32, bool)`

GetCategorizedByOk returns a tuple with the CategorizedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorizedBy

`func (o *EnhancedRawDetails) SetCategorizedBy(v int32)`

SetCategorizedBy sets CategorizedBy field to given value.

### HasCategorizedBy

`func (o *EnhancedRawDetails) HasCategorizedBy() bool`

HasCategorizedBy returns a boolean if a field has been set.

### GetCategory

`func (o *EnhancedRawDetails) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EnhancedRawDetails) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EnhancedRawDetails) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EnhancedRawDetails) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryGuid

`func (o *EnhancedRawDetails) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *EnhancedRawDetails) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *EnhancedRawDetails) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *EnhancedRawDetails) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### GetDescribedBy

`func (o *EnhancedRawDetails) GetDescribedBy() int32`

GetDescribedBy returns the DescribedBy field if non-nil, zero value otherwise.

### GetDescribedByOk

`func (o *EnhancedRawDetails) GetDescribedByOk() (*int32, bool)`

GetDescribedByOk returns a tuple with the DescribedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribedBy

`func (o *EnhancedRawDetails) SetDescribedBy(v int32)`

SetDescribedBy sets DescribedBy field to given value.

### HasDescribedBy

`func (o *EnhancedRawDetails) HasDescribedBy() bool`

HasDescribedBy returns a boolean if a field has been set.

### GetDescription

`func (o *EnhancedRawDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnhancedRawDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnhancedRawDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnhancedRawDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtendedTransactionType

`func (o *EnhancedRawDetails) GetExtendedTransactionType() string`

GetExtendedTransactionType returns the ExtendedTransactionType field if non-nil, zero value otherwise.

### GetExtendedTransactionTypeOk

`func (o *EnhancedRawDetails) GetExtendedTransactionTypeOk() (*string, bool)`

GetExtendedTransactionTypeOk returns a tuple with the ExtendedTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedTransactionType

`func (o *EnhancedRawDetails) SetExtendedTransactionType(v string)`

SetExtendedTransactionType sets ExtendedTransactionType field to given value.

### HasExtendedTransactionType

`func (o *EnhancedRawDetails) HasExtendedTransactionType() bool`

HasExtendedTransactionType returns a boolean if a field has been set.

### GetId

`func (o *EnhancedRawDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhancedRawDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhancedRawDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnhancedRawDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBillPay

`func (o *EnhancedRawDetails) GetIsBillPay() bool`

GetIsBillPay returns the IsBillPay field if non-nil, zero value otherwise.

### GetIsBillPayOk

`func (o *EnhancedRawDetails) GetIsBillPayOk() (*bool, bool)`

GetIsBillPayOk returns a tuple with the IsBillPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillPay

`func (o *EnhancedRawDetails) SetIsBillPay(v bool)`

SetIsBillPay sets IsBillPay field to given value.

### HasIsBillPay

`func (o *EnhancedRawDetails) HasIsBillPay() bool`

HasIsBillPay returns a boolean if a field has been set.

### GetIsDirectDeposit

`func (o *EnhancedRawDetails) GetIsDirectDeposit() bool`

GetIsDirectDeposit returns the IsDirectDeposit field if non-nil, zero value otherwise.

### GetIsDirectDepositOk

`func (o *EnhancedRawDetails) GetIsDirectDepositOk() (*bool, bool)`

GetIsDirectDepositOk returns a tuple with the IsDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectDeposit

`func (o *EnhancedRawDetails) SetIsDirectDeposit(v bool)`

SetIsDirectDeposit sets IsDirectDeposit field to given value.

### HasIsDirectDeposit

`func (o *EnhancedRawDetails) HasIsDirectDeposit() bool`

HasIsDirectDeposit returns a boolean if a field has been set.

### GetIsExpense

`func (o *EnhancedRawDetails) GetIsExpense() bool`

GetIsExpense returns the IsExpense field if non-nil, zero value otherwise.

### GetIsExpenseOk

`func (o *EnhancedRawDetails) GetIsExpenseOk() (*bool, bool)`

GetIsExpenseOk returns a tuple with the IsExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpense

`func (o *EnhancedRawDetails) SetIsExpense(v bool)`

SetIsExpense sets IsExpense field to given value.

### HasIsExpense

`func (o *EnhancedRawDetails) HasIsExpense() bool`

HasIsExpense returns a boolean if a field has been set.

### GetIsFee

`func (o *EnhancedRawDetails) GetIsFee() bool`

GetIsFee returns the IsFee field if non-nil, zero value otherwise.

### GetIsFeeOk

`func (o *EnhancedRawDetails) GetIsFeeOk() (*bool, bool)`

GetIsFeeOk returns a tuple with the IsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFee

`func (o *EnhancedRawDetails) SetIsFee(v bool)`

SetIsFee sets IsFee field to given value.

### HasIsFee

`func (o *EnhancedRawDetails) HasIsFee() bool`

HasIsFee returns a boolean if a field has been set.

### GetIsIncome

`func (o *EnhancedRawDetails) GetIsIncome() bool`

GetIsIncome returns the IsIncome field if non-nil, zero value otherwise.

### GetIsIncomeOk

`func (o *EnhancedRawDetails) GetIsIncomeOk() (*bool, bool)`

GetIsIncomeOk returns a tuple with the IsIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncome

`func (o *EnhancedRawDetails) SetIsIncome(v bool)`

SetIsIncome sets IsIncome field to given value.

### HasIsIncome

`func (o *EnhancedRawDetails) HasIsIncome() bool`

HasIsIncome returns a boolean if a field has been set.

### GetIsInternational

`func (o *EnhancedRawDetails) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *EnhancedRawDetails) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *EnhancedRawDetails) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *EnhancedRawDetails) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### GetIsOverdraftFee

`func (o *EnhancedRawDetails) GetIsOverdraftFee() bool`

GetIsOverdraftFee returns the IsOverdraftFee field if non-nil, zero value otherwise.

### GetIsOverdraftFeeOk

`func (o *EnhancedRawDetails) GetIsOverdraftFeeOk() (*bool, bool)`

GetIsOverdraftFeeOk returns a tuple with the IsOverdraftFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdraftFee

`func (o *EnhancedRawDetails) SetIsOverdraftFee(v bool)`

SetIsOverdraftFee sets IsOverdraftFee field to given value.

### HasIsOverdraftFee

`func (o *EnhancedRawDetails) HasIsOverdraftFee() bool`

HasIsOverdraftFee returns a boolean if a field has been set.

### GetIsPayrollAdvance

`func (o *EnhancedRawDetails) GetIsPayrollAdvance() bool`

GetIsPayrollAdvance returns the IsPayrollAdvance field if non-nil, zero value otherwise.

### GetIsPayrollAdvanceOk

`func (o *EnhancedRawDetails) GetIsPayrollAdvanceOk() (*bool, bool)`

GetIsPayrollAdvanceOk returns a tuple with the IsPayrollAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPayrollAdvance

`func (o *EnhancedRawDetails) SetIsPayrollAdvance(v bool)`

SetIsPayrollAdvance sets IsPayrollAdvance field to given value.

### HasIsPayrollAdvance

`func (o *EnhancedRawDetails) HasIsPayrollAdvance() bool`

HasIsPayrollAdvance returns a boolean if a field has been set.

### GetIsSubscription

`func (o *EnhancedRawDetails) GetIsSubscription() bool`

GetIsSubscription returns the IsSubscription field if non-nil, zero value otherwise.

### GetIsSubscriptionOk

`func (o *EnhancedRawDetails) GetIsSubscriptionOk() (*bool, bool)`

GetIsSubscriptionOk returns a tuple with the IsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscription

`func (o *EnhancedRawDetails) SetIsSubscription(v bool)`

SetIsSubscription sets IsSubscription field to given value.

### HasIsSubscription

`func (o *EnhancedRawDetails) HasIsSubscription() bool`

HasIsSubscription returns a boolean if a field has been set.

### GetMemo

`func (o *EnhancedRawDetails) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *EnhancedRawDetails) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *EnhancedRawDetails) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *EnhancedRawDetails) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMerchantCategoryCode

`func (o *EnhancedRawDetails) GetMerchantCategoryCode() int32`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *EnhancedRawDetails) GetMerchantCategoryCodeOk() (*int32, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *EnhancedRawDetails) SetMerchantCategoryCode(v int32)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *EnhancedRawDetails) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetMerchantGuid

`func (o *EnhancedRawDetails) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *EnhancedRawDetails) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *EnhancedRawDetails) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *EnhancedRawDetails) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### GetMerchantLocationGuid

`func (o *EnhancedRawDetails) GetMerchantLocationGuid() string`

GetMerchantLocationGuid returns the MerchantLocationGuid field if non-nil, zero value otherwise.

### GetMerchantLocationGuidOk

`func (o *EnhancedRawDetails) GetMerchantLocationGuidOk() (*string, bool)`

GetMerchantLocationGuidOk returns a tuple with the MerchantLocationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLocationGuid

`func (o *EnhancedRawDetails) SetMerchantLocationGuid(v string)`

SetMerchantLocationGuid sets MerchantLocationGuid field to given value.

### HasMerchantLocationGuid

`func (o *EnhancedRawDetails) HasMerchantLocationGuid() bool`

HasMerchantLocationGuid returns a boolean if a field has been set.

### GetOriginalDescription

`func (o *EnhancedRawDetails) GetOriginalDescription() string`

GetOriginalDescription returns the OriginalDescription field if non-nil, zero value otherwise.

### GetOriginalDescriptionOk

`func (o *EnhancedRawDetails) GetOriginalDescriptionOk() (*string, bool)`

GetOriginalDescriptionOk returns a tuple with the OriginalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDescription

`func (o *EnhancedRawDetails) SetOriginalDescription(v string)`

SetOriginalDescription sets OriginalDescription field to given value.

### HasOriginalDescription

`func (o *EnhancedRawDetails) HasOriginalDescription() bool`

HasOriginalDescription returns a boolean if a field has been set.

### GetType

`func (o *EnhancedRawDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnhancedRawDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnhancedRawDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnhancedRawDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


