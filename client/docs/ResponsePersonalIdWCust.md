# ResponsePersonalIdWCust

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Id of the customer having this personal identifier | 
**CountryCode** | **string** | The ISO 3166 Alpha-2 country code for the country that issued the personal identifier.  | 
**Id** | **string** | UUID for the personal identifier for subsequent changes and deletion | [readonly] 
**IdType** | [**PersonalIdType**](PersonalIdType.md) |  | 
**Identifier** | **string** | The personal identifier. Format varies by personal identifier type. | 
**SystemProvided** | **bool** | True if the identifier was provided by the system, e.g. via SSN Prefill. | [readonly] 
**Tenant** | **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | 

## Methods

### NewResponsePersonalIdWCust

`func NewResponsePersonalIdWCust(customerId string, countryCode string, id string, idType PersonalIdType, identifier string, systemProvided bool, tenant string, ) *ResponsePersonalIdWCust`

NewResponsePersonalIdWCust instantiates a new ResponsePersonalIdWCust object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePersonalIdWCustWithDefaults

`func NewResponsePersonalIdWCustWithDefaults() *ResponsePersonalIdWCust`

NewResponsePersonalIdWCustWithDefaults instantiates a new ResponsePersonalIdWCust object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *ResponsePersonalIdWCust) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ResponsePersonalIdWCust) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ResponsePersonalIdWCust) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCountryCode

`func (o *ResponsePersonalIdWCust) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ResponsePersonalIdWCust) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ResponsePersonalIdWCust) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetId

`func (o *ResponsePersonalIdWCust) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePersonalIdWCust) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePersonalIdWCust) SetId(v string)`

SetId sets Id field to given value.


### GetIdType

`func (o *ResponsePersonalIdWCust) GetIdType() PersonalIdType`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *ResponsePersonalIdWCust) GetIdTypeOk() (*PersonalIdType, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *ResponsePersonalIdWCust) SetIdType(v PersonalIdType)`

SetIdType sets IdType field to given value.


### GetIdentifier

`func (o *ResponsePersonalIdWCust) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ResponsePersonalIdWCust) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ResponsePersonalIdWCust) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetSystemProvided

`func (o *ResponsePersonalIdWCust) GetSystemProvided() bool`

GetSystemProvided returns the SystemProvided field if non-nil, zero value otherwise.

### GetSystemProvidedOk

`func (o *ResponsePersonalIdWCust) GetSystemProvidedOk() (*bool, bool)`

GetSystemProvidedOk returns a tuple with the SystemProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemProvided

`func (o *ResponsePersonalIdWCust) SetSystemProvided(v bool)`

SetSystemProvided sets SystemProvided field to given value.


### GetTenant

`func (o *ResponsePersonalIdWCust) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ResponsePersonalIdWCust) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ResponsePersonalIdWCust) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


