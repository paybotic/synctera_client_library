# PatchAccountTemplateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountProgramId** | Pointer to **string** |  | [optional] 
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly determines the due date for a payment.  | [optional] 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 
**AutoPaymentPeriod** | Pointer to **int32** | The number of days past the billing period to initiate an auto payment. Only applicable for accounts with type &#x60;CHARGE_SECURED&#x60;, where the account holder has opted in for auto payment functionality. This value must be lower than or equal the &#x60;grace_period&#x60; setting on the account. If this value is 0, the auto payment will happen on the same day as the statement is generated. Auto payment only occurs if regular payments are not received on time.  | [optional] 
**MinimumPayment** | Pointer to [**MinimumPaymentTypeRateOrAmount**](MinimumPaymentTypeRateOrAmount.md) |  | [optional] 
**ApplicationWorkflowId** | Pointer to **string** | Taktile workflow ID for credit application processing | [optional] 
**VendorInfo** | Pointer to [**TemplateVendorInfo**](TemplateVendorInfo.md) |  | [optional] 
**IsAchEnabled** | Pointer to **bool** | A flag to indicate whether ACH transactions are enabled. | [optional] 
**IsCardEnabled** | Pointer to **bool** | A flag to indicate whether card transactions are enabled. | [optional] 
**IsEftCaEnabled** | Pointer to **bool** | A flag to indicate whether EFT Canada transactions are enabled. | [optional] 
**IsExternalCardEnabled** | Pointer to **bool** | A flag to indicate whether external card transactions are enabled. | [optional] 
**IsP2pEnabled** | Pointer to **bool** | A flag to indicate whether P2P transactions are enabled. | [optional] 
**IsSyncteraPayEnabled** | Pointer to **bool** | A flag to indicate whether Synctera Pay transactions are enabled. | [optional] 
**IsWireEnabled** | Pointer to **bool** | A flag to indicate whether wire transactions are enabled. | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with.  | [optional] 
**IsSarEnabled** | Pointer to **bool** | Enable SAR report. | [optional] 

## Methods

### NewPatchAccountTemplateFields

`func NewPatchAccountTemplateFields() *PatchAccountTemplateFields`

NewPatchAccountTemplateFields instantiates a new PatchAccountTemplateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchAccountTemplateFieldsWithDefaults

`func NewPatchAccountTemplateFieldsWithDefaults() *PatchAccountTemplateFields`

NewPatchAccountTemplateFieldsWithDefaults instantiates a new PatchAccountTemplateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountProgramId

`func (o *PatchAccountTemplateFields) GetAccountProgramId() string`

GetAccountProgramId returns the AccountProgramId field if non-nil, zero value otherwise.

### GetAccountProgramIdOk

`func (o *PatchAccountTemplateFields) GetAccountProgramIdOk() (*string, bool)`

GetAccountProgramIdOk returns a tuple with the AccountProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProgramId

`func (o *PatchAccountTemplateFields) SetAccountProgramId(v string)`

SetAccountProgramId sets AccountProgramId field to given value.

### HasAccountProgramId

`func (o *PatchAccountTemplateFields) HasAccountProgramId() bool`

HasAccountProgramId returns a boolean if a field has been set.

### GetGracePeriod

`func (o *PatchAccountTemplateFields) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *PatchAccountTemplateFields) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *PatchAccountTemplateFields) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *PatchAccountTemplateFields) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetSpendControlIds

`func (o *PatchAccountTemplateFields) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *PatchAccountTemplateFields) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *PatchAccountTemplateFields) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *PatchAccountTemplateFields) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.

### GetAutoPaymentPeriod

`func (o *PatchAccountTemplateFields) GetAutoPaymentPeriod() int32`

GetAutoPaymentPeriod returns the AutoPaymentPeriod field if non-nil, zero value otherwise.

### GetAutoPaymentPeriodOk

`func (o *PatchAccountTemplateFields) GetAutoPaymentPeriodOk() (*int32, bool)`

GetAutoPaymentPeriodOk returns a tuple with the AutoPaymentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPaymentPeriod

`func (o *PatchAccountTemplateFields) SetAutoPaymentPeriod(v int32)`

SetAutoPaymentPeriod sets AutoPaymentPeriod field to given value.

### HasAutoPaymentPeriod

`func (o *PatchAccountTemplateFields) HasAutoPaymentPeriod() bool`

