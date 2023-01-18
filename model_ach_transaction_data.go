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

// AchTransactionData Transaction metadata specific to transactions with type `ACH`
type AchTransactionData struct {
	// The ACH company entry description
	Description *string `json:"description,omitempty"`
	// For outgoing ACH, the name of the file the entry went out in. For incoming ACH, the name of the file the entry came from. The value will be omitted for outgoing payments that have not yet been written into a file.
	FileName *string `json:"file_name,omitempty"`
	// The ACH payment uuid (used in `/v0/ach` endpoint)
	Id string `json:"id"`
	// The uuid of the account originating the ACH. This will be a customer account uuid for outgoing ACH, and omitted for incoming ACH.
	OriginatingAccountId *string `json:"originating_account_id,omitempty"`
	// The name of the originator according to the ACH entry. This should map to the ACH `Identification Number` field, if provided.
	OriginatorName *string `json:"originator_name,omitempty"`
	// The uuid of the account receiving the ACH entry. In the case of an outgoing ACH, this will be an external_account uuid. For incoming ACH, this will be an account uuid.
	ReceivingAccountId *string `json:"receiving_account_id,omitempty"`
	// The name of the recipient according to the ACH entry. This should map to the ACH `Individual Name` field.
	RecipientName *string `json:"recipient_name,omitempty"`
	// The ACH return code, if this transaction was a return
	ReturnCode *string `json:"return_code,omitempty"`
	// The ACH trace number associated with the transaction
	TraceNumber *string `json:"trace_number,omitempty"`
}

// NewAchTransactionData instantiates a new AchTransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchTransactionData(id string) *AchTransactionData {
	this := AchTransactionData{}
	this.Id = id
	return &this
}

// NewAchTransactionDataWithDefaults instantiates a new AchTransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchTransactionDataWithDefaults() *AchTransactionData {
	this := AchTransactionData{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AchTransactionData) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchTransactionData) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AchTransactionData) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AchTransactionData) SetDescription(v string) {
	o.Description = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AchTransactionData) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchTransactionData) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AchTransactionData) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AchTransactionData) SetFileName(v string) {
	o.FileName = &v
}

// GetId returns the Id field value
func (o *AchTransactionData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AchTransactionData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AchTransactionData) SetId(v string) {
	o.Id = v
}

// GetOriginatingAccountId returns the OriginatingAccountId field value if set, zero value otherwise.
func (o *AchTransactionData) GetOriginatingAccountId() string {
	if o == nil || o.OriginatingAccountId == nil {
		var ret string
		return ret
	}
	return *o.OriginatingAccountId
}

// GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchTransactionData) GetOriginatingAccountIdOk() (*string, bool) {
	if o == nil || o.OriginatingAccountId == nil {
		return nil, false
	}
	return o.OriginatingAccountId, true
}

// HasOriginatingAccountId returns a boolean if a field has been set.
func (o *AchTransactionData) HasOriginatingAccountId() bool {
	if o != nil && o.OriginatingAccountId != nil {
		return true
	}

	return false
}

// SetOriginatingAccountId gets a reference to the given string and assigns it to the OriginatingAccountId field.
func (o *AchTransactionData) SetOriginatingAccountId(v string) {
	o.OriginatingAccountId = &v
}

// GetOriginatorName returns the OriginatorName field value if set, zero value otherwise.
func (o *AchTransactionData) GetOriginatorName() string {
	if o == nil || o.OriginatorName == nil {
		var ret string
		return ret
	}
	return *o.OriginatorName
}

// GetOriginatorNameOk returns a tuple with the OriginatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchTransactionData) GetOriginatorNameOk() (*string, bool) {
	if o == nil || o.OriginatorName == nil {
		return nil, false
	}
	return o.OriginatorName, true
}

// HasOriginatorName returns a boolean if a field has been set.
func (o *AchTransactionData) HasOriginatorName() bool {
	if o != nil && o.OriginatorName != nil {
		return true
	}

	return false
}

// SetOriginatorName gets a reference to the given string and assigns it to the OriginatorName field.
func (o *AchTransactionData) SetOriginatorName(v string) {
	o.OriginatorName = &v
}

