# FeeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Fees** | [**[]FeeResponse**](FeeResponse.md) | Array of fees | 

## Methods

### NewFeeList

`func NewFeeList(fees []FeeResponse, ) *FeeList`

NewFeeList instantiates a new FeeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeListWithDefaults

`func NewFeeListWithDefaults() *FeeList`

NewFeeListWithDefaults instantiates a new FeeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *FeeList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FeeList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FeeList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FeeList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetFees

`func (o *FeeList) GetFees() []FeeResponse`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *FeeList) GetFeesOk() (*[]FeeResponse, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *FeeList) SetFees(v []FeeResponse)`

SetFees sets Fees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


