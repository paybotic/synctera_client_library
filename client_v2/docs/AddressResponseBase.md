# AddressResponseBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | The identifier for the business customer associated with address.  | [optional] 
**CreationDate** | **time.Time** | Date and time when the address was created. | 
**CustomerId** | Pointer to **string** | The identifier for the personal customer associated with address.  | [optional] 
**IsActive** | **bool** | Whether the address is active or not | 
**LastUpdatedTime** | **time.Time** |  | 

## Methods

### NewAddressResponseBase

`func NewAddressResponseBase(creationDate time.Time, isActive bool, lastUpdatedTime time.Time, ) *AddressResponseBase`

NewAddressResponseBase instantiates a new AddressResponseBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressResponseBaseWithDefaults

`func NewAddressResponseBaseWithDefaults() *AddressResponseBase`

NewAddressResponseBaseWithDefaults instantiates a new AddressResponseBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *AddressResponseBase) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AddressResponseBase) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AddressResponseBase) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AddressResponseBase) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationDate

`func (o *AddressResponseBase) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AddressResponseBase) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AddressResponseBase) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.


### GetCustomerId

`func (o *AddressResponseBase) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AddressResponseBase) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AddressResponseBase) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AddressResponseBase) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetIsActive

`func (o *AddressResponseBase) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AddressResponseBase) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AddressResponseBase) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastUpdatedTime

`func (o *AddressResponseBase) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AddressResponseBase) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AddressResponseBase) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


