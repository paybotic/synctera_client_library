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

// CardTransactionDataPos Information about the point of sale
type CardTransactionDataPos struct {
	// Terminal Card data acceptance method
	CardDataInputCapability *string `json:"card_data_input_capability,omitempty"`
	// Cardholder presence
	CardHolderPresence *bool `json:"card_holder_presence,omitempty"`
	// Card presence
	CardPresence *bool `json:"card_presence,omitempty"`
	// Cardholder authentication method
	CardholderAuthenticationMethod *string `json:"cardholder_authentication_method,omitempty"`
	// Terminal country code
	CountryCode *string `json:"country_code,omitempty"`
	// Transaction is an installment payment
	IsInstallment *bool `json:"is_installment,omitempty"`
	// Transaction is recurring
	IsRecurring *bool `json:"is_recurring,omitempty"`
	// Card pan capture method
	PanEntryMode *string `json:"pan_entry_mode,omitempty"`
	// Terminal partial approval capability
	PartialApprovalCapable *bool `json:"partial_approval_capable,omitempty"`
	// Card pin capture method
	PinEntryMode *string `json:"pin_entry_mode,omitempty"`
	// Pin presence
	PinPresent *bool `json:"pin_present,omitempty"`
	// Terminal purchase amount only
	PurchaseAmountOnly *bool `json:"purchase_amount_only,omitempty"`
	// Terminal attendance
	TerminalAttendance *string `json:"terminal_attendance,omitempty"`
	// Terminal identifier
	TerminalId *string `json:"terminal_id,omitempty"`
	// Terminal location
	TerminalLocation *string `json:"terminal_location,omitempty"`
	// Terminal type
	TerminalType *string `json:"terminal_type,omitempty"`
	// Terminal zip code
	Zip *string `json:"zip,omitempty"`
}

// NewCardTransactionDataPos instantiates a new CardTransactionDataPos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardTransactionDataPos() *CardTransactionDataPos {
	this := CardTransactionDataPos{}
	return &this
}

// NewCardTransactionDataPosWithDefaults instantiates a new CardTransactionDataPos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardTransactionDataPosWithDefaults() *CardTransactionDataPos {
	this := CardTransactionDataPos{}
	return &this
}

// GetCardDataInputCapability returns the CardDataInputCapability field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetCardDataInputCapability() string {
	if o == nil || o.CardDataInputCapability == nil {
		var ret string
		return ret
	}
	return *o.CardDataInputCapability
}

// GetCardDataInputCapabilityOk returns a tuple with the CardDataInputCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetCardDataInputCapabilityOk() (*string, bool) {
	if o == nil || o.CardDataInputCapability == nil {
		return nil, false
	}
	return o.CardDataInputCapability, true
}

// HasCardDataInputCapability returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasCardDataInputCapability() bool {
	if o != nil && o.CardDataInputCapability != nil {
		return true
	}

	return false
}

// SetCardDataInputCapability gets a reference to the given string and assigns it to the CardDataInputCapability field.
func (o *CardTransactionDataPos) SetCardDataInputCapability(v string) {
	o.CardDataInputCapability = &v
}

// GetCardHolderPresence returns the CardHolderPresence field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetCardHolderPresence() bool {
	if o == nil || o.CardHolderPresence == nil {
		var ret bool
		return ret
	}
	return *o.CardHolderPresence
}

// GetCardHolderPresenceOk returns a tuple with the CardHolderPresence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetCardHolderPresenceOk() (*bool, bool) {
	if o == nil || o.CardHolderPresence == nil {
		return nil, false
	}
	return o.CardHolderPresence, true
}

// HasCardHolderPresence returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasCardHolderPresence() bool {
	if o != nil && o.CardHolderPresence != nil {
		return true
	}

	return false
}

// SetCardHolderPresence gets a reference to the given bool and assigns it to the CardHolderPresence field.
func (o *CardTransactionDataPos) SetCardHolderPresence(v bool) {
	o.CardHolderPresence = &v
}

// GetCardPresence returns the CardPresence field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetCardPresence() bool {
	if o == nil || o.CardPresence == nil {
		var ret bool
		return ret
	}
	return *o.CardPresence
}

// GetCardPresenceOk returns a tuple with the CardPresence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetCardPresenceOk() (*bool, bool) {
	if o == nil || o.CardPresence == nil {
		return nil, false
	}
	return o.CardPresence, true
}

