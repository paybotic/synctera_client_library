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

// checks if the CustomerServiceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerServiceDetails{}

// CustomerServiceDetails The customer service details of the fintech partner, e.g. phone number, email address, etc.
type CustomerServiceDetails struct {
	Address *Address2 `json:"address,omitempty"`
	// The customer service email address
	Email *string `json:"email,omitempty"`
	// The customer service phone number
	PhoneNumber          *string `json:"phone_number,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerServiceDetails CustomerServiceDetails

// NewCustomerServiceDetails instantiates a new CustomerServiceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerServiceDetails() *CustomerServiceDetails {
	this := CustomerServiceDetails{}
	return &this
}

// NewCustomerServiceDetailsWithDefaults instantiates a new CustomerServiceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerServiceDetailsWithDefaults() *CustomerServiceDetails {
	this := CustomerServiceDetails{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CustomerServiceDetails) GetAddress() Address2 {
	if o == nil || IsNil(o.Address) {
		var ret Address2
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceDetails) GetAddressOk() (*Address2, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CustomerServiceDetails) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address2 and assigns it to the Address field.
func (o *CustomerServiceDetails) SetAddress(v Address2) {
	o.Address = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerServiceDetails) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceDetails) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerServiceDetails) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerServiceDetails) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CustomerServiceDetails) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceDetails) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerServiceDetails) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CustomerServiceDetails) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o CustomerServiceDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerServiceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerServiceDetails) UnmarshalJSON(data []byte) (err error) {
	varCustomerServiceDetails := _CustomerServiceDetails{}

	err = json.Unmarshal(data, &varCustomerServiceDetails)

	if err != nil {
		return err
	}

	*o = CustomerServiceDetails(varCustomerServiceDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerServiceDetails struct {
	value *CustomerServiceDetails
	isSet bool
}

func (v NullableCustomerServiceDetails) Get() *CustomerServiceDetails {
	return v.value
}

func (v *NullableCustomerServiceDetails) Set(val *CustomerServiceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerServiceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerServiceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerServiceDetails(val *CustomerServiceDetails) *NullableCustomerServiceDetails {
	return &NullableCustomerServiceDetails{value: val, isSet: true}
}

func (v NullableCustomerServiceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerServiceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
