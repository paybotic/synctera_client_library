# SpendControlCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionCase** | **bool** | If set, create a case for transactions that do not conform to the spend control | 
**ActionDecline** | **bool** | If set, decline transactions that do not conform to the spend control | 
**AmountLimit** | Pointer to **int64** | Monetary limit for the spend control in the smallest currency unit (eg cents). At least one of amount_limit and transaction_count_limit must be specified. | [optional] 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the spend control was created | [optional] [readonly] 
**Direction** | [**SpendControlDirection**](SpendControlDirection.md) |  | 
**GeoRestrictions** | Pointer to [**SpendControlGeoRestrictions**](SpendControlGeoRestrictions.md) |  | [optional] 
**Id** | Pointer to **string** | Spend Control ID | [optional] [readonly] 
**IsActive** | **bool** | Indicates if spend control is active | 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the spend control was last modified | [optional] [readonly] 
**ManagedBy** | Pointer to [**ManagedByTypes**](ManagedByTypes.md) |  | [optional] 
**MerchantCategoryCodes** | Pointer to **[]string** | merchant category codes for spend control | [optional] 
**Name** | **string** | Name assigned to spend control | 
**PaymentSubTypes** | Pointer to **[]string** | A list of payment sub-types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all sub-types. | [optional] 
**PaymentTypes** | Pointer to [**[]PaymentType**](PaymentType.md) | A list of payment types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all types of payments. | [optional] 
**TimeRange** | [**SpendControlTimeRange**](SpendControlTimeRange.md) |  | 
**TransactionCountLimit** | Pointer to **int64** | Number of transactions allowed by the spend control, regardless of transaction amounts. At least one of amount_limit and transaction_count_limit must be specified. | [optional] 

## Methods

### NewSpendControlCreationRequest

`func NewSpendControlCreationRequest(actionCase bool, actionDecline bool, direction SpendControlDirection, isActive bool, name string, timeRange SpendControlTimeRange, ) *SpendControlCreationRequest`

NewSpendControlCreationRequest instantiates a new SpendControlCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlCreationRequestWithDefaults

`func NewSpendControlCreationRequestWithDefaults() *SpendControlCreationRequest`

NewSpendControlCreationRequestWithDefaults instantiates a new SpendControlCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionCase

`func (o *SpendControlCreationRequest) GetActionCase() bool`

GetActionCase returns the ActionCase field if non-nil, zero value otherwise.

### GetActionCaseOk

`func (o *SpendControlCreationRequest) GetActionCaseOk() (*bool, bool)`

GetActionCaseOk returns a tuple with the ActionCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCase

`func (o *SpendControlCreationRequest) SetActionCase(v bool)`

SetActionCase sets ActionCase field to given value.


### GetActionDecline

`func (o *SpendControlCreationRequest) GetActionDecline() bool`

GetActionDecline returns the ActionDecline field if non-nil, zero value otherwise.

### GetActionDeclineOk

`func (o *SpendControlCreationRequest) GetActionDeclineOk() (*bool, bool)`

GetActionDeclineOk returns a tuple with the ActionDecline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDecline

`func (o *SpendControlCreationRequest) SetActionDecline(v bool)`

SetActionDecline sets ActionDecline field to given value.


### GetAmountLimit

`func (o *SpendControlCreationRequest) GetAmountLimit() int64`

GetAmountLimit returns the AmountLimit field if non-nil, zero value otherwise.

### GetAmountLimitOk

`func (o *SpendControlCreationRequest) GetAmountLimitOk() (*int64, bool)`

GetAmountLimitOk returns a tuple with the AmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimit

`func (o *SpendControlCreationRequest) SetAmountLimit(v int64)`

SetAmountLimit sets AmountLimit field to given value.

### HasAmountLimit

`func (o *SpendControlCreationRequest) HasAmountLimit() bool`

HasAmountLimit returns a boolean if a field has been set.

### GetCreationTime

