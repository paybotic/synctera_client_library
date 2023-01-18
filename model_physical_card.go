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

// PhysicalCard Details pertaining to issuing a new physical card
type PhysicalCard struct {
	// PHYSICAL or VIRTUAL.
	Form string `json:"form"`
	// The ID of the account to which the card will be linked
	AccountId *string `json:"account_id,omitempty"`
	// The card product to which the card is attached
	CardProductId *string `json:"card_product_id,omitempty"`
	// The timestamp representing when the card issuance request was made
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The ID of the customer to whom the card will be issued
	CustomerId *string `json:"customer_id,omitempty"`
	EmbossName *EmbossName `json:"emboss_name,omitempty"`
	ExpirationMonth *string `json:"expiration_month,omitempty"`
	// The timestamp representing when the card would expire at
	ExpirationTime *time.Time `json:"expiration_time,omitempty"`
	ExpirationYear *string `json:"expiration_year,omitempty"`
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
	// Indicates the type of card to be issued
	Type *string `json:"type,omitempty"`
	// The ID of the custom card image used for this card
	CardImageId *string `json:"card_image_id,omitempty"`
	Shipping *Shipping `json:"shipping,omitempty"`
}

// NewPhysicalCard instantiates a new PhysicalCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalCard(form string) *PhysicalCard {
	this := PhysicalCard{}
	this.Form = form
	return &this
}

// NewPhysicalCardWithDefaults instantiates a new PhysicalCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalCardWithDefaults() *PhysicalCard {
	this := PhysicalCard{}
	return &this
}

// GetForm returns the Form field value
func (o *PhysicalCard) GetForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *PhysicalCard) SetForm(v string) {
	o.Form = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PhysicalCard) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PhysicalCard) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *PhysicalCard) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCardProductId returns the CardProductId field value if set, zero value otherwise.
func (o *PhysicalCard) GetCardProductId() string {
	if o == nil || o.CardProductId == nil {
		var ret string
		return ret
	}
	return *o.CardProductId
}

// GetCardProductIdOk returns a tuple with the CardProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetCardProductIdOk() (*string, bool) {
	if o == nil || o.CardProductId == nil {
		return nil, false
	}
	return o.CardProductId, true
}

// HasCardProductId returns a boolean if a field has been set.
func (o *PhysicalCard) HasCardProductId() bool {
	if o != nil && o.CardProductId != nil {
		return true
	}

	return false
}

// SetCardProductId gets a reference to the given string and assigns it to the CardProductId field.
func (o *PhysicalCard) SetCardProductId(v string) {
	o.CardProductId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PhysicalCard) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PhysicalCard) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PhysicalCard) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *PhysicalCard) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PhysicalCard) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *PhysicalCard) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEmbossName returns the EmbossName field value if set, zero value otherwise.
func (o *PhysicalCard) GetEmbossName() EmbossName {
	if o == nil || o.EmbossName == nil {
		var ret EmbossName
		return ret
	}
	return *o.EmbossName
}

// GetEmbossNameOk returns a tuple with the EmbossName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetEmbossNameOk() (*EmbossName, bool) {
	if o == nil || o.EmbossName == nil {
		return nil, false
	}
	return o.EmbossName, true
}

// HasEmbossName returns a boolean if a field has been set.
func (o *PhysicalCard) HasEmbossName() bool {
	if o != nil && o.EmbossName != nil {
		return true
	}

	return false
}

