/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TemplateFields - struct for TemplateFields
type TemplateFields struct {
	TemplateFieldsDepository *TemplateFieldsDepository
	TemplateFieldsLineOfCredit *TemplateFieldsLineOfCredit
}

// TemplateFieldsDepositoryAsTemplateFields is a convenience function that returns TemplateFieldsDepository wrapped in TemplateFields
func TemplateFieldsDepositoryAsTemplateFields(v *TemplateFieldsDepository) TemplateFields {
	return TemplateFields{
		TemplateFieldsDepository: v,
	}
}

// TemplateFieldsLineOfCreditAsTemplateFields is a convenience function that returns TemplateFieldsLineOfCredit wrapped in TemplateFields
func TemplateFieldsLineOfCreditAsTemplateFields(v *TemplateFieldsLineOfCredit) TemplateFields {
	return TemplateFields{
		TemplateFieldsLineOfCredit: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TemplateFields) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TemplateFieldsDepository
	err = newStrictDecoder(data).Decode(&dst.TemplateFieldsDepository)
	if err == nil {
		jsonTemplateFieldsDepository, _ := json.Marshal(dst.TemplateFieldsDepository)
		if string(jsonTemplateFieldsDepository) == "{}" { // empty struct
			dst.TemplateFieldsDepository = nil
		} else {
			match++
		}
	} else {
		dst.TemplateFieldsDepository = nil
	}

	// try to unmarshal data into TemplateFieldsLineOfCredit
	err = newStrictDecoder(data).Decode(&dst.TemplateFieldsLineOfCredit)
	if err == nil {
		jsonTemplateFieldsLineOfCredit, _ := json.Marshal(dst.TemplateFieldsLineOfCredit)
		if string(jsonTemplateFieldsLineOfCredit) == "{}" { // empty struct
			dst.TemplateFieldsLineOfCredit = nil
		} else {
			match++
		}
	} else {
		dst.TemplateFieldsLineOfCredit = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TemplateFieldsDepository = nil
		dst.TemplateFieldsLineOfCredit = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(TemplateFields)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(TemplateFields)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TemplateFields) MarshalJSON() ([]byte, error) {
	if src.TemplateFieldsDepository != nil {
		return json.Marshal(&src.TemplateFieldsDepository)
	}

	if src.TemplateFieldsLineOfCredit != nil {
		return json.Marshal(&src.TemplateFieldsLineOfCredit)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TemplateFields) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TemplateFieldsDepository != nil {
		return obj.TemplateFieldsDepository
	}

	if obj.TemplateFieldsLineOfCredit != nil {
		return obj.TemplateFieldsLineOfCredit
	}

	// all schemas are nil
	return nil
}

type NullableTemplateFields struct {
	value *TemplateFields
	isSet bool
}

func (v NullableTemplateFields) Get() *TemplateFields {
	return v.value
}

func (v *NullableTemplateFields) Set(val *TemplateFields) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFields) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFields(val *TemplateFields) *NullableTemplateFields {
	return &NullableTemplateFields{value: val, isSet: true}
}

func (v NullableTemplateFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


