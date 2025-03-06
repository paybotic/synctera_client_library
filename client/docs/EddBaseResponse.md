# EddBaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | **time.Time** |  | [readonly] 
**DeletionTime** | **NullableTime** |  | [readonly] 
**Id** | **string** | EDD record unique identifier | [readonly] 

## Methods

### NewEddBaseResponse

`func NewEddBaseResponse(creationTime time.Time, deletionTime NullableTime, id string, ) *EddBaseResponse`

NewEddBaseResponse instantiates a new EddBaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEddBaseResponseWithDefaults

`func NewEddBaseResponseWithDefaults() *EddBaseResponse`

NewEddBaseResponseWithDefaults instantiates a new EddBaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *EddBaseResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *EddBaseResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *EddBaseResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetDeletionTime

`func (o *EddBaseResponse) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *EddBaseResponse) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *EddBaseResponse) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.


### SetDeletionTimeNil

`func (o *EddBaseResponse) SetDeletionTimeNil(b bool)`

 SetDeletionTimeNil sets the value for DeletionTime to be an explicit nil

### UnsetDeletionTime
`func (o *EddBaseResponse) UnsetDeletionTime()`

UnsetDeletionTime ensures that no value is present for DeletionTime, not even an explicit nil
### GetId

`func (o *EddBaseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EddBaseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EddBaseResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


