/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the TransactionDispute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDispute{}

// TransactionDispute struct for TransactionDispute
type TransactionDispute struct {
	// The creation time of the dispute
	Created time.Time `json:"created"`
	// The external case number or id for the dispute (eg: from a vendor such as Marqeta), if one exists.
	ExternalCaseReference *string `json:"external_case_reference,omitempty"`
	// The unique identifier of the dispute.
	Id string `json:"id"`
	// The internal case number or id for the dispute in the Synctera platform, if one exists.
	InternalCaseReference *string `json:"internal_case_reference,omitempty"`
	Status                string  `json:"status"`
	// The time the dispute was last updated
	Updated time.Time `json:"updated"`
}

type _TransactionDispute TransactionDispute

// NewTransactionDispute instantiates a new TransactionDispute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDispute(created time.Time, id string, status string, updated time.Time) *TransactionDispute {
	this := TransactionDispute{}
	this.Created = created
	this.Id = id
	this.Status = status
	this.Updated = updated
	return &this
}

// NewTransactionDisputeWithDefaults instantiates a new TransactionDispute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDisputeWithDefaults() *TransactionDispute {
	this := TransactionDispute{}
	return &this
}

// GetCreated returns the Created field value
func (o *TransactionDispute) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TransactionDispute) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TransactionDispute) SetCreated(v time.Time) {
	o.Created = v
}

// GetExternalCaseReference returns the ExternalCaseReference field value if set, zero value otherwise.
func (o *TransactionDispute) GetExternalCaseReference() string {
	if o == nil || IsNil(o.ExternalCaseReference) {
		var ret string
		return ret
	}
	return *o.ExternalCaseReference
}

// GetExternalCaseReferenceOk returns a tuple with the ExternalCaseReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDispute) GetExternalCaseReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalCaseReference) {
		return nil, false
	}
	return o.ExternalCaseReference, true
}

// HasExternalCaseReference returns a boolean if a field has been set.
func (o *TransactionDispute) HasExternalCaseReference() bool {
	if o != nil && !IsNil(o.ExternalCaseReference) {
		return true
	}

	return false
}

// SetExternalCaseReference gets a reference to the given string and assigns it to the ExternalCaseReference field.
func (o *TransactionDispute) SetExternalCaseReference(v string) {
	o.ExternalCaseReference = &v
}

// GetId returns the Id field value
func (o *TransactionDispute) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionDispute) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionDispute) SetId(v string) {
	o.Id = v
}

// GetInternalCaseReference returns the InternalCaseReference field value if set, zero value otherwise.
func (o *TransactionDispute) GetInternalCaseReference() string {
	if o == nil || IsNil(o.InternalCaseReference) {
		var ret string
		return ret
	}
	return *o.InternalCaseReference
}

// GetInternalCaseReferenceOk returns a tuple with the InternalCaseReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionDispute) GetInternalCaseReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.InternalCaseReference) {
		return nil, false
	}
	return o.InternalCaseReference, true
}

// HasInternalCaseReference returns a boolean if a field has been set.
func (o *TransactionDispute) HasInternalCaseReference() bool {
	if o != nil && !IsNil(o.InternalCaseReference) {
		return true
	}

	return false
}

// SetInternalCaseReference gets a reference to the given string and assigns it to the InternalCaseReference field.
func (o *TransactionDispute) SetInternalCaseReference(v string) {
	o.InternalCaseReference = &v
}

// GetStatus returns the Status field value
func (o *TransactionDispute) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TransactionDispute) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TransactionDispute) SetStatus(v string) {
	o.Status = v
}

// GetUpdated returns the Updated field value
func (o *TransactionDispute) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *TransactionDispute) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *TransactionDispute) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o TransactionDispute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDispute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	if !IsNil(o.ExternalCaseReference) {
		toSerialize["external_case_reference"] = o.ExternalCaseReference
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.InternalCaseReference) {
		toSerialize["internal_case_reference"] = o.InternalCaseReference
	}
	toSerialize["status"] = o.Status
	toSerialize["updated"] = o.Updated
	return toSerialize, nil
}

func (o *TransactionDispute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"id",
		"status",
		"updated",
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

	varTransactionDispute := _TransactionDispute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionDispute)

	if err != nil {
		return err
	}

	*o = TransactionDispute(varTransactionDispute)

	return err
}

type NullableTransactionDispute struct {
	value *TransactionDispute
	isSet bool
}

func (v NullableTransactionDispute) Get() *TransactionDispute {
	return v.value
}

func (v *NullableTransactionDispute) Set(val *TransactionDispute) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDispute) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDispute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDispute(val *TransactionDispute) *NullableTransactionDispute {
	return &NullableTransactionDispute{value: val, isSet: true}
}

func (v NullableTransactionDispute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDispute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
