# PersonalIdCountryCodePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | The ISO 3166 Alpha-2 country code for the country that issued the personal identifier. This is optional for personal identifier types that have an implicit country, e.g. SSN. This is required for other types, e.g. PASSPORT_NUMBER.  | [optional] 

## Methods

### NewPersonalIdCountryCodePost

`func NewPersonalIdCountryCodePost() *PersonalIdCountryCodePost`

NewPersonalIdCountryCodePost instantiates a new PersonalIdCountryCodePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalIdCountryCodePostWithDefaults

`func NewPersonalIdCountryCodePostWithDefaults() *PersonalIdCountryCodePost`

NewPersonalIdCountryCodePostWithDefaults instantiates a new PersonalIdCountryCodePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *PersonalIdCountryCodePost) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PersonalIdCountryCodePost) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PersonalIdCountryCodePost) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PersonalIdCountryCodePost) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


