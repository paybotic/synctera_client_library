# RestrictedApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDetails** | [**RestrictedApplicationDetails**](RestrictedApplicationDetails.md) |  | 
**BusinessId** | Pointer to **string** | Business ID for the application. An application must have either a Business or a customer associated with it. | [optional] 
**CustomerId** | Pointer to **string** | Customer ID for the application. An application must have either a Business or a customer associated with it. | [optional] 
**Description** | Pointer to **string** | A description of the restricted account application | [optional] 
**Status** | [**RestrictedApplicationStatus**](RestrictedApplicationStatus.md) |  | 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 

## Methods

### NewRestrictedApplication

`func NewRestrictedApplication(applicationDetails RestrictedApplicationDetails, status RestrictedApplicationStatus, type_ ApplicationType, ) *RestrictedApplication`

NewRestrictedApplication instantiates a new RestrictedApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedApplicationWithDefaults

`func NewRestrictedApplicationWithDefaults() *RestrictedApplication`

NewRestrictedApplicationWithDefaults instantiates a new RestrictedApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationDetails

`func (o *RestrictedApplication) GetApplicationDetails() RestrictedApplicationDetails`

GetApplicationDetails returns the ApplicationDetails field if non-nil, zero value otherwise.

### GetApplicationDetailsOk

`func (o *RestrictedApplication) GetApplicationDetailsOk() (*RestrictedApplicationDetails, bool)`

GetApplicationDetailsOk returns a tuple with the ApplicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDetails

`func (o *RestrictedApplication) SetApplicationDetails(v RestrictedApplicationDetails)`

SetApplicationDetails sets ApplicationDetails field to given value.


### GetBusinessId

`func (o *RestrictedApplication) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *RestrictedApplication) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *RestrictedApplication) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *RestrictedApplication) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *RestrictedApplication) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *RestrictedApplication) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *RestrictedApplication) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *RestrictedApplication) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *RestrictedApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestrictedApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestrictedApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RestrictedApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *RestrictedApplication) GetStatus() RestrictedApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestrictedApplication) GetStatusOk() (*RestrictedApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestrictedApplication) SetStatus(v RestrictedApplicationStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *RestrictedApplication) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestrictedApplication) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestrictedApplication) SetType(v ApplicationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


