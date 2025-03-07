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

// checks if the TemplateFieldsChargeSecured type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateFieldsChargeSecured{}

// TemplateFieldsChargeSecured For creating CHARGE_SECURED accounts, e.g. for use in a Smart Charge Card offering.
type TemplateFieldsChargeSecured struct {
	AccountType AccountType `json:"account_type"`
	// The bank account ID for this account. This is a unique identifier for the bank side account that this Synctera account belongs to. This field can be auto filled if only one bank account of the appropriate type exist for the tenant of concern.
	BankAccountId *string `json:"bank_account_id,omitempty"`
	// Bank country of the account. ISO 3166-1 Alpha-2 or Alpha-3 country code.
	BankCountry string `json:"bank_country" validate:"regexp=^[A-Z]{2,3}$"`
	// Account currency. ISO 4217 alphabetic currency code
	Currency string `json:"currency" validate:"regexp=^[A-Z]{3}$"`
	// The number of days past the billing period to initiate an auto payment. Only applicable for accounts with type `CHARGE_SECURED`, where the account holder has opted in for auto payment functionality. This value must be lower than or equal the `grace_period` setting on the account. If this value is 0, the auto payment will happen on the same day as the statement is generated. Auto payment only occurs if regular payments are not received on time.
	AutoPaymentPeriod *int32 `json:"auto_payment_period,omitempty"`
	// The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment. The default will be set to 21 days.
	GracePeriod    *int32             `json:"grace_period,omitempty"`
	MinimumPayment MinimumPaymentFull `json:"minimum_payment"`
	// List of spend control IDs to control spending for the account
	SpendControlIds      []string `json:"spend_control_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateFieldsChargeSecured TemplateFieldsChargeSecured

// NewTemplateFieldsChargeSecured instantiates a new TemplateFieldsChargeSecured object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateFieldsChargeSecured(accountType AccountType, bankCountry string, currency string, minimumPayment MinimumPaymentFull) *TemplateFieldsChargeSecured {
	this := TemplateFieldsChargeSecured{}
	this.AccountType = accountType
	this.BankCountry = bankCountry
	this.Currency = currency
	var gracePeriod int32 = 21
	this.GracePeriod = &gracePeriod
	this.MinimumPayment = minimumPayment
	return &this
}

// NewTemplateFieldsChargeSecuredWithDefaults instantiates a new TemplateFieldsChargeSecured object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateFieldsChargeSecuredWithDefaults() *TemplateFieldsChargeSecured {
	this := TemplateFieldsChargeSecured{}
	var gracePeriod int32 = 21
	this.GracePeriod = &gracePeriod
	return &this
}

// GetAccountType returns the AccountType field value
func (o *TemplateFieldsChargeSecured) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsChargeSecured) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *TemplateFieldsChargeSecured) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetBankAccountId returns the BankAccountId field value if set, zero value otherwise.
func (o *TemplateFieldsChargeSecured) GetBankAccountId() string {
	if o == nil || IsNil(o.BankAccountId) {
		var ret string
		return ret
	}
	return *o.BankAccountId
}

// GetBankAccountIdOk returns a tuple with the BankAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsChargeSecured) GetBankAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountId) {
		return nil, false
	}
	return o.BankAccountId, true
}

// HasBankAccountId returns a boolean if a field has been set.
func (o *TemplateFieldsChargeSecured) HasBankAccountId() bool {
	if o != nil && !IsNil(o.BankAccountId) {
		return true
	}

	return false
}

// SetBankAccountId gets a reference to the given string and assigns it to the BankAccountId field.
func (o *TemplateFieldsChargeSecured) SetBankAccountId(v string) {
	o.BankAccountId = &v
}

// GetBankCountry returns the BankCountry field value
func (o *TemplateFieldsChargeSecured) GetBankCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCountry
}

// GetBankCountryOk returns a tuple with the BankCountry field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsChargeSecured) GetBankCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCountry, true
}

// SetBankCountry sets field value
func (o *TemplateFieldsChargeSecured) SetBankCountry(v string) {
	o.BankCountry = v
}

// GetCurrency returns the Currency field value
func (o *TemplateFieldsChargeSecured) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsChargeSecured) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TemplateFieldsChargeSecured) SetCurrency(v string) {
	o.Currency = v
}

// GetAutoPaymentPeriod returns the AutoPaymentPeriod field value if set, zero value otherwise.
func (o *TemplateFieldsChargeSecured) GetAutoPaymentPeriod() int32 {
	if o == nil || IsNil(o.AutoPaymentPeriod) {
		var ret int32
		return ret
	}
	return *o.AutoPaymentPeriod
}

// GetAutoPaymentPeriodOk returns a tuple with the AutoPaymentPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsChargeSecured) GetAutoPaymentPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoPaymentPeriod) {
		return nil, false
	}
	return o.AutoPaymentPeriod, true
}

// HasAutoPaymentPeriod returns a boolean if a field has been set.
func (o *TemplateFieldsChargeSecured) HasAutoPaymentPeriod() bool {
	if o != nil && !IsNil(o.AutoPaymentPeriod) {
		return true
	}

	return false
}

// SetAutoPaymentPeriod gets a reference to the given int32 and assigns it to the AutoPaymentPeriod field.
func (o *TemplateFieldsChargeSecured) SetAutoPaymentPeriod(v int32) {
	o.AutoPaymentPeriod = &v
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *TemplateFieldsChargeSecured) GetGracePeriod() int32 {
	if o == nil || IsNil(o.GracePeriod) {
		var ret int32
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsChargeSecured) GetGracePeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.GracePeriod) {
		return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *TemplateFieldsChargeSecured) HasGracePeriod() bool {
	if o != nil && !IsNil(o.GracePeriod) {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given int32 and assigns it to the GracePeriod field.
func (o *TemplateFieldsChargeSecured) SetGracePeriod(v int32) {
	o.GracePeriod = &v
}

// GetMinimumPayment returns the MinimumPayment field value
func (o *TemplateFieldsChargeSecured) GetMinimumPayment() MinimumPaymentFull {
	if o == nil {
		var ret MinimumPaymentFull
		return ret
	}

	return o.MinimumPayment
}

// GetMinimumPaymentOk returns a tuple with the MinimumPayment field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsChargeSecured) GetMinimumPaymentOk() (*MinimumPaymentFull, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumPayment, true
}

// SetMinimumPayment sets field value
func (o *TemplateFieldsChargeSecured) SetMinimumPayment(v MinimumPaymentFull) {
	o.MinimumPayment = v
}

// GetSpendControlIds returns the SpendControlIds field value if set, zero value otherwise.
func (o *TemplateFieldsChargeSecured) GetSpendControlIds() []string {
	if o == nil || IsNil(o.SpendControlIds) {
		var ret []string
		return ret
	}
	return o.SpendControlIds
}

// GetSpendControlIdsOk returns a tuple with the SpendControlIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsChargeSecured) GetSpendControlIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SpendControlIds) {
		return nil, false
	}
	return o.SpendControlIds, true
}

// HasSpendControlIds returns a boolean if a field has been set.
func (o *TemplateFieldsChargeSecured) HasSpendControlIds() bool {
	if o != nil && !IsNil(o.SpendControlIds) {
		return true
	}

	return false
}

// SetSpendControlIds gets a reference to the given []string and assigns it to the SpendControlIds field.
func (o *TemplateFieldsChargeSecured) SetSpendControlIds(v []string) {
	o.SpendControlIds = v
}

func (o TemplateFieldsChargeSecured) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateFieldsChargeSecured) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_type"] = o.AccountType
	if !IsNil(o.BankAccountId) {
		toSerialize["bank_account_id"] = o.BankAccountId
	}
	toSerialize["bank_country"] = o.BankCountry
	toSerialize["currency"] = o.Currency
	if !IsNil(o.AutoPaymentPeriod) {
		toSerialize["auto_payment_period"] = o.AutoPaymentPeriod
	}
	if !IsNil(o.GracePeriod) {
		toSerialize["grace_period"] = o.GracePeriod
	}
	toSerialize["minimum_payment"] = o.MinimumPayment
	if !IsNil(o.SpendControlIds) {
		toSerialize["spend_control_ids"] = o.SpendControlIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TemplateFieldsChargeSecured) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_type",
		"bank_country",
		"currency",
		"minimum_payment",
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

	varTemplateFieldsChargeSecured := _TemplateFieldsChargeSecured{}

	err = json.Unmarshal(data, &varTemplateFieldsChargeSecured)

	if err != nil {
		return err
	}

	*o = TemplateFieldsChargeSecured(varTemplateFieldsChargeSecured)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_type")
		delete(additionalProperties, "bank_account_id")
		delete(additionalProperties, "bank_country")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "auto_payment_period")
		delete(additionalProperties, "grace_period")
		delete(additionalProperties, "minimum_payment")
		delete(additionalProperties, "spend_control_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateFieldsChargeSecured struct {
	value *TemplateFieldsChargeSecured
	isSet bool
}

func (v NullableTemplateFieldsChargeSecured) Get() *TemplateFieldsChargeSecured {
	return v.value
}

func (v *NullableTemplateFieldsChargeSecured) Set(val *TemplateFieldsChargeSecured) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFieldsChargeSecured) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFieldsChargeSecured) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFieldsChargeSecured(val *TemplateFieldsChargeSecured) *NullableTemplateFieldsChargeSecured {
	return &NullableTemplateFieldsChargeSecured{value: val, isSet: true}
}

func (v NullableTemplateFieldsChargeSecured) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFieldsChargeSecured) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
