/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PhysicalCardAllOf struct for PhysicalCardAllOf
type PhysicalCardAllOf struct {
	// The ID of the custom card image used for this card
	CardImageId *string `json:"card_image_id,omitempty"`
	Shipping *Shipping `json:"shipping,omitempty"`
}

// NewPhysicalCardAllOf instantiates a new PhysicalCardAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalCardAllOf() *PhysicalCardAllOf {
	this := PhysicalCardAllOf{}
	return &this
}

// NewPhysicalCardAllOfWithDefaults instantiates a new PhysicalCardAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalCardAllOfWithDefaults() *PhysicalCardAllOf {
	this := PhysicalCardAllOf{}
	return &this
}

// GetCardImageId returns the CardImageId field value if set, zero value otherwise.
func (o *PhysicalCardAllOf) GetCardImageId() string {
	if o == nil || o.CardImageId == nil {
		var ret string
		return ret
	}
	return *o.CardImageId
}

// GetCardImageIdOk returns a tuple with the CardImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardAllOf) GetCardImageIdOk() (*string, bool) {
	if o == nil || o.CardImageId == nil {
		return nil, false
	}
	return o.CardImageId, true
}

// HasCardImageId returns a boolean if a field has been set.
func (o *PhysicalCardAllOf) HasCardImageId() bool {
	if o != nil && o.CardImageId != nil {
		return true
	}

	return false
}

// SetCardImageId gets a reference to the given string and assigns it to the CardImageId field.
func (o *PhysicalCardAllOf) SetCardImageId(v string) {
	o.CardImageId = &v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *PhysicalCardAllOf) GetShipping() Shipping {
	if o == nil || o.Shipping == nil {
		var ret Shipping
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardAllOf) GetShippingOk() (*Shipping, bool) {
	if o == nil || o.Shipping == nil {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *PhysicalCardAllOf) HasShipping() bool {
	if o != nil && o.Shipping != nil {
		return true
	}

	return false
}

// SetShipping gets a reference to the given Shipping and assigns it to the Shipping field.
func (o *PhysicalCardAllOf) SetShipping(v Shipping) {
	o.Shipping = &v
}

func (o PhysicalCardAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CardImageId != nil {
		toSerialize["card_image_id"] = o.CardImageId
	}
	if o.Shipping != nil {
		toSerialize["shipping"] = o.Shipping
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalCardAllOf struct {
	value *PhysicalCardAllOf
	isSet bool
}

func (v NullablePhysicalCardAllOf) Get() *PhysicalCardAllOf {
	return v.value
}

func (v *NullablePhysicalCardAllOf) Set(val *PhysicalCardAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalCardAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalCardAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalCardAllOf(val *PhysicalCardAllOf) *NullablePhysicalCardAllOf {
	return &NullablePhysicalCardAllOf{value: val, isSet: true}
}

func (v NullablePhysicalCardAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalCardAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


