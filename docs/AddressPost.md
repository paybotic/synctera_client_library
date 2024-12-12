# AddressPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | Street address line 1 | 
**AddressLine2** | Pointer to **string** | Street address line 2 | [optional] 
**City** | **string** | City | 
**CountryCode** | **string** | ISO-3166-1 Alpha-2 country code | 
**IsActive** | Pointer to **bool** | Whether the address is active or not | [optional] 
**Nickname** | Pointer to **string** | A nickname for the address. This is used to identify the address in the UI.  | [optional] 
**PostalCode** | Pointer to **string** | Postal code. For US, formats of 12345 or 12345-1234 are accepted. For CA, formats of A1A 1A1 or A1A1A1 (regardless of case) are accepted, and will be converted to A1A 1A1 format.  | [optional] 
**State** | **string** | State, region, province, or prefecture. This is the ISO-3166-2 subdivision code, excluding the country prefix. For example, TX for Texas USA or TAM for Tamaulipas Mexico. Its length varies by country, e.g. 2 characters for US, 3 for MX.  | 
**AddressType** | **string** | Specifies the address type.  | 
**BusinessId** | Pointer to **string** | The identifier for the business customer associated with address.  | [optional] 
**CustomerId** | Pointer to **string** | The identifier for the personal customer associated with address.  | [optional] 

## Methods

### NewAddressPost

`func NewAddressPost(addressLine1 string, city string, countryCode string, state string, addressType string, ) *AddressPost`

NewAddressPost instantiates a new AddressPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressPostWithDefaults

`func NewAddressPostWithDefaults() *AddressPost`

NewAddressPostWithDefaults instantiates a new AddressPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *AddressPost) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AddressPost) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AddressPost) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *AddressPost) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressPost) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressPost) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AddressPost) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *AddressPost) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressPost) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressPost) SetCity(v string)`

SetCity sets City field to given value.


### GetCountryCode

`func (o *AddressPost) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressPost) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressPost) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetIsActive

`func (o *AddressPost) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AddressPost) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AddressPost) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AddressPost) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetNickname

`func (o *AddressPost) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AddressPost) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AddressPost) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AddressPost) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressPost) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressPost) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressPost) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressPost) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *AddressPost) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressPost) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressPost) SetState(v string)`

SetState sets State field to given value.


### GetAddressType

`func (o *AddressPost) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *AddressPost) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *AddressPost) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetBusinessId

`func (o *AddressPost) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AddressPost) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AddressPost) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AddressPost) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AddressPost) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AddressPost) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AddressPost) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AddressPost) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


