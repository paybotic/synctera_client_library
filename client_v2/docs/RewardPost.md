# RewardPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the business or customer account being charged the reward. | 
**Amount** | Pointer to **int32** | The amount of the reward in ISO 4217 minor currency units, e.g. cents. The internal account referenced by the reward template will be debited this amount. Defaults to the value in the reward template.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Note** | Pointer to **string** | An optional note for this instance of the reward. | [optional] 
**TemplateId** | **string** | The ID of the reward template to use to create the reward. Values from the reward template will be used as defaults for the reward. Note that the reward template may have been updated since the reward was created and that such subsequent updates to the reward template do not affect existing rewards.  | 

## Methods

### NewRewardPost

`func NewRewardPost(accountId string, templateId string, ) *RewardPost`

NewRewardPost instantiates a new RewardPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardPostWithDefaults

`func NewRewardPostWithDefaults() *RewardPost`

NewRewardPostWithDefaults instantiates a new RewardPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RewardPost) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RewardPost) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RewardPost) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *RewardPost) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardPost) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardPost) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RewardPost) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *RewardPost) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RewardPost) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RewardPost) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RewardPost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNote

`func (o *RewardPost) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *RewardPost) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *RewardPost) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *RewardPost) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTemplateId

`func (o *RewardPost) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *RewardPost) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *RewardPost) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


