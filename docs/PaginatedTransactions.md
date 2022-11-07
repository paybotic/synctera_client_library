# PaginatedTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | **NullableString** | If returned, use the value of &#x60;next_page_token&#x60; in the &#x60;page_token&#x60; query parameter to query for the next page of results. This will be &#x60;null&#x60; if there are no more pages. | 
**Transactions** | [**[]Transaction1**](Transaction1.md) | List of transactions | 

## Methods

### NewPaginatedTransactions

`func NewPaginatedTransactions(nextPageToken NullableString, transactions []Transaction1, ) *PaginatedTransactions`

NewPaginatedTransactions instantiates a new PaginatedTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedTransactionsWithDefaults

`func NewPaginatedTransactionsWithDefaults() *PaginatedTransactions`

NewPaginatedTransactionsWithDefaults instantiates a new PaginatedTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *PaginatedTransactions) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *PaginatedTransactions) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *PaginatedTransactions) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.


### SetNextPageTokenNil

`func (o *PaginatedTransactions) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *PaginatedTransactions) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil
### GetTransactions

`func (o *PaginatedTransactions) GetTransactions() []Transaction1`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PaginatedTransactions) GetTransactionsOk() (*[]Transaction1, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PaginatedTransactions) SetTransactions(v []Transaction1)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