HasAutoPaymentPeriod returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *PatchAccountTemplateFields) GetMinimumPayment() MinimumPaymentTypeRateOrAmount`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *PatchAccountTemplateFields) GetMinimumPaymentOk() (*MinimumPaymentTypeRateOrAmount, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *PatchAccountTemplateFields) SetMinimumPayment(v MinimumPaymentTypeRateOrAmount)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *PatchAccountTemplateFields) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.

### GetApplicationWorkflowId

`func (o *PatchAccountTemplateFields) GetApplicationWorkflowId() string`

GetApplicationWorkflowId returns the ApplicationWorkflowId field if non-nil, zero value otherwise.

### GetApplicationWorkflowIdOk

`func (o *PatchAccountTemplateFields) GetApplicationWorkflowIdOk() (*string, bool)`

GetApplicationWorkflowIdOk returns a tuple with the ApplicationWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationWorkflowId

`func (o *PatchAccountTemplateFields) SetApplicationWorkflowId(v string)`

SetApplicationWorkflowId sets ApplicationWorkflowId field to given value.

### HasApplicationWorkflowId

`func (o *PatchAccountTemplateFields) HasApplicationWorkflowId() bool`

HasApplicationWorkflowId returns a boolean if a field has been set.

### GetVendorInfo

`func (o *PatchAccountTemplateFields) GetVendorInfo() TemplateVendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *PatchAccountTemplateFields) GetVendorInfoOk() (*TemplateVendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *PatchAccountTemplateFields) SetVendorInfo(v TemplateVendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *PatchAccountTemplateFields) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.

### GetIsAchEnabled

`func (o *PatchAccountTemplateFields) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *PatchAccountTemplateFields) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *PatchAccountTemplateFields) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *PatchAccountTemplateFields) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *PatchAccountTemplateFields) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *PatchAccountTemplateFields) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *PatchAccountTemplateFields) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *PatchAccountTemplateFields) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsEftCaEnabled

`func (o *PatchAccountTemplateFields) GetIsEftCaEnabled() bool`

GetIsEftCaEnabled returns the IsEftCaEnabled field if non-nil, zero value otherwise.

### GetIsEftCaEnabledOk

`func (o *PatchAccountTemplateFields) GetIsEftCaEnabledOk() (*bool, bool)`

GetIsEftCaEnabledOk returns a tuple with the IsEftCaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEftCaEnabled

`func (o *PatchAccountTemplateFields) SetIsEftCaEnabled(v bool)`

SetIsEftCaEnabled sets IsEftCaEnabled field to given value.

### HasIsEftCaEnabled

`func (o *PatchAccountTemplateFields) HasIsEftCaEnabled() bool`

HasIsEftCaEnabled returns a boolean if a field has been set.

### GetIsExternalCardEnabled

`func (o *PatchAccountTemplateFields) GetIsExternalCardEnabled() bool`

GetIsExternalCardEnabled returns the IsExternalCardEnabled field if non-nil, zero value otherwise.

### GetIsExternalCardEnabledOk

`func (o *PatchAccountTemplateFields) GetIsExternalCardEnabledOk() (*bool, bool)`

GetIsExternalCardEnabledOk returns a tuple with the IsExternalCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalCardEnabled

`func (o *PatchAccountTemplateFields) SetIsExternalCardEnabled(v bool)`

SetIsExternalCardEnabled sets IsExternalCardEnabled field to given value.

### HasIsExternalCardEnabled

`func (o *PatchAccountTemplateFields) HasIsExternalCardEnabled() bool`

HasIsExternalCardEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *PatchAccountTemplateFields) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *PatchAccountTemplateFields) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *PatchAccountTemplateFields) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *PatchAccountTemplateFields) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsSyncteraPayEnabled

`func (o *PatchAccountTemplateFields) GetIsSyncteraPayEnabled() bool`

GetIsSyncteraPayEnabled returns the IsSyncteraPayEnabled field if non-nil, zero value otherwise.

### GetIsSyncteraPayEnabledOk

`func (o *PatchAccountTemplateFields) GetIsSyncteraPayEnabledOk() (*bool, bool)`

GetIsSyncteraPayEnabledOk returns a tuple with the IsSyncteraPayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncteraPayEnabled

`func (o *PatchAccountTemplateFields) SetIsSyncteraPayEnabled(v bool)`

SetIsSyncteraPayEnabled sets IsSyncteraPayEnabled field to given value.

### HasIsSyncteraPayEnabled

`func (o *PatchAccountTemplateFields) HasIsSyncteraPayEnabled() bool`

HasIsSyncteraPayEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *PatchAccountTemplateFields) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *PatchAccountTemplateFields) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *PatchAccountTemplateFields) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *PatchAccountTemplateFields) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetAccountType

`func (o *PatchAccountTemplateFields) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PatchAccountTemplateFields) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PatchAccountTemplateFields) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *PatchAccountTemplateFields) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetInterestProductId

`func (o *PatchAccountTemplateFields) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *PatchAccountTemplateFields) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *PatchAccountTemplateFields) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *PatchAccountTemplateFields) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetIsSarEnabled

`func (o *PatchAccountTemplateFields) GetIsSarEnabled() bool`

GetIsSarEnabled returns the IsSarEnabled field if non-nil, zero value otherwise.

### GetIsSarEnabledOk

`func (o *PatchAccountTemplateFields) GetIsSarEnabledOk() (*bool, bool)`

GetIsSarEnabledOk returns a tuple with the IsSarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSarEnabled

`func (o *PatchAccountTemplateFields) SetIsSarEnabled(v bool)`

SetIsSarEnabled sets IsSarEnabled field to given value.

### HasIsSarEnabled

`func (o *PatchAccountTemplateFields) HasIsSarEnabled() bool`

HasIsSarEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


