# SyncteraPayVendorBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**SyncteraPayVendorConfig**](SyncteraPayVendorConfig.md) |  | 
**Description** | Pointer to **string** | A description of the vendor. | [optional] 
**Name** | **string** | The name of the vendor. | 

## Methods

### NewSyncteraPayVendorBase

`func NewSyncteraPayVendorBase(config SyncteraPayVendorConfig, name string, ) *SyncteraPayVendorBase`

NewSyncteraPayVendorBase instantiates a new SyncteraPayVendorBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayVendorBaseWithDefaults

`func NewSyncteraPayVendorBaseWithDefaults() *SyncteraPayVendorBase`

NewSyncteraPayVendorBaseWithDefaults instantiates a new SyncteraPayVendorBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SyncteraPayVendorBase) GetConfig() SyncteraPayVendorConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SyncteraPayVendorBase) GetConfigOk() (*SyncteraPayVendorConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SyncteraPayVendorBase) SetConfig(v SyncteraPayVendorConfig)`

SetConfig sets Config field to given value.


### GetDescription

`func (o *SyncteraPayVendorBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyncteraPayVendorBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyncteraPayVendorBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyncteraPayVendorBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SyncteraPayVendorBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyncteraPayVendorBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyncteraPayVendorBase) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


