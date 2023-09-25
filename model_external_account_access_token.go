/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ExternalAccountAccessToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalAccountAccessToken{}

// ExternalAccountAccessToken struct for ExternalAccountAccessToken
type ExternalAccountAccessToken struct {
	// The identifier for the business customer associated with this external account. Exactly one of `business_id` or `customer_id` must be specified. 
	BusinessId *string `json:"business_id,omitempty"`
	// The identifier for the personal customer associated with this external account. Exactly one of `customer_id` or `business_id` must be specified. 
	CustomerId *string `json:"customer_id,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting
	RequestId *string `json:"request_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	VendorAccessToken *string `json:"vendor_access_token,omitempty"`
	// An alias for `customer_id` (deprecated).
	VendorCustomerId *string `json:"vendor_customer_id,omitempty"`
	// The ID of the institution the access token is requested for 
	VendorInstitutionId string `json:"vendor_institution_id"`
	// The user's public token obtained from successful link login. 
	VendorPublicToken string `json:"vendor_public_token"`
}

// NewExternalAccountAccessToken instantiates a new ExternalAccountAccessToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccountAccessToken(vendorInstitutionId string, vendorPublicToken string) *ExternalAccountAccessToken {
	this := ExternalAccountAccessToken{}
	this.VendorInstitutionId = vendorInstitutionId
	this.VendorPublicToken = vendorPublicToken
	return &this
}

// NewExternalAccountAccessTokenWithDefaults instantiates a new ExternalAccountAccessToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountAccessTokenWithDefaults() *ExternalAccountAccessToken {
	this := ExternalAccountAccessToken{}
	return &this
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *ExternalAccountAccessToken) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *ExternalAccountAccessToken) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *ExternalAccountAccessToken) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *ExternalAccountAccessToken) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *ExternalAccountAccessToken) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *ExternalAccountAccessToken) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ExternalAccountAccessToken) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ExternalAccountAccessToken) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ExternalAccountAccessToken) SetRequestId(v string) {
	o.RequestId = &v
}

// GetVendorAccessToken returns the VendorAccessToken field value if set, zero value otherwise.
func (o *ExternalAccountAccessToken) GetVendorAccessToken() string {
	if o == nil || IsNil(o.VendorAccessToken) {
		var ret string
		return ret
	}
	return *o.VendorAccessToken
}

// GetVendorAccessTokenOk returns a tuple with the VendorAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetVendorAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.VendorAccessToken) {
		return nil, false
	}
	return o.VendorAccessToken, true
}

// HasVendorAccessToken returns a boolean if a field has been set.
func (o *ExternalAccountAccessToken) HasVendorAccessToken() bool {
	if o != nil && !IsNil(o.VendorAccessToken) {
		return true
	}

	return false
}

// SetVendorAccessToken gets a reference to the given string and assigns it to the VendorAccessToken field.
func (o *ExternalAccountAccessToken) SetVendorAccessToken(v string) {
	o.VendorAccessToken = &v
}

// GetVendorCustomerId returns the VendorCustomerId field value if set, zero value otherwise.
func (o *ExternalAccountAccessToken) GetVendorCustomerId() string {
	if o == nil || IsNil(o.VendorCustomerId) {
		var ret string
		return ret
	}
	return *o.VendorCustomerId
}

// GetVendorCustomerIdOk returns a tuple with the VendorCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetVendorCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorCustomerId) {
		return nil, false
	}
	return o.VendorCustomerId, true
}

// HasVendorCustomerId returns a boolean if a field has been set.
func (o *ExternalAccountAccessToken) HasVendorCustomerId() bool {
	if o != nil && !IsNil(o.VendorCustomerId) {
		return true
	}

	return false
}

// SetVendorCustomerId gets a reference to the given string and assigns it to the VendorCustomerId field.
func (o *ExternalAccountAccessToken) SetVendorCustomerId(v string) {
	o.VendorCustomerId = &v
}

// GetVendorInstitutionId returns the VendorInstitutionId field value
func (o *ExternalAccountAccessToken) GetVendorInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorInstitutionId
}

// GetVendorInstitutionIdOk returns a tuple with the VendorInstitutionId field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetVendorInstitutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorInstitutionId, true
}

// SetVendorInstitutionId sets field value
func (o *ExternalAccountAccessToken) SetVendorInstitutionId(v string) {
	o.VendorInstitutionId = v
}

// GetVendorPublicToken returns the VendorPublicToken field value
func (o *ExternalAccountAccessToken) GetVendorPublicToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorPublicToken
}

// GetVendorPublicTokenOk returns a tuple with the VendorPublicToken field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetVendorPublicTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VendorPublicToken, true
}

// SetVendorPublicToken sets field value
func (o *ExternalAccountAccessToken) SetVendorPublicToken(v string) {
	o.VendorPublicToken = v
}

func (o ExternalAccountAccessToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalAccountAccessToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.VendorAccessToken) {
		toSerialize["vendor_access_token"] = o.VendorAccessToken
	}
	if !IsNil(o.VendorCustomerId) {
		toSerialize["vendor_customer_id"] = o.VendorCustomerId
	}
	toSerialize["vendor_institution_id"] = o.VendorInstitutionId
	toSerialize["vendor_public_token"] = o.VendorPublicToken
	return toSerialize, nil
}

type NullableExternalAccountAccessToken struct {
	value *ExternalAccountAccessToken
	isSet bool
}

func (v NullableExternalAccountAccessToken) Get() *ExternalAccountAccessToken {
	return v.value
}

func (v *NullableExternalAccountAccessToken) Set(val *ExternalAccountAccessToken) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccountAccessToken) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccountAccessToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccountAccessToken(val *ExternalAccountAccessToken) *NullableExternalAccountAccessToken {
	return &NullableExternalAccountAccessToken{value: val, isSet: true}
}

func (v NullableExternalAccountAccessToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccountAccessToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