// HasCardPresence returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasCardPresence() bool {
	if o != nil && o.CardPresence != nil {
		return true
	}

	return false
}

// SetCardPresence gets a reference to the given bool and assigns it to the CardPresence field.
func (o *CardTransactionDataPos) SetCardPresence(v bool) {
	o.CardPresence = &v
}

// GetCardholderAuthenticationMethod returns the CardholderAuthenticationMethod field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetCardholderAuthenticationMethod() string {
	if o == nil || o.CardholderAuthenticationMethod == nil {
		var ret string
		return ret
	}
	return *o.CardholderAuthenticationMethod
}

// GetCardholderAuthenticationMethodOk returns a tuple with the CardholderAuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetCardholderAuthenticationMethodOk() (*string, bool) {
	if o == nil || o.CardholderAuthenticationMethod == nil {
		return nil, false
	}
	return o.CardholderAuthenticationMethod, true
}

// HasCardholderAuthenticationMethod returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasCardholderAuthenticationMethod() bool {
	if o != nil && o.CardholderAuthenticationMethod != nil {
		return true
	}

	return false
}

// SetCardholderAuthenticationMethod gets a reference to the given string and assigns it to the CardholderAuthenticationMethod field.
func (o *CardTransactionDataPos) SetCardholderAuthenticationMethod(v string) {
	o.CardholderAuthenticationMethod = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *CardTransactionDataPos) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetIsInstallment returns the IsInstallment field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetIsInstallment() bool {
	if o == nil || o.IsInstallment == nil {
		var ret bool
		return ret
	}
	return *o.IsInstallment
}

// GetIsInstallmentOk returns a tuple with the IsInstallment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetIsInstallmentOk() (*bool, bool) {
	if o == nil || o.IsInstallment == nil {
		return nil, false
	}
	return o.IsInstallment, true
}

// HasIsInstallment returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasIsInstallment() bool {
	if o != nil && o.IsInstallment != nil {
		return true
	}

	return false
}

// SetIsInstallment gets a reference to the given bool and assigns it to the IsInstallment field.
func (o *CardTransactionDataPos) SetIsInstallment(v bool) {
	o.IsInstallment = &v
}

// GetIsRecurring returns the IsRecurring field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetIsRecurring() bool {
	if o == nil || o.IsRecurring == nil {
		var ret bool
		return ret
	}
	return *o.IsRecurring
}

// GetIsRecurringOk returns a tuple with the IsRecurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetIsRecurringOk() (*bool, bool) {
	if o == nil || o.IsRecurring == nil {
		return nil, false
	}
	return o.IsRecurring, true
}

// HasIsRecurring returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasIsRecurring() bool {
	if o != nil && o.IsRecurring != nil {
		return true
	}

	return false
}

// SetIsRecurring gets a reference to the given bool and assigns it to the IsRecurring field.
func (o *CardTransactionDataPos) SetIsRecurring(v bool) {
	o.IsRecurring = &v
}

// GetPanEntryMode returns the PanEntryMode field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetPanEntryMode() string {
	if o == nil || o.PanEntryMode == nil {
		var ret string
		return ret
	}
	return *o.PanEntryMode
}

// GetPanEntryModeOk returns a tuple with the PanEntryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetPanEntryModeOk() (*string, bool) {
	if o == nil || o.PanEntryMode == nil {
		return nil, false
	}
	return o.PanEntryMode, true
}

// HasPanEntryMode returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasPanEntryMode() bool {
	if o != nil && o.PanEntryMode != nil {
		return true
	}

	return false
}

// SetPanEntryMode gets a reference to the given string and assigns it to the PanEntryMode field.
func (o *CardTransactionDataPos) SetPanEntryMode(v string) {
	o.PanEntryMode = &v
}

// GetPartialApprovalCapable returns the PartialApprovalCapable field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetPartialApprovalCapable() bool {
	if o == nil || o.PartialApprovalCapable == nil {
		var ret bool
		return ret
	}
	return *o.PartialApprovalCapable
}

// GetPartialApprovalCapableOk returns a tuple with the PartialApprovalCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetPartialApprovalCapableOk() (*bool, bool) {
	if o == nil || o.PartialApprovalCapable == nil {
		return nil, false
	}
	return o.PartialApprovalCapable, true
}

// HasPartialApprovalCapable returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasPartialApprovalCapable() bool {
	if o != nil && o.PartialApprovalCapable != nil {
		return true
	}

	return false
}