`func (o *SpendControlCreationRequest) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *SpendControlCreationRequest) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *SpendControlCreationRequest) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *SpendControlCreationRequest) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDirection

`func (o *SpendControlCreationRequest) GetDirection() SpendControlDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SpendControlCreationRequest) GetDirectionOk() (*SpendControlDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SpendControlCreationRequest) SetDirection(v SpendControlDirection)`

SetDirection sets Direction field to given value.


### GetGeoRestrictions

`func (o *SpendControlCreationRequest) GetGeoRestrictions() SpendControlGeoRestrictions`

GetGeoRestrictions returns the GeoRestrictions field if non-nil, zero value otherwise.

### GetGeoRestrictionsOk

`func (o *SpendControlCreationRequest) GetGeoRestrictionsOk() (*SpendControlGeoRestrictions, bool)`

GetGeoRestrictionsOk returns a tuple with the GeoRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoRestrictions

`func (o *SpendControlCreationRequest) SetGeoRestrictions(v SpendControlGeoRestrictions)`

SetGeoRestrictions sets GeoRestrictions field to given value.

### HasGeoRestrictions

`func (o *SpendControlCreationRequest) HasGeoRestrictions() bool`

HasGeoRestrictions returns a boolean if a field has been set.

### GetId

`func (o *SpendControlCreationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpendControlCreationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpendControlCreationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpendControlCreationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *SpendControlCreationRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SpendControlCreationRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SpendControlCreationRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastModifiedTime

`func (o *SpendControlCreationRequest) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *SpendControlCreationRequest) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *SpendControlCreationRequest) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *SpendControlCreationRequest) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetManagedBy

`func (o *SpendControlCreationRequest) GetManagedBy() ManagedByTypes`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *SpendControlCreationRequest) GetManagedByOk() (*ManagedByTypes, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *SpendControlCreationRequest) SetManagedBy(v ManagedByTypes)`

SetManagedBy sets ManagedBy field to given value.

### HasManagedBy

`func (o *SpendControlCreationRequest) HasManagedBy() bool`

HasManagedBy returns a boolean if a field has been set.

### GetMerchantCategoryCodes

`func (o *SpendControlCreationRequest) GetMerchantCategoryCodes() []string`

GetMerchantCategoryCodes returns the MerchantCategoryCodes field if non-nil, zero value otherwise.

### GetMerchantCategoryCodesOk

`func (o *SpendControlCreationRequest) GetMerchantCategoryCodesOk() (*[]string, bool)`

GetMerchantCategoryCodesOk returns a tuple with the MerchantCategoryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCodes

`func (o *SpendControlCreationRequest) SetMerchantCategoryCodes(v []string)`

SetMerchantCategoryCodes sets MerchantCategoryCodes field to given value.

### HasMerchantCategoryCodes

`func (o *SpendControlCreationRequest) HasMerchantCategoryCodes() bool`

HasMerchantCategoryCodes returns a boolean if a field has been set.

### GetName

`func (o *SpendControlCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpendControlCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpendControlCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPaymentSubTypes

`func (o *SpendControlCreationRequest) GetPaymentSubTypes() []string`

GetPaymentSubTypes returns the PaymentSubTypes field if non-nil, zero value otherwise.

### GetPaymentSubTypesOk

`func (o *SpendControlCreationRequest) GetPaymentSubTypesOk() (*[]string, bool)`

GetPaymentSubTypesOk returns a tuple with the PaymentSubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSubTypes

`func (o *SpendControlCreationRequest) SetPaymentSubTypes(v []string)`

SetPaymentSubTypes sets PaymentSubTypes field to given value.

### HasPaymentSubTypes

`func (o *SpendControlCreationRequest) HasPaymentSubTypes() bool`

HasPaymentSubTypes returns a boolean if a field has been set.

### GetPaymentTypes

`func (o *SpendControlCreationRequest) GetPaymentTypes() []PaymentType`

GetPaymentTypes returns the PaymentTypes field if non-nil, zero value otherwise.

### GetPaymentTypesOk

`func (o *SpendControlCreationRequest) GetPaymentTypesOk() (*[]PaymentType, bool)`

GetPaymentTypesOk returns a tuple with the PaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypes

`func (o *SpendControlCreationRequest) SetPaymentTypes(v []PaymentType)`

SetPaymentTypes sets PaymentTypes field to given value.

### HasPaymentTypes

`func (o *SpendControlCreationRequest) HasPaymentTypes() bool`

HasPaymentTypes returns a boolean if a field has been set.

### GetTimeRange

`func (o *SpendControlCreationRequest) GetTimeRange() SpendControlTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SpendControlCreationRequest) GetTimeRangeOk() (*SpendControlTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SpendControlCreationRequest) SetTimeRange(v SpendControlTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetTransactionCountLimit

`func (o *SpendControlCreationRequest) GetTransactionCountLimit() int64`

GetTransactionCountLimit returns the TransactionCountLimit field if non-nil, zero value otherwise.

### GetTransactionCountLimitOk

`func (o *SpendControlCreationRequest) GetTransactionCountLimitOk() (*int64, bool)`

GetTransactionCountLimitOk returns a tuple with the TransactionCountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCountLimit

`func (o *SpendControlCreationRequest) SetTransactionCountLimit(v int64)`

SetTransactionCountLimit sets TransactionCountLimit field to given value.

### HasTransactionCountLimit

`func (o *SpendControlCreationRequest) HasTransactionCountLimit() bool`

HasTransactionCountLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


