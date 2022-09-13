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

// FulfillmentDetails struct for FulfillmentDetails
type FulfillmentDetails struct {
	// The date that the card was shipped as reported by the card fulfillment provider
	ShipDate *string `json:"ship_date,omitempty"`
	// The specific shipping method as reported by the card fulfillment provider
	ShippingMethod *string `json:"shipping_method,omitempty"`
	// The shipment tracking number
	TrackingNumber *string `json:"tracking_number,omitempty"`
}

// NewFulfillmentDetails instantiates a new FulfillmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentDetails() *FulfillmentDetails {
	this := FulfillmentDetails{}
	return &this
}

// NewFulfillmentDetailsWithDefaults instantiates a new FulfillmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentDetailsWithDefaults() *FulfillmentDetails {
	this := FulfillmentDetails{}
	return &this
}

// GetShipDate returns the ShipDate field value if set, zero value otherwise.
func (o *FulfillmentDetails) GetShipDate() string {
	if o == nil || o.ShipDate == nil {
		var ret string
		return ret
	}
	return *o.ShipDate
}

// GetShipDateOk returns a tuple with the ShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentDetails) GetShipDateOk() (*string, bool) {
	if o == nil || o.ShipDate == nil {
		return nil, false
	}
	return o.ShipDate, true
}

// HasShipDate returns a boolean if a field has been set.
func (o *FulfillmentDetails) HasShipDate() bool {
	if o != nil && o.ShipDate != nil {
		return true
	}

	return false
}

// SetShipDate gets a reference to the given string and assigns it to the ShipDate field.
func (o *FulfillmentDetails) SetShipDate(v string) {
	o.ShipDate = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *FulfillmentDetails) GetShippingMethod() string {
	if o == nil || o.ShippingMethod == nil {
		var ret string
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentDetails) GetShippingMethodOk() (*string, bool) {
	if o == nil || o.ShippingMethod == nil {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *FulfillmentDetails) HasShippingMethod() bool {
	if o != nil && o.ShippingMethod != nil {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given string and assigns it to the ShippingMethod field.
func (o *FulfillmentDetails) SetShippingMethod(v string) {
	o.ShippingMethod = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *FulfillmentDetails) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber == nil {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentDetails) GetTrackingNumberOk() (*string, bool) {
	if o == nil || o.TrackingNumber == nil {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *FulfillmentDetails) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber != nil {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *FulfillmentDetails) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

func (o FulfillmentDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShipDate != nil {
		toSerialize["ship_date"] = o.ShipDate
	}
	if o.ShippingMethod != nil {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	if o.TrackingNumber != nil {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	return json.Marshal(toSerialize)
}

type NullableFulfillmentDetails struct {
	value *FulfillmentDetails
	isSet bool
}

func (v NullableFulfillmentDetails) Get() *FulfillmentDetails {
	return v.value
}

func (v *NullableFulfillmentDetails) Set(val *FulfillmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentDetails(val *FulfillmentDetails) *NullableFulfillmentDetails {
	return &NullableFulfillmentDetails{value: val, isSet: true}
}

func (v NullableFulfillmentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


