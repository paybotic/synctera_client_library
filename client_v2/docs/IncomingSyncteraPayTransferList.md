# IncomingSyncteraPayTransferList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Transfers** | [**[]IncomingSyncteraPayTransferResponse**](IncomingSyncteraPayTransferResponse.md) |  | 

## Methods

### NewIncomingSyncteraPayTransferList

`func NewIncomingSyncteraPayTransferList(transfers []IncomingSyncteraPayTransferResponse, ) *IncomingSyncteraPayTransferList`

NewIncomingSyncteraPayTransferList instantiates a new IncomingSyncteraPayTransferList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingSyncteraPayTransferListWithDefaults

`func NewIncomingSyncteraPayTransferListWithDefaults() *IncomingSyncteraPayTransferList`

NewIncomingSyncteraPayTransferListWithDefaults instantiates a new IncomingSyncteraPayTransferList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *IncomingSyncteraPayTransferList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *IncomingSyncteraPayTransferList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *IncomingSyncteraPayTransferList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *IncomingSyncteraPayTransferList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTransfers

`func (o *IncomingSyncteraPayTransferList) GetTransfers() []IncomingSyncteraPayTransferResponse`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *IncomingSyncteraPayTransferList) GetTransfersOk() (*[]IncomingSyncteraPayTransferResponse, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *IncomingSyncteraPayTransferList) SetTransfers(v []IncomingSyncteraPayTransferResponse)`

SetTransfers sets Transfers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


