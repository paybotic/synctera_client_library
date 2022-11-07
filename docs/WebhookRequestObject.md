# WebhookRequestObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventResource** | Pointer to **string** | Json string of object associated with the event. For example, if your event is ACCOUNT.CREATED, You can refer to Acccount to parse the account event to obtain the ID, status etc.  | [optional] 
**EventResourceChangedFields** | Pointer to **string** | Json string of object associated with the event related to a resource change. This only contains those fields that have value changed on the event, and the field values are prior to the resource change event.  | [optional] 
**EventTime** | **time.Time** | Timestamp of the current event raised | 
**Id** | **string** | The unique ID of the current event | 
**Metadata** | **string** | Metadata that stored in the webhook subscription | 
**ResponseHistory** | Pointer to [**[]ResponseHistoryItem**](ResponseHistoryItem.md) | Response history of the webhook request | [optional] 
**Status** | Pointer to **string** | Current event status. Failing event will keep retry until it is purged. | [optional] 
**Type** | [**EventTypeExplicit**](EventTypeExplicit.md) |  | 
**Url** | **string** | URL that you specified for the webhook and where this request will be sent | 
**WebhookId** | **string** | Id of the Webhook the current request belongs to | 

## Methods

### NewWebhookRequestObject

`func NewWebhookRequestObject(eventTime time.Time, id string, metadata string, type_ EventTypeExplicit, url string, webhookId string, ) *WebhookRequestObject`

NewWebhookRequestObject instantiates a new WebhookRequestObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRequestObjectWithDefaults

`func NewWebhookRequestObjectWithDefaults() *WebhookRequestObject`

NewWebhookRequestObjectWithDefaults instantiates a new WebhookRequestObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventResource

`func (o *WebhookRequestObject) GetEventResource() string`

GetEventResource returns the EventResource field if non-nil, zero value otherwise.

### GetEventResourceOk

`func (o *WebhookRequestObject) GetEventResourceOk() (*string, bool)`

GetEventResourceOk returns a tuple with the EventResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventResource

`func (o *WebhookRequestObject) SetEventResource(v string)`

SetEventResource sets EventResource field to given value.

### HasEventResource

`func (o *WebhookRequestObject) HasEventResource() bool`

HasEventResource returns a boolean if a field has been set.

### GetEventResourceChangedFields

`func (o *WebhookRequestObject) GetEventResourceChangedFields() string`

GetEventResourceChangedFields returns the EventResourceChangedFields field if non-nil, zero value otherwise.

### GetEventResourceChangedFieldsOk

`func (o *WebhookRequestObject) GetEventResourceChangedFieldsOk() (*string, bool)`

GetEventResourceChangedFieldsOk returns a tuple with the EventResourceChangedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventResourceChangedFields

`func (o *WebhookRequestObject) SetEventResourceChangedFields(v string)`

SetEventResourceChangedFields sets EventResourceChangedFields field to given value.

### HasEventResourceChangedFields

`func (o *WebhookRequestObject) HasEventResourceChangedFields() bool`

HasEventResourceChangedFields returns a boolean if a field has been set.

### GetEventTime

`func (o *WebhookRequestObject) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *WebhookRequestObject) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *WebhookRequestObject) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetId

`func (o *WebhookRequestObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookRequestObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookRequestObject) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *WebhookRequestObject) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebhookRequestObject) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebhookRequestObject) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetResponseHistory

`func (o *WebhookRequestObject) GetResponseHistory() []ResponseHistoryItem`

GetResponseHistory returns the ResponseHistory field if non-nil, zero value otherwise.

### GetResponseHistoryOk

`func (o *WebhookRequestObject) GetResponseHistoryOk() (*[]ResponseHistoryItem, bool)`

GetResponseHistoryOk returns a tuple with the ResponseHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHistory

`func (o *WebhookRequestObject) SetResponseHistory(v []ResponseHistoryItem)`

SetResponseHistory sets ResponseHistory field to given value.

### HasResponseHistory

`func (o *WebhookRequestObject) HasResponseHistory() bool`

HasResponseHistory returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookRequestObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookRequestObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookRequestObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookRequestObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *WebhookRequestObject) GetType() EventTypeExplicit`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookRequestObject) GetTypeOk() (*EventTypeExplicit, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookRequestObject) SetType(v EventTypeExplicit)`

SetType sets Type field to given value.


### GetUrl

`func (o *WebhookRequestObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookRequestObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookRequestObject) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWebhookId

`func (o *WebhookRequestObject) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookRequestObject) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookRequestObject) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


