# PatchLineOfCreditTemplateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountProgramId** | Pointer to **string** |  | [optional] 
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly determines the due date for a payment.  | [optional] 
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with.  | [optional] 
**MinimumPayment** | Pointer to [**MinimumPaymentTypeRateOrAmount**](MinimumPaymentTypeRateOrAmount.md) |  | [optional] 

## Methods

### NewPatchLineOfCreditTemplateFields

`func NewPatchLineOfCreditTemplateFields() *PatchLineOfCreditTemplateFields`

NewPatchLineOfCreditTemplateFields instantiates a new PatchLineOfCreditTemplateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchLineOfCreditTemplateFieldsWithDefaults

`func NewPatchLineOfCreditTemplateFieldsWithDefaults() *PatchLineOfCreditTemplateFields`

NewPatchLineOfCreditTemplateFieldsWithDefaults instantiates a new PatchLineOfCreditTemplateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountProgramId

`func (o *PatchLineOfCreditTemplateFields) GetAccountProgramId() string`

GetAccountProgramId returns the AccountProgramId field if non-nil, zero value otherwise.

### GetAccountProgramIdOk

`func (o *PatchLineOfCreditTemplateFields) GetAccountProgramIdOk() (*string, bool)`

GetAccountProgramIdOk returns a tuple with the AccountProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProgramId

`func (o *PatchLineOfCreditTemplateFields) SetAccountProgramId(v string)`

SetAccountProgramId sets AccountProgramId field to given value.

### HasAccountProgramId

`func (o *PatchLineOfCreditTemplateFields) HasAccountProgramId() bool`

HasAccountProgramId returns a boolean if a field has been set.

### GetGracePeriod

`func (o *PatchLineOfCreditTemplateFields) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *PatchLineOfCreditTemplateFields) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *PatchLineOfCreditTemplateFields) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *PatchLineOfCreditTemplateFields) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetInterestProductId

`func (o *PatchLineOfCreditTemplateFields) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *PatchLineOfCreditTemplateFields) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *PatchLineOfCreditTemplateFields) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *PatchLineOfCreditTemplateFields) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *PatchLineOfCreditTemplateFields) GetMinimumPayment() MinimumPaymentTypeRateOrAmount`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *PatchLineOfCreditTemplateFields) GetMinimumPaymentOk() (*MinimumPaymentTypeRateOrAmount, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *PatchLineOfCreditTemplateFields) SetMinimumPayment(v MinimumPaymentTypeRateOrAmount)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *PatchLineOfCreditTemplateFields) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