// SetPartialApprovalCapable gets a reference to the given bool and assigns it to the PartialApprovalCapable field.
func (o *CardTransactionDataPos) SetPartialApprovalCapable(v bool) {
	o.PartialApprovalCapable = &v
}

// GetPinEntryMode returns the PinEntryMode field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetPinEntryMode() string {
	if o == nil || o.PinEntryMode == nil {
		var ret string
		return ret
	}
	return *o.PinEntryMode
}

// GetPinEntryModeOk returns a tuple with the PinEntryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetPinEntryModeOk() (*string, bool) {
	if o == nil || o.PinEntryMode == nil {
		return nil, false
	}
	return o.PinEntryMode, true
}

// HasPinEntryMode returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasPinEntryMode() bool {
	if o != nil && o.PinEntryMode != nil {
		return true
	}

	return false
}

// SetPinEntryMode gets a reference to the given string and assigns it to the PinEntryMode field.
func (o *CardTransactionDataPos) SetPinEntryMode(v string) {
	o.PinEntryMode = &v
}

// GetPinPresent returns the PinPresent field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetPinPresent() bool {
	if o == nil || o.PinPresent == nil {
		var ret bool
		return ret
	}
	return *o.PinPresent
}

// GetPinPresentOk returns a tuple with the PinPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetPinPresentOk() (*bool, bool) {
	if o == nil || o.PinPresent == nil {
		return nil, false
	}
	return o.PinPresent, true
}

// HasPinPresent returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasPinPresent() bool {
	if o != nil && o.PinPresent != nil {
		return true
	}

	return false
}

// SetPinPresent gets a reference to the given bool and assigns it to the PinPresent field.
func (o *CardTransactionDataPos) SetPinPresent(v bool) {
	o.PinPresent = &v
}

// GetPurchaseAmountOnly returns the PurchaseAmountOnly field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetPurchaseAmountOnly() bool {
	if o == nil || o.PurchaseAmountOnly == nil {
		var ret bool
		return ret
	}
	return *o.PurchaseAmountOnly
}

// GetPurchaseAmountOnlyOk returns a tuple with the PurchaseAmountOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetPurchaseAmountOnlyOk() (*bool, bool) {
	if o == nil || o.PurchaseAmountOnly == nil {
		return nil, false
	}
	return o.PurchaseAmountOnly, true
}

// HasPurchaseAmountOnly returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasPurchaseAmountOnly() bool {
	if o != nil && o.PurchaseAmountOnly != nil {
		return true
	}

	return false
}

// SetPurchaseAmountOnly gets a reference to the given bool and assigns it to the PurchaseAmountOnly field.
func (o *CardTransactionDataPos) SetPurchaseAmountOnly(v bool) {
	o.PurchaseAmountOnly = &v
}

// GetTerminalAttendance returns the TerminalAttendance field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetTerminalAttendance() string {
	if o == nil || o.TerminalAttendance == nil {
		var ret string
		return ret
	}
	return *o.TerminalAttendance
}

// GetTerminalAttendanceOk returns a tuple with the TerminalAttendance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetTerminalAttendanceOk() (*string, bool) {
	if o == nil || o.TerminalAttendance == nil {
		return nil, false
	}
	return o.TerminalAttendance, true
}

// HasTerminalAttendance returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasTerminalAttendance() bool {
	if o != nil && o.TerminalAttendance != nil {
		return true
	}

	return false
}

// SetTerminalAttendance gets a reference to the given string and assigns it to the TerminalAttendance field.
func (o *CardTransactionDataPos) SetTerminalAttendance(v string) {
	o.TerminalAttendance = &v
}

// GetTerminalId returns the TerminalId field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetTerminalId() string {
	if o == nil || o.TerminalId == nil {
		var ret string
		return ret
	}
	return *o.TerminalId
}

// GetTerminalIdOk returns a tuple with the TerminalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetTerminalIdOk() (*string, bool) {
	if o == nil || o.TerminalId == nil {
		return nil, false
	}
	return o.TerminalId, true
}

// HasTerminalId returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasTerminalId() bool {
	if o != nil && o.TerminalId != nil {
		return true
	}

	return false
}

// SetTerminalId gets a reference to the given string and assigns it to the TerminalId field.
func (o *CardTransactionDataPos) SetTerminalId(v string) {
	o.TerminalId = &v
}

