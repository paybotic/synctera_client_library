# GatewayAuthorizationEnhancedTransactionRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableFloat32** |  | [optional] 
**CategorizedBy** | Pointer to **NullableInt32** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**CategoryGuid** | Pointer to **NullableString** |  | [optional] 
**DescribedBy** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExtendedTransactionType** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**IsBillPay** | Pointer to **NullableBool** |  | [optional] 
**IsDirectDeposit** | Pointer to **NullableBool** |  | [optional] 
**IsExpense** | Pointer to **NullableBool** |  | [optional] 
**IsFee** | Pointer to **NullableBool** |  | [optional] 
**IsIncome** | Pointer to **NullableBool** |  | [optional] 
**IsInternational** | Pointer to **NullableBool** |  | [optional] 
**IsOverdraftFee** | Pointer to **NullableBool** |  | [optional] 
**IsPayrollAdvance** | Pointer to **NullableBool** |  | [optional] 
**IsSubscription** | Pointer to **NullableBool** |  | [optional] 
**Memo** | Pointer to **NullableString** |  | [optional] 
**MerchantCategoryCode** | Pointer to **string** |  | [optional] 
**MerchantGuid** | Pointer to **NullableString** |  | [optional] 
**MerchantLocationGuid** | Pointer to **NullableString** |  | [optional] 
**OriginalDescription** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGatewayAuthorizationEnhancedTransactionRaw

`func NewGatewayAuthorizationEnhancedTransactionRaw() *GatewayAuthorizationEnhancedTransactionRaw`

NewGatewayAuthorizationEnhancedTransactionRaw instantiates a new GatewayAuthorizationEnhancedTransactionRaw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayAuthorizationEnhancedTransactionRawWithDefaults

`func NewGatewayAuthorizationEnhancedTransactionRawWithDefaults() *GatewayAuthorizationEnhancedTransactionRaw`

NewGatewayAuthorizationEnhancedTransactionRawWithDefaults instantiates a new GatewayAuthorizationEnhancedTransactionRaw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCategorizedBy

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetCategorizedBy() int32`

GetCategorizedBy returns the CategorizedBy field if non-nil, zero value otherwise.

### GetCategorizedByOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetCategorizedByOk() (*int32, bool)`

GetCategorizedByOk returns a tuple with the CategorizedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorizedBy

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetCategorizedBy(v int32)`

SetCategorizedBy sets CategorizedBy field to given value.

### HasCategorizedBy

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasCategorizedBy() bool`

HasCategorizedBy returns a boolean if a field has been set.

### SetCategorizedByNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetCategorizedByNil(b bool)`

 SetCategorizedByNil sets the value for CategorizedBy to be an explicit nil

### UnsetCategorizedBy
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetCategorizedBy()`

UnsetCategorizedBy ensures that no value is present for CategorizedBy, not even an explicit nil
### GetCategory

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCategoryGuid

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetCategoryGuid() string`

GetCategoryGuid returns the CategoryGuid field if non-nil, zero value otherwise.

### GetCategoryGuidOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetCategoryGuidOk() (*string, bool)`

GetCategoryGuidOk returns a tuple with the CategoryGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryGuid

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetCategoryGuid(v string)`

SetCategoryGuid sets CategoryGuid field to given value.

### HasCategoryGuid

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasCategoryGuid() bool`

HasCategoryGuid returns a boolean if a field has been set.

### SetCategoryGuidNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetCategoryGuidNil(b bool)`

 SetCategoryGuidNil sets the value for CategoryGuid to be an explicit nil

### UnsetCategoryGuid
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetCategoryGuid()`

UnsetCategoryGuid ensures that no value is present for CategoryGuid, not even an explicit nil
### GetDescribedBy

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetDescribedBy() int32`

GetDescribedBy returns the DescribedBy field if non-nil, zero value otherwise.

### GetDescribedByOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetDescribedByOk() (*int32, bool)`

