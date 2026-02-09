# SyncteraPayVendorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**SyncteraPayVendorConfig**](SyncteraPayVendorConfig.md) |  | 
**Description** | Pointer to **string** | A description of the vendor. | [optional] 
**Name** | **string** | The name of the vendor. | 
**Enabled** | **bool** | Whether the vendor is enabled. | 
**Id** | **string** | Unique identifier for the vendor. | 
**TenantId** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewSyncteraPayVendorResponse

`func NewSyncteraPayVendorResponse(config SyncteraPayVendorConfig, name string, enabled bool, id string, tenantId string, ) *SyncteraPayVendorResponse`

NewSyncteraPayVendorResponse instantiates a new SyncteraPayVendorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayVendorResponseWithDefaults

`func NewSyncteraPayVendorResponseWithDefaults() *SyncteraPayVendorResponse`

NewSyncteraPayVendorResponseWithDefaults instantiates a new SyncteraPayVendorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SyncteraPayVendorResponse) GetConfig() SyncteraPayVendorConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SyncteraPayVendorResponse) GetConfigOk() (*SyncteraPayVendorConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SyncteraPayVendorResponse) SetConfig(v SyncteraPayVendorConfig)`

SetConfig sets Config field to given value.


### GetDescription

`func (o *SyncteraPayVendorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyncteraPayVendorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyncteraPayVendorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyncteraPayVendorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SyncteraPayVendorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyncteraPayVendorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyncteraPayVendorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *SyncteraPayVendorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyncteraPayVendorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyncteraPayVendorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *SyncteraPayVendorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncteraPayVendorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncteraPayVendorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *SyncteraPayVendorResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SyncteraPayVendorResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SyncteraPayVendorResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


