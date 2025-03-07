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

// checks if the MerchantResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantResponse{}

// MerchantResponse struct for MerchantResponse
type MerchantResponse struct {
	CreatedAt            NullableString `json:"created_at,omitempty"`
	Guid                 NullableString `json:"guid,omitempty"`
	LogoUrl              NullableString `json:"logo_url,omitempty"`
	Name                 NullableString `json:"name,omitempty"`
	UpdatedAt            NullableString `json:"updated_at,omitempty"`
	WebsiteUrl           NullableString `json:"website_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MerchantResponse MerchantResponse

// NewMerchantResponse instantiates a new MerchantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantResponse() *MerchantResponse {
	this := MerchantResponse{}
	return &this
}

// NewMerchantResponseWithDefaults instantiates a new MerchantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantResponseWithDefaults() *MerchantResponse {
	this := MerchantResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MerchantResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *MerchantResponse) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *MerchantResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *MerchantResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetGuid returns the Guid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantResponse) GetGuid() string {
	if o == nil || IsNil(o.Guid.Get()) {
		var ret string
		return ret
	}
	return *o.Guid.Get()
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantResponse) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guid.Get(), o.Guid.IsSet()
}

// HasGuid returns a boolean if a field has been set.
func (o *MerchantResponse) HasGuid() bool {
	if o != nil && o.Guid.IsSet() {
		return true
	}

	return false
}

// SetGuid gets a reference to the given NullableString and assigns it to the Guid field.
func (o *MerchantResponse) SetGuid(v string) {
	o.Guid.Set(&v)
}

// SetGuidNil sets the value for Guid to be an explicit nil
func (o *MerchantResponse) SetGuidNil() {
	o.Guid.Set(nil)
}

// UnsetGuid ensures that no value is present for Guid, not even an explicit nil
func (o *MerchantResponse) UnsetGuid() {
	o.Guid.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantResponse) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantResponse) GetLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *MerchantResponse) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *MerchantResponse) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *MerchantResponse) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *MerchantResponse) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantResponse) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MerchantResponse) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MerchantResponse) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *MerchantResponse) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MerchantResponse) UnsetName() {
	o.Name.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MerchantResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *MerchantResponse) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *MerchantResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *MerchantResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantResponse) GetWebsiteUrl() string {
	if o == nil || IsNil(o.WebsiteUrl.Get()) {
		var ret string
		return ret
	}
	return *o.WebsiteUrl.Get()
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantResponse) GetWebsiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebsiteUrl.Get(), o.WebsiteUrl.IsSet()
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *MerchantResponse) HasWebsiteUrl() bool {
	if o != nil && o.WebsiteUrl.IsSet() {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given NullableString and assigns it to the WebsiteUrl field.
func (o *MerchantResponse) SetWebsiteUrl(v string) {
	o.WebsiteUrl.Set(&v)
}

// SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil
func (o *MerchantResponse) SetWebsiteUrlNil() {
	o.WebsiteUrl.Set(nil)
}

// UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
func (o *MerchantResponse) UnsetWebsiteUrl() {
	o.WebsiteUrl.Unset()
}

func (o MerchantResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Guid.IsSet() {
		toSerialize["guid"] = o.Guid.Get()
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.WebsiteUrl.IsSet() {
		toSerialize["website_url"] = o.WebsiteUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MerchantResponse) UnmarshalJSON(data []byte) (err error) {
	varMerchantResponse := _MerchantResponse{}

	err = json.Unmarshal(data, &varMerchantResponse)

	if err != nil {
		return err
	}

	*o = MerchantResponse(varMerchantResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "guid")
		delete(additionalProperties, "logo_url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "website_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMerchantResponse struct {
	value *MerchantResponse
	isSet bool
}

func (v NullableMerchantResponse) Get() *MerchantResponse {
	return v.value
}

func (v *NullableMerchantResponse) Set(val *MerchantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantResponse(val *MerchantResponse) *NullableMerchantResponse {
	return &NullableMerchantResponse{value: val, isSet: true}
}

func (v NullableMerchantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
