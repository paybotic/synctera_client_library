/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the WireList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireList{}

// WireList struct for WireList
type WireList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of wires
	Wires                []Wire `json:"wires"`
	AdditionalProperties map[string]interface{}
}

type _WireList WireList

// NewWireList instantiates a new WireList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireList(wires []Wire) *WireList {
	this := WireList{}
	this.Wires = wires
	return &this
}

// NewWireListWithDefaults instantiates a new WireList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireListWithDefaults() *WireList {
	this := WireList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *WireList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *WireList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *WireList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetWires returns the Wires field value
func (o *WireList) GetWires() []Wire {
	if o == nil {
		var ret []Wire
		return ret
	}

	return o.Wires
}

// GetWiresOk returns a tuple with the Wires field value
// and a boolean to check if the value has been set.
func (o *WireList) GetWiresOk() ([]Wire, bool) {
	if o == nil {
		return nil, false
	}
	return o.Wires, true
}

// SetWires sets field value
func (o *WireList) SetWires(v []Wire) {
	o.Wires = v
}

func (o WireList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WireList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["wires"] = o.Wires

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WireList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wires",
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

	varWireList := _WireList{}

	err = json.Unmarshal(data, &varWireList)

	if err != nil {
		return err
	}

	*o = WireList(varWireList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_token")
		delete(additionalProperties, "wires")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWireList struct {
	value *WireList
	isSet bool
}

func (v NullableWireList) Get() *WireList {
	return v.value
}

func (v *NullableWireList) Set(val *WireList) {
	v.value = val
	v.isSet = true
}

func (v NullableWireList) IsSet() bool {
	return v.isSet
}

func (v *NullableWireList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireList(val *WireList) *NullableWireList {
	return &NullableWireList{value: val, isSet: true}
}

func (v NullableWireList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
