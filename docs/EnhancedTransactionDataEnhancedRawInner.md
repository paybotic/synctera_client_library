# EnhancedTransactionDataEnhancedRawInner

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

### NewEnhancedTransactionDataEnhancedRawInner

`func NewEnhancedTransactionDataEnhancedRawInner() *EnhancedTransactionDataEnhancedRawInner`

NewEnhancedTransactionDataEnhancedRawInner instantiates a new EnhancedTransactionDataEnhancedRawInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnhancedTransactionDataEnhancedRawInnerWithDefaults

`func NewEnhancedTransactionDataEnhancedRawInnerWithDefaults() *EnhancedTransactionDataEnhancedRawInner`

NewEnhancedTransactionDataEnhancedRawInnerWithDefaults instantiates a new EnhancedTransactionDataEnhancedRawInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EnhancedTransactionDataEnhancedRawInner) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EnhancedTransactionDataEnhancedRawInner) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EnhancedTransactionDataEnhancedRawInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCategorizedBy

`func (o *EnhancedTransactionDataEnhancedRawInner) GetCategorizedBy() int32`

GetCategorizedBy returns the CategorizedBy field if non-nil, zero value otherwise.

### GetCategorizedByOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetCategorizedByOk() (*int32, bool)`

GetCategorizedByOk returns a tuple with the CategorizedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorizedBy

`func (o *EnhancedTransactionDataEnhancedRawInner) SetCategorizedBy(v int32)`

SetCategorizedBy sets CategorizedBy field to given value.

### HasCategorizedBy

`func (o *EnhancedTransactionDataEnhancedRawInner) HasCategorizedBy() bool`

HasCategorizedBy returns a boolean if a field has been set.

### GetCategory

`func (o *EnhancedTransactionDataEnhancedRawInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EnhancedTransactionDataEnhancedRawInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EnhancedTransactionDataEnhancedRawInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCategoryGuid

`func (o *EnhancedTransactionDataEnhancedRawInner) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *EnhancedTransactionDataEnhancedRawInner) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *EnhancedTransactionDataEnhancedRawInner) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### GetDescribedBy

`func (o *EnhancedTransactionDataEnhancedRawInner) GetDescribedBy() int32`

GetDescribedBy returns the DescribedBy field if non-nil, zero value otherwise.

### GetDescribedByOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetDescribedByOk() (*int32, bool)`

GetDescribedByOk returns a tuple with the DescribedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribedBy

`func (o *EnhancedTransactionDataEnhancedRawInner) SetDescribedBy(v int32)`

SetDescribedBy sets DescribedBy field to given value.

### HasDescribedBy

`func (o *EnhancedTransactionDataEnhancedRawInner) HasDescribedBy() bool`

HasDescribedBy returns a boolean if a field has been set.

### GetDescription

`func (o *EnhancedTransactionDataEnhancedRawInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnhancedTransactionDataEnhancedRawInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnhancedTransactionDataEnhancedRawInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtendedTransactionType

`func (o *EnhancedTransactionDataEnhancedRawInner) GetExtendedTransactionType() string`

GetExtendedTransactionType returns the ExtendedTransactionType field if non-nil, zero value otherwise.

### GetExtendedTransactionTypeOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetExtendedTransactionTypeOk() (*string, bool)`

GetExtendedTransactionTypeOk returns a tuple with the ExtendedTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedTransactionType

`func (o *EnhancedTransactionDataEnhancedRawInner) SetExtendedTransactionType(v string)`

SetExtendedTransactionType sets ExtendedTransactionType field to given value.

### HasExtendedTransactionType

`func (o *EnhancedTransactionDataEnhancedRawInner) HasExtendedTransactionType() bool`

HasExtendedTransactionType returns a boolean if a field has been set.

### GetId

`func (o *EnhancedTransactionDataEnhancedRawInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnhancedTransactionDataEnhancedRawInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnhancedTransactionDataEnhancedRawInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBillPay

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsBillPay() bool`

