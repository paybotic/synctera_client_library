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
	"time"
)

// checks if the CustomerVerificationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerVerificationResult{}

// CustomerVerificationResult Verification result
type CustomerVerificationResult struct {
	// Unique ID for this verification result.
	Id *string `json:"id,omitempty"`
	// List of potential problems found. These are subject to change.
	// Deprecated
	Issues []string `json:"issues,omitempty"`
	// Deprecated
	RawResponse *RawResponse `json:"raw_response,omitempty"`
	// The determination of this verification.
	Result     string                  `json:"result"`
	VendorInfo *VerificationVendorInfo `json:"vendor_info,omitempty"`
	// The date and time the verification was completed.
	VerificationTime time.Time           `json:"verification_time"`
	VerificationType KycVerificationType `json:"verification_type"`
}

type _CustomerVerificationResult CustomerVerificationResult

// NewCustomerVerificationResult instantiates a new CustomerVerificationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerVerificationResult(result string, verificationTime time.Time, verificationType KycVerificationType) *CustomerVerificationResult {
	this := CustomerVerificationResult{}
	this.Result = result
	this.VerificationTime = verificationTime
	this.VerificationType = verificationType
	return &this
}

// NewCustomerVerificationResultWithDefaults instantiates a new CustomerVerificationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerVerificationResultWithDefaults() *CustomerVerificationResult {
	this := CustomerVerificationResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerVerificationResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerVerificationResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerVerificationResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerVerificationResult) SetId(v string) {
	o.Id = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
// Deprecated
func (o *CustomerVerificationResult) GetIssues() []string {
	if o == nil || IsNil(o.Issues) {
		var ret []string
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CustomerVerificationResult) GetIssuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *CustomerVerificationResult) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []string and assigns it to the Issues field.
// Deprecated
func (o *CustomerVerificationResult) SetIssues(v []string) {
	o.Issues = v
}

// GetRawResponse returns the RawResponse field value if set, zero value otherwise.
// Deprecated
func (o *CustomerVerificationResult) GetRawResponse() RawResponse {
	if o == nil || IsNil(o.RawResponse) {
		var ret RawResponse
		return ret
	}
	return *o.RawResponse
}

// GetRawResponseOk returns a tuple with the RawResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CustomerVerificationResult) GetRawResponseOk() (*RawResponse, bool) {
	if o == nil || IsNil(o.RawResponse) {
		return nil, false
	}
	return o.RawResponse, true
}

// HasRawResponse returns a boolean if a field has been set.
func (o *CustomerVerificationResult) HasRawResponse() bool {
	if o != nil && !IsNil(o.RawResponse) {
		return true
	}

	return false
}

// SetRawResponse gets a reference to the given RawResponse and assigns it to the RawResponse field.
// Deprecated
func (o *CustomerVerificationResult) SetRawResponse(v RawResponse) {
	o.RawResponse = &v
}

// GetResult returns the Result field value
func (o *CustomerVerificationResult) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CustomerVerificationResult) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CustomerVerificationResult) SetResult(v string) {
	o.Result = v
}

// GetVendorInfo returns the VendorInfo field value if set, zero value otherwise.
func (o *CustomerVerificationResult) GetVendorInfo() VerificationVendorInfo {
	if o == nil || IsNil(o.VendorInfo) {
		var ret VerificationVendorInfo
		return ret
	}
	return *o.VendorInfo
}

// GetVendorInfoOk returns a tuple with the VendorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerVerificationResult) GetVendorInfoOk() (*VerificationVendorInfo, bool) {
	if o == nil || IsNil(o.VendorInfo) {
		return nil, false
	}
	return o.VendorInfo, true
}

// HasVendorInfo returns a boolean if a field has been set.
func (o *CustomerVerificationResult) HasVendorInfo() bool {
	if o != nil && !IsNil(o.VendorInfo) {
		return true
	}

	return false
}

// SetVendorInfo gets a reference to the given VerificationVendorInfo and assigns it to the VendorInfo field.
func (o *CustomerVerificationResult) SetVendorInfo(v VerificationVendorInfo) {
	o.VendorInfo = &v
}

// GetVerificationTime returns the VerificationTime field value
func (o *CustomerVerificationResult) GetVerificationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.VerificationTime
}

// GetVerificationTimeOk returns a tuple with the VerificationTime field value
// and a boolean to check if the value has been set.
func (o *CustomerVerificationResult) GetVerificationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationTime, true
}

// SetVerificationTime sets field value
func (o *CustomerVerificationResult) SetVerificationTime(v time.Time) {
	o.VerificationTime = v
}

// GetVerificationType returns the VerificationType field value
func (o *CustomerVerificationResult) GetVerificationType() KycVerificationType {
	if o == nil {
		var ret KycVerificationType
		return ret
	}

	return o.VerificationType
}

// GetVerificationTypeOk returns a tuple with the VerificationType field value
// and a boolean to check if the value has been set.
func (o *CustomerVerificationResult) GetVerificationTypeOk() (*KycVerificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationType, true
}

// SetVerificationType sets field value
func (o *CustomerVerificationResult) SetVerificationType(v KycVerificationType) {
	o.VerificationType = v
}

func (o CustomerVerificationResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerVerificationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.RawResponse) {
		toSerialize["raw_response"] = o.RawResponse
	}
	toSerialize["result"] = o.Result
	if !IsNil(o.VendorInfo) {
		toSerialize["vendor_info"] = o.VendorInfo
	}
	toSerialize["verification_time"] = o.VerificationTime
	toSerialize["verification_type"] = o.VerificationType
	return toSerialize, nil
}

func (o *CustomerVerificationResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
		"verification_time",
		"verification_type",
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

	varCustomerVerificationResult := _CustomerVerificationResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerVerificationResult)

	if err != nil {
		return err
	}

	*o = CustomerVerificationResult(varCustomerVerificationResult)

	return err
}

type NullableCustomerVerificationResult struct {
	value *CustomerVerificationResult
	isSet bool
}

func (v NullableCustomerVerificationResult) Get() *CustomerVerificationResult {
	return v.value
}

func (v *NullableCustomerVerificationResult) Set(val *CustomerVerificationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerVerificationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerVerificationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerVerificationResult(val *CustomerVerificationResult) *NullableCustomerVerificationResult {
	return &NullableCustomerVerificationResult{value: val, isSet: true}
}

func (v NullableCustomerVerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerVerificationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
