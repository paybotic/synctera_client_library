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

// checks if the CardImageDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardImageDetails{}

// CardImageDetails struct for CardImageDetails
type CardImageDetails struct {
	// The unique identifier of a cards product
	CardProductId string `json:"card_product_id"`
	// The unique identifier of a customer
	CustomerId string `json:"customer_id"`
	// The unique identifier of a card image
	Id                   string                    `json:"id"`
	RejectionMemo        *string                   `json:"rejection_memo,omitempty"`
	RejectionReason      *CardImageRejectionReason `json:"rejection_reason,omitempty"`
	Status               CardImageStatus           `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _CardImageDetails CardImageDetails

// NewCardImageDetails instantiates a new CardImageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardImageDetails(cardProductId string, customerId string, id string, status CardImageStatus) *CardImageDetails {
	this := CardImageDetails{}
	this.CardProductId = cardProductId
	this.CustomerId = customerId
	this.Id = id
	this.Status = status
	return &this
}

// NewCardImageDetailsWithDefaults instantiates a new CardImageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardImageDetailsWithDefaults() *CardImageDetails {
	this := CardImageDetails{}
	return &this
}

// GetCardProductId returns the CardProductId field value
func (o *CardImageDetails) GetCardProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardProductId
}

// GetCardProductIdOk returns a tuple with the CardProductId field value
// and a boolean to check if the value has been set.
func (o *CardImageDetails) GetCardProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProductId, true
}

// SetCardProductId sets field value
func (o *CardImageDetails) SetCardProductId(v string) {
	o.CardProductId = v
}

// GetCustomerId returns the CustomerId field value
func (o *CardImageDetails) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CardImageDetails) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CardImageDetails) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetId returns the Id field value
func (o *CardImageDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CardImageDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CardImageDetails) SetId(v string) {
	o.Id = v
}

// GetRejectionMemo returns the RejectionMemo field value if set, zero value otherwise.
func (o *CardImageDetails) GetRejectionMemo() string {
	if o == nil || IsNil(o.RejectionMemo) {
		var ret string
		return ret
	}
	return *o.RejectionMemo
}

// GetRejectionMemoOk returns a tuple with the RejectionMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardImageDetails) GetRejectionMemoOk() (*string, bool) {
	if o == nil || IsNil(o.RejectionMemo) {
		return nil, false
	}
	return o.RejectionMemo, true
}

// HasRejectionMemo returns a boolean if a field has been set.
func (o *CardImageDetails) HasRejectionMemo() bool {
	if o != nil && !IsNil(o.RejectionMemo) {
		return true
	}

	return false
}

// SetRejectionMemo gets a reference to the given string and assigns it to the RejectionMemo field.
func (o *CardImageDetails) SetRejectionMemo(v string) {
	o.RejectionMemo = &v
}

// GetRejectionReason returns the RejectionReason field value if set, zero value otherwise.
func (o *CardImageDetails) GetRejectionReason() CardImageRejectionReason {
	if o == nil || IsNil(o.RejectionReason) {
		var ret CardImageRejectionReason
		return ret
	}
	return *o.RejectionReason
}

// GetRejectionReasonOk returns a tuple with the RejectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardImageDetails) GetRejectionReasonOk() (*CardImageRejectionReason, bool) {
	if o == nil || IsNil(o.RejectionReason) {
		return nil, false
	}
	return o.RejectionReason, true
}

// HasRejectionReason returns a boolean if a field has been set.
func (o *CardImageDetails) HasRejectionReason() bool {
	if o != nil && !IsNil(o.RejectionReason) {
		return true
	}

	return false
}

// SetRejectionReason gets a reference to the given CardImageRejectionReason and assigns it to the RejectionReason field.
func (o *CardImageDetails) SetRejectionReason(v CardImageRejectionReason) {
	o.RejectionReason = &v
}

// GetStatus returns the Status field value
func (o *CardImageDetails) GetStatus() CardImageStatus {
	if o == nil {
		var ret CardImageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CardImageDetails) GetStatusOk() (*CardImageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CardImageDetails) SetStatus(v CardImageStatus) {
	o.Status = v
}

func (o CardImageDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardImageDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["card_product_id"] = o.CardProductId
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["id"] = o.Id
	if !IsNil(o.RejectionMemo) {
		toSerialize["rejection_memo"] = o.RejectionMemo
	}
	if !IsNil(o.RejectionReason) {
		toSerialize["rejection_reason"] = o.RejectionReason
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardImageDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"card_product_id",
		"customer_id",
		"id",
		"status",
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

	varCardImageDetails := _CardImageDetails{}

	err = json.Unmarshal(data, &varCardImageDetails)

	if err != nil {
		return err
	}

	*o = CardImageDetails(varCardImageDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "card_product_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "rejection_memo")
		delete(additionalProperties, "rejection_reason")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardImageDetails struct {
	value *CardImageDetails
	isSet bool
}

func (v NullableCardImageDetails) Get() *CardImageDetails {
	return v.value
}

func (v *NullableCardImageDetails) Set(val *CardImageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCardImageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCardImageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardImageDetails(val *CardImageDetails) *NullableCardImageDetails {
	return &NullableCardImageDetails{value: val, isSet: true}
}

func (v NullableCardImageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardImageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
