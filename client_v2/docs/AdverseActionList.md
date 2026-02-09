# AdverseActionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**AdverseActions** | [**[]AdverseActionResponse**](AdverseActionResponse.md) |  | 

## Methods

### NewAdverseActionList

`func NewAdverseActionList(adverseActions []AdverseActionResponse, ) *AdverseActionList`

NewAdverseActionList instantiates a new AdverseActionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdverseActionListWithDefaults

`func NewAdverseActionListWithDefaults() *AdverseActionList`

NewAdverseActionListWithDefaults instantiates a new AdverseActionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *AdverseActionList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AdverseActionList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AdverseActionList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AdverseActionList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetAdverseActions

`func (o *AdverseActionList) GetAdverseActions() []AdverseActionResponse`

GetAdverseActions returns the AdverseActions field if non-nil, zero value otherwise.

### GetAdverseActionsOk

`func (o *AdverseActionList) GetAdverseActionsOk() (*[]AdverseActionResponse, bool)`

GetAdverseActionsOk returns a tuple with the AdverseActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdverseActions

`func (o *AdverseActionList) SetAdverseActions(v []AdverseActionResponse)`

SetAdverseActions sets AdverseActions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


