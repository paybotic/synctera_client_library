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

// checks if the TransactionOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionOptions{}

// TransactionOptions struct for TransactionOptions
type TransactionOptions struct {
	AdditionalData                     *string `json:"additional_data,omitempty"`
	CardExpirationDateYymm             *string `json:"card_expiration_date_yymm,omitempty"`
	DatabaseTransactionTimeout         *int32  `json:"database_transaction_timeout,omitempty"`
	EncryptionKeyId                    *string `json:"encryption_key_id,omitempty"`
	IsAsync                            *bool   `json:"is_async,omitempty"`
	PreAuthTimeLimit                   *string `json:"pre_auth_time_limit,omitempty"`
	SendExpirationDate                 *bool   `json:"send_expiration_date,omitempty"`
	SendTrackData                      *bool   `json:"send_track_data,omitempty"`
	TransactionId                      *string `json:"transaction_id,omitempty"`
	TransactionTimeoutThresholdSeconds *int64  `json:"transaction_timeout_threshold_seconds,omitempty"`
	AdditionalProperties               map[string]interface{}
}

type _TransactionOptions TransactionOptions

// NewTransactionOptions instantiates a new TransactionOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionOptions() *TransactionOptions {
	this := TransactionOptions{}
	var isAsync bool = false
	this.IsAsync = &isAsync
	var sendExpirationDate bool = false
	this.SendExpirationDate = &sendExpirationDate
	var sendTrackData bool = false
	this.SendTrackData = &sendTrackData
	return &this
}

// NewTransactionOptionsWithDefaults instantiates a new TransactionOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionOptionsWithDefaults() *TransactionOptions {
	this := TransactionOptions{}
	var isAsync bool = false
	this.IsAsync = &isAsync
	var sendExpirationDate bool = false
	this.SendExpirationDate = &sendExpirationDate
	var sendTrackData bool = false
	this.SendTrackData = &sendTrackData
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *TransactionOptions) GetAdditionalData() string {
	if o == nil || IsNil(o.AdditionalData) {
		var ret string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetAdditionalDataOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *TransactionOptions) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given string and assigns it to the AdditionalData field.
func (o *TransactionOptions) SetAdditionalData(v string) {
	o.AdditionalData = &v
}

// GetCardExpirationDateYymm returns the CardExpirationDateYymm field value if set, zero value otherwise.
func (o *TransactionOptions) GetCardExpirationDateYymm() string {
	if o == nil || IsNil(o.CardExpirationDateYymm) {
		var ret string
		return ret
	}
	return *o.CardExpirationDateYymm
}

// GetCardExpirationDateYymmOk returns a tuple with the CardExpirationDateYymm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetCardExpirationDateYymmOk() (*string, bool) {
	if o == nil || IsNil(o.CardExpirationDateYymm) {
		return nil, false
	}
	return o.CardExpirationDateYymm, true
}

// HasCardExpirationDateYymm returns a boolean if a field has been set.
func (o *TransactionOptions) HasCardExpirationDateYymm() bool {
	if o != nil && !IsNil(o.CardExpirationDateYymm) {
		return true
	}

	return false
}

// SetCardExpirationDateYymm gets a reference to the given string and assigns it to the CardExpirationDateYymm field.
func (o *TransactionOptions) SetCardExpirationDateYymm(v string) {
	o.CardExpirationDateYymm = &v
}

// GetDatabaseTransactionTimeout returns the DatabaseTransactionTimeout field value if set, zero value otherwise.
func (o *TransactionOptions) GetDatabaseTransactionTimeout() int32 {
	if o == nil || IsNil(o.DatabaseTransactionTimeout) {
		var ret int32
		return ret
	}
	return *o.DatabaseTransactionTimeout
}

// GetDatabaseTransactionTimeoutOk returns a tuple with the DatabaseTransactionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetDatabaseTransactionTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.DatabaseTransactionTimeout) {
		return nil, false
	}
	return o.DatabaseTransactionTimeout, true
}

// HasDatabaseTransactionTimeout returns a boolean if a field has been set.
func (o *TransactionOptions) HasDatabaseTransactionTimeout() bool {
	if o != nil && !IsNil(o.DatabaseTransactionTimeout) {
		return true
	}

	return false
}

