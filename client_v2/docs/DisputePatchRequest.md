# DisputePatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkEligibilityOverride** | Pointer to **bool** | Override network eligibility restrictions and forces the dispute to be filed with the network. | [optional] 
**Status** | Pointer to [**DisputeStatus**](DisputeStatus.md) |  | [optional] 

## Methods

### NewDisputePatchRequest

`func NewDisputePatchRequest() *DisputePatchRequest`

NewDisputePatchRequest instantiates a new DisputePatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputePatchRequestWithDefaults

`func NewDisputePatchRequestWithDefaults() *DisputePatchRequest`

NewDisputePatchRequestWithDefaults instantiates a new DisputePatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkEligibilityOverride

`func (o *DisputePatchRequest) GetNetworkEligibilityOverride() bool`

GetNetworkEligibilityOverride returns the NetworkEligibilityOverride field if non-nil, zero value otherwise.

### GetNetworkEligibilityOverrideOk

`func (o *DisputePatchRequest) GetNetworkEligibilityOverrideOk() (*bool, bool)`

GetNetworkEligibilityOverrideOk returns a tuple with the NetworkEligibilityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEligibilityOverride

`func (o *DisputePatchRequest) SetNetworkEligibilityOverride(v bool)`

SetNetworkEligibilityOverride sets NetworkEligibilityOverride field to given value.

### HasNetworkEligibilityOverride

`func (o *DisputePatchRequest) HasNetworkEligibilityOverride() bool`

HasNetworkEligibilityOverride returns a boolean if a field has been set.

### GetStatus

`func (o *DisputePatchRequest) GetStatus() DisputeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisputePatchRequest) GetStatusOk() (*DisputeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisputePatchRequest) SetStatus(v DisputeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DisputePatchRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


