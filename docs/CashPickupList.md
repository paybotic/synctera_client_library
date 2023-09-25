# CashPickupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**CashPickups** | [**[]CashPickup**](CashPickup.md) | Array of cash pickups | 

## Methods

### NewCashPickupList

`func NewCashPickupList(cashPickups []CashPickup, ) *CashPickupList`

NewCashPickupList instantiates a new CashPickupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashPickupListWithDefaults

`func NewCashPickupListWithDefaults() *CashPickupList`

NewCashPickupListWithDefaults instantiates a new CashPickupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *CashPickupList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CashPickupList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CashPickupList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CashPickupList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetCashPickups

`func (o *CashPickupList) GetCashPickups() []CashPickup`

GetCashPickups returns the CashPickups field if non-nil, zero value otherwise.

### GetCashPickupsOk

`func (o *CashPickupList) GetCashPickupsOk() (*[]CashPickup, bool)`

GetCashPickupsOk returns a tuple with the CashPickups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashPickups

`func (o *CashPickupList) SetCashPickups(v []CashPickup)`

SetCashPickups sets CashPickups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


