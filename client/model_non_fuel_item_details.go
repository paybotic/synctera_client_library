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

// checks if the NonFuelItemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonFuelItemDetails{}

// NonFuelItemDetails struct for NonFuelItemDetails
type NonFuelItemDetails struct {
	Description          *string `json:"description,omitempty"`
	ProductCode          *string `json:"product_code,omitempty"`
	Quantity             *int64  `json:"quantity,omitempty"`
	UnitOfMeasure        *string `json:"unit_of_measure,omitempty"`
	UnitPrice            *int64  `json:"unit_price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonFuelItemDetails NonFuelItemDetails

// NewNonFuelItemDetails instantiates a new NonFuelItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonFuelItemDetails() *NonFuelItemDetails {
	this := NonFuelItemDetails{}
	return &this
}

// NewNonFuelItemDetailsWithDefaults instantiates a new NonFuelItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonFuelItemDetailsWithDefaults() *NonFuelItemDetails {
	this := NonFuelItemDetails{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NonFuelItemDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFuelItemDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NonFuelItemDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NonFuelItemDetails) SetDescription(v string) {
	o.Description = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *NonFuelItemDetails) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFuelItemDetails) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *NonFuelItemDetails) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *NonFuelItemDetails) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *NonFuelItemDetails) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFuelItemDetails) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *NonFuelItemDetails) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *NonFuelItemDetails) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *NonFuelItemDetails) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFuelItemDetails) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *NonFuelItemDetails) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *NonFuelItemDetails) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *NonFuelItemDetails) GetUnitPrice() int64 {
	if o == nil || IsNil(o.UnitPrice) {
		var ret int64
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonFuelItemDetails) GetUnitPriceOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *NonFuelItemDetails) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given int64 and assigns it to the UnitPrice field.
func (o *NonFuelItemDetails) SetUnitPrice(v int64) {
	o.UnitPrice = &v
}

func (o NonFuelItemDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonFuelItemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.UnitOfMeasure) {
		toSerialize["unit_of_measure"] = o.UnitOfMeasure
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unit_price"] = o.UnitPrice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NonFuelItemDetails) UnmarshalJSON(data []byte) (err error) {
	varNonFuelItemDetails := _NonFuelItemDetails{}

	err = json.Unmarshal(data, &varNonFuelItemDetails)

	if err != nil {
		return err
	}

	*o = NonFuelItemDetails(varNonFuelItemDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "product_code")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unit_of_measure")
		delete(additionalProperties, "unit_price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonFuelItemDetails struct {
	value *NonFuelItemDetails
	isSet bool
}

func (v NullableNonFuelItemDetails) Get() *NonFuelItemDetails {
	return v.value
}

func (v *NullableNonFuelItemDetails) Set(val *NonFuelItemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNonFuelItemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNonFuelItemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonFuelItemDetails(val *NonFuelItemDetails) *NullableNonFuelItemDetails {
	return &NullableNonFuelItemDetails{value: val, isSet: true}
}

func (v NullableNonFuelItemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonFuelItemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
