# WebhooksTriggerBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**EventType1**](EventType1.md) |  | [optional] 

## Methods

### NewWebhooksTriggerBody

`func NewWebhooksTriggerBody() *WebhooksTriggerBody`

NewWebhooksTriggerBody instantiates a new WebhooksTriggerBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksTriggerBodyWithDefaults

`func NewWebhooksTriggerBodyWithDefaults() *WebhooksTriggerBody`

NewWebhooksTriggerBodyWithDefaults instantiates a new WebhooksTriggerBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhooksTriggerBody) GetEvent() EventType1`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhooksTriggerBody) GetEventOk() (*EventType1, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhooksTriggerBody) SetEvent(v EventType1)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WebhooksTriggerBody) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


