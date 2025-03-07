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

// checks if the DepositPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepositPost{}

// DepositPost struct for DepositPost
type DepositPost struct {
	// The ID of the account
	AccountId string `json:"account_id"`
	// ID of the uploaded image of the back of the check
	BackImageId string `json:"back_image_id"`
	// Unique ID for the business. Exactly one of `business_id` or `person_id` must be set.
	BusinessId *string `json:"business_id,omitempty"`
	// Amount on check in ISO 4217 minor currency units
	CheckAmount int32 `json:"check_amount"`
	// ISO 4217 currency code for the deposit amount
	DepositCurrency string `json:"deposit_currency"`
	// ID of the uploaded image of the front of the check
	FrontImageId string `json:"front_image_id"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Unique ID for the person. Exactly one of `person_id` or `business_id` must be set.
	PersonId             *string `json:"person_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepositPost DepositPost

// NewDepositPost instantiates a new DepositPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositPost(accountId string, backImageId string, checkAmount int32, depositCurrency string, frontImageId string) *DepositPost {
	this := DepositPost{}
	this.AccountId = accountId
	this.BackImageId = backImageId
	this.CheckAmount = checkAmount
	this.DepositCurrency = depositCurrency
	this.FrontImageId = frontImageId
	return &this
}

// NewDepositPostWithDefaults instantiates a new DepositPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositPostWithDefaults() *DepositPost {
	this := DepositPost{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *DepositPost) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *DepositPost) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *DepositPost) SetAccountId(v string) {
	o.AccountId = v
}

// GetBackImageId returns the BackImageId field value
func (o *DepositPost) GetBackImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackImageId
}

// GetBackImageIdOk returns a tuple with the BackImageId field value
// and a boolean to check if the value has been set.
func (o *DepositPost) GetBackImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackImageId, true
}

// SetBackImageId sets field value
func (o *DepositPost) SetBackImageId(v string) {
	o.BackImageId = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *DepositPost) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPost) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *DepositPost) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *DepositPost) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCheckAmount returns the CheckAmount field value
func (o *DepositPost) GetCheckAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CheckAmount
}

// GetCheckAmountOk returns a tuple with the CheckAmount field value
// and a boolean to check if the value has been set.
func (o *DepositPost) GetCheckAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckAmount, true
}

// SetCheckAmount sets field value
func (o *DepositPost) SetCheckAmount(v int32) {
	o.CheckAmount = v
}

// GetDepositCurrency returns the DepositCurrency field value
func (o *DepositPost) GetDepositCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositCurrency
}

// GetDepositCurrencyOk returns a tuple with the DepositCurrency field value
// and a boolean to check if the value has been set.
func (o *DepositPost) GetDepositCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositCurrency, true
}

// SetDepositCurrency sets field value
func (o *DepositPost) SetDepositCurrency(v string) {
	o.DepositCurrency = v
}

// GetFrontImageId returns the FrontImageId field value
func (o *DepositPost) GetFrontImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FrontImageId
}

// GetFrontImageIdOk returns a tuple with the FrontImageId field value
// and a boolean to check if the value has been set.
func (o *DepositPost) GetFrontImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrontImageId, true
}

// SetFrontImageId sets field value
func (o *DepositPost) SetFrontImageId(v string) {
	o.FrontImageId = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DepositPost) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPost) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DepositPost) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *DepositPost) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *DepositPost) GetPersonId() string {
	if o == nil || IsNil(o.PersonId) {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPost) GetPersonIdOk() (*string, bool) {
	if o == nil || IsNil(o.PersonId) {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *DepositPost) HasPersonId() bool {
	if o != nil && !IsNil(o.PersonId) {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
func (o *DepositPost) SetPersonId(v string) {
	o.PersonId = &v
}

func (o DepositPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["back_image_id"] = o.BackImageId
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	toSerialize["check_amount"] = o.CheckAmount
	toSerialize["deposit_currency"] = o.DepositCurrency
	toSerialize["front_image_id"] = o.FrontImageId
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PersonId) {
		toSerialize["person_id"] = o.PersonId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DepositPost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"back_image_id",
		"check_amount",
		"deposit_currency",
		"front_image_id",
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

	varDepositPost := _DepositPost{}

	err = json.Unmarshal(data, &varDepositPost)

	if err != nil {
		return err
	}

	*o = DepositPost(varDepositPost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "back_image_id")
		delete(additionalProperties, "business_id")
		delete(additionalProperties, "check_amount")
		delete(additionalProperties, "deposit_currency")
		delete(additionalProperties, "front_image_id")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "person_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositPost struct {
	value *DepositPost
	isSet bool
}

func (v NullableDepositPost) Get() *DepositPost {
	return v.value
}

func (v *NullableDepositPost) Set(val *DepositPost) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositPost) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositPost(val *DepositPost) *NullableDepositPost {
	return &NullableDepositPost{value: val, isSet: true}
}

func (v NullableDepositPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
