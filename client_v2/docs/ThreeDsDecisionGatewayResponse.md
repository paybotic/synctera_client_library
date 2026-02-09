# ThreeDsDecisionGatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProducts** | **[]string** | List of Card Product IDs that will use the 3DS decision gateway | 
**CreationTime** | **time.Time** | The timestamp when the 3DS decision gateway was created | [readonly] 
**CustomHeaders** | Pointer to **map[string]string** | These key-value pairs define custom HTTP headers that will be included in every HTTP request to the gateway. Note that when updating this field, all key-value pairs will be replaced. They are not merged with existing data.  | [optional] 
**DecisionUrl** | **string** | URL of the 3DS decision gateway | 
**FallbackDecision** | [**ThreeDsDecision**](ThreeDsDecision.md) |  | 
**Id** | **string** | The unique identifier of an 3DS decision gateway | 
**IsActive** | **bool** | The 3DS decision gateway will only be used if this is true. | 
**LastUpdatedTime** | **time.Time** | The timestamp when the 3DS decision gateway was last modified | [readonly] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewThreeDsDecisionGatewayResponse

`func NewThreeDsDecisionGatewayResponse(cardProducts []string, creationTime time.Time, decisionUrl string, fallbackDecision ThreeDsDecision, id string, isActive bool, lastUpdatedTime time.Time, tenant string, ) *ThreeDsDecisionGatewayResponse`

NewThreeDsDecisionGatewayResponse instantiates a new ThreeDsDecisionGatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDsDecisionGatewayResponseWithDefaults

`func NewThreeDsDecisionGatewayResponseWithDefaults() *ThreeDsDecisionGatewayResponse`

NewThreeDsDecisionGatewayResponseWithDefaults instantiates a new ThreeDsDecisionGatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *ThreeDsDecisionGatewayResponse) GetCardProducts() []string`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *ThreeDsDecisionGatewayResponse) GetCardProductsOk() (*[]string, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *ThreeDsDecisionGatewayResponse) SetCardProducts(v []string)`

SetCardProducts sets CardProducts field to given value.


### GetCreationTime

`func (o *ThreeDsDecisionGatewayResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ThreeDsDecisionGatewayResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ThreeDsDecisionGatewayResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomHeaders

`func (o *ThreeDsDecisionGatewayResponse) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *ThreeDsDecisionGatewayResponse) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *ThreeDsDecisionGatewayResponse) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *ThreeDsDecisionGatewayResponse) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetDecisionUrl

`func (o *ThreeDsDecisionGatewayResponse) GetDecisionUrl() string`

GetDecisionUrl returns the DecisionUrl field if non-nil, zero value otherwise.

### GetDecisionUrlOk

`func (o *ThreeDsDecisionGatewayResponse) GetDecisionUrlOk() (*string, bool)`

GetDecisionUrlOk returns a tuple with the DecisionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionUrl

`func (o *ThreeDsDecisionGatewayResponse) SetDecisionUrl(v string)`

SetDecisionUrl sets DecisionUrl field to given value.


### GetFallbackDecision

`func (o *ThreeDsDecisionGatewayResponse) GetFallbackDecision() ThreeDsDecision`

GetFallbackDecision returns the FallbackDecision field if non-nil, zero value otherwise.

### GetFallbackDecisionOk

`func (o *ThreeDsDecisionGatewayResponse) GetFallbackDecisionOk() (*ThreeDsDecision, bool)`

GetFallbackDecisionOk returns a tuple with the FallbackDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackDecision

`func (o *ThreeDsDecisionGatewayResponse) SetFallbackDecision(v ThreeDsDecision)`

SetFallbackDecision sets FallbackDecision field to given value.


### GetId

`func (o *ThreeDsDecisionGatewayResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreeDsDecisionGatewayResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreeDsDecisionGatewayResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *ThreeDsDecisionGatewayResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ThreeDsDecisionGatewayResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ThreeDsDecisionGatewayResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastUpdatedTime

`func (o *ThreeDsDecisionGatewayResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ThreeDsDecisionGatewayResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ThreeDsDecisionGatewayResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetTenant

`func (o *ThreeDsDecisionGatewayResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ThreeDsDecisionGatewayResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ThreeDsDecisionGatewayResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


