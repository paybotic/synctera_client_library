# IncomingWireList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Wires** | [**[]IncomingWire**](IncomingWire.md) | Array of incoming wires | 

## Methods

### NewIncomingWireList

`func NewIncomingWireList(wires []IncomingWire, ) *IncomingWireList`

NewIncomingWireList instantiates a new IncomingWireList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingWireListWithDefaults

`func NewIncomingWireListWithDefaults() *IncomingWireList`

NewIncomingWireListWithDefaults instantiates a new IncomingWireList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *IncomingWireList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *IncomingWireList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *IncomingWireList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *IncomingWireList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetWires

`func (o *IncomingWireList) GetWires() []IncomingWire`

GetWires returns the Wires field if non-nil, zero value otherwise.

### GetWiresOk

`func (o *IncomingWireList) GetWiresOk() (*[]IncomingWire, bool)`

GetWiresOk returns a tuple with the Wires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWires

`func (o *IncomingWireList) SetWires(v []IncomingWire)`

SetWires sets Wires field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


