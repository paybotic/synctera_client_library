# ExternalCardVerifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressVerificationResult** | **string** | Address verification results  Status | Description --- | --- VERIFIED | AVS verified NOT_VERIFIED | AVS not verified ADDRESS_MISMATCH | ZIP code match, address no match ZIP_MISMATCH | Address match, ZIP code no match ADDRESS_AND_ZIP_MISMATCH | Address and ZIP code no match  | 
**Cvv2Result** | **string** | Card Verification Value results  Status | Description --- | --- VERIFIED | CVV and expiration date verified NOT_VERIFIED | CVV and expiration date not verified CVV_MISMATCH | Either CVV or expiration date does not match NOT_SUPPORTED | Issuer does not participate in CVV2 service  | 
**NameVerificationResult** | **string** | Issuer cardholder name verification result with Account Name Inquiry (ANI) service The result of verifying the cardholder name against the name on file at the issuing institution. If this fails, it means the issuing institution has a different person&#39;s name on file as the cardholder.  Status | Description --- | --- VERIFIED | ANI Name verified NOT_VERIFIED | ANI Name not verified NOT_SUPPORTED | Issuer does not participate in ANI service NAME_MISMATCH | ANI Name does not match  | 
**PullDetails** | Pointer to [**PullDetails**](PullDetails.md) |  | [optional] 
**PullEnabled** | **bool** | Indicates if the card is able to perform PULL transfers. | 
**PushDetails** | Pointer to [**PushDetails**](PushDetails.md) |  | [optional] 
**PushEnabled** | **bool** | Indicates if the card is able to perform PUSH transfers. | 
**State** | **string** |  | 

## Methods

### NewExternalCardVerifications

`func NewExternalCardVerifications(addressVerificationResult string, cvv2Result string, nameVerificationResult string, pullEnabled bool, pushEnabled bool, state string, ) *ExternalCardVerifications`

NewExternalCardVerifications instantiates a new ExternalCardVerifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardVerificationsWithDefaults

`func NewExternalCardVerificationsWithDefaults() *ExternalCardVerifications`

NewExternalCardVerificationsWithDefaults instantiates a new ExternalCardVerifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressVerificationResult

`func (o *ExternalCardVerifications) GetAddressVerificationResult() string`

GetAddressVerificationResult returns the AddressVerificationResult field if non-nil, zero value otherwise.

### GetAddressVerificationResultOk

`func (o *ExternalCardVerifications) GetAddressVerificationResultOk() (*string, bool)`

GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerificationResult

`func (o *ExternalCardVerifications) SetAddressVerificationResult(v string)`

SetAddressVerificationResult sets AddressVerificationResult field to given value.


### GetCvv2Result

`func (o *ExternalCardVerifications) GetCvv2Result() string`

GetCvv2Result returns the Cvv2Result field if non-nil, zero value otherwise.

### GetCvv2ResultOk

`func (o *ExternalCardVerifications) GetCvv2ResultOk() (*string, bool)`

GetCvv2ResultOk returns a tuple with the Cvv2Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv2Result

`func (o *ExternalCardVerifications) SetCvv2Result(v string)`

SetCvv2Result sets Cvv2Result field to given value.


### GetNameVerificationResult

`func (o *ExternalCardVerifications) GetNameVerificationResult() string`

GetNameVerificationResult returns the NameVerificationResult field if non-nil, zero value otherwise.

### GetNameVerificationResultOk

`func (o *ExternalCardVerifications) GetNameVerificationResultOk() (*string, bool)`

GetNameVerificationResultOk returns a tuple with the NameVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameVerificationResult

`func (o *ExternalCardVerifications) SetNameVerificationResult(v string)`

SetNameVerificationResult sets NameVerificationResult field to given value.


### GetPullDetails

`func (o *ExternalCardVerifications) GetPullDetails() PullDetails`

GetPullDetails returns the PullDetails field if non-nil, zero value otherwise.

### GetPullDetailsOk

`func (o *ExternalCardVerifications) GetPullDetailsOk() (*PullDetails, bool)`

GetPullDetailsOk returns a tuple with the PullDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullDetails

`func (o *ExternalCardVerifications) SetPullDetails(v PullDetails)`

SetPullDetails sets PullDetails field to given value.

### HasPullDetails

`func (o *ExternalCardVerifications) HasPullDetails() bool`

HasPullDetails returns a boolean if a field has been set.

### GetPullEnabled

`func (o *ExternalCardVerifications) GetPullEnabled() bool`

GetPullEnabled returns the PullEnabled field if non-nil, zero value otherwise.

### GetPullEnabledOk

`func (o *ExternalCardVerifications) GetPullEnabledOk() (*bool, bool)`

GetPullEnabledOk returns a tuple with the PullEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullEnabled

`func (o *ExternalCardVerifications) SetPullEnabled(v bool)`

SetPullEnabled sets PullEnabled field to given value.


### GetPushDetails

`func (o *ExternalCardVerifications) GetPushDetails() PushDetails`

GetPushDetails returns the PushDetails field if non-nil, zero value otherwise.

### GetPushDetailsOk

`func (o *ExternalCardVerifications) GetPushDetailsOk() (*PushDetails, bool)`

GetPushDetailsOk returns a tuple with the PushDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushDetails

`func (o *ExternalCardVerifications) SetPushDetails(v PushDetails)`

SetPushDetails sets PushDetails field to given value.

### HasPushDetails

`func (o *ExternalCardVerifications) HasPushDetails() bool`

HasPushDetails returns a boolean if a field has been set.

### GetPushEnabled

`func (o *ExternalCardVerifications) GetPushEnabled() bool`

GetPushEnabled returns the PushEnabled field if non-nil, zero value otherwise.

### GetPushEnabledOk

`func (o *ExternalCardVerifications) GetPushEnabledOk() (*bool, bool)`

GetPushEnabledOk returns a tuple with the PushEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushEnabled

`func (o *ExternalCardVerifications) SetPushEnabled(v bool)`

SetPushEnabled sets PushEnabled field to given value.


### GetState

`func (o *ExternalCardVerifications) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExternalCardVerifications) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExternalCardVerifications) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


