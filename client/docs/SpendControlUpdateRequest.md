# SpendControlUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionCase** | Pointer to **bool** | If set, create a case for transactions that do not conform to the spend control | [optional] 
**ActionDecline** | Pointer to **bool** | If set, decline transactions that do not conform to the spend control | [optional] 
**AmountLimit** | Pointer to **int64** | Monetary limit for the spend control in the smallest currency unit (eg cents) | [optional] 
**Direction** | Pointer to [**SpendControlDirection**](SpendControlDirection.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates if spend control is active | [optional] 
**MerchantCategoryCodes** | Pointer to **[]string** | merchant category codes for spend control | [optional] 
**Name** | Pointer to **string** | Name assigned to spend control | [optional] 
**PaymentSubTypes** | Pointer to [**[]PaymentSubType**](PaymentSubType.md) | A list of payment sub-types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all sub-types. | [optional] 
**PaymentTypes** | Pointer to [**[]PaymentType**](PaymentType.md) | A list of payment types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all types of payments. | [optional] 
**TimeRange** | Pointer to [**SpendControlTimeRange**](SpendControlTimeRange.md) |  | [optional] 

## Methods

### NewSpendControlUpdateRequest

`func NewSpendControlUpdateRequest() *SpendControlUpdateRequest`

NewSpendControlUpdateRequest instantiates a new SpendControlUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlUpdateRequestWithDefaults

`func NewSpendControlUpdateRequestWithDefaults() *SpendControlUpdateRequest`

NewSpendControlUpdateRequestWithDefaults instantiates a new SpendControlUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionCase

`func (o *SpendControlUpdateRequest) GetActionCase() bool`

GetActionCase returns the ActionCase field if non-nil, zero value otherwise.

### GetActionCaseOk

`func (o *SpendControlUpdateRequest) GetActionCaseOk() (*bool, bool)`

GetActionCaseOk returns a tuple with the ActionCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCase

`func (o *SpendControlUpdateRequest) SetActionCase(v bool)`

SetActionCase sets ActionCase field to given value.

### HasActionCase

`func (o *SpendControlUpdateRequest) HasActionCase() bool`

HasActionCase returns a boolean if a field has been set.

### GetActionDecline

`func (o *SpendControlUpdateRequest) GetActionDecline() bool`

GetActionDecline returns the ActionDecline field if non-nil, zero value otherwise.

### GetActionDeclineOk

`func (o *SpendControlUpdateRequest) GetActionDeclineOk() (*bool, bool)`

GetActionDeclineOk returns a tuple with the ActionDecline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDecline

`func (o *SpendControlUpdateRequest) SetActionDecline(v bool)`

SetActionDecline sets ActionDecline field to given value.

### HasActionDecline

`func (o *SpendControlUpdateRequest) HasActionDecline() bool`

HasActionDecline returns a boolean if a field has been set.

### GetAmountLimit

`func (o *SpendControlUpdateRequest) GetAmountLimit() int64`

GetAmountLimit returns the AmountLimit field if non-nil, zero value otherwise.

### GetAmountLimitOk

`func (o *SpendControlUpdateRequest) GetAmountLimitOk() (*int64, bool)`

GetAmountLimitOk returns a tuple with the AmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimit

`func (o *SpendControlUpdateRequest) SetAmountLimit(v int64)`

SetAmountLimit sets AmountLimit field to given value.

### HasAmountLimit

`func (o *SpendControlUpdateRequest) HasAmountLimit() bool`

HasAmountLimit returns a boolean if a field has been set.

### GetDirection

`func (o *SpendControlUpdateRequest) GetDirection() SpendControlDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SpendControlUpdateRequest) GetDirectionOk() (*SpendControlDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SpendControlUpdateRequest) SetDirection(v SpendControlDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SpendControlUpdateRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetIsActive

`func (o *SpendControlUpdateRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SpendControlUpdateRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SpendControlUpdateRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SpendControlUpdateRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetMerchantCategoryCodes

`func (o *SpendControlUpdateRequest) GetMerchantCategoryCodes() []string`

GetMerchantCategoryCodes returns the MerchantCategoryCodes field if non-nil, zero value otherwise.

### GetMerchantCategoryCodesOk

`func (o *SpendControlUpdateRequest) GetMerchantCategoryCodesOk() (*[]string, bool)`

GetMerchantCategoryCodesOk returns a tuple with the MerchantCategoryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCodes

`func (o *SpendControlUpdateRequest) SetMerchantCategoryCodes(v []string)`

SetMerchantCategoryCodes sets MerchantCategoryCodes field to given value.

### HasMerchantCategoryCodes

`func (o *SpendControlUpdateRequest) HasMerchantCategoryCodes() bool`

HasMerchantCategoryCodes returns a boolean if a field has been set.

### GetName

`func (o *SpendControlUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpendControlUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpendControlUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpendControlUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentSubTypes

`func (o *SpendControlUpdateRequest) GetPaymentSubTypes() []PaymentSubType`

GetPaymentSubTypes returns the PaymentSubTypes field if non-nil, zero value otherwise.

### GetPaymentSubTypesOk

`func (o *SpendControlUpdateRequest) GetPaymentSubTypesOk() (*[]PaymentSubType, bool)`

GetPaymentSubTypesOk returns a tuple with the PaymentSubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSubTypes

`func (o *SpendControlUpdateRequest) SetPaymentSubTypes(v []PaymentSubType)`

SetPaymentSubTypes sets PaymentSubTypes field to given value.

### HasPaymentSubTypes

`func (o *SpendControlUpdateRequest) HasPaymentSubTypes() bool`

HasPaymentSubTypes returns a boolean if a field has been set.

### GetPaymentTypes

`func (o *SpendControlUpdateRequest) GetPaymentTypes() []PaymentType`

GetPaymentTypes returns the PaymentTypes field if non-nil, zero value otherwise.

### GetPaymentTypesOk

`func (o *SpendControlUpdateRequest) GetPaymentTypesOk() (*[]PaymentType, bool)`

GetPaymentTypesOk returns a tuple with the PaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypes

`func (o *SpendControlUpdateRequest) SetPaymentTypes(v []PaymentType)`

SetPaymentTypes sets PaymentTypes field to given value.

### HasPaymentTypes

`func (o *SpendControlUpdateRequest) HasPaymentTypes() bool`

HasPaymentTypes returns a boolean if a field has been set.

### GetTimeRange

`func (o *SpendControlUpdateRequest) GetTimeRange() SpendControlTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SpendControlUpdateRequest) GetTimeRangeOk() (*SpendControlTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SpendControlUpdateRequest) SetTimeRange(v SpendControlTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *SpendControlUpdateRequest) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


