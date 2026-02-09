# ExternalCardVerificationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressVerificationResult** | Pointer to **string** | Address verification results  Status | Description --- | --- VERIFIED | AVS verified NOT_VERIFIED | AVS not verified ADDRESS_MISMATCH | ZIP code match, address no match ZIP_MISMATCH | Address match, ZIP code no match ADDRESS_AND_ZIP_MISMATCH | Address and ZIP code no match  | [optional] 
**Cvv2Result** | Pointer to **string** | Card Verification Value results  Status | Description --- | --- VERIFIED | CVV and expiration date verified NOT_VERIFIED | CVV and expiration date not verified CVV_MISMATCH | Either CVV or expiration date does not match NOT_SUPPORTED | Issuer does not participate in CVV2 service  | [optional] 
**NameVerificationResult** | Pointer to **string** | Issuer cardholder name verification result with Account Name Inquiry (ANI) service The result of verifying the cardholder name against the name on file at the issuing institution. If this fails, it means the issuing institution has a different person&#39;s name on file as the cardholder.  Status | Description --- | --- VERIFIED | ANI Name verified NOT_VERIFIED | ANI Name not verified NOT_SUPPORTED | Issuer does not participate in ANI service NAME_MISMATCH | ANI Name does not match  | [optional] 
**PullDetails** | Pointer to [**PullDetails**](PullDetails.md) |  | [optional] 
**PullEnabled** | Pointer to **bool** | Indicates if the card is able to perform PULL transfers. | [optional] 
**PushDetails** | Pointer to [**PushDetails**](PushDetails.md) |  | [optional] 
**PushEnabled** | Pointer to **bool** | Indicates if the card is able to perform PUSH transfers. | [optional] 

## Methods

### NewExternalCardVerificationDetails

`func NewExternalCardVerificationDetails() *ExternalCardVerificationDetails`

NewExternalCardVerificationDetails instantiates a new ExternalCardVerificationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardVerificationDetailsWithDefaults

`func NewExternalCardVerificationDetailsWithDefaults() *ExternalCardVerificationDetails`

NewExternalCardVerificationDetailsWithDefaults instantiates a new ExternalCardVerificationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressVerificationResult

`func (o *ExternalCardVerificationDetails) GetAddressVerificationResult() string`

GetAddressVerificationResult returns the AddressVerificationResult field if non-nil, zero value otherwise.

### GetAddressVerificationResultOk

`func (o *ExternalCardVerificationDetails) GetAddressVerificationResultOk() (*string, bool)`

GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerificationResult

`func (o *ExternalCardVerificationDetails) SetAddressVerificationResult(v string)`

SetAddressVerificationResult sets AddressVerificationResult field to given value.

### HasAddressVerificationResult

`func (o *ExternalCardVerificationDetails) HasAddressVerificationResult() bool`

HasAddressVerificationResult returns a boolean if a field has been set.

### GetCvv2Result

`func (o *ExternalCardVerificationDetails) GetCvv2Result() string`

GetCvv2Result returns the Cvv2Result field if non-nil, zero value otherwise.

### GetCvv2ResultOk

`func (o *ExternalCardVerificationDetails) GetCvv2ResultOk() (*string, bool)`

GetCvv2ResultOk returns a tuple with the Cvv2Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv2Result

`func (o *ExternalCardVerificationDetails) SetCvv2Result(v string)`

SetCvv2Result sets Cvv2Result field to given value.

### HasCvv2Result

`func (o *ExternalCardVerificationDetails) HasCvv2Result() bool`

HasCvv2Result returns a boolean if a field has been set.

### GetNameVerificationResult

`func (o *ExternalCardVerificationDetails) GetNameVerificationResult() string`

GetNameVerificationResult returns the NameVerificationResult field if non-nil, zero value otherwise.

### GetNameVerificationResultOk

`func (o *ExternalCardVerificationDetails) GetNameVerificationResultOk() (*string, bool)`

GetNameVerificationResultOk returns a tuple with the NameVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameVerificationResult

`func (o *ExternalCardVerificationDetails) SetNameVerificationResult(v string)`

SetNameVerificationResult sets NameVerificationResult field to given value.

### HasNameVerificationResult

`func (o *ExternalCardVerificationDetails) HasNameVerificationResult() bool`

HasNameVerificationResult returns a boolean if a field has been set.

### GetPullDetails

`func (o *ExternalCardVerificationDetails) GetPullDetails() PullDetails`

GetPullDetails returns the PullDetails field if non-nil, zero value otherwise.

### GetPullDetailsOk

`func (o *ExternalCardVerificationDetails) GetPullDetailsOk() (*PullDetails, bool)`

GetPullDetailsOk returns a tuple with the PullDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullDetails

`func (o *ExternalCardVerificationDetails) SetPullDetails(v PullDetails)`

SetPullDetails sets PullDetails field to given value.

### HasPullDetails

`func (o *ExternalCardVerificationDetails) HasPullDetails() bool`

HasPullDetails returns a boolean if a field has been set.

### GetPullEnabled

`func (o *ExternalCardVerificationDetails) GetPullEnabled() bool`

GetPullEnabled returns the PullEnabled field if non-nil, zero value otherwise.

### GetPullEnabledOk

`func (o *ExternalCardVerificationDetails) GetPullEnabledOk() (*bool, bool)`

GetPullEnabledOk returns a tuple with the PullEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullEnabled

`func (o *ExternalCardVerificationDetails) SetPullEnabled(v bool)`

SetPullEnabled sets PullEnabled field to given value.

### HasPullEnabled

`func (o *ExternalCardVerificationDetails) HasPullEnabled() bool`

HasPullEnabled returns a boolean if a field has been set.

### GetPushDetails

`func (o *ExternalCardVerificationDetails) GetPushDetails() PushDetails`

GetPushDetails returns the PushDetails field if non-nil, zero value otherwise.

### GetPushDetailsOk

`func (o *ExternalCardVerificationDetails) GetPushDetailsOk() (*PushDetails, bool)`

GetPushDetailsOk returns a tuple with the PushDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushDetails

`func (o *ExternalCardVerificationDetails) SetPushDetails(v PushDetails)`

SetPushDetails sets PushDetails field to given value.

### HasPushDetails

`func (o *ExternalCardVerificationDetails) HasPushDetails() bool`

HasPushDetails returns a boolean if a field has been set.

### GetPushEnabled

`func (o *ExternalCardVerificationDetails) GetPushEnabled() bool`

GetPushEnabled returns the PushEnabled field if non-nil, zero value otherwise.

### GetPushEnabledOk

`func (o *ExternalCardVerificationDetails) GetPushEnabledOk() (*bool, bool)`

GetPushEnabledOk returns a tuple with the PushEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushEnabled

`func (o *ExternalCardVerificationDetails) SetPushEnabled(v bool)`

SetPushEnabled sets PushEnabled field to given value.

### HasPushEnabled

`func (o *ExternalCardVerificationDetails) HasPushEnabled() bool`

HasPushEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


