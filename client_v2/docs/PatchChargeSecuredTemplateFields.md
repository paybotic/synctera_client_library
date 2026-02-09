# PatchChargeSecuredTemplateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountProgramId** | Pointer to **string** |  | [optional] 
**AutoPaymentPeriod** | Pointer to **int32** | The number of days past the billing period to initiate an auto payment. Only applicable for accounts with type &#x60;CHARGE_SECURED&#x60;, where the account holder has opted in for auto payment functionality. This value must be lower than or equal the &#x60;grace_period&#x60; setting on the account. If this value is 0, the auto payment will happen on the same day as the statement is generated. Auto payment only occurs if regular payments are not received on time.  | [optional] 
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly determines the due date for a payment.  | [optional] 
**MinimumPayment** | Pointer to [**MinimumPaymentTypeFull**](MinimumPaymentTypeFull.md) |  | [optional] 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 

## Methods

### NewPatchChargeSecuredTemplateFields

`func NewPatchChargeSecuredTemplateFields() *PatchChargeSecuredTemplateFields`

NewPatchChargeSecuredTemplateFields instantiates a new PatchChargeSecuredTemplateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchChargeSecuredTemplateFieldsWithDefaults

`func NewPatchChargeSecuredTemplateFieldsWithDefaults() *PatchChargeSecuredTemplateFields`

NewPatchChargeSecuredTemplateFieldsWithDefaults instantiates a new PatchChargeSecuredTemplateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountProgramId

`func (o *PatchChargeSecuredTemplateFields) GetAccountProgramId() string`

GetAccountProgramId returns the AccountProgramId field if non-nil, zero value otherwise.

### GetAccountProgramIdOk

`func (o *PatchChargeSecuredTemplateFields) GetAccountProgramIdOk() (*string, bool)`

GetAccountProgramIdOk returns a tuple with the AccountProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProgramId

`func (o *PatchChargeSecuredTemplateFields) SetAccountProgramId(v string)`

SetAccountProgramId sets AccountProgramId field to given value.

### HasAccountProgramId

`func (o *PatchChargeSecuredTemplateFields) HasAccountProgramId() bool`

HasAccountProgramId returns a boolean if a field has been set.

### GetAutoPaymentPeriod

`func (o *PatchChargeSecuredTemplateFields) GetAutoPaymentPeriod() int32`

GetAutoPaymentPeriod returns the AutoPaymentPeriod field if non-nil, zero value otherwise.

### GetAutoPaymentPeriodOk

`func (o *PatchChargeSecuredTemplateFields) GetAutoPaymentPeriodOk() (*int32, bool)`

GetAutoPaymentPeriodOk returns a tuple with the AutoPaymentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPaymentPeriod

`func (o *PatchChargeSecuredTemplateFields) SetAutoPaymentPeriod(v int32)`

SetAutoPaymentPeriod sets AutoPaymentPeriod field to given value.

### HasAutoPaymentPeriod

`func (o *PatchChargeSecuredTemplateFields) HasAutoPaymentPeriod() bool`

HasAutoPaymentPeriod returns a boolean if a field has been set.

### GetGracePeriod

`func (o *PatchChargeSecuredTemplateFields) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *PatchChargeSecuredTemplateFields) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *PatchChargeSecuredTemplateFields) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *PatchChargeSecuredTemplateFields) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *PatchChargeSecuredTemplateFields) GetMinimumPayment() MinimumPaymentTypeFull`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *PatchChargeSecuredTemplateFields) GetMinimumPaymentOk() (*MinimumPaymentTypeFull, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *PatchChargeSecuredTemplateFields) SetMinimumPayment(v MinimumPaymentTypeFull)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *PatchChargeSecuredTemplateFields) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.

### GetSpendControlIds

`func (o *PatchChargeSecuredTemplateFields) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *PatchChargeSecuredTemplateFields) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *PatchChargeSecuredTemplateFields) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *PatchChargeSecuredTemplateFields) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


