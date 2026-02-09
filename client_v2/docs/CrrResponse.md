# CrrResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | ID of the customer that the risk score applies to. Either customer_id or business_id will be provided but not both | [optional] 
**CalculationsBreakdown** | Pointer to [**[]CalculationBreakdown**](CalculationBreakdown.md) | An array that holds the calculated weights for each risk parameter that was used in the score calculation | [optional] 
**CreatorId** | Pointer to **string** | ID of the user that created risk score. creator_id will be null if score was calculated automatically | [optional] 
**CustomerId** | Pointer to **string** | ID of the customer that the risk score applies to. Either customer_id or business_id will be provided but not both | [optional] 
**Id** | **string** | Risk score record unique ID | 
**Note** | Pointer to **string** | The attached note for the risk score. An note is always added when a score has been overridden by a user | [optional] 
**RiskLevel** | **string** | Risk score classification | 
**RiskScore** | **int32** | The calculated risk score for the customer/business | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**ValidFrom** | Pointer to **time.Time** | The start date that this risk score came into effect. | [optional] 
**ValidTo** | Pointer to **time.Time** | The end date that this risk score came into effect. | [optional] 

## Methods

### NewCrrResponse

`func NewCrrResponse(id string, riskLevel string, riskScore int32, tenant string, ) *CrrResponse`

NewCrrResponse instantiates a new CrrResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrrResponseWithDefaults

`func NewCrrResponseWithDefaults() *CrrResponse`

NewCrrResponseWithDefaults instantiates a new CrrResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *CrrResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *CrrResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *CrrResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *CrrResponse) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCalculationsBreakdown

`func (o *CrrResponse) GetCalculationsBreakdown() []CalculationBreakdown`

GetCalculationsBreakdown returns the CalculationsBreakdown field if non-nil, zero value otherwise.

### GetCalculationsBreakdownOk

`func (o *CrrResponse) GetCalculationsBreakdownOk() (*[]CalculationBreakdown, bool)`

GetCalculationsBreakdownOk returns a tuple with the CalculationsBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationsBreakdown

`func (o *CrrResponse) SetCalculationsBreakdown(v []CalculationBreakdown)`

SetCalculationsBreakdown sets CalculationsBreakdown field to given value.

### HasCalculationsBreakdown

`func (o *CrrResponse) HasCalculationsBreakdown() bool`

HasCalculationsBreakdown returns a boolean if a field has been set.

### GetCreatorId

`func (o *CrrResponse) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CrrResponse) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CrrResponse) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *CrrResponse) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCustomerId

`func (o *CrrResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CrrResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CrrResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CrrResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *CrrResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrrResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrrResponse) SetId(v string)`

SetId sets Id field to given value.


### GetNote

`func (o *CrrResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CrrResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CrrResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CrrResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRiskLevel

`func (o *CrrResponse) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *CrrResponse) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *CrrResponse) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.


### GetRiskScore

`func (o *CrrResponse) GetRiskScore() int32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *CrrResponse) GetRiskScoreOk() (*int32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *CrrResponse) SetRiskScore(v int32)`

SetRiskScore sets RiskScore field to given value.


### GetTenant

`func (o *CrrResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CrrResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CrrResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetValidFrom

`func (o *CrrResponse) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *CrrResponse) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *CrrResponse) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *CrrResponse) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *CrrResponse) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *CrrResponse) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *CrrResponse) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *CrrResponse) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


