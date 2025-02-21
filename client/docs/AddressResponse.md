# AddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | The identifier for the business customer associated with address.  | [optional] 
**CreationDate** | **time.Time** | Date and time when the address was created. | 
**CustomerId** | Pointer to **string** | The identifier for the personal customer associated with address.  | [optional] 
**IsActive** | **bool** | Whether the address is active or not | 
**LastUpdatedTime** | **time.Time** |  | 
**AddressLine1** | **string** | Street address line 1 | 
**AddressLine2** | Pointer to **string** | Street address line 2 | [optional] 
**AddressType** | Pointer to **string** | Specifies the address type.  | [optional] [readonly] 
**City** | Pointer to **string** | City | [optional] 
**CountryCode** | **string** | ISO-3166-1 Alpha-2 country code | 
**Id** | **string** | The unique identifier for this resource. | [readonly] 
**IsRegisteredAgent** | Pointer to **bool** | Indicates whether an address is a registered agent. Omitted if the address is not a registered agent. | [optional] 
**Nickname** | Pointer to **string** | A nickname for the address. This is used to identify the address in the UI.  | [optional] 
**PostalCode** | Pointer to **string** | Postal code. For US, formats of 12345 or 12345-1234 are accepted. For CA, formats of A1A 1A1 or A1A1A1 (regardless of case) are accepted, and will be converted to A1A 1A1 format.  | [optional] 
**State** | Pointer to **string** | State, region, province, or prefecture. This is the ISO-3166-2 subdivision code, excluding the country prefix. For example, TX for Texas USA or TAM for Tamaulipas Mexico. Its length varies by country, e.g. 2 characters for US, 3 for MX.  | [optional] 

## Methods

### NewAddressResponse

`func NewAddressResponse(creationDate time.Time, isActive bool, lastUpdatedTime time.Time, addressLine1 string, countryCode string, id string, ) *AddressResponse`

NewAddressResponse instantiates a new AddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressResponseWithDefaults

`func NewAddressResponseWithDefaults() *AddressResponse`

NewAddressResponseWithDefaults instantiates a new AddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *AddressResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AddressResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AddressResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AddressResponse) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationDate

`func (o *AddressResponse) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AddressResponse) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AddressResponse) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.


### GetCustomerId

`func (o *AddressResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AddressResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AddressResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AddressResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetIsActive

`func (o *AddressResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AddressResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AddressResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastUpdatedTime

`func (o *AddressResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AddressResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AddressResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetAddressLine1

`func (o *AddressResponse) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AddressResponse) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AddressResponse) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *AddressResponse) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *AddressResponse) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *AddressResponse) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *AddressResponse) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetAddressType

`func (o *AddressResponse) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *AddressResponse) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *AddressResponse) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *AddressResponse) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetCity

`func (o *AddressResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressResponse) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressResponse) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetId

`func (o *AddressResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsRegisteredAgent

`func (o *AddressResponse) GetIsRegisteredAgent() bool`

GetIsRegisteredAgent returns the IsRegisteredAgent field if non-nil, zero value otherwise.

### GetIsRegisteredAgentOk

`func (o *AddressResponse) GetIsRegisteredAgentOk() (*bool, bool)`

GetIsRegisteredAgentOk returns a tuple with the IsRegisteredAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegisteredAgent

`func (o *AddressResponse) SetIsRegisteredAgent(v bool)`

SetIsRegisteredAgent sets IsRegisteredAgent field to given value.

### HasIsRegisteredAgent

`func (o *AddressResponse) HasIsRegisteredAgent() bool`

HasIsRegisteredAgent returns a boolean if a field has been set.

### GetNickname

`func (o *AddressResponse) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AddressResponse) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AddressResponse) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AddressResponse) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressResponse) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressResponse) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressResponse) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressResponse) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *AddressResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


