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
	"time"
)

// checks if the VirtualCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualCard{}

// VirtualCard A virtual card
type VirtualCard struct {
	// The ID of the account to which the card will be linked
	AccountId *string `json:"account_id,omitempty"`
	// The business ID associated with this card. If no customer_id is supplied, a card can still be issued to a business, but cannot be activated or used until a customer is assigned via the PATCH /cards/{card_id} endpoint.
	BusinessId *string `json:"business_id,omitempty"`
	// The card product to which the card is attached
	CardProductId *string `json:"card_product_id,omitempty"`
	// The timestamp representing when the card issuance request was made
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The ID of the customer to whom the card will be issued. If a business_id is passed, but a customer_id not passed at the time of card creation the card cannot be activated or used for spend until it's assigned to a human customer via the PATCH /cards/{card_id} endpoint. If no business_id is passed, a customer_id is required.
	CustomerId      *string     `json:"customer_id,omitempty"`
	EmbossName      *EmbossName `json:"emboss_name,omitempty"`
	ExpirationMonth *string     `json:"expiration_month,omitempty"`
	// The timestamp representing when the card would expire at
	ExpirationTime *time.Time `json:"expiration_time,omitempty"`
	ExpirationYear *string    `json:"expiration_year,omitempty"`
	// Card ID
	Id *string `json:"id,omitempty"`
	// indicates whether a pin has been set on the card
	IsPinSet *bool `json:"is_pin_set,omitempty"`
	// The last 4 digits of the card PAN
	LastFour *string `json:"last_four,omitempty"`
	// The timestamp representing when the card was last modified at
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// Additional data to include in the request structured as key-value pairs
	Metadata *map[string]string `json:"metadata,omitempty"`
	// This is the reason the card needs to be reissued, if any. The reason determines several behaviours:   - whether or not the new card will use the same PAN as the original card   - the old card will be terminated and if so, when it will be terminated  Reason                 | Same PAN | Terminate Old Card ---------------------- | -------- | ------------------ EXPIRATION             | yes      | on activation LOST                   | no       | immediately STOLEN                 | no       | immediately DAMAGED                | yes      | on activation VIRTUAL_TO_PHYSICAL(*) | yes      | on activation PRODUCT_CHANGE         | yes      | on activation NAME_CHANGE(**)        | yes      | on activation APPEARANCE             | yes      | on activation  (*) VIRTUAL_TO_PHYSICAL is deprecated. Please use PRODUCT_CHANGE whenever reissuing from one card product to another, including from a virtual product to a physical product.  (**) NAME_CHANGE is deprecated. Please use APPEARANCE whenever reissuing in order to change the appearance of a card, such as the printed name or custom image.  For all reasons, the new card will use the same PIN as the original card and digital wallet tokens will reassigned to the new card
	ReissueReason *string `json:"reissue_reason,omitempty"`
	// When reissuing a card, specify the card to be replaced here. When getting a card's details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced. If this field is set, then reissue_reason must also be set.
	ReissuedFromId *string `json:"reissued_from_id,omitempty"`
	// If this card was reissued, this ID refers to the card that replaced it.
	ReissuedToId *string `json:"reissued_to_id,omitempty"`
	// Time when the PIN was last set or changed.
	TimestampPinSet *time.Time `json:"timestamp_pin_set,omitempty"`
	Type            *CardType  `json:"type,omitempty"`
	// PHYSICAL or VIRTUAL.
	Form string `json:"form"`
}

type _VirtualCard VirtualCard

// NewVirtualCard instantiates a new VirtualCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualCard(form string) *VirtualCard {
	this := VirtualCard{}
	this.Form = form
	return &this
}

// NewVirtualCardWithDefaults instantiates a new VirtualCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualCardWithDefaults() *VirtualCard {
	this := VirtualCard{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *VirtualCard) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *VirtualCard) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *VirtualCard) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *VirtualCard) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *VirtualCard) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *VirtualCard) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCardProductId returns the CardProductId field value if set, zero value otherwise.
func (o *VirtualCard) GetCardProductId() string {
	if o == nil || IsNil(o.CardProductId) {
		var ret string
		return ret
	}
	return *o.CardProductId
}

// GetCardProductIdOk returns a tuple with the CardProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetCardProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.CardProductId) {
		return nil, false
	}
	return o.CardProductId, true
}

// HasCardProductId returns a boolean if a field has been set.
func (o *VirtualCard) HasCardProductId() bool {
	if o != nil && !IsNil(o.CardProductId) {
		return true
	}

	return false
}

