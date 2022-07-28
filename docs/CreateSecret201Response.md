# CreateSecret201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | Generated secret. Do not share. This secret will be used to verify that webhook requests were sent from Synctera. | [optional] 

## Methods

### NewCreateSecret201Response

`func NewCreateSecret201Response() *CreateSecret201Response`

NewCreateSecret201Response instantiates a new CreateSecret201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecret201ResponseWithDefaults

`func NewCreateSecret201ResponseWithDefaults() *CreateSecret201Response`

NewCreateSecret201ResponseWithDefaults instantiates a new CreateSecret201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *CreateSecret201Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateSecret201Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateSecret201Response) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CreateSecret201Response) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)


