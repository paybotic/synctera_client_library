# RestrictedApplicationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDetails** | Pointer to [**RestrictedApplicationDetails**](RestrictedApplicationDetails.md) |  | [optional] 
**Description** | Pointer to **string** | A description of the restricted account application | [optional] 
**Status** | Pointer to [**RestrictedApplicationStatus**](RestrictedApplicationStatus.md) |  | [optional] 

## Methods

### NewRestrictedApplicationPatch

`func NewRestrictedApplicationPatch() *RestrictedApplicationPatch`

NewRestrictedApplicationPatch instantiates a new RestrictedApplicationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedApplicationPatchWithDefaults

`func NewRestrictedApplicationPatchWithDefaults() *RestrictedApplicationPatch`

NewRestrictedApplicationPatchWithDefaults instantiates a new RestrictedApplicationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationDetails

`func (o *RestrictedApplicationPatch) GetApplicationDetails() RestrictedApplicationDetails`

GetApplicationDetails returns the ApplicationDetails field if non-nil, zero value otherwise.

### GetApplicationDetailsOk

`func (o *RestrictedApplicationPatch) GetApplicationDetailsOk() (*RestrictedApplicationDetails, bool)`

GetApplicationDetailsOk returns a tuple with the ApplicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDetails

`func (o *RestrictedApplicationPatch) SetApplicationDetails(v RestrictedApplicationDetails)`

SetApplicationDetails sets ApplicationDetails field to given value.

### HasApplicationDetails

`func (o *RestrictedApplicationPatch) HasApplicationDetails() bool`

HasApplicationDetails returns a boolean if a field has been set.

### GetDescription

`func (o *RestrictedApplicationPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestrictedApplicationPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestrictedApplicationPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RestrictedApplicationPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *RestrictedApplicationPatch) GetStatus() RestrictedApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestrictedApplicationPatch) GetStatusOk() (*RestrictedApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestrictedApplicationPatch) SetStatus(v RestrictedApplicationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestrictedApplicationPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


