# PersonalIdConfigurationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicEncryptionKey** | **string** | A public encryption key used for encrypting personal identifiers before transport. | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 

## Methods

### NewPersonalIdConfigurationCreate

`func NewPersonalIdConfigurationCreate(publicEncryptionKey string, ) *PersonalIdConfigurationCreate`

NewPersonalIdConfigurationCreate instantiates a new PersonalIdConfigurationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalIdConfigurationCreateWithDefaults

`func NewPersonalIdConfigurationCreateWithDefaults() *PersonalIdConfigurationCreate`

NewPersonalIdConfigurationCreateWithDefaults instantiates a new PersonalIdConfigurationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicEncryptionKey

`func (o *PersonalIdConfigurationCreate) GetPublicEncryptionKey() string`

GetPublicEncryptionKey returns the PublicEncryptionKey field if non-nil, zero value otherwise.

### GetPublicEncryptionKeyOk

`func (o *PersonalIdConfigurationCreate) GetPublicEncryptionKeyOk() (*string, bool)`

GetPublicEncryptionKeyOk returns a tuple with the PublicEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEncryptionKey

`func (o *PersonalIdConfigurationCreate) SetPublicEncryptionKey(v string)`

SetPublicEncryptionKey sets PublicEncryptionKey field to given value.


### GetTenant

`func (o *PersonalIdConfigurationCreate) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PersonalIdConfigurationCreate) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PersonalIdConfigurationCreate) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PersonalIdConfigurationCreate) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


