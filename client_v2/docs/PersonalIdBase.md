# PersonalIdBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | UUID for the personal identifier for subsequent changes and deletion | [optional] [readonly] 
**IdType** | Pointer to [**PersonalIdType**](PersonalIdType.md) |  | [optional] 
**Identifier** | Pointer to **string** | The personal identifier. Format varies by personal identifier type. | [optional] 
**SystemProvided** | Pointer to **bool** | True if the identifier was provided by the system, e.g. via SSN Prefill. | [optional] [readonly] 

## Methods

### NewPersonalIdBase

`func NewPersonalIdBase() *PersonalIdBase`

NewPersonalIdBase instantiates a new PersonalIdBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalIdBaseWithDefaults

`func NewPersonalIdBaseWithDefaults() *PersonalIdBase`

NewPersonalIdBaseWithDefaults instantiates a new PersonalIdBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonalIdBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonalIdBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonalIdBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PersonalIdBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdType

`func (o *PersonalIdBase) GetIdType() PersonalIdType`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *PersonalIdBase) GetIdTypeOk() (*PersonalIdType, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *PersonalIdBase) SetIdType(v PersonalIdType)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *PersonalIdBase) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetIdentifier

`func (o *PersonalIdBase) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PersonalIdBase) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PersonalIdBase) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PersonalIdBase) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSystemProvided

`func (o *PersonalIdBase) GetSystemProvided() bool`

GetSystemProvided returns the SystemProvided field if non-nil, zero value otherwise.

### GetSystemProvidedOk

`func (o *PersonalIdBase) GetSystemProvidedOk() (*bool, bool)`

GetSystemProvidedOk returns a tuple with the SystemProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemProvided

`func (o *PersonalIdBase) SetSystemProvided(v bool)`

SetSystemProvided sets SystemProvided field to given value.

### HasSystemProvided

`func (o *PersonalIdBase) HasSystemProvided() bool`

HasSystemProvided returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


