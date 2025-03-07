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

// checks if the PostLicense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLicense{}

// PostLicense struct for PostLicense
type PostLicense struct {
	// Unique ID for the business. Exactly one of `business_id` or `person_id` must be set.
	BusinessId *string `json:"business_id,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	// The entity's license number
	LicenseNumber *string      `json:"license_number,omitempty"`
	LicenseType   *LicenseType `json:"license_type,omitempty"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant *string `json:"tenant,omitempty"`
}

// NewPostLicense instantiates a new PostLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLicense() *PostLicense {
	this := PostLicense{}
	return &this
}

// NewPostLicenseWithDefaults instantiates a new PostLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLicenseWithDefaults() *PostLicense {
	this := PostLicense{}
	return &this
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *PostLicense) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLicense) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *PostLicense) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *PostLicense) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *PostLicense) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLicense) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PostLicense) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *PostLicense) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetLicenseNumber returns the LicenseNumber field value if set, zero value otherwise.
func (o *PostLicense) GetLicenseNumber() string {
	if o == nil || IsNil(o.LicenseNumber) {
		var ret string
		return ret
	}
	return *o.LicenseNumber
}

// GetLicenseNumberOk returns a tuple with the LicenseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLicense) GetLicenseNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseNumber) {
		return nil, false
	}
	return o.LicenseNumber, true
}

// HasLicenseNumber returns a boolean if a field has been set.
func (o *PostLicense) HasLicenseNumber() bool {
	if o != nil && !IsNil(o.LicenseNumber) {
		return true
	}

	return false
}

// SetLicenseNumber gets a reference to the given string and assigns it to the LicenseNumber field.
func (o *PostLicense) SetLicenseNumber(v string) {
	o.LicenseNumber = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *PostLicense) GetLicenseType() LicenseType {
	if o == nil || IsNil(o.LicenseType) {
		var ret LicenseType
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLicense) GetLicenseTypeOk() (*LicenseType, bool) {
	if o == nil || IsNil(o.LicenseType) {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *PostLicense) HasLicenseType() bool {
	if o != nil && !IsNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given LicenseType and assigns it to the LicenseType field.
func (o *PostLicense) SetLicenseType(v LicenseType) {
	o.LicenseType = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *PostLicense) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLicense) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *PostLicense) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *PostLicense) SetTenant(v string) {
	o.Tenant = &v
}

func (o PostLicense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLicense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.LicenseNumber) {
		toSerialize["license_number"] = o.LicenseNumber
	}
	if !IsNil(o.LicenseType) {
		toSerialize["license_type"] = o.LicenseType
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	return toSerialize, nil
}

type NullablePostLicense struct {
	value *PostLicense
	isSet bool
}

func (v NullablePostLicense) Get() *PostLicense {
	return v.value
}

func (v *NullablePostLicense) Set(val *PostLicense) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLicense) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLicense(val *PostLicense) *NullablePostLicense {
	return &NullablePostLicense{value: val, isSet: true}
}

func (v NullablePostLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
