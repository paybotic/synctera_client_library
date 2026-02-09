# CreditCardResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | 
**AccountId** | **string** | The ID of the account to which the card will be linked | 
**BusinessId** | Pointer to **string** | The business ID associated with this card. If no customer_id is supplied, a card can still be issued to a business, but cannot be activated or used until a customer is assigned via the PATCH /cards/{card_id} endpoint. | [optional] 
**CardProductId** | **string** | The card product to which the card is attached | 
**CreationTime** | **time.Time** | The timestamp representing when the card issuance request was made | [readonly] 
**CustomerId** | Pointer to **string** | The ID of the customer to whom the card will be issued. If this is not populated with a valid customer_id the card cannot be activated or used for spend until it&#39;s assigned to a human customer via the PATCH /cards/{card_id} endpoint. If no business_id is passed, a customer_id is required. | [optional] 
**EmbossName** | [**EmbossName**](EmbossName.md) |  | 
**ExpirationMonth** | **string** |  | [readonly] 
**ExpirationTime** | Pointer to **time.Time** | The timestamp representing when the card would expire at | [optional] [readonly] 
**ExpirationYear** | **string** |  | [readonly] 
**Id** | **string** | Card ID | [readonly] 
**IsPinSet** | Pointer to **bool** | indicates whether a pin has been set on the card | [optional] [readonly] [default to false]
**LastFour** | **string** | The last 4 digits of the card PAN | [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The timestamp representing when the card was last modified at | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional data to include in the request structured as key-value pairs | [optional] 
**ReissueReason** | Pointer to **string** | This is the reason the card needs to be reissued, if any. The reason determines several behaviours:   - whether or not the new card will use the same PAN as the original card   - the old card will be terminated and if so, when it will be terminated  Reason                 | Same PAN | Terminate Old Card ---------------------- | -------- | ------------------ APPEARANCE             | yes      | on activation BANK_MIGRATION         | yes      | on activation DAMAGED                | yes      | on activation EXPIRATION             | yes      | on activation LOST                   | no       | immediately PRODUCT_CHANGE         | yes      | on activation PROGRAM_CHANGE         | yes      | on activation STOLEN                 | no       | immediately  For all reasons, the new card will use the same PIN as the original card and digital wallet tokens will reassigned to the new card  | [optional] 
**ReissuedFromId** | Pointer to **string** | When reissuing a card, specify the card to be replaced here. When getting a card&#39;s details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced. If this field is set, then reissue_reason must also be set.  | [optional] 
**ReissuedToId** | Pointer to **string** | If this card was reissued, this ID refers to the card that replaced it. | [optional] [readonly] 
**TimestampPinSet** | Pointer to **time.Time** | Time when the PIN was last set or changed. | [optional] [readonly] 
**BulkOrderConfigId** | Pointer to **string** | The ID of the bulk order config which should be used for shipping this card as part of a bulk order. Refer to Bulk Issuance for details on ordering cards in bulk.  | [optional] 
**CardImageId** | Pointer to **string** | The ID of the custom card image used for this card | [optional] 
**Shipping** | [**Shipping**](Shipping.md) |  | 
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**PendingReasons** | Pointer to [**CardStatusPendingReasons**](CardStatusPendingReasons.md) |  | [optional] 
**StatusReason** | [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | 
**CardFulfillmentStatus** | [**CardFulfillmentStatus**](CardFulfillmentStatus.md) |  | 
**FulfillmentDetails** | Pointer to [**FulfillmentDetails**](FulfillmentDetails.md) |  | [optional] 
**Bin** | Pointer to **string** | The bin number | [optional] 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 
**PhysicalCardFormat** | [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | 

## Methods

### NewCreditCardResponseDetails

`func NewCreditCardResponseDetails(form string, accountId string, cardProductId string, creationTime time.Time, embossName EmbossName, expirationMonth string, expirationYear string, id string, lastFour string, shipping Shipping, cardStatus CardStatus, statusReason CardStatusReasonCode, cardFulfillmentStatus CardFulfillmentStatus, cardBrand CardBrand, physicalCardFormat PhysicalCardFormat, ) *CreditCardResponseDetails`

NewCreditCardResponseDetails instantiates a new CreditCardResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardResponseDetailsWithDefaults

`func NewCreditCardResponseDetailsWithDefaults() *CreditCardResponseDetails`

NewCreditCardResponseDetailsWithDefaults instantiates a new CreditCardResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *CreditCardResponseDetails) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CreditCardResponseDetails) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CreditCardResponseDetails) SetForm(v string)`

SetForm sets Form field to given value.


### GetAccountId

`func (o *CreditCardResponseDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreditCardResponseDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreditCardResponseDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBusinessId

`func (o *CreditCardResponseDetails) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *CreditCardResponseDetails) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *CreditCardResponseDetails) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *CreditCardResponseDetails) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCardProductId

`func (o *CreditCardResponseDetails) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *CreditCardResponseDetails) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *CreditCardResponseDetails) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.


### GetCreationTime