// SetCardProductId gets a reference to the given string and assigns it to the CardProductId field.
func (o *VirtualCard) SetCardProductId(v string) {
	o.CardProductId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *VirtualCard) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *VirtualCard) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *VirtualCard) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *VirtualCard) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *VirtualCard) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *VirtualCard) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEmbossName returns the EmbossName field value if set, zero value otherwise.
func (o *VirtualCard) GetEmbossName() EmbossName {
	if o == nil || IsNil(o.EmbossName) {
		var ret EmbossName
		return ret
	}
	return *o.EmbossName
}

// GetEmbossNameOk returns a tuple with the EmbossName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetEmbossNameOk() (*EmbossName, bool) {
	if o == nil || IsNil(o.EmbossName) {
		return nil, false
	}
	return o.EmbossName, true
}

// HasEmbossName returns a boolean if a field has been set.
func (o *VirtualCard) HasEmbossName() bool {
	if o != nil && !IsNil(o.EmbossName) {
		return true
	}

	return false
}

// SetEmbossName gets a reference to the given EmbossName and assigns it to the EmbossName field.
func (o *VirtualCard) SetEmbossName(v EmbossName) {
	o.EmbossName = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *VirtualCard) GetExpirationMonth() string {
	if o == nil || IsNil(o.ExpirationMonth) {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetExpirationMonthOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationMonth) {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *VirtualCard) HasExpirationMonth() bool {
	if o != nil && !IsNil(o.ExpirationMonth) {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *VirtualCard) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *VirtualCard) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *VirtualCard) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *VirtualCard) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *VirtualCard) GetExpirationYear() string {
	if o == nil || IsNil(o.ExpirationYear) {
		var ret string
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetExpirationYearOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationYear) {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *VirtualCard) HasExpirationYear() bool {
	if o != nil && !IsNil(o.ExpirationYear) {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given string and assigns it to the ExpirationYear field.
func (o *VirtualCard) SetExpirationYear(v string) {
	o.ExpirationYear = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualCard) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualCard) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VirtualCard) SetId(v string) {
	o.Id = &v
}

// GetIsPinSet returns the IsPinSet field value if set, zero value otherwise.
func (o *VirtualCard) GetIsPinSet() bool {
	if o == nil || IsNil(o.IsPinSet) {
		var ret bool
		return ret
	}
	return *o.IsPinSet
}

// GetIsPinSetOk returns a tuple with the IsPinSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetIsPinSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPinSet) {
		return nil, false
	}
	return o.IsPinSet, true
}

// HasIsPinSet returns a boolean if a field has been set.
func (o *VirtualCard) HasIsPinSet() bool {
	if o != nil && !IsNil(o.IsPinSet) {
		return true
	}

	return false
}

// SetIsPinSet gets a reference to the given bool and assigns it to the IsPinSet field.
func (o *VirtualCard) SetIsPinSet(v bool) {
	o.IsPinSet = &v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *VirtualCard) GetLastFour() string {
	if o == nil || IsNil(o.LastFour) {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetLastFourOk() (*string, bool) {
	if o == nil || IsNil(o.LastFour) {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *VirtualCard) HasLastFour() bool {
	if o != nil && !IsNil(o.LastFour) {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *VirtualCard) SetLastFour(v string) {
	o.LastFour = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *VirtualCard) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *VirtualCard) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *VirtualCard) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *VirtualCard) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *VirtualCard) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *VirtualCard) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetReissueReason returns the ReissueReason field value if set, zero value otherwise.
func (o *VirtualCard) GetReissueReason() string {
	if o == nil || IsNil(o.ReissueReason) {
		var ret string
		return ret
	}
	return *o.ReissueReason
}

// GetReissueReasonOk returns a tuple with the ReissueReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetReissueReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ReissueReason) {
		return nil, false
	}
	return o.ReissueReason, true
}

// HasReissueReason returns a boolean if a field has been set.
func (o *VirtualCard) HasReissueReason() bool {
	if o != nil && !IsNil(o.ReissueReason) {
		return true
	}

	return false
}

// SetReissueReason gets a reference to the given string and assigns it to the ReissueReason field.
func (o *VirtualCard) SetReissueReason(v string) {
	o.ReissueReason = &v
}

// GetReissuedFromId returns the ReissuedFromId field value if set, zero value otherwise.
func (o *VirtualCard) GetReissuedFromId() string {
	if o == nil || IsNil(o.ReissuedFromId) {
		var ret string
		return ret
	}
	return *o.ReissuedFromId
}

// GetReissuedFromIdOk returns a tuple with the ReissuedFromId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetReissuedFromIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReissuedFromId) {
		return nil, false
	}
	return o.ReissuedFromId, true
}

