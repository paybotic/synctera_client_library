# EftCaList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Transfers** | [**[]EftCaResponse**](EftCaResponse.md) | Array of transfers. | 

## Methods

### NewEftCaList

`func NewEftCaList(transfers []EftCaResponse, ) *EftCaList`

NewEftCaList instantiates a new EftCaList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEftCaListWithDefaults

`func NewEftCaListWithDefaults() *EftCaList`

NewEftCaListWithDefaults instantiates a new EftCaList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *EftCaList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *EftCaList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *EftCaList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *EftCaList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTransfers

`func (o *EftCaList) GetTransfers() []EftCaResponse`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *EftCaList) GetTransfersOk() (*[]EftCaResponse, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *EftCaList) SetTransfers(v []EftCaResponse)`

SetTransfers sets Transfers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


