# RelationshipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | Business associated with the current account. One of business_id or customer_id must be specified. | [optional] 
**CustomerId** | Pointer to **string** | Personal customer associated with the current account. One of customer_id or business_id must be specified. | [optional] 
**Id** | Pointer to **string** | ID of account relationship | [optional] [readonly] 
**PersonId** | Pointer to **string** | Person associated with the current account. This attribute is deprecated and will be removed in a future API version. Use customer_id instead. | [optional] 
**RelationshipType** | [**AccountRelationshipType**](AccountRelationshipType.md) |  | 
**AccountId** | Pointer to **string** | Account ID | [optional] 

## Methods

### NewRelationshipResponse

`func NewRelationshipResponse(relationshipType AccountRelationshipType, ) *RelationshipResponse`

NewRelationshipResponse instantiates a new RelationshipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipResponseWithDefaults

`func NewRelationshipResponseWithDefaults() *RelationshipResponse`

NewRelationshipResponseWithDefaults instantiates a new RelationshipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *RelationshipResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *RelationshipResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *RelationshipResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *RelationshipResponse) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *RelationshipResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *RelationshipResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *RelationshipResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *RelationshipResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *RelationshipResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPersonId

`func (o *RelationshipResponse) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *RelationshipResponse) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *RelationshipResponse) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *RelationshipResponse) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetRelationshipType

`func (o *RelationshipResponse) GetRelationshipType() AccountRelationshipType`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *RelationshipResponse) GetRelationshipTypeOk() (*AccountRelationshipType, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *RelationshipResponse) SetRelationshipType(v AccountRelationshipType)`

SetRelationshipType sets RelationshipType field to given value.


### GetAccountId

`func (o *RelationshipResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RelationshipResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RelationshipResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *RelationshipResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


