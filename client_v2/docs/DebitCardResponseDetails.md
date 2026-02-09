# DebitCardResponseDetails

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

### NewDebitCardResponseDetails

`func NewDebitCardResponseDetails(form string, accountId string, cardProductId string, creationTime time.Time, embossName EmbossName, expirationMonth string, expirationYear string, id string, lastFour string, shipping Shipping, cardStatus CardStatus, statusReason CardStatusReasonCode, cardFulfillmentStatus CardFulfillmentStatus, cardBrand CardBrand, physicalCardFormat PhysicalCardFormat, ) *DebitCardResponseDetails`

NewDebitCardResponseDetails instantiates a new DebitCardResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebitCardResponseDetailsWithDefaults

`func NewDebitCardResponseDetailsWithDefaults() *DebitCardResponseDetails`

NewDebitCardResponseDetailsWithDefaults instantiates a new DebitCardResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *DebitCardResponseDetails) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *DebitCardResponseDetails) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *DebitCardResponseDetails) SetForm(v string)`

SetForm sets Form field to given value.


### GetAccountId

`func (o *DebitCardResponseDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DebitCardResponseDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DebitCardResponseDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBusinessId

`func (o *DebitCardResponseDetails) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *DebitCardResponseDetails) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *DebitCardResponseDetails) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *DebitCardResponseDetails) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCardProductId

`func (o *DebitCardResponseDetails) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *DebitCardResponseDetails) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *DebitCardResponseDetails) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.


### GetCreationTime

`func (o *DebitCardResponseDetails) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DebitCardResponseDetails) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DebitCardResponseDetails) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomerId

`func (o *DebitCardResponseDetails) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DebitCardResponseDetails) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DebitCardResponseDetails) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DebitCardResponseDetails) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEmbossName

`func (o *DebitCardResponseDetails) GetEmbossName() EmbossName`

GetEmbossName returns the EmbossName field if non-nil, zero value otherwise.

### GetEmbossNameOk

`func (o *DebitCardResponseDetails) GetEmbossNameOk() (*EmbossName, bool)`

GetEmbossNameOk returns a tuple with the EmbossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbossName

`func (o *DebitCardResponseDetails) SetEmbossName(v EmbossName)`

SetEmbossName sets EmbossName field to given value.


### GetExpirationMonth

`func (o *DebitCardResponseDetails) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *DebitCardResponseDetails) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *DebitCardResponseDetails) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.


### GetExpirationTime

`func (o *DebitCardResponseDetails) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *DebitCardResponseDetails) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *DebitCardResponseDetails) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *DebitCardResponseDetails) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetExpirationYear

`func (o *DebitCardResponseDetails) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *DebitCardResponseDetails) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *DebitCardResponseDetails) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.


### GetId

`func (o *DebitCardResponseDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DebitCardResponseDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DebitCardResponseDetails) SetId(v string)`

SetId sets Id field to given value.


### GetIsPinSet

`func (o *DebitCardResponseDetails) GetIsPinSet() bool`

GetIsPinSet returns the IsPinSet field if non-nil, zero value otherwise.

### GetIsPinSetOk

`func (o *DebitCardResponseDetails) GetIsPinSetOk() (*bool, bool)`

GetIsPinSetOk returns a tuple with the IsPinSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinSet

`func (o *DebitCardResponseDetails) SetIsPinSet(v bool)`

SetIsPinSet sets IsPinSet field to given value.

### HasIsPinSet

`func (o *DebitCardResponseDetails) HasIsPinSet() bool`

HasIsPinSet returns a boolean if a field has been set.

### GetLastFour

`func (o *DebitCardResponseDetails) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *DebitCardResponseDetails) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *DebitCardResponseDetails) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.


### GetLastUpdatedTime

