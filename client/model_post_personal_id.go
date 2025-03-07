/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PostPersonalId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPersonalId{}

// PostPersonalId struct for PostPersonalId
type PostPersonalId struct {
	// UUID for the personal identifier for subsequent changes and deletion
	Id     *string        `json:"id,omitempty"`
	IdType PersonalIdType `json:"id_type"`
	// The personal identifier. Format varies by personal identifier type.
	Identifier string `json:"identifier"`
	// True if the identifier was provided by the system, e.g. via SSN Prefill.
	SystemProvided *bool `json:"system_provided,omitempty"`
	// The ISO 3166 Alpha-2 country code for the country that issued the personal identifier. This is optional for personal identifier types that have an implicit country, e.g. SSN. This is required for other types, e.g. PASSPORT_NUMBER.
	CountryCode *string `json:"country_code,omitempty"`
}

type _PostPersonalId PostPersonalId

// NewPostPersonalId instantiates a new PostPersonalId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPersonalId(idType PersonalIdType, identifier string) *PostPersonalId {
	this := PostPersonalId{}
	this.IdType = idType
	this.Identifier = identifier
	return &this
}

// NewPostPersonalIdWithDefaults instantiates a new PostPersonalId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPersonalIdWithDefaults() *PostPersonalId {
	this := PostPersonalId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostPersonalId) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPersonalId) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostPersonalId) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PostPersonalId) SetId(v string) {
	o.Id = &v
}

// GetIdType returns the IdType field value
func (o *PostPersonalId) GetIdType() PersonalIdType {
	if o == nil {
		var ret PersonalIdType
		return ret
	}

	return o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value
// and a boolean to check if the value has been set.
func (o *PostPersonalId) GetIdTypeOk() (*PersonalIdType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdType, true
}

// SetIdType sets field value
func (o *PostPersonalId) SetIdType(v PersonalIdType) {
	o.IdType = v
}

// GetIdentifier returns the Identifier field value
func (o *PostPersonalId) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *PostPersonalId) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *PostPersonalId) SetIdentifier(v string) {
	o.Identifier = v
}

// GetSystemProvided returns the SystemProvided field value if set, zero value otherwise.
func (o *PostPersonalId) GetSystemProvided() bool {
	if o == nil || IsNil(o.SystemProvided) {
		var ret bool
		return ret
	}
	return *o.SystemProvided
}

// GetSystemProvidedOk returns a tuple with the SystemProvided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPersonalId) GetSystemProvidedOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemProvided) {
		return nil, false
	}
	return o.SystemProvided, true
}

// HasSystemProvided returns a boolean if a field has been set.
func (o *PostPersonalId) HasSystemProvided() bool {
	if o != nil && !IsNil(o.SystemProvided) {
		return true
	}

	return false
}

// SetSystemProvided gets a reference to the given bool and assigns it to the SystemProvided field.
func (o *PostPersonalId) SetSystemProvided(v bool) {
	o.SystemProvided = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PostPersonalId) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPersonalId) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PostPersonalId) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PostPersonalId) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o PostPersonalId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPersonalId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["id_type"] = o.IdType
	toSerialize["identifier"] = o.Identifier
	if !IsNil(o.SystemProvided) {
		toSerialize["system_provided"] = o.SystemProvided
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	return toSerialize, nil
}

func (o *PostPersonalId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id_type",
		"identifier",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostPersonalId := _PostPersonalId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostPersonalId)

	if err != nil {
		return err
	}

	*o = PostPersonalId(varPostPersonalId)

	return err
}

type NullablePostPersonalId struct {
	value *PostPersonalId
	isSet bool
}

func (v NullablePostPersonalId) Get() *PostPersonalId {
	return v.value
}

func (v *NullablePostPersonalId) Set(val *PostPersonalId) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPersonalId) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPersonalId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPersonalId(val *PostPersonalId) *NullablePostPersonalId {
	return &NullablePostPersonalId{value: val, isSet: true}
}

func (v NullablePostPersonalId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPersonalId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
