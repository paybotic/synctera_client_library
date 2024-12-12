# ResponsePersonalId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | The ISO 3166 Alpha-2 country code for the country that issued the personal identifier.  | 
**Id** | **string** | UUID for the personal identifier for subsequent changes and deletion | [readonly] 
**IdType** | [**PersonalIdType**](PersonalIdType.md) |  | 
**Identifier** | **string** | The personal identifier. Format varies by personal identifier type. | 
**SystemProvided** | **bool** | True if the identifier was provided by the system, e.g. via SSN Prefill. | [readonly] 

## Methods

### NewResponsePersonalId

`func NewResponsePersonalId(countryCode string, id string, idType PersonalIdType, identifier string, systemProvided bool, ) *ResponsePersonalId`

NewResponsePersonalId instantiates a new ResponsePersonalId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePersonalIdWithDefaults

`func NewResponsePersonalIdWithDefaults() *ResponsePersonalId`

NewResponsePersonalIdWithDefaults instantiates a new ResponsePersonalId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *ResponsePersonalId) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ResponsePersonalId) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ResponsePersonalId) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetId

`func (o *ResponsePersonalId) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePersonalId) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePersonalId) SetId(v string)`

SetId sets Id field to given value.


### GetIdType

`func (o *ResponsePersonalId) GetIdType() PersonalIdType`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *ResponsePersonalId) GetIdTypeOk() (*PersonalIdType, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *ResponsePersonalId) SetIdType(v PersonalIdType)`

SetIdType sets IdType field to given value.


### GetIdentifier

`func (o *ResponsePersonalId) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ResponsePersonalId) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ResponsePersonalId) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetSystemProvided

`func (o *ResponsePersonalId) GetSystemProvided() bool`

GetSystemProvided returns the SystemProvided field if non-nil, zero value otherwise.

### GetSystemProvidedOk

`func (o *ResponsePersonalId) GetSystemProvidedOk() (*bool, bool)`

GetSystemProvidedOk returns a tuple with the SystemProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemProvided

`func (o *ResponsePersonalId) SetSystemProvided(v bool)`

SetSystemProvided sets SystemProvided field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


