# AdverseActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**NotificationTime** | **time.Time** | The date and time the adverse action notice was sent. | 
**Purpose** | [**Purpose**](Purpose.md) |  | 
**Reasons** | **[]string** | Reasons (up to 5) provided to customers when adverse action is taken. | 
**RelatedResourceId** | **string** | Unique identifier for the related resource. | 
**RelatedResourceType** | [**RelatedResourceType**](RelatedResourceType.md) |  | 
**CreationTime** | **time.Time** | The date and time the resource was created. | [readonly] 
**Id** | **string** | Unique identifier for this adverse action notice. | [readonly] 
**LastUpdatedTime** | **time.Time** | The date and time the resource was last update. | [readonly] 

## Methods

### NewAdverseActionResponse

`func NewAdverseActionResponse(notificationTime time.Time, purpose Purpose, reasons []string, relatedResourceId string, relatedResourceType RelatedResourceType, creationTime time.Time, id string, lastUpdatedTime time.Time, ) *AdverseActionResponse`

NewAdverseActionResponse instantiates a new AdverseActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdverseActionResponseWithDefaults

`func NewAdverseActionResponseWithDefaults() *AdverseActionResponse`

NewAdverseActionResponseWithDefaults instantiates a new AdverseActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AdverseActionResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AdverseActionResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AdverseActionResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AdverseActionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotificationTime

`func (o *AdverseActionResponse) GetNotificationTime() time.Time`

GetNotificationTime returns the NotificationTime field if non-nil, zero value otherwise.

### GetNotificationTimeOk

`func (o *AdverseActionResponse) GetNotificationTimeOk() (*time.Time, bool)`

GetNotificationTimeOk returns a tuple with the NotificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTime

`func (o *AdverseActionResponse) SetNotificationTime(v time.Time)`

SetNotificationTime sets NotificationTime field to given value.


### GetPurpose

`func (o *AdverseActionResponse) GetPurpose() Purpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *AdverseActionResponse) GetPurposeOk() (*Purpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *AdverseActionResponse) SetPurpose(v Purpose)`

SetPurpose sets Purpose field to given value.


### GetReasons

`func (o *AdverseActionResponse) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *AdverseActionResponse) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *AdverseActionResponse) SetReasons(v []string)`

SetReasons sets Reasons field to given value.


### GetRelatedResourceId

`func (o *AdverseActionResponse) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *AdverseActionResponse) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *AdverseActionResponse) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.


### GetRelatedResourceType

`func (o *AdverseActionResponse) GetRelatedResourceType() RelatedResourceType`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *AdverseActionResponse) GetRelatedResourceTypeOk() (*RelatedResourceType, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *AdverseActionResponse) SetRelatedResourceType(v RelatedResourceType)`

SetRelatedResourceType sets RelatedResourceType field to given value.


### GetCreationTime

`func (o *AdverseActionResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AdverseActionResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AdverseActionResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *AdverseActionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdverseActionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdverseActionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *AdverseActionResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AdverseActionResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AdverseActionResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


