# FdxAuthRequestList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**FdxAuthRequests** | [**[]FdxAuthRequestResponse**](FdxAuthRequestResponse.md) | Array of FDX authorization requests | 

## Methods

### NewFdxAuthRequestList

`func NewFdxAuthRequestList(fdxAuthRequests []FdxAuthRequestResponse, ) *FdxAuthRequestList`

NewFdxAuthRequestList instantiates a new FdxAuthRequestList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdxAuthRequestListWithDefaults

`func NewFdxAuthRequestListWithDefaults() *FdxAuthRequestList`

NewFdxAuthRequestListWithDefaults instantiates a new FdxAuthRequestList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *FdxAuthRequestList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FdxAuthRequestList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FdxAuthRequestList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FdxAuthRequestList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetFdxAuthRequests

`func (o *FdxAuthRequestList) GetFdxAuthRequests() []FdxAuthRequestResponse`

GetFdxAuthRequests returns the FdxAuthRequests field if non-nil, zero value otherwise.

### GetFdxAuthRequestsOk

`func (o *FdxAuthRequestList) GetFdxAuthRequestsOk() (*[]FdxAuthRequestResponse, bool)`

GetFdxAuthRequestsOk returns a tuple with the FdxAuthRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdxAuthRequests

`func (o *FdxAuthRequestList) SetFdxAuthRequests(v []FdxAuthRequestResponse)`

SetFdxAuthRequests sets FdxAuthRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