GetDescribedByOk returns a tuple with the DescribedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribedBy

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetDescribedBy(v int32)`

SetDescribedBy sets DescribedBy field to given value.

### HasDescribedBy

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasDescribedBy() bool`

HasDescribedBy returns a boolean if a field has been set.

### SetDescribedByNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetDescribedByNil(b bool)`

 SetDescribedByNil sets the value for DescribedBy to be an explicit nil

### UnsetDescribedBy
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetDescribedBy()`

UnsetDescribedBy ensures that no value is present for DescribedBy, not even an explicit nil
### GetDescription

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedTransactionType

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetExtendedTransactionType() string`

GetExtendedTransactionType returns the ExtendedTransactionType field if non-nil, zero value otherwise.

### GetExtendedTransactionTypeOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetExtendedTransactionTypeOk() (*string, bool)`

GetExtendedTransactionTypeOk returns a tuple with the ExtendedTransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedTransactionType

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetExtendedTransactionType(v string)`

SetExtendedTransactionType sets ExtendedTransactionType field to given value.

### HasExtendedTransactionType

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasExtendedTransactionType() bool`

HasExtendedTransactionType returns a boolean if a field has been set.

### SetExtendedTransactionTypeNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetExtendedTransactionTypeNil(b bool)`

 SetExtendedTransactionTypeNil sets the value for ExtendedTransactionType to be an explicit nil

### UnsetExtendedTransactionType
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetExtendedTransactionType()`

UnsetExtendedTransactionType ensures that no value is present for ExtendedTransactionType, not even an explicit nil
### GetId

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsBillPay

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsBillPay() bool`

GetIsBillPay returns the IsBillPay field if non-nil, zero value otherwise.

### GetIsBillPayOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsBillPayOk() (*bool, bool)`

GetIsBillPayOk returns a tuple with the IsBillPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillPay

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsBillPay(v bool)`

SetIsBillPay sets IsBillPay field to given value.

### HasIsBillPay

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasIsBillPay() bool`

HasIsBillPay returns a boolean if a field has been set.

### SetIsBillPayNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsBillPayNil(b bool)`

 SetIsBillPayNil sets the value for IsBillPay to be an explicit nil

### UnsetIsBillPay
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetIsBillPay()`

UnsetIsBillPay ensures that no value is present for IsBillPay, not even an explicit nil
### GetIsDirectDeposit

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsDirectDeposit() bool`

GetIsDirectDeposit returns the IsDirectDeposit field if non-nil, zero value otherwise.

### GetIsDirectDepositOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsDirectDepositOk() (*bool, bool)`

GetIsDirectDepositOk returns a tuple with the IsDirectDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectDeposit

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsDirectDeposit(v bool)`

SetIsDirectDeposit sets IsDirectDeposit field to given value.

### HasIsDirectDeposit

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasIsDirectDeposit() bool`

HasIsDirectDeposit returns a boolean if a field has been set.

### SetIsDirectDepositNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsDirectDepositNil(b bool)`

 SetIsDirectDepositNil sets the value for IsDirectDeposit to be an explicit nil

### UnsetIsDirectDeposit
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetIsDirectDeposit()`

UnsetIsDirectDeposit ensures that no value is present for IsDirectDeposit, not even an explicit nil
### GetIsExpense

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsExpense() bool`

GetIsExpense returns the IsExpense field if non-nil, zero value otherwise.

### GetIsExpenseOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsExpenseOk() (*bool, bool)`

GetIsExpenseOk returns a tuple with the IsExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpense

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsExpense(v bool)`

SetIsExpense sets IsExpense field to given value.

### HasIsExpense

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasIsExpense() bool`

HasIsExpense returns a boolean if a field has been set.

### SetIsExpenseNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsExpenseNil(b bool)`

 SetIsExpenseNil sets the value for IsExpense to be an explicit nil

### UnsetIsExpense
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetIsExpense()`

UnsetIsExpense ensures that no value is present for IsExpense, not even an explicit nil
### GetIsFee

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsFee() bool`