// SetEmbossName gets a reference to the given EmbossName and assigns it to the EmbossName field.
func (o *PhysicalCard) SetEmbossName(v EmbossName) {
	o.EmbossName = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *PhysicalCard) GetExpirationMonth() string {
	if o == nil || o.ExpirationMonth == nil {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetExpirationMonthOk() (*string, bool) {
	if o == nil || o.ExpirationMonth == nil {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *PhysicalCard) HasExpirationMonth() bool {
	if o != nil && o.ExpirationMonth != nil {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *PhysicalCard) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *PhysicalCard) GetExpirationTime() time.Time {
	if o == nil || o.ExpirationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationTime == nil {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *PhysicalCard) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime != nil {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *PhysicalCard) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *PhysicalCard) GetExpirationYear() string {
	if o == nil || o.ExpirationYear == nil {
		var ret string
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetExpirationYearOk() (*string, bool) {
	if o == nil || o.ExpirationYear == nil {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *PhysicalCard) HasExpirationYear() bool {
	if o != nil && o.ExpirationYear != nil {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given string and assigns it to the ExpirationYear field.
func (o *PhysicalCard) SetExpirationYear(v string) {
	o.ExpirationYear = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PhysicalCard) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PhysicalCard) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PhysicalCard) SetId(v string) {
	o.Id = &v
}

// GetIsPinSet returns the IsPinSet field value if set, zero value otherwise.
func (o *PhysicalCard) GetIsPinSet() bool {
	if o == nil || o.IsPinSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPinSet
}

// GetIsPinSetOk returns a tuple with the IsPinSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetIsPinSetOk() (*bool, bool) {
	if o == nil || o.IsPinSet == nil {
		return nil, false
	}
	return o.IsPinSet, true
}

// HasIsPinSet returns a boolean if a field has been set.
func (o *PhysicalCard) HasIsPinSet() bool {
	if o != nil && o.IsPinSet != nil {
		return true
	}

	return false
}

// SetIsPinSet gets a reference to the given bool and assigns it to the IsPinSet field.
func (o *PhysicalCard) SetIsPinSet(v bool) {
	o.IsPinSet = &v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *PhysicalCard) GetLastFour() string {
	if o == nil || o.LastFour == nil {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetLastFourOk() (*string, bool) {
	if o == nil || o.LastFour == nil {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *PhysicalCard) HasLastFour() bool {
	if o != nil && o.LastFour != nil {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *PhysicalCard) SetLastFour(v string) {
	o.LastFour = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *PhysicalCard) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *PhysicalCard) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *PhysicalCard) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PhysicalCard) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PhysicalCard) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *PhysicalCard) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetReissueReason returns the ReissueReason field value if set, zero value otherwise.
func (o *PhysicalCard) GetReissueReason() string {
	if o == nil || o.ReissueReason == nil {
		var ret string
		return ret
	}
	return *o.ReissueReason
}

// GetReissueReasonOk returns a tuple with the ReissueReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetReissueReasonOk() (*string, bool) {
	if o == nil || o.ReissueReason == nil {
		return nil, false
	}
	return o.ReissueReason, true
}

// HasReissueReason returns a boolean if a field has been set.
func (o *PhysicalCard) HasReissueReason() bool {
	if o != nil && o.ReissueReason != nil {
		return true
	}

	return false
}

// SetReissueReason gets a reference to the given string and assigns it to the ReissueReason field.
func (o *PhysicalCard) SetReissueReason(v string) {
	o.ReissueReason = &v
}

// GetReissuedFromId returns the ReissuedFromId field value if set, zero value otherwise.
func (o *PhysicalCard) GetReissuedFromId() string {
	if o == nil || o.ReissuedFromId == nil {
		var ret string
		return ret
	}
	return *o.ReissuedFromId
}

// GetReissuedFromIdOk returns a tuple with the ReissuedFromId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetReissuedFromIdOk() (*string, bool) {
	if o == nil || o.ReissuedFromId == nil {
		return nil, false
	}
	return o.ReissuedFromId, true
}

// HasReissuedFromId returns a boolean if a field has been set.
func (o *PhysicalCard) HasReissuedFromId() bool {
	if o != nil && o.ReissuedFromId != nil {
		return true
	}

	return false
}

// SetReissuedFromId gets a reference to the given string and assigns it to the ReissuedFromId field.
func (o *PhysicalCard) SetReissuedFromId(v string) {
	o.ReissuedFromId = &v
}

// GetReissuedToId returns the ReissuedToId field value if set, zero value otherwise.
func (o *PhysicalCard) GetReissuedToId() string {
	if o == nil || o.ReissuedToId == nil {
		var ret string
		return ret
	}
	return *o.ReissuedToId
}

// GetReissuedToIdOk returns a tuple with the ReissuedToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetReissuedToIdOk() (*string, bool) {
	if o == nil || o.ReissuedToId == nil {
		return nil, false
	}
	return o.ReissuedToId, true
}

// HasReissuedToId returns a boolean if a field has been set.
func (o *PhysicalCard) HasReissuedToId() bool {
	if o != nil && o.ReissuedToId != nil {
		return true
	}

	return false
}

// SetReissuedToId gets a reference to the given string and assigns it to the ReissuedToId field.
func (o *PhysicalCard) SetReissuedToId(v string) {
	o.ReissuedToId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PhysicalCard) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PhysicalCard) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PhysicalCard) SetType(v string) {
	o.Type = &v
}

