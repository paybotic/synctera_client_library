# IatAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CityStateProvince** | **string** | City and State / Province. Should be separated with an asterisk (*) as a delimiter. | 
**CountryPostalCode** | **string** | Country and Postal Code. Should be separated with an asterisk (*) as a delimiter. | 
**Street** | **string** | The street address | 

## Methods

### NewIatAddress

`func NewIatAddress(cityStateProvince string, countryPostalCode string, street string, ) *IatAddress`

NewIatAddress instantiates a new IatAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIatAddressWithDefaults

`func NewIatAddressWithDefaults() *IatAddress`

NewIatAddressWithDefaults instantiates a new IatAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCityStateProvince

`func (o *IatAddress) GetCityStateProvince() string`

GetCityStateProvince returns the CityStateProvince field if non-nil, zero value otherwise.

### GetCityStateProvinceOk

`func (o *IatAddress) GetCityStateProvinceOk() (*string, bool)`

GetCityStateProvinceOk returns a tuple with the CityStateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityStateProvince

`func (o *IatAddress) SetCityStateProvince(v string)`

SetCityStateProvince sets CityStateProvince field to given value.


### GetCountryPostalCode

`func (o *IatAddress) GetCountryPostalCode() string`

GetCountryPostalCode returns the CountryPostalCode field if non-nil, zero value otherwise.

### GetCountryPostalCodeOk

`func (o *IatAddress) GetCountryPostalCodeOk() (*string, bool)`

GetCountryPostalCodeOk returns a tuple with the CountryPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryPostalCode

`func (o *IatAddress) SetCountryPostalCode(v string)`

SetCountryPostalCode sets CountryPostalCode field to given value.


### GetStreet

`func (o *IatAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *IatAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *IatAddress) SetStreet(v string)`

SetStreet sets Street field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


