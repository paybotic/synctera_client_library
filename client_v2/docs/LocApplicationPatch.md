# LocApplicationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditLimit** | Pointer to **int32** | Credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 
**Decision** | Pointer to [**LocDecision**](LocDecision.md) |  | [optional] 
**ExternalId** | Pointer to **string** | External identifier (e.g., Taktile decision_id) | [optional] 
**InterestRate** | Pointer to **int32** | Interest rate in basis points (e.g., 1250 &#x3D; 12.50%) | [optional] 
**RejectionReasons** | Pointer to **[]string** | Reasons for rejection (when decision is declined) | [optional] 
**Status** | Pointer to [**LocApplicationStatus**](LocApplicationStatus.md) |  | [optional] 

## Methods

### NewLocApplicationPatch

`func NewLocApplicationPatch() *LocApplicationPatch`

NewLocApplicationPatch instantiates a new LocApplicationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocApplicationPatchWithDefaults

`func NewLocApplicationPatchWithDefaults() *LocApplicationPatch`

NewLocApplicationPatchWithDefaults instantiates a new LocApplicationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditLimit

`func (o *LocApplicationPatch) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *LocApplicationPatch) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *LocApplicationPatch) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *LocApplicationPatch) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetDecision

`func (o *LocApplicationPatch) GetDecision() LocDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *LocApplicationPatch) GetDecisionOk() (*LocDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *LocApplicationPatch) SetDecision(v LocDecision)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *LocApplicationPatch) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetExternalId

`func (o *LocApplicationPatch) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *LocApplicationPatch) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *LocApplicationPatch) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *LocApplicationPatch) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInterestRate

`func (o *LocApplicationPatch) GetInterestRate() int32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *LocApplicationPatch) GetInterestRateOk() (*int32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *LocApplicationPatch) SetInterestRate(v int32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *LocApplicationPatch) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetRejectionReasons

`func (o *LocApplicationPatch) GetRejectionReasons() []string`

GetRejectionReasons returns the RejectionReasons field if non-nil, zero value otherwise.

### GetRejectionReasonsOk

`func (o *LocApplicationPatch) GetRejectionReasonsOk() (*[]string, bool)`

GetRejectionReasonsOk returns a tuple with the RejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReasons

`func (o *LocApplicationPatch) SetRejectionReasons(v []string)`

SetRejectionReasons sets RejectionReasons field to given value.

### HasRejectionReasons

`func (o *LocApplicationPatch) HasRejectionReasons() bool`

HasRejectionReasons returns a boolean if a field has been set.

### GetStatus

`func (o *LocApplicationPatch) GetStatus() LocApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LocApplicationPatch) GetStatusOk() (*LocApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LocApplicationPatch) SetStatus(v LocApplicationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LocApplicationPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


