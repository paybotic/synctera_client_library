/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CashOrderAuthorizationPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashOrderAuthorizationPatch{}

// CashOrderAuthorizationPatch Request body to patch a cash order authorization
type CashOrderAuthorizationPatch struct {
	// The account number of the client business the cash order is for. Known as \"Location ID\" in the cash order CSV files.
	AccountNumber     *string               `json:"account_number,omitempty"`
	AuthorizationType CashAuthorizationType `json:"authorization_type"`
	// The name of the client business the cash order is for.
	ClientName *string `json:"client_name,omitempty"`
	// The date the cash order was placed with Empyreal
	OrderDate *string `json:"order_date,omitempty"`
	Status    *string `json:"status,omitempty"`
}

type _CashOrderAuthorizationPatch CashOrderAuthorizationPatch

// NewCashOrderAuthorizationPatch instantiates a new CashOrderAuthorizationPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashOrderAuthorizationPatch(authorizationType CashAuthorizationType) *CashOrderAuthorizationPatch {
	this := CashOrderAuthorizationPatch{}
	this.AuthorizationType = authorizationType
	return &this
}

// NewCashOrderAuthorizationPatchWithDefaults instantiates a new CashOrderAuthorizationPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashOrderAuthorizationPatchWithDefaults() *CashOrderAuthorizationPatch {
	this := CashOrderAuthorizationPatch{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *CashOrderAuthorizationPatch) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashOrderAuthorizationPatch) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *CashOrderAuthorizationPatch) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *CashOrderAuthorizationPatch) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAuthorizationType returns the AuthorizationType field value
func (o *CashOrderAuthorizationPatch) GetAuthorizationType() CashAuthorizationType {
	if o == nil {
		var ret CashAuthorizationType
		return ret
	}

	return o.AuthorizationType
}

// GetAuthorizationTypeOk returns a tuple with the AuthorizationType field value
// and a boolean to check if the value has been set.
func (o *CashOrderAuthorizationPatch) GetAuthorizationTypeOk() (*CashAuthorizationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationType, true
}

// SetAuthorizationType sets field value
func (o *CashOrderAuthorizationPatch) SetAuthorizationType(v CashAuthorizationType) {
	o.AuthorizationType = v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *CashOrderAuthorizationPatch) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashOrderAuthorizationPatch) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *CashOrderAuthorizationPatch) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *CashOrderAuthorizationPatch) SetClientName(v string) {
	o.ClientName = &v
}

// GetOrderDate returns the OrderDate field value if set, zero value otherwise.
func (o *CashOrderAuthorizationPatch) GetOrderDate() string {
	if o == nil || IsNil(o.OrderDate) {
		var ret string
		return ret
	}
	return *o.OrderDate
}

// GetOrderDateOk returns a tuple with the OrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashOrderAuthorizationPatch) GetOrderDateOk() (*string, bool) {
	if o == nil || IsNil(o.OrderDate) {
		return nil, false
	}
	return o.OrderDate, true
}

// HasOrderDate returns a boolean if a field has been set.
func (o *CashOrderAuthorizationPatch) HasOrderDate() bool {
	if o != nil && !IsNil(o.OrderDate) {
		return true
	}

	return false
}

// SetOrderDate gets a reference to the given string and assigns it to the OrderDate field.
func (o *CashOrderAuthorizationPatch) SetOrderDate(v string) {
	o.OrderDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CashOrderAuthorizationPatch) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashOrderAuthorizationPatch) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CashOrderAuthorizationPatch) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CashOrderAuthorizationPatch) SetStatus(v string) {
	o.Status = &v
}

func (o CashOrderAuthorizationPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashOrderAuthorizationPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountNumber) {
		toSerialize["account_number"] = o.AccountNumber
	}
	toSerialize["authorization_type"] = o.AuthorizationType
	if !IsNil(o.ClientName) {
		toSerialize["client_name"] = o.ClientName
	}
	if !IsNil(o.OrderDate) {
		toSerialize["order_date"] = o.OrderDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *CashOrderAuthorizationPatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorization_type",
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

	varCashOrderAuthorizationPatch := _CashOrderAuthorizationPatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCashOrderAuthorizationPatch)

	if err != nil {
		return err
	}

	*o = CashOrderAuthorizationPatch(varCashOrderAuthorizationPatch)

	return err
}

type NullableCashOrderAuthorizationPatch struct {
	value *CashOrderAuthorizationPatch
	isSet bool
}

func (v NullableCashOrderAuthorizationPatch) Get() *CashOrderAuthorizationPatch {
	return v.value
}

func (v *NullableCashOrderAuthorizationPatch) Set(val *CashOrderAuthorizationPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableCashOrderAuthorizationPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableCashOrderAuthorizationPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashOrderAuthorizationPatch(val *CashOrderAuthorizationPatch) *NullableCashOrderAuthorizationPatch {
	return &NullableCashOrderAuthorizationPatch{value: val, isSet: true}
}

func (v NullableCashOrderAuthorizationPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashOrderAuthorizationPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}