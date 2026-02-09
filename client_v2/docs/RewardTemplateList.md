# RewardTemplateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**RewardTemplates** | [**[]RewardTemplateResponse**](RewardTemplateResponse.md) | Array of reward templates | 

## Methods

### NewRewardTemplateList

`func NewRewardTemplateList(rewardTemplates []RewardTemplateResponse, ) *RewardTemplateList`

NewRewardTemplateList instantiates a new RewardTemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardTemplateListWithDefaults

`func NewRewardTemplateListWithDefaults() *RewardTemplateList`

NewRewardTemplateListWithDefaults instantiates a new RewardTemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *RewardTemplateList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *RewardTemplateList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *RewardTemplateList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *RewardTemplateList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetRewardTemplates

`func (o *RewardTemplateList) GetRewardTemplates() []RewardTemplateResponse`

GetRewardTemplates returns the RewardTemplates field if non-nil, zero value otherwise.

### GetRewardTemplatesOk

`func (o *RewardTemplateList) GetRewardTemplatesOk() (*[]RewardTemplateResponse, bool)`

GetRewardTemplatesOk returns a tuple with the RewardTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardTemplates

`func (o *RewardTemplateList) SetRewardTemplates(v []RewardTemplateResponse)`

SetRewardTemplates sets RewardTemplates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


