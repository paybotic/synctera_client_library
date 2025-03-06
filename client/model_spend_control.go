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

// checks if the SpendControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpendControl{}

// SpendControl struct for SpendControl
type SpendControl struct {
	// If set, create a case for transactions that do not conform to the spend control
	ActionCase bool `json:"action_case"`
	// If set, decline transactions that do not conform to the spend control
	ActionDecline bool `json:"action_decline"`
	// Monetary limit for the spend control in the smallest currency unit (eg cents)
	AmountLimit int64 `json:"amount_limit"`
	// The timestamp representing when the spend control was created
	CreationTime *time.Time             `json:"creation_time,omitempty"`
	Direction    *SpendControlDirection `json:"direction,omitempty"`
	// Spend Control ID
	Id *string `json:"id,omitempty"`
	// Indicates if spend control is active
	IsActive bool `json:"is_active"`
	// The timestamp representing when the spend control was last modified
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// merchant category codes for spend control
	MerchantCategoryCodes []string `json:"merchant_category_codes,omitempty"`
	// Name assigned to spend control
	Name string `json:"name"`
	// A list of payment sub-types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all sub-types.
	PaymentSubTypes []PaymentSubType `json:"payment_sub_types,omitempty"`
	// A list of payment types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all types of payments.
	PaymentTypes []PaymentType         `json:"payment_types,omitempty"`
	TimeRange    SpendControlTimeRange `json:"time_range"`
}

type _SpendControl SpendControl

// NewSpendControl instantiates a new SpendControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendControl(actionCase bool, actionDecline bool, amountLimit int64, isActive bool, name string, timeRange SpendControlTimeRange) *SpendControl {
	this := SpendControl{}
	this.ActionCase = actionCase
	this.ActionDecline = actionDecline
	this.AmountLimit = amountLimit
	this.IsActive = isActive
	this.Name = name
	this.TimeRange = timeRange
	return &this
}

// NewSpendControlWithDefaults instantiates a new SpendControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendControlWithDefaults() *SpendControl {
	this := SpendControl{}
	return &this
}

// GetActionCase returns the ActionCase field value
func (o *SpendControl) GetActionCase() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ActionCase
}

// GetActionCaseOk returns a tuple with the ActionCase field value
// and a boolean to check if the value has been set.
func (o *SpendControl) GetActionCaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionCase, true
}

// SetActionCase sets field value
func (o *SpendControl) SetActionCase(v bool) {
	o.ActionCase = v
}

// GetActionDecline returns the ActionDecline field value
func (o *SpendControl) GetActionDecline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ActionDecline
}

// GetActionDeclineOk returns a tuple with the ActionDecline field value
// and a boolean to check if the value has been set.
func (o *SpendControl) GetActionDeclineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionDecline, true
}

// SetActionDecline sets field value
func (o *SpendControl) SetActionDecline(v bool) {
	o.ActionDecline = v
}

// GetAmountLimit returns the AmountLimit field value
func (o *SpendControl) GetAmountLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AmountLimit
}

// GetAmountLimitOk returns a tuple with the AmountLimit field value
// and a boolean to check if the value has been set.
func (o *SpendControl) GetAmountLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountLimit, true
}

// SetAmountLimit sets field value
func (o *SpendControl) SetAmountLimit(v int64) {
	o.AmountLimit = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *SpendControl) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControl) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *SpendControl) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *SpendControl) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *SpendControl) GetDirection() SpendControlDirection {
	if o == nil || IsNil(o.Direction) {
		var ret SpendControlDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControl) GetDirectionOk() (*SpendControlDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *SpendControl) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given SpendControlDirection and assigns it to the Direction field.
func (o *SpendControl) SetDirection(v SpendControlDirection) {
	o.Direction = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpendControl) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControl) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpendControl) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SpendControl) SetId(v string) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value
func (o *SpendControl) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *SpendControl) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *SpendControl) SetIsActive(v bool) {
	o.IsActive = v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *SpendControl) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControl) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *SpendControl) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *SpendControl) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetMerchantCategoryCodes returns the MerchantCategoryCodes field value if set, zero value otherwise.
func (o *SpendControl) GetMerchantCategoryCodes() []string {
	if o == nil || IsNil(o.MerchantCategoryCodes) {
		var ret []string
		return ret
	}
	return o.MerchantCategoryCodes
}

// GetMerchantCategoryCodesOk returns a tuple with the MerchantCategoryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControl) GetMerchantCategoryCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.MerchantCategoryCodes) {
		return nil, false
	}
	return o.MerchantCategoryCodes, true
}

