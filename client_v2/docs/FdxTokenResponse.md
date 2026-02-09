# FdxTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | The UUID of the business associated with the FDX token. One of customer_id or business_id must be provided.  | [optional] 
**CreationTime** | **time.Time** | The date and time the token was created. | [readonly] 
**CustomerId** | Pointer to **string** | The UUID of the customer associated with the FDX token. One of customer_id or business_id must be provided.  | [optional] 
**Id** | **string** | FDX token ID | 
**LastUpdatedTime** | **time.Time** | The date and time the token was last updated. | [readonly] 
**ParentTokenId** | Pointer to **string** | If the token was created via refresh, this is the ID of the refreshing token.  | [optional] 
**Status** | [**FdxTokenStatus**](FdxTokenStatus.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**TokenExpiryTime** | **time.Time** | The date and time the token expires. | [readonly] 
**TokenHash** | **string** | The non-secret hash of the FDX token.  | 
**TokenType** | [**FdxTokenType**](FdxTokenType.md) |  | 

## Methods

### NewFdxTokenResponse

`func NewFdxTokenResponse(creationTime time.Time, id string, lastUpdatedTime time.Time, status FdxTokenStatus, tenant string, tokenExpiryTime time.Time, tokenHash string, tokenType FdxTokenType, ) *FdxTokenResponse`

NewFdxTokenResponse instantiates a new FdxTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdxTokenResponseWithDefaults

`func NewFdxTokenResponseWithDefaults() *FdxTokenResponse`

NewFdxTokenResponseWithDefaults instantiates a new FdxTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *FdxTokenResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *FdxTokenResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *FdxTokenResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *FdxTokenResponse) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *FdxTokenResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *FdxTokenResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *FdxTokenResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomerId

`func (o *FdxTokenResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *FdxTokenResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *FdxTokenResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *FdxTokenResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *FdxTokenResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FdxTokenResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FdxTokenResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *FdxTokenResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *FdxTokenResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *FdxTokenResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetParentTokenId

`func (o *FdxTokenResponse) GetParentTokenId() string`

GetParentTokenId returns the ParentTokenId field if non-nil, zero value otherwise.

### GetParentTokenIdOk

`func (o *FdxTokenResponse) GetParentTokenIdOk() (*string, bool)`

GetParentTokenIdOk returns a tuple with the ParentTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTokenId

`func (o *FdxTokenResponse) SetParentTokenId(v string)`

SetParentTokenId sets ParentTokenId field to given value.

### HasParentTokenId

`func (o *FdxTokenResponse) HasParentTokenId() bool`

HasParentTokenId returns a boolean if a field has been set.

### GetStatus

`func (o *FdxTokenResponse) GetStatus() FdxTokenStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FdxTokenResponse) GetStatusOk() (*FdxTokenStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FdxTokenResponse) SetStatus(v FdxTokenStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *FdxTokenResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *FdxTokenResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *FdxTokenResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTokenExpiryTime

`func (o *FdxTokenResponse) GetTokenExpiryTime() time.Time`

GetTokenExpiryTime returns the TokenExpiryTime field if non-nil, zero value otherwise.

### GetTokenExpiryTimeOk

`func (o *FdxTokenResponse) GetTokenExpiryTimeOk() (*time.Time, bool)`

GetTokenExpiryTimeOk returns a tuple with the TokenExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiryTime

`func (o *FdxTokenResponse) SetTokenExpiryTime(v time.Time)`

SetTokenExpiryTime sets TokenExpiryTime field to given value.


### GetTokenHash

`func (o *FdxTokenResponse) GetTokenHash() string`

GetTokenHash returns the TokenHash field if non-nil, zero value otherwise.

### GetTokenHashOk

`func (o *FdxTokenResponse) GetTokenHashOk() (*string, bool)`

GetTokenHashOk returns a tuple with the TokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenHash

`func (o *FdxTokenResponse) SetTokenHash(v string)`

SetTokenHash sets TokenHash field to given value.


### GetTokenType

`func (o *FdxTokenResponse) GetTokenType() FdxTokenType`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *FdxTokenResponse) GetTokenTypeOk() (*FdxTokenType, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *FdxTokenResponse) SetTokenType(v FdxTokenType)`

SetTokenType sets TokenType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


