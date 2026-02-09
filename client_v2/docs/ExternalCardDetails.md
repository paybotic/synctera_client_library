# ExternalCardDetails

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
**Bin** | Pointer to **string** | Bank Identification Number | [optional] 
**Issuer** | Pointer to **string** | Name of the issuing financial institution | [optional] 
**LastFour** | Pointer to **string** | The last four digits of the card PAN | [optional] 
**PaymentAccountReference** | Pointer to **string** | A unique identifier associated with a specific cardholder PAN and its affiliated tokens | [optional] 

## Methods

### NewExternalCardDetails

`func NewExternalCardDetails() *ExternalCardDetails`

NewExternalCardDetails instantiates a new ExternalCardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardDetailsWithDefaults

`func NewExternalCardDetailsWithDefaults() *ExternalCardDetails`

NewExternalCardDetailsWithDefaults instantiates a new ExternalCardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressVerificationResult

`func (o *ExternalCardDetails) GetAddressVerificationResult() string`

GetAddressVerificationResult returns the AddressVerificationResult field if non-nil, zero value otherwise.

### GetAddressVerificationResultOk

`func (o *ExternalCardDetails) GetAddressVerificationResultOk() (*string, bool)`

GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerificationResult

`func (o *ExternalCardDetails) SetAddressVerificationResult(v string)`

SetAddressVerificationResult sets AddressVerificationResult field to given value.

### HasAddressVerificationResult

`func (o *ExternalCardDetails) HasAddressVerificationResult() bool`

HasAddressVerificationResult returns a boolean if a field has been set.

### GetCvv2Result

`func (o *ExternalCardDetails) GetCvv2Result() string`

GetCvv2Result returns the Cvv2Result field if non-nil, zero value otherwise.

### GetCvv2ResultOk

`func (o *ExternalCardDetails) GetCvv2ResultOk() (*string, bool)`

GetCvv2ResultOk returns a tuple with the Cvv2Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv2Result

`func (o *ExternalCardDetails) SetCvv2Result(v string)`

SetCvv2Result sets Cvv2Result field to given value.

### HasCvv2Result

`func (o *ExternalCardDetails) HasCvv2Result() bool`

HasCvv2Result returns a boolean if a field has been set.

### GetNameVerificationResult

`func (o *ExternalCardDetails) GetNameVerificationResult() string`

GetNameVerificationResult returns the NameVerificationResult field if non-nil, zero value otherwise.

### GetNameVerificationResultOk

`func (o *ExternalCardDetails) GetNameVerificationResultOk() (*string, bool)`

GetNameVerificationResultOk returns a tuple with the NameVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameVerificationResult

`func (o *ExternalCardDetails) SetNameVerificationResult(v string)`

SetNameVerificationResult sets NameVerificationResult field to given value.

### HasNameVerificationResult

`func (o *ExternalCardDetails) HasNameVerificationResult() bool`

HasNameVerificationResult returns a boolean if a field has been set.

### GetPullDetails

`func (o *ExternalCardDetails) GetPullDetails() PullDetails`

GetPullDetails returns the PullDetails field if non-nil, zero value otherwise.

### GetPullDetailsOk

`func (o *ExternalCardDetails) GetPullDetailsOk() (*PullDetails, bool)`

GetPullDetailsOk returns a tuple with the PullDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullDetails

`func (o *ExternalCardDetails) SetPullDetails(v PullDetails)`

SetPullDetails sets PullDetails field to given value.

### HasPullDetails

`func (o *ExternalCardDetails) HasPullDetails() bool`

HasPullDetails returns a boolean if a field has been set.

### GetPullEnabled

`func (o *ExternalCardDetails) GetPullEnabled() bool`

GetPullEnabled returns the PullEnabled field if non-nil, zero value otherwise.

### GetPullEnabledOk

`func (o *ExternalCardDetails) GetPullEnabledOk() (*bool, bool)`

GetPullEnabledOk returns a tuple with the PullEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullEnabled

`func (o *ExternalCardDetails) SetPullEnabled(v bool)`

SetPullEnabled sets PullEnabled field to given value.

### HasPullEnabled

`func (o *ExternalCardDetails) HasPullEnabled() bool`

HasPullEnabled returns a boolean if a field has been set.

### GetPushDetails

`func (o *ExternalCardDetails) GetPushDetails() PushDetails`

GetPushDetails returns the PushDetails field if non-nil, zero value otherwise.

### GetPushDetailsOk

`func (o *ExternalCardDetails) GetPushDetailsOk() (*PushDetails, bool)`

GetPushDetailsOk returns a tuple with the PushDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushDetails

`func (o *ExternalCardDetails) SetPushDetails(v PushDetails)`

SetPushDetails sets PushDetails field to given value.

### HasPushDetails

`func (o *ExternalCardDetails) HasPushDetails() bool`

HasPushDetails returns a boolean if a field has been set.

### GetPushEnabled

`func (o *ExternalCardDetails) GetPushEnabled() bool`

GetPushEnabled returns the PushEnabled field if non-nil, zero value otherwise.

### GetPushEnabledOk

`func (o *ExternalCardDetails) GetPushEnabledOk() (*bool, bool)`

GetPushEnabledOk returns a tuple with the PushEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushEnabled

`func (o *ExternalCardDetails) SetPushEnabled(v bool)`

SetPushEnabled sets PushEnabled field to given value.

### HasPushEnabled

`func (o *ExternalCardDetails) HasPushEnabled() bool`

HasPushEnabled returns a boolean if a field has been set.

### GetBin

`func (o *ExternalCardDetails) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *ExternalCardDetails) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *ExternalCardDetails) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *ExternalCardDetails) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetIssuer

`func (o *ExternalCardDetails) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ExternalCardDetails) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ExternalCardDetails) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ExternalCardDetails) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLastFour

`func (o *ExternalCardDetails) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *ExternalCardDetails) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *ExternalCardDetails) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *ExternalCardDetails) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetPaymentAccountReference

`func (o *ExternalCardDetails) GetPaymentAccountReference() string`

GetPaymentAccountReference returns the PaymentAccountReference field if non-nil, zero value otherwise.

### GetPaymentAccountReferenceOk

`func (o *ExternalCardDetails) GetPaymentAccountReferenceOk() (*string, bool)`

GetPaymentAccountReferenceOk returns a tuple with the PaymentAccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountReference

`func (o *ExternalCardDetails) SetPaymentAccountReference(v string)`

SetPaymentAccountReference sets PaymentAccountReference field to given value.

### HasPaymentAccountReference

`func (o *ExternalCardDetails) HasPaymentAccountReference() bool`

HasPaymentAccountReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


