# FeeTemplateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**FeeTemplates** | [**[]FeeTemplateResponse**](FeeTemplateResponse.md) | Array of fee templates | 

## Methods

### NewFeeTemplateList

`func NewFeeTemplateList(feeTemplates []FeeTemplateResponse, ) *FeeTemplateList`

NewFeeTemplateList instantiates a new FeeTemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeTemplateListWithDefaults

`func NewFeeTemplateListWithDefaults() *FeeTemplateList`

NewFeeTemplateListWithDefaults instantiates a new FeeTemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *FeeTemplateList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FeeTemplateList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FeeTemplateList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FeeTemplateList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetFeeTemplates

`func (o *FeeTemplateList) GetFeeTemplates() []FeeTemplateResponse`

GetFeeTemplates returns the FeeTemplates field if non-nil, zero value otherwise.

### GetFeeTemplatesOk

`func (o *FeeTemplateList) GetFeeTemplatesOk() (*[]FeeTemplateResponse, bool)`

GetFeeTemplatesOk returns a tuple with the FeeTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTemplates

`func (o *FeeTemplateList) SetFeeTemplates(v []FeeTemplateResponse)`

SetFeeTemplates sets FeeTemplates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


