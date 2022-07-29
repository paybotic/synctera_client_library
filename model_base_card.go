/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.17.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type BaseCard struct {
	// PHYSICAL or VIRTUAL.
	Form string `json:"form"`
	// The ID of the account to which the card will be linked
	AccountId string `json:"account_id,omitempty"`
	// The bin number
	Bin string `json:"bin,omitempty"`
	CardBrand *CardBrand `json:"card_brand,omitempty"`
	// The card product to which the card is attached
	CardProductId string `json:"card_product_id,omitempty"`
	// The timestamp representing when the card issuance request was made
	CreationTime time.Time `json:"creation_time,omitempty"`
	// The ID of the customer to whom the card will be issued
	CustomerId string `json:"customer_id,omitempty"`
	EmbossName *EmbossName `json:"emboss_name,omitempty"`
	ExpirationMonth string `json:"expiration_month,omitempty"`
	// The timestamp representing when the card would expire at
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	ExpirationYear string `json:"expiration_year,omitempty"`
	// Card ID
	Id string `json:"id,omitempty"`
	// The last 4 digits of the card PAN
	LastFour string `json:"last_four,omitempty"`
	// The timestamp representing when the card was last modified at
	LastModifiedTime time.Time `json:"last_modified_time,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	// This is the reason the card needs to be reissued, if any. The reason determines several behaviours:   - whether or not the new card will use the same PAN as the original card   - the old card will be terminated and if so, when it will be terminated  Reason                 | Same PAN | Terminate Old Card ---------------------- | -------- | ------------------ EXPIRATION             | yes      | on activation LOST                   | no       | immediately STOLEN                 | no       | immediately DAMAGED                | yes      | on activation VIRTUAL_TO_PHYSICAL(*) | yes      | on activation PRODUCT_CHANGE         | yes      | on activation NAME_CHANGE(**)        | yes      | on activation APPEARANCE             | yes      | on activation  (*) VIRTUAL_TO_PHYSICAL is deprecated. Please use PRODUCT_CHANGE whenever reissuing from one card product to another, including from a virtual product to a physical product.  (**) NAME_CHANGE is deprecated. Please use APPEARANCE whenever reissuing in order to change the appearance of a card, such as the printed name or custom image.  For all reasons, the new card will use the same PIN as the original card and digital wallet tokens will reassigned to the new card 
	ReissueReason string `json:"reissue_reason,omitempty"`
	// When reissuing a card, specify the card to be replaced here. When getting a card's details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced. If this field is set, then reissue_reason must also be set. 
	ReissuedFromId string `json:"reissued_from_id,omitempty"`
	// If this card was reissued, this ID refers to the card that replaced it.
	ReissuedToId string `json:"reissued_to_id,omitempty"`
	// Indicates the type of card to be issued
	Type_ string `json:"type,omitempty"`
}
