# LegalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | Street address line 1 | 
**AddressLine2** | Pointer to **string** | Street address line 2 | [optional] 
**AddressType** | Pointer to **string** | Specifies the address type.  | [optional] [readonly] 
**City** | Pointer to **string** | City | [optional] 
**CountryCode** | **string** | ISO-3166-1 Alpha-2 country code | 
**Id** | Pointer to **string** | The unique identifier for this resource. | [optional] [readonly] 
**IsRegisteredAgent** | Pointer to **bool** | Indicates whether an address is a registered agent. Omitted if the address is not a registered agent. | [optional] 
**Nickname** | Pointer to **string** | A nickname for the address. This is used to identify the address in the UI.  | [optional] 
**PostalCode** | Pointer to **string** | Postal code. For US, formats of 12345 or 12345-1234 are accepted. For CA, formats of A1A 1A1 or A1A1A1 (regardless of case) are accepted, and will be converted to A1A 1A1 format.  | [optional] 
**State** | Pointer to **string** | State, region, province, or prefecture. This is the ISO-3166-2 subdivision code, excluding the country prefix. For example, TX for Texas USA or TAM for Tamaulipas Mexico. Its length varies by country, e.g. 2 characters for US, 3 for MX.  | [optional] 

## Methods

### NewLegalAddress

`func NewLegalAddress(addressLine1 string, countryCode string, ) *LegalAddress`

NewLegalAddress instantiates a new LegalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalAddressWithDefaults

`func NewLegalAddressWithDefaults() *LegalAddress`

NewLegalAddressWithDefaults instantiates a new LegalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *LegalAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *LegalAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *LegalAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *LegalAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *LegalAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *LegalAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *LegalAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressType

`func (o *LegalAddress) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *LegalAddress) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *LegalAddress) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *LegalAddress) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetCity

`func (o *LegalAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LegalAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LegalAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LegalAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *LegalAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *LegalAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *LegalAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetId

`func (o *LegalAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegalAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegalAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegalAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRegisteredAgent

`func (o *LegalAddress) GetIsRegisteredAgent() bool`

GetIsRegisteredAgent returns the IsRegisteredAgent field if non-nil, zero value otherwise.

### GetIsRegisteredAgentOk

`func (o *LegalAddress) GetIsRegisteredAgentOk() (*bool, bool)`

GetIsRegisteredAgentOk returns a tuple with the IsRegisteredAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegisteredAgent

`func (o *LegalAddress) SetIsRegisteredAgent(v bool)`

SetIsRegisteredAgent sets IsRegisteredAgent field to given value.

### HasIsRegisteredAgent

`func (o *LegalAddress) HasIsRegisteredAgent() bool`

HasIsRegisteredAgent returns a boolean if a field has been set.

### GetNickname

`func (o *LegalAddress) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *LegalAddress) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *LegalAddress) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *LegalAddress) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPostalCode

`func (o *LegalAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *LegalAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *LegalAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *LegalAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *LegalAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LegalAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LegalAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LegalAddress) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