GetIsFee returns the IsFee field if non-nil, zero value otherwise.

### GetIsFeeOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsFeeOk() (*bool, bool)`

GetIsFeeOk returns a tuple with the IsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFee

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsFee(v bool)`

SetIsFee sets IsFee field to given value.

### HasIsFee

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasIsFee() bool`

HasIsFee returns a boolean if a field has been set.

### SetIsFeeNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsFeeNil(b bool)`

 SetIsFeeNil sets the value for IsFee to be an explicit nil

### UnsetIsFee
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetIsFee()`

UnsetIsFee ensures that no value is present for IsFee, not even an explicit nil
### GetIsIncome

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsIncome() bool`

GetIsIncome returns the IsIncome field if non-nil, zero value otherwise.

### GetIsIncomeOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsIncomeOk() (*bool, bool)`

GetIsIncomeOk returns a tuple with the IsIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncome

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsIncome(v bool)`

SetIsIncome sets IsIncome field to given value.

### HasIsIncome

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasIsIncome() bool`

HasIsIncome returns a boolean if a field has been set.

### SetIsIncomeNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsIncomeNil(b bool)`

 SetIsIncomeNil sets the value for IsIncome to be an explicit nil

### UnsetIsIncome
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetIsIncome()`

UnsetIsIncome ensures that no value is present for IsIncome, not even an explicit nil
### GetIsInternational

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsInternational() bool`

GetIsInternational returns the IsInternational field if non-nil, zero value otherwise.

### GetIsInternationalOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsInternationalOk() (*bool, bool)`

GetIsInternationalOk returns a tuple with the IsInternational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternational

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsInternational(v bool)`

SetIsInternational sets IsInternational field to given value.

### HasIsInternational

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasIsInternational() bool`

HasIsInternational returns a boolean if a field has been set.

### SetIsInternationalNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsInternationalNil(b bool)`

 SetIsInternationalNil sets the value for IsInternational to be an explicit nil

### UnsetIsInternational
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetIsInternational()`

UnsetIsInternational ensures that no value is present for IsInternational, not even an explicit nil
### GetIsOverdraftFee

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsOverdraftFee() bool`

GetIsOverdraftFee returns the IsOverdraftFee field if non-nil, zero value otherwise.

### GetIsOverdraftFeeOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsOverdraftFeeOk() (*bool, bool)`

GetIsOverdraftFeeOk returns a tuple with the IsOverdraftFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdraftFee

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsOverdraftFee(v bool)`

SetIsOverdraftFee sets IsOverdraftFee field to given value.

### HasIsOverdraftFee

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasIsOverdraftFee() bool`

HasIsOverdraftFee returns a boolean if a field has been set.

### SetIsOverdraftFeeNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsOverdraftFeeNil(b bool)`

 SetIsOverdraftFeeNil sets the value for IsOverdraftFee to be an explicit nil

### UnsetIsOverdraftFee
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetIsOverdraftFee()`

UnsetIsOverdraftFee ensures that no value is present for IsOverdraftFee, not even an explicit nil
### GetIsPayrollAdvance

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsPayrollAdvance() bool`

GetIsPayrollAdvance returns the IsPayrollAdvance field if non-nil, zero value otherwise.

### GetIsPayrollAdvanceOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsPayrollAdvanceOk() (*bool, bool)`

GetIsPayrollAdvanceOk returns a tuple with the IsPayrollAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPayrollAdvance

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsPayrollAdvance(v bool)`

SetIsPayrollAdvance sets IsPayrollAdvance field to given value.

### HasIsPayrollAdvance

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasIsPayrollAdvance() bool`

HasIsPayrollAdvance returns a boolean if a field has been set.

### SetIsPayrollAdvanceNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsPayrollAdvanceNil(b bool)`

 SetIsPayrollAdvanceNil sets the value for IsPayrollAdvance to be an explicit nil

### UnsetIsPayrollAdvance
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetIsPayrollAdvance()`

UnsetIsPayrollAdvance ensures that no value is present for IsPayrollAdvance, not even an explicit nil
### GetIsSubscription

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsSubscription() bool`