`func (o *CreditCardResponseDetails) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreditCardResponseDetails) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreditCardResponseDetails) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomerId

`func (o *CreditCardResponseDetails) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreditCardResponseDetails) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreditCardResponseDetails) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CreditCardResponseDetails) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEmbossName

`func (o *CreditCardResponseDetails) GetEmbossName() EmbossName`

GetEmbossName returns the EmbossName field if non-nil, zero value otherwise.

### GetEmbossNameOk

`func (o *CreditCardResponseDetails) GetEmbossNameOk() (*EmbossName, bool)`

GetEmbossNameOk returns a tuple with the EmbossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbossName

`func (o *CreditCardResponseDetails) SetEmbossName(v EmbossName)`

SetEmbossName sets EmbossName field to given value.


### GetExpirationMonth

`func (o *CreditCardResponseDetails) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *CreditCardResponseDetails) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *CreditCardResponseDetails) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.


### GetExpirationTime

`func (o *CreditCardResponseDetails) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CreditCardResponseDetails) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CreditCardResponseDetails) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *CreditCardResponseDetails) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetExpirationYear

`func (o *CreditCardResponseDetails) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *CreditCardResponseDetails) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *CreditCardResponseDetails) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.


### GetId

`func (o *CreditCardResponseDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditCardResponseDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditCardResponseDetails) SetId(v string)`

SetId sets Id field to given value.


### GetIsPinSet

`func (o *CreditCardResponseDetails) GetIsPinSet() bool`

GetIsPinSet returns the IsPinSet field if non-nil, zero value otherwise.

### GetIsPinSetOk

`func (o *CreditCardResponseDetails) GetIsPinSetOk() (*bool, bool)`

GetIsPinSetOk returns a tuple with the IsPinSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinSet

`func (o *CreditCardResponseDetails) SetIsPinSet(v bool)`

SetIsPinSet sets IsPinSet field to given value.

### HasIsPinSet

`func (o *CreditCardResponseDetails) HasIsPinSet() bool`

HasIsPinSet returns a boolean if a field has been set.

### GetLastFour

`func (o *CreditCardResponseDetails) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *CreditCardResponseDetails) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *CreditCardResponseDetails) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.


### GetLastUpdatedTime

`func (o *CreditCardResponseDetails) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CreditCardResponseDetails) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CreditCardResponseDetails) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *CreditCardResponseDetails) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *CreditCardResponseDetails) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreditCardResponseDetails) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreditCardResponseDetails) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreditCardResponseDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReissueReason

`func (o *CreditCardResponseDetails) GetReissueReason() string`

GetReissueReason returns the ReissueReason field if non-nil, zero value otherwise.

### GetReissueReasonOk

`func (o *CreditCardResponseDetails) GetReissueReasonOk() (*string, bool)`

GetReissueReasonOk returns a tuple with the ReissueReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissueReason

`func (o *CreditCardResponseDetails) SetReissueReason(v string)`

SetReissueReason sets ReissueReason field to given value.

### HasReissueReason

`func (o *CreditCardResponseDetails) HasReissueReason() bool`

HasReissueReason returns a boolean if a field has been set.

### GetReissuedFromId

`func (o *CreditCardResponseDetails) GetReissuedFromId() string`

GetReissuedFromId returns the ReissuedFromId field if non-nil, zero value otherwise.

### GetReissuedFromIdOk

`func (o *CreditCardResponseDetails) GetReissuedFromIdOk() (*string, bool)`

GetReissuedFromIdOk returns a tuple with the ReissuedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedFromId

`func (o *CreditCardResponseDetails) SetReissuedFromId(v string)`

SetReissuedFromId sets ReissuedFromId field to given value.

### HasReissuedFromId

`func (o *CreditCardResponseDetails) HasReissuedFromId() bool`

HasReissuedFromId returns a boolean if a field has been set.

### GetReissuedToId

`func (o *CreditCardResponseDetails) GetReissuedToId() string`

GetReissuedToId returns the ReissuedToId field if non-nil, zero value otherwise.

### GetReissuedToIdOk

`func (o *CreditCardResponseDetails) GetReissuedToIdOk() (*string, bool)`

GetReissuedToIdOk returns a tuple with the ReissuedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedToId

`func (o *CreditCardResponseDetails) SetReissuedToId(v string)`

SetReissuedToId sets ReissuedToId field to given value.

### HasReissuedToId

`func (o *CreditCardResponseDetails) HasReissuedToId() bool`

HasReissuedToId returns a boolean if a field has been set.

### GetTimestampPinSet

`func (o *CreditCardResponseDetails) GetTimestampPinSet() time.Time`

GetTimestampPinSet returns the TimestampPinSet field if non-nil, zero value otherwise.

### GetTimestampPinSetOk

`func (o *CreditCardResponseDetails) GetTimestampPinSetOk() (*time.Time, bool)`

GetTimestampPinSetOk returns a tuple with the TimestampPinSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampPinSet

`func (o *CreditCardResponseDetails) SetTimestampPinSet(v time.Time)`

SetTimestampPinSet sets TimestampPinSet field to given value.

### HasTimestampPinSet

`func (o *CreditCardResponseDetails) HasTimestampPinSet() bool`

HasTimestampPinSet returns a boolean if a field has been set.

### GetBulkOrderConfigId

`func (o *CreditCardResponseDetails) GetBulkOrderConfigId() string`

GetBulkOrderConfigId returns the BulkOrderConfigId field if non-nil, zero value otherwise.

### GetBulkOrderConfigIdOk

`func (o *CreditCardResponseDetails) GetBulkOrderConfigIdOk() (*string, bool)`

GetBulkOrderConfigIdOk returns a tuple with the BulkOrderConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkOrderConfigId

`func (o *CreditCardResponseDetails) SetBulkOrderConfigId(v string)`

SetBulkOrderConfigId sets BulkOrderConfigId field to given value.

### HasBulkOrderConfigId

`func (o *CreditCardResponseDetails) HasBulkOrderConfigId() bool`

HasBulkOrderConfigId returns a boolean if a field has been set.

### GetCardImageId

`func (o *CreditCardResponseDetails) GetCardImageId() string`

GetCardImageId returns the CardImageId field if non-nil, zero value otherwise.

### GetCardImageIdOk

`func (o *CreditCardResponseDetails) GetCardImageIdOk() (*string, bool)`

GetCardImageIdOk returns a tuple with the CardImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardImageId

`func (o *CreditCardResponseDetails) SetCardImageId(v string)`

SetCardImageId sets CardImageId field to given value.

### HasCardImageId

`func (o *CreditCardResponseDetails) HasCardImageId() bool`

HasCardImageId returns a boolean if a field has been set.

### GetShipping

`func (o *CreditCardResponseDetails) GetShipping() Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *CreditCardResponseDetails) GetShippingOk() (*Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *CreditCardResponseDetails) SetShipping(v Shipping)`

SetShipping sets Shipping field to given value.


### GetCardStatus

`func (o *CreditCardResponseDetails) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *CreditCardResponseDetails) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *CreditCardResponseDetails) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetMemo

