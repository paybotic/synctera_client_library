# BaseCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | 
**AccountId** | Pointer to **string** | The ID of the account to which the card will be linked | [optional] 
**BusinessId** | Pointer to **string** | The business ID associated with this card. If no customer_id is supplied, a card can still be issued to a business, but cannot be activated or used until a customer is assigned via the PATCH /cards/{card_id} endpoint. | [optional] 
**CardProductId** | Pointer to **string** | The card product to which the card is attached | [optional] 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the card issuance request was made | [optional] [readonly] 
**CustomerId** | Pointer to **string** | The ID of the customer to whom the card will be issued. If a business_id is passed, but a customer_id not passed at the time of card creation the card cannot be activated or used for spend until it&#39;s assigned to a human customer via the PATCH /cards/{card_id} endpoint. If no business_id is passed, a customer_id is required. | [optional] 
**EmbossName** | Pointer to [**EmbossName**](EmbossName.md) |  | [optional] 
**ExpirationMonth** | Pointer to **string** |  | [optional] [readonly] 
**ExpirationTime** | Pointer to **time.Time** | The timestamp representing when the card would expire at | [optional] [readonly] 
**ExpirationYear** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** | Card ID | [optional] [readonly] 
**IsPinSet** | Pointer to **bool** | indicates whether a pin has been set on the card | [optional] [readonly] [default to false]
**LastFour** | Pointer to **string** | The last 4 digits of the card PAN | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the card was last modified at | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** | Additional data to include in the request structured as key-value pairs | [optional] 
**ReissueReason** | Pointer to **string** | This is the reason the card needs to be reissued, if any. The reason determines several behaviours:   - whether or not the new card will use the same PAN as the original card   - the old card will be terminated and if so, when it will be terminated  Reason                 | Same PAN | Terminate Old Card ---------------------- | -------- | ------------------ EXPIRATION             | yes      | on activation LOST                   | no       | immediately STOLEN                 | no       | immediately DAMAGED                | yes      | on activation VIRTUAL_TO_PHYSICAL(*) | yes      | on activation PRODUCT_CHANGE         | yes      | on activation NAME_CHANGE(**)        | yes      | on activation APPEARANCE             | yes      | on activation  (*) VIRTUAL_TO_PHYSICAL is deprecated. Please use PRODUCT_CHANGE whenever reissuing from one card product to another, including from a virtual product to a physical product.  (**) NAME_CHANGE is deprecated. Please use APPEARANCE whenever reissuing in order to change the appearance of a card, such as the printed name or custom image.  For all reasons, the new card will use the same PIN as the original card and digital wallet tokens will reassigned to the new card  | [optional] 
**ReissuedFromId** | Pointer to **string** | When reissuing a card, specify the card to be replaced here. When getting a card&#39;s details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced. If this field is set, then reissue_reason must also be set.  | [optional] 
**ReissuedToId** | Pointer to **string** | If this card was reissued, this ID refers to the card that replaced it. | [optional] [readonly] 
**TimestampPinSet** | Pointer to **time.Time** | Time when the PIN was last set or changed. | [optional] [readonly] 
**Type** | Pointer to [**CardType**](CardType.md) |  | [optional] 

## Methods

### NewBaseCard

`func NewBaseCard(form string, ) *BaseCard`

NewBaseCard instantiates a new BaseCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCardWithDefaults

`func NewBaseCardWithDefaults() *BaseCard`

NewBaseCardWithDefaults instantiates a new BaseCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *BaseCard) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *BaseCard) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *BaseCard) SetForm(v string)`

SetForm sets Form field to given value.


### GetAccountId

`func (o *BaseCard) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BaseCard) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BaseCard) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BaseCard) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBusinessId

`func (o *BaseCard) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *BaseCard) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *BaseCard) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *BaseCard) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCardProductId

`func (o *BaseCard) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *BaseCard) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *BaseCard) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.

### HasCardProductId

`func (o *BaseCard) HasCardProductId() bool`

HasCardProductId returns a boolean if a field has been set.

### GetCreationTime

`func (o *BaseCard) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BaseCard) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BaseCard) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BaseCard) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCustomerId

`func (o *BaseCard) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *BaseCard) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *BaseCard) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *BaseCard) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEmbossName

`func (o *BaseCard) GetEmbossName() EmbossName`

GetEmbossName returns the EmbossName field if non-nil, zero value otherwise.

### GetEmbossNameOk

`func (o *BaseCard) GetEmbossNameOk() (*EmbossName, bool)`

GetEmbossNameOk returns a tuple with the EmbossName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbossName

