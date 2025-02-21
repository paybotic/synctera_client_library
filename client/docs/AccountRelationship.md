# AccountRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | Business associated with the current account. One of business_id or customer_id must be specified. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when this association was created. | [optional] 
**CustomerId** | Pointer to **string** | Personal customer associated with the current account. One of customer_id or business_id must be specified. | [optional] 
**DeletedAt** | Pointer to **time.Time** | Date and time when this association was deleted. | [optional] 
**Id** | Pointer to **string** | ID of account relationship | [optional] [readonly] 
**PersonId** | Pointer to **string** | Person associated with the current account. This attribute is deprecated and will be removed in a future API version. Use customer_id instead. | [optional] 
**RelationshipType** | [**AccountRelationshipType**](AccountRelationshipType.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | Date and time when this association was last updated. | [optional] 

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

### GetCreatedAt

`func (o *AccountRelationship) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountRelationship) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountRelationship) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountRelationship) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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

### GetDeletedAt

`func (o *AccountRelationship) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AccountRelationship) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AccountRelationship) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AccountRelationship) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

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


### GetUpdatedAt

`func (o *AccountRelationship) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountRelationship) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountRelationship) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountRelationship) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