// GetReceivingAccountId returns the ReceivingAccountId field value if set, zero value otherwise.
func (o *AchTransactionData) GetReceivingAccountId() string {
	if o == nil || o.ReceivingAccountId == nil {
		var ret string
		return ret
	}
	return *o.ReceivingAccountId
}

// GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchTransactionData) GetReceivingAccountIdOk() (*string, bool) {
	if o == nil || o.ReceivingAccountId == nil {
		return nil, false
	}
	return o.ReceivingAccountId, true
}

// HasReceivingAccountId returns a boolean if a field has been set.
func (o *AchTransactionData) HasReceivingAccountId() bool {
	if o != nil && o.ReceivingAccountId != nil {
		return true
	}

	return false
}

// SetReceivingAccountId gets a reference to the given string and assigns it to the ReceivingAccountId field.
func (o *AchTransactionData) SetReceivingAccountId(v string) {
	o.ReceivingAccountId = &v
}

// GetRecipientName returns the RecipientName field value if set, zero value otherwise.
func (o *AchTransactionData) GetRecipientName() string {
	if o == nil || o.RecipientName == nil {
		var ret string
		return ret
	}
	return *o.RecipientName
}

// GetRecipientNameOk returns a tuple with the RecipientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchTransactionData) GetRecipientNameOk() (*string, bool) {
	if o == nil || o.RecipientName == nil {
		return nil, false
	}
	return o.RecipientName, true
}

// HasRecipientName returns a boolean if a field has been set.
func (o *AchTransactionData) HasRecipientName() bool {
	if o != nil && o.RecipientName != nil {
		return true
	}

	return false
}

// SetRecipientName gets a reference to the given string and assigns it to the RecipientName field.
func (o *AchTransactionData) SetRecipientName(v string) {
	o.RecipientName = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *AchTransactionData) GetReturnCode() string {
	if o == nil || o.ReturnCode == nil {
		var ret string
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchTransactionData) GetReturnCodeOk() (*string, bool) {
	if o == nil || o.ReturnCode == nil {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *AchTransactionData) HasReturnCode() bool {
	if o != nil && o.ReturnCode != nil {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given string and assigns it to the ReturnCode field.
func (o *AchTransactionData) SetReturnCode(v string) {
	o.ReturnCode = &v
}

// GetTraceNumber returns the TraceNumber field value if set, zero value otherwise.
func (o *AchTransactionData) GetTraceNumber() string {
	if o == nil || o.TraceNumber == nil {
		var ret string
		return ret
	}
	return *o.TraceNumber
}

// GetTraceNumberOk returns a tuple with the TraceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchTransactionData) GetTraceNumberOk() (*string, bool) {
	if o == nil || o.TraceNumber == nil {
		return nil, false
	}
	return o.TraceNumber, true
}

// HasTraceNumber returns a boolean if a field has been set.
func (o *AchTransactionData) HasTraceNumber() bool {
	if o != nil && o.TraceNumber != nil {
		return true
	}

	return false
}

// SetTraceNumber gets a reference to the given string and assigns it to the TraceNumber field.
func (o *AchTransactionData) SetTraceNumber(v string) {
	o.TraceNumber = &v
}

func (o AchTransactionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FileName != nil {
		toSerialize["file_name"] = o.FileName
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.OriginatingAccountId != nil {
		toSerialize["originating_account_id"] = o.OriginatingAccountId
	}
	if o.OriginatorName != nil {
		toSerialize["originator_name"] = o.OriginatorName
	}
	if o.ReceivingAccountId != nil {
		toSerialize["receiving_account_id"] = o.ReceivingAccountId
	}
	if o.RecipientName != nil {
		toSerialize["recipient_name"] = o.RecipientName
	}
	if o.ReturnCode != nil {
		toSerialize["return_code"] = o.ReturnCode
	}
	if o.TraceNumber != nil {
		toSerialize["trace_number"] = o.TraceNumber
	}
	return json.Marshal(toSerialize)
}

type NullableAchTransactionData struct {
	value *AchTransactionData
	isSet bool
}

func (v NullableAchTransactionData) Get() *AchTransactionData {
	return v.value
}

func (v *NullableAchTransactionData) Set(val *AchTransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullableAchTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullableAchTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchTransactionData(val *AchTransactionData) *NullableAchTransactionData {
	return &NullableAchTransactionData{value: val, isSet: true}
}

func (v NullableAchTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


