# GooglePayAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** | Address line 1 | [optional] 
**Address2** | Pointer to **string** | Address line 2 | [optional] 
**Address3** | Pointer to **string** | Address line 3 | [optional] 
**AdministrativeArea** | Pointer to **string** | Country subdivision, such as state or province | [optional] 
**CountryCode** | Pointer to **string** | ISO-3166-1 Alpha-2 country code | [optional] 
**Locality** | Pointer to **string** | City, town, neighborhood or suburb | [optional] 
**Name** | Pointer to **string** | Full name of addressee | [optional] 
**PhoneNumber** | Pointer to **string** | Telephone number | [optional] 
**PostalCode** | Pointer to **string** | Postal or ZIP code | [optional] 
**SortingCode** | Pointer to **string** | Sorting code | [optional] 

## Methods

### NewGooglePayAddress

`func NewGooglePayAddress() *GooglePayAddress`

NewGooglePayAddress instantiates a new GooglePayAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayAddressWithDefaults

`func NewGooglePayAddressWithDefaults() *GooglePayAddress`

NewGooglePayAddressWithDefaults instantiates a new GooglePayAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *GooglePayAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *GooglePayAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *GooglePayAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *GooglePayAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *GooglePayAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *GooglePayAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *GooglePayAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *GooglePayAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetAddress3

`func (o *GooglePayAddress) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *GooglePayAddress) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *GooglePayAddress) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *GooglePayAddress) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### GetAdministrativeArea

`func (o *GooglePayAddress) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *GooglePayAddress) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *GooglePayAddress) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *GooglePayAddress) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetCountryCode

`func (o *GooglePayAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *GooglePayAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *GooglePayAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *GooglePayAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLocality

`func (o *GooglePayAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *GooglePayAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *GooglePayAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *GooglePayAddress) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetName

`func (o *GooglePayAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GooglePayAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GooglePayAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GooglePayAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *GooglePayAddress) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GooglePayAddress) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GooglePayAddress) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GooglePayAddress) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPostalCode

`func (o *GooglePayAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *GooglePayAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *GooglePayAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *GooglePayAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetSortingCode

`func (o *GooglePayAddress) GetSortingCode() string`

GetSortingCode returns the SortingCode field if non-nil, zero value otherwise.

### GetSortingCodeOk

`func (o *GooglePayAddress) GetSortingCodeOk() (*string, bool)`

GetSortingCodeOk returns a tuple with the SortingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingCode

`func (o *GooglePayAddress) SetSortingCode(v string)`

SetSortingCode sets SortingCode field to given value.

### HasSortingCode

`func (o *GooglePayAddress) HasSortingCode() bool`

HasSortingCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


