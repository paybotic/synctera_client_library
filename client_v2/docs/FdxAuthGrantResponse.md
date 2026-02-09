# FdxAuthGrantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUri** | **string** | URI to redirect to after successful authorization.  | 

## Methods

### NewFdxAuthGrantResponse

`func NewFdxAuthGrantResponse(redirectUri string, ) *FdxAuthGrantResponse`

NewFdxAuthGrantResponse instantiates a new FdxAuthGrantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdxAuthGrantResponseWithDefaults

`func NewFdxAuthGrantResponseWithDefaults() *FdxAuthGrantResponse`

NewFdxAuthGrantResponseWithDefaults instantiates a new FdxAuthGrantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUri

`func (o *FdxAuthGrantResponse) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *FdxAuthGrantResponse) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *FdxAuthGrantResponse) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


