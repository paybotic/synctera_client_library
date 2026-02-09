# CreateGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProducts** | **[]string** | List of Card Product unique identifiers that will utilize the Gateway | 
**CustomHeaders** | Pointer to **map[string]string** | These key-value pairs define custom HTTP headers that will be included in every HTTP request to the gateway. Note that when updating this field, all key-value pairs will be replaced. They are not merged with existing data.  | [optional] 
**IsActive** | **bool** | Current status of the Authorization gateway | 
**Standin** | Pointer to [**GatewayStandin**](GatewayStandin.md) |  | [optional] 
**Url** | **string** | URL of the Authorization gateway | 

## Methods

### NewCreateGatewayRequest

`func NewCreateGatewayRequest(cardProducts []string, isActive bool, url string, ) *CreateGatewayRequest`

NewCreateGatewayRequest instantiates a new CreateGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGatewayRequestWithDefaults

`func NewCreateGatewayRequestWithDefaults() *CreateGatewayRequest`

NewCreateGatewayRequestWithDefaults instantiates a new CreateGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *CreateGatewayRequest) GetCardProducts() []string`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *CreateGatewayRequest) GetCardProductsOk() (*[]string, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *CreateGatewayRequest) SetCardProducts(v []string)`

SetCardProducts sets CardProducts field to given value.


### GetCustomHeaders

`func (o *CreateGatewayRequest) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CreateGatewayRequest) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CreateGatewayRequest) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *CreateGatewayRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetIsActive

`func (o *CreateGatewayRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateGatewayRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateGatewayRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetStandin

`func (o *CreateGatewayRequest) GetStandin() GatewayStandin`

GetStandin returns the Standin field if non-nil, zero value otherwise.

### GetStandinOk

`func (o *CreateGatewayRequest) GetStandinOk() (*GatewayStandin, bool)`

GetStandinOk returns a tuple with the Standin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandin

`func (o *CreateGatewayRequest) SetStandin(v GatewayStandin)`

SetStandin sets Standin field to given value.

### HasStandin

`func (o *CreateGatewayRequest) HasStandin() bool`

HasStandin returns a boolean if a field has been set.

### GetUrl

`func (o *CreateGatewayRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateGatewayRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateGatewayRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


