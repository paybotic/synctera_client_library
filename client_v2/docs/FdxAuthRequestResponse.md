# FdxAuthRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | The UUID of the business associated with the FDX token. One of customer_id or business_id must be provided.  | [optional] 
**Code** | Pointer to **string** | authorization code returned to Plaid after user completes auth steps  | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the authorization request was created. | [optional] [readonly] 
**CustomerId** | Pointer to **string** | The UUID of the customer associated with the FDX token. One of customer_id or business_id must be provided.  | [optional] 
**ExpiryTime** | Pointer to **time.Time** | The date and time the authorization request will expire. | [optional] [readonly] 
**Id** | **string** | The ID of the FDX authorization request.  | 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the authorization request was last udpated. | [optional] [readonly] 
**Oauth2State** | Pointer to **string** | OAuth2 state, an opaque string to be returned to Plaid  | [optional] 
**RedirectUri** | Pointer to **string** | URI to redirect to after successful authorization.  | [optional] 
**Status** | [**FdxAuthRequestStatus**](FdxAuthRequestStatus.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 

## Methods

### NewFdxAuthRequestResponse

`func NewFdxAuthRequestResponse(id string, status FdxAuthRequestStatus, ) *FdxAuthRequestResponse`

NewFdxAuthRequestResponse instantiates a new FdxAuthRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdxAuthRequestResponseWithDefaults

`func NewFdxAuthRequestResponseWithDefaults() *FdxAuthRequestResponse`

NewFdxAuthRequestResponseWithDefaults instantiates a new FdxAuthRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *FdxAuthRequestResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *FdxAuthRequestResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *FdxAuthRequestResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *FdxAuthRequestResponse) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCode

`func (o *FdxAuthRequestResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FdxAuthRequestResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FdxAuthRequestResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FdxAuthRequestResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreationTime

`func (o *FdxAuthRequestResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *FdxAuthRequestResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *FdxAuthRequestResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *FdxAuthRequestResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCustomerId

`func (o *FdxAuthRequestResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *FdxAuthRequestResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *FdxAuthRequestResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *FdxAuthRequestResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetExpiryTime

`func (o *FdxAuthRequestResponse) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *FdxAuthRequestResponse) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *FdxAuthRequestResponse) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *FdxAuthRequestResponse) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetId

`func (o *FdxAuthRequestResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FdxAuthRequestResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FdxAuthRequestResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *FdxAuthRequestResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *FdxAuthRequestResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *FdxAuthRequestResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *FdxAuthRequestResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetOauth2State

`func (o *FdxAuthRequestResponse) GetOauth2State() string`

GetOauth2State returns the Oauth2State field if non-nil, zero value otherwise.

### GetOauth2StateOk

`func (o *FdxAuthRequestResponse) GetOauth2StateOk() (*string, bool)`

GetOauth2StateOk returns a tuple with the Oauth2State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2State

`func (o *FdxAuthRequestResponse) SetOauth2State(v string)`

SetOauth2State sets Oauth2State field to given value.

### HasOauth2State

`func (o *FdxAuthRequestResponse) HasOauth2State() bool`

HasOauth2State returns a boolean if a field has been set.

### GetRedirectUri

`func (o *FdxAuthRequestResponse) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *FdxAuthRequestResponse) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *FdxAuthRequestResponse) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *FdxAuthRequestResponse) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetStatus

`func (o *FdxAuthRequestResponse) GetStatus() FdxAuthRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FdxAuthRequestResponse) GetStatusOk() (*FdxAuthRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FdxAuthRequestResponse) SetStatus(v FdxAuthRequestStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *FdxAuthRequestResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *FdxAuthRequestResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *FdxAuthRequestResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *FdxAuthRequestResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


