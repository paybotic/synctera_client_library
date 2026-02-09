# InstitutionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Institutions** | Pointer to [**[]Institution**](Institution.md) |  | [optional] 

## Methods

### NewInstitutionList

`func NewInstitutionList() *InstitutionList`

NewInstitutionList instantiates a new InstitutionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionListWithDefaults

`func NewInstitutionListWithDefaults() *InstitutionList`

NewInstitutionListWithDefaults instantiates a new InstitutionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *InstitutionList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *InstitutionList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *InstitutionList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *InstitutionList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetInstitutions

`func (o *InstitutionList) GetInstitutions() []Institution`

GetInstitutions returns the Institutions field if non-nil, zero value otherwise.

### GetInstitutionsOk

`func (o *InstitutionList) GetInstitutionsOk() (*[]Institution, bool)`

GetInstitutionsOk returns a tuple with the Institutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutions

`func (o *InstitutionList) SetInstitutions(v []Institution)`

SetInstitutions sets Institutions field to given value.

### HasInstitutions

`func (o *InstitutionList) HasInstitutions() bool`

HasInstitutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


