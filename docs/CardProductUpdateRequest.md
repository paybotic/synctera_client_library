# CardProductUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DigitalWalletTokenization** | Pointer to [**DigitalWalletTokenization**](DigitalWalletTokenization.md) |  | [optional] 
**IssueWithoutKyc** | Pointer to **bool** | Allow issuing cards on this product without requiring KYC (customer and account statuses must still be valid, and the customer must still be associated to the account). If set to true on a virtual card product, activation of newly issued cards on that product will no longer be created in an activated state automatically. With this setting enabled, cards will be issued in an unactivated state, and activation of the cards will be blocked until the customer has passed KYC.  | [optional] 

## Methods

### NewCardProductUpdateRequest

`func NewCardProductUpdateRequest() *CardProductUpdateRequest`

NewCardProductUpdateRequest instantiates a new CardProductUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductUpdateRequestWithDefaults

`func NewCardProductUpdateRequestWithDefaults() *CardProductUpdateRequest`

NewCardProductUpdateRequestWithDefaults instantiates a new CardProductUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigitalWalletTokenization

`func (o *CardProductUpdateRequest) GetDigitalWalletTokenization() DigitalWalletTokenization`

GetDigitalWalletTokenization returns the DigitalWalletTokenization field if non-nil, zero value otherwise.

### GetDigitalWalletTokenizationOk

`func (o *CardProductUpdateRequest) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool)`

GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokenization

`func (o *CardProductUpdateRequest) SetDigitalWalletTokenization(v DigitalWalletTokenization)`

SetDigitalWalletTokenization sets DigitalWalletTokenization field to given value.

### HasDigitalWalletTokenization

`func (o *CardProductUpdateRequest) HasDigitalWalletTokenization() bool`

HasDigitalWalletTokenization returns a boolean if a field has been set.

### GetIssueWithoutKyc

`func (o *CardProductUpdateRequest) GetIssueWithoutKyc() bool`

GetIssueWithoutKyc returns the IssueWithoutKyc field if non-nil, zero value otherwise.

### GetIssueWithoutKycOk

`func (o *CardProductUpdateRequest) GetIssueWithoutKycOk() (*bool, bool)`

GetIssueWithoutKycOk returns a tuple with the IssueWithoutKyc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueWithoutKyc

`func (o *CardProductUpdateRequest) SetIssueWithoutKyc(v bool)`

SetIssueWithoutKyc sets IssueWithoutKyc field to given value.

### HasIssueWithoutKyc

`func (o *CardProductUpdateRequest) HasIssueWithoutKyc() bool`

HasIssueWithoutKyc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


