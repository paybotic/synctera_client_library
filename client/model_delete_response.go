/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// checks if the DeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteResponse{}

// DeleteResponse Deleted object information
type DeleteResponse struct {
	// Object ID
	Id *string `json:"id,omitempty"`
	// The resource name
	Resource             *string `json:"resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteResponse DeleteResponse

// NewDeleteResponse instantiates a new DeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteResponse() *DeleteResponse {
	this := DeleteResponse{}
	return &this
}

// NewDeleteResponseWithDefaults instantiates a new DeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteResponseWithDefaults() *DeleteResponse {
	this := DeleteResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeleteResponse) SetId(v string) {
	o.Id = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *DeleteResponse) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteResponse) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *DeleteResponse) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *DeleteResponse) SetResource(v string) {
	o.Resource = &v
}

func (o DeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteResponse) UnmarshalJSON(data []byte) (err error) {
	varDeleteResponse := _DeleteResponse{}

	err = json.Unmarshal(data, &varDeleteResponse)

	if err != nil {
		return err
	}

	*o = DeleteResponse(varDeleteResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "resource")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteResponse struct {
	value *DeleteResponse
	isSet bool
}

func (v NullableDeleteResponse) Get() *DeleteResponse {
	return v.value
}

func (v *NullableDeleteResponse) Set(val *DeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteResponse(val *DeleteResponse) *NullableDeleteResponse {
	return &NullableDeleteResponse{value: val, isSet: true}
}

func (v NullableDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
