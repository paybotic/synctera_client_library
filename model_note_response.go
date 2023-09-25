/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the NoteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoteResponse{}

// NoteResponse struct for NoteResponse
type NoteResponse struct {
	// The note's author.
	Author string `json:"author"`
	// The note's text content.
	Content string `json:"content"`
	// The date and time the note was created.
	CreationTime time.Time `json:"creation_time"`
	// note ID
	Id string `json:"id"`
	// The date and time the note was last updated.
	LastUpdatedTime time.Time `json:"last_updated_time"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// 🚧 Beta This is a Beta property. Feedback from the community is welcome. We may make breaking changes to this property. Path to the field in the related resource that the note pertains to. This uses a dot notation like the following: Examples: * a field in the resource: first_name * a sub-field: legal_address.city * nested arrays: application_details.sections[1].pages[2].items[0].answer 
	RelatedResourceField *string `json:"related_resource_field,omitempty"`
	// The id of the resource that is associated with the note. This is typically a UUID. For TENANT it is a string tenant ID. 
	RelatedResourceId string `json:"related_resource_id"`
	RelatedResourceType RelatedResourceType1 `json:"related_resource_type"`
	Status *Status1 `json:"status,omitempty"`
	// The tenant containing the resource. Tenancy is represented as bank_id_partner_id. This attribute is included on all responses. For requests, it is optional for clients with access to a single tenant. 
	Tenant string `json:"tenant"`
	Type *ModelType `json:"type,omitempty"`
}

// NewNoteResponse instantiates a new NoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteResponse(author string, content string, creationTime time.Time, id string, lastUpdatedTime time.Time, relatedResourceId string, relatedResourceType RelatedResourceType1, tenant string) *NoteResponse {
	this := NoteResponse{}
	this.Author = author
	this.Content = content
	this.CreationTime = creationTime
	this.Id = id
	this.LastUpdatedTime = lastUpdatedTime
	this.RelatedResourceId = relatedResourceId
	this.RelatedResourceType = relatedResourceType
	this.Tenant = tenant
	var type_ ModelType = MODELTYPE_NOTE
	this.Type = &type_
	return &this
}

// NewNoteResponseWithDefaults instantiates a new NoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteResponseWithDefaults() *NoteResponse {
	this := NoteResponse{}
	var type_ ModelType = MODELTYPE_NOTE
	this.Type = &type_
	return &this
}

// GetAuthor returns the Author field value
func (o *NoteResponse) GetAuthor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *NoteResponse) SetAuthor(v string) {
	o.Author = v
}

// GetContent returns the Content field value
func (o *NoteResponse) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *NoteResponse) SetContent(v string) {
	o.Content = v
}

// GetCreationTime returns the CreationTime field value
func (o *NoteResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *NoteResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetId returns the Id field value
func (o *NoteResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NoteResponse) SetId(v string) {
	o.Id = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *NoteResponse) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *NoteResponse) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NoteResponse) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NoteResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *NoteResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetRelatedResourceField returns the RelatedResourceField field value if set, zero value otherwise.
func (o *NoteResponse) GetRelatedResourceField() string {
	if o == nil || IsNil(o.RelatedResourceField) {
		var ret string
		return ret
	}
	return *o.RelatedResourceField
}

// GetRelatedResourceFieldOk returns a tuple with the RelatedResourceField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetRelatedResourceFieldOk() (*string, bool) {
	if o == nil || IsNil(o.RelatedResourceField) {
		return nil, false
	}
	return o.RelatedResourceField, true
}

// HasRelatedResourceField returns a boolean if a field has been set.
func (o *NoteResponse) HasRelatedResourceField() bool {
	if o != nil && !IsNil(o.RelatedResourceField) {
		return true
	}

	return false
}

// SetRelatedResourceField gets a reference to the given string and assigns it to the RelatedResourceField field.
func (o *NoteResponse) SetRelatedResourceField(v string) {
	o.RelatedResourceField = &v
}

// GetRelatedResourceId returns the RelatedResourceId field value
func (o *NoteResponse) GetRelatedResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelatedResourceId
}

// GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field value
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetRelatedResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceId, true
}

// SetRelatedResourceId sets field value
func (o *NoteResponse) SetRelatedResourceId(v string) {
	o.RelatedResourceId = v
}

// GetRelatedResourceType returns the RelatedResourceType field value
func (o *NoteResponse) GetRelatedResourceType() RelatedResourceType1 {
	if o == nil {
		var ret RelatedResourceType1
		return ret
	}

	return o.RelatedResourceType
}

// GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field value
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceType, true
}

// SetRelatedResourceType sets field value
func (o *NoteResponse) SetRelatedResourceType(v RelatedResourceType1) {
	o.RelatedResourceType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NoteResponse) GetStatus() Status1 {
	if o == nil || IsNil(o.Status) {
		var ret Status1
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetStatusOk() (*Status1, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NoteResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status1 and assigns it to the Status field.
func (o *NoteResponse) SetStatus(v Status1) {
	o.Status = &v
}

// GetTenant returns the Tenant field value
func (o *NoteResponse) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *NoteResponse) SetTenant(v string) {
	o.Tenant = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NoteResponse) GetType() ModelType {
	if o == nil || IsNil(o.Type) {
		var ret ModelType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteResponse) GetTypeOk() (*ModelType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NoteResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ModelType and assigns it to the Type field.
func (o *NoteResponse) SetType(v ModelType) {
	o.Type = &v
}

func (o NoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["author"] = o.Author
	toSerialize["content"] = o.Content
	toSerialize["creation_time"] = o.CreationTime
	toSerialize["id"] = o.Id
	toSerialize["last_updated_time"] = o.LastUpdatedTime
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.RelatedResourceField) {
		toSerialize["related_resource_field"] = o.RelatedResourceField
	}
	toSerialize["related_resource_id"] = o.RelatedResourceId
	toSerialize["related_resource_type"] = o.RelatedResourceType
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["tenant"] = o.Tenant
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableNoteResponse struct {
	value *NoteResponse
	isSet bool
}

func (v NullableNoteResponse) Get() *NoteResponse {
	return v.value
}

func (v *NullableNoteResponse) Set(val *NoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteResponse(val *NoteResponse) *NullableNoteResponse {
	return &NullableNoteResponse{value: val, isSet: true}
}

func (v NullableNoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