GetIsBillPay returns the IsBillPay field if non-nil, zero value otherwise.

### GetIsBillPayOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsBillPayOk() (*bool, bool)`

GetIsBillPayOk returns a tuple with the IsBillPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillPay

`func (o *EnhancedTransactionDataEnhancedRawInner) SetIsBillPay(v bool)`

SetIsBillPay sets IsBillPay field to given value.

### HasIsBillPay

`func (o *EnhancedTransactionDataEnhancedRawInner) HasIsBillPay() bool`

HasIsBillPay returns a boolean if a field has been set.

### GetIsDirectDeposit

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsDirectDeposit() bool`

GetIsDirectDeposit returns the IsDirectDeposit field if non-nil, zero value otherwise.

### GetIsDirectDepositOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsDirectDepositOk() (*bool, bool)`

GetIsDirectDepositOk returns a tuple with the IsDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectDeposit

`func (o *EnhancedTransactionDataEnhancedRawInner) SetIsDirectDeposit(v bool)`

SetIsDirectDeposit sets IsDirectDeposit field to given value.

### HasIsDirectDeposit

`func (o *EnhancedTransactionDataEnhancedRawInner) HasIsDirectDeposit() bool`

HasIsDirectDeposit returns a boolean if a field has been set.

### GetIsExpense

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsExpense() bool`

GetIsExpense returns the IsExpense field if non-nil, zero value otherwise.

### GetIsExpenseOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsExpenseOk() (*bool, bool)`

GetIsExpenseOk returns a tuple with the IsExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpense

`func (o *EnhancedTransactionDataEnhancedRawInner) SetIsExpense(v bool)`

SetIsExpense sets IsExpense field to given value.

### HasIsExpense

`func (o *EnhancedTransactionDataEnhancedRawInner) HasIsExpense() bool`

HasIsExpense returns a boolean if a field has been set.

### GetIsFee

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsFee() bool`

GetIsFee returns the IsFee field if non-nil, zero value otherwise.

### GetIsFeeOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsFeeOk() (*bool, bool)`

GetIsFeeOk returns a tuple with the IsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFee

`func (o *EnhancedTransactionDataEnhancedRawInner) SetIsFee(v bool)`

SetIsFee sets IsFee field to given value.

### HasIsFee

`func (o *EnhancedTransactionDataEnhancedRawInner) HasIsFee() bool`

HasIsFee returns a boolean if a field has been set.

### GetIsIncome

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsIncome() bool`

GetIsIncome returns the IsIncome field if non-nil, zero value otherwise.

### GetIsIncomeOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsIncomeOk() (*bool, bool)`

GetIsIncomeOk returns a tuple with the IsIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncome

`func (o *EnhancedTransactionDataEnhancedRawInner) SetIsIncome(v bool)`

SetIsIncome sets IsIncome field to given value.

### HasIsIncome

`func (o *EnhancedTransactionDataEnhancedRawInner) HasIsIncome() bool`

HasIsIncome returns a boolean if a field has been set.

### GetIsInternational

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *EnhancedTransactionDataEnhancedRawInner) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *EnhancedTransactionDataEnhancedRawInner) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### GetIsOverdraftFee

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsOverdraftFee() bool`

GetIsOverdraftFee returns the IsOverdraftFee field if non-nil, zero value otherwise.

### GetIsOverdraftFeeOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsOverdraftFeeOk() (*bool, bool)`

GetIsOverdraftFeeOk returns a tuple with the IsOverdraftFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdraftFee

`func (o *EnhancedTransactionDataEnhancedRawInner) SetIsOverdraftFee(v bool)`

SetIsOverdraftFee sets IsOverdraftFee field to given value.

### HasIsOverdraftFee

`func (o *EnhancedTransactionDataEnhancedRawInner) HasIsOverdraftFee() bool`

HasIsOverdraftFee returns a boolean if a field has been set.

