# Initialize3dsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount in cents of the External Card Transfer to be authenticated | 
**Currency** | **string** | ISO 4217  Alpha-3 currency code | 
**ExternalCardId** | **string** | The ID of the External Card for which the 3DS Authentication will be performed | 

## Methods

### NewInitialize3dsRequest

`func NewInitialize3dsRequest(amount int32, currency string, externalCardId string, ) *Initialize3dsRequest`

NewInitialize3dsRequest instantiates a new Initialize3dsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialize3dsRequestWithDefaults

`func NewInitialize3dsRequestWithDefaults() *Initialize3dsRequest`

NewInitialize3dsRequestWithDefaults instantiates a new Initialize3dsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Initialize3dsRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Initialize3dsRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Initialize3dsRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *Initialize3dsRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Initialize3dsRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Initialize3dsRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExternalCardId

`func (o *Initialize3dsRequest) GetExternalCardId() string`

GetExternalCardId returns the ExternalCardId field if non-nil, zero value otherwise.

### GetExternalCardIdOk

`func (o *Initialize3dsRequest) GetExternalCardIdOk() (*string, bool)`

GetExternalCardIdOk returns a tuple with the ExternalCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCardId

`func (o *Initialize3dsRequest) SetExternalCardId(v string)`

SetExternalCardId sets ExternalCardId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


