# AdverseAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**NotificationTime** | **time.Time** | The date and time the adverse action notice was sent. | 
**Purpose** | [**Purpose**](Purpose.md) |  | 
**Reasons** | **[]string** | Reasons (up to 5) provided to customers when adverse action is taken. | 
**RelatedResourceId** | **string** | Unique identifier for the related resource. | 
**RelatedResourceType** | [**RelatedResourceType**](RelatedResourceType.md) |  | 

## Methods

### NewAdverseAction

`func NewAdverseAction(notificationTime time.Time, purpose Purpose, reasons []string, relatedResourceId string, relatedResourceType RelatedResourceType, ) *AdverseAction`

NewAdverseAction instantiates a new AdverseAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdverseActionWithDefaults

`func NewAdverseActionWithDefaults() *AdverseAction`

NewAdverseActionWithDefaults instantiates a new AdverseAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AdverseAction) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AdverseAction) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AdverseAction) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AdverseAction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotificationTime

`func (o *AdverseAction) GetNotificationTime() time.Time`

GetNotificationTime returns the NotificationTime field if non-nil, zero value otherwise.

### GetNotificationTimeOk

`func (o *AdverseAction) GetNotificationTimeOk() (*time.Time, bool)`

GetNotificationTimeOk returns a tuple with the NotificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTime

`func (o *AdverseAction) SetNotificationTime(v time.Time)`

SetNotificationTime sets NotificationTime field to given value.


### GetPurpose

`func (o *AdverseAction) GetPurpose() Purpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *AdverseAction) GetPurposeOk() (*Purpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *AdverseAction) SetPurpose(v Purpose)`

SetPurpose sets Purpose field to given value.


### GetReasons

`func (o *AdverseAction) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *AdverseAction) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *AdverseAction) SetReasons(v []string)`

SetReasons sets Reasons field to given value.


### GetRelatedResourceId

`func (o *AdverseAction) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *AdverseAction) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *AdverseAction) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *AdverseAction) GetRelatedResourceType() RelatedResourceType`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *AdverseAction) GetRelatedResourceTypeOk() (*RelatedResourceType, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *AdverseAction) SetRelatedResourceType(v RelatedResourceType)`

SetRelatedResourceType sets RelatedResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


