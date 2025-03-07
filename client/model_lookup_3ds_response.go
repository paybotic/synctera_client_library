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

// checks if the Lookup3dsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Lookup3dsResponse{}

// Lookup3dsResponse Details for an External Card Transfer 3-D Secure Authentication lookup response
type Lookup3dsResponse struct {
	// 3DS challenge payload, returned if challenge is required
	ChallengePayload *string `json:"challenge_payload,omitempty"`
	// 3DS challenge URL, returned if challenge is required
	ChallengeUrl *string `json:"challenge_url,omitempty"`
	// The unique identifier of the 3DS authentication
	Id string `json:"id"`
	// Processor Transaction ID, returned if challenge is required
	ProcessorTransactionId *string `json:"processor_transaction_id,omitempty"`
	// Status of the 3DS authentication
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _Lookup3dsResponse Lookup3dsResponse

// NewLookup3dsResponse instantiates a new Lookup3dsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookup3dsResponse(id string, status string) *Lookup3dsResponse {
	this := Lookup3dsResponse{}
	this.Id = id
	this.Status = status
	return &this
}

// NewLookup3dsResponseWithDefaults instantiates a new Lookup3dsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookup3dsResponseWithDefaults() *Lookup3dsResponse {
	this := Lookup3dsResponse{}
	return &this
}

// GetChallengePayload returns the ChallengePayload field value if set, zero value otherwise.
func (o *Lookup3dsResponse) GetChallengePayload() string {
	if o == nil || IsNil(o.ChallengePayload) {
		var ret string
		return ret
	}
	return *o.ChallengePayload
}

// GetChallengePayloadOk returns a tuple with the ChallengePayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lookup3dsResponse) GetChallengePayloadOk() (*string, bool) {
	if o == nil || IsNil(o.ChallengePayload) {
		return nil, false
	}
	return o.ChallengePayload, true
}

// HasChallengePayload returns a boolean if a field has been set.
func (o *Lookup3dsResponse) HasChallengePayload() bool {
	if o != nil && !IsNil(o.ChallengePayload) {
		return true
	}

	return false
}

// SetChallengePayload gets a reference to the given string and assigns it to the ChallengePayload field.
func (o *Lookup3dsResponse) SetChallengePayload(v string) {
	o.ChallengePayload = &v
}

// GetChallengeUrl returns the ChallengeUrl field value if set, zero value otherwise.
func (o *Lookup3dsResponse) GetChallengeUrl() string {
	if o == nil || IsNil(o.ChallengeUrl) {
		var ret string
		return ret
	}
	return *o.ChallengeUrl
}

// GetChallengeUrlOk returns a tuple with the ChallengeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lookup3dsResponse) GetChallengeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ChallengeUrl) {
		return nil, false
	}
	return o.ChallengeUrl, true
}

// HasChallengeUrl returns a boolean if a field has been set.
func (o *Lookup3dsResponse) HasChallengeUrl() bool {
	if o != nil && !IsNil(o.ChallengeUrl) {
		return true
	}

	return false
}

// SetChallengeUrl gets a reference to the given string and assigns it to the ChallengeUrl field.
func (o *Lookup3dsResponse) SetChallengeUrl(v string) {
	o.ChallengeUrl = &v
}

// GetId returns the Id field value
func (o *Lookup3dsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Lookup3dsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Lookup3dsResponse) SetId(v string) {
	o.Id = v
}

// GetProcessorTransactionId returns the ProcessorTransactionId field value if set, zero value otherwise.
func (o *Lookup3dsResponse) GetProcessorTransactionId() string {
	if o == nil || IsNil(o.ProcessorTransactionId) {
		var ret string
		return ret
	}
	return *o.ProcessorTransactionId
}

// GetProcessorTransactionIdOk returns a tuple with the ProcessorTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Lookup3dsResponse) GetProcessorTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorTransactionId) {
		return nil, false
	}
	return o.ProcessorTransactionId, true
}

// HasProcessorTransactionId returns a boolean if a field has been set.
func (o *Lookup3dsResponse) HasProcessorTransactionId() bool {
	if o != nil && !IsNil(o.ProcessorTransactionId) {
		return true
	}

	return false
}

// SetProcessorTransactionId gets a reference to the given string and assigns it to the ProcessorTransactionId field.
func (o *Lookup3dsResponse) SetProcessorTransactionId(v string) {
	o.ProcessorTransactionId = &v
}

// GetStatus returns the Status field value
func (o *Lookup3dsResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Lookup3dsResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Lookup3dsResponse) SetStatus(v string) {
	o.Status = v
}

func (o Lookup3dsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Lookup3dsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChallengePayload) {
		toSerialize["challenge_payload"] = o.ChallengePayload
	}
	if !IsNil(o.ChallengeUrl) {
		toSerialize["challenge_url"] = o.ChallengeUrl
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.ProcessorTransactionId) {
		toSerialize["processor_transaction_id"] = o.ProcessorTransactionId
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Lookup3dsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
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

	varLookup3dsResponse := _Lookup3dsResponse{}

	err = json.Unmarshal(data, &varLookup3dsResponse)

	if err != nil {
		return err
	}

	*o = Lookup3dsResponse(varLookup3dsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "challenge_payload")
		delete(additionalProperties, "challenge_url")
		delete(additionalProperties, "id")
		delete(additionalProperties, "processor_transaction_id")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLookup3dsResponse struct {
	value *Lookup3dsResponse
	isSet bool
}

func (v NullableLookup3dsResponse) Get() *Lookup3dsResponse {
	return v.value
}

func (v *NullableLookup3dsResponse) Set(val *Lookup3dsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLookup3dsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLookup3dsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookup3dsResponse(val *Lookup3dsResponse) *NullableLookup3dsResponse {
	return &NullableLookup3dsResponse{value: val, isSet: true}
}

func (v NullableLookup3dsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookup3dsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
