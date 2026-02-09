# RewardList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Rewards** | [**[]RewardResponse**](RewardResponse.md) | Array of rewards | 

## Methods

### NewRewardList

`func NewRewardList(rewards []RewardResponse, ) *RewardList`

NewRewardList instantiates a new RewardList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardListWithDefaults

`func NewRewardListWithDefaults() *RewardList`

NewRewardListWithDefaults instantiates a new RewardList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *RewardList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *RewardList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *RewardList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *RewardList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetRewards

`func (o *RewardList) GetRewards() []RewardResponse`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *RewardList) GetRewardsOk() (*[]RewardResponse, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *RewardList) SetRewards(v []RewardResponse)`

SetRewards sets Rewards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


