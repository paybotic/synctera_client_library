# SpendControlResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**SpendControls** | [**[]SpendControlResponse**](SpendControlResponse.md) | Array of Spend Controls | 

## Methods

### NewSpendControlResponseList

`func NewSpendControlResponseList(spendControls []SpendControlResponse, ) *SpendControlResponseList`

NewSpendControlResponseList instantiates a new SpendControlResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlResponseListWithDefaults

`func NewSpendControlResponseListWithDefaults() *SpendControlResponseList`

NewSpendControlResponseListWithDefaults instantiates a new SpendControlResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *SpendControlResponseList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *SpendControlResponseList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *SpendControlResponseList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *SpendControlResponseList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetSpendControls

`func (o *SpendControlResponseList) GetSpendControls() []SpendControlResponse`

GetSpendControls returns the SpendControls field if non-nil, zero value otherwise.

### GetSpendControlsOk

`func (o *SpendControlResponseList) GetSpendControlsOk() (*[]SpendControlResponse, bool)`

GetSpendControlsOk returns a tuple with the SpendControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendControls

`func (o *SpendControlResponseList) SetSpendControls(v []SpendControlResponse)`

SetSpendControls sets SpendControls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


