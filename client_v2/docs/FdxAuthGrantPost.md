# FdxAuthGrantPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthRequestId** | **string** | The ID of the FDX authorization request.  | 
**BusinessId** | Pointer to **string** | The UUID of the business associated with the FDX token. One of customer_id or business_id must be provided.  | [optional] 
**CustomerId** | Pointer to **string** | The UUID of the customer associated with the FDX token. One of customer_id or business_id must be provided.  | [optional] 
**Status** | [**FdxAuthGrantStatus**](FdxAuthGrantStatus.md) |  | 

## Methods

### NewFdxAuthGrantPost

`func NewFdxAuthGrantPost(authRequestId string, status FdxAuthGrantStatus, ) *FdxAuthGrantPost`

NewFdxAuthGrantPost instantiates a new FdxAuthGrantPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdxAuthGrantPostWithDefaults

`func NewFdxAuthGrantPostWithDefaults() *FdxAuthGrantPost`

NewFdxAuthGrantPostWithDefaults instantiates a new FdxAuthGrantPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthRequestId

`func (o *FdxAuthGrantPost) GetAuthRequestId() string`

GetAuthRequestId returns the AuthRequestId field if non-nil, zero value otherwise.

### GetAuthRequestIdOk

`func (o *FdxAuthGrantPost) GetAuthRequestIdOk() (*string, bool)`

GetAuthRequestIdOk returns a tuple with the AuthRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRequestId

`func (o *FdxAuthGrantPost) SetAuthRequestId(v string)`

SetAuthRequestId sets AuthRequestId field to given value.


### GetBusinessId

`func (o *FdxAuthGrantPost) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *FdxAuthGrantPost) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *FdxAuthGrantPost) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *FdxAuthGrantPost) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *FdxAuthGrantPost) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *FdxAuthGrantPost) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *FdxAuthGrantPost) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *FdxAuthGrantPost) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetStatus

`func (o *FdxAuthGrantPost) GetStatus() FdxAuthGrantStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FdxAuthGrantPost) GetStatusOk() (*FdxAuthGrantStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FdxAuthGrantPost) SetStatus(v FdxAuthGrantStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


