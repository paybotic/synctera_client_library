# SpendControlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionCase** | **bool** | If set, create a case for transactions that do not conform to the spend control | 
**ActionDecline** | **bool** | If set, decline transactions that do not conform to the spend control | 
**AmountLimit** | Pointer to **int64** | Monetary limit for the spend control in the smallest currency unit (eg cents). At least one of amount_limit and transaction_count_limit must be specified. | [optional] 
**CreationTime** | **time.Time** | The timestamp representing when the spend control was created | [readonly] 
**Direction** | [**SpendControlDirection**](SpendControlDirection.md) |  | 
**GeoRestrictions** | Pointer to [**SpendControlGeoRestrictions**](SpendControlGeoRestrictions.md) |  | [optional] 
**Id** | **string** | Spend Control ID | [readonly] 
**IsActive** | **bool** | Indicates if spend control is active | 
**LastModifiedTime** | **time.Time** | The timestamp representing when the spend control was last modified | [readonly] 
**ManagedBy** | Pointer to [**ManagedByTypes**](ManagedByTypes.md) |  | [optional] 
**MerchantCategoryCodes** | Pointer to **[]string** | merchant category codes for spend control | [optional] 
**Name** | **string** | Name assigned to spend control | 
**PaymentSubTypes** | Pointer to **[]string** | A list of payment sub-types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all sub-types. | [optional] 
**PaymentTypes** | Pointer to [**[]PaymentType**](PaymentType.md) | A list of payment types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all types of payments. | [optional] 
**TimeRange** | [**SpendControlTimeRange**](SpendControlTimeRange.md) |  | 
**TransactionCountLimit** | Pointer to **int64** | Number of transactions allowed by the spend control, regardless of transaction amounts. At least one of amount_limit and transaction_count_limit must be specified. | [optional] 
**NumberOfRelatedAccounts** | **int32** | A count of how many accounts are using this spend control | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewSpendControlResponse

`func NewSpendControlResponse(actionCase bool, actionDecline bool, creationTime time.Time, direction SpendControlDirection, id string, isActive bool, lastModifiedTime time.Time, name string, timeRange SpendControlTimeRange, numberOfRelatedAccounts int32, tenant string, ) *SpendControlResponse`

NewSpendControlResponse instantiates a new SpendControlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlResponseWithDefaults

`func NewSpendControlResponseWithDefaults() *SpendControlResponse`

NewSpendControlResponseWithDefaults instantiates a new SpendControlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionCase

`func (o *SpendControlResponse) GetActionCase() bool`

GetActionCase returns the ActionCase field if non-nil, zero value otherwise.

### GetActionCaseOk

`func (o *SpendControlResponse) GetActionCaseOk() (*bool, bool)`

GetActionCaseOk returns a tuple with the ActionCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCase

`func (o *SpendControlResponse) SetActionCase(v bool)`

SetActionCase sets ActionCase field to given value.


### GetActionDecline

`func (o *SpendControlResponse) GetActionDecline() bool`

GetActionDecline returns the ActionDecline field if non-nil, zero value otherwise.

### GetActionDeclineOk

`func (o *SpendControlResponse) GetActionDeclineOk() (*bool, bool)`

GetActionDeclineOk returns a tuple with the ActionDecline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDecline

`func (o *SpendControlResponse) SetActionDecline(v bool)`

SetActionDecline sets ActionDecline field to given value.


### GetAmountLimit

`func (o *SpendControlResponse) GetAmountLimit() int64`

GetAmountLimit returns the AmountLimit field if non-nil, zero value otherwise.

### GetAmountLimitOk

`func (o *SpendControlResponse) GetAmountLimitOk() (*int64, bool)`

GetAmountLimitOk returns a tuple with the AmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimit

`func (o *SpendControlResponse) SetAmountLimit(v int64)`

SetAmountLimit sets AmountLimit field to given value.

### HasAmountLimit

`func (o *SpendControlResponse) HasAmountLimit() bool`

HasAmountLimit returns a boolean if a field has been set.

### GetCreationTime

