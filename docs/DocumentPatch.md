# DocumentPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletionReason** | Pointer to **string** | An explanation why the file was deleted. You must set a document&#39;s deletion_reason before deleting it. | [optional] 
**Description** | Pointer to **string** | A description of the document | [optional] 
**Name** | Pointer to **string** | A user-friendly name for the document | [optional] 
**Type** | Pointer to [**DocumentType**](DocumentType.md) |  | [optional] 

## Methods

### NewDocumentPatch

`func NewDocumentPatch() *DocumentPatch`

NewDocumentPatch instantiates a new DocumentPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentPatchWithDefaults

`func NewDocumentPatchWithDefaults() *DocumentPatch`

NewDocumentPatchWithDefaults instantiates a new DocumentPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletionReason

`func (o *DocumentPatch) GetDeletionReason() string`

GetDeletionReason returns the DeletionReason field if non-nil, zero value otherwise.

### GetDeletionReasonOk

`func (o *DocumentPatch) GetDeletionReasonOk() (*string, bool)`

GetDeletionReasonOk returns a tuple with the DeletionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionReason

`func (o *DocumentPatch) SetDeletionReason(v string)`

SetDeletionReason sets DeletionReason field to given value.

### HasDeletionReason

`func (o *DocumentPatch) HasDeletionReason() bool`

HasDeletionReason returns a boolean if a field has been set.

### GetDescription

`func (o *DocumentPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DocumentPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DocumentPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DocumentPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DocumentPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DocumentPatch) GetType() DocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentPatch) GetTypeOk() (*DocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentPatch) SetType(v DocumentType)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentPatch) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


