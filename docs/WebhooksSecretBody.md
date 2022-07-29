# WebhooksSecretBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRollingSecret** | Pointer to **bool** | Set true to let the current secret expire in the next 24 hours. Set false to let the current secret expire immediately. | [optional] 

## Methods

### NewWebhooksSecretBody

`func NewWebhooksSecretBody() *WebhooksSecretBody`

NewWebhooksSecretBody instantiates a new WebhooksSecretBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksSecretBodyWithDefaults

`func NewWebhooksSecretBodyWithDefaults() *WebhooksSecretBody`

NewWebhooksSecretBodyWithDefaults instantiates a new WebhooksSecretBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRollingSecret

`func (o *WebhooksSecretBody) GetIsRollingSecret() bool`

GetIsRollingSecret returns the IsRollingSecret field if non-nil, zero value otherwise.

### GetIsRollingSecretOk

`func (o *WebhooksSecretBody) GetIsRollingSecretOk() (*bool, bool)`

GetIsRollingSecretOk returns a tuple with the IsRollingSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRollingSecret

`func (o *WebhooksSecretBody) SetIsRollingSecret(v bool)`

SetIsRollingSecret sets IsRollingSecret field to given value.

### HasIsRollingSecret

`func (o *WebhooksSecretBody) HasIsRollingSecret() bool`

HasIsRollingSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

