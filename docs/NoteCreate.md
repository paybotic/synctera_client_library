# NoteCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The note&#39;s text content. | 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**RelatedResourceField** | Pointer to **string** | 🚧 Beta This is a Beta property. Feedback from the community is welcome. We may make breaking changes to this property. Path to the field in the related resource that the note pertains to. This uses a dot notation like the following: Examples: * a field in the resource: first_name * a sub-field: legal_address.city * nested arrays: application_details.sections[1].pages[2].items[0].answer  | [optional] 
**RelatedResourceId** | **string** | The id of the resource that is associated with the note. This is typically a UUID. For TENANT it is a string tenant ID.  | 
**RelatedResourceType** | [**RelatedResourceType2**](RelatedResourceType2.md) |  | 
**Status** | Pointer to [**NoteStatus**](NoteStatus.md) |  | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**Type** | Pointer to [**ModelType**](ModelType.md) |  | [optional] [default to MODELTYPE_NOTE]

## Methods

### NewNoteCreate

`func NewNoteCreate(content string, relatedResourceId string, relatedResourceType RelatedResourceType2, ) *NoteCreate`

NewNoteCreate instantiates a new NoteCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteCreateWithDefaults

`func NewNoteCreateWithDefaults() *NoteCreate`

NewNoteCreateWithDefaults instantiates a new NoteCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *NoteCreate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteCreate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteCreate) SetContent(v string)`

SetContent sets Content field to given value.


### GetMetadata

`func (o *NoteCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NoteCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NoteCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NoteCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelatedResourceField

`func (o *NoteCreate) GetRelatedResourceField() string`

GetRelatedResourceField returns the RelatedResourceField field if non-nil, zero value otherwise.

### GetRelatedResourceFieldOk

`func (o *NoteCreate) GetRelatedResourceFieldOk() (*string, bool)`

GetRelatedResourceFieldOk returns a tuple with the RelatedResourceField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceField

`func (o *NoteCreate) SetRelatedResourceField(v string)`

SetRelatedResourceField sets RelatedResourceField field to given value.

### HasRelatedResourceField

`func (o *NoteCreate) HasRelatedResourceField() bool`

HasRelatedResourceField returns a boolean if a field has been set.

### GetRelatedResourceId

`func (o *NoteCreate) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *NoteCreate) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *NoteCreate) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *NoteCreate) GetRelatedResourceType() RelatedResourceType2`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *NoteCreate) GetRelatedResourceTypeOk() (*RelatedResourceType2, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *NoteCreate) SetRelatedResourceType(v RelatedResourceType2)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetStatus

`func (o *NoteCreate) GetStatus() NoteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NoteCreate) GetStatusOk() (*NoteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NoteCreate) SetStatus(v NoteStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NoteCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *NoteCreate) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *NoteCreate) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *NoteCreate) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *NoteCreate) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *NoteCreate) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoteCreate) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoteCreate) SetType(v ModelType)`

SetType sets Type field to given value.

### HasType

`func (o *NoteCreate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


