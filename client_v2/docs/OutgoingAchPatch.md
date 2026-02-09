# OutgoingAchPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundsAvailabilityTime** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOutgoingAchPatch

`func NewOutgoingAchPatch() *OutgoingAchPatch`

NewOutgoingAchPatch instantiates a new OutgoingAchPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingAchPatchWithDefaults

`func NewOutgoingAchPatchWithDefaults() *OutgoingAchPatch`

NewOutgoingAchPatchWithDefaults instantiates a new OutgoingAchPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundsAvailabilityTime

`func (o *OutgoingAchPatch) GetFundsAvailabilityTime() time.Time`

GetFundsAvailabilityTime returns the FundsAvailabilityTime field if non-nil, zero value otherwise.

### GetFundsAvailabilityTimeOk

`func (o *OutgoingAchPatch) GetFundsAvailabilityTimeOk() (*time.Time, bool)`

GetFundsAvailabilityTimeOk returns a tuple with the FundsAvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsAvailabilityTime

`func (o *OutgoingAchPatch) SetFundsAvailabilityTime(v time.Time)`

SetFundsAvailabilityTime sets FundsAvailabilityTime field to given value.

### HasFundsAvailabilityTime

`func (o *OutgoingAchPatch) HasFundsAvailabilityTime() bool`

HasFundsAvailabilityTime returns a boolean if a field has been set.

### SetFundsAvailabilityTimeNil

`func (o *OutgoingAchPatch) SetFundsAvailabilityTimeNil(b bool)`

 SetFundsAvailabilityTimeNil sets the value for FundsAvailabilityTime to be an explicit nil

### UnsetFundsAvailabilityTime
`func (o *OutgoingAchPatch) UnsetFundsAvailabilityTime()`

UnsetFundsAvailabilityTime ensures that no value is present for FundsAvailabilityTime, not even an explicit nil
### GetStatus

`func (o *OutgoingAchPatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutgoingAchPatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutgoingAchPatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OutgoingAchPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *OutgoingAchPatch) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *OutgoingAchPatch) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


