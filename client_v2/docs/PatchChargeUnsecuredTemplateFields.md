# PatchChargeUnsecuredTemplateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountProgramId** | Pointer to **string** |  | [optional] 
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly determines the due date for a payment.  | [optional] 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 

## Methods

### NewPatchChargeUnsecuredTemplateFields

`func NewPatchChargeUnsecuredTemplateFields() *PatchChargeUnsecuredTemplateFields`

NewPatchChargeUnsecuredTemplateFields instantiates a new PatchChargeUnsecuredTemplateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchChargeUnsecuredTemplateFieldsWithDefaults

`func NewPatchChargeUnsecuredTemplateFieldsWithDefaults() *PatchChargeUnsecuredTemplateFields`

NewPatchChargeUnsecuredTemplateFieldsWithDefaults instantiates a new PatchChargeUnsecuredTemplateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountProgramId

`func (o *PatchChargeUnsecuredTemplateFields) GetAccountProgramId() string`

GetAccountProgramId returns the AccountProgramId field if non-nil, zero value otherwise.

### GetAccountProgramIdOk

`func (o *PatchChargeUnsecuredTemplateFields) GetAccountProgramIdOk() (*string, bool)`

GetAccountProgramIdOk returns a tuple with the AccountProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProgramId

`func (o *PatchChargeUnsecuredTemplateFields) SetAccountProgramId(v string)`

SetAccountProgramId sets AccountProgramId field to given value.

### HasAccountProgramId

`func (o *PatchChargeUnsecuredTemplateFields) HasAccountProgramId() bool`

HasAccountProgramId returns a boolean if a field has been set.

### GetGracePeriod

`func (o *PatchChargeUnsecuredTemplateFields) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *PatchChargeUnsecuredTemplateFields) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *PatchChargeUnsecuredTemplateFields) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *PatchChargeUnsecuredTemplateFields) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetSpendControlIds

`func (o *PatchChargeUnsecuredTemplateFields) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *PatchChargeUnsecuredTemplateFields) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *PatchChargeUnsecuredTemplateFields) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *PatchChargeUnsecuredTemplateFields) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


