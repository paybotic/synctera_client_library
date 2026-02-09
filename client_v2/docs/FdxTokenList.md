# FdxTokenList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**FdxTokens** | [**[]FdxTokenResponse**](FdxTokenResponse.md) | Array of FDX tokens | 

## Methods

### NewFdxTokenList

`func NewFdxTokenList(fdxTokens []FdxTokenResponse, ) *FdxTokenList`

NewFdxTokenList instantiates a new FdxTokenList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdxTokenListWithDefaults

`func NewFdxTokenListWithDefaults() *FdxTokenList`

NewFdxTokenListWithDefaults instantiates a new FdxTokenList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *FdxTokenList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FdxTokenList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FdxTokenList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FdxTokenList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetFdxTokens

`func (o *FdxTokenList) GetFdxTokens() []FdxTokenResponse`

GetFdxTokens returns the FdxTokens field if non-nil, zero value otherwise.

### GetFdxTokensOk

`func (o *FdxTokenList) GetFdxTokensOk() (*[]FdxTokenResponse, bool)`

GetFdxTokensOk returns a tuple with the FdxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdxTokens

`func (o *FdxTokenList) SetFdxTokens(v []FdxTokenResponse)`

SetFdxTokens sets FdxTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


