# AccountRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | Business associated with the current account. One of business_id or customer_id must be specified. | [optional] 
**CustomerId** | Pointer to **string** | Personal customer associated with the current account. One of customer_id or business_id must be specified. | [optional] 
**Id** | Pointer to **string** | ID of account relationship | [optional] [readonly] 
**PersonId** | Pointer to **string** | Person associated with the current account. This attribute is deprecated and will be removed in a future API version. Use customer_id instead. | [optional] 
**RelationshipType** | [**AccountRelationshipType**](AccountRelationshipType.md) |  | 

## Methods

### NewAccountRelationship

`func NewAccountRelationship(relationshipType AccountRelationshipType, ) *AccountRelationship`

NewAccountRelationship instantiates a new AccountRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRelationshipWithDefaults

`func NewAccountRelationshipWithDefaults() *AccountRelationship`

NewAccountRelationshipWithDefaults instantiates a new AccountRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *AccountRelationship) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AccountRelationship) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AccountRelationship) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AccountRelationship) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AccountRelationship) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AccountRelationship) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AccountRelationship) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AccountRelationship) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *AccountRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountRelationship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountRelationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPersonId

`func (o *AccountRelationship) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *AccountRelationship) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *AccountRelationship) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *AccountRelationship) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetRelationshipType

`func (o *AccountRelationship) GetRelationshipType() AccountRelationshipType`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *AccountRelationship) GetRelationshipTypeOk() (*AccountRelationshipType, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *AccountRelationship) SetRelationshipType(v AccountRelationshipType)`

SetRelationshipType sets RelationshipType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