### GetIsPayrollAdvance

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsPayrollAdvance() bool`

GetIsPayrollAdvance returns the IsPayrollAdvance field if non-nil, zero value otherwise.

### GetIsPayrollAdvanceOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsPayrollAdvanceOk() (*bool, bool)`

GetIsPayrollAdvanceOk returns a tuple with the IsPayrollAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPayrollAdvance

`func (o *EnhancedTransactionDataEnhancedRawInner) SetIsPayrollAdvance(v bool)`

SetIsPayrollAdvance sets IsPayrollAdvance field to given value.

### HasIsPayrollAdvance

`func (o *EnhancedTransactionDataEnhancedRawInner) HasIsPayrollAdvance() bool`

HasIsPayrollAdvance returns a boolean if a field has been set.

### GetIsSubscription

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsSubscription() bool`

GetIsSubscription returns the IsSubscription field if non-nil, zero value otherwise.

### GetIsSubscriptionOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetIsSubscriptionOk() (*bool, bool)`

GetIsSubscriptionOk returns a tuple with the IsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscription

`func (o *EnhancedTransactionDataEnhancedRawInner) SetIsSubscription(v bool)`

SetIsSubscription sets IsSubscription field to given value.

### HasIsSubscription

`func (o *EnhancedTransactionDataEnhancedRawInner) HasIsSubscription() bool`

HasIsSubscription returns a boolean if a field has been set.

### GetMemo

`func (o *EnhancedTransactionDataEnhancedRawInner) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *EnhancedTransactionDataEnhancedRawInner) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *EnhancedTransactionDataEnhancedRawInner) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMerchantCategoryCode

`func (o *EnhancedTransactionDataEnhancedRawInner) GetMerchantCategoryCode() int32`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetMerchantCategoryCodeOk() (*int32, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *EnhancedTransactionDataEnhancedRawInner) SetMerchantCategoryCode(v int32)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *EnhancedTransactionDataEnhancedRawInner) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetMerchantGuid

`func (o *EnhancedTransactionDataEnhancedRawInner) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *EnhancedTransactionDataEnhancedRawInner) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *EnhancedTransactionDataEnhancedRawInner) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### GetMerchantLocationGuid

`func (o *EnhancedTransactionDataEnhancedRawInner) GetMerchantLocationGuid() string`

GetMerchantLocationGuid returns the MerchantLocationGuid field if non-nil, zero value otherwise.

### GetMerchantLocationGuidOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetMerchantLocationGuidOk() (*string, bool)`

GetMerchantLocationGuidOk returns a tuple with the MerchantLocationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLocationGuid

`func (o *EnhancedTransactionDataEnhancedRawInner) SetMerchantLocationGuid(v string)`

SetMerchantLocationGuid sets MerchantLocationGuid field to given value.

### HasMerchantLocationGuid

`func (o *EnhancedTransactionDataEnhancedRawInner) HasMerchantLocationGuid() bool`

HasMerchantLocationGuid returns a boolean if a field has been set.

### GetOriginalDescription

`func (o *EnhancedTransactionDataEnhancedRawInner) GetOriginalDescription() string`

GetOriginalDescription returns the OriginalDescription field if non-nil, zero value otherwise.

### GetOriginalDescriptionOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetOriginalDescriptionOk() (*string, bool)`

GetOriginalDescriptionOk returns a tuple with the OriginalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDescription

`func (o *EnhancedTransactionDataEnhancedRawInner) SetOriginalDescription(v string)`

SetOriginalDescription sets OriginalDescription field to given value.

### HasOriginalDescription

`func (o *EnhancedTransactionDataEnhancedRawInner) HasOriginalDescription() bool`

HasOriginalDescription returns a boolean if a field has been set.

### GetType

`func (o *EnhancedTransactionDataEnhancedRawInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnhancedTransactionDataEnhancedRawInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnhancedTransactionDataEnhancedRawInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnhancedTransactionDataEnhancedRawInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


