# PatchNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NoteStatus**](NoteStatus.md) |  | [optional] 

## Methods

### NewPatchNote

`func NewPatchNote() *PatchNote`

NewPatchNote instantiates a new PatchNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchNoteWithDefaults

`func NewPatchNoteWithDefaults() *PatchNote`

NewPatchNoteWithDefaults instantiates a new PatchNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PatchNote) GetStatus() NoteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchNote) GetStatusOk() (*NoteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchNote) SetStatus(v NoteStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchNote) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