// HasMerchantCategoryCodes returns a boolean if a field has been set.
func (o *SpendControl) HasMerchantCategoryCodes() bool {
	if o != nil && !IsNil(o.MerchantCategoryCodes) {
		return true
	}

	return false
}

// SetMerchantCategoryCodes gets a reference to the given []string and assigns it to the MerchantCategoryCodes field.
func (o *SpendControl) SetMerchantCategoryCodes(v []string) {
	o.MerchantCategoryCodes = v
}

// GetName returns the Name field value
func (o *SpendControl) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpendControl) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpendControl) SetName(v string) {
	o.Name = v
}

// GetPaymentSubTypes returns the PaymentSubTypes field value if set, zero value otherwise.
func (o *SpendControl) GetPaymentSubTypes() []PaymentSubType {
	if o == nil || IsNil(o.PaymentSubTypes) {
		var ret []PaymentSubType
		return ret
	}
	return o.PaymentSubTypes
}

// GetPaymentSubTypesOk returns a tuple with the PaymentSubTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControl) GetPaymentSubTypesOk() ([]PaymentSubType, bool) {
	if o == nil || IsNil(o.PaymentSubTypes) {
		return nil, false
	}
	return o.PaymentSubTypes, true
}

// HasPaymentSubTypes returns a boolean if a field has been set.
func (o *SpendControl) HasPaymentSubTypes() bool {
	if o != nil && !IsNil(o.PaymentSubTypes) {
		return true
	}

	return false
}

// SetPaymentSubTypes gets a reference to the given []PaymentSubType and assigns it to the PaymentSubTypes field.
func (o *SpendControl) SetPaymentSubTypes(v []PaymentSubType) {
	o.PaymentSubTypes = v
}

// GetPaymentTypes returns the PaymentTypes field value if set, zero value otherwise.
func (o *SpendControl) GetPaymentTypes() []PaymentType {
	if o == nil || IsNil(o.PaymentTypes) {
		var ret []PaymentType
		return ret
	}
	return o.PaymentTypes
}

// GetPaymentTypesOk returns a tuple with the PaymentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControl) GetPaymentTypesOk() ([]PaymentType, bool) {
	if o == nil || IsNil(o.PaymentTypes) {
		return nil, false
	}
	return o.PaymentTypes, true
}

// HasPaymentTypes returns a boolean if a field has been set.
func (o *SpendControl) HasPaymentTypes() bool {
	if o != nil && !IsNil(o.PaymentTypes) {
		return true
	}

	return false
}

// SetPaymentTypes gets a reference to the given []PaymentType and assigns it to the PaymentTypes field.
func (o *SpendControl) SetPaymentTypes(v []PaymentType) {
	o.PaymentTypes = v
}

// GetTimeRange returns the TimeRange field value
func (o *SpendControl) GetTimeRange() SpendControlTimeRange {
	if o == nil {
		var ret SpendControlTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *SpendControl) GetTimeRangeOk() (*SpendControlTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *SpendControl) SetTimeRange(v SpendControlTimeRange) {
	o.TimeRange = v
}

func (o SpendControl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpendControl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action_case"] = o.ActionCase
	toSerialize["action_decline"] = o.ActionDecline
	toSerialize["amount_limit"] = o.AmountLimit
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["is_active"] = o.IsActive
	if !IsNil(o.LastModifiedTime) {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if !IsNil(o.MerchantCategoryCodes) {
		toSerialize["merchant_category_codes"] = o.MerchantCategoryCodes
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PaymentSubTypes) {
		toSerialize["payment_sub_types"] = o.PaymentSubTypes
	}
	if !IsNil(o.PaymentTypes) {
		toSerialize["payment_types"] = o.PaymentTypes
	}
	toSerialize["time_range"] = o.TimeRange
	return toSerialize, nil
}

func (o *SpendControl) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action_case",
		"action_decline",
		"amount_limit",
		"is_active",
		"name",
		"time_range",
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

	varSpendControl := _SpendControl{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpendControl)

	if err != nil {
		return err
	}

	*o = SpendControl(varSpendControl)

	return err
}

type NullableSpendControl struct {
	value *SpendControl
	isSet bool
}

func (v NullableSpendControl) Get() *SpendControl {
	return v.value
}

func (v *NullableSpendControl) Set(val *SpendControl) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendControl) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendControl(val *SpendControl) *NullableSpendControl {
	return &NullableSpendControl{value: val, isSet: true}
}

func (v NullableSpendControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
