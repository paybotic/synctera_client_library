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

// checks if the CashOrderAuthorizationPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashOrderAuthorizationPost{}

// CashOrderAuthorizationPost struct for CashOrderAuthorizationPost
type CashOrderAuthorizationPost struct {
	// Transfer amount in cents
	Amount            int64                 `json:"amount"`
	AuthorizationType CashAuthorizationType `json:"authorization_type"`
	// The UUID of the Synctera account resource that is the destination of the transfer.
	DestinationAccountId string `json:"destination_account_id"`
	// The date the cash order was placed with cash distribution provider
	OrderDate            string `json:"order_date"`
	AdditionalProperties map[string]interface{}
}

type _CashOrderAuthorizationPost CashOrderAuthorizationPost

// NewCashOrderAuthorizationPost instantiates a new CashOrderAuthorizationPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashOrderAuthorizationPost(amount int64, authorizationType CashAuthorizationType, destinationAccountId string, orderDate string) *CashOrderAuthorizationPost {
	this := CashOrderAuthorizationPost{}
	this.Amount = amount
	this.AuthorizationType = authorizationType
	this.DestinationAccountId = destinationAccountId
	this.OrderDate = orderDate
	return &this
}

// NewCashOrderAuthorizationPostWithDefaults instantiates a new CashOrderAuthorizationPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashOrderAuthorizationPostWithDefaults() *CashOrderAuthorizationPost {
	this := CashOrderAuthorizationPost{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CashOrderAuthorizationPost) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CashOrderAuthorizationPost) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CashOrderAuthorizationPost) SetAmount(v int64) {
	o.Amount = v
}

// GetAuthorizationType returns the AuthorizationType field value
func (o *CashOrderAuthorizationPost) GetAuthorizationType() CashAuthorizationType {
	if o == nil {
		var ret CashAuthorizationType
		return ret
	}

	return o.AuthorizationType
}

// GetAuthorizationTypeOk returns a tuple with the AuthorizationType field value
// and a boolean to check if the value has been set.
func (o *CashOrderAuthorizationPost) GetAuthorizationTypeOk() (*CashAuthorizationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationType, true
}

// SetAuthorizationType sets field value
func (o *CashOrderAuthorizationPost) SetAuthorizationType(v CashAuthorizationType) {
	o.AuthorizationType = v
}

// GetDestinationAccountId returns the DestinationAccountId field value
func (o *CashOrderAuthorizationPost) GetDestinationAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationAccountId
}

// GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field value
// and a boolean to check if the value has been set.
func (o *CashOrderAuthorizationPost) GetDestinationAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationAccountId, true
}

// SetDestinationAccountId sets field value
func (o *CashOrderAuthorizationPost) SetDestinationAccountId(v string) {
	o.DestinationAccountId = v
}

// GetOrderDate returns the OrderDate field value
func (o *CashOrderAuthorizationPost) GetOrderDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderDate
}

// GetOrderDateOk returns a tuple with the OrderDate field value
// and a boolean to check if the value has been set.
func (o *CashOrderAuthorizationPost) GetOrderDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderDate, true
}

// SetOrderDate sets field value
func (o *CashOrderAuthorizationPost) SetOrderDate(v string) {
	o.OrderDate = v
}

func (o CashOrderAuthorizationPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashOrderAuthorizationPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["authorization_type"] = o.AuthorizationType
	toSerialize["destination_account_id"] = o.DestinationAccountId
	toSerialize["order_date"] = o.OrderDate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CashOrderAuthorizationPost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"authorization_type",
		"destination_account_id",
		"order_date",
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

	varCashOrderAuthorizationPost := _CashOrderAuthorizationPost{}

	err = json.Unmarshal(data, &varCashOrderAuthorizationPost)

	if err != nil {
		return err
	}

	*o = CashOrderAuthorizationPost(varCashOrderAuthorizationPost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "authorization_type")
		delete(additionalProperties, "destination_account_id")
		delete(additionalProperties, "order_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCashOrderAuthorizationPost struct {
	value *CashOrderAuthorizationPost
	isSet bool
}

func (v NullableCashOrderAuthorizationPost) Get() *CashOrderAuthorizationPost {
	return v.value
}

func (v *NullableCashOrderAuthorizationPost) Set(val *CashOrderAuthorizationPost) {
	v.value = val
	v.isSet = true
}

func (v NullableCashOrderAuthorizationPost) IsSet() bool {
	return v.isSet
}

func (v *NullableCashOrderAuthorizationPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashOrderAuthorizationPost(val *CashOrderAuthorizationPost) *NullableCashOrderAuthorizationPost {
	return &NullableCashOrderAuthorizationPost{value: val, isSet: true}
}

func (v NullableCashOrderAuthorizationPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashOrderAuthorizationPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
