# SyncteraPayPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transfer amount in cents | 
**ConfigurationId** | Pointer to **string** | The ID of the Outgoing Synctera Pay configuration that will be used to match the Synctera Pay transfer to the appropriate configuration. Only required if the transfer is not automatically matched to a configuration.  | [optional] 
**Currency** | **string** | The currency of the transfer in ISO 4217 format | 
**DcSign** | **string** | The debit/credit sign of the transfer. This is a legacy field for backward compatibility. For outgoing synctera pay, customer accounts are always debited. For incoming synctera pay, customer accounts are always credited. The direction of the transfer is determined by the API route that is used to create the transfer. For outgoing synctera pay, the dc_sign must be specified and it must always be &#39;CREDIT&#39;.  | 
**Direction** | [**SyncteraPayDirection**](SyncteraPayDirection.md) |  | 
**ExchangeDetails** | Pointer to [**ExchangeDetails**](ExchangeDetails.md) |  | [optional] 
**FinalExternalAccountId** | Pointer to **string** | The ID of the final external account that will be the receiver of the Outgoing Synctera Pay transfer. Required unless previously agreed with Synctera.  | [optional] 
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 
**Subtype** | [**SyncteraPaySubtype**](SyncteraPaySubtype.md) |  | 
**SyncteraPayNetwork** | **string** | The network of the transfer.  | 
**SyncteraPayVendorId** | Pointer to **string** | The ID of the vendor that will be used to process the Synctera Pay transfer. | [optional] 
**AccountId** | **string** | The UUID of the Synctera account resource of the customer.  | 
**CustomerId** | **string** | The UUID of the Synctera customer resource.  | 
**EffectiveDate** | **string** | The effective date of the transaction once it gets posted | 
**ReferenceId** | Pointer to **string** | The network reference id of the transfer, this must be supplied by the vendor. | [optional] 

## Methods

### NewSyncteraPayPost

`func NewSyncteraPayPost(amount int64, currency string, dcSign string, direction SyncteraPayDirection, subtype SyncteraPaySubtype, syncteraPayNetwork string, accountId string, customerId string, effectiveDate string, ) *SyncteraPayPost`

NewSyncteraPayPost instantiates a new SyncteraPayPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayPostWithDefaults

`func NewSyncteraPayPostWithDefaults() *SyncteraPayPost`

NewSyncteraPayPostWithDefaults instantiates a new SyncteraPayPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SyncteraPayPost) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SyncteraPayPost) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SyncteraPayPost) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetConfigurationId

`func (o *SyncteraPayPost) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *SyncteraPayPost) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *SyncteraPayPost) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *SyncteraPayPost) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetCurrency

`func (o *SyncteraPayPost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SyncteraPayPost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SyncteraPayPost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDcSign

`func (o *SyncteraPayPost) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *SyncteraPayPost) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *SyncteraPayPost) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetDirection

`func (o *SyncteraPayPost) GetDirection() SyncteraPayDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SyncteraPayPost) GetDirectionOk() (*SyncteraPayDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SyncteraPayPost) SetDirection(v SyncteraPayDirection)`

SetDirection sets Direction field to given value.


### GetExchangeDetails

`func (o *SyncteraPayPost) GetExchangeDetails() ExchangeDetails`

GetExchangeDetails returns the ExchangeDetails field if non-nil, zero value otherwise.

### GetExchangeDetailsOk

`func (o *SyncteraPayPost) GetExchangeDetailsOk() (*ExchangeDetails, bool)`

GetExchangeDetailsOk returns a tuple with the ExchangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDetails

`func (o *SyncteraPayPost) SetExchangeDetails(v ExchangeDetails)`

SetExchangeDetails sets ExchangeDetails field to given value.

### HasExchangeDetails

`func (o *SyncteraPayPost) HasExchangeDetails() bool`

HasExchangeDetails returns a boolean if a field has been set.

### GetFinalExternalAccountId

`func (o *SyncteraPayPost) GetFinalExternalAccountId() string`

GetFinalExternalAccountId returns the FinalExternalAccountId field if non-nil, zero value otherwise.

### GetFinalExternalAccountIdOk

`func (o *SyncteraPayPost) GetFinalExternalAccountIdOk() (*string, bool)`

GetFinalExternalAccountIdOk returns a tuple with the FinalExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalExternalAccountId

`func (o *SyncteraPayPost) SetFinalExternalAccountId(v string)`

SetFinalExternalAccountId sets FinalExternalAccountId field to given value.

### HasFinalExternalAccountId

`func (o *SyncteraPayPost) HasFinalExternalAccountId() bool`

HasFinalExternalAccountId returns a boolean if a field has been set.

### GetSourceData

`func (o *SyncteraPayPost) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *SyncteraPayPost) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *SyncteraPayPost) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *SyncteraPayPost) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSubtype

`func (o *SyncteraPayPost) GetSubtype() SyncteraPaySubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SyncteraPayPost) GetSubtypeOk() (*SyncteraPaySubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SyncteraPayPost) SetSubtype(v SyncteraPaySubtype)`

SetSubtype sets Subtype field to given value.


### GetSyncteraPayNetwork

`func (o *SyncteraPayPost) GetSyncteraPayNetwork() string`

GetSyncteraPayNetwork returns the SyncteraPayNetwork field if non-nil, zero value otherwise.

### GetSyncteraPayNetworkOk

`func (o *SyncteraPayPost) GetSyncteraPayNetworkOk() (*string, bool)`

GetSyncteraPayNetworkOk returns a tuple with the SyncteraPayNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncteraPayNetwork

`func (o *SyncteraPayPost) SetSyncteraPayNetwork(v string)`

SetSyncteraPayNetwork sets SyncteraPayNetwork field to given value.


### GetSyncteraPayVendorId

`func (o *SyncteraPayPost) GetSyncteraPayVendorId() string`

GetSyncteraPayVendorId returns the SyncteraPayVendorId field if non-nil, zero value otherwise.

### GetSyncteraPayVendorIdOk

`func (o *SyncteraPayPost) GetSyncteraPayVendorIdOk() (*string, bool)`

GetSyncteraPayVendorIdOk returns a tuple with the SyncteraPayVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncteraPayVendorId

`func (o *SyncteraPayPost) SetSyncteraPayVendorId(v string)`

SetSyncteraPayVendorId sets SyncteraPayVendorId field to given value.

### HasSyncteraPayVendorId

`func (o *SyncteraPayPost) HasSyncteraPayVendorId() bool`

HasSyncteraPayVendorId returns a boolean if a field has been set.

### GetAccountId

`func (o *SyncteraPayPost) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SyncteraPayPost) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SyncteraPayPost) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomerId

`func (o *SyncteraPayPost) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SyncteraPayPost) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SyncteraPayPost) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEffectiveDate

`func (o *SyncteraPayPost) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *SyncteraPayPost) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *SyncteraPayPost) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetReferenceId

`func (o *SyncteraPayPost) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SyncteraPayPost) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SyncteraPayPost) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *SyncteraPayPost) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