// GetTerminalLocation returns the TerminalLocation field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetTerminalLocation() string {
	if o == nil || o.TerminalLocation == nil {
		var ret string
		return ret
	}
	return *o.TerminalLocation
}

// GetTerminalLocationOk returns a tuple with the TerminalLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetTerminalLocationOk() (*string, bool) {
	if o == nil || o.TerminalLocation == nil {
		return nil, false
	}
	return o.TerminalLocation, true
}

// HasTerminalLocation returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasTerminalLocation() bool {
	if o != nil && o.TerminalLocation != nil {
		return true
	}

	return false
}

// SetTerminalLocation gets a reference to the given string and assigns it to the TerminalLocation field.
func (o *CardTransactionDataPos) SetTerminalLocation(v string) {
	o.TerminalLocation = &v
}

// GetTerminalType returns the TerminalType field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetTerminalType() string {
	if o == nil || o.TerminalType == nil {
		var ret string
		return ret
	}
	return *o.TerminalType
}

// GetTerminalTypeOk returns a tuple with the TerminalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetTerminalTypeOk() (*string, bool) {
	if o == nil || o.TerminalType == nil {
		return nil, false
	}
	return o.TerminalType, true
}

// HasTerminalType returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasTerminalType() bool {
	if o != nil && o.TerminalType != nil {
		return true
	}

	return false
}

// SetTerminalType gets a reference to the given string and assigns it to the TerminalType field.
func (o *CardTransactionDataPos) SetTerminalType(v string) {
	o.TerminalType = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *CardTransactionDataPos) GetZip() string {
	if o == nil || o.Zip == nil {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTransactionDataPos) GetZipOk() (*string, bool) {
	if o == nil || o.Zip == nil {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *CardTransactionDataPos) HasZip() bool {
	if o != nil && o.Zip != nil {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *CardTransactionDataPos) SetZip(v string) {
	o.Zip = &v
}

func (o CardTransactionDataPos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CardDataInputCapability != nil {
		toSerialize["card_data_input_capability"] = o.CardDataInputCapability
	}
	if o.CardHolderPresence != nil {
		toSerialize["card_holder_presence"] = o.CardHolderPresence
	}
	if o.CardPresence != nil {
		toSerialize["card_presence"] = o.CardPresence
	}
	if o.CardholderAuthenticationMethod != nil {
		toSerialize["cardholder_authentication_method"] = o.CardholderAuthenticationMethod
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.IsInstallment != nil {
		toSerialize["is_installment"] = o.IsInstallment
	}
	if o.IsRecurring != nil {
		toSerialize["is_recurring"] = o.IsRecurring
	}
	if o.PanEntryMode != nil {
		toSerialize["pan_entry_mode"] = o.PanEntryMode
	}
	if o.PartialApprovalCapable != nil {
		toSerialize["partial_approval_capable"] = o.PartialApprovalCapable
	}
	if o.PinEntryMode != nil {
		toSerialize["pin_entry_mode"] = o.PinEntryMode
	}
	if o.PinPresent != nil {
		toSerialize["pin_present"] = o.PinPresent
	}
	if o.PurchaseAmountOnly != nil {
		toSerialize["purchase_amount_only"] = o.PurchaseAmountOnly
	}
	if o.TerminalAttendance != nil {
		toSerialize["terminal_attendance"] = o.TerminalAttendance
	}
	if o.TerminalId != nil {
		toSerialize["terminal_id"] = o.TerminalId
	}
	if o.TerminalLocation != nil {
		toSerialize["terminal_location"] = o.TerminalLocation
	}
	if o.TerminalType != nil {
		toSerialize["terminal_type"] = o.TerminalType
	}
	if o.Zip != nil {
		toSerialize["zip"] = o.Zip
	}
	return json.Marshal(toSerialize)
}

type NullableCardTransactionDataPos struct {
	value *CardTransactionDataPos
	isSet bool
}

func (v NullableCardTransactionDataPos) Get() *CardTransactionDataPos {
	return v.value
}

func (v *NullableCardTransactionDataPos) Set(val *CardTransactionDataPos) {
	v.value = val
	v.isSet = true
}

func (v NullableCardTransactionDataPos) IsSet() bool {
	return v.isSet
}

func (v *NullableCardTransactionDataPos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardTransactionDataPos(val *CardTransactionDataPos) *NullableCardTransactionDataPos {
	return &NullableCardTransactionDataPos{value: val, isSet: true}
}

func (v NullableCardTransactionDataPos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardTransactionDataPos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

