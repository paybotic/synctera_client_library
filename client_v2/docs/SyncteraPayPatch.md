# SyncteraPayPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The amount in the source currency&#39;s minor unit. For example, 10000 would be $100 for USD.  One of the amount or percentage is required.  | [optional] 
**ExchangeDetails** | Pointer to [**ExchangeDetails**](ExchangeDetails.md) |  | [optional] 
**FinalExternalAccountId** | Pointer to **string** | The ID of the final external account that will be the receiver of the Synctera Pay transfer. | [optional] 
**ReferenceId** | Pointer to **string** | The reference id of the transfer. | [optional] 
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SyncteraPayVendorId** | Pointer to **string** | The ID of the vendor that will be used to process the Synctera Pay transfer. | [optional] 

## Methods

### NewSyncteraPayPatch

`func NewSyncteraPayPatch() *SyncteraPayPatch`

NewSyncteraPayPatch instantiates a new SyncteraPayPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayPatchWithDefaults

`func NewSyncteraPayPatchWithDefaults() *SyncteraPayPatch`

NewSyncteraPayPatchWithDefaults instantiates a new SyncteraPayPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SyncteraPayPatch) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SyncteraPayPatch) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SyncteraPayPatch) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SyncteraPayPatch) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExchangeDetails

`func (o *SyncteraPayPatch) GetExchangeDetails() ExchangeDetails`

GetExchangeDetails returns the ExchangeDetails field if non-nil, zero value otherwise.

### GetExchangeDetailsOk

`func (o *SyncteraPayPatch) GetExchangeDetailsOk() (*ExchangeDetails, bool)`

GetExchangeDetailsOk returns a tuple with the ExchangeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDetails

`func (o *SyncteraPayPatch) SetExchangeDetails(v ExchangeDetails)`

SetExchangeDetails sets ExchangeDetails field to given value.

### HasExchangeDetails

`func (o *SyncteraPayPatch) HasExchangeDetails() bool`

HasExchangeDetails returns a boolean if a field has been set.

### GetFinalExternalAccountId

`func (o *SyncteraPayPatch) GetFinalExternalAccountId() string`

GetFinalExternalAccountId returns the FinalExternalAccountId field if non-nil, zero value otherwise.

### GetFinalExternalAccountIdOk

`func (o *SyncteraPayPatch) GetFinalExternalAccountIdOk() (*string, bool)`

GetFinalExternalAccountIdOk returns a tuple with the FinalExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalExternalAccountId

`func (o *SyncteraPayPatch) SetFinalExternalAccountId(v string)`

SetFinalExternalAccountId sets FinalExternalAccountId field to given value.

### HasFinalExternalAccountId

`func (o *SyncteraPayPatch) HasFinalExternalAccountId() bool`

HasFinalExternalAccountId returns a boolean if a field has been set.

### GetReferenceId

`func (o *SyncteraPayPatch) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SyncteraPayPatch) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SyncteraPayPatch) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *SyncteraPayPatch) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetSourceData

`func (o *SyncteraPayPatch) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *SyncteraPayPatch) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *SyncteraPayPatch) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *SyncteraPayPatch) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetStatus

`func (o *SyncteraPayPatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncteraPayPatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncteraPayPatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyncteraPayPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncteraPayVendorId

`func (o *SyncteraPayPatch) GetSyncteraPayVendorId() string`

GetSyncteraPayVendorId returns the SyncteraPayVendorId field if non-nil, zero value otherwise.

### GetSyncteraPayVendorIdOk

`func (o *SyncteraPayPatch) GetSyncteraPayVendorIdOk() (*string, bool)`

GetSyncteraPayVendorIdOk returns a tuple with the SyncteraPayVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncteraPayVendorId

`func (o *SyncteraPayPatch) SetSyncteraPayVendorId(v string)`

SetSyncteraPayVendorId sets SyncteraPayVendorId field to given value.

### HasSyncteraPayVendorId

`func (o *SyncteraPayPatch) HasSyncteraPayVendorId() bool`

HasSyncteraPayVendorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


