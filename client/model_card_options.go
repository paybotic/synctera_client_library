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

// checks if the CardOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardOptions{}

// CardOptions struct for CardOptions
type CardOptions struct {
	BillingAddress       *BillingAddress `json:"billing_address,omitempty"`
	CardPresent          *bool           `json:"card_present,omitempty"`
	Cvv                  *string         `json:"cvv,omitempty"`
	Expiration           *string         `json:"expiration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CardOptions CardOptions

// NewCardOptions instantiates a new CardOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardOptions() *CardOptions {
	this := CardOptions{}
	var cardPresent bool = false
	this.CardPresent = &cardPresent
	return &this
}

// NewCardOptionsWithDefaults instantiates a new CardOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardOptionsWithDefaults() *CardOptions {
	this := CardOptions{}
	var cardPresent bool = false
	this.CardPresent = &cardPresent
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *CardOptions) GetBillingAddress() BillingAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret BillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOptions) GetBillingAddressOk() (*BillingAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *CardOptions) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given BillingAddress and assigns it to the BillingAddress field.
func (o *CardOptions) SetBillingAddress(v BillingAddress) {
	o.BillingAddress = &v
}

// GetCardPresent returns the CardPresent field value if set, zero value otherwise.
func (o *CardOptions) GetCardPresent() bool {
	if o == nil || IsNil(o.CardPresent) {
		var ret bool
		return ret
	}
	return *o.CardPresent
}

// GetCardPresentOk returns a tuple with the CardPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOptions) GetCardPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.CardPresent) {
		return nil, false
	}
	return o.CardPresent, true
}

// HasCardPresent returns a boolean if a field has been set.
func (o *CardOptions) HasCardPresent() bool {
	if o != nil && !IsNil(o.CardPresent) {
		return true
	}

	return false
}

// SetCardPresent gets a reference to the given bool and assigns it to the CardPresent field.
func (o *CardOptions) SetCardPresent(v bool) {
	o.CardPresent = &v
}

// GetCvv returns the Cvv field value if set, zero value otherwise.
func (o *CardOptions) GetCvv() string {
	if o == nil || IsNil(o.Cvv) {
		var ret string
		return ret
	}
	return *o.Cvv
}

// GetCvvOk returns a tuple with the Cvv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOptions) GetCvvOk() (*string, bool) {
	if o == nil || IsNil(o.Cvv) {
		return nil, false
	}
	return o.Cvv, true
}

// HasCvv returns a boolean if a field has been set.
func (o *CardOptions) HasCvv() bool {
	if o != nil && !IsNil(o.Cvv) {
		return true
	}

	return false
}

// SetCvv gets a reference to the given string and assigns it to the Cvv field.
func (o *CardOptions) SetCvv(v string) {
	o.Cvv = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *CardOptions) GetExpiration() string {
	if o == nil || IsNil(o.Expiration) {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOptions) GetExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *CardOptions) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *CardOptions) SetExpiration(v string) {
	o.Expiration = &v
}

func (o CardOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if !IsNil(o.CardPresent) {
		toSerialize["card_present"] = o.CardPresent
	}
	if !IsNil(o.Cvv) {
		toSerialize["cvv"] = o.Cvv
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardOptions) UnmarshalJSON(data []byte) (err error) {
	varCardOptions := _CardOptions{}

	err = json.Unmarshal(data, &varCardOptions)

	if err != nil {
		return err
	}

	*o = CardOptions(varCardOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "billing_address")
		delete(additionalProperties, "card_present")
		delete(additionalProperties, "cvv")
		delete(additionalProperties, "expiration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardOptions struct {
	value *CardOptions
	isSet bool
}

func (v NullableCardOptions) Get() *CardOptions {
	return v.value
}

func (v *NullableCardOptions) Set(val *CardOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCardOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCardOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardOptions(val *CardOptions) *NullableCardOptions {
	return &NullableCardOptions{value: val, isSet: true}
}

func (v NullableCardOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
