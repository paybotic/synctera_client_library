# StopPaymentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**StopPayments** | [**[]StopPaymentResponseWAccount**](StopPaymentResponseWAccount.md) |  | 

## Methods

### NewStopPaymentList

`func NewStopPaymentList(stopPayments []StopPaymentResponseWAccount, ) *StopPaymentList`

NewStopPaymentList instantiates a new StopPaymentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopPaymentListWithDefaults

`func NewStopPaymentListWithDefaults() *StopPaymentList`

NewStopPaymentListWithDefaults instantiates a new StopPaymentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *StopPaymentList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *StopPaymentList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *StopPaymentList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *StopPaymentList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetStopPayments

`func (o *StopPaymentList) GetStopPayments() []StopPaymentResponseWAccount`

GetStopPayments returns the StopPayments field if non-nil, zero value otherwise.

### GetStopPaymentsOk

`func (o *StopPaymentList) GetStopPaymentsOk() (*[]StopPaymentResponseWAccount, bool)`

GetStopPaymentsOk returns a tuple with the StopPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPayments

`func (o *StopPaymentList) SetStopPayments(v []StopPaymentResponseWAccount)`

SetStopPayments sets StopPayments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


