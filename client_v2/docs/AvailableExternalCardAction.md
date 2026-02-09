# AvailableExternalCardAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ExternalCardAction**](ExternalCardAction.md) |  | 
**State** | [**ExternalCardActionState**](ExternalCardActionState.md) |  | 
**TimestampValidTo** | Pointer to **time.Time** | The time by which the action must be taken against the dispute. * Only applicable for time constrained actions.  | [optional] 

## Methods

### NewAvailableExternalCardAction

`func NewAvailableExternalCardAction(action ExternalCardAction, state ExternalCardActionState, ) *AvailableExternalCardAction`

NewAvailableExternalCardAction instantiates a new AvailableExternalCardAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableExternalCardActionWithDefaults

`func NewAvailableExternalCardActionWithDefaults() *AvailableExternalCardAction`

NewAvailableExternalCardActionWithDefaults instantiates a new AvailableExternalCardAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AvailableExternalCardAction) GetAction() ExternalCardAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AvailableExternalCardAction) GetActionOk() (*ExternalCardAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AvailableExternalCardAction) SetAction(v ExternalCardAction)`

SetAction sets Action field to given value.


### GetState

`func (o *AvailableExternalCardAction) GetState() ExternalCardActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AvailableExternalCardAction) GetStateOk() (*ExternalCardActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AvailableExternalCardAction) SetState(v ExternalCardActionState)`

SetState sets State field to given value.


### GetTimestampValidTo

`func (o *AvailableExternalCardAction) GetTimestampValidTo() time.Time`

GetTimestampValidTo returns the TimestampValidTo field if non-nil, zero value otherwise.

### GetTimestampValidToOk

`func (o *AvailableExternalCardAction) GetTimestampValidToOk() (*time.Time, bool)`

GetTimestampValidToOk returns a tuple with the TimestampValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampValidTo

`func (o *AvailableExternalCardAction) SetTimestampValidTo(v time.Time)`

SetTimestampValidTo sets TimestampValidTo field to given value.

### HasTimestampValidTo

`func (o *AvailableExternalCardAction) HasTimestampValidTo() bool`

HasTimestampValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


