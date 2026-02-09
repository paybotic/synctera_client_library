# CreditScoreList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**CreditScores** | [**[]CreditScoreResponse**](CreditScoreResponse.md) |  | 

## Methods

### NewCreditScoreList

`func NewCreditScoreList(creditScores []CreditScoreResponse, ) *CreditScoreList`

NewCreditScoreList instantiates a new CreditScoreList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditScoreListWithDefaults

`func NewCreditScoreListWithDefaults() *CreditScoreList`

NewCreditScoreListWithDefaults instantiates a new CreditScoreList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *CreditScoreList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CreditScoreList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CreditScoreList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CreditScoreList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetCreditScores

`func (o *CreditScoreList) GetCreditScores() []CreditScoreResponse`

GetCreditScores returns the CreditScores field if non-nil, zero value otherwise.

### GetCreditScoresOk

`func (o *CreditScoreList) GetCreditScoresOk() (*[]CreditScoreResponse, bool)`

GetCreditScoresOk returns a tuple with the CreditScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditScores

`func (o *CreditScoreList) SetCreditScores(v []CreditScoreResponse)`

SetCreditScores sets CreditScores field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