// SetDatabaseTransactionTimeout gets a reference to the given int32 and assigns it to the DatabaseTransactionTimeout field.
func (o *TransactionOptions) SetDatabaseTransactionTimeout(v int32) {
	o.DatabaseTransactionTimeout = &v
}

// GetEncryptionKeyId returns the EncryptionKeyId field value if set, zero value otherwise.
func (o *TransactionOptions) GetEncryptionKeyId() string {
	if o == nil || IsNil(o.EncryptionKeyId) {
		var ret string
		return ret
	}
	return *o.EncryptionKeyId
}

// GetEncryptionKeyIdOk returns a tuple with the EncryptionKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetEncryptionKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionKeyId) {
		return nil, false
	}
	return o.EncryptionKeyId, true
}

// HasEncryptionKeyId returns a boolean if a field has been set.
func (o *TransactionOptions) HasEncryptionKeyId() bool {
	if o != nil && !IsNil(o.EncryptionKeyId) {
		return true
	}

	return false
}

// SetEncryptionKeyId gets a reference to the given string and assigns it to the EncryptionKeyId field.
func (o *TransactionOptions) SetEncryptionKeyId(v string) {
	o.EncryptionKeyId = &v
}

// GetIsAsync returns the IsAsync field value if set, zero value otherwise.
func (o *TransactionOptions) GetIsAsync() bool {
	if o == nil || IsNil(o.IsAsync) {
		var ret bool
		return ret
	}
	return *o.IsAsync
}

// GetIsAsyncOk returns a tuple with the IsAsync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetIsAsyncOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAsync) {
		return nil, false
	}
	return o.IsAsync, true
}

// HasIsAsync returns a boolean if a field has been set.
func (o *TransactionOptions) HasIsAsync() bool {
	if o != nil && !IsNil(o.IsAsync) {
		return true
	}

	return false
}

// SetIsAsync gets a reference to the given bool and assigns it to the IsAsync field.
func (o *TransactionOptions) SetIsAsync(v bool) {
	o.IsAsync = &v
}

// GetPreAuthTimeLimit returns the PreAuthTimeLimit field value if set, zero value otherwise.
func (o *TransactionOptions) GetPreAuthTimeLimit() string {
	if o == nil || IsNil(o.PreAuthTimeLimit) {
		var ret string
		return ret
	}
	return *o.PreAuthTimeLimit
}

// GetPreAuthTimeLimitOk returns a tuple with the PreAuthTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetPreAuthTimeLimitOk() (*string, bool) {
	if o == nil || IsNil(o.PreAuthTimeLimit) {
		return nil, false
	}
	return o.PreAuthTimeLimit, true
}

// HasPreAuthTimeLimit returns a boolean if a field has been set.
func (o *TransactionOptions) HasPreAuthTimeLimit() bool {
	if o != nil && !IsNil(o.PreAuthTimeLimit) {
		return true
	}

	return false
}

// SetPreAuthTimeLimit gets a reference to the given string and assigns it to the PreAuthTimeLimit field.
func (o *TransactionOptions) SetPreAuthTimeLimit(v string) {
	o.PreAuthTimeLimit = &v
}

// GetSendExpirationDate returns the SendExpirationDate field value if set, zero value otherwise.
func (o *TransactionOptions) GetSendExpirationDate() bool {
	if o == nil || IsNil(o.SendExpirationDate) {
		var ret bool
		return ret
	}
	return *o.SendExpirationDate
}

// GetSendExpirationDateOk returns a tuple with the SendExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetSendExpirationDateOk() (*bool, bool) {
	if o == nil || IsNil(o.SendExpirationDate) {
		return nil, false
	}
	return o.SendExpirationDate, true
}

// HasSendExpirationDate returns a boolean if a field has been set.
func (o *TransactionOptions) HasSendExpirationDate() bool {
	if o != nil && !IsNil(o.SendExpirationDate) {
		return true
	}

	return false
}

// SetSendExpirationDate gets a reference to the given bool and assigns it to the SendExpirationDate field.
func (o *TransactionOptions) SetSendExpirationDate(v bool) {
	o.SendExpirationDate = &v
}

// GetSendTrackData returns the SendTrackData field value if set, zero value otherwise.
func (o *TransactionOptions) GetSendTrackData() bool {
	if o == nil || IsNil(o.SendTrackData) {
		var ret bool
		return ret
	}
	return *o.SendTrackData
}

