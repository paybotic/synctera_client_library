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

// checks if the RailsDepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RailsDepository{}

// RailsDepository Rails available to enable on depository accounts
type RailsDepository struct {
	// A flag to indicate whether ACH transactions are enabled.
	IsAchEnabled *bool `json:"is_ach_enabled,omitempty"`
	// A flag to indicate whether card transactions are enabled.
	IsCardEnabled *bool `json:"is_card_enabled,omitempty"`
	// A flag to indicate whether EFT Canada transactions are enabled.
	IsEftCaEnabled *bool `json:"is_eft_ca_enabled,omitempty"`
	// A flag to indicate whether external card transactions are enabled.
	IsExternalCardEnabled *bool `json:"is_external_card_enabled,omitempty"`
	// A flag to indicate whether P2P transactions are enabled.
	IsP2pEnabled *bool `json:"is_p2p_enabled,omitempty"`
	// A flag to indicate whether Synctera Pay transactions are enabled.
	IsSyncteraPayEnabled *bool `json:"is_synctera_pay_enabled,omitempty"`
	// A flag to indicate whether wire transactions are enabled.
	IsWireEnabled *bool `json:"is_wire_enabled,omitempty"`
}

// NewRailsDepository instantiates a new RailsDepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRailsDepository() *RailsDepository {
	this := RailsDepository{}
	return &this
}

// NewRailsDepositoryWithDefaults instantiates a new RailsDepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRailsDepositoryWithDefaults() *RailsDepository {
	this := RailsDepository{}
	return &this
}

// GetIsAchEnabled returns the IsAchEnabled field value if set, zero value otherwise.
func (o *RailsDepository) GetIsAchEnabled() bool {
	if o == nil || IsNil(o.IsAchEnabled) {
		var ret bool
		return ret
	}
	return *o.IsAchEnabled
}

// GetIsAchEnabledOk returns a tuple with the IsAchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RailsDepository) GetIsAchEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAchEnabled) {
		return nil, false
	}
	return o.IsAchEnabled, true
}

// HasIsAchEnabled returns a boolean if a field has been set.
func (o *RailsDepository) HasIsAchEnabled() bool {
	if o != nil && !IsNil(o.IsAchEnabled) {
		return true
	}

	return false
}

// SetIsAchEnabled gets a reference to the given bool and assigns it to the IsAchEnabled field.
func (o *RailsDepository) SetIsAchEnabled(v bool) {
	o.IsAchEnabled = &v
}

// GetIsCardEnabled returns the IsCardEnabled field value if set, zero value otherwise.
func (o *RailsDepository) GetIsCardEnabled() bool {
	if o == nil || IsNil(o.IsCardEnabled) {
		var ret bool
		return ret
	}
	return *o.IsCardEnabled
}

// GetIsCardEnabledOk returns a tuple with the IsCardEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RailsDepository) GetIsCardEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCardEnabled) {
		return nil, false
	}
	return o.IsCardEnabled, true
}

// HasIsCardEnabled returns a boolean if a field has been set.
func (o *RailsDepository) HasIsCardEnabled() bool {
	if o != nil && !IsNil(o.IsCardEnabled) {
		return true
	}

	return false
}

// SetIsCardEnabled gets a reference to the given bool and assigns it to the IsCardEnabled field.
func (o *RailsDepository) SetIsCardEnabled(v bool) {
	o.IsCardEnabled = &v
}

// GetIsEftCaEnabled returns the IsEftCaEnabled field value if set, zero value otherwise.
func (o *RailsDepository) GetIsEftCaEnabled() bool {
	if o == nil || IsNil(o.IsEftCaEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEftCaEnabled
}

// GetIsEftCaEnabledOk returns a tuple with the IsEftCaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RailsDepository) GetIsEftCaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEftCaEnabled) {
		return nil, false
	}
	return o.IsEftCaEnabled, true
}

// HasIsEftCaEnabled returns a boolean if a field has been set.
func (o *RailsDepository) HasIsEftCaEnabled() bool {
	if o != nil && !IsNil(o.IsEftCaEnabled) {
		return true
	}

	return false
}

// SetIsEftCaEnabled gets a reference to the given bool and assigns it to the IsEftCaEnabled field.
func (o *RailsDepository) SetIsEftCaEnabled(v bool) {
	o.IsEftCaEnabled = &v
}

// GetIsExternalCardEnabled returns the IsExternalCardEnabled field value if set, zero value otherwise.
func (o *RailsDepository) GetIsExternalCardEnabled() bool {
	if o == nil || IsNil(o.IsExternalCardEnabled) {
		var ret bool
		return ret
	}
	return *o.IsExternalCardEnabled
}

// GetIsExternalCardEnabledOk returns a tuple with the IsExternalCardEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RailsDepository) GetIsExternalCardEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExternalCardEnabled) {
		return nil, false
	}
	return o.IsExternalCardEnabled, true
}

