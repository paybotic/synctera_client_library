# RestrictedApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDetails** | [**RestrictedApplicationDetails**](RestrictedApplicationDetails.md) |  | 
**BusinessId** | Pointer to **string** | Business ID for the application. An application must have either a Business or a customer associated with it. | [optional] 
**CustomerId** | Pointer to **string** | Customer ID for the application. An application must have either a Business or a customer associated with it. | [optional] 
**Description** | Pointer to **string** | A description of the restricted account application | [optional] 
**Status** | [**RestrictedApplicationStatus**](RestrictedApplicationStatus.md) |  | 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 
**CreationTime** | **time.Time** | Application creation timestamp in RFC3339 format | [readonly] 
**Id** | **string** | Generated ID for the application | [readonly] 
**LastUpdatedTime** | **time.Time** | Timestamp of the last application modification in RFC3339 format | [readonly] 

## Methods

### NewRestrictedApplicationResponse

`func NewRestrictedApplicationResponse(applicationDetails RestrictedApplicationDetails, status RestrictedApplicationStatus, type_ ApplicationType, creationTime time.Time, id string, lastUpdatedTime time.Time, ) *RestrictedApplicationResponse`

NewRestrictedApplicationResponse instantiates a new RestrictedApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedApplicationResponseWithDefaults

`func NewRestrictedApplicationResponseWithDefaults() *RestrictedApplicationResponse`

NewRestrictedApplicationResponseWithDefaults instantiates a new RestrictedApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationDetails

`func (o *RestrictedApplicationResponse) GetApplicationDetails() RestrictedApplicationDetails`

GetApplicationDetails returns the ApplicationDetails field if non-nil, zero value otherwise.

### GetApplicationDetailsOk

`func (o *RestrictedApplicationResponse) GetApplicationDetailsOk() (*RestrictedApplicationDetails, bool)`

GetApplicationDetailsOk returns a tuple with the ApplicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDetails

`func (o *RestrictedApplicationResponse) SetApplicationDetails(v RestrictedApplicationDetails)`

SetApplicationDetails sets ApplicationDetails field to given value.


### GetBusinessId

`func (o *RestrictedApplicationResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *RestrictedApplicationResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *RestrictedApplicationResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *RestrictedApplicationResponse) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *RestrictedApplicationResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *RestrictedApplicationResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *RestrictedApplicationResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *RestrictedApplicationResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *RestrictedApplicationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestrictedApplicationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestrictedApplicationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RestrictedApplicationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *RestrictedApplicationResponse) GetStatus() RestrictedApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestrictedApplicationResponse) GetStatusOk() (*RestrictedApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestrictedApplicationResponse) SetStatus(v RestrictedApplicationStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *RestrictedApplicationResponse) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestrictedApplicationResponse) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestrictedApplicationResponse) SetType(v ApplicationType)`

SetType sets Type field to given value.


### GetCreationTime

`func (o *RestrictedApplicationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RestrictedApplicationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RestrictedApplicationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *RestrictedApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestrictedApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestrictedApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *RestrictedApplicationResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *RestrictedApplicationResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *RestrictedApplicationResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


