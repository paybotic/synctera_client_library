# CashAuthorizationResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Authorizations** | [**[]CashAuthorizationResponse**](CashAuthorizationResponse.md) | Array of authorizations. | 

## Methods

### NewCashAuthorizationResponseList

`func NewCashAuthorizationResponseList(authorizations []CashAuthorizationResponse, ) *CashAuthorizationResponseList`

NewCashAuthorizationResponseList instantiates a new CashAuthorizationResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashAuthorizationResponseListWithDefaults

`func NewCashAuthorizationResponseListWithDefaults() *CashAuthorizationResponseList`

NewCashAuthorizationResponseListWithDefaults instantiates a new CashAuthorizationResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *CashAuthorizationResponseList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CashAuthorizationResponseList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CashAuthorizationResponseList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CashAuthorizationResponseList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetAuthorizations

`func (o *CashAuthorizationResponseList) GetAuthorizations() []CashAuthorizationResponse`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *CashAuthorizationResponseList) GetAuthorizationsOk() (*[]CashAuthorizationResponse, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *CashAuthorizationResponseList) SetAuthorizations(v []CashAuthorizationResponse)`

SetAuthorizations sets Authorizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


