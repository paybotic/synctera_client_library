# AccountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Accounts** | [**[]AccountGenericResponse**](AccountGenericResponse.md) | Array of Accounts | 

## Methods

### NewAccountList

`func NewAccountList(accounts []AccountGenericResponse, ) *AccountList`

NewAccountList instantiates a new AccountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountListWithDefaults

`func NewAccountListWithDefaults() *AccountList`

NewAccountListWithDefaults instantiates a new AccountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *AccountList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AccountList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AccountList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AccountList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetAccounts

`func (o *AccountList) GetAccounts() []AccountGenericResponse`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountList) GetAccountsOk() (*[]AccountGenericResponse, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountList) SetAccounts(v []AccountGenericResponse)`

SetAccounts sets Accounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


