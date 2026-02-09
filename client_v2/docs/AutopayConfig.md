# AutopayConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | The business ID that owns this autopay configuration (mutually exclusive with person_id) | [optional] 
**Config** | [**AutopayConfigData**](AutopayConfigData.md) |  | 
**CreationTime** | **time.Time** | Timestamp when the configuration was created | 
**Id** | **string** | Unique identifier for the autopay configuration | 
**LastUpdatedTime** | **time.Time** | Timestamp when the configuration was last updated | 
**LendingAccountId** | **string** | The lending account ID this configuration belongs to | 
**PersonId** | Pointer to **string** | The person ID who owns this autopay configuration (mutually exclusive with business_id) | [optional] 
**Status** | [**AutopayConfigStatus**](AutopayConfigStatus.md) |  | 
**Tenant** | Pointer to [**Tenant**](Tenant.md) |  | [optional] 

## Methods

### NewAutopayConfig

`func NewAutopayConfig(config AutopayConfigData, creationTime time.Time, id string, lastUpdatedTime time.Time, lendingAccountId string, status AutopayConfigStatus, ) *AutopayConfig`

NewAutopayConfig instantiates a new AutopayConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayConfigWithDefaults

`func NewAutopayConfigWithDefaults() *AutopayConfig`

NewAutopayConfigWithDefaults instantiates a new AutopayConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *AutopayConfig) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AutopayConfig) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AutopayConfig) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AutopayConfig) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetConfig

`func (o *AutopayConfig) GetConfig() AutopayConfigData`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AutopayConfig) GetConfigOk() (*AutopayConfigData, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AutopayConfig) SetConfig(v AutopayConfigData)`

SetConfig sets Config field to given value.


### GetCreationTime

`func (o *AutopayConfig) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AutopayConfig) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AutopayConfig) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *AutopayConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutopayConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutopayConfig) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *AutopayConfig) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AutopayConfig) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AutopayConfig) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetLendingAccountId

`func (o *AutopayConfig) GetLendingAccountId() string`

GetLendingAccountId returns the LendingAccountId field if non-nil, zero value otherwise.

### GetLendingAccountIdOk

`func (o *AutopayConfig) GetLendingAccountIdOk() (*string, bool)`

GetLendingAccountIdOk returns a tuple with the LendingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLendingAccountId

`func (o *AutopayConfig) SetLendingAccountId(v string)`

SetLendingAccountId sets LendingAccountId field to given value.


### GetPersonId

`func (o *AutopayConfig) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *AutopayConfig) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *AutopayConfig) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *AutopayConfig) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetStatus

`func (o *AutopayConfig) GetStatus() AutopayConfigStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutopayConfig) GetStatusOk() (*AutopayConfigStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutopayConfig) SetStatus(v AutopayConfigStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *AutopayConfig) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AutopayConfig) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AutopayConfig) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AutopayConfig) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


