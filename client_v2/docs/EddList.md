# EddList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**EddReports** | [**[]CreateEddResponse**](CreateEddResponse.md) | Array of edd reports. | 

## Methods

### NewEddList

`func NewEddList(eddReports []CreateEddResponse, ) *EddList`

NewEddList instantiates a new EddList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEddListWithDefaults

`func NewEddListWithDefaults() *EddList`

NewEddListWithDefaults instantiates a new EddList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *EddList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *EddList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *EddList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *EddList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetEddReports

`func (o *EddList) GetEddReports() []CreateEddResponse`

GetEddReports returns the EddReports field if non-nil, zero value otherwise.

### GetEddReportsOk

`func (o *EddList) GetEddReportsOk() (*[]CreateEddResponse, bool)`

GetEddReportsOk returns a tuple with the EddReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEddReports

`func (o *EddList) SetEddReports(v []CreateEddResponse)`

SetEddReports sets EddReports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


