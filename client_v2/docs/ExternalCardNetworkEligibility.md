# ExternalCardNetworkEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDigitalWalletToken** | **bool** | Transaction was processed using a digital wallet token. | 
**IsThreeDs** | **bool** | Transaction was authenticated with 3DS. | 

## Methods

### NewExternalCardNetworkEligibility

`func NewExternalCardNetworkEligibility(isDigitalWalletToken bool, isThreeDs bool, ) *ExternalCardNetworkEligibility`

NewExternalCardNetworkEligibility instantiates a new ExternalCardNetworkEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardNetworkEligibilityWithDefaults

`func NewExternalCardNetworkEligibilityWithDefaults() *ExternalCardNetworkEligibility`

NewExternalCardNetworkEligibilityWithDefaults instantiates a new ExternalCardNetworkEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDigitalWalletToken

`func (o *ExternalCardNetworkEligibility) GetIsDigitalWalletToken() bool`

GetIsDigitalWalletToken returns the IsDigitalWalletToken field if non-nil, zero value otherwise.

### GetIsDigitalWalletTokenOk

`func (o *ExternalCardNetworkEligibility) GetIsDigitalWalletTokenOk() (*bool, bool)`

GetIsDigitalWalletTokenOk returns a tuple with the IsDigitalWalletToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDigitalWalletToken

`func (o *ExternalCardNetworkEligibility) SetIsDigitalWalletToken(v bool)`

SetIsDigitalWalletToken sets IsDigitalWalletToken field to given value.


### GetIsThreeDs

`func (o *ExternalCardNetworkEligibility) GetIsThreeDs() bool`

GetIsThreeDs returns the IsThreeDs field if non-nil, zero value otherwise.

### GetIsThreeDsOk

`func (o *ExternalCardNetworkEligibility) GetIsThreeDsOk() (*bool, bool)`

GetIsThreeDsOk returns a tuple with the IsThreeDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsThreeDs

`func (o *ExternalCardNetworkEligibility) SetIsThreeDs(v bool)`

SetIsThreeDs sets IsThreeDs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


