/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// SpendControlResponse struct for SpendControlResponse
type SpendControlResponse struct {
	// If set, create a case for transactions that do not conform to the spend control
	ActionCase bool `json:"action_case"`
	// If set, decline transactions that do not conform to the spend control
	ActionDecline bool `json:"action_decline"`
	// Monetary limit for the spend control in the smallest currency unit (eg cents)
	AmountLimit int64 `json:"amount_limit"`
	// The timestamp representing when the spend control was created
	CreationTime time.Time `json:"creation_time"`
	Direction *SpendControlDirection `json:"direction,omitempty"`
	// Spend Control ID
	Id string `json:"id"`
	// Indicates if spend control is active
	IsActive bool `json:"is_active"`
	// The timestamp representing when the spend control was last modified
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Name assigned to spend control
	Name string `json:"name"`
	// A list of payment types to which a spend control will apply, if set. If not set or the array is empty, then the spend control will apply to all types of payments.
	PaymentTypes []PaymentType `json:"payment_types,omitempty"`
	TimeRange SpendControlTimeRange `json:"time_range"`
	// A count of how many accounts are using this spend control
	NumberOfRelatedAccounts int32 `json:"number_of_related_accounts"`
}

// NewSpendControlResponse instantiates a new SpendControlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendControlResponse(actionCase bool, actionDecline bool, amountLimit int64, creationTime time.Time, id string, isActive bool, lastModifiedTime time.Time, name string, timeRange SpendControlTimeRange, numberOfRelatedAccounts int32) *SpendControlResponse {
	this := SpendControlResponse{}
	this.ActionCase = actionCase
	this.ActionDecline = actionDecline
	this.AmountLimit = amountLimit
	this.CreationTime = creationTime
	this.Id = id
	this.IsActive = isActive
	this.LastModifiedTime = lastModifiedTime
	this.Name = name
	this.TimeRange = timeRange
	this.NumberOfRelatedAccounts = numberOfRelatedAccounts
	return &this
}

// NewSpendControlResponseWithDefaults instantiates a new SpendControlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendControlResponseWithDefaults() *SpendControlResponse {
	this := SpendControlResponse{}
	return &this
}

// GetActionCase returns the ActionCase field value
func (o *SpendControlResponse) GetActionCase() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ActionCase
}

// GetActionCaseOk returns a tuple with the ActionCase field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetActionCaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionCase, true
}

// SetActionCase sets field value
func (o *SpendControlResponse) SetActionCase(v bool) {
	o.ActionCase = v
}

// GetActionDecline returns the ActionDecline field value
func (o *SpendControlResponse) GetActionDecline() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ActionDecline
}

// GetActionDeclineOk returns a tuple with the ActionDecline field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetActionDeclineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionDecline, true
}

// SetActionDecline sets field value
func (o *SpendControlResponse) SetActionDecline(v bool) {
	o.ActionDecline = v
}

// GetAmountLimit returns the AmountLimit field value
func (o *SpendControlResponse) GetAmountLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AmountLimit
}

// GetAmountLimitOk returns a tuple with the AmountLimit field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetAmountLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountLimit, true
}

// SetAmountLimit sets field value
func (o *SpendControlResponse) SetAmountLimit(v int64) {
	o.AmountLimit = v
}

// GetCreationTime returns the CreationTime field value
func (o *SpendControlResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *SpendControlResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *SpendControlResponse) GetDirection() SpendControlDirection {
	if o == nil || o.Direction == nil {
		var ret SpendControlDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetDirectionOk() (*SpendControlDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *SpendControlResponse) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given SpendControlDirection and assigns it to the Direction field.
func (o *SpendControlResponse) SetDirection(v SpendControlDirection) {
	o.Direction = &v
}

// GetId returns the Id field value
func (o *SpendControlResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SpendControlResponse) SetId(v string) {
	o.Id = v
}

// GetIsActive returns the IsActive field value
func (o *SpendControlResponse) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *SpendControlResponse) SetIsActive(v bool) {
	o.IsActive = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *SpendControlResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *SpendControlResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetName returns the Name field value
func (o *SpendControlResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpendControlResponse) SetName(v string) {
	o.Name = v
}

// GetPaymentTypes returns the PaymentTypes field value if set, zero value otherwise.
func (o *SpendControlResponse) GetPaymentTypes() []PaymentType {
	if o == nil || o.PaymentTypes == nil {
		var ret []PaymentType
		return ret
	}
	return o.PaymentTypes
}

// GetPaymentTypesOk returns a tuple with the PaymentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetPaymentTypesOk() ([]PaymentType, bool) {
	if o == nil || o.PaymentTypes == nil {
		return nil, false
	}
	return o.PaymentTypes, true
}

// HasPaymentTypes returns a boolean if a field has been set.
func (o *SpendControlResponse) HasPaymentTypes() bool {
	if o != nil && o.PaymentTypes != nil {
		return true
	}

	return false
}

// SetPaymentTypes gets a reference to the given []PaymentType and assigns it to the PaymentTypes field.
func (o *SpendControlResponse) SetPaymentTypes(v []PaymentType) {
	o.PaymentTypes = v
}

// GetTimeRange returns the TimeRange field value
func (o *SpendControlResponse) GetTimeRange() SpendControlTimeRange {
	if o == nil {
		var ret SpendControlTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetTimeRangeOk() (*SpendControlTimeRange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *SpendControlResponse) SetTimeRange(v SpendControlTimeRange) {
	o.TimeRange = v
}

// GetNumberOfRelatedAccounts returns the NumberOfRelatedAccounts field value
func (o *SpendControlResponse) GetNumberOfRelatedAccounts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfRelatedAccounts
}

// GetNumberOfRelatedAccountsOk returns a tuple with the NumberOfRelatedAccounts field value
// and a boolean to check if the value has been set.
func (o *SpendControlResponse) GetNumberOfRelatedAccountsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfRelatedAccounts, true
}

// SetNumberOfRelatedAccounts sets field value
func (o *SpendControlResponse) SetNumberOfRelatedAccounts(v int32) {
	o.NumberOfRelatedAccounts = v
}

func (o SpendControlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action_case"] = o.ActionCase
	}
	if true {
		toSerialize["action_decline"] = o.ActionDecline
	}
	if true {
		toSerialize["amount_limit"] = o.AmountLimit
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["is_active"] = o.IsActive
	}
	if true {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.PaymentTypes != nil {
		toSerialize["payment_types"] = o.PaymentTypes
	}
	if true {
		toSerialize["time_range"] = o.TimeRange
	}
	if true {
		toSerialize["number_of_related_accounts"] = o.NumberOfRelatedAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableSpendControlResponse struct {
	value *SpendControlResponse
	isSet bool
}

func (v NullableSpendControlResponse) Get() *SpendControlResponse {
	return v.value
}

func (v *NullableSpendControlResponse) Set(val *SpendControlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendControlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendControlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendControlResponse(val *SpendControlResponse) *NullableSpendControlResponse {
	return &NullableSpendControlResponse{value: val, isSet: true}
}

func (v NullableSpendControlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendControlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


