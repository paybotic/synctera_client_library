/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// BinUpdateRequest struct for BinUpdateRequest
type BinUpdateRequest struct {
	BinStatus *BinStatus `json:"bin_status,omitempty"`
	CardBrand *CardBrand `json:"card_brand,omitempty"`
	CardCategory *CardCategory `json:"card_category,omitempty"`
	CardProductType *CardProductType `json:"card_product_type,omitempty"`
	// The time when bin is decommissioned
	EndDate *time.Time `json:"end_date,omitempty"`
	// Controls whether bin allows tokenization
	IsTokenizationEnabled *bool `json:"is_tokenization_enabled,omitempty"`
	PhysicalCardFormat *PhysicalCardFormat `json:"physical_card_format,omitempty"`
	// The time when bin goes live
	StartDate *time.Time `json:"start_date,omitempty"`
}

// NewBinUpdateRequest instantiates a new BinUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinUpdateRequest() *BinUpdateRequest {
	this := BinUpdateRequest{}
	return &this
}

// NewBinUpdateRequestWithDefaults instantiates a new BinUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinUpdateRequestWithDefaults() *BinUpdateRequest {
	this := BinUpdateRequest{}
	return &this
}

// GetBinStatus returns the BinStatus field value if set, zero value otherwise.
func (o *BinUpdateRequest) GetBinStatus() BinStatus {
	if o == nil || o.BinStatus == nil {
		var ret BinStatus
		return ret
	}
	return *o.BinStatus
}

// GetBinStatusOk returns a tuple with the BinStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinUpdateRequest) GetBinStatusOk() (*BinStatus, bool) {
	if o == nil || o.BinStatus == nil {
		return nil, false
	}
	return o.BinStatus, true
}

// HasBinStatus returns a boolean if a field has been set.
func (o *BinUpdateRequest) HasBinStatus() bool {
	if o != nil && o.BinStatus != nil {
		return true
	}

	return false
}

// SetBinStatus gets a reference to the given BinStatus and assigns it to the BinStatus field.
func (o *BinUpdateRequest) SetBinStatus(v BinStatus) {
	o.BinStatus = &v
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise.
func (o *BinUpdateRequest) GetCardBrand() CardBrand {
	if o == nil || o.CardBrand == nil {
		var ret CardBrand
		return ret
	}
	return *o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinUpdateRequest) GetCardBrandOk() (*CardBrand, bool) {
	if o == nil || o.CardBrand == nil {
		return nil, false
	}
	return o.CardBrand, true
}

// HasCardBrand returns a boolean if a field has been set.
func (o *BinUpdateRequest) HasCardBrand() bool {
	if o != nil && o.CardBrand != nil {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given CardBrand and assigns it to the CardBrand field.
func (o *BinUpdateRequest) SetCardBrand(v CardBrand) {
	o.CardBrand = &v
}

// GetCardCategory returns the CardCategory field value if set, zero value otherwise.
func (o *BinUpdateRequest) GetCardCategory() CardCategory {
	if o == nil || o.CardCategory == nil {
		var ret CardCategory
		return ret
	}
	return *o.CardCategory
}

// GetCardCategoryOk returns a tuple with the CardCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinUpdateRequest) GetCardCategoryOk() (*CardCategory, bool) {
	if o == nil || o.CardCategory == nil {
		return nil, false
	}
	return o.CardCategory, true
}

// HasCardCategory returns a boolean if a field has been set.
func (o *BinUpdateRequest) HasCardCategory() bool {
	if o != nil && o.CardCategory != nil {
		return true
	}

	return false
}

// SetCardCategory gets a reference to the given CardCategory and assigns it to the CardCategory field.
func (o *BinUpdateRequest) SetCardCategory(v CardCategory) {
	o.CardCategory = &v
}

// GetCardProductType returns the CardProductType field value if set, zero value otherwise.
func (o *BinUpdateRequest) GetCardProductType() CardProductType {
	if o == nil || o.CardProductType == nil {
		var ret CardProductType
		return ret
	}
	return *o.CardProductType
}

// GetCardProductTypeOk returns a tuple with the CardProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinUpdateRequest) GetCardProductTypeOk() (*CardProductType, bool) {
	if o == nil || o.CardProductType == nil {
		return nil, false
	}
	return o.CardProductType, true
}

// HasCardProductType returns a boolean if a field has been set.
func (o *BinUpdateRequest) HasCardProductType() bool {
	if o != nil && o.CardProductType != nil {
		return true
	}

	return false
}

// SetCardProductType gets a reference to the given CardProductType and assigns it to the CardProductType field.
func (o *BinUpdateRequest) SetCardProductType(v CardProductType) {
	o.CardProductType = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BinUpdateRequest) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinUpdateRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BinUpdateRequest) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BinUpdateRequest) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetIsTokenizationEnabled returns the IsTokenizationEnabled field value if set, zero value otherwise.
