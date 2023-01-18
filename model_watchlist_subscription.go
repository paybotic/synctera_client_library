/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// WatchlistSubscription struct for WatchlistSubscription
type WatchlistSubscription struct {
	// Whether this subscription should automatically renew when the subscription period is over (default: vendor-dependent). 
	AutoRenew *bool `json:"auto_renew,omitempty"`
	// When this subscription was created
	Created *time.Time `json:"created,omitempty"`
	// Whether this customer has consented to being enrolled for watchlist monitoring 
	CustomerConsent bool `json:"customer_consent"`
	// Unique identifier for this subscription
	Id *string `json:"id,omitempty"`
	// The date when monitoring of this individual should end.
	PeriodEnd *string `json:"period_end,omitempty"`
	// The date when monitoring of this individual should begin (default: today).
	PeriodStart *string `json:"period_start,omitempty"`
	// External provider subscription id
	ProviderSubscriptionId *string `json:"provider_subscription_id,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewWatchlistSubscription instantiates a new WatchlistSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistSubscription(customerConsent bool) *WatchlistSubscription {
	this := WatchlistSubscription{}
	this.CustomerConsent = customerConsent
	return &this
}

// NewWatchlistSubscriptionWithDefaults instantiates a new WatchlistSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistSubscriptionWithDefaults() *WatchlistSubscription {
	this := WatchlistSubscription{}
	return &this
}

// GetAutoRenew returns the AutoRenew field value if set, zero value otherwise.
func (o *WatchlistSubscription) GetAutoRenew() bool {
	if o == nil || o.AutoRenew == nil {
		var ret bool
		return ret
	}
	return *o.AutoRenew
}

// GetAutoRenewOk returns a tuple with the AutoRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistSubscription) GetAutoRenewOk() (*bool, bool) {
	if o == nil || o.AutoRenew == nil {
		return nil, false
	}
	return o.AutoRenew, true
}

// HasAutoRenew returns a boolean if a field has been set.
func (o *WatchlistSubscription) HasAutoRenew() bool {
	if o != nil && o.AutoRenew != nil {
		return true
	}

	return false
}

// SetAutoRenew gets a reference to the given bool and assigns it to the AutoRenew field.
func (o *WatchlistSubscription) SetAutoRenew(v bool) {
	o.AutoRenew = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WatchlistSubscription) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistSubscription) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WatchlistSubscription) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *WatchlistSubscription) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCustomerConsent returns the CustomerConsent field value
func (o *WatchlistSubscription) GetCustomerConsent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CustomerConsent
}

// GetCustomerConsentOk returns a tuple with the CustomerConsent field value
// and a boolean to check if the value has been set.
func (o *WatchlistSubscription) GetCustomerConsentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerConsent, true
}

// SetCustomerConsent sets field value
func (o *WatchlistSubscription) SetCustomerConsent(v bool) {
	o.CustomerConsent = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WatchlistSubscription) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistSubscription) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WatchlistSubscription) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WatchlistSubscription) SetId(v string) {
	o.Id = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *WatchlistSubscription) GetPeriodEnd() string {
	if o == nil || o.PeriodEnd == nil {
		var ret string
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistSubscription) GetPeriodEndOk() (*string, bool) {
	if o == nil || o.PeriodEnd == nil {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *WatchlistSubscription) HasPeriodEnd() bool {
	if o != nil && o.PeriodEnd != nil {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given string and assigns it to the PeriodEnd field.
func (o *WatchlistSubscription) SetPeriodEnd(v string) {
	o.PeriodEnd = &v
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *WatchlistSubscription) GetPeriodStart() string {
	if o == nil || o.PeriodStart == nil {
		var ret string
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistSubscription) GetPeriodStartOk() (*string, bool) {
	if o == nil || o.PeriodStart == nil {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *WatchlistSubscription) HasPeriodStart() bool {
	if o != nil && o.PeriodStart != nil {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given string and assigns it to the PeriodStart field.
func (o *WatchlistSubscription) SetPeriodStart(v string) {
	o.PeriodStart = &v
}

// GetProviderSubscriptionId returns the ProviderSubscriptionId field value if set, zero value otherwise.
func (o *WatchlistSubscription) GetProviderSubscriptionId() string {
	if o == nil || o.ProviderSubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.ProviderSubscriptionId
}

// GetProviderSubscriptionIdOk returns a tuple with the ProviderSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistSubscription) GetProviderSubscriptionIdOk() (*string, bool) {
	if o == nil || o.ProviderSubscriptionId == nil {
		return nil, false
	}
	return o.ProviderSubscriptionId, true
}

// HasProviderSubscriptionId returns a boolean if a field has been set.
func (o *WatchlistSubscription) HasProviderSubscriptionId() bool {
	if o != nil && o.ProviderSubscriptionId != nil {
		return true
	}

	return false
}

// SetProviderSubscriptionId gets a reference to the given string and assigns it to the ProviderSubscriptionId field.
func (o *WatchlistSubscription) SetProviderSubscriptionId(v string) {
	o.ProviderSubscriptionId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WatchlistSubscription) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistSubscription) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WatchlistSubscription) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WatchlistSubscription) SetStatus(v string) {
	o.Status = &v
}

func (o WatchlistSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoRenew != nil {
		toSerialize["auto_renew"] = o.AutoRenew
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["customer_consent"] = o.CustomerConsent
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PeriodEnd != nil {
		toSerialize["period_end"] = o.PeriodEnd
	}
	if o.PeriodStart != nil {
		toSerialize["period_start"] = o.PeriodStart
	}
	if o.ProviderSubscriptionId != nil {
		toSerialize["provider_subscription_id"] = o.ProviderSubscriptionId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistSubscription struct {
	value *WatchlistSubscription
	isSet bool
}

func (v NullableWatchlistSubscription) Get() *WatchlistSubscription {
	return v.value
}

func (v *NullableWatchlistSubscription) Set(val *WatchlistSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistSubscription(val *WatchlistSubscription) *NullableWatchlistSubscription {
	return &NullableWatchlistSubscription{value: val, isSet: true}
}

func (v NullableWatchlistSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


