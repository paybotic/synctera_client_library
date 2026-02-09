# CreditCardApplicationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditLimit** | Pointer to **int32** | Credit limit in cents (e.g., 250000000 &#x3D; $2,500,000) | [optional] 
**Decision** | Pointer to [**CreditCardDecision**](CreditCardDecision.md) |  | [optional] 
**ExternalId** | Pointer to **string** | External identifier (e.g., Taktile decision_id) | [optional] 
**InterestRate** | Pointer to **int32** | Interest rate in basis points (e.g., 1250 &#x3D; 12.50%) | [optional] 
**RejectionReasons** | Pointer to **[]string** | Reasons for rejection (when decision is declined) | [optional] 
**Status** | Pointer to [**CreditCardApplicationStatus**](CreditCardApplicationStatus.md) |  | [optional] 

## Methods

### NewCreditCardApplicationPatch

`func NewCreditCardApplicationPatch() *CreditCardApplicationPatch`

NewCreditCardApplicationPatch instantiates a new CreditCardApplicationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardApplicationPatchWithDefaults

`func NewCreditCardApplicationPatchWithDefaults() *CreditCardApplicationPatch`

NewCreditCardApplicationPatchWithDefaults instantiates a new CreditCardApplicationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditLimit

`func (o *CreditCardApplicationPatch) GetCreditLimit() int32`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *CreditCardApplicationPatch) GetCreditLimitOk() (*int32, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *CreditCardApplicationPatch) SetCreditLimit(v int32)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *CreditCardApplicationPatch) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetDecision

`func (o *CreditCardApplicationPatch) GetDecision() CreditCardDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *CreditCardApplicationPatch) GetDecisionOk() (*CreditCardDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *CreditCardApplicationPatch) SetDecision(v CreditCardDecision)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *CreditCardApplicationPatch) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetExternalId

`func (o *CreditCardApplicationPatch) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreditCardApplicationPatch) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreditCardApplicationPatch) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreditCardApplicationPatch) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetInterestRate

`func (o *CreditCardApplicationPatch) GetInterestRate() int32`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *CreditCardApplicationPatch) GetInterestRateOk() (*int32, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *CreditCardApplicationPatch) SetInterestRate(v int32)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *CreditCardApplicationPatch) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetRejectionReasons

`func (o *CreditCardApplicationPatch) GetRejectionReasons() []string`

GetRejectionReasons returns the RejectionReasons field if non-nil, zero value otherwise.

### GetRejectionReasonsOk

`func (o *CreditCardApplicationPatch) GetRejectionReasonsOk() (*[]string, bool)`

GetRejectionReasonsOk returns a tuple with the RejectionReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReasons

`func (o *CreditCardApplicationPatch) SetRejectionReasons(v []string)`

SetRejectionReasons sets RejectionReasons field to given value.

### HasRejectionReasons

`func (o *CreditCardApplicationPatch) HasRejectionReasons() bool`

HasRejectionReasons returns a boolean if a field has been set.

### GetStatus

`func (o *CreditCardApplicationPatch) GetStatus() CreditCardApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditCardApplicationPatch) GetStatusOk() (*CreditCardApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditCardApplicationPatch) SetStatus(v CreditCardApplicationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreditCardApplicationPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


