# PostPersonalIdWCust

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Id of the customer having this personal identifier | 
**Id** | Pointer to **string** | UUID for the personal identifier for subsequent changes and deletion | [optional] [readonly] 
**IdType** | [**PersonalIdType**](PersonalIdType.md) |  | 
**Identifier** | **string** | The personal identifier. Format varies by personal identifier type. | 
**SystemProvided** | Pointer to **bool** | True if the identifier was provided by the system, e.g. via SSN Prefill. | [optional] [readonly] 
**CountryCode** | Pointer to **string** | The ISO 3166 Alpha-2 country code for the country that issued the personal identifier. This is optional for personal identifier types that have an implicit country, e.g. SSN. This is required for other types, e.g. PASSPORT_NUMBER.  | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 

## Methods

### NewPostPersonalIdWCust

`func NewPostPersonalIdWCust(customerId string, idType PersonalIdType, identifier string, ) *PostPersonalIdWCust`

NewPostPersonalIdWCust instantiates a new PostPersonalIdWCust object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPersonalIdWCustWithDefaults

`func NewPostPersonalIdWCustWithDefaults() *PostPersonalIdWCust`

NewPostPersonalIdWCustWithDefaults instantiates a new PostPersonalIdWCust object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *PostPersonalIdWCust) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PostPersonalIdWCust) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PostPersonalIdWCust) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *PostPersonalIdWCust) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostPersonalIdWCust) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostPersonalIdWCust) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostPersonalIdWCust) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdType

`func (o *PostPersonalIdWCust) GetIdType() PersonalIdType`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *PostPersonalIdWCust) GetIdTypeOk() (*PersonalIdType, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *PostPersonalIdWCust) SetIdType(v PersonalIdType)`

SetIdType sets IdType field to given value.


### GetIdentifier

`func (o *PostPersonalIdWCust) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PostPersonalIdWCust) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PostPersonalIdWCust) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetSystemProvided

`func (o *PostPersonalIdWCust) GetSystemProvided() bool`

GetSystemProvided returns the SystemProvided field if non-nil, zero value otherwise.

### GetSystemProvidedOk

`func (o *PostPersonalIdWCust) GetSystemProvidedOk() (*bool, bool)`

GetSystemProvidedOk returns a tuple with the SystemProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemProvided

`func (o *PostPersonalIdWCust) SetSystemProvided(v bool)`

SetSystemProvided sets SystemProvided field to given value.

### HasSystemProvided

`func (o *PostPersonalIdWCust) HasSystemProvided() bool`

HasSystemProvided returns a boolean if a field has been set.

### GetCountryCode

`func (o *PostPersonalIdWCust) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PostPersonalIdWCust) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PostPersonalIdWCust) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PostPersonalIdWCust) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetTenant

`func (o *PostPersonalIdWCust) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PostPersonalIdWCust) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PostPersonalIdWCust) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PostPersonalIdWCust) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