`func (o *DebitCardResponseDetails) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *DebitCardResponseDetails) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *DebitCardResponseDetails) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *DebitCardResponseDetails) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *DebitCardResponseDetails) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DebitCardResponseDetails) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DebitCardResponseDetails) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DebitCardResponseDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReissueReason

`func (o *DebitCardResponseDetails) GetReissueReason() string`

GetReissueReason returns the ReissueReason field if non-nil, zero value otherwise.

### GetReissueReasonOk

`func (o *DebitCardResponseDetails) GetReissueReasonOk() (*string, bool)`

GetReissueReasonOk returns a tuple with the ReissueReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissueReason

`func (o *DebitCardResponseDetails) SetReissueReason(v string)`

SetReissueReason sets ReissueReason field to given value.

### HasReissueReason

`func (o *DebitCardResponseDetails) HasReissueReason() bool`

HasReissueReason returns a boolean if a field has been set.

### GetReissuedFromId

`func (o *DebitCardResponseDetails) GetReissuedFromId() string`

GetReissuedFromId returns the ReissuedFromId field if non-nil, zero value otherwise.

### GetReissuedFromIdOk

`func (o *DebitCardResponseDetails) GetReissuedFromIdOk() (*string, bool)`

GetReissuedFromIdOk returns a tuple with the ReissuedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedFromId

`func (o *DebitCardResponseDetails) SetReissuedFromId(v string)`

SetReissuedFromId sets ReissuedFromId field to given value.

### HasReissuedFromId

`func (o *DebitCardResponseDetails) HasReissuedFromId() bool`

HasReissuedFromId returns a boolean if a field has been set.

### GetReissuedToId

`func (o *DebitCardResponseDetails) GetReissuedToId() string`

GetReissuedToId returns the ReissuedToId field if non-nil, zero value otherwise.

### GetReissuedToIdOk

`func (o *DebitCardResponseDetails) GetReissuedToIdOk() (*string, bool)`

GetReissuedToIdOk returns a tuple with the ReissuedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedToId

`func (o *DebitCardResponseDetails) SetReissuedToId(v string)`

SetReissuedToId sets ReissuedToId field to given value.

### HasReissuedToId

`func (o *DebitCardResponseDetails) HasReissuedToId() bool`

HasReissuedToId returns a boolean if a field has been set.

### GetTimestampPinSet

`func (o *DebitCardResponseDetails) GetTimestampPinSet() time.Time`

GetTimestampPinSet returns the TimestampPinSet field if non-nil, zero value otherwise.

### GetTimestampPinSetOk

`func (o *DebitCardResponseDetails) GetTimestampPinSetOk() (*time.Time, bool)`

GetTimestampPinSetOk returns a tuple with the TimestampPinSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampPinSet

`func (o *DebitCardResponseDetails) SetTimestampPinSet(v time.Time)`

SetTimestampPinSet sets TimestampPinSet field to given value.

### HasTimestampPinSet

`func (o *DebitCardResponseDetails) HasTimestampPinSet() bool`

HasTimestampPinSet returns a boolean if a field has been set.

### GetBulkOrderConfigId

`func (o *DebitCardResponseDetails) GetBulkOrderConfigId() string`

GetBulkOrderConfigId returns the BulkOrderConfigId field if non-nil, zero value otherwise.

### GetBulkOrderConfigIdOk

`func (o *DebitCardResponseDetails) GetBulkOrderConfigIdOk() (*string, bool)`

GetBulkOrderConfigIdOk returns a tuple with the BulkOrderConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkOrderConfigId

`func (o *DebitCardResponseDetails) SetBulkOrderConfigId(v string)`

SetBulkOrderConfigId sets BulkOrderConfigId field to given value.

### HasBulkOrderConfigId

`func (o *DebitCardResponseDetails) HasBulkOrderConfigId() bool`

HasBulkOrderConfigId returns a boolean if a field has been set.

### GetCardImageId

`func (o *DebitCardResponseDetails) GetCardImageId() string`

GetCardImageId returns the CardImageId field if non-nil, zero value otherwise.

### GetCardImageIdOk

`func (o *DebitCardResponseDetails) GetCardImageIdOk() (*string, bool)`

GetCardImageIdOk returns a tuple with the CardImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardImageId

`func (o *DebitCardResponseDetails) SetCardImageId(v string)`

SetCardImageId sets CardImageId field to given value.

### HasCardImageId

`func (o *DebitCardResponseDetails) HasCardImageId() bool`

HasCardImageId returns a boolean if a field has been set.

### GetShipping

`func (o *DebitCardResponseDetails) GetShipping() Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *DebitCardResponseDetails) GetShippingOk() (*Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *DebitCardResponseDetails) SetShipping(v Shipping)`

SetShipping sets Shipping field to given value.


### GetCardStatus

`func (o *DebitCardResponseDetails) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *DebitCardResponseDetails) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *DebitCardResponseDetails) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetMemo

