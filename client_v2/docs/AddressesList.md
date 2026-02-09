# AddressesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Addresses** | [**[]AddressResponse**](AddressResponse.md) | Array of Addresses | 

## Methods

### NewAddressesList

`func NewAddressesList(addresses []AddressResponse, ) *AddressesList`

NewAddressesList instantiates a new AddressesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressesListWithDefaults

`func NewAddressesListWithDefaults() *AddressesList`

NewAddressesListWithDefaults instantiates a new AddressesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *AddressesList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AddressesList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AddressesList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AddressesList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetAddresses

`func (o *AddressesList) GetAddresses() []AddressResponse`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *AddressesList) GetAddressesOk() (*[]AddressResponse, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *AddressesList) SetAddresses(v []AddressResponse)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


