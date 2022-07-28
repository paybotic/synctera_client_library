/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ExternalCardVerifications Verify card passed AVS and CVV checks in order to use PULL transfer. Verify card is eligible to receive push-to-card disbursements. For example, the inquiry_details object will contain a push_funds_block_indicator attribute that indicates if it is eligible for push-to-card disbursements.
type ExternalCardVerifications struct {
	// Address verification results  Status | Description --- | --- VERIFIED | Address verified NOT_VERIFIED | Address not verified ADDRESS_NO_MATCH | Postal/ZIP match, street addresses do not match or street address not included in request 
	AddressVerificationResults *string `json:"address_verification_results,omitempty"`
	// Indicates whether the external card is credit, debit, prepaid, deferred debit, or charge.
	CardType *string `json:"card_type,omitempty"`
	// Card Verification Value results  Status | Description --- | --- VERIFIED | CVV and expiration dates verified INCORRECT | Either CVV or expiration date is incorrect NOT_SUPPORTED |  Issuer does not participate in CVV2 service 
	Cvv2Result *string `json:"cvv2_result,omitempty"`
	// Indicates if card is Fast Funds eligible (i.e. if the funds will settle in 30 mins or less). If not eligible, typically funds will settle within 2 business days.
	FastFundsIndicator *bool `json:"fast_funds_indicator,omitempty"`
	// Indicates if the card can receive push-payments for online gambling payouts.
	OnlineGamblingBlockIndicator *bool `json:"online_gambling_block_indicator,omitempty"`
	// The name of the processor
	Processor *string `json:"processor,omitempty"`
	// Indicates if the associated card can receive push-to-card disbursements.
	PushFundsBlockIndicator *bool `json:"push_funds_block_indicator,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewExternalCardVerifications instantiates a new ExternalCardVerifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalCardVerifications() *ExternalCardVerifications {
	this := ExternalCardVerifications{}
	return &this
}

// NewExternalCardVerificationsWithDefaults instantiates a new ExternalCardVerifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalCardVerificationsWithDefaults() *ExternalCardVerifications {
	this := ExternalCardVerifications{}
	return &this
}

// GetAddressVerificationResults returns the AddressVerificationResults field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetAddressVerificationResults() string {
	if o == nil || o.AddressVerificationResults == nil {
		var ret string
		return ret
	}
	return *o.AddressVerificationResults
}

// GetAddressVerificationResultsOk returns a tuple with the AddressVerificationResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetAddressVerificationResultsOk() (*string, bool) {
	if o == nil || o.AddressVerificationResults == nil {
		return nil, false
	}
	return o.AddressVerificationResults, true
}

// HasAddressVerificationResults returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasAddressVerificationResults() bool {
	if o != nil && o.AddressVerificationResults != nil {
		return true
	}

	return false
}

// SetAddressVerificationResults gets a reference to the given string and assigns it to the AddressVerificationResults field.
func (o *ExternalCardVerifications) SetAddressVerificationResults(v string) {
	o.AddressVerificationResults = &v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetCardType() string {
	if o == nil || o.CardType == nil {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetCardTypeOk() (*string, bool) {
	if o == nil || o.CardType == nil {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasCardType() bool {
	if o != nil && o.CardType != nil {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *ExternalCardVerifications) SetCardType(v string) {
	o.CardType = &v
}

// GetCvv2Result returns the Cvv2Result field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetCvv2Result() string {
	if o == nil || o.Cvv2Result == nil {
		var ret string
		return ret
	}
	return *o.Cvv2Result
}

// GetCvv2ResultOk returns a tuple with the Cvv2Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetCvv2ResultOk() (*string, bool) {
	if o == nil || o.Cvv2Result == nil {
		return nil, false
	}
	return o.Cvv2Result, true
}

// HasCvv2Result returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasCvv2Result() bool {
	if o != nil && o.Cvv2Result != nil {
		return true
	}

	return false
}

// SetCvv2Result gets a reference to the given string and assigns it to the Cvv2Result field.
func (o *ExternalCardVerifications) SetCvv2Result(v string) {
	o.Cvv2Result = &v
}

// GetFastFundsIndicator returns the FastFundsIndicator field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetFastFundsIndicator() bool {
	if o == nil || o.FastFundsIndicator == nil {
		var ret bool
		return ret
	}
	return *o.FastFundsIndicator
}

// GetFastFundsIndicatorOk returns a tuple with the FastFundsIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetFastFundsIndicatorOk() (*bool, bool) {
	if o == nil || o.FastFundsIndicator == nil {
		return nil, false
	}
	return o.FastFundsIndicator, true
}

// HasFastFundsIndicator returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasFastFundsIndicator() bool {
	if o != nil && o.FastFundsIndicator != nil {
		return true
	}

	return false
}

// SetFastFundsIndicator gets a reference to the given bool and assigns it to the FastFundsIndicator field.
func (o *ExternalCardVerifications) SetFastFundsIndicator(v bool) {
	o.FastFundsIndicator = &v
}

// GetOnlineGamblingBlockIndicator returns the OnlineGamblingBlockIndicator field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetOnlineGamblingBlockIndicator() bool {
	if o == nil || o.OnlineGamblingBlockIndicator == nil {
		var ret bool
		return ret
	}
	return *o.OnlineGamblingBlockIndicator
}

// GetOnlineGamblingBlockIndicatorOk returns a tuple with the OnlineGamblingBlockIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetOnlineGamblingBlockIndicatorOk() (*bool, bool) {
	if o == nil || o.OnlineGamblingBlockIndicator == nil {
		return nil, false
	}
	return o.OnlineGamblingBlockIndicator, true
}

// HasOnlineGamblingBlockIndicator returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasOnlineGamblingBlockIndicator() bool {
	if o != nil && o.OnlineGamblingBlockIndicator != nil {
		return true
	}

	return false
}

// SetOnlineGamblingBlockIndicator gets a reference to the given bool and assigns it to the OnlineGamblingBlockIndicator field.
func (o *ExternalCardVerifications) SetOnlineGamblingBlockIndicator(v bool) {
	o.OnlineGamblingBlockIndicator = &v
}

// GetProcessor returns the Processor field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetProcessor() string {
	if o == nil || o.Processor == nil {
		var ret string
		return ret
	}
	return *o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetProcessorOk() (*string, bool) {
	if o == nil || o.Processor == nil {
		return nil, false
	}
	return o.Processor, true
}

// HasProcessor returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasProcessor() bool {
	if o != nil && o.Processor != nil {
		return true
	}

	return false
}

// SetProcessor gets a reference to the given string and assigns it to the Processor field.
func (o *ExternalCardVerifications) SetProcessor(v string) {
	o.Processor = &v
}

// GetPushFundsBlockIndicator returns the PushFundsBlockIndicator field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetPushFundsBlockIndicator() bool {
	if o == nil || o.PushFundsBlockIndicator == nil {
		var ret bool
		return ret
	}
	return *o.PushFundsBlockIndicator
}

// GetPushFundsBlockIndicatorOk returns a tuple with the PushFundsBlockIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetPushFundsBlockIndicatorOk() (*bool, bool) {
	if o == nil || o.PushFundsBlockIndicator == nil {
		return nil, false
	}
	return o.PushFundsBlockIndicator, true
}

// HasPushFundsBlockIndicator returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasPushFundsBlockIndicator() bool {
	if o != nil && o.PushFundsBlockIndicator != nil {
		return true
	}

	return false
}

// SetPushFundsBlockIndicator gets a reference to the given bool and assigns it to the PushFundsBlockIndicator field.
func (o *ExternalCardVerifications) SetPushFundsBlockIndicator(v bool) {
	o.PushFundsBlockIndicator = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ExternalCardVerifications) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardVerifications) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ExternalCardVerifications) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ExternalCardVerifications) SetState(v string) {
	o.State = &v
}

func (o ExternalCardVerifications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressVerificationResults != nil {
		toSerialize["address_verification_results"] = o.AddressVerificationResults
	}
	if o.CardType != nil {
		toSerialize["card_type"] = o.CardType
	}
	if o.Cvv2Result != nil {
		toSerialize["cvv2_result"] = o.Cvv2Result
	}
	if o.FastFundsIndicator != nil {
		toSerialize["fast_funds_indicator"] = o.FastFundsIndicator
	}
	if o.OnlineGamblingBlockIndicator != nil {
		toSerialize["online_gambling_block_indicator"] = o.OnlineGamblingBlockIndicator
	}
	if o.Processor != nil {
		toSerialize["processor"] = o.Processor
	}
	if o.PushFundsBlockIndicator != nil {
		toSerialize["push_funds_block_indicator"] = o.PushFundsBlockIndicator
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableExternalCardVerifications struct {
	value *ExternalCardVerifications
	isSet bool
}

func (v NullableExternalCardVerifications) Get() *ExternalCardVerifications {
	return v.value
}

func (v *NullableExternalCardVerifications) Set(val *ExternalCardVerifications) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalCardVerifications) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalCardVerifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalCardVerifications(val *ExternalCardVerifications) *NullableExternalCardVerifications {
	return &NullableExternalCardVerifications{value: val, isSet: true}
}

func (v NullableExternalCardVerifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalCardVerifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


