# DisputeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Disputes** | [**[]DisputeResponse**](DisputeResponse.md) | Array of disputes | 

## Methods

### NewDisputeList

`func NewDisputeList(disputes []DisputeResponse, ) *DisputeList`

NewDisputeList instantiates a new DisputeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeListWithDefaults

`func NewDisputeListWithDefaults() *DisputeList`

NewDisputeListWithDefaults instantiates a new DisputeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *DisputeList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *DisputeList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *DisputeList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *DisputeList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetDisputes

`func (o *DisputeList) GetDisputes() []DisputeResponse`

GetDisputes returns the Disputes field if non-nil, zero value otherwise.

### GetDisputesOk

`func (o *DisputeList) GetDisputesOk() (*[]DisputeResponse, bool)`

GetDisputesOk returns a tuple with the Disputes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputes

`func (o *DisputeList) SetDisputes(v []DisputeResponse)`

SetDisputes sets Disputes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


