# ThreeDsDecisionGatewayList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Gateways** | [**[]ThreeDsDecisionGatewayResponse**](ThreeDsDecisionGatewayResponse.md) | Array of 3DS decision gateways | 

## Methods

### NewThreeDsDecisionGatewayList

`func NewThreeDsDecisionGatewayList(gateways []ThreeDsDecisionGatewayResponse, ) *ThreeDsDecisionGatewayList`

NewThreeDsDecisionGatewayList instantiates a new ThreeDsDecisionGatewayList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDsDecisionGatewayListWithDefaults

`func NewThreeDsDecisionGatewayListWithDefaults() *ThreeDsDecisionGatewayList`

NewThreeDsDecisionGatewayListWithDefaults instantiates a new ThreeDsDecisionGatewayList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ThreeDsDecisionGatewayList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ThreeDsDecisionGatewayList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ThreeDsDecisionGatewayList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ThreeDsDecisionGatewayList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetGateways

`func (o *ThreeDsDecisionGatewayList) GetGateways() []ThreeDsDecisionGatewayResponse`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *ThreeDsDecisionGatewayList) GetGatewaysOk() (*[]ThreeDsDecisionGatewayResponse, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *ThreeDsDecisionGatewayList) SetGateways(v []ThreeDsDecisionGatewayResponse)`

SetGateways sets Gateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


