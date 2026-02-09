# SyncteraPayConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Details** | [**SyncteraPayConfigurationDetails**](SyncteraPayConfigurationDetails.md) |  | 
**Enabled** | Pointer to **bool** | Whether or not the configuration is enabled. If the configuration is not enabled, it will not be used when creating an Outgoing Synctera Pay transfer.  | [optional] 
**Id** | Pointer to **string** | The ID of the Outgoing Synctera Pay configuration  | [optional] [readonly] 
**Name** | **string** |  | 
**Rules** | [**SyncteraPayConfigurationRules**](SyncteraPayConfigurationRules.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewSyncteraPayConfiguration

`func NewSyncteraPayConfiguration(details SyncteraPayConfigurationDetails, name string, rules SyncteraPayConfigurationRules, tenant string, ) *SyncteraPayConfiguration`

NewSyncteraPayConfiguration instantiates a new SyncteraPayConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayConfigurationWithDefaults

`func NewSyncteraPayConfigurationWithDefaults() *SyncteraPayConfiguration`

NewSyncteraPayConfigurationWithDefaults instantiates a new SyncteraPayConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SyncteraPayConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyncteraPayConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyncteraPayConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyncteraPayConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *SyncteraPayConfiguration) GetDetails() SyncteraPayConfigurationDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SyncteraPayConfiguration) GetDetailsOk() (*SyncteraPayConfigurationDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SyncteraPayConfiguration) SetDetails(v SyncteraPayConfigurationDetails)`

SetDetails sets Details field to given value.


### GetEnabled

`func (o *SyncteraPayConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyncteraPayConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyncteraPayConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SyncteraPayConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *SyncteraPayConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncteraPayConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncteraPayConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SyncteraPayConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SyncteraPayConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyncteraPayConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyncteraPayConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetRules

`func (o *SyncteraPayConfiguration) GetRules() SyncteraPayConfigurationRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SyncteraPayConfiguration) GetRulesOk() (*SyncteraPayConfigurationRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SyncteraPayConfiguration) SetRules(v SyncteraPayConfigurationRules)`

SetRules sets Rules field to given value.


### GetTenant

`func (o *SyncteraPayConfiguration) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *SyncteraPayConfiguration) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *SyncteraPayConfiguration) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


