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

// checks if the StopPaymentResponseWAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopPaymentResponseWAccount{}

// StopPaymentResponseWAccount created stop payment
type StopPaymentResponseWAccount struct {
	// ID of the dispute that created the stop payment
	DisputeId *string `json:"dispute_id,omitempty"`
	// The date when this stop payment is no longer valid. This is only for business accounts.
	ExpiresOn *time.Time `json:"expires_on,omitempty"`
	// Name of the originator
	OriginatorName string `json:"originator_name"`
	StopPaymentId  string `json:"stop_payment_id"`
	// If this stop payment was created from a disputed transaction, transaction_id references the posted transaction.
	TransactionId *string `json:"transaction_id,omitempty"`
	AccountId     *string `json:"account_id,omitempty"`
	// Timestamp when the stop payment was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Timestamp when stop payment was updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant *string `json:"tenant,omitempty"`
}

type _StopPaymentResponseWAccount StopPaymentResponseWAccount

// NewStopPaymentResponseWAccount instantiates a new StopPaymentResponseWAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopPaymentResponseWAccount(originatorName string, stopPaymentId string) *StopPaymentResponseWAccount {
	this := StopPaymentResponseWAccount{}
	this.OriginatorName = originatorName
	this.StopPaymentId = stopPaymentId
	return &this
}

// NewStopPaymentResponseWAccountWithDefaults instantiates a new StopPaymentResponseWAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopPaymentResponseWAccountWithDefaults() *StopPaymentResponseWAccount {
	this := StopPaymentResponseWAccount{}
	return &this
}

// GetDisputeId returns the DisputeId field value if set, zero value otherwise.
func (o *StopPaymentResponseWAccount) GetDisputeId() string {
	if o == nil || IsNil(o.DisputeId) {
		var ret string
		return ret
	}
	return *o.DisputeId
}

// GetDisputeIdOk returns a tuple with the DisputeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPaymentResponseWAccount) GetDisputeIdOk() (*string, bool) {
	if o == nil || IsNil(o.DisputeId) {
		return nil, false
	}
	return o.DisputeId, true
}

// HasDisputeId returns a boolean if a field has been set.
func (o *StopPaymentResponseWAccount) HasDisputeId() bool {
	if o != nil && !IsNil(o.DisputeId) {
		return true
	}

	return false
}

// SetDisputeId gets a reference to the given string and assigns it to the DisputeId field.
func (o *StopPaymentResponseWAccount) SetDisputeId(v string) {
	o.DisputeId = &v
}

// GetExpiresOn returns the ExpiresOn field value if set, zero value otherwise.
func (o *StopPaymentResponseWAccount) GetExpiresOn() time.Time {
	if o == nil || IsNil(o.ExpiresOn) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPaymentResponseWAccount) GetExpiresOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresOn) {
		return nil, false
	}
	return o.ExpiresOn, true
}

// HasExpiresOn returns a boolean if a field has been set.
func (o *StopPaymentResponseWAccount) HasExpiresOn() bool {
	if o != nil && !IsNil(o.ExpiresOn) {
		return true
	}

	return false
}

// SetExpiresOn gets a reference to the given time.Time and assigns it to the ExpiresOn field.
func (o *StopPaymentResponseWAccount) SetExpiresOn(v time.Time) {
	o.ExpiresOn = &v
}

// GetOriginatorName returns the OriginatorName field value
func (o *StopPaymentResponseWAccount) GetOriginatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatorName
}

// GetOriginatorNameOk returns a tuple with the OriginatorName field value
// and a boolean to check if the value has been set.
func (o *StopPaymentResponseWAccount) GetOriginatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatorName, true
}

// SetOriginatorName sets field value
func (o *StopPaymentResponseWAccount) SetOriginatorName(v string) {
	o.OriginatorName = v
}

// GetStopPaymentId returns the StopPaymentId field value
func (o *StopPaymentResponseWAccount) GetStopPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StopPaymentId
}

// GetStopPaymentIdOk returns a tuple with the StopPaymentId field value
// and a boolean to check if the value has been set.
func (o *StopPaymentResponseWAccount) GetStopPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopPaymentId, true
}

// SetStopPaymentId sets field value
func (o *StopPaymentResponseWAccount) SetStopPaymentId(v string) {
	o.StopPaymentId = v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *StopPaymentResponseWAccount) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPaymentResponseWAccount) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *StopPaymentResponseWAccount) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *StopPaymentResponseWAccount) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *StopPaymentResponseWAccount) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPaymentResponseWAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *StopPaymentResponseWAccount) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *StopPaymentResponseWAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *StopPaymentResponseWAccount) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPaymentResponseWAccount) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *StopPaymentResponseWAccount) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *StopPaymentResponseWAccount) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *StopPaymentResponseWAccount) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPaymentResponseWAccount) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *StopPaymentResponseWAccount) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *StopPaymentResponseWAccount) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StopPaymentResponseWAccount) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopPaymentResponseWAccount) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StopPaymentResponseWAccount) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *StopPaymentResponseWAccount) SetTenant(v string) {
	o.Tenant = &v
}

func (o StopPaymentResponseWAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopPaymentResponseWAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisputeId) {
		toSerialize["dispute_id"] = o.DisputeId
	}
	if !IsNil(o.ExpiresOn) {
		toSerialize["expires_on"] = o.ExpiresOn
	}
	toSerialize["originator_name"] = o.OriginatorName
	toSerialize["stop_payment_id"] = o.StopPaymentId
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	return toSerialize, nil
}

func (o *StopPaymentResponseWAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"originator_name",
		"stop_payment_id",
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

	varStopPaymentResponseWAccount := _StopPaymentResponseWAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStopPaymentResponseWAccount)

	if err != nil {
		return err
	}

	*o = StopPaymentResponseWAccount(varStopPaymentResponseWAccount)

	return err
}

type NullableStopPaymentResponseWAccount struct {
	value *StopPaymentResponseWAccount
	isSet bool
}

func (v NullableStopPaymentResponseWAccount) Get() *StopPaymentResponseWAccount {
	return v.value
}

func (v *NullableStopPaymentResponseWAccount) Set(val *StopPaymentResponseWAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableStopPaymentResponseWAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableStopPaymentResponseWAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopPaymentResponseWAccount(val *StopPaymentResponseWAccount) *NullableStopPaymentResponseWAccount {
	return &NullableStopPaymentResponseWAccount{value: val, isSet: true}
}

func (v NullableStopPaymentResponseWAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopPaymentResponseWAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