`func (o *SpendControlResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *SpendControlResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *SpendControlResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetDirection

`func (o *SpendControlResponse) GetDirection() SpendControlDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SpendControlResponse) GetDirectionOk() (*SpendControlDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SpendControlResponse) SetDirection(v SpendControlDirection)`

SetDirection sets Direction field to given value.


### GetGeoRestrictions

`func (o *SpendControlResponse) GetGeoRestrictions() SpendControlGeoRestrictions`

GetGeoRestrictions returns the GeoRestrictions field if non-nil, zero value otherwise.

### GetGeoRestrictionsOk

`func (o *SpendControlResponse) GetGeoRestrictionsOk() (*SpendControlGeoRestrictions, bool)`

GetGeoRestrictionsOk returns a tuple with the GeoRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoRestrictions

`func (o *SpendControlResponse) SetGeoRestrictions(v SpendControlGeoRestrictions)`

SetGeoRestrictions sets GeoRestrictions field to given value.

### HasGeoRestrictions

`func (o *SpendControlResponse) HasGeoRestrictions() bool`

HasGeoRestrictions returns a boolean if a field has been set.

### GetId

`func (o *SpendControlResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpendControlResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpendControlResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *SpendControlResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SpendControlResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SpendControlResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastModifiedTime

`func (o *SpendControlResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *SpendControlResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *SpendControlResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetManagedBy

`func (o *SpendControlResponse) GetManagedBy() ManagedByTypes`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *SpendControlResponse) GetManagedByOk() (*ManagedByTypes, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *SpendControlResponse) SetManagedBy(v ManagedByTypes)`

SetManagedBy sets ManagedBy field to given value.

### HasManagedBy

`func (o *SpendControlResponse) HasManagedBy() bool`

HasManagedBy returns a boolean if a field has been set.

### GetMerchantCategoryCodes

`func (o *SpendControlResponse) GetMerchantCategoryCodes() []string`

GetMerchantCategoryCodes returns the MerchantCategoryCodes field if non-nil, zero value otherwise.

### GetMerchantCategoryCodesOk

`func (o *SpendControlResponse) GetMerchantCategoryCodesOk() (*[]string, bool)`

GetMerchantCategoryCodesOk returns a tuple with the MerchantCategoryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCodes

`func (o *SpendControlResponse) SetMerchantCategoryCodes(v []string)`

SetMerchantCategoryCodes sets MerchantCategoryCodes field to given value.

### HasMerchantCategoryCodes

`func (o *SpendControlResponse) HasMerchantCategoryCodes() bool`

HasMerchantCategoryCodes returns a boolean if a field has been set.

### GetName

`func (o *SpendControlResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpendControlResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpendControlResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPaymentSubTypes

`func (o *SpendControlResponse) GetPaymentSubTypes() []string`

GetPaymentSubTypes returns the PaymentSubTypes field if non-nil, zero value otherwise.

### GetPaymentSubTypesOk

`func (o *SpendControlResponse) GetPaymentSubTypesOk() (*[]string, bool)`

GetPaymentSubTypesOk returns a tuple with the PaymentSubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSubTypes

`func (o *SpendControlResponse) SetPaymentSubTypes(v []string)`

SetPaymentSubTypes sets PaymentSubTypes field to given value.

### HasPaymentSubTypes

`func (o *SpendControlResponse) HasPaymentSubTypes() bool`

HasPaymentSubTypes returns a boolean if a field has been set.

### GetPaymentTypes

`func (o *SpendControlResponse) GetPaymentTypes() []PaymentType`

GetPaymentTypes returns the PaymentTypes field if non-nil, zero value otherwise.

### GetPaymentTypesOk

`func (o *SpendControlResponse) GetPaymentTypesOk() (*[]PaymentType, bool)`

GetPaymentTypesOk returns a tuple with the PaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypes

`func (o *SpendControlResponse) SetPaymentTypes(v []PaymentType)`

SetPaymentTypes sets PaymentTypes field to given value.

### HasPaymentTypes

`func (o *SpendControlResponse) HasPaymentTypes() bool`

HasPaymentTypes returns a boolean if a field has been set.

### GetTimeRange

`func (o *SpendControlResponse) GetTimeRange() SpendControlTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SpendControlResponse) GetTimeRangeOk() (*SpendControlTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SpendControlResponse) SetTimeRange(v SpendControlTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetTransactionCountLimit

`func (o *SpendControlResponse) GetTransactionCountLimit() int64`

GetTransactionCountLimit returns the TransactionCountLimit field if non-nil, zero value otherwise.

### GetTransactionCountLimitOk

`func (o *SpendControlResponse) GetTransactionCountLimitOk() (*int64, bool)`

GetTransactionCountLimitOk returns a tuple with the TransactionCountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCountLimit

`func (o *SpendControlResponse) SetTransactionCountLimit(v int64)`

SetTransactionCountLimit sets TransactionCountLimit field to given value.

### HasTransactionCountLimit

`func (o *SpendControlResponse) HasTransactionCountLimit() bool`

HasTransactionCountLimit returns a boolean if a field has been set.

### GetNumberOfRelatedAccounts

`func (o *SpendControlResponse) GetNumberOfRelatedAccounts() int32`

GetNumberOfRelatedAccounts returns the NumberOfRelatedAccounts field if non-nil, zero value otherwise.

### GetNumberOfRelatedAccountsOk

`func (o *SpendControlResponse) GetNumberOfRelatedAccountsOk() (*int32, bool)`

GetNumberOfRelatedAccountsOk returns a tuple with the NumberOfRelatedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRelatedAccounts

`func (o *SpendControlResponse) SetNumberOfRelatedAccounts(v int32)`

SetNumberOfRelatedAccounts sets NumberOfRelatedAccounts field to given value.


### GetTenant

`func (o *SpendControlResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *SpendControlResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *SpendControlResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


