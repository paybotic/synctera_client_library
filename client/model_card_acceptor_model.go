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

// checks if the CardAcceptorModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardAcceptorModel{}

// CardAcceptorModel struct for CardAcceptorModel
type CardAcceptorModel struct {
	Address                         *string `json:"address,omitempty"`
	City                            *string `json:"city,omitempty"`
	Country                         *string `json:"country,omitempty"`
	EcommerceSecurityLevelIndicator *string `json:"ecommerce_security_level_indicator,omitempty"`
	Mcc                             *string `json:"mcc,omitempty"`
	Name                            *string `json:"name,omitempty"`
	PartialApprovalCapable          *bool   `json:"partial_approval_capable,omitempty"`
	// Two-Letter USPS State Abbreviation
	State                *string `json:"state,omitempty" validate:"regexp=\\\\b(A[KLRZ]|C[AOT]|D[CE]|FL|GA|HI|I[ADLN]|K[SY]|LA|M[ADEINOST]|N[CDEHJMVY]|O[HKR]|PA|RI|S[CD]|T[NX]|UT|V[AT]|W[AIVY])\\\\b"`
	Zip                  *string `json:"zip,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CardAcceptorModel CardAcceptorModel

// NewCardAcceptorModel instantiates a new CardAcceptorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardAcceptorModel() *CardAcceptorModel {
	this := CardAcceptorModel{}
	var partialApprovalCapable bool = false
	this.PartialApprovalCapable = &partialApprovalCapable
	return &this
}

// NewCardAcceptorModelWithDefaults instantiates a new CardAcceptorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardAcceptorModelWithDefaults() *CardAcceptorModel {
	this := CardAcceptorModel{}
	var partialApprovalCapable bool = false
	this.PartialApprovalCapable = &partialApprovalCapable
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CardAcceptorModel) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAcceptorModel) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CardAcceptorModel) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *CardAcceptorModel) SetAddress(v string) {
	o.Address = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CardAcceptorModel) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAcceptorModel) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CardAcceptorModel) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CardAcceptorModel) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CardAcceptorModel) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAcceptorModel) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CardAcceptorModel) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CardAcceptorModel) SetCountry(v string) {
	o.Country = &v
}

// GetEcommerceSecurityLevelIndicator returns the EcommerceSecurityLevelIndicator field value if set, zero value otherwise.
func (o *CardAcceptorModel) GetEcommerceSecurityLevelIndicator() string {
	if o == nil || IsNil(o.EcommerceSecurityLevelIndicator) {
		var ret string
		return ret
	}
	return *o.EcommerceSecurityLevelIndicator
}

// GetEcommerceSecurityLevelIndicatorOk returns a tuple with the EcommerceSecurityLevelIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAcceptorModel) GetEcommerceSecurityLevelIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.EcommerceSecurityLevelIndicator) {
		return nil, false
	}
	return o.EcommerceSecurityLevelIndicator, true
}

// HasEcommerceSecurityLevelIndicator returns a boolean if a field has been set.
func (o *CardAcceptorModel) HasEcommerceSecurityLevelIndicator() bool {
	if o != nil && !IsNil(o.EcommerceSecurityLevelIndicator) {
		return true
	}

	return false
}

// SetEcommerceSecurityLevelIndicator gets a reference to the given string and assigns it to the EcommerceSecurityLevelIndicator field.
func (o *CardAcceptorModel) SetEcommerceSecurityLevelIndicator(v string) {
	o.EcommerceSecurityLevelIndicator = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *CardAcceptorModel) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAcceptorModel) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *CardAcceptorModel) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *CardAcceptorModel) SetMcc(v string) {
	o.Mcc = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CardAcceptorModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAcceptorModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CardAcceptorModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CardAcceptorModel) SetName(v string) {
	o.Name = &v
}

// GetPartialApprovalCapable returns the PartialApprovalCapable field value if set, zero value otherwise.
func (o *CardAcceptorModel) GetPartialApprovalCapable() bool {
	if o == nil || IsNil(o.PartialApprovalCapable) {
		var ret bool
		return ret
	}
	return *o.PartialApprovalCapable
}

// GetPartialApprovalCapableOk returns a tuple with the PartialApprovalCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAcceptorModel) GetPartialApprovalCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.PartialApprovalCapable) {
		return nil, false
	}
	return o.PartialApprovalCapable, true
}

// HasPartialApprovalCapable returns a boolean if a field has been set.
func (o *CardAcceptorModel) HasPartialApprovalCapable() bool {
	if o != nil && !IsNil(o.PartialApprovalCapable) {
		return true
	}

	return false
}

// SetPartialApprovalCapable gets a reference to the given bool and assigns it to the PartialApprovalCapable field.
func (o *CardAcceptorModel) SetPartialApprovalCapable(v bool) {
	o.PartialApprovalCapable = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CardAcceptorModel) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAcceptorModel) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CardAcceptorModel) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CardAcceptorModel) SetState(v string) {
	o.State = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *CardAcceptorModel) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardAcceptorModel) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *CardAcceptorModel) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *CardAcceptorModel) SetZip(v string) {
	o.Zip = &v
}

func (o CardAcceptorModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardAcceptorModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.EcommerceSecurityLevelIndicator) {
		toSerialize["ecommerce_security_level_indicator"] = o.EcommerceSecurityLevelIndicator
	}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PartialApprovalCapable) {
		toSerialize["partial_approval_capable"] = o.PartialApprovalCapable
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardAcceptorModel) UnmarshalJSON(data []byte) (err error) {
	varCardAcceptorModel := _CardAcceptorModel{}

	err = json.Unmarshal(data, &varCardAcceptorModel)

	if err != nil {
		return err
	}

	*o = CardAcceptorModel(varCardAcceptorModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "city")
		delete(additionalProperties, "country")
		delete(additionalProperties, "ecommerce_security_level_indicator")
		delete(additionalProperties, "mcc")
		delete(additionalProperties, "name")
		delete(additionalProperties, "partial_approval_capable")
		delete(additionalProperties, "state")
		delete(additionalProperties, "zip")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardAcceptorModel struct {
	value *CardAcceptorModel
	isSet bool
}

func (v NullableCardAcceptorModel) Get() *CardAcceptorModel {
	return v.value
}

func (v *NullableCardAcceptorModel) Set(val *CardAcceptorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCardAcceptorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCardAcceptorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardAcceptorModel(val *CardAcceptorModel) *NullableCardAcceptorModel {
	return &NullableCardAcceptorModel{value: val, isSet: true}
}

func (v NullableCardAcceptorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardAcceptorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