// HasIsExternalCardEnabled returns a boolean if a field has been set.
func (o *RailsDepository) HasIsExternalCardEnabled() bool {
	if o != nil && !IsNil(o.IsExternalCardEnabled) {
		return true
	}

	return false
}

// SetIsExternalCardEnabled gets a reference to the given bool and assigns it to the IsExternalCardEnabled field.
func (o *RailsDepository) SetIsExternalCardEnabled(v bool) {
	o.IsExternalCardEnabled = &v
}

// GetIsP2pEnabled returns the IsP2pEnabled field value if set, zero value otherwise.
func (o *RailsDepository) GetIsP2pEnabled() bool {
	if o == nil || IsNil(o.IsP2pEnabled) {
		var ret bool
		return ret
	}
	return *o.IsP2pEnabled
}

// GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RailsDepository) GetIsP2pEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsP2pEnabled) {
		return nil, false
	}
	return o.IsP2pEnabled, true
}

// HasIsP2pEnabled returns a boolean if a field has been set.
func (o *RailsDepository) HasIsP2pEnabled() bool {
	if o != nil && !IsNil(o.IsP2pEnabled) {
		return true
	}

	return false
}

// SetIsP2pEnabled gets a reference to the given bool and assigns it to the IsP2pEnabled field.
func (o *RailsDepository) SetIsP2pEnabled(v bool) {
	o.IsP2pEnabled = &v
}

// GetIsSyncteraPayEnabled returns the IsSyncteraPayEnabled field value if set, zero value otherwise.
func (o *RailsDepository) GetIsSyncteraPayEnabled() bool {
	if o == nil || IsNil(o.IsSyncteraPayEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSyncteraPayEnabled
}

// GetIsSyncteraPayEnabledOk returns a tuple with the IsSyncteraPayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RailsDepository) GetIsSyncteraPayEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncteraPayEnabled) {
		return nil, false
	}
	return o.IsSyncteraPayEnabled, true
}

// HasIsSyncteraPayEnabled returns a boolean if a field has been set.
func (o *RailsDepository) HasIsSyncteraPayEnabled() bool {
	if o != nil && !IsNil(o.IsSyncteraPayEnabled) {
		return true
	}

	return false
}

// SetIsSyncteraPayEnabled gets a reference to the given bool and assigns it to the IsSyncteraPayEnabled field.
func (o *RailsDepository) SetIsSyncteraPayEnabled(v bool) {
	o.IsSyncteraPayEnabled = &v
}

// GetIsWireEnabled returns the IsWireEnabled field value if set, zero value otherwise.
func (o *RailsDepository) GetIsWireEnabled() bool {
	if o == nil || IsNil(o.IsWireEnabled) {
		var ret bool
		return ret
	}
	return *o.IsWireEnabled
}

// GetIsWireEnabledOk returns a tuple with the IsWireEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RailsDepository) GetIsWireEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWireEnabled) {
		return nil, false
	}
	return o.IsWireEnabled, true
}

// HasIsWireEnabled returns a boolean if a field has been set.
func (o *RailsDepository) HasIsWireEnabled() bool {
	if o != nil && !IsNil(o.IsWireEnabled) {
		return true
	}

	return false
}

// SetIsWireEnabled gets a reference to the given bool and assigns it to the IsWireEnabled field.
func (o *RailsDepository) SetIsWireEnabled(v bool) {
	o.IsWireEnabled = &v
}

func (o RailsDepository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RailsDepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAchEnabled) {
		toSerialize["is_ach_enabled"] = o.IsAchEnabled
	}
	if !IsNil(o.IsCardEnabled) {
		toSerialize["is_card_enabled"] = o.IsCardEnabled
	}
	if !IsNil(o.IsEftCaEnabled) {
		toSerialize["is_eft_ca_enabled"] = o.IsEftCaEnabled
	}
	if !IsNil(o.IsExternalCardEnabled) {
		toSerialize["is_external_card_enabled"] = o.IsExternalCardEnabled
	}
	if !IsNil(o.IsP2pEnabled) {
		toSerialize["is_p2p_enabled"] = o.IsP2pEnabled
	}
	if !IsNil(o.IsSyncteraPayEnabled) {
		toSerialize["is_synctera_pay_enabled"] = o.IsSyncteraPayEnabled
	}
	if !IsNil(o.IsWireEnabled) {
		toSerialize["is_wire_enabled"] = o.IsWireEnabled
	}
	return toSerialize, nil
}

type NullableRailsDepository struct {
	value *RailsDepository
	isSet bool
}

func (v NullableRailsDepository) Get() *RailsDepository {
	return v.value
}

func (v *NullableRailsDepository) Set(val *RailsDepository) {
	v.value = val
	v.isSet = true
}

func (v NullableRailsDepository) IsSet() bool {
	return v.isSet
}

func (v *NullableRailsDepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRailsDepository(val *RailsDepository) *NullableRailsDepository {
	return &NullableRailsDepository{value: val, isSet: true}
}

func (v NullableRailsDepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRailsDepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