// HasReissuedFromId returns a boolean if a field has been set.
func (o *VirtualCard) HasReissuedFromId() bool {
	if o != nil && !IsNil(o.ReissuedFromId) {
		return true
	}

	return false
}

// SetReissuedFromId gets a reference to the given string and assigns it to the ReissuedFromId field.
func (o *VirtualCard) SetReissuedFromId(v string) {
	o.ReissuedFromId = &v
}

// GetReissuedToId returns the ReissuedToId field value if set, zero value otherwise.
func (o *VirtualCard) GetReissuedToId() string {
	if o == nil || IsNil(o.ReissuedToId) {
		var ret string
		return ret
	}
	return *o.ReissuedToId
}

// GetReissuedToIdOk returns a tuple with the ReissuedToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetReissuedToIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReissuedToId) {
		return nil, false
	}
	return o.ReissuedToId, true
}

// HasReissuedToId returns a boolean if a field has been set.
func (o *VirtualCard) HasReissuedToId() bool {
	if o != nil && !IsNil(o.ReissuedToId) {
		return true
	}

	return false
}

// SetReissuedToId gets a reference to the given string and assigns it to the ReissuedToId field.
func (o *VirtualCard) SetReissuedToId(v string) {
	o.ReissuedToId = &v
}

// GetTimestampPinSet returns the TimestampPinSet field value if set, zero value otherwise.
func (o *VirtualCard) GetTimestampPinSet() time.Time {
	if o == nil || IsNil(o.TimestampPinSet) {
		var ret time.Time
		return ret
	}
	return *o.TimestampPinSet
}

// GetTimestampPinSetOk returns a tuple with the TimestampPinSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetTimestampPinSetOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampPinSet) {
		return nil, false
	}
	return o.TimestampPinSet, true
}

// HasTimestampPinSet returns a boolean if a field has been set.
func (o *VirtualCard) HasTimestampPinSet() bool {
	if o != nil && !IsNil(o.TimestampPinSet) {
		return true
	}

	return false
}

// SetTimestampPinSet gets a reference to the given time.Time and assigns it to the TimestampPinSet field.
func (o *VirtualCard) SetTimestampPinSet(v time.Time) {
	o.TimestampPinSet = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualCard) GetType() CardType {
	if o == nil || IsNil(o.Type) {
		var ret CardType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetTypeOk() (*CardType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualCard) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CardType and assigns it to the Type field.
func (o *VirtualCard) SetType(v CardType) {
	o.Type = &v
}

// GetForm returns the Form field value
func (o *VirtualCard) GetForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *VirtualCard) GetFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *VirtualCard) SetForm(v string) {
	o.Form = v
}

func (o VirtualCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	if !IsNil(o.CardProductId) {
		toSerialize["card_product_id"] = o.CardProductId
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.EmbossName) {
		toSerialize["emboss_name"] = o.EmbossName
	}
	if !IsNil(o.ExpirationMonth) {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if !IsNil(o.ExpirationTime) {
		toSerialize["expiration_time"] = o.ExpirationTime
	}
	if !IsNil(o.ExpirationYear) {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsPinSet) {
		toSerialize["is_pin_set"] = o.IsPinSet
	}
	if !IsNil(o.LastFour) {
		toSerialize["last_four"] = o.LastFour
	}
	if !IsNil(o.LastModifiedTime) {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ReissueReason) {
		toSerialize["reissue_reason"] = o.ReissueReason
	}
	if !IsNil(o.ReissuedFromId) {
		toSerialize["reissued_from_id"] = o.ReissuedFromId
	}
	if !IsNil(o.ReissuedToId) {
		toSerialize["reissued_to_id"] = o.ReissuedToId
	}
	if !IsNil(o.TimestampPinSet) {
		toSerialize["timestamp_pin_set"] = o.TimestampPinSet
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["form"] = o.Form
	return toSerialize, nil
}

func (o *VirtualCard) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"form",
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

	varVirtualCard := _VirtualCard{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVirtualCard)

	if err != nil {
		return err
	}

	*o = VirtualCard(varVirtualCard)

	return err
}

type NullableVirtualCard struct {
	value *VirtualCard
	isSet bool
}

func (v NullableVirtualCard) Get() *VirtualCard {
	return v.value
}

func (v *NullableVirtualCard) Set(val *VirtualCard) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualCard) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualCard(val *VirtualCard) *NullableVirtualCard {
	return &NullableVirtualCard{value: val, isSet: true}
}

func (v NullableVirtualCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
