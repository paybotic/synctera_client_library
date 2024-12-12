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

// checks if the CustomerVerificationResultList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerVerificationResultList{}

// CustomerVerificationResultList struct for CustomerVerificationResultList
type CustomerVerificationResultList struct {
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// Array of verification results.
	Verifications []CustomerVerificationResult `json:"verifications"`
}

type _CustomerVerificationResultList CustomerVerificationResultList

// NewCustomerVerificationResultList instantiates a new CustomerVerificationResultList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerVerificationResultList(verifications []CustomerVerificationResult) *CustomerVerificationResultList {
	this := CustomerVerificationResultList{}
	this.Verifications = verifications
	return &this
}

// NewCustomerVerificationResultListWithDefaults instantiates a new CustomerVerificationResultList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerVerificationResultListWithDefaults() *CustomerVerificationResultList {
	this := CustomerVerificationResultList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CustomerVerificationResultList) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerVerificationResultList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CustomerVerificationResultList) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CustomerVerificationResultList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetVerifications returns the Verifications field value
func (o *CustomerVerificationResultList) GetVerifications() []CustomerVerificationResult {
	if o == nil {
		var ret []CustomerVerificationResult
		return ret
	}

	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value
// and a boolean to check if the value has been set.
func (o *CustomerVerificationResultList) GetVerificationsOk() ([]CustomerVerificationResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verifications, true
}

// SetVerifications sets field value
func (o *CustomerVerificationResultList) SetVerifications(v []CustomerVerificationResult) {
	o.Verifications = v
}

func (o CustomerVerificationResultList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerVerificationResultList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	toSerialize["verifications"] = o.Verifications
	return toSerialize, nil
}

func (o *CustomerVerificationResultList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"verifications",
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

	varCustomerVerificationResultList := _CustomerVerificationResultList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerVerificationResultList)

	if err != nil {
		return err
	}

	*o = CustomerVerificationResultList(varCustomerVerificationResultList)

	return err
}

type NullableCustomerVerificationResultList struct {
	value *CustomerVerificationResultList
	isSet bool
}

func (v NullableCustomerVerificationResultList) Get() *CustomerVerificationResultList {
	return v.value
}

func (v *NullableCustomerVerificationResultList) Set(val *CustomerVerificationResultList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerVerificationResultList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerVerificationResultList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerVerificationResultList(val *CustomerVerificationResultList) *NullableCustomerVerificationResultList {
	return &NullableCustomerVerificationResultList{value: val, isSet: true}
}

func (v NullableCustomerVerificationResultList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerVerificationResultList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
