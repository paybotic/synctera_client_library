# CreateThreeDsDecisionGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProducts** | **[]string** | List of Card Product IDs that will use the 3DS decision gateway | 
**CustomHeaders** | Pointer to **map[string]string** | These key-value pairs define custom HTTP headers that will be included in every HTTP request to the gateway. Note that when updating this field, all key-value pairs will be replaced. They are not merged with existing data.  | [optional] 
**DecisionUrl** | **string** | URL of the 3DS decision gateway | 
**FallbackDecision** | [**ThreeDsDecision**](ThreeDsDecision.md) |  | 
**IsActive** | **bool** | The 3DS decision gateway will only be used if this is true. | 

## Methods

### NewCreateThreeDsDecisionGatewayRequest

`func NewCreateThreeDsDecisionGatewayRequest(cardProducts []string, decisionUrl string, fallbackDecision ThreeDsDecision, isActive bool, ) *CreateThreeDsDecisionGatewayRequest`

NewCreateThreeDsDecisionGatewayRequest instantiates a new CreateThreeDsDecisionGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateThreeDsDecisionGatewayRequestWithDefaults

`func NewCreateThreeDsDecisionGatewayRequestWithDefaults() *CreateThreeDsDecisionGatewayRequest`

NewCreateThreeDsDecisionGatewayRequestWithDefaults instantiates a new CreateThreeDsDecisionGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *CreateThreeDsDecisionGatewayRequest) GetCardProducts() []string`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *CreateThreeDsDecisionGatewayRequest) GetCardProductsOk() (*[]string, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *CreateThreeDsDecisionGatewayRequest) SetCardProducts(v []string)`

SetCardProducts sets CardProducts field to given value.


### GetCustomHeaders

`func (o *CreateThreeDsDecisionGatewayRequest) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CreateThreeDsDecisionGatewayRequest) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CreateThreeDsDecisionGatewayRequest) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *CreateThreeDsDecisionGatewayRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetDecisionUrl

`func (o *CreateThreeDsDecisionGatewayRequest) GetDecisionUrl() string`

GetDecisionUrl returns the DecisionUrl field if non-nil, zero value otherwise.

### GetDecisionUrlOk

`func (o *CreateThreeDsDecisionGatewayRequest) GetDecisionUrlOk() (*string, bool)`

GetDecisionUrlOk returns a tuple with the DecisionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionUrl

`func (o *CreateThreeDsDecisionGatewayRequest) SetDecisionUrl(v string)`

SetDecisionUrl sets DecisionUrl field to given value.


### GetFallbackDecision

`func (o *CreateThreeDsDecisionGatewayRequest) GetFallbackDecision() ThreeDsDecision`

GetFallbackDecision returns the FallbackDecision field if non-nil, zero value otherwise.

### GetFallbackDecisionOk

`func (o *CreateThreeDsDecisionGatewayRequest) GetFallbackDecisionOk() (*ThreeDsDecision, bool)`

GetFallbackDecisionOk returns a tuple with the FallbackDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackDecision

`func (o *CreateThreeDsDecisionGatewayRequest) SetFallbackDecision(v ThreeDsDecision)`

SetFallbackDecision sets FallbackDecision field to given value.


### GetIsActive

`func (o *CreateThreeDsDecisionGatewayRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateThreeDsDecisionGatewayRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateThreeDsDecisionGatewayRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


