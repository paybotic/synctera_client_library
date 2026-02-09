# GatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProducts** | **[]string** | List of Card Product unique identifiers that will utilize the Gateway | 
**CreationTime** | **time.Time** | The timestamp representing when the gateway config request was made | [readonly] 
**CustomHeaders** | Pointer to **map[string]string** | These key-value pairs define custom HTTP headers that will be included in every HTTP request to the gateway. Note that when updating this field, all key-value pairs will be replaced. They are not merged with existing data.  | [optional] 
**Id** | **string** | Gateway ID | 
**IsActive** | **bool** | Current status of the Authorization gateway | 
**LastUpdatedTime** | **time.Time** | The timestamp representing when the gateway config was last modified at | [readonly] 
**Standin** | Pointer to [**GatewayStandin**](GatewayStandin.md) |  | [optional] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**Url** | **string** | URL of the Authorization gateway | 

## Methods

### NewGatewayResponse

`func NewGatewayResponse(cardProducts []string, creationTime time.Time, id string, isActive bool, lastUpdatedTime time.Time, tenant string, url string, ) *GatewayResponse`

NewGatewayResponse instantiates a new GatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayResponseWithDefaults

`func NewGatewayResponseWithDefaults() *GatewayResponse`

NewGatewayResponseWithDefaults instantiates a new GatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *GatewayResponse) GetCardProducts() []string`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *GatewayResponse) GetCardProductsOk() (*[]string, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *GatewayResponse) SetCardProducts(v []string)`

SetCardProducts sets CardProducts field to given value.


### GetCreationTime

`func (o *GatewayResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GatewayResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GatewayResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomHeaders

`func (o *GatewayResponse) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *GatewayResponse) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *GatewayResponse) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *GatewayResponse) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetId

`func (o *GatewayResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *GatewayResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GatewayResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GatewayResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastUpdatedTime

`func (o *GatewayResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *GatewayResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *GatewayResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetStandin

`func (o *GatewayResponse) GetStandin() GatewayStandin`

GetStandin returns the Standin field if non-nil, zero value otherwise.

### GetStandinOk

`func (o *GatewayResponse) GetStandinOk() (*GatewayStandin, bool)`

GetStandinOk returns a tuple with the Standin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandin

`func (o *GatewayResponse) SetStandin(v GatewayStandin)`

SetStandin sets Standin field to given value.

### HasStandin

`func (o *GatewayResponse) HasStandin() bool`

HasStandin returns a boolean if a field has been set.

### GetTenant

`func (o *GatewayResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *GatewayResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *GatewayResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetUrl

`func (o *GatewayResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GatewayResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GatewayResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


