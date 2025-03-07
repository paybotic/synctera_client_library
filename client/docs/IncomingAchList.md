# IncomingAchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Transactions** | [**[]IncomingAch**](IncomingAch.md) | Array of incoming ACH transactions | 

## Methods

### NewIncomingAchList

`func NewIncomingAchList(transactions []IncomingAch, ) *IncomingAchList`

NewIncomingAchList instantiates a new IncomingAchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingAchListWithDefaults

`func NewIncomingAchListWithDefaults() *IncomingAchList`

NewIncomingAchListWithDefaults instantiates a new IncomingAchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *IncomingAchList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *IncomingAchList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *IncomingAchList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *IncomingAchList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTransactions

`func (o *IncomingAchList) GetTransactions() []IncomingAch`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *IncomingAchList) GetTransactionsOk() (*[]IncomingAch, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *IncomingAchList) SetTransactions(v []IncomingAch)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


