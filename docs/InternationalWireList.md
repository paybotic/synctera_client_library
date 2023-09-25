# InternationalWireList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**InternationalWires** | [**[]InternationalWireResponse**](InternationalWireResponse.md) | Array of international wires | 

## Methods

### NewInternationalWireList

`func NewInternationalWireList(internationalWires []InternationalWireResponse, ) *InternationalWireList`

NewInternationalWireList instantiates a new InternationalWireList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalWireListWithDefaults

`func NewInternationalWireListWithDefaults() *InternationalWireList`

NewInternationalWireListWithDefaults instantiates a new InternationalWireList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *InternationalWireList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *InternationalWireList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *InternationalWireList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *InternationalWireList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetInternationalWires

`func (o *InternationalWireList) GetInternationalWires() []InternationalWireResponse`

GetInternationalWires returns the InternationalWires field if non-nil, zero value otherwise.

### GetInternationalWiresOk

`func (o *InternationalWireList) GetInternationalWiresOk() (*[]InternationalWireResponse, bool)`

GetInternationalWiresOk returns a tuple with the InternationalWires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalWires

`func (o *InternationalWireList) SetInternationalWires(v []InternationalWireResponse)`

SetInternationalWires sets InternationalWires field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


