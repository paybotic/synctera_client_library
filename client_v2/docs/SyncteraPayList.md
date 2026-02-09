# SyncteraPayList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Transfers** | [**[]SyncteraPayResponse**](SyncteraPayResponse.md) | Array of Outgoing Synctera Pay transfers. | 

## Methods

### NewSyncteraPayList

`func NewSyncteraPayList(transfers []SyncteraPayResponse, ) *SyncteraPayList`

NewSyncteraPayList instantiates a new SyncteraPayList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayListWithDefaults

`func NewSyncteraPayListWithDefaults() *SyncteraPayList`

NewSyncteraPayListWithDefaults instantiates a new SyncteraPayList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *SyncteraPayList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *SyncteraPayList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *SyncteraPayList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *SyncteraPayList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTransfers

`func (o *SyncteraPayList) GetTransfers() []SyncteraPayResponse`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *SyncteraPayList) GetTransfersOk() (*[]SyncteraPayResponse, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *SyncteraPayList) SetTransfers(v []SyncteraPayResponse)`

SetTransfers sets Transfers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


