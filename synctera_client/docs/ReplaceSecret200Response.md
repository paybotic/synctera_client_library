# ReplaceSecret200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteAt** | Pointer to **time.Time** | Timestamp that the old secret is delete | [optional] 
**Secret** | Pointer to **string** | Generated secret. Do not share. This secret will be used to verify that webhook requests were sent from Synctera. | [optional] 

## Methods

### NewReplaceSecret200Response

`func NewReplaceSecret200Response() *ReplaceSecret200Response`

NewReplaceSecret200Response instantiates a new ReplaceSecret200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceSecret200ResponseWithDefaults

`func NewReplaceSecret200ResponseWithDefaults() *ReplaceSecret200Response`

NewReplaceSecret200ResponseWithDefaults instantiates a new ReplaceSecret200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteAt

`func (o *ReplaceSecret200Response) GetDeleteAt() time.Time`

GetDeleteAt returns the DeleteAt field if non-nil, zero value otherwise.

### GetDeleteAtOk

`func (o *ReplaceSecret200Response) GetDeleteAtOk() (*time.Time, bool)`

GetDeleteAtOk returns a tuple with the DeleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAt

`func (o *ReplaceSecret200Response) SetDeleteAt(v time.Time)`

SetDeleteAt sets DeleteAt field to given value.

### HasDeleteAt

`func (o *ReplaceSecret200Response) HasDeleteAt() bool`

HasDeleteAt returns a boolean if a field has been set.

### GetSecret

`func (o *ReplaceSecret200Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ReplaceSecret200Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ReplaceSecret200Response) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ReplaceSecret200Response) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)


