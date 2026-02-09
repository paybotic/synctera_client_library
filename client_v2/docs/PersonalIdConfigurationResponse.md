# PersonalIdConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | **time.Time** | The date and time the configuration was created. | [readonly] 
**Id** | **string** | Personal ID configuration ID | [readonly] 
**LastUpdatedTime** | **time.Time** | The date and time the configuration was last updated. | [readonly] 
**PublicEncryptionKey** | **string** | A public encryption key used for encrypting personal identifiers before transport. | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewPersonalIdConfigurationResponse

`func NewPersonalIdConfigurationResponse(creationTime time.Time, id string, lastUpdatedTime time.Time, publicEncryptionKey string, tenant string, ) *PersonalIdConfigurationResponse`

NewPersonalIdConfigurationResponse instantiates a new PersonalIdConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalIdConfigurationResponseWithDefaults

`func NewPersonalIdConfigurationResponseWithDefaults() *PersonalIdConfigurationResponse`

NewPersonalIdConfigurationResponseWithDefaults instantiates a new PersonalIdConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *PersonalIdConfigurationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PersonalIdConfigurationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PersonalIdConfigurationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *PersonalIdConfigurationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonalIdConfigurationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonalIdConfigurationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *PersonalIdConfigurationResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *PersonalIdConfigurationResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *PersonalIdConfigurationResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetPublicEncryptionKey

`func (o *PersonalIdConfigurationResponse) GetPublicEncryptionKey() string`

GetPublicEncryptionKey returns the PublicEncryptionKey field if non-nil, zero value otherwise.

### GetPublicEncryptionKeyOk

`func (o *PersonalIdConfigurationResponse) GetPublicEncryptionKeyOk() (*string, bool)`

GetPublicEncryptionKeyOk returns a tuple with the PublicEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEncryptionKey

`func (o *PersonalIdConfigurationResponse) SetPublicEncryptionKey(v string)`

SetPublicEncryptionKey sets PublicEncryptionKey field to given value.


### GetTenant

`func (o *PersonalIdConfigurationResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PersonalIdConfigurationResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PersonalIdConfigurationResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


