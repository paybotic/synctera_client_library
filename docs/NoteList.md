# NoteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | [**[]NoteResponse**](NoteResponse.md) | Array of notes | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewNoteList

`func NewNoteList(notes []NoteResponse, ) *NoteList`

NewNoteList instantiates a new NoteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteListWithDefaults

`func NewNoteListWithDefaults() *NoteList`

NewNoteListWithDefaults instantiates a new NoteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *NoteList) GetNotes() []NoteResponse`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *NoteList) GetNotesOk() (*[]NoteResponse, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *NoteList) SetNotes(v []NoteResponse)`

SetNotes sets Notes field to given value.


### GetNextPageToken

`func (o *NoteList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *NoteList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *NoteList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *NoteList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


