# TemplateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**AccountTemplates** | [**[]AccountTemplateResponse**](AccountTemplateResponse.md) | Array of account templates | 

## Methods

### NewTemplateList

`func NewTemplateList(accountTemplates []AccountTemplateResponse, ) *TemplateList`

NewTemplateList instantiates a new TemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateListWithDefaults

`func NewTemplateListWithDefaults() *TemplateList`

NewTemplateListWithDefaults instantiates a new TemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *TemplateList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *TemplateList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *TemplateList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *TemplateList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetAccountTemplates

`func (o *TemplateList) GetAccountTemplates() []AccountTemplateResponse`

GetAccountTemplates returns the AccountTemplates field if non-nil, zero value otherwise.

### GetAccountTemplatesOk

`func (o *TemplateList) GetAccountTemplatesOk() (*[]AccountTemplateResponse, bool)`

GetAccountTemplatesOk returns a tuple with the AccountTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplates

`func (o *TemplateList) SetAccountTemplates(v []AccountTemplateResponse)`

SetAccountTemplates sets AccountTemplates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


