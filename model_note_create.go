/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NoteCreate struct for NoteCreate
type NoteCreate struct {
	// The note's text content.
	Content string `json:"content"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The id of the resource that is associated with the note. This is typically a UUID. For TENANT it is a string tenant ID. 
	RelatedResourceId string `json:"related_resource_id"`
	RelatedResourceType RelatedResourceType1 `json:"related_resource_type"`
	// The tenant containing the resource. Tenancy is represented as bank_id_partner_id. This attribute is included on all responses. For requests, it is optional for clients with access to a single tenant. 
	Tenant *string `json:"tenant,omitempty"`
}

// NewNoteCreate instantiates a new NoteCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteCreate(content string, relatedResourceId string, relatedResourceType RelatedResourceType1) *NoteCreate {
	this := NoteCreate{}
	this.Content = content
	this.RelatedResourceId = relatedResourceId
	this.RelatedResourceType = relatedResourceType
	return &this
}

// NewNoteCreateWithDefaults instantiates a new NoteCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteCreateWithDefaults() *NoteCreate {
	this := NoteCreate{}
	return &this
}

// GetContent returns the Content field value
func (o *NoteCreate) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *NoteCreate) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *NoteCreate) SetContent(v string) {
	o.Content = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NoteCreate) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteCreate) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NoteCreate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *NoteCreate) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetRelatedResourceId returns the RelatedResourceId field value
func (o *NoteCreate) GetRelatedResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelatedResourceId
}

// GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field value
// and a boolean to check if the value has been set.
func (o *NoteCreate) GetRelatedResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceId, true
}

// SetRelatedResourceId sets field value
func (o *NoteCreate) SetRelatedResourceId(v string) {
	o.RelatedResourceId = v
}

// GetRelatedResourceType returns the RelatedResourceType field value
func (o *NoteCreate) GetRelatedResourceType() RelatedResourceType1 {
	if o == nil {
		var ret RelatedResourceType1
		return ret
	}

	return o.RelatedResourceType
}

// GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field value
// and a boolean to check if the value has been set.
func (o *NoteCreate) GetRelatedResourceTypeOk() (*RelatedResourceType1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResourceType, true
}

// SetRelatedResourceType sets field value
func (o *NoteCreate) SetRelatedResourceType(v RelatedResourceType1) {
	o.RelatedResourceType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *NoteCreate) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteCreate) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *NoteCreate) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *NoteCreate) SetTenant(v string) {
	o.Tenant = &v
}

func (o NoteCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["related_resource_id"] = o.RelatedResourceId
	}
	if true {
		toSerialize["related_resource_type"] = o.RelatedResourceType
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullableNoteCreate struct {
	value *NoteCreate
	isSet bool
}

func (v NullableNoteCreate) Get() *NoteCreate {
	return v.value
}

func (v *NullableNoteCreate) Set(val *NoteCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteCreate(val *NoteCreate) *NullableNoteCreate {
	return &NullableNoteCreate{value: val, isSet: true}
}

func (v NullableNoteCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