`func (o *BaseCard) SetEmbossName(v EmbossName)`

SetEmbossName sets EmbossName field to given value.

### HasEmbossName

`func (o *BaseCard) HasEmbossName() bool`

HasEmbossName returns a boolean if a field has been set.

### GetExpirationMonth

`func (o *BaseCard) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *BaseCard) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *BaseCard) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.

### HasExpirationMonth

`func (o *BaseCard) HasExpirationMonth() bool`

HasExpirationMonth returns a boolean if a field has been set.

### GetExpirationTime

`func (o *BaseCard) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *BaseCard) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *BaseCard) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *BaseCard) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetExpirationYear

`func (o *BaseCard) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *BaseCard) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *BaseCard) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.

### HasExpirationYear

`func (o *BaseCard) HasExpirationYear() bool`

HasExpirationYear returns a boolean if a field has been set.

### GetId

`func (o *BaseCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPinSet

`func (o *BaseCard) GetIsPinSet() bool`

GetIsPinSet returns the IsPinSet field if non-nil, zero value otherwise.

### GetIsPinSetOk

`func (o *BaseCard) GetIsPinSetOk() (*bool, bool)`

GetIsPinSetOk returns a tuple with the IsPinSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinSet

`func (o *BaseCard) SetIsPinSet(v bool)`

SetIsPinSet sets IsPinSet field to given value.

### HasIsPinSet

`func (o *BaseCard) HasIsPinSet() bool`

HasIsPinSet returns a boolean if a field has been set.

### GetLastFour

`func (o *BaseCard) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *BaseCard) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *BaseCard) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *BaseCard) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *BaseCard) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BaseCard) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BaseCard) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *BaseCard) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *BaseCard) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BaseCard) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BaseCard) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BaseCard) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReissueReason

`func (o *BaseCard) GetReissueReason() string`

GetReissueReason returns the ReissueReason field if non-nil, zero value otherwise.

### GetReissueReasonOk

`func (o *BaseCard) GetReissueReasonOk() (*string, bool)`

GetReissueReasonOk returns a tuple with the ReissueReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissueReason

`func (o *BaseCard) SetReissueReason(v string)`

SetReissueReason sets ReissueReason field to given value.

### HasReissueReason

`func (o *BaseCard) HasReissueReason() bool`

HasReissueReason returns a boolean if a field has been set.

### GetReissuedFromId

`func (o *BaseCard) GetReissuedFromId() string`

GetReissuedFromId returns the ReissuedFromId field if non-nil, zero value otherwise.

### GetReissuedFromIdOk

`func (o *BaseCard) GetReissuedFromIdOk() (*string, bool)`

GetReissuedFromIdOk returns a tuple with the ReissuedFromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedFromId

`func (o *BaseCard) SetReissuedFromId(v string)`

SetReissuedFromId sets ReissuedFromId field to given value.

### HasReissuedFromId

`func (o *BaseCard) HasReissuedFromId() bool`

HasReissuedFromId returns a boolean if a field has been set.

### GetReissuedToId

`func (o *BaseCard) GetReissuedToId() string`

GetReissuedToId returns the ReissuedToId field if non-nil, zero value otherwise.

### GetReissuedToIdOk

`func (o *BaseCard) GetReissuedToIdOk() (*string, bool)`

GetReissuedToIdOk returns a tuple with the ReissuedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuedToId

`func (o *BaseCard) SetReissuedToId(v string)`

SetReissuedToId sets ReissuedToId field to given value.

### HasReissuedToId

`func (o *BaseCard) HasReissuedToId() bool`

HasReissuedToId returns a boolean if a field has been set.

### GetTimestampPinSet

`func (o *BaseCard) GetTimestampPinSet() time.Time`

GetTimestampPinSet returns the TimestampPinSet field if non-nil, zero value otherwise.

### GetTimestampPinSetOk

`func (o *BaseCard) GetTimestampPinSetOk() (*time.Time, bool)`

GetTimestampPinSetOk returns a tuple with the TimestampPinSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampPinSet

`func (o *BaseCard) SetTimestampPinSet(v time.Time)`

SetTimestampPinSet sets TimestampPinSet field to given value.

### HasTimestampPinSet

`func (o *BaseCard) HasTimestampPinSet() bool`

HasTimestampPinSet returns a boolean if a field has been set.

### GetType

`func (o *BaseCard) GetType() CardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseCard) GetTypeOk() (*CardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseCard) SetType(v CardType)`

SetType sets Type field to given value.

### HasType

`func (o *BaseCard) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