// GetSendTrackDataOk returns a tuple with the SendTrackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetSendTrackDataOk() (*bool, bool) {
	if o == nil || IsNil(o.SendTrackData) {
		return nil, false
	}
	return o.SendTrackData, true
}

// HasSendTrackData returns a boolean if a field has been set.
func (o *TransactionOptions) HasSendTrackData() bool {
	if o != nil && !IsNil(o.SendTrackData) {
		return true
	}

	return false
}

// SetSendTrackData gets a reference to the given bool and assigns it to the SendTrackData field.
func (o *TransactionOptions) SetSendTrackData(v bool) {
	o.SendTrackData = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *TransactionOptions) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *TransactionOptions) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *TransactionOptions) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionTimeoutThresholdSeconds returns the TransactionTimeoutThresholdSeconds field value if set, zero value otherwise.
func (o *TransactionOptions) GetTransactionTimeoutThresholdSeconds() int64 {
	if o == nil || IsNil(o.TransactionTimeoutThresholdSeconds) {
		var ret int64
		return ret
	}
	return *o.TransactionTimeoutThresholdSeconds
}

// GetTransactionTimeoutThresholdSecondsOk returns a tuple with the TransactionTimeoutThresholdSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptions) GetTransactionTimeoutThresholdSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionTimeoutThresholdSeconds) {
		return nil, false
	}
	return o.TransactionTimeoutThresholdSeconds, true
}

// HasTransactionTimeoutThresholdSeconds returns a boolean if a field has been set.
func (o *TransactionOptions) HasTransactionTimeoutThresholdSeconds() bool {
	if o != nil && !IsNil(o.TransactionTimeoutThresholdSeconds) {
		return true
	}

	return false
}

// SetTransactionTimeoutThresholdSeconds gets a reference to the given int64 and assigns it to the TransactionTimeoutThresholdSeconds field.
func (o *TransactionOptions) SetTransactionTimeoutThresholdSeconds(v int64) {
	o.TransactionTimeoutThresholdSeconds = &v
}

func (o TransactionOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalData) {
		toSerialize["additional_data"] = o.AdditionalData
	}
	if !IsNil(o.CardExpirationDateYymm) {
		toSerialize["card_expiration_date_yymm"] = o.CardExpirationDateYymm
	}
	if !IsNil(o.DatabaseTransactionTimeout) {
		toSerialize["database_transaction_timeout"] = o.DatabaseTransactionTimeout
	}
	if !IsNil(o.EncryptionKeyId) {
		toSerialize["encryption_key_id"] = o.EncryptionKeyId
	}
	if !IsNil(o.IsAsync) {
		toSerialize["is_async"] = o.IsAsync
	}
	if !IsNil(o.PreAuthTimeLimit) {
		toSerialize["pre_auth_time_limit"] = o.PreAuthTimeLimit
	}
	if !IsNil(o.SendExpirationDate) {
		toSerialize["send_expiration_date"] = o.SendExpirationDate
	}
	if !IsNil(o.SendTrackData) {
		toSerialize["send_track_data"] = o.SendTrackData
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !IsNil(o.TransactionTimeoutThresholdSeconds) {
		toSerialize["transaction_timeout_threshold_seconds"] = o.TransactionTimeoutThresholdSeconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransactionOptions) UnmarshalJSON(data []byte) (err error) {
	varTransactionOptions := _TransactionOptions{}

	err = json.Unmarshal(data, &varTransactionOptions)

	if err != nil {
		return err
	}

	*o = TransactionOptions(varTransactionOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_data")
		delete(additionalProperties, "card_expiration_date_yymm")
		delete(additionalProperties, "database_transaction_timeout")
		delete(additionalProperties, "encryption_key_id")
		delete(additionalProperties, "is_async")
		delete(additionalProperties, "pre_auth_time_limit")
		delete(additionalProperties, "send_expiration_date")
		delete(additionalProperties, "send_track_data")
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "transaction_timeout_threshold_seconds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionOptions struct {
	value *TransactionOptions
	isSet bool
}

func (v NullableTransactionOptions) Get() *TransactionOptions {
	return v.value
}

func (v *NullableTransactionOptions) Set(val *TransactionOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionOptions(val *TransactionOptions) *NullableTransactionOptions {
	return &NullableTransactionOptions{value: val, isSet: true}
}

func (v NullableTransactionOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
