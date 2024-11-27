/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the IncomingWireList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomingWireList{}

// IncomingWireList struct for IncomingWireList
type IncomingWireList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of incoming wires
	Wires []IncomingWire `json:"wires"`
}

type _IncomingWireList IncomingWireList

// NewIncomingWireList instantiates a new IncomingWireList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomingWireList(wires []IncomingWire) *IncomingWireList {
	this := IncomingWireList{}
	this.Wires = wires
	return &this
}

// NewIncomingWireListWithDefaults instantiates a new IncomingWireList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomingWireListWithDefaults() *IncomingWireList {
	this := IncomingWireList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *IncomingWireList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomingWireList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *IncomingWireList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *IncomingWireList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetWires returns the Wires field value
func (o *IncomingWireList) GetWires() []IncomingWire {
	if o == nil {
		var ret []IncomingWire
		return ret
	}

	return o.Wires
}

// GetWiresOk returns a tuple with the Wires field value
// and a boolean to check if the value has been set.
func (o *IncomingWireList) GetWiresOk() ([]IncomingWire, bool) {
	if o == nil {
		return nil, false
	}
	return o.Wires, true
}

// SetWires sets field value
func (o *IncomingWireList) SetWires(v []IncomingWire) {
	o.Wires = v
}

func (o IncomingWireList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomingWireList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["wires"] = o.Wires
	return toSerialize, nil
}

func (o *IncomingWireList) UnmarshalJSON(data []byte) (err error) {
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

	varIncomingWireList := _IncomingWireList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIncomingWireList)

	if err != nil {
		return err
	}

	*o = IncomingWireList(varIncomingWireList)

	return err
}

type NullableIncomingWireList struct {
	value *IncomingWireList
	isSet bool
}

func (v NullableIncomingWireList) Get() *IncomingWireList {
	return v.value
}

func (v *NullableIncomingWireList) Set(val *IncomingWireList) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomingWireList) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomingWireList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomingWireList(val *IncomingWireList) *NullableIncomingWireList {
	return &NullableIncomingWireList{value: val, isSet: true}
}

func (v NullableIncomingWireList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomingWireList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}