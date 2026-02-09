# UpdateThreeDsDecisionGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProducts** | Pointer to **[]string** | List of card product unique IDs that will use the 3DS decision gateway. Note that the list of card product IDs from this request does not merge with the 3DS gateway&#39;s previous list of card product IDs. The list of card product IDs for the gateway will be fully replaced by the list from the update request, if any such list is included in the request.  | [optional] 
**CustomHeaders** | Pointer to **map[string]string** | These key-value pairs define custom HTTP headers that will be included in every HTTP request to the gateway. Note that when updating this field, all key-value pairs will be replaced. They are not merged with existing data.  | [optional] 
**DecisionUrl** | Pointer to **string** | URL of the 3DS decision gateway | [optional] 
**FallbackDecision** | Pointer to [**ThreeDsDecision**](ThreeDsDecision.md) |  | [optional] 
**IsActive** | Pointer to **bool** | The 3DS decision gateway will only be used if this is true. | [optional] 

## Methods

### NewUpdateThreeDsDecisionGatewayRequest

`func NewUpdateThreeDsDecisionGatewayRequest() *UpdateThreeDsDecisionGatewayRequest`

NewUpdateThreeDsDecisionGatewayRequest instantiates a new UpdateThreeDsDecisionGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateThreeDsDecisionGatewayRequestWithDefaults

`func NewUpdateThreeDsDecisionGatewayRequestWithDefaults() *UpdateThreeDsDecisionGatewayRequest`

NewUpdateThreeDsDecisionGatewayRequestWithDefaults instantiates a new UpdateThreeDsDecisionGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *UpdateThreeDsDecisionGatewayRequest) GetCardProducts() []string`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *UpdateThreeDsDecisionGatewayRequest) GetCardProductsOk() (*[]string, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *UpdateThreeDsDecisionGatewayRequest) SetCardProducts(v []string)`

SetCardProducts sets CardProducts field to given value.

### HasCardProducts

`func (o *UpdateThreeDsDecisionGatewayRequest) HasCardProducts() bool`

HasCardProducts returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *UpdateThreeDsDecisionGatewayRequest) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *UpdateThreeDsDecisionGatewayRequest) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *UpdateThreeDsDecisionGatewayRequest) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *UpdateThreeDsDecisionGatewayRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetDecisionUrl

`func (o *UpdateThreeDsDecisionGatewayRequest) GetDecisionUrl() string`

GetDecisionUrl returns the DecisionUrl field if non-nil, zero value otherwise.

### GetDecisionUrlOk

`func (o *UpdateThreeDsDecisionGatewayRequest) GetDecisionUrlOk() (*string, bool)`

GetDecisionUrlOk returns a tuple with the DecisionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionUrl

`func (o *UpdateThreeDsDecisionGatewayRequest) SetDecisionUrl(v string)`

SetDecisionUrl sets DecisionUrl field to given value.

### HasDecisionUrl

`func (o *UpdateThreeDsDecisionGatewayRequest) HasDecisionUrl() bool`

HasDecisionUrl returns a boolean if a field has been set.

### GetFallbackDecision

`func (o *UpdateThreeDsDecisionGatewayRequest) GetFallbackDecision() ThreeDsDecision`

GetFallbackDecision returns the FallbackDecision field if non-nil, zero value otherwise.

### GetFallbackDecisionOk

`func (o *UpdateThreeDsDecisionGatewayRequest) GetFallbackDecisionOk() (*ThreeDsDecision, bool)`

GetFallbackDecisionOk returns a tuple with the FallbackDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackDecision

`func (o *UpdateThreeDsDecisionGatewayRequest) SetFallbackDecision(v ThreeDsDecision)`

SetFallbackDecision sets FallbackDecision field to given value.

### HasFallbackDecision

`func (o *UpdateThreeDsDecisionGatewayRequest) HasFallbackDecision() bool`

HasFallbackDecision returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateThreeDsDecisionGatewayRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateThreeDsDecisionGatewayRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateThreeDsDecisionGatewayRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateThreeDsDecisionGatewayRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


