# DocumentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Documents** | [**[]DocumentResponse**](DocumentResponse.md) | Array of documents | 

## Methods

### NewDocumentList

`func NewDocumentList(documents []DocumentResponse, ) *DocumentList`

NewDocumentList instantiates a new DocumentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentListWithDefaults

`func NewDocumentListWithDefaults() *DocumentList`

NewDocumentListWithDefaults instantiates a new DocumentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *DocumentList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *DocumentList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *DocumentList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *DocumentList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetDocuments

`func (o *DocumentList) GetDocuments() []DocumentResponse`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DocumentList) GetDocumentsOk() (*[]DocumentResponse, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DocumentList) SetDocuments(v []DocumentResponse)`

SetDocuments sets Documents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


