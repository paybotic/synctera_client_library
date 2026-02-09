# ApplePayContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLines** | Pointer to **[]string** | Street portion of the contact’s address | [optional] 
**AdministrativeArea** | Pointer to **string** | The state for the contact | [optional] 
**Country** | Pointer to **string** | The name of the country or region for the contact | [optional] 
**CountryCode** | Pointer to **string** | ISO-3166-1 Alpha-2 country code | [optional] 
**EmailAddress** | Pointer to **string** | Email address | [optional] 
**FamilyName** | Pointer to **string** | Family name | [optional] 
**GivenName** | Pointer to **string** | Given name | [optional] 
**Locality** | Pointer to **string** | The city for the contact | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number | [optional] 
**PhoneticFamilyName** | Pointer to **string** | Phonetic spelling of the contact’s family name | [optional] 
**PhoneticGivenName** | Pointer to **string** | Phonetic spelling of the contact’s given name | [optional] 
**PostalCode** | Pointer to **string** | Postal code | [optional] 
**SubAdministrativeArea** | Pointer to **string** | The subadministrative area (such as a county or other region) in a postal address | [optional] 
**SubLocality** | Pointer to **string** | Additional information associated with the location, typically defined at the city or town level (such as district or neighborhood), in a postal address | [optional] 

## Methods

### NewApplePayContact

`func NewApplePayContact() *ApplePayContact`

NewApplePayContact instantiates a new ApplePayContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePayContactWithDefaults

`func NewApplePayContactWithDefaults() *ApplePayContact`

NewApplePayContactWithDefaults instantiates a new ApplePayContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLines

`func (o *ApplePayContact) GetAddressLines() []string`

GetAddressLines returns the AddressLines field if non-nil, zero value otherwise.

### GetAddressLinesOk

`func (o *ApplePayContact) GetAddressLinesOk() (*[]string, bool)`

GetAddressLinesOk returns a tuple with the AddressLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLines

`func (o *ApplePayContact) SetAddressLines(v []string)`

SetAddressLines sets AddressLines field to given value.

### HasAddressLines

`func (o *ApplePayContact) HasAddressLines() bool`

HasAddressLines returns a boolean if a field has been set.

### GetAdministrativeArea

`func (o *ApplePayContact) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *ApplePayContact) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *ApplePayContact) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *ApplePayContact) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetCountry

`func (o *ApplePayContact) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ApplePayContact) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ApplePayContact) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ApplePayContact) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *ApplePayContact) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ApplePayContact) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ApplePayContact) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ApplePayContact) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmailAddress

`func (o *ApplePayContact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ApplePayContact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ApplePayContact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ApplePayContact) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetFamilyName

`func (o *ApplePayContact) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *ApplePayContact) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *ApplePayContact) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *ApplePayContact) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetGivenName

`func (o *ApplePayContact) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *ApplePayContact) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *ApplePayContact) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *ApplePayContact) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetLocality

`func (o *ApplePayContact) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *ApplePayContact) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *ApplePayContact) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *ApplePayContact) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ApplePayContact) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ApplePayContact) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ApplePayContact) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ApplePayContact) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneticFamilyName

`func (o *ApplePayContact) GetPhoneticFamilyName() string`

GetPhoneticFamilyName returns the PhoneticFamilyName field if non-nil, zero value otherwise.

### GetPhoneticFamilyNameOk

`func (o *ApplePayContact) GetPhoneticFamilyNameOk() (*string, bool)`

GetPhoneticFamilyNameOk returns a tuple with the PhoneticFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneticFamilyName

`func (o *ApplePayContact) SetPhoneticFamilyName(v string)`

SetPhoneticFamilyName sets PhoneticFamilyName field to given value.

### HasPhoneticFamilyName

`func (o *ApplePayContact) HasPhoneticFamilyName() bool`

HasPhoneticFamilyName returns a boolean if a field has been set.

### GetPhoneticGivenName

`func (o *ApplePayContact) GetPhoneticGivenName() string`

GetPhoneticGivenName returns the PhoneticGivenName field if non-nil, zero value otherwise.

### GetPhoneticGivenNameOk

`func (o *ApplePayContact) GetPhoneticGivenNameOk() (*string, bool)`

GetPhoneticGivenNameOk returns a tuple with the PhoneticGivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneticGivenName

`func (o *ApplePayContact) SetPhoneticGivenName(v string)`

SetPhoneticGivenName sets PhoneticGivenName field to given value.

### HasPhoneticGivenName

`func (o *ApplePayContact) HasPhoneticGivenName() bool`

HasPhoneticGivenName returns a boolean if a field has been set.

### GetPostalCode

`func (o *ApplePayContact) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ApplePayContact) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ApplePayContact) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ApplePayContact) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetSubAdministrativeArea

`func (o *ApplePayContact) GetSubAdministrativeArea() string`

GetSubAdministrativeArea returns the SubAdministrativeArea field if non-nil, zero value otherwise.

### GetSubAdministrativeAreaOk

`func (o *ApplePayContact) GetSubAdministrativeAreaOk() (*string, bool)`

GetSubAdministrativeAreaOk returns a tuple with the SubAdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAdministrativeArea

`func (o *ApplePayContact) SetSubAdministrativeArea(v string)`

SetSubAdministrativeArea sets SubAdministrativeArea field to given value.

### HasSubAdministrativeArea

`func (o *ApplePayContact) HasSubAdministrativeArea() bool`

HasSubAdministrativeArea returns a boolean if a field has been set.

### GetSubLocality

`func (o *ApplePayContact) GetSubLocality() string`

GetSubLocality returns the SubLocality field if non-nil, zero value otherwise.

### GetSubLocalityOk

`func (o *ApplePayContact) GetSubLocalityOk() (*string, bool)`

GetSubLocalityOk returns a tuple with the SubLocality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLocality

`func (o *ApplePayContact) SetSubLocality(v string)`

SetSubLocality sets SubLocality field to given value.

### HasSubLocality

`func (o *ApplePayContact) HasSubLocality() bool`

HasSubLocality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


