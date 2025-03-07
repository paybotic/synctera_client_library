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

// checks if the PhysicalCardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhysicalCardResponse{}

// PhysicalCardResponse struct for PhysicalCardResponse
type PhysicalCardResponse struct {
	// The ID of the account to which the card will be linked
	AccountId string `json:"account_id"`
	// The business ID associated with this card. If no customer_id is supplied, a card can still be issued to a business, but cannot be activated or used until a customer is assigned via the PATCH /cards/{card_id} endpoint.
	BusinessId *string `json:"business_id,omitempty"`
	// The card product to which the card is attached
	CardProductId string `json:"card_product_id"`
	// The timestamp representing when the card issuance request was made
	CreationTime time.Time `json:"creation_time"`
	// The ID of the customer to whom the card will be issued. If a business_id is passed, but a customer_id not passed at the time of card creation the card cannot be activated or used for spend until it's assigned to a human customer via the PATCH /cards/{card_id} endpoint. If no business_id is passed, a customer_id is required.
	CustomerId      *string    `json:"customer_id,omitempty"`
	EmbossName      EmbossName `json:"emboss_name"`
	ExpirationMonth *string    `json:"expiration_month,omitempty"`
	// The timestamp representing when the card would expire at
	ExpirationTime *time.Time `json:"expiration_time,omitempty"`
	ExpirationYear *string    `json:"expiration_year,omitempty"`
	// Card ID
	Id string `json:"id"`
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
	CardImageId *string    `json:"card_image_id,omitempty"`
	Shipping    Shipping   `json:"shipping"`
	CardStatus  CardStatus `json:"card_status"`
	// Additional details about the reason for the status change
	Memo                  *string                   `json:"memo,omitempty"`
	PendingReasons        *CardStatusPendingReasons `json:"pending_reasons,omitempty"`
	StatusReason          CardStatusReasonCode      `json:"status_reason"`
	CardFulfillmentStatus CardFulfillmentStatus     `json:"card_fulfillment_status"`
	FulfillmentDetails    *FulfillmentDetails       `json:"fulfillment_details,omitempty"`
	// This contains all shipping details as provided by the card fulfillment provider, including the tracking number. This field is deprecated. Instead, please use the fulfillment_details object, which includes a field for just the tracking number.
	// Deprecated
	TrackingNumber *string `json:"tracking_number,omitempty"`
	// The bin number
	Bin                *string            `json:"bin,omitempty"`
	CardBrand          CardBrand          `json:"card_brand"`
	PhysicalCardFormat PhysicalCardFormat `json:"physical_card_format"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant               string `json:"tenant"`
	AdditionalProperties map[string]interface{}
}

type _PhysicalCardResponse PhysicalCardResponse

// NewPhysicalCardResponse instantiates a new PhysicalCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalCardResponse(accountId string, cardProductId string, creationTime time.Time, embossName EmbossName, id string, type_ CardType, form string, shipping Shipping, cardStatus CardStatus, statusReason CardStatusReasonCode, cardFulfillmentStatus CardFulfillmentStatus, cardBrand CardBrand, physicalCardFormat PhysicalCardFormat, tenant string) *PhysicalCardResponse {
	this := PhysicalCardResponse{}
	this.AccountId = accountId
	this.CardProductId = cardProductId
	this.CreationTime = creationTime
	this.EmbossName = embossName
	this.Id = id
	this.Type = type_
	this.Form = form
	this.Shipping = shipping
	this.CardStatus = cardStatus
	this.StatusReason = statusReason
	this.CardFulfillmentStatus = cardFulfillmentStatus
	this.CardBrand = cardBrand
	this.PhysicalCardFormat = physicalCardFormat
	this.Tenant = tenant
	return &this
}

// NewPhysicalCardResponseWithDefaults instantiates a new PhysicalCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalCardResponseWithDefaults() *PhysicalCardResponse {
	this := PhysicalCardResponse{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *PhysicalCardResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *PhysicalCardResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *PhysicalCardResponse) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCardProductId returns the CardProductId field value
func (o *PhysicalCardResponse) GetCardProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardProductId
}

// GetCardProductIdOk returns a tuple with the CardProductId field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetCardProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProductId, true
}

// SetCardProductId sets field value
func (o *PhysicalCardResponse) SetCardProductId(v string) {
	o.CardProductId = v
}

// GetCreationTime returns the CreationTime field value
func (o *PhysicalCardResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *PhysicalCardResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *PhysicalCardResponse) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEmbossName returns the EmbossName field value
func (o *PhysicalCardResponse) GetEmbossName() EmbossName {
	if o == nil {
		var ret EmbossName
		return ret
	}

	return o.EmbossName
}

// GetEmbossNameOk returns a tuple with the EmbossName field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetEmbossNameOk() (*EmbossName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmbossName, true
}

// SetEmbossName sets field value
func (o *PhysicalCardResponse) SetEmbossName(v EmbossName) {
	o.EmbossName = v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetExpirationMonth() string {
	if o == nil || IsNil(o.ExpirationMonth) {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetExpirationMonthOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationMonth) {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasExpirationMonth() bool {
	if o != nil && !IsNil(o.ExpirationMonth) {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *PhysicalCardResponse) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *PhysicalCardResponse) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetExpirationYear() string {
	if o == nil || IsNil(o.ExpirationYear) {
		var ret string
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetExpirationYearOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationYear) {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasExpirationYear() bool {
	if o != nil && !IsNil(o.ExpirationYear) {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given string and assigns it to the ExpirationYear field.
func (o *PhysicalCardResponse) SetExpirationYear(v string) {
	o.ExpirationYear = &v
}

// GetId returns the Id field value
func (o *PhysicalCardResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PhysicalCardResponse) SetId(v string) {
	o.Id = v
}

// GetIsPinSet returns the IsPinSet field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetIsPinSet() bool {
	if o == nil || IsNil(o.IsPinSet) {
		var ret bool
		return ret
	}
	return *o.IsPinSet
}

// GetIsPinSetOk returns a tuple with the IsPinSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetIsPinSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPinSet) {
		return nil, false
	}
	return o.IsPinSet, true
}

// HasIsPinSet returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasIsPinSet() bool {
	if o != nil && !IsNil(o.IsPinSet) {
		return true
	}

	return false
}

// SetIsPinSet gets a reference to the given bool and assigns it to the IsPinSet field.
func (o *PhysicalCardResponse) SetIsPinSet(v bool) {
	o.IsPinSet = &v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetLastFour() string {
	if o == nil || IsNil(o.LastFour) {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetLastFourOk() (*string, bool) {
	if o == nil || IsNil(o.LastFour) {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasLastFour() bool {
	if o != nil && !IsNil(o.LastFour) {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *PhysicalCardResponse) SetLastFour(v string) {
	o.LastFour = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *PhysicalCardResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *PhysicalCardResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetReissueReason returns the ReissueReason field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetReissueReason() string {
	if o == nil || IsNil(o.ReissueReason) {
		var ret string
		return ret
	}
	return *o.ReissueReason
}

// GetReissueReasonOk returns a tuple with the ReissueReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetReissueReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ReissueReason) {
		return nil, false
	}
	return o.ReissueReason, true
}

// HasReissueReason returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasReissueReason() bool {
	if o != nil && !IsNil(o.ReissueReason) {
		return true
	}

	return false
}

// SetReissueReason gets a reference to the given string and assigns it to the ReissueReason field.
func (o *PhysicalCardResponse) SetReissueReason(v string) {
	o.ReissueReason = &v
}

// GetReissuedFromId returns the ReissuedFromId field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetReissuedFromId() string {
	if o == nil || IsNil(o.ReissuedFromId) {
		var ret string
		return ret
	}
	return *o.ReissuedFromId
}

// GetReissuedFromIdOk returns a tuple with the ReissuedFromId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetReissuedFromIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReissuedFromId) {
		return nil, false
	}
	return o.ReissuedFromId, true
}

// HasReissuedFromId returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasReissuedFromId() bool {
	if o != nil && !IsNil(o.ReissuedFromId) {
		return true
	}

	return false
}

// SetReissuedFromId gets a reference to the given string and assigns it to the ReissuedFromId field.
func (o *PhysicalCardResponse) SetReissuedFromId(v string) {
	o.ReissuedFromId = &v
}

// GetReissuedToId returns the ReissuedToId field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetReissuedToId() string {
	if o == nil || IsNil(o.ReissuedToId) {
		var ret string
		return ret
	}
	return *o.ReissuedToId
}

// GetReissuedToIdOk returns a tuple with the ReissuedToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetReissuedToIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReissuedToId) {
		return nil, false
	}
	return o.ReissuedToId, true
}

// HasReissuedToId returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasReissuedToId() bool {
	if o != nil && !IsNil(o.ReissuedToId) {
		return true
	}

	return false
}

// SetReissuedToId gets a reference to the given string and assigns it to the ReissuedToId field.
func (o *PhysicalCardResponse) SetReissuedToId(v string) {
	o.ReissuedToId = &v
}

// GetTimestampPinSet returns the TimestampPinSet field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetTimestampPinSet() time.Time {
	if o == nil || IsNil(o.TimestampPinSet) {
		var ret time.Time
		return ret
	}
	return *o.TimestampPinSet
}

// GetTimestampPinSetOk returns a tuple with the TimestampPinSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetTimestampPinSetOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampPinSet) {
		return nil, false
	}
	return o.TimestampPinSet, true
}

// HasTimestampPinSet returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasTimestampPinSet() bool {
	if o != nil && !IsNil(o.TimestampPinSet) {
		return true
	}

	return false
}

// SetTimestampPinSet gets a reference to the given time.Time and assigns it to the TimestampPinSet field.
func (o *PhysicalCardResponse) SetTimestampPinSet(v time.Time) {
	o.TimestampPinSet = &v
}

// GetType returns the Type field value
func (o *PhysicalCardResponse) GetType() CardType {
	if o == nil {
		var ret CardType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetTypeOk() (*CardType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PhysicalCardResponse) SetType(v CardType) {
	o.Type = v
}

// GetForm returns the Form field value
func (o *PhysicalCardResponse) GetForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *PhysicalCardResponse) SetForm(v string) {
	o.Form = v
}

// GetCardImageId returns the CardImageId field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetCardImageId() string {
	if o == nil || IsNil(o.CardImageId) {
		var ret string
		return ret
	}
	return *o.CardImageId
}

// GetCardImageIdOk returns a tuple with the CardImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetCardImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.CardImageId) {
		return nil, false
	}
	return o.CardImageId, true
}

// HasCardImageId returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasCardImageId() bool {
	if o != nil && !IsNil(o.CardImageId) {
		return true
	}

	return false
}

// SetCardImageId gets a reference to the given string and assigns it to the CardImageId field.
func (o *PhysicalCardResponse) SetCardImageId(v string) {
	o.CardImageId = &v
}

// GetShipping returns the Shipping field value
func (o *PhysicalCardResponse) GetShipping() Shipping {
	if o == nil {
		var ret Shipping
		return ret
	}

	return o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetShippingOk() (*Shipping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shipping, true
}

// SetShipping sets field value
func (o *PhysicalCardResponse) SetShipping(v Shipping) {
	o.Shipping = v
}

// GetCardStatus returns the CardStatus field value
func (o *PhysicalCardResponse) GetCardStatus() CardStatus {
	if o == nil {
		var ret CardStatus
		return ret
	}

	return o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetCardStatusOk() (*CardStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardStatus, true
}

// SetCardStatus sets field value
func (o *PhysicalCardResponse) SetCardStatus(v CardStatus) {
	o.CardStatus = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *PhysicalCardResponse) SetMemo(v string) {
	o.Memo = &v
}

// GetPendingReasons returns the PendingReasons field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetPendingReasons() CardStatusPendingReasons {
	if o == nil || IsNil(o.PendingReasons) {
		var ret CardStatusPendingReasons
		return ret
	}
	return *o.PendingReasons
}

// GetPendingReasonsOk returns a tuple with the PendingReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetPendingReasonsOk() (*CardStatusPendingReasons, bool) {
	if o == nil || IsNil(o.PendingReasons) {
		return nil, false
	}
	return o.PendingReasons, true
}

// HasPendingReasons returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasPendingReasons() bool {
	if o != nil && !IsNil(o.PendingReasons) {
		return true
	}

	return false
}

// SetPendingReasons gets a reference to the given CardStatusPendingReasons and assigns it to the PendingReasons field.
func (o *PhysicalCardResponse) SetPendingReasons(v CardStatusPendingReasons) {
	o.PendingReasons = &v
}

// GetStatusReason returns the StatusReason field value
func (o *PhysicalCardResponse) GetStatusReason() CardStatusReasonCode {
	if o == nil {
		var ret CardStatusReasonCode
		return ret
	}

	return o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetStatusReasonOk() (*CardStatusReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusReason, true
}

// SetStatusReason sets field value
func (o *PhysicalCardResponse) SetStatusReason(v CardStatusReasonCode) {
	o.StatusReason = v
}

// GetCardFulfillmentStatus returns the CardFulfillmentStatus field value
func (o *PhysicalCardResponse) GetCardFulfillmentStatus() CardFulfillmentStatus {
	if o == nil {
		var ret CardFulfillmentStatus
		return ret
	}

	return o.CardFulfillmentStatus
}

// GetCardFulfillmentStatusOk returns a tuple with the CardFulfillmentStatus field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetCardFulfillmentStatusOk() (*CardFulfillmentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardFulfillmentStatus, true
}

// SetCardFulfillmentStatus sets field value
func (o *PhysicalCardResponse) SetCardFulfillmentStatus(v CardFulfillmentStatus) {
	o.CardFulfillmentStatus = v
}

// GetFulfillmentDetails returns the FulfillmentDetails field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetFulfillmentDetails() FulfillmentDetails {
	if o == nil || IsNil(o.FulfillmentDetails) {
		var ret FulfillmentDetails
		return ret
	}
	return *o.FulfillmentDetails
}

// GetFulfillmentDetailsOk returns a tuple with the FulfillmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetFulfillmentDetailsOk() (*FulfillmentDetails, bool) {
	if o == nil || IsNil(o.FulfillmentDetails) {
		return nil, false
	}
	return o.FulfillmentDetails, true
}

// HasFulfillmentDetails returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasFulfillmentDetails() bool {
	if o != nil && !IsNil(o.FulfillmentDetails) {
		return true
	}

	return false
}

// SetFulfillmentDetails gets a reference to the given FulfillmentDetails and assigns it to the FulfillmentDetails field.
func (o *PhysicalCardResponse) SetFulfillmentDetails(v FulfillmentDetails) {
	o.FulfillmentDetails = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
// Deprecated
func (o *PhysicalCardResponse) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhysicalCardResponse) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
// Deprecated
func (o *PhysicalCardResponse) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *PhysicalCardResponse) GetBin() string {
	if o == nil || IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetBinOk() (*string, bool) {
	if o == nil || IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *PhysicalCardResponse) HasBin() bool {
	if o != nil && !IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *PhysicalCardResponse) SetBin(v string) {
	o.Bin = &v
}

// GetCardBrand returns the CardBrand field value
func (o *PhysicalCardResponse) GetCardBrand() CardBrand {
	if o == nil {
		var ret CardBrand
		return ret
	}

	return o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetCardBrandOk() (*CardBrand, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardBrand, true
}

// SetCardBrand sets field value
func (o *PhysicalCardResponse) SetCardBrand(v CardBrand) {
	o.CardBrand = v
}

// GetPhysicalCardFormat returns the PhysicalCardFormat field value
func (o *PhysicalCardResponse) GetPhysicalCardFormat() PhysicalCardFormat {
	if o == nil {
		var ret PhysicalCardFormat
		return ret
	}

	return o.PhysicalCardFormat
}

// GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhysicalCardFormat, true
}

// SetPhysicalCardFormat sets field value
func (o *PhysicalCardResponse) SetPhysicalCardFormat(v PhysicalCardFormat) {
	o.PhysicalCardFormat = v
}

// GetTenant returns the Tenant field value
func (o *PhysicalCardResponse) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponse) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *PhysicalCardResponse) SetTenant(v string) {
	o.Tenant = v
}

func (o PhysicalCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhysicalCardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	toSerialize["card_product_id"] = o.CardProductId
	toSerialize["creation_time"] = o.CreationTime
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	toSerialize["emboss_name"] = o.EmbossName
	if !IsNil(o.ExpirationMonth) {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if !IsNil(o.ExpirationTime) {
		toSerialize["expiration_time"] = o.ExpirationTime
	}
	if !IsNil(o.ExpirationYear) {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	toSerialize["id"] = o.Id
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
	toSerialize["shipping"] = o.Shipping
	toSerialize["card_status"] = o.CardStatus
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.PendingReasons) {
		toSerialize["pending_reasons"] = o.PendingReasons
	}
	toSerialize["status_reason"] = o.StatusReason
	toSerialize["card_fulfillment_status"] = o.CardFulfillmentStatus
	if !IsNil(o.FulfillmentDetails) {
		toSerialize["fulfillment_details"] = o.FulfillmentDetails
	}
	if !IsNil(o.TrackingNumber) {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	if !IsNil(o.Bin) {
		toSerialize["bin"] = o.Bin
	}
	toSerialize["card_brand"] = o.CardBrand
	toSerialize["physical_card_format"] = o.PhysicalCardFormat
	toSerialize["tenant"] = o.Tenant

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PhysicalCardResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"card_product_id",
		"creation_time",
		"emboss_name",
		"id",
		"type",
		"form",
		"shipping",
		"card_status",
		"status_reason",
		"card_fulfillment_status",
		"card_brand",
		"physical_card_format",
		"tenant",
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

	varPhysicalCardResponse := _PhysicalCardResponse{}

	err = json.Unmarshal(data, &varPhysicalCardResponse)

	if err != nil {
		return err
	}

	*o = PhysicalCardResponse(varPhysicalCardResponse)

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
		delete(additionalProperties, "card_status")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "pending_reasons")
		delete(additionalProperties, "status_reason")
		delete(additionalProperties, "card_fulfillment_status")
		delete(additionalProperties, "fulfillment_details")
		delete(additionalProperties, "tracking_number")
		delete(additionalProperties, "bin")
		delete(additionalProperties, "card_brand")
		delete(additionalProperties, "physical_card_format")
		delete(additionalProperties, "tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePhysicalCardResponse struct {
	value *PhysicalCardResponse
	isSet bool
}

func (v NullablePhysicalCardResponse) Get() *PhysicalCardResponse {
	return v.value
}

func (v *NullablePhysicalCardResponse) Set(val *PhysicalCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalCardResponse(val *PhysicalCardResponse) *NullablePhysicalCardResponse {
	return &NullablePhysicalCardResponse{value: val, isSet: true}
}

func (v NullablePhysicalCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
