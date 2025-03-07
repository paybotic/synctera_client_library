/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// checks if the AccountTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountTemplate{}

// AccountTemplate struct for AccountTemplate
type AccountTemplate struct {
	ApplicationType *ApplicationType `json:"application_type,omitempty"`
	// User provided account template description
	Description *string `json:"description,omitempty"`
	// Generated ID for the template
	Id *string `json:"id,omitempty"`
	// Whether this template can be used for account creation
	IsEnabled bool `json:"is_enabled"`
	// Unique account template name
	Name                 string         `json:"name"`
	Template             TemplateFields `json:"template"`
	AdditionalProperties map[string]interface{}
}

type _AccountTemplate AccountTemplate

// NewAccountTemplate instantiates a new AccountTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountTemplate(isEnabled bool, name string, template TemplateFields) *AccountTemplate {
	this := AccountTemplate{}
	this.IsEnabled = isEnabled
	this.Name = name
	this.Template = template
	return &this
}

// NewAccountTemplateWithDefaults instantiates a new AccountTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTemplateWithDefaults() *AccountTemplate {
	this := AccountTemplate{}
	return &this
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *AccountTemplate) GetApplicationType() ApplicationType {
	if o == nil || IsNil(o.ApplicationType) {
		var ret ApplicationType
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTemplate) GetApplicationTypeOk() (*ApplicationType, bool) {
	if o == nil || IsNil(o.ApplicationType) {
		return nil, false
	}
	return o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *AccountTemplate) HasApplicationType() bool {
	if o != nil && !IsNil(o.ApplicationType) {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given ApplicationType and assigns it to the ApplicationType field.
func (o *AccountTemplate) SetApplicationType(v ApplicationType) {
	o.ApplicationType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountTemplate) SetId(v string) {
	o.Id = &v
}

// GetIsEnabled returns the IsEnabled field value
func (o *AccountTemplate) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *AccountTemplate) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *AccountTemplate) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetName returns the Name field value
func (o *AccountTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountTemplate) SetName(v string) {
	o.Name = v
}

// GetTemplate returns the Template field value
func (o *AccountTemplate) GetTemplate() TemplateFields {
	if o == nil {
		var ret TemplateFields
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *AccountTemplate) GetTemplateOk() (*TemplateFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *AccountTemplate) SetTemplate(v TemplateFields) {
	o.Template = v
}

func (o AccountTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationType) {
		toSerialize["application_type"] = o.ApplicationType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["is_enabled"] = o.IsEnabled
	toSerialize["name"] = o.Name
	toSerialize["template"] = o.Template

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"is_enabled",
		"name",
		"template",
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

	varAccountTemplate := _AccountTemplate{}

	err = json.Unmarshal(data, &varAccountTemplate)

	if err != nil {
		return err
	}

	*o = AccountTemplate(varAccountTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "application_type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_enabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "template")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountTemplate struct {
	value *AccountTemplate
	isSet bool
}

func (v NullableAccountTemplate) Get() *AccountTemplate {
	return v.value
}

func (v *NullableAccountTemplate) Set(val *AccountTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTemplate(val *AccountTemplate) *NullableAccountTemplate {
	return &NullableAccountTemplate{value: val, isSet: true}
}

func (v NullableAccountTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
