# RewardBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The ID of the business or customer account being charged the reward. | [optional] 
**Amount** | Pointer to **int32** | The amount of the reward in ISO 4217 minor currency units, e.g. cents. The internal account referenced by the reward template will be debited this amount. Defaults to the value in the reward template.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Note** | Pointer to **string** | An optional note for this instance of the reward. | [optional] 
**TemplateId** | Pointer to **string** | The ID of the reward template to use to create the reward. Values from the reward template will be used as defaults for the reward. Note that the reward template may have been updated since the reward was created and that such subsequent updates to the reward template do not affect existing rewards.  | [optional] 

## Methods

### NewRewardBase

`func NewRewardBase() *RewardBase`

NewRewardBase instantiates a new RewardBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardBaseWithDefaults

`func NewRewardBaseWithDefaults() *RewardBase`

NewRewardBaseWithDefaults instantiates a new RewardBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RewardBase) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RewardBase) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RewardBase) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *RewardBase) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *RewardBase) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardBase) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardBase) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RewardBase) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *RewardBase) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RewardBase) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RewardBase) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RewardBase) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNote

`func (o *RewardBase) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RewardBase) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RewardBase) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *RewardBase) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTemplateId

`func (o *RewardBase) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *RewardBase) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *RewardBase) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *RewardBase) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


