# AutopayList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Autopays** | [**[]Autopay**](Autopay.md) | Array of autopays | 

## Methods

### NewAutopayList

`func NewAutopayList(autopays []Autopay, ) *AutopayList`

NewAutopayList instantiates a new AutopayList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayListWithDefaults

`func NewAutopayListWithDefaults() *AutopayList`

NewAutopayListWithDefaults instantiates a new AutopayList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *AutopayList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AutopayList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AutopayList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AutopayList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetAutopays

`func (o *AutopayList) GetAutopays() []Autopay`

GetAutopays returns the Autopays field if non-nil, zero value otherwise.

### GetAutopaysOk

`func (o *AutopayList) GetAutopaysOk() (*[]Autopay, bool)`

GetAutopaysOk returns a tuple with the Autopays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopays

`func (o *AutopayList) SetAutopays(v []Autopay)`

SetAutopays sets Autopays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


