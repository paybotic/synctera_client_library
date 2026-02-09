# PersonList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Persons** | [**[]ResponsePerson**](ResponsePerson.md) | Array of persons. | 

## Methods

### NewPersonList

`func NewPersonList(persons []ResponsePerson, ) *PersonList`

NewPersonList instantiates a new PersonList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonListWithDefaults

`func NewPersonListWithDefaults() *PersonList`

NewPersonListWithDefaults instantiates a new PersonList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *PersonList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *PersonList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *PersonList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *PersonList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPersons

`func (o *PersonList) GetPersons() []ResponsePerson`

GetPersons returns the Persons field if non-nil, zero value otherwise.

### GetPersonsOk

`func (o *PersonList) GetPersonsOk() (*[]ResponsePerson, bool)`

GetPersonsOk returns a tuple with the Persons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersons

`func (o *PersonList) SetPersons(v []ResponsePerson)`

SetPersons sets Persons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


