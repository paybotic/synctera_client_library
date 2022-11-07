/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// WebhookRequestObject Request body for webhook http request
type WebhookRequestObject struct {
	// Json string of object associated with the event. For example, if your event is ACCOUNT.CREATED, You can refer to Acccount to parse the account event to obtain the ID, status etc. 
	EventResource *string `json:"event_resource,omitempty"`
	// Json string of object associated with the event related to a resource change. This only contains those fields that have value changed on the event, and the field values are prior to the resource change event. 
	EventResourceChangedFields *string `json:"event_resource_changed_fields,omitempty"`
	// Timestamp of the current event raised
	EventTime time.Time `json:"event_time"`
	// The unique ID of the current event
	Id string `json:"id"`
	// Metadata that stored in the webhook subscription
	Metadata string `json:"metadata"`
	// Response history of the webhook request
	// Deprecated
	ResponseHistory []ResponseHistoryItem `json:"response_history,omitempty"`
	// Current event status. Failing event will keep retry until it is purged.
	// Deprecated
	Status *string `json:"status,omitempty"`
	Type EventTypeExplicit `json:"type"`
	// URL that you specified for the webhook and where this request will be sent
	Url string `json:"url"`
	// Id of the Webhook the current request belongs to
	WebhookId string `json:"webhook_id"`
}

// NewWebhookRequestObject instantiates a new WebhookRequestObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookRequestObject(eventTime time.Time, id string, metadata string, type_ EventTypeExplicit, url string, webhookId string) *WebhookRequestObject {
	this := WebhookRequestObject{}
	this.EventTime = eventTime
	this.Id = id
	this.Metadata = metadata
	this.Type = type_
	this.Url = url
	this.WebhookId = webhookId
	return &this
}

// NewWebhookRequestObjectWithDefaults instantiates a new WebhookRequestObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookRequestObjectWithDefaults() *WebhookRequestObject {
	this := WebhookRequestObject{}
	return &this
}

// GetEventResource returns the EventResource field value if set, zero value otherwise.
func (o *WebhookRequestObject) GetEventResource() string {
	if o == nil || o.EventResource == nil {
		var ret string
		return ret
	}
	return *o.EventResource
}

// GetEventResourceOk returns a tuple with the EventResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequestObject) GetEventResourceOk() (*string, bool) {
	if o == nil || o.EventResource == nil {
		return nil, false
	}
	return o.EventResource, true
}

// HasEventResource returns a boolean if a field has been set.
func (o *WebhookRequestObject) HasEventResource() bool {
	if o != nil && o.EventResource != nil {
		return true
	}

	return false
}

// SetEventResource gets a reference to the given string and assigns it to the EventResource field.
func (o *WebhookRequestObject) SetEventResource(v string) {
	o.EventResource = &v
}

// GetEventResourceChangedFields returns the EventResourceChangedFields field value if set, zero value otherwise.
func (o *WebhookRequestObject) GetEventResourceChangedFields() string {
	if o == nil || o.EventResourceChangedFields == nil {
		var ret string
		return ret
	}
	return *o.EventResourceChangedFields
}

// GetEventResourceChangedFieldsOk returns a tuple with the EventResourceChangedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequestObject) GetEventResourceChangedFieldsOk() (*string, bool) {
	if o == nil || o.EventResourceChangedFields == nil {
		return nil, false
	}
	return o.EventResourceChangedFields, true
}

// HasEventResourceChangedFields returns a boolean if a field has been set.
func (o *WebhookRequestObject) HasEventResourceChangedFields() bool {
	if o != nil && o.EventResourceChangedFields != nil {
		return true
	}

	return false
}

// SetEventResourceChangedFields gets a reference to the given string and assigns it to the EventResourceChangedFields field.
func (o *WebhookRequestObject) SetEventResourceChangedFields(v string) {
	o.EventResourceChangedFields = &v
}

// GetEventTime returns the EventTime field value
func (o *WebhookRequestObject) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *WebhookRequestObject) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *WebhookRequestObject) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetId returns the Id field value
func (o *WebhookRequestObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookRequestObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookRequestObject) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value
func (o *WebhookRequestObject) GetMetadata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *WebhookRequestObject) GetMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *WebhookRequestObject) SetMetadata(v string) {
	o.Metadata = v
}

// GetResponseHistory returns the ResponseHistory field value if set, zero value otherwise.
// Deprecated
func (o *WebhookRequestObject) GetResponseHistory() []ResponseHistoryItem {
	if o == nil || o.ResponseHistory == nil {
		var ret []ResponseHistoryItem
		return ret
	}
	return o.ResponseHistory
}

// GetResponseHistoryOk returns a tuple with the ResponseHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *WebhookRequestObject) GetResponseHistoryOk() ([]ResponseHistoryItem, bool) {
	if o == nil || o.ResponseHistory == nil {
		return nil, false
	}
	return o.ResponseHistory, true
}

// HasResponseHistory returns a boolean if a field has been set.
func (o *WebhookRequestObject) HasResponseHistory() bool {
	if o != nil && o.ResponseHistory != nil {
		return true
	}

	return false
}

// SetResponseHistory gets a reference to the given []ResponseHistoryItem and assigns it to the ResponseHistory field.
// Deprecated
func (o *WebhookRequestObject) SetResponseHistory(v []ResponseHistoryItem) {
	o.ResponseHistory = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
// Deprecated
func (o *WebhookRequestObject) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *WebhookRequestObject) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhookRequestObject) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
// Deprecated
func (o *WebhookRequestObject) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value
func (o *WebhookRequestObject) GetType() EventTypeExplicit {
	if o == nil {
		var ret EventTypeExplicit
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookRequestObject) GetTypeOk() (*EventTypeExplicit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookRequestObject) SetType(v EventTypeExplicit) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *WebhookRequestObject) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookRequestObject) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookRequestObject) SetUrl(v string) {
	o.Url = v
}

// GetWebhookId returns the WebhookId field value
func (o *WebhookRequestObject) GetWebhookId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value
// and a boolean to check if the value has been set.
func (o *WebhookRequestObject) GetWebhookIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookId, true
}

// SetWebhookId sets field value
func (o *WebhookRequestObject) SetWebhookId(v string) {
	o.WebhookId = v
}

func (o WebhookRequestObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventResource != nil {
		toSerialize["event_resource"] = o.EventResource
	}
	if o.EventResourceChangedFields != nil {
		toSerialize["event_resource_changed_fields"] = o.EventResourceChangedFields
	}
	if true {
		toSerialize["event_time"] = o.EventTime
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ResponseHistory != nil {
		toSerialize["response_history"] = o.ResponseHistory
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["webhook_id"] = o.WebhookId
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookRequestObject struct {
	value *WebhookRequestObject
	isSet bool
}

func (v NullableWebhookRequestObject) Get() *WebhookRequestObject {
	return v.value
}

func (v *NullableWebhookRequestObject) Set(val *WebhookRequestObject) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookRequestObject) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookRequestObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookRequestObject(val *WebhookRequestObject) *NullableWebhookRequestObject {
	return &NullableWebhookRequestObject{value: val, isSet: true}
}

func (v NullableWebhookRequestObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookRequestObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


