# ExternalCardRequestPan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**BusinessId** | Pointer to **string** | Business ID for the application. An application must have either a Business or a customer associated with it. | [optional] 
**CustomerId** | **string** | Customer ID for the application | 
**Cvv** | Pointer to **string** |  | [optional] 
**ExpirationMonth** | **string** |  | 
**ExpirationYear** | **string** |  | 
**Name** | **string** | The cardholder name | 
**Pan** | **string** |  | 

## Methods

### NewExternalCardRequestPan

`func NewExternalCardRequestPan(customerId string, expirationMonth string, expirationYear string, name string, pan string, ) *ExternalCardRequestPan`

NewExternalCardRequestPan instantiates a new ExternalCardRequestPan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardRequestPanWithDefaults

`func NewExternalCardRequestPanWithDefaults() *ExternalCardRequestPan`

NewExternalCardRequestPanWithDefaults instantiates a new ExternalCardRequestPan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *ExternalCardRequestPan) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ExternalCardRequestPan) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ExternalCardRequestPan) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *ExternalCardRequestPan) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBusinessId

`func (o *ExternalCardRequestPan) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *ExternalCardRequestPan) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *ExternalCardRequestPan) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *ExternalCardRequestPan) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *ExternalCardRequestPan) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExternalCardRequestPan) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExternalCardRequestPan) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCvv

`func (o *ExternalCardRequestPan) GetCvv() string`

GetCvv returns the Cvv field if non-nil, zero value otherwise.

### GetCvvOk

`func (o *ExternalCardRequestPan) GetCvvOk() (*string, bool)`

GetCvvOk returns a tuple with the Cvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv

`func (o *ExternalCardRequestPan) SetCvv(v string)`

SetCvv sets Cvv field to given value.

### HasCvv

`func (o *ExternalCardRequestPan) HasCvv() bool`

HasCvv returns a boolean if a field has been set.

### GetExpirationMonth

`func (o *ExternalCardRequestPan) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *ExternalCardRequestPan) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *ExternalCardRequestPan) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.


### GetExpirationYear

`func (o *ExternalCardRequestPan) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *ExternalCardRequestPan) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *ExternalCardRequestPan) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.


### GetName

`func (o *ExternalCardRequestPan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalCardRequestPan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalCardRequestPan) SetName(v string)`

SetName sets Name field to given value.


### GetPan

`func (o *ExternalCardRequestPan) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *ExternalCardRequestPan) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *ExternalCardRequestPan) SetPan(v string)`

SetPan sets Pan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


