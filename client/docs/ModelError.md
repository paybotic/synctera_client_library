# ModelError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | An optional “sneak preview” of our future API v1 error responses. This is provided to give integrators a chance to work with our future error codes. Error codes for the same error may change between v0 and v1. | [optional] 
**Detail** | Pointer to **string** | A human-readable string explaining this particular error.  | [optional] 
**Status** | Pointer to **int32** | the HTTP status code for this response | [optional] 
**Title** | Pointer to **string** | A human-readable string for this general category of error, which corresponds 1-to-1 with error types (&#x60;title&#x60; is the human-readable version of &#x60;type&#x60;). There can be multiple distinct titles for the same HTTP status code, and the same &#x60;title&#x60; can result in many different &#x60;detail&#x60; strings. This field will be removed in API v1.  | [optional] 
**Type** | Pointer to **string** | A machine-readable string that identifies the error for programmatic use. This is a URI, i.e. a globally unique identifier. It is _not_ necessarily a URL, so do not expect it to resolve to a web page. You can use this whole string as an error code, or just everything after the last slash. This field will be removed in API v1.  | [optional] 

## Methods

### NewModelError

`func NewModelError() *ModelError`

NewModelError instantiates a new ModelError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelErrorWithDefaults

`func NewModelErrorWithDefaults() *ModelError`

NewModelErrorWithDefaults instantiates a new ModelError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ModelError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModelError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModelError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ModelError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *ModelError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ModelError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ModelError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ModelError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatus

`func (o *ModelError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *ModelError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ModelError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


