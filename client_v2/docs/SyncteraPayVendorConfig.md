# SyncteraPayVendorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueReferenceIds** | Pointer to **bool** | Whether the vendor supports unique reference IDs. If true, synctera pay will check for unique reference IDs. If a transfer with the same reference ID is found, the transfer will be rejected with a duplicate reference ID error.  | [optional] 

## Methods

### NewSyncteraPayVendorConfig

`func NewSyncteraPayVendorConfig() *SyncteraPayVendorConfig`

NewSyncteraPayVendorConfig instantiates a new SyncteraPayVendorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayVendorConfigWithDefaults

`func NewSyncteraPayVendorConfigWithDefaults() *SyncteraPayVendorConfig`

NewSyncteraPayVendorConfigWithDefaults instantiates a new SyncteraPayVendorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueReferenceIds

`func (o *SyncteraPayVendorConfig) GetUniqueReferenceIds() bool`

GetUniqueReferenceIds returns the UniqueReferenceIds field if non-nil, zero value otherwise.

### GetUniqueReferenceIdsOk

`func (o *SyncteraPayVendorConfig) GetUniqueReferenceIdsOk() (*bool, bool)`

GetUniqueReferenceIdsOk returns a tuple with the UniqueReferenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueReferenceIds

`func (o *SyncteraPayVendorConfig) SetUniqueReferenceIds(v bool)`

SetUniqueReferenceIds sets UniqueReferenceIds field to given value.

### HasUniqueReferenceIds

`func (o *SyncteraPayVendorConfig) HasUniqueReferenceIds() bool`

HasUniqueReferenceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


