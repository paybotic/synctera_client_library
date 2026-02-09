# CreditScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The unique identifier of a customer | 
**Score** | Pointer to **int32** | The credit score value | [optional] 
**ScoreRequestedTime** | Pointer to **time.Time** | The time the credit score was requested | [optional] 
**SourceOfScore** | [**SourceOfScore**](SourceOfScore.md) |  | 
**Type** | Pointer to **string** | The type of the credit score | [optional] 
**VendorName** | Pointer to [**VendorName**](VendorName.md) |  | [optional] 
**Version** | Pointer to **string** | The version of the credit score | [optional] 

## Methods

### NewCreditScore

`func NewCreditScore(customerId string, sourceOfScore SourceOfScore, ) *CreditScore`

NewCreditScore instantiates a new CreditScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditScoreWithDefaults

`func NewCreditScoreWithDefaults() *CreditScore`

NewCreditScoreWithDefaults instantiates a new CreditScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CreditScore) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreditScore) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreditScore) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetScore

`func (o *CreditScore) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreditScore) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreditScore) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *CreditScore) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetScoreRequestedTime

`func (o *CreditScore) GetScoreRequestedTime() time.Time`

GetScoreRequestedTime returns the ScoreRequestedTime field if non-nil, zero value otherwise.

### GetScoreRequestedTimeOk

`func (o *CreditScore) GetScoreRequestedTimeOk() (*time.Time, bool)`

GetScoreRequestedTimeOk returns a tuple with the ScoreRequestedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreRequestedTime

`func (o *CreditScore) SetScoreRequestedTime(v time.Time)`

SetScoreRequestedTime sets ScoreRequestedTime field to given value.

### HasScoreRequestedTime

`func (o *CreditScore) HasScoreRequestedTime() bool`

HasScoreRequestedTime returns a boolean if a field has been set.

### GetSourceOfScore

`func (o *CreditScore) GetSourceOfScore() SourceOfScore`

GetSourceOfScore returns the SourceOfScore field if non-nil, zero value otherwise.

### GetSourceOfScoreOk

`func (o *CreditScore) GetSourceOfScoreOk() (*SourceOfScore, bool)`

GetSourceOfScoreOk returns a tuple with the SourceOfScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfScore

`func (o *CreditScore) SetSourceOfScore(v SourceOfScore)`

SetSourceOfScore sets SourceOfScore field to given value.


### GetType

`func (o *CreditScore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditScore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditScore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreditScore) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendorName

`func (o *CreditScore) GetVendorName() VendorName`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *CreditScore) GetVendorNameOk() (*VendorName, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *CreditScore) SetVendorName(v VendorName)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *CreditScore) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetVersion

`func (o *CreditScore) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreditScore) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreditScore) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreditScore) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


