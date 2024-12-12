# NoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | **string** | The note&#39;s author. | 
**Content** | **string** | The note&#39;s text content. | 
**CreationTime** | **time.Time** | The date and time the note was created. | [readonly] 
**Id** | **string** | note ID | [readonly] 
**LastUpdatedTime** | **time.Time** | The date and time the note was last updated. | [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**RelatedResourceField** | Pointer to **string** | 🚧 Beta This is a Beta property. Feedback from the community is welcome. We may make breaking changes to this property. Path to the field in the related resource that the note pertains to. This uses a dot notation like the following: Examples: * a field in the resource: first_name * a sub-field: legal_address.city * nested arrays: application_details.sections[1].pages[2].items[0].answer  | [optional] 
**RelatedResourceId** | **string** | The id of the resource that is associated with the note. This is typically a UUID. For TENANT it is a string tenant ID.  | 
**RelatedResourceType** | [**RelatedResourceType2**](RelatedResourceType2.md) |  | 
**Status** | Pointer to [**NoteStatus**](NoteStatus.md) |  | [optional] 
**Tenant** | **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | 
**Type** | Pointer to [**ModelType**](ModelType.md) |  | [optional] [default to MODELTYPE_NOTE]

## Methods

### NewNoteResponse

`func NewNoteResponse(author string, content string, creationTime time.Time, id string, lastUpdatedTime time.Time, relatedResourceId string, relatedResourceType RelatedResourceType2, tenant string, ) *NoteResponse`

NewNoteResponse instantiates a new NoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteResponseWithDefaults

`func NewNoteResponseWithDefaults() *NoteResponse`

NewNoteResponseWithDefaults instantiates a new NoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *NoteResponse) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NoteResponse) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NoteResponse) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetContent

`func (o *NoteResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NoteResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NoteResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreationTime

`func (o *NoteResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *NoteResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *NoteResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *NoteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *NoteResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *NoteResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *NoteResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMetadata

`func (o *NoteResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NoteResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NoteResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NoteResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelatedResourceField

`func (o *NoteResponse) GetRelatedResourceField() string`

GetRelatedResourceField returns the RelatedResourceField field if non-nil, zero value otherwise.

### GetRelatedResourceFieldOk

`func (o *NoteResponse) GetRelatedResourceFieldOk() (*string, bool)`

GetRelatedResourceFieldOk returns a tuple with the RelatedResourceField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceField

`func (o *NoteResponse) SetRelatedResourceField(v string)`

SetRelatedResourceField sets RelatedResourceField field to given value.

### HasRelatedResourceField

`func (o *NoteResponse) HasRelatedResourceField() bool`

HasRelatedResourceField returns a boolean if a field has been set.

### GetRelatedResourceId

`func (o *NoteResponse) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *NoteResponse) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *NoteResponse) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *NoteResponse) GetRelatedResourceType() RelatedResourceType2`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *NoteResponse) GetRelatedResourceTypeOk() (*RelatedResourceType2, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *NoteResponse) SetRelatedResourceType(v RelatedResourceType2)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetStatus

`func (o *NoteResponse) GetStatus() NoteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NoteResponse) GetStatusOk() (*NoteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NoteResponse) SetStatus(v NoteStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NoteResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *NoteResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *NoteResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *NoteResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetType

`func (o *NoteResponse) GetType() ModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NoteResponse) GetTypeOk() (*ModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NoteResponse) SetType(v ModelType)`

SetType sets Type field to given value.

### HasType

`func (o *NoteResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


