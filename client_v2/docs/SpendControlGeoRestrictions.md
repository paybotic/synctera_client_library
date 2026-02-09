# SpendControlGeoRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | **[]string** | ISO 3166-1 alpha-2 country codes to block or allow | 
**CountrySources** | Pointer to [**[]CountrySource**](CountrySource.md) | Which country attributes on the transaction to evaluate | [optional] 
**Mode** | [**GeoRestrictionMode**](GeoRestrictionMode.md) |  | 

## Methods

### NewSpendControlGeoRestrictions

`func NewSpendControlGeoRestrictions(countries []string, mode GeoRestrictionMode, ) *SpendControlGeoRestrictions`

NewSpendControlGeoRestrictions instantiates a new SpendControlGeoRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendControlGeoRestrictionsWithDefaults

`func NewSpendControlGeoRestrictionsWithDefaults() *SpendControlGeoRestrictions`

NewSpendControlGeoRestrictionsWithDefaults instantiates a new SpendControlGeoRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *SpendControlGeoRestrictions) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *SpendControlGeoRestrictions) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *SpendControlGeoRestrictions) SetCountries(v []string)`

SetCountries sets Countries field to given value.


### GetCountrySources

`func (o *SpendControlGeoRestrictions) GetCountrySources() []CountrySource`

GetCountrySources returns the CountrySources field if non-nil, zero value otherwise.

### GetCountrySourcesOk

`func (o *SpendControlGeoRestrictions) GetCountrySourcesOk() (*[]CountrySource, bool)`

GetCountrySourcesOk returns a tuple with the CountrySources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySources

`func (o *SpendControlGeoRestrictions) SetCountrySources(v []CountrySource)`

SetCountrySources sets CountrySources field to given value.

### HasCountrySources

`func (o *SpendControlGeoRestrictions) HasCountrySources() bool`

HasCountrySources returns a boolean if a field has been set.

### GetMode

`func (o *SpendControlGeoRestrictions) GetMode() GeoRestrictionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SpendControlGeoRestrictions) GetModeOk() (*GeoRestrictionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SpendControlGeoRestrictions) SetMode(v GeoRestrictionMode)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


