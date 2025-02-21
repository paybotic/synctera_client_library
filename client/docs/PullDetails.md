# PullDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | ISO-3166-1 Alpha-2 country code | [optional] 
**Currency** | Pointer to **string** | ISO 4217  Alpha-3 currency code | [optional] 
**Network** | Pointer to **string** | Payment network | [optional] 
**ProductType** | Pointer to [**ExternalCardProductType**](ExternalCardProductType.md) |  | [optional] 
**Regulated** | Pointer to **bool** | Exemption status from debit card interchange fee standards | [optional] 

## Methods

### NewPullDetails

`func NewPullDetails() *PullDetails`

NewPullDetails instantiates a new PullDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullDetailsWithDefaults

`func NewPullDetailsWithDefaults() *PullDetails`

NewPullDetailsWithDefaults instantiates a new PullDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PullDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PullDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PullDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PullDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *PullDetails) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PullDetails) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PullDetails) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PullDetails) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetNetwork

`func (o *PullDetails) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PullDetails) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PullDetails) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PullDetails) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProductType

`func (o *PullDetails) GetProductType() ExternalCardProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PullDetails) GetProductTypeOk() (*ExternalCardProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PullDetails) SetProductType(v ExternalCardProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *PullDetails) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetRegulated

`func (o *PullDetails) GetRegulated() bool`

GetRegulated returns the Regulated field if non-nil, zero value otherwise.

### GetRegulatedOk

`func (o *PullDetails) GetRegulatedOk() (*bool, bool)`

GetRegulatedOk returns a tuple with the Regulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulated

`func (o *PullDetails) SetRegulated(v bool)`

SetRegulated sets Regulated field to given value.

### HasRegulated

`func (o *PullDetails) HasRegulated() bool`

HasRegulated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


