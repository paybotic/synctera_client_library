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

// checks if the ModelCase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCase{}

// ModelCase struct for ModelCase
type ModelCase struct {
	// Case ID
	CaseId string `json:"case_id"`
	// Case Reason
	CaseReason           string `json:"case_reason"`
	AdditionalProperties map[string]interface{}
}

type _ModelCase ModelCase

// NewModelCase instantiates a new ModelCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCase(caseId string, caseReason string) *ModelCase {
	this := ModelCase{}
	this.CaseId = caseId
	this.CaseReason = caseReason
	return &this
}

// NewModelCaseWithDefaults instantiates a new ModelCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCaseWithDefaults() *ModelCase {
	this := ModelCase{}
	return &this
}

// GetCaseId returns the CaseId field value
func (o *ModelCase) GetCaseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value
// and a boolean to check if the value has been set.
func (o *ModelCase) GetCaseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseId, true
}

// SetCaseId sets field value
func (o *ModelCase) SetCaseId(v string) {
	o.CaseId = v
}

// GetCaseReason returns the CaseReason field value
func (o *ModelCase) GetCaseReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaseReason
}

// GetCaseReasonOk returns a tuple with the CaseReason field value
// and a boolean to check if the value has been set.
func (o *ModelCase) GetCaseReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseReason, true
}

// SetCaseReason sets field value
func (o *ModelCase) SetCaseReason(v string) {
	o.CaseReason = v
}

func (o ModelCase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["case_id"] = o.CaseId
	toSerialize["case_reason"] = o.CaseReason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelCase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"case_id",
		"case_reason",
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

	varModelCase := _ModelCase{}

	err = json.Unmarshal(data, &varModelCase)

	if err != nil {
		return err
	}

	*o = ModelCase(varModelCase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "case_id")
		delete(additionalProperties, "case_reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelCase struct {
	value *ModelCase
	isSet bool
}

func (v NullableModelCase) Get() *ModelCase {
	return v.value
}

func (v *NullableModelCase) Set(val *ModelCase) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCase) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCase(val *ModelCase) *NullableModelCase {
	return &NullableModelCase{value: val, isSet: true}
}

func (v NullableModelCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
