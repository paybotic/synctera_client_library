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

// DishonorAch Dishonored return
type DishonorAch struct {
	FieldErrors []string `json:"field_errors,omitempty"`
	Type        string   `json:"type"`
}

// NewDishonorAch instantiates a new DishonorAch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDishonorAch(type_ string) *DishonorAch {
	this := DishonorAch{}
	this.Type = type_
	return &this
}

// NewDishonorAchWithDefaults instantiates a new DishonorAch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDishonorAchWithDefaults() *DishonorAch {
	this := DishonorAch{}
	return &this
}

// GetFieldErrors returns the FieldErrors field value if set, zero value otherwise.
func (o *DishonorAch) GetFieldErrors() []string {
	if o == nil || o.FieldErrors == nil {
		var ret []string
		return ret
	}
	return o.FieldErrors
}

// GetFieldErrorsOk returns a tuple with the FieldErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DishonorAch) GetFieldErrorsOk() ([]string, bool) {
	if o == nil || o.FieldErrors == nil {
		return nil, false
	}
	return o.FieldErrors, true
}

// HasFieldErrors returns a boolean if a field has been set.
func (o *DishonorAch) HasFieldErrors() bool {
	if o != nil && o.FieldErrors != nil {
		return true
	}

	return false
}

// SetFieldErrors gets a reference to the given []string and assigns it to the FieldErrors field.
func (o *DishonorAch) SetFieldErrors(v []string) {
	o.FieldErrors = v
}

// GetType returns the Type field value
func (o *DishonorAch) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DishonorAch) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DishonorAch) SetType(v string) {
	o.Type = v
}

func (o DishonorAch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldErrors != nil {
		toSerialize["field_errors"] = o.FieldErrors
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDishonorAch struct {
	value *DishonorAch
	isSet bool
}

func (v NullableDishonorAch) Get() *DishonorAch {
	return v.value
}

func (v *NullableDishonorAch) Set(val *DishonorAch) {
	v.value = val
	v.isSet = true
}

func (v NullableDishonorAch) IsSet() bool {
	return v.isSet
}

func (v *NullableDishonorAch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDishonorAch(val *DishonorAch) *NullableDishonorAch {
	return &NullableDishonorAch{value: val, isSet: true}
}

func (v NullableDishonorAch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDishonorAch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
