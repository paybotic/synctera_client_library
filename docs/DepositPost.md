# DepositPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The ID of the account | 
**BackImageId** | **string** | ID of the uploaded image of the back of the check | 
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**CheckAmount** | **int32** | Amount on check in ISO 4217 minor currency units | 
**DepositCurrency** | **string** | ISO 4217 currency code for the deposit amount | 
**FrontImageId** | **string** | ID of the uploaded image of the front of the check | 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 

## Methods

### NewDepositPost

`func NewDepositPost(accountId string, backImageId string, checkAmount int32, depositCurrency string, frontImageId string, ) *DepositPost`

NewDepositPost instantiates a new DepositPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositPostWithDefaults

`func NewDepositPostWithDefaults() *DepositPost`

NewDepositPostWithDefaults instantiates a new DepositPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DepositPost) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DepositPost) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DepositPost) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBackImageId

`func (o *DepositPost) GetBackImageId() string`

GetBackImageId returns the BackImageId field if non-nil, zero value otherwise.

### GetBackImageIdOk

`func (o *DepositPost) GetBackImageIdOk() (*string, bool)`

GetBackImageIdOk returns a tuple with the BackImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackImageId

`func (o *DepositPost) SetBackImageId(v string)`

SetBackImageId sets BackImageId field to given value.


### GetBusinessId

`func (o *DepositPost) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *DepositPost) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *DepositPost) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *DepositPost) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCheckAmount

`func (o *DepositPost) GetCheckAmount() int32`

GetCheckAmount returns the CheckAmount field if non-nil, zero value otherwise.

### GetCheckAmountOk

`func (o *DepositPost) GetCheckAmountOk() (*int32, bool)`

GetCheckAmountOk returns a tuple with the CheckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAmount

`func (o *DepositPost) SetCheckAmount(v int32)`

SetCheckAmount sets CheckAmount field to given value.


### GetDepositCurrency

`func (o *DepositPost) GetDepositCurrency() string`

GetDepositCurrency returns the DepositCurrency field if non-nil, zero value otherwise.

### GetDepositCurrencyOk

`func (o *DepositPost) GetDepositCurrencyOk() (*string, bool)`

GetDepositCurrencyOk returns a tuple with the DepositCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositCurrency

`func (o *DepositPost) SetDepositCurrency(v string)`

SetDepositCurrency sets DepositCurrency field to given value.


### GetFrontImageId

`func (o *DepositPost) GetFrontImageId() string`

GetFrontImageId returns the FrontImageId field if non-nil, zero value otherwise.

### GetFrontImageIdOk

`func (o *DepositPost) GetFrontImageIdOk() (*string, bool)`

GetFrontImageIdOk returns a tuple with the FrontImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImageId

`func (o *DepositPost) SetFrontImageId(v string)`

SetFrontImageId sets FrontImageId field to given value.


### GetMetadata

`func (o *DepositPost) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DepositPost) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DepositPost) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DepositPost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPersonId

`func (o *DepositPost) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *DepositPost) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *DepositPost) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *DepositPost) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


