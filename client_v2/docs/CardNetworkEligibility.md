# CardNetworkEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDigitalWalletToken** | **bool** | Transaction was processed using a digital wallet token. | 
**IsNetworkEligibilityOverridden** | **bool** | Network eligibility restrictions have been overridden. | 
**IsThreeDs** | **bool** | Transaction was authenticated with 3DS. | 

## Methods

### NewCardNetworkEligibility

`func NewCardNetworkEligibility(isDigitalWalletToken bool, isNetworkEligibilityOverridden bool, isThreeDs bool, ) *CardNetworkEligibility`

NewCardNetworkEligibility instantiates a new CardNetworkEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardNetworkEligibilityWithDefaults

`func NewCardNetworkEligibilityWithDefaults() *CardNetworkEligibility`

NewCardNetworkEligibilityWithDefaults instantiates a new CardNetworkEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDigitalWalletToken

`func (o *CardNetworkEligibility) GetIsDigitalWalletToken() bool`

GetIsDigitalWalletToken returns the IsDigitalWalletToken field if non-nil, zero value otherwise.

### GetIsDigitalWalletTokenOk

`func (o *CardNetworkEligibility) GetIsDigitalWalletTokenOk() (*bool, bool)`

GetIsDigitalWalletTokenOk returns a tuple with the IsDigitalWalletToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDigitalWalletToken

`func (o *CardNetworkEligibility) SetIsDigitalWalletToken(v bool)`

SetIsDigitalWalletToken sets IsDigitalWalletToken field to given value.


### GetIsNetworkEligibilityOverridden

`func (o *CardNetworkEligibility) GetIsNetworkEligibilityOverridden() bool`

GetIsNetworkEligibilityOverridden returns the IsNetworkEligibilityOverridden field if non-nil, zero value otherwise.

### GetIsNetworkEligibilityOverriddenOk

`func (o *CardNetworkEligibility) GetIsNetworkEligibilityOverriddenOk() (*bool, bool)`

GetIsNetworkEligibilityOverriddenOk returns a tuple with the IsNetworkEligibilityOverridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNetworkEligibilityOverridden

`func (o *CardNetworkEligibility) SetIsNetworkEligibilityOverridden(v bool)`

SetIsNetworkEligibilityOverridden sets IsNetworkEligibilityOverridden field to given value.


### GetIsThreeDs

`func (o *CardNetworkEligibility) GetIsThreeDs() bool`

GetIsThreeDs returns the IsThreeDs field if non-nil, zero value otherwise.

### GetIsThreeDsOk

`func (o *CardNetworkEligibility) GetIsThreeDsOk() (*bool, bool)`

GetIsThreeDsOk returns a tuple with the IsThreeDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsThreeDs

`func (o *CardNetworkEligibility) SetIsThreeDs(v bool)`

SetIsThreeDs sets IsThreeDs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


