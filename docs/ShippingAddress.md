# ShippingAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | Street address line 1 | 
**AddressLine2** | Pointer to **string** | Street address line 2 | [optional] 
**AddressType** | Pointer to **string** | Specifies the address type.  | [optional] [readonly] 
**City** | Pointer to **string** | City | [optional] 
**CountryCode** | **string** | ISO-3166-1 Alpha-2 country code | 
**Id** | Pointer to **string** | The unique identifier for this resource. | [optional] [readonly] 
**Nickname** | Pointer to **string** | A nickname for the address. This is used to identify the address in the UI.  | [optional] 
**PostalCode** | Pointer to **string** | Postal code. For US, formats of 12345 or 12345-1234 are accepted. For CA, formats of A1A 1A1 or A1A1A1 (regardless of case) are accepted, and will be converted to A1A 1A1 format.  | [optional] 
**State** | Pointer to **string** | State, region, province, or prefecture. This is the ISO-3166-2 subdivision code, excluding the country prefix. For example, TX for Texas USA or TAM for Tamaulipas Mexico. Its length varies by country, e.g. 2 characters for US, 3 for MX.  | [optional] 

## Methods

### NewShippingAddress

`func NewShippingAddress(addressLine1 string, countryCode string, ) *ShippingAddress`

NewShippingAddress instantiates a new ShippingAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingAddressWithDefaults

`func NewShippingAddressWithDefaults() *ShippingAddress`

NewShippingAddressWithDefaults instantiates a new ShippingAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *ShippingAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ShippingAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ShippingAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *ShippingAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ShippingAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ShippingAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ShippingAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressType

`func (o *ShippingAddress) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *ShippingAddress) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *ShippingAddress) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *ShippingAddress) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetCity

`func (o *ShippingAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ShippingAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ShippingAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ShippingAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *ShippingAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ShippingAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ShippingAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetId

`func (o *ShippingAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShippingAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNickname

`func (o *ShippingAddress) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *ShippingAddress) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *ShippingAddress) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *ShippingAddress) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPostalCode

`func (o *ShippingAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ShippingAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ShippingAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ShippingAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *ShippingAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ShippingAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ShippingAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ShippingAddress) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


