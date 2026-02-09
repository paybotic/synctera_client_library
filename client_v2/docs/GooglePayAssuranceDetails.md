# GooglePayAssuranceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountVerified** | Pointer to **bool** | Indicates if the cardholder possession validation has been performed on the returned payment credential | [optional] 
**CardholderAuthenticated** | Pointer to **bool** | Indicates if identification and verifications (ID&amp;V) was performed on the returned payment credential | [optional] 

## Methods

### NewGooglePayAssuranceDetails

`func NewGooglePayAssuranceDetails() *GooglePayAssuranceDetails`

NewGooglePayAssuranceDetails instantiates a new GooglePayAssuranceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayAssuranceDetailsWithDefaults

`func NewGooglePayAssuranceDetailsWithDefaults() *GooglePayAssuranceDetails`

NewGooglePayAssuranceDetailsWithDefaults instantiates a new GooglePayAssuranceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountVerified

`func (o *GooglePayAssuranceDetails) GetAccountVerified() bool`

GetAccountVerified returns the AccountVerified field if non-nil, zero value otherwise.

### GetAccountVerifiedOk

`func (o *GooglePayAssuranceDetails) GetAccountVerifiedOk() (*bool, bool)`

GetAccountVerifiedOk returns a tuple with the AccountVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountVerified

`func (o *GooglePayAssuranceDetails) SetAccountVerified(v bool)`

SetAccountVerified sets AccountVerified field to given value.

### HasAccountVerified

`func (o *GooglePayAssuranceDetails) HasAccountVerified() bool`

HasAccountVerified returns a boolean if a field has been set.

### GetCardholderAuthenticated

`func (o *GooglePayAssuranceDetails) GetCardholderAuthenticated() bool`

GetCardholderAuthenticated returns the CardholderAuthenticated field if non-nil, zero value otherwise.

### GetCardholderAuthenticatedOk

`func (o *GooglePayAssuranceDetails) GetCardholderAuthenticatedOk() (*bool, bool)`

GetCardholderAuthenticatedOk returns a tuple with the CardholderAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderAuthenticated

`func (o *GooglePayAssuranceDetails) SetCardholderAuthenticated(v bool)`

SetCardholderAuthenticated sets CardholderAuthenticated field to given value.

### HasCardholderAuthenticated

`func (o *GooglePayAssuranceDetails) HasCardholderAuthenticated() bool`

HasCardholderAuthenticated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


