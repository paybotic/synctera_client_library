# FeePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the business or customer account being charged the fee. | 
**Amount** | Pointer to **int32** | The amount of the fee in ISO 4217 minor currency units, e.g. cents. The internal account referenced by the fee template will be debited this amount. Defaults to the value in the fee template.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Note** | Pointer to **string** | An optional note for this instance of the fee. | [optional] 
**ReferenceId** | Pointer to **string** | An optional reference ID that uniquely identifies this fee instance. This ID will be included in the posted fee transaction. | [optional] 
**TemplateId** | **string** | The ID of the fee template to use to create the fee. Values from the fee template will be used as defaults for the fee. Note that the fee template may have been updated since the fee was created and that such subsequent updates to the fee template do not affect existing fees.  | 

## Methods

### NewFeePost

`func NewFeePost(accountId string, templateId string, ) *FeePost`

NewFeePost instantiates a new FeePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeePostWithDefaults

`func NewFeePostWithDefaults() *FeePost`

NewFeePostWithDefaults instantiates a new FeePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *FeePost) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *FeePost) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *FeePost) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *FeePost) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeePost) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeePost) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeePost) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMetadata

`func (o *FeePost) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeePost) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeePost) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeePost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNote

`func (o *FeePost) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *FeePost) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *FeePost) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *FeePost) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReferenceId

`func (o *FeePost) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *FeePost) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *FeePost) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *FeePost) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetTemplateId

`func (o *FeePost) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *FeePost) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *FeePost) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


