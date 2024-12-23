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

// VerificationRequestAllOf struct for VerificationRequestAllOf
type VerificationRequestAllOf struct {
	// Unique ID for the business. Exactly one of `business_id` or `person_id` must be set.
	BusinessId *string `json:"business_id,omitempty"`
	// Whether this customer has consented to be verified.
	CustomerConsent *bool `json:"customer_consent,omitempty"`
	// The customer's IP address.
	CustomerIpAddress *string `json:"customer_ip_address,omitempty"`
	// The ID of the uploaded government-issued identification document provided by the Socure SDK.
	DocumentId *string `json:"document_id,omitempty"`
	// Unique ID for the person. Exactly one of `person_id` or `business_id` must be set.
	PersonId *string `json:"person_id,omitempty"`
}

// NewVerificationRequestAllOf instantiates a new VerificationRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationRequestAllOf() *VerificationRequestAllOf {
	this := VerificationRequestAllOf{}
	return &this
}

// NewVerificationRequestAllOfWithDefaults instantiates a new VerificationRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationRequestAllOfWithDefaults() *VerificationRequestAllOf {
	this := VerificationRequestAllOf{}
	return &this
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *VerificationRequestAllOf) GetBusinessId() string {
	if o == nil || o.BusinessId == nil {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationRequestAllOf) GetBusinessIdOk() (*string, bool) {
	if o == nil || o.BusinessId == nil {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *VerificationRequestAllOf) HasBusinessId() bool {
	if o != nil && o.BusinessId != nil {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *VerificationRequestAllOf) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCustomerConsent returns the CustomerConsent field value if set, zero value otherwise.
func (o *VerificationRequestAllOf) GetCustomerConsent() bool {
	if o == nil || o.CustomerConsent == nil {
		var ret bool
		return ret
	}
	return *o.CustomerConsent
}

// GetCustomerConsentOk returns a tuple with the CustomerConsent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationRequestAllOf) GetCustomerConsentOk() (*bool, bool) {
	if o == nil || o.CustomerConsent == nil {
		return nil, false
	}
	return o.CustomerConsent, true
}

// HasCustomerConsent returns a boolean if a field has been set.
func (o *VerificationRequestAllOf) HasCustomerConsent() bool {
	if o != nil && o.CustomerConsent != nil {
		return true
	}

	return false
}

// SetCustomerConsent gets a reference to the given bool and assigns it to the CustomerConsent field.
func (o *VerificationRequestAllOf) SetCustomerConsent(v bool) {
	o.CustomerConsent = &v
}

// GetCustomerIpAddress returns the CustomerIpAddress field value if set, zero value otherwise.
func (o *VerificationRequestAllOf) GetCustomerIpAddress() string {
	if o == nil || o.CustomerIpAddress == nil {
		var ret string
		return ret
	}
	return *o.CustomerIpAddress
}

// GetCustomerIpAddressOk returns a tuple with the CustomerIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationRequestAllOf) GetCustomerIpAddressOk() (*string, bool) {
	if o == nil || o.CustomerIpAddress == nil {
		return nil, false
	}
	return o.CustomerIpAddress, true
}

// HasCustomerIpAddress returns a boolean if a field has been set.
func (o *VerificationRequestAllOf) HasCustomerIpAddress() bool {
	if o != nil && o.CustomerIpAddress != nil {
		return true
	}

	return false
}

// SetCustomerIpAddress gets a reference to the given string and assigns it to the CustomerIpAddress field.
func (o *VerificationRequestAllOf) SetCustomerIpAddress(v string) {
	o.CustomerIpAddress = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *VerificationRequestAllOf) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationRequestAllOf) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *VerificationRequestAllOf) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *VerificationRequestAllOf) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *VerificationRequestAllOf) GetPersonId() string {
	if o == nil || o.PersonId == nil {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationRequestAllOf) GetPersonIdOk() (*string, bool) {
	if o == nil || o.PersonId == nil {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *VerificationRequestAllOf) HasPersonId() bool {
	if o != nil && o.PersonId != nil {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
func (o *VerificationRequestAllOf) SetPersonId(v string) {
	o.PersonId = &v
}

func (o VerificationRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BusinessId != nil {
		toSerialize["business_id"] = o.BusinessId
	}
	if o.CustomerConsent != nil {
		toSerialize["customer_consent"] = o.CustomerConsent
	}
	if o.CustomerIpAddress != nil {
		toSerialize["customer_ip_address"] = o.CustomerIpAddress
	}
	if o.DocumentId != nil {
		toSerialize["document_id"] = o.DocumentId
	}
	if o.PersonId != nil {
		toSerialize["person_id"] = o.PersonId
	}
	return json.Marshal(toSerialize)
}

type NullableVerificationRequestAllOf struct {
	value *VerificationRequestAllOf
	isSet bool
}

func (v NullableVerificationRequestAllOf) Get() *VerificationRequestAllOf {
	return v.value
}

func (v *NullableVerificationRequestAllOf) Set(val *VerificationRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationRequestAllOf(val *VerificationRequestAllOf) *NullableVerificationRequestAllOf {
	return &NullableVerificationRequestAllOf{value: val, isSet: true}
}

func (v NullableVerificationRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
