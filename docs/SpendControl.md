# SpendControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionCase** | **bool** | If set, create a case for transactions that do not conform to the spend control | 
**ActionDecline** | **bool** | If set, decline transactions that do not conform to the spend control | 
**AmountLimit** | **int64** | Monetary limit for the spend control in the smallest currency unit (eg cents) | 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the spend control was created | [optional] [readonly] 
**Direction** | Pointer to [**SpendControlDirection**](SpendControlDirection.md) |  | [optional] 
**Id** | Pointer to **string** | Spend Control ID | [optional] [readonly] 
**IsActive** | **bool** | Indicates if spend control is active | 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the spend control was last modified | [optional] [readonly] 
**Name** | **string** | Name assigned to spend control | 
**PaymentTypes** | Pointer to [**[]PaymentType**](PaymentType.md) | A list of payment types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all types of payments. | [optional] 
**TimeRange** | [**SpendControlTimeRange**](SpendControlTimeRange.md) |  | 

## Methods

### NewSpendControl

`func NewSpendControl(actionCase bool, actionDecline bool, amountLimit int64, isActive bool, name string, timeRange SpendControlTimeRange, ) *SpendControl`

NewSpendControl instantiates a new SpendControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlWithDefaults

`func NewSpendControlWithDefaults() *SpendControl`

NewSpendControlWithDefaults instantiates a new SpendControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionCase

`func (o *SpendControl) GetActionCase() bool`

GetActionCase returns the ActionCase field if non-nil, zero value otherwise.

### GetActionCaseOk

`func (o *SpendControl) GetActionCaseOk() (*bool, bool)`

GetActionCaseOk returns a tuple with the ActionCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCase

`func (o *SpendControl) SetActionCase(v bool)`

SetActionCase sets ActionCase field to given value.


### GetActionDecline

`func (o *SpendControl) GetActionDecline() bool`

GetActionDecline returns the ActionDecline field if non-nil, zero value otherwise.

### GetActionDeclineOk

`func (o *SpendControl) GetActionDeclineOk() (*bool, bool)`

GetActionDeclineOk returns a tuple with the ActionDecline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDecline

`func (o *SpendControl) SetActionDecline(v bool)`

SetActionDecline sets ActionDecline field to given value.


### GetAmountLimit

`func (o *SpendControl) GetAmountLimit() int64`

GetAmountLimit returns the AmountLimit field if non-nil, zero value otherwise.

### GetAmountLimitOk

`func (o *SpendControl) GetAmountLimitOk() (*int64, bool)`

GetAmountLimitOk returns a tuple with the AmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimit

`func (o *SpendControl) SetAmountLimit(v int64)`

SetAmountLimit sets AmountLimit field to given value.


### GetCreationTime

`func (o *SpendControl) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *SpendControl) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *SpendControl) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *SpendControl) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDirection

`func (o *SpendControl) GetDirection() SpendControlDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SpendControl) GetDirectionOk() (*SpendControlDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SpendControl) SetDirection(v SpendControlDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SpendControl) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *SpendControl) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpendControl) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpendControl) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpendControl) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *SpendControl) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SpendControl) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SpendControl) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastModifiedTime

`func (o *SpendControl) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *SpendControl) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *SpendControl) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *SpendControl) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *SpendControl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpendControl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpendControl) SetName(v string)`

SetName sets Name field to given value.


### GetPaymentTypes

`func (o *SpendControl) GetPaymentTypes() []PaymentType`

GetPaymentTypes returns the PaymentTypes field if non-nil, zero value otherwise.

### GetPaymentTypesOk

`func (o *SpendControl) GetPaymentTypesOk() (*[]PaymentType, bool)`

GetPaymentTypesOk returns a tuple with the PaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypes

`func (o *SpendControl) SetPaymentTypes(v []PaymentType)`

SetPaymentTypes sets PaymentTypes field to given value.

### HasPaymentTypes

`func (o *SpendControl) HasPaymentTypes() bool`

HasPaymentTypes returns a boolean if a field has been set.

### GetTimeRange

`func (o *SpendControl) GetTimeRange() SpendControlTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SpendControl) GetTimeRangeOk() (*SpendControlTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SpendControl) SetTimeRange(v SpendControlTimeRange)`

SetTimeRange sets TimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