`func (o *CreditCardResponseDetails) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CreditCardResponseDetails) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CreditCardResponseDetails) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CreditCardResponseDetails) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPendingReasons

`func (o *CreditCardResponseDetails) GetPendingReasons() CardStatusPendingReasons`

GetPendingReasons returns the PendingReasons field if non-nil, zero value otherwise.

### GetPendingReasonsOk

`func (o *CreditCardResponseDetails) GetPendingReasonsOk() (*CardStatusPendingReasons, bool)`

GetPendingReasonsOk returns a tuple with the PendingReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReasons

`func (o *CreditCardResponseDetails) SetPendingReasons(v CardStatusPendingReasons)`

SetPendingReasons sets PendingReasons field to given value.

### HasPendingReasons

`func (o *CreditCardResponseDetails) HasPendingReasons() bool`

HasPendingReasons returns a boolean if a field has been set.

### GetStatusReason

`func (o *CreditCardResponseDetails) GetStatusReason() CardStatusReasonCode`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *CreditCardResponseDetails) GetStatusReasonOk() (*CardStatusReasonCode, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *CreditCardResponseDetails) SetStatusReason(v CardStatusReasonCode)`

SetStatusReason sets StatusReason field to given value.


### GetCardFulfillmentStatus

`func (o *CreditCardResponseDetails) GetCardFulfillmentStatus() CardFulfillmentStatus`

GetCardFulfillmentStatus returns the CardFulfillmentStatus field if non-nil, zero value otherwise.

### GetCardFulfillmentStatusOk

`func (o *CreditCardResponseDetails) GetCardFulfillmentStatusOk() (*CardFulfillmentStatus, bool)`

GetCardFulfillmentStatusOk returns a tuple with the CardFulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentStatus

`func (o *CreditCardResponseDetails) SetCardFulfillmentStatus(v CardFulfillmentStatus)`

SetCardFulfillmentStatus sets CardFulfillmentStatus field to given value.


### GetFulfillmentDetails

`func (o *CreditCardResponseDetails) GetFulfillmentDetails() FulfillmentDetails`

GetFulfillmentDetails returns the FulfillmentDetails field if non-nil, zero value otherwise.

### GetFulfillmentDetailsOk

`func (o *CreditCardResponseDetails) GetFulfillmentDetailsOk() (*FulfillmentDetails, bool)`

GetFulfillmentDetailsOk returns a tuple with the FulfillmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentDetails

`func (o *CreditCardResponseDetails) SetFulfillmentDetails(v FulfillmentDetails)`

SetFulfillmentDetails sets FulfillmentDetails field to given value.

### HasFulfillmentDetails

`func (o *CreditCardResponseDetails) HasFulfillmentDetails() bool`

HasFulfillmentDetails returns a boolean if a field has been set.

### GetBin

`func (o *CreditCardResponseDetails) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *CreditCardResponseDetails) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *CreditCardResponseDetails) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *CreditCardResponseDetails) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCardBrand

`func (o *CreditCardResponseDetails) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *CreditCardResponseDetails) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *CreditCardResponseDetails) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetPhysicalCardFormat

`func (o *CreditCardResponseDetails) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *CreditCardResponseDetails) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *CreditCardResponseDetails) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