func (o *BinUpdateRequest) GetIsTokenizationEnabled() bool {
	if o == nil || o.IsTokenizationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsTokenizationEnabled
}

// GetIsTokenizationEnabledOk returns a tuple with the IsTokenizationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinUpdateRequest) GetIsTokenizationEnabledOk() (*bool, bool) {
	if o == nil || o.IsTokenizationEnabled == nil {
		return nil, false
	}
	return o.IsTokenizationEnabled, true
}

// HasIsTokenizationEnabled returns a boolean if a field has been set.
func (o *BinUpdateRequest) HasIsTokenizationEnabled() bool {
	if o != nil && o.IsTokenizationEnabled != nil {
		return true
	}

	return false
}

// SetIsTokenizationEnabled gets a reference to the given bool and assigns it to the IsTokenizationEnabled field.
func (o *BinUpdateRequest) SetIsTokenizationEnabled(v bool) {
	o.IsTokenizationEnabled = &v
}

// GetPhysicalCardFormat returns the PhysicalCardFormat field value if set, zero value otherwise.
func (o *BinUpdateRequest) GetPhysicalCardFormat() PhysicalCardFormat {
	if o == nil || o.PhysicalCardFormat == nil {
		var ret PhysicalCardFormat
		return ret
	}
	return *o.PhysicalCardFormat
}

// GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinUpdateRequest) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool) {
	if o == nil || o.PhysicalCardFormat == nil {
		return nil, false
	}
	return o.PhysicalCardFormat, true
}

// HasPhysicalCardFormat returns a boolean if a field has been set.
func (o *BinUpdateRequest) HasPhysicalCardFormat() bool {
	if o != nil && o.PhysicalCardFormat != nil {
		return true
	}

	return false
}

// SetPhysicalCardFormat gets a reference to the given PhysicalCardFormat and assigns it to the PhysicalCardFormat field.
func (o *BinUpdateRequest) SetPhysicalCardFormat(v PhysicalCardFormat) {
	o.PhysicalCardFormat = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BinUpdateRequest) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinUpdateRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BinUpdateRequest) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BinUpdateRequest) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o BinUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BinStatus != nil {
		toSerialize["bin_status"] = o.BinStatus
	}
	if o.CardBrand != nil {
		toSerialize["card_brand"] = o.CardBrand
	}
	if o.CardCategory != nil {
		toSerialize["card_category"] = o.CardCategory
	}
	if o.CardProductType != nil {
		toSerialize["card_product_type"] = o.CardProductType
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.IsTokenizationEnabled != nil {
		toSerialize["is_tokenization_enabled"] = o.IsTokenizationEnabled
	}
	if o.PhysicalCardFormat != nil {
		toSerialize["physical_card_format"] = o.PhysicalCardFormat
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableBinUpdateRequest struct {
	value *BinUpdateRequest
	isSet bool
}

func (v NullableBinUpdateRequest) Get() *BinUpdateRequest {
	return v.value
}

func (v *NullableBinUpdateRequest) Set(val *BinUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBinUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBinUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinUpdateRequest(val *BinUpdateRequest) *NullableBinUpdateRequest {
	return &NullableBinUpdateRequest{value: val, isSet: true}
}

func (v NullableBinUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