GetIsSubscription returns the IsSubscription field if non-nil, zero value otherwise.

### GetIsSubscriptionOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetIsSubscriptionOk() (*bool, bool)`

GetIsSubscriptionOk returns a tuple with the IsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscription

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsSubscription(v bool)`

SetIsSubscription sets IsSubscription field to given value.

### HasIsSubscription

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasIsSubscription() bool`

HasIsSubscription returns a boolean if a field has been set.

### SetIsSubscriptionNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetIsSubscriptionNil(b bool)`

 SetIsSubscriptionNil sets the value for IsSubscription to be an explicit nil

### UnsetIsSubscription
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetIsSubscription()`

UnsetIsSubscription ensures that no value is present for IsSubscription, not even an explicit nil
### GetMemo

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetMerchantCategoryCode

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetMerchantCategoryCode() string`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetMerchantCategoryCodeOk() (*string, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetMerchantCategoryCode(v string)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetMerchantGuid

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetMerchantGuid() string`

GetMerchantGuid returns the MerchantGuid field if non-nil, zero value otherwise.

### GetMerchantGuidOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetMerchantGuidOk() (*string, bool)`

GetMerchantGuidOk returns a tuple with the MerchantGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGuid

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetMerchantGuid(v string)`

SetMerchantGuid sets MerchantGuid field to given value.

### HasMerchantGuid

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasMerchantGuid() bool`

HasMerchantGuid returns a boolean if a field has been set.

### SetMerchantGuidNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetMerchantGuidNil(b bool)`

 SetMerchantGuidNil sets the value for MerchantGuid to be an explicit nil

### UnsetMerchantGuid
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetMerchantGuid()`

UnsetMerchantGuid ensures that no value is present for MerchantGuid, not even an explicit nil
### GetMerchantLocationGuid

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetMerchantLocationGuid() string`

GetMerchantLocationGuid returns the MerchantLocationGuid field if non-nil, zero value otherwise.

### GetMerchantLocationGuidOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetMerchantLocationGuidOk() (*string, bool)`

GetMerchantLocationGuidOk returns a tuple with the MerchantLocationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantLocationGuid

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetMerchantLocationGuid(v string)`

SetMerchantLocationGuid sets MerchantLocationGuid field to given value.

### HasMerchantLocationGuid

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasMerchantLocationGuid() bool`

HasMerchantLocationGuid returns a boolean if a field has been set.

### SetMerchantLocationGuidNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetMerchantLocationGuidNil(b bool)`

 SetMerchantLocationGuidNil sets the value for MerchantLocationGuid to be an explicit nil

### UnsetMerchantLocationGuid
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetMerchantLocationGuid()`

UnsetMerchantLocationGuid ensures that no value is present for MerchantLocationGuid, not even an explicit nil
### GetOriginalDescription

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetOriginalDescription() string`

GetOriginalDescription returns the OriginalDescription field if non-nil, zero value otherwise.

### GetOriginalDescriptionOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetOriginalDescriptionOk() (*string, bool)`

GetOriginalDescriptionOk returns a tuple with the OriginalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDescription

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetOriginalDescription(v string)`

SetOriginalDescription sets OriginalDescription field to given value.

### HasOriginalDescription

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasOriginalDescription() bool`

HasOriginalDescription returns a boolean if a field has been set.

### SetOriginalDescriptionNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetOriginalDescriptionNil(b bool)`

 SetOriginalDescriptionNil sets the value for OriginalDescription to be an explicit nil

### UnsetOriginalDescription
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetOriginalDescription()`

UnsetOriginalDescription ensures that no value is present for OriginalDescription, not even an explicit nil
### GetType

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayAuthorizationEnhancedTransactionRaw) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayAuthorizationEnhancedTransactionRaw) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GatewayAuthorizationEnhancedTransactionRaw) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GatewayAuthorizationEnhancedTransactionRaw) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


