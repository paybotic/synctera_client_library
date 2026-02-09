# PatchDepositoryTemplateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAchEnabled** | Pointer to **bool** | A flag to indicate whether ACH transactions are enabled. | [optional] 
**IsCardEnabled** | Pointer to **bool** | A flag to indicate whether card transactions are enabled. | [optional] 
**IsEftCaEnabled** | Pointer to **bool** | A flag to indicate whether EFT Canada transactions are enabled. | [optional] 
**IsExternalCardEnabled** | Pointer to **bool** | A flag to indicate whether external card transactions are enabled. | [optional] 
**IsP2pEnabled** | Pointer to **bool** | A flag to indicate whether P2P transactions are enabled. | [optional] 
**IsSyncteraPayEnabled** | Pointer to **bool** | A flag to indicate whether Synctera Pay transactions are enabled. | [optional] 
**IsWireEnabled** | Pointer to **bool** | A flag to indicate whether wire transactions are enabled. | [optional] 
**AccountProgramId** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with.  | [optional] 
**IsSarEnabled** | Pointer to **bool** | Enable SAR report. | [optional] 
**SpendControlIds** | Pointer to **[]string** | List of spend control IDs to control spending for the account | [optional] 

## Methods

### NewPatchDepositoryTemplateFields

`func NewPatchDepositoryTemplateFields() *PatchDepositoryTemplateFields`

NewPatchDepositoryTemplateFields instantiates a new PatchDepositoryTemplateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchDepositoryTemplateFieldsWithDefaults

`func NewPatchDepositoryTemplateFieldsWithDefaults() *PatchDepositoryTemplateFields`

NewPatchDepositoryTemplateFieldsWithDefaults instantiates a new PatchDepositoryTemplateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAchEnabled

`func (o *PatchDepositoryTemplateFields) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *PatchDepositoryTemplateFields) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *PatchDepositoryTemplateFields) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *PatchDepositoryTemplateFields) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *PatchDepositoryTemplateFields) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *PatchDepositoryTemplateFields) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *PatchDepositoryTemplateFields) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *PatchDepositoryTemplateFields) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsEftCaEnabled

`func (o *PatchDepositoryTemplateFields) GetIsEftCaEnabled() bool`

GetIsEftCaEnabled returns the IsEftCaEnabled field if non-nil, zero value otherwise.

### GetIsEftCaEnabledOk

`func (o *PatchDepositoryTemplateFields) GetIsEftCaEnabledOk() (*bool, bool)`

GetIsEftCaEnabledOk returns a tuple with the IsEftCaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEftCaEnabled

`func (o *PatchDepositoryTemplateFields) SetIsEftCaEnabled(v bool)`

SetIsEftCaEnabled sets IsEftCaEnabled field to given value.

### HasIsEftCaEnabled

`func (o *PatchDepositoryTemplateFields) HasIsEftCaEnabled() bool`

HasIsEftCaEnabled returns a boolean if a field has been set.

### GetIsExternalCardEnabled

`func (o *PatchDepositoryTemplateFields) GetIsExternalCardEnabled() bool`

GetIsExternalCardEnabled returns the IsExternalCardEnabled field if non-nil, zero value otherwise.

### GetIsExternalCardEnabledOk

`func (o *PatchDepositoryTemplateFields) GetIsExternalCardEnabledOk() (*bool, bool)`

GetIsExternalCardEnabledOk returns a tuple with the IsExternalCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalCardEnabled

`func (o *PatchDepositoryTemplateFields) SetIsExternalCardEnabled(v bool)`

SetIsExternalCardEnabled sets IsExternalCardEnabled field to given value.

### HasIsExternalCardEnabled

`func (o *PatchDepositoryTemplateFields) HasIsExternalCardEnabled() bool`

HasIsExternalCardEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *PatchDepositoryTemplateFields) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *PatchDepositoryTemplateFields) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *PatchDepositoryTemplateFields) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *PatchDepositoryTemplateFields) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsSyncteraPayEnabled

`func (o *PatchDepositoryTemplateFields) GetIsSyncteraPayEnabled() bool`

GetIsSyncteraPayEnabled returns the IsSyncteraPayEnabled field if non-nil, zero value otherwise.

### GetIsSyncteraPayEnabledOk

`func (o *PatchDepositoryTemplateFields) GetIsSyncteraPayEnabledOk() (*bool, bool)`

GetIsSyncteraPayEnabledOk returns a tuple with the IsSyncteraPayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncteraPayEnabled

`func (o *PatchDepositoryTemplateFields) SetIsSyncteraPayEnabled(v bool)`

SetIsSyncteraPayEnabled sets IsSyncteraPayEnabled field to given value.

### HasIsSyncteraPayEnabled

`func (o *PatchDepositoryTemplateFields) HasIsSyncteraPayEnabled() bool`

HasIsSyncteraPayEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *PatchDepositoryTemplateFields) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *PatchDepositoryTemplateFields) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *PatchDepositoryTemplateFields) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *PatchDepositoryTemplateFields) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetAccountProgramId

`func (o *PatchDepositoryTemplateFields) GetAccountProgramId() string`

GetAccountProgramId returns the AccountProgramId field if non-nil, zero value otherwise.

### GetAccountProgramIdOk

`func (o *PatchDepositoryTemplateFields) GetAccountProgramIdOk() (*string, bool)`

GetAccountProgramIdOk returns a tuple with the AccountProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProgramId

`func (o *PatchDepositoryTemplateFields) SetAccountProgramId(v string)`

SetAccountProgramId sets AccountProgramId field to given value.

### HasAccountProgramId

`func (o *PatchDepositoryTemplateFields) HasAccountProgramId() bool`

HasAccountProgramId returns a boolean if a field has been set.

### GetAccountType

`func (o *PatchDepositoryTemplateFields) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PatchDepositoryTemplateFields) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PatchDepositoryTemplateFields) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *PatchDepositoryTemplateFields) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetInterestProductId

`func (o *PatchDepositoryTemplateFields) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *PatchDepositoryTemplateFields) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *PatchDepositoryTemplateFields) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *PatchDepositoryTemplateFields) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetIsSarEnabled

`func (o *PatchDepositoryTemplateFields) GetIsSarEnabled() bool`

GetIsSarEnabled returns the IsSarEnabled field if non-nil, zero value otherwise.

### GetIsSarEnabledOk

`func (o *PatchDepositoryTemplateFields) GetIsSarEnabledOk() (*bool, bool)`

GetIsSarEnabledOk returns a tuple with the IsSarEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSarEnabled

`func (o *PatchDepositoryTemplateFields) SetIsSarEnabled(v bool)`

SetIsSarEnabled sets IsSarEnabled field to given value.

### HasIsSarEnabled

`func (o *PatchDepositoryTemplateFields) HasIsSarEnabled() bool`

HasIsSarEnabled returns a boolean if a field has been set.

### GetSpendControlIds

`func (o *PatchDepositoryTemplateFields) GetSpendControlIds() []string`

GetSpendControlIds returns the SpendControlIds field if non-nil, zero value otherwise.

### GetSpendControlIdsOk

`func (o *PatchDepositoryTemplateFields) GetSpendControlIdsOk() (*[]string, bool)`

GetSpendControlIdsOk returns a tuple with the SpendControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControlIds

`func (o *PatchDepositoryTemplateFields) SetSpendControlIds(v []string)`

SetSpendControlIds sets SpendControlIds field to given value.

### HasSpendControlIds

`func (o *PatchDepositoryTemplateFields) HasSpendControlIds() bool`

HasSpendControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


