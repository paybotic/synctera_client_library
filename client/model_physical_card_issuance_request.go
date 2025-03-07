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
	"time"
)

// checks if the PhysicalCardIssuanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhysicalCardIssuanceRequest{}

// PhysicalCardIssuanceRequest struct for PhysicalCardIssuanceRequest
type PhysicalCardIssuanceRequest struct {
	// The ID of the account to which the card will be linked
	AccountId string `json:"account_id"`
	// The business ID associated with this card. If no customer_id is supplied, a card can still be issued to a business, but cannot be activated or used until a customer is assigned via the PATCH /cards/{card_id} endpoint.
	BusinessId *string `json:"business_id,omitempty"`
	// The card product to which the card is attached
	CardProductId string `json:"card_product_id"`
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
	Type            CardType   `json:"type"`
	// PHYSICAL or VIRTUAL.
	Form string `json:"form"`
	// The ID of the custom card image used for this card
	CardImageId          *string   `json:"card_image_id,omitempty"`
	Shipping             *Shipping `json:"shipping,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PhysicalCardIssuanceRequest PhysicalCardIssuanceRequest

// NewPhysicalCardIssuanceRequest instantiates a new PhysicalCardIssuanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalCardIssuanceRequest(accountId string, cardProductId string, type_ CardType, form string) *PhysicalCardIssuanceRequest {
	this := PhysicalCardIssuanceRequest{}
	this.AccountId = accountId
	this.CardProductId = cardProductId
	this.Type = type_
	this.Form = form
	return &this
}

// NewPhysicalCardIssuanceRequestWithDefaults instantiates a new PhysicalCardIssuanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalCardIssuanceRequestWithDefaults() *PhysicalCardIssuanceRequest {
	this := PhysicalCardIssuanceRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *PhysicalCardIssuanceRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *PhysicalCardIssuanceRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *PhysicalCardIssuanceRequest) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCardProductId returns the CardProductId field value
func (o *PhysicalCardIssuanceRequest) GetCardProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardProductId
}

// GetCardProductIdOk returns a tuple with the CardProductId field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetCardProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProductId, true
}

// SetCardProductId sets field value
func (o *PhysicalCardIssuanceRequest) SetCardProductId(v string) {
	o.CardProductId = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PhysicalCardIssuanceRequest) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *PhysicalCardIssuanceRequest) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEmbossName returns the EmbossName field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetEmbossName() EmbossName {
	if o == nil || IsNil(o.EmbossName) {
		var ret EmbossName
		return ret
	}
	return *o.EmbossName
}

// GetEmbossNameOk returns a tuple with the EmbossName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetEmbossNameOk() (*EmbossName, bool) {
	if o == nil || IsNil(o.EmbossName) {
		return nil, false
	}
	return o.EmbossName, true
}

// HasEmbossName returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasEmbossName() bool {
	if o != nil && !IsNil(o.EmbossName) {
		return true
	}

	return false
}

// SetEmbossName gets a reference to the given EmbossName and assigns it to the EmbossName field.
func (o *PhysicalCardIssuanceRequest) SetEmbossName(v EmbossName) {
	o.EmbossName = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetExpirationMonth() string {
	if o == nil || IsNil(o.ExpirationMonth) {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetExpirationMonthOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationMonth) {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasExpirationMonth() bool {
	if o != nil && !IsNil(o.ExpirationMonth) {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *PhysicalCardIssuanceRequest) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *PhysicalCardIssuanceRequest) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetExpirationYear() string {
	if o == nil || IsNil(o.ExpirationYear) {
		var ret string
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetExpirationYearOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationYear) {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasExpirationYear() bool {
	if o != nil && !IsNil(o.ExpirationYear) {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given string and assigns it to the ExpirationYear field.
func (o *PhysicalCardIssuanceRequest) SetExpirationYear(v string) {
	o.ExpirationYear = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PhysicalCardIssuanceRequest) SetId(v string) {
	o.Id = &v
}

// GetIsPinSet returns the IsPinSet field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetIsPinSet() bool {
	if o == nil || IsNil(o.IsPinSet) {
		var ret bool
		return ret
	}
	return *o.IsPinSet
}

// GetIsPinSetOk returns a tuple with the IsPinSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetIsPinSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPinSet) {
		return nil, false
	}
	return o.IsPinSet, true
}

// HasIsPinSet returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasIsPinSet() bool {
	if o != nil && !IsNil(o.IsPinSet) {
		return true
	}

	return false
}

// SetIsPinSet gets a reference to the given bool and assigns it to the IsPinSet field.
func (o *PhysicalCardIssuanceRequest) SetIsPinSet(v bool) {
	o.IsPinSet = &v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetLastFour() string {
	if o == nil || IsNil(o.LastFour) {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetLastFourOk() (*string, bool) {
	if o == nil || IsNil(o.LastFour) {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasLastFour() bool {
	if o != nil && !IsNil(o.LastFour) {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *PhysicalCardIssuanceRequest) SetLastFour(v string) {
	o.LastFour = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *PhysicalCardIssuanceRequest) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *PhysicalCardIssuanceRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetReissueReason returns the ReissueReason field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetReissueReason() string {
	if o == nil || IsNil(o.ReissueReason) {
		var ret string
		return ret
	}
	return *o.ReissueReason
}

// GetReissueReasonOk returns a tuple with the ReissueReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetReissueReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ReissueReason) {
		return nil, false
	}
	return o.ReissueReason, true
}

// HasReissueReason returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasReissueReason() bool {
	if o != nil && !IsNil(o.ReissueReason) {
		return true
	}

	return false
}

// SetReissueReason gets a reference to the given string and assigns it to the ReissueReason field.
func (o *PhysicalCardIssuanceRequest) SetReissueReason(v string) {
	o.ReissueReason = &v
}

// GetReissuedFromId returns the ReissuedFromId field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetReissuedFromId() string {
	if o == nil || IsNil(o.ReissuedFromId) {
		var ret string
		return ret
	}
	return *o.ReissuedFromId
}

// GetReissuedFromIdOk returns a tuple with the ReissuedFromId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetReissuedFromIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReissuedFromId) {
		return nil, false
	}
	return o.ReissuedFromId, true
}

// HasReissuedFromId returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasReissuedFromId() bool {
	if o != nil && !IsNil(o.ReissuedFromId) {
		return true
	}

	return false
}

// SetReissuedFromId gets a reference to the given string and assigns it to the ReissuedFromId field.
func (o *PhysicalCardIssuanceRequest) SetReissuedFromId(v string) {
	o.ReissuedFromId = &v
}

// GetReissuedToId returns the ReissuedToId field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetReissuedToId() string {
	if o == nil || IsNil(o.ReissuedToId) {
		var ret string
		return ret
	}
	return *o.ReissuedToId
}

// GetReissuedToIdOk returns a tuple with the ReissuedToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetReissuedToIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReissuedToId) {
		return nil, false
	}
	return o.ReissuedToId, true
}

// HasReissuedToId returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasReissuedToId() bool {
	if o != nil && !IsNil(o.ReissuedToId) {
		return true
	}

	return false
}

// SetReissuedToId gets a reference to the given string and assigns it to the ReissuedToId field.
func (o *PhysicalCardIssuanceRequest) SetReissuedToId(v string) {
	o.ReissuedToId = &v
}

// GetTimestampPinSet returns the TimestampPinSet field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetTimestampPinSet() time.Time {
	if o == nil || IsNil(o.TimestampPinSet) {
		var ret time.Time
		return ret
	}
	return *o.TimestampPinSet
}

// GetTimestampPinSetOk returns a tuple with the TimestampPinSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetTimestampPinSetOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampPinSet) {
		return nil, false
	}
	return o.TimestampPinSet, true
}

// HasTimestampPinSet returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasTimestampPinSet() bool {
	if o != nil && !IsNil(o.TimestampPinSet) {
		return true
	}

	return false
}

// SetTimestampPinSet gets a reference to the given time.Time and assigns it to the TimestampPinSet field.
func (o *PhysicalCardIssuanceRequest) SetTimestampPinSet(v time.Time) {
	o.TimestampPinSet = &v
}

// GetType returns the Type field value
func (o *PhysicalCardIssuanceRequest) GetType() CardType {
	if o == nil {
		var ret CardType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetTypeOk() (*CardType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PhysicalCardIssuanceRequest) SetType(v CardType) {
	o.Type = v
}

// GetForm returns the Form field value
func (o *PhysicalCardIssuanceRequest) GetForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *PhysicalCardIssuanceRequest) SetForm(v string) {
	o.Form = v
}

// GetCardImageId returns the CardImageId field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetCardImageId() string {
	if o == nil || IsNil(o.CardImageId) {
		var ret string
		return ret
	}
	return *o.CardImageId
}

// GetCardImageIdOk returns a tuple with the CardImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetCardImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.CardImageId) {
		return nil, false
	}
	return o.CardImageId, true
}

// HasCardImageId returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasCardImageId() bool {
	if o != nil && !IsNil(o.CardImageId) {
		return true
	}

	return false
}

// SetCardImageId gets a reference to the given string and assigns it to the CardImageId field.
func (o *PhysicalCardIssuanceRequest) SetCardImageId(v string) {
	o.CardImageId = &v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *PhysicalCardIssuanceRequest) GetShipping() Shipping {
	if o == nil || IsNil(o.Shipping) {
		var ret Shipping
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardIssuanceRequest) GetShippingOk() (*Shipping, bool) {
	if o == nil || IsNil(o.Shipping) {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *PhysicalCardIssuanceRequest) HasShipping() bool {
	if o != nil && !IsNil(o.Shipping) {
		return true
	}

	return false
}

// SetShipping gets a reference to the given Shipping and assigns it to the Shipping field.
func (o *PhysicalCardIssuanceRequest) SetShipping(v Shipping) {
	o.Shipping = &v
}

func (o PhysicalCardIssuanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhysicalCardIssuanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	toSerialize["card_product_id"] = o.CardProductId
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
	toSerialize["type"] = o.Type
	toSerialize["form"] = o.Form
	if !IsNil(o.CardImageId) {
		toSerialize["card_image_id"] = o.CardImageId
	}
	if !IsNil(o.Shipping) {
		toSerialize["shipping"] = o.Shipping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PhysicalCardIssuanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"card_product_id",
		"type",
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

	varPhysicalCardIssuanceRequest := _PhysicalCardIssuanceRequest{}

	err = json.Unmarshal(data, &varPhysicalCardIssuanceRequest)

	if err != nil {
		return err
	}

	*o = PhysicalCardIssuanceRequest(varPhysicalCardIssuanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "business_id")
		delete(additionalProperties, "card_product_id")
		delete(additionalProperties, "creation_time")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "emboss_name")
		delete(additionalProperties, "expiration_month")
		delete(additionalProperties, "expiration_time")
		delete(additionalProperties, "expiration_year")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_pin_set")
		delete(additionalProperties, "last_four")
		delete(additionalProperties, "last_modified_time")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "reissue_reason")
		delete(additionalProperties, "reissued_from_id")
		delete(additionalProperties, "reissued_to_id")
		delete(additionalProperties, "timestamp_pin_set")
		delete(additionalProperties, "type")
		delete(additionalProperties, "form")
		delete(additionalProperties, "card_image_id")
		delete(additionalProperties, "shipping")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePhysicalCardIssuanceRequest struct {
	value *PhysicalCardIssuanceRequest
	isSet bool
}

func (v NullablePhysicalCardIssuanceRequest) Get() *PhysicalCardIssuanceRequest {
	return v.value
}

func (v *NullablePhysicalCardIssuanceRequest) Set(val *PhysicalCardIssuanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalCardIssuanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalCardIssuanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalCardIssuanceRequest(val *PhysicalCardIssuanceRequest) *NullablePhysicalCardIssuanceRequest {
	return &NullablePhysicalCardIssuanceRequest{value: val, isSet: true}
}

func (v NullablePhysicalCardIssuanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalCardIssuanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
