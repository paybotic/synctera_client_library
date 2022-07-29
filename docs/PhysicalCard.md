# PhysicalCard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account to which the card will be linked | [optional] [default to null]
**Bin** | **string** | The bin number | [optional] [default to null]
**CardBrand** | [***CardBrand**](card_brand.md) |  | [optional] [default to null]
**CardProductId** | **string** | The card product to which the card is attached | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | The timestamp representing when the card issuance request was made | [optional] [default to null]
**CustomerId** | **string** | The ID of the customer to whom the card will be issued | [optional] [default to null]
**EmbossName** | [***EmbossName**](emboss_name.md) |  | [optional] [default to null]
**ExpirationMonth** | **string** |  | [optional] [default to null]
**ExpirationTime** | [**time.Time**](time.Time.md) | The timestamp representing when the card would expire at | [optional] [default to null]
**ExpirationYear** | **string** |  | [optional] [default to null]
**Id** | **string** | Card ID | [optional] [default to null]
**LastFour** | **string** | The last 4 digits of the card PAN | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | The timestamp representing when the card was last modified at | [optional] [default to null]
**Metadata** | [***map[string]string**](map.md) |  | [optional] [default to null]
**ReissueReason** | **string** | This is the reason the card needs to be reissued, if any. The reason determines several behaviours:   - whether or not the new card will use the same PAN as the original card   - the old card will be terminated and if so, when it will be terminated  Reason                 | Same PAN | Terminate Old Card ---------------------- | -------- | ------------------ EXPIRATION             | yes      | on activation LOST                   | no       | immediately STOLEN                 | no       | immediately DAMAGED                | yes      | on activation VIRTUAL_TO_PHYSICAL(*) | yes      | on activation PRODUCT_CHANGE         | yes      | on activation NAME_CHANGE(**)        | yes      | on activation APPEARANCE             | yes      | on activation  (*) VIRTUAL_TO_PHYSICAL is deprecated. Please use PRODUCT_CHANGE whenever reissuing from one card product to another, including from a virtual product to a physical product.  (**) NAME_CHANGE is deprecated. Please use APPEARANCE whenever reissuing in order to change the appearance of a card, such as the printed name or custom image.  For all reasons, the new card will use the same PIN as the original card and digital wallet tokens will reassigned to the new card  | [optional] [default to null]
**ReissuedFromId** | **string** | When reissuing a card, specify the card to be replaced here. When getting a card&#x27;s details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced. If this field is set, then reissue_reason must also be set.  | [optional] [default to null]
**ReissuedToId** | **string** | If this card was reissued, this ID refers to the card that replaced it. | [optional] [default to null]
**Type_** | **string** | Indicates the type of card to be issued | [optional] [default to null]
**CardImageId** | **string** | The ID of the custom card image used for this card | [optional] [default to null]
**IsPinSet** | **bool** | indicates whether a pin has been set on the card | [optional] [default to false]
**Shipping** | [***Shipping**](shipping.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

