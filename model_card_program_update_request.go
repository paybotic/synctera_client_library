/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CardProgramUpdateRequest struct for CardProgramUpdateRequest
type CardProgramUpdateRequest struct {
	CardBrand *CardBrand `json:"card_brand,omitempty"`
	CardCategory *CardCategory `json:"card_category,omitempty"`
	CardProductType *CardProductType `json:"card_product_type,omitempty"`
	// The time when program became inactive
	EndDate *time.Time `json:"end_date,omitempty"`
	// Program name
	Name *string `json:"name,omitempty"`
	// The time when program becomes active
	StartDate *time.Time `json:"start_date,omitempty"`
}

// NewCardProgramUpdateRequest instantiates a new CardProgramUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProgramUpdateRequest() *CardProgramUpdateRequest {
	this := CardProgramUpdateRequest{}
	return &this
}

// NewCardProgramUpdateRequestWithDefaults instantiates a new CardProgramUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProgramUpdateRequestWithDefaults() *CardProgramUpdateRequest {
	this := CardProgramUpdateRequest{}
	return &this
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise.
func (o *CardProgramUpdateRequest) GetCardBrand() CardBrand {
	if o == nil || o.CardBrand == nil {
		var ret CardBrand
		return ret
	}
	return *o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramUpdateRequest) GetCardBrandOk() (*CardBrand, bool) {
	if o == nil || o.CardBrand == nil {
		return nil, false
	}
	return o.CardBrand, true
}

// HasCardBrand returns a boolean if a field has been set.
func (o *CardProgramUpdateRequest) HasCardBrand() bool {
	if o != nil && o.CardBrand != nil {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given CardBrand and assigns it to the CardBrand field.
func (o *CardProgramUpdateRequest) SetCardBrand(v CardBrand) {
	o.CardBrand = &v
}

// GetCardCategory returns the CardCategory field value if set, zero value otherwise.
func (o *CardProgramUpdateRequest) GetCardCategory() CardCategory {
	if o == nil || o.CardCategory == nil {
		var ret CardCategory
		return ret
	}
	return *o.CardCategory
}

// GetCardCategoryOk returns a tuple with the CardCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramUpdateRequest) GetCardCategoryOk() (*CardCategory, bool) {
	if o == nil || o.CardCategory == nil {
		return nil, false
	}
	return o.CardCategory, true
}

// HasCardCategory returns a boolean if a field has been set.
func (o *CardProgramUpdateRequest) HasCardCategory() bool {
	if o != nil && o.CardCategory != nil {
		return true
	}

	return false
}

// SetCardCategory gets a reference to the given CardCategory and assigns it to the CardCategory field.
func (o *CardProgramUpdateRequest) SetCardCategory(v CardCategory) {
	o.CardCategory = &v
}

// GetCardProductType returns the CardProductType field value if set, zero value otherwise.
func (o *CardProgramUpdateRequest) GetCardProductType() CardProductType {
	if o == nil || o.CardProductType == nil {
		var ret CardProductType
		return ret
	}
	return *o.CardProductType
}

// GetCardProductTypeOk returns a tuple with the CardProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramUpdateRequest) GetCardProductTypeOk() (*CardProductType, bool) {
	if o == nil || o.CardProductType == nil {
		return nil, false
	}
	return o.CardProductType, true
}

// HasCardProductType returns a boolean if a field has been set.
func (o *CardProgramUpdateRequest) HasCardProductType() bool {
	if o != nil && o.CardProductType != nil {
		return true
	}

	return false
}

// SetCardProductType gets a reference to the given CardProductType and assigns it to the CardProductType field.
func (o *CardProgramUpdateRequest) SetCardProductType(v CardProductType) {
	o.CardProductType = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CardProgramUpdateRequest) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramUpdateRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CardProgramUpdateRequest) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *CardProgramUpdateRequest) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CardProgramUpdateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CardProgramUpdateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CardProgramUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CardProgramUpdateRequest) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramUpdateRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CardProgramUpdateRequest) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CardProgramUpdateRequest) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o CardProgramUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableCardProgramUpdateRequest struct {
	value *CardProgramUpdateRequest
	isSet bool
}

func (v NullableCardProgramUpdateRequest) Get() *CardProgramUpdateRequest {
	return v.value
}

func (v *NullableCardProgramUpdateRequest) Set(val *CardProgramUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProgramUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProgramUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProgramUpdateRequest(val *CardProgramUpdateRequest) *NullableCardProgramUpdateRequest {
	return &NullableCardProgramUpdateRequest{value: val, isSet: true}
}

func (v NullableCardProgramUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProgramUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


