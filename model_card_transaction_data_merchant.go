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

// CardTransactionDataMerchant Information about the merchant
type CardTransactionDataMerchant struct {
	// Merchant address
	Address *string `json:"address,omitempty"`
	// Merchant city
	City *string `json:"city,omitempty"`
	// Merchant country code
	CountryCode *string `json:"country_code,omitempty"`
	// Independent sales organization identifier
	IndependentSalesOrganizationId *string `json:"independent_sales_organization_id,omitempty"`
	// Merchant category code
	Mcc *string `json:"mcc,omitempty"`
	// Merchant identifier
	Mid *string `json:"mid,omitempty"`
	// Merchant name
	Name *string `json:"name,omitempty"`
	// Payment facilitator identifier
	PaymentFacilitatorId *string `json:"payment_facilitator_id,omitempty"`
	// Merchant postal code
	PostalCode *string `json:"postal_code,omitempty"`
	// Merchant state
	State *string `json:"state,omitempty"`
	// Sub merchant identifier
	SubMerchantId *string `json:"sub_merchant_id,omitempty"`
}

// NewCardTransactionDataMerchant instantiates a new CardTransactionDataMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardTransactionDataMerchant() *CardTransactionDataMerchant {
	this := CardTransactionDataMerchant{}
	return &this
}

// NewCardTransactionDataMerchantWithDefaults instantiates a new CardTransactionDataMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardTransactionDataMerchantWithDefaults() *CardTransactionDataMerchant {
	this := CardTransactionDataMerchant{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CardTransactionDataMerchant) SetAddress(v string) {
	o.Address = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CardTransactionDataMerchant) SetCity(v string) {
	o.City = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *CardTransactionDataMerchant) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetIndependentSalesOrganizationId returns the IndependentSalesOrganizationId field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetIndependentSalesOrganizationId() string {
	if o == nil || o.IndependentSalesOrganizationId == nil {
		var ret string
		return ret
	}
	return *o.IndependentSalesOrganizationId
}

// GetIndependentSalesOrganizationIdOk returns a tuple with the IndependentSalesOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetIndependentSalesOrganizationIdOk() (*string, bool) {
	if o == nil || o.IndependentSalesOrganizationId == nil {
		return nil, false
	}
	return o.IndependentSalesOrganizationId, true
}

// HasIndependentSalesOrganizationId returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasIndependentSalesOrganizationId() bool {
	if o != nil && o.IndependentSalesOrganizationId != nil {
		return true
	}

	return false
}

// SetIndependentSalesOrganizationId gets a reference to the given string and assigns it to the IndependentSalesOrganizationId field.
func (o *CardTransactionDataMerchant) SetIndependentSalesOrganizationId(v string) {
	o.IndependentSalesOrganizationId = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetMcc() string {
	if o == nil || o.Mcc == nil {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetMccOk() (*string, bool) {
	if o == nil || o.Mcc == nil {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasMcc() bool {
	if o != nil && o.Mcc != nil {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *CardTransactionDataMerchant) SetMcc(v string) {
	o.Mcc = &v
}

// GetMid returns the Mid field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetMid() string {
	if o == nil || o.Mid == nil {
		var ret string
		return ret
	}
	return *o.Mid
}

// GetMidOk returns a tuple with the Mid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetMidOk() (*string, bool) {
	if o == nil || o.Mid == nil {
		return nil, false
	}
	return o.Mid, true
}

// HasMid returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasMid() bool {
	if o != nil && o.Mid != nil {
		return true
	}

	return false
}

// SetMid gets a reference to the given string and assigns it to the Mid field.
func (o *CardTransactionDataMerchant) SetMid(v string) {
	o.Mid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CardTransactionDataMerchant) SetName(v string) {
	o.Name = &v
}

// GetPaymentFacilitatorId returns the PaymentFacilitatorId field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetPaymentFacilitatorId() string {
	if o == nil || o.PaymentFacilitatorId == nil {
		var ret string
		return ret
	}
	return *o.PaymentFacilitatorId
}

// GetPaymentFacilitatorIdOk returns a tuple with the PaymentFacilitatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetPaymentFacilitatorIdOk() (*string, bool) {
	if o == nil || o.PaymentFacilitatorId == nil {
		return nil, false
	}
	return o.PaymentFacilitatorId, true
}

// HasPaymentFacilitatorId returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasPaymentFacilitatorId() bool {
	if o != nil && o.PaymentFacilitatorId != nil {
		return true
	}

	return false
}

// SetPaymentFacilitatorId gets a reference to the given string and assigns it to the PaymentFacilitatorId field.
func (o *CardTransactionDataMerchant) SetPaymentFacilitatorId(v string) {
	o.PaymentFacilitatorId = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *CardTransactionDataMerchant) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CardTransactionDataMerchant) SetState(v string) {
	o.State = &v
}

// GetSubMerchantId returns the SubMerchantId field value if set, zero value otherwise.
func (o *CardTransactionDataMerchant) GetSubMerchantId() string {
	if o == nil || o.SubMerchantId == nil {
		var ret string
		return ret
	}
	return *o.SubMerchantId
}

// GetSubMerchantIdOk returns a tuple with the SubMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataMerchant) GetSubMerchantIdOk() (*string, bool) {
	if o == nil || o.SubMerchantId == nil {
		return nil, false
	}
	return o.SubMerchantId, true
}

// HasSubMerchantId returns a boolean if a field has been set.
func (o *CardTransactionDataMerchant) HasSubMerchantId() bool {
	if o != nil && o.SubMerchantId != nil {
		return true
	}

	return false
}

// SetSubMerchantId gets a reference to the given string and assigns it to the SubMerchantId field.
func (o *CardTransactionDataMerchant) SetSubMerchantId(v string) {
	o.SubMerchantId = &v
}

func (o CardTransactionDataMerchant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.IndependentSalesOrganizationId != nil {
		toSerialize["independent_sales_organization_id"] = o.IndependentSalesOrganizationId
	}
	if o.Mcc != nil {
		toSerialize["mcc"] = o.Mcc
	}
	if o.Mid != nil {
		toSerialize["mid"] = o.Mid
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PaymentFacilitatorId != nil {
		toSerialize["payment_facilitator_id"] = o.PaymentFacilitatorId
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.SubMerchantId != nil {
		toSerialize["sub_merchant_id"] = o.SubMerchantId
	}
	return json.Marshal(toSerialize)
}

type NullableCardTransactionDataMerchant struct {
	value *CardTransactionDataMerchant
	isSet bool
}

func (v NullableCardTransactionDataMerchant) Get() *CardTransactionDataMerchant {
	return v.value
}

func (v *NullableCardTransactionDataMerchant) Set(val *CardTransactionDataMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableCardTransactionDataMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableCardTransactionDataMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardTransactionDataMerchant(val *CardTransactionDataMerchant) *NullableCardTransactionDataMerchant {
	return &NullableCardTransactionDataMerchant{value: val, isSet: true}
}

func (v NullableCardTransactionDataMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardTransactionDataMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
