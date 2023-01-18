/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CardIssuanceRequest - struct for CardIssuanceRequest
type CardIssuanceRequest struct {
	PhysicalCardIssuanceRequest *PhysicalCardIssuanceRequest
	VirtualCardIssuanceRequest *VirtualCardIssuanceRequest
}

// PhysicalCardIssuanceRequestAsCardIssuanceRequest is a convenience function that returns PhysicalCardIssuanceRequest wrapped in CardIssuanceRequest
func PhysicalCardIssuanceRequestAsCardIssuanceRequest(v *PhysicalCardIssuanceRequest) CardIssuanceRequest {
	return CardIssuanceRequest{
		PhysicalCardIssuanceRequest: v,
	}
}

// VirtualCardIssuanceRequestAsCardIssuanceRequest is a convenience function that returns VirtualCardIssuanceRequest wrapped in CardIssuanceRequest
func VirtualCardIssuanceRequestAsCardIssuanceRequest(v *VirtualCardIssuanceRequest) CardIssuanceRequest {
	return CardIssuanceRequest{
		VirtualCardIssuanceRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CardIssuanceRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'PHYSICAL'
	if jsonDict["form"] == "PHYSICAL" {
		// try to unmarshal JSON data into PhysicalCardIssuanceRequest
		err = json.Unmarshal(data, &dst.PhysicalCardIssuanceRequest)
		if err == nil {
			return nil // data stored in dst.PhysicalCardIssuanceRequest, return on the first match
		} else {
			dst.PhysicalCardIssuanceRequest = nil
			return fmt.Errorf("Failed to unmarshal CardIssuanceRequest as PhysicalCardIssuanceRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'VIRTUAL'
	if jsonDict["form"] == "VIRTUAL" {
		// try to unmarshal JSON data into VirtualCardIssuanceRequest
		err = json.Unmarshal(data, &dst.VirtualCardIssuanceRequest)
		if err == nil {
			return nil // data stored in dst.VirtualCardIssuanceRequest, return on the first match
		} else {
			dst.VirtualCardIssuanceRequest = nil
			return fmt.Errorf("Failed to unmarshal CardIssuanceRequest as VirtualCardIssuanceRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'physical_card_issuance_request'
	if jsonDict["form"] == "physical_card_issuance_request" {
		// try to unmarshal JSON data into PhysicalCardIssuanceRequest
		err = json.Unmarshal(data, &dst.PhysicalCardIssuanceRequest)
		if err == nil {
			return nil // data stored in dst.PhysicalCardIssuanceRequest, return on the first match
		} else {
			dst.PhysicalCardIssuanceRequest = nil
			return fmt.Errorf("Failed to unmarshal CardIssuanceRequest as PhysicalCardIssuanceRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'virtual_card_issuance_request'
	if jsonDict["form"] == "virtual_card_issuance_request" {
		// try to unmarshal JSON data into VirtualCardIssuanceRequest
		err = json.Unmarshal(data, &dst.VirtualCardIssuanceRequest)
		if err == nil {
			return nil // data stored in dst.VirtualCardIssuanceRequest, return on the first match
		} else {
			dst.VirtualCardIssuanceRequest = nil
			return fmt.Errorf("Failed to unmarshal CardIssuanceRequest as VirtualCardIssuanceRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CardIssuanceRequest) MarshalJSON() ([]byte, error) {
	if src.PhysicalCardIssuanceRequest != nil {
		return json.Marshal(&src.PhysicalCardIssuanceRequest)
	}

	if src.VirtualCardIssuanceRequest != nil {
		return json.Marshal(&src.VirtualCardIssuanceRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CardIssuanceRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PhysicalCardIssuanceRequest != nil {
		return obj.PhysicalCardIssuanceRequest
	}

	if obj.VirtualCardIssuanceRequest != nil {
		return obj.VirtualCardIssuanceRequest
	}

	// all schemas are nil
	return nil
}

type NullableCardIssuanceRequest struct {
	value *CardIssuanceRequest
	isSet bool
}

func (v NullableCardIssuanceRequest) Get() *CardIssuanceRequest {
	return v.value
}

func (v *NullableCardIssuanceRequest) Set(val *CardIssuanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardIssuanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardIssuanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardIssuanceRequest(val *CardIssuanceRequest) *NullableCardIssuanceRequest {
	return &NullableCardIssuanceRequest{value: val, isSet: true}
}

func (v NullableCardIssuanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardIssuanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


