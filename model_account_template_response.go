/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AccountTemplateResponse struct for AccountTemplateResponse
type AccountTemplateResponse struct {
	ApplicationType *ApplicationType `json:"application_type,omitempty"`
	// Account template description
	Description *string `json:"description,omitempty"`
	// Generated ID for the template
	Id *string `json:"id,omitempty"`
	// Whether this template can be used for account creation
	IsEnabled bool `json:"is_enabled"`
	// Unique account template name
	Name string `json:"name"`
	Template TemplateFieldsGenericResponse `json:"template"`
}

// NewAccountTemplateResponse instantiates a new AccountTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountTemplateResponse(isEnabled bool, name string, template TemplateFieldsGenericResponse) *AccountTemplateResponse {
	this := AccountTemplateResponse{}
	this.IsEnabled = isEnabled
	this.Name = name
	this.Template = template
	return &this
}

// NewAccountTemplateResponseWithDefaults instantiates a new AccountTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTemplateResponseWithDefaults() *AccountTemplateResponse {
	this := AccountTemplateResponse{}
	return &this
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *AccountTemplateResponse) GetApplicationType() ApplicationType {
	if o == nil || o.ApplicationType == nil {
		var ret ApplicationType
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTemplateResponse) GetApplicationTypeOk() (*ApplicationType, bool) {
	if o == nil || o.ApplicationType == nil {
		return nil, false
	}
	return o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *AccountTemplateResponse) HasApplicationType() bool {
	if o != nil && o.ApplicationType != nil {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given ApplicationType and assigns it to the ApplicationType field.
func (o *AccountTemplateResponse) SetApplicationType(v ApplicationType) {
	o.ApplicationType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountTemplateResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTemplateResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountTemplateResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountTemplateResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountTemplateResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTemplateResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountTemplateResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountTemplateResponse) SetId(v string) {
	o.Id = &v
}

// GetIsEnabled returns the IsEnabled field value
func (o *AccountTemplateResponse) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *AccountTemplateResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *AccountTemplateResponse) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetName returns the Name field value
func (o *AccountTemplateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountTemplateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountTemplateResponse) SetName(v string) {
	o.Name = v
}

// GetTemplate returns the Template field value
func (o *AccountTemplateResponse) GetTemplate() TemplateFieldsGenericResponse {
	if o == nil {
		var ret TemplateFieldsGenericResponse
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *AccountTemplateResponse) GetTemplateOk() (*TemplateFieldsGenericResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *AccountTemplateResponse) SetTemplate(v TemplateFieldsGenericResponse) {
	o.Template = v
}

func (o AccountTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationType != nil {
		toSerialize["application_type"] = o.ApplicationType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableAccountTemplateResponse struct {
	value *AccountTemplateResponse
	isSet bool
}

func (v NullableAccountTemplateResponse) Get() *AccountTemplateResponse {
	return v.value
}

func (v *NullableAccountTemplateResponse) Set(val *AccountTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTemplateResponse(val *AccountTemplateResponse) *NullableAccountTemplateResponse {
	return &NullableAccountTemplateResponse{value: val, isSet: true}
}

func (v NullableAccountTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