// GetCardImageId returns the CardImageId field value if set, zero value otherwise.
func (o *PhysicalCard) GetCardImageId() string {
	if o == nil || o.CardImageId == nil {
		var ret string
		return ret
	}
	return *o.CardImageId
}

// GetCardImageIdOk returns a tuple with the CardImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetCardImageIdOk() (*string, bool) {
	if o == nil || o.CardImageId == nil {
		return nil, false
	}
	return o.CardImageId, true
}

// HasCardImageId returns a boolean if a field has been set.
func (o *PhysicalCard) HasCardImageId() bool {
	if o != nil && o.CardImageId != nil {
		return true
	}

	return false
}

// SetCardImageId gets a reference to the given string and assigns it to the CardImageId field.
func (o *PhysicalCard) SetCardImageId(v string) {
	o.CardImageId = &v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *PhysicalCard) GetShipping() Shipping {
	if o == nil || o.Shipping == nil {
		var ret Shipping
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCard) GetShippingOk() (*Shipping, bool) {
	if o == nil || o.Shipping == nil {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *PhysicalCard) HasShipping() bool {
	if o != nil && o.Shipping != nil {
		return true
	}

	return false
}

// SetShipping gets a reference to the given Shipping and assigns it to the Shipping field.
func (o *PhysicalCard) SetShipping(v Shipping) {
	o.Shipping = &v
}

func (o PhysicalCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["form"] = o.Form
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.CardProductId != nil {
		toSerialize["card_product_id"] = o.CardProductId
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.EmbossName != nil {
		toSerialize["emboss_name"] = o.EmbossName
	}
	if o.ExpirationMonth != nil {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if o.ExpirationTime != nil {
		toSerialize["expiration_time"] = o.ExpirationTime
	}
	if o.ExpirationYear != nil {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsPinSet != nil {
		toSerialize["is_pin_set"] = o.IsPinSet
	}
	if o.LastFour != nil {
		toSerialize["last_four"] = o.LastFour
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ReissueReason != nil {
		toSerialize["reissue_reason"] = o.ReissueReason
	}
	if o.ReissuedFromId != nil {
		toSerialize["reissued_from_id"] = o.ReissuedFromId
	}
	if o.ReissuedToId != nil {
		toSerialize["reissued_to_id"] = o.ReissuedToId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.CardImageId != nil {
		toSerialize["card_image_id"] = o.CardImageId
	}
	if o.Shipping != nil {
		toSerialize["shipping"] = o.Shipping
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalCard struct {
	value *PhysicalCard
	isSet bool
}

func (v NullablePhysicalCard) Get() *PhysicalCard {
	return v.value
}

func (v *NullablePhysicalCard) Set(val *PhysicalCard) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalCard) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalCard(val *PhysicalCard) *NullablePhysicalCard {
	return &NullablePhysicalCard{value: val, isSet: true}
}

func (v NullablePhysicalCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