`func (o *DebitCardResponseDetails) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *DebitCardResponseDetails) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *DebitCardResponseDetails) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *DebitCardResponseDetails) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPendingReasons

`func (o *DebitCardResponseDetails) GetPendingReasons() CardStatusPendingReasons`

GetPendingReasons returns the PendingReasons field if non-nil, zero value otherwise.

### GetPendingReasonsOk

`func (o *DebitCardResponseDetails) GetPendingReasonsOk() (*CardStatusPendingReasons, bool)`

GetPendingReasonsOk returns a tuple with the PendingReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReasons

`func (o *DebitCardResponseDetails) SetPendingReasons(v CardStatusPendingReasons)`

SetPendingReasons sets PendingReasons field to given value.

### HasPendingReasons

`func (o *DebitCardResponseDetails) HasPendingReasons() bool`

HasPendingReasons returns a boolean if a field has been set.

### GetStatusReason

`func (o *DebitCardResponseDetails) GetStatusReason() CardStatusReasonCode`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *DebitCardResponseDetails) GetStatusReasonOk() (*CardStatusReasonCode, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *DebitCardResponseDetails) SetStatusReason(v CardStatusReasonCode)`

SetStatusReason sets StatusReason field to given value.


### GetCardFulfillmentStatus

`func (o *DebitCardResponseDetails) GetCardFulfillmentStatus() CardFulfillmentStatus`

GetCardFulfillmentStatus returns the CardFulfillmentStatus field if non-nil, zero value otherwise.

### GetCardFulfillmentStatusOk

`func (o *DebitCardResponseDetails) GetCardFulfillmentStatusOk() (*CardFulfillmentStatus, bool)`

GetCardFulfillmentStatusOk returns a tuple with the CardFulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentStatus

`func (o *DebitCardResponseDetails) SetCardFulfillmentStatus(v CardFulfillmentStatus)`

SetCardFulfillmentStatus sets CardFulfillmentStatus field to given value.


### GetFulfillmentDetails

`func (o *DebitCardResponseDetails) GetFulfillmentDetails() FulfillmentDetails`

GetFulfillmentDetails returns the FulfillmentDetails field if non-nil, zero value otherwise.

### GetFulfillmentDetailsOk

`func (o *DebitCardResponseDetails) GetFulfillmentDetailsOk() (*FulfillmentDetails, bool)`

GetFulfillmentDetailsOk returns a tuple with the FulfillmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentDetails

`func (o *DebitCardResponseDetails) SetFulfillmentDetails(v FulfillmentDetails)`

SetFulfillmentDetails sets FulfillmentDetails field to given value.

### HasFulfillmentDetails

`func (o *DebitCardResponseDetails) HasFulfillmentDetails() bool`

HasFulfillmentDetails returns a boolean if a field has been set.

### GetBin

`func (o *DebitCardResponseDetails) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *DebitCardResponseDetails) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *DebitCardResponseDetails) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *DebitCardResponseDetails) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCardBrand

`func (o *DebitCardResponseDetails) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *DebitCardResponseDetails) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *DebitCardResponseDetails) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetPhysicalCardFormat

`func (o *DebitCardResponseDetails) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *DebitCardResponseDetails) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *DebitCardResponseDetails) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


