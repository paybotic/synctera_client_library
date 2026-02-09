# FraudRiskData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decision** | **string** | The overall decision, based on the vendor response | 
**Reasons** | Pointer to **[]string** | Set of machine-readable descriptions of why the transaction was accepted or declined | [optional] 
**Vendor** | **string** | The external vendor used for risk evaluation | 
**VendorReasons** | Pointer to **[]string** | List of reasons supplied by the provider | [optional] 

## Methods

### NewFraudRiskData

`func NewFraudRiskData(decision string, vendor string, ) *FraudRiskData`

NewFraudRiskData instantiates a new FraudRiskData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFraudRiskDataWithDefaults

`func NewFraudRiskDataWithDefaults() *FraudRiskData`

NewFraudRiskDataWithDefaults instantiates a new FraudRiskData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecision

`func (o *FraudRiskData) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *FraudRiskData) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *FraudRiskData) SetDecision(v string)`

SetDecision sets Decision field to given value.


### GetReasons

`func (o *FraudRiskData) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *FraudRiskData) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *FraudRiskData) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *FraudRiskData) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetVendor

`func (o *FraudRiskData) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *FraudRiskData) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *FraudRiskData) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetVendorReasons

`func (o *FraudRiskData) GetVendorReasons() []string`

GetVendorReasons returns the VendorReasons field if non-nil, zero value otherwise.

### GetVendorReasonsOk

`func (o *FraudRiskData) GetVendorReasonsOk() (*[]string, bool)`

GetVendorReasonsOk returns a tuple with the VendorReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorReasons

`func (o *FraudRiskData) SetVendorReasons(v []string)`

SetVendorReasons sets VendorReasons field to given value.

### HasVendorReasons

`func (o *FraudRiskData) HasVendorReasons() bool`

HasVendorReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


