# CrrList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**RiskScores** | [**[]CrrResponse**](CrrResponse.md) | Array of risk scores | 

## Methods

### NewCrrList

`func NewCrrList(riskScores []CrrResponse, ) *CrrList`

NewCrrList instantiates a new CrrList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrrListWithDefaults

`func NewCrrListWithDefaults() *CrrList`

NewCrrListWithDefaults instantiates a new CrrList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *CrrList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CrrList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CrrList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CrrList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetRiskScores

`func (o *CrrList) GetRiskScores() []CrrResponse`

GetRiskScores returns the RiskScores field if non-nil, zero value otherwise.

### GetRiskScoresOk

`func (o *CrrList) GetRiskScoresOk() (*[]CrrResponse, bool)`

GetRiskScoresOk returns a tuple with the RiskScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScores

`func (o *CrrList) SetRiskScores(v []CrrResponse)`

SetRiskScores sets RiskScores field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


