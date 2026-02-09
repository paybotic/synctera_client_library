# DocumentVerificationSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionToken** | Pointer to **string** | ID of the document verification session to be used in &#x60;/verifications/verify&#x60; request. | [optional] 
**Url** | Pointer to **string** | URL of the document verification session which can be shared with the end-customer. | [optional] 

## Methods

### NewDocumentVerificationSessionResponse

`func NewDocumentVerificationSessionResponse() *DocumentVerificationSessionResponse`

NewDocumentVerificationSessionResponse instantiates a new DocumentVerificationSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentVerificationSessionResponseWithDefaults

`func NewDocumentVerificationSessionResponseWithDefaults() *DocumentVerificationSessionResponse`

NewDocumentVerificationSessionResponseWithDefaults instantiates a new DocumentVerificationSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionToken

`func (o *DocumentVerificationSessionResponse) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *DocumentVerificationSessionResponse) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *DocumentVerificationSessionResponse) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *DocumentVerificationSessionResponse) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetUrl

`func (o *DocumentVerificationSessionResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DocumentVerificationSessionResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DocumentVerificationSessionResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DocumentVerificationSessionResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


