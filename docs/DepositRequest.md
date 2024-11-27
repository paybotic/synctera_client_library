# DepositRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The ID of the account | [optional] 
**BackImageId** | Pointer to **string** | ID of the uploaded image of the back of the check | [optional] 
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**CheckAmount** | Pointer to **int32** | Amount on check in ISO 4217 minor currency units | [optional] 
**DepositCurrency** | Pointer to **string** | ISO 4217 currency code for the deposit amount | [optional] 
**FrontImageId** | Pointer to **string** | ID of the uploaded image of the front of the check | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 

## Methods

### NewDepositRequest

`func NewDepositRequest() *DepositRequest`

NewDepositRequest instantiates a new DepositRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositRequestWithDefaults

`func NewDepositRequestWithDefaults() *DepositRequest`

NewDepositRequestWithDefaults instantiates a new DepositRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DepositRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DepositRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DepositRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *DepositRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBackImageId

`func (o *DepositRequest) GetBackImageId() string`

GetBackImageId returns the BackImageId field if non-nil, zero value otherwise.

### GetBackImageIdOk

`func (o *DepositRequest) GetBackImageIdOk() (*string, bool)`

GetBackImageIdOk returns a tuple with the BackImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackImageId

`func (o *DepositRequest) SetBackImageId(v string)`

SetBackImageId sets BackImageId field to given value.

### HasBackImageId

`func (o *DepositRequest) HasBackImageId() bool`

HasBackImageId returns a boolean if a field has been set.

### GetBusinessId

`func (o *DepositRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *DepositRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *DepositRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *DepositRequest) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCheckAmount

`func (o *DepositRequest) GetCheckAmount() int32`

GetCheckAmount returns the CheckAmount field if non-nil, zero value otherwise.

### GetCheckAmountOk

`func (o *DepositRequest) GetCheckAmountOk() (*int32, bool)`

GetCheckAmountOk returns a tuple with the CheckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAmount

`func (o *DepositRequest) SetCheckAmount(v int32)`

SetCheckAmount sets CheckAmount field to given value.

### HasCheckAmount

`func (o *DepositRequest) HasCheckAmount() bool`

HasCheckAmount returns a boolean if a field has been set.

### GetDepositCurrency

`func (o *DepositRequest) GetDepositCurrency() string`

GetDepositCurrency returns the DepositCurrency field if non-nil, zero value otherwise.

### GetDepositCurrencyOk

`func (o *DepositRequest) GetDepositCurrencyOk() (*string, bool)`

GetDepositCurrencyOk returns a tuple with the DepositCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositCurrency

`func (o *DepositRequest) SetDepositCurrency(v string)`

SetDepositCurrency sets DepositCurrency field to given value.

### HasDepositCurrency

`func (o *DepositRequest) HasDepositCurrency() bool`

HasDepositCurrency returns a boolean if a field has been set.

### GetFrontImageId

`func (o *DepositRequest) GetFrontImageId() string`

GetFrontImageId returns the FrontImageId field if non-nil, zero value otherwise.

### GetFrontImageIdOk

`func (o *DepositRequest) GetFrontImageIdOk() (*string, bool)`

GetFrontImageIdOk returns a tuple with the FrontImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImageId

`func (o *DepositRequest) SetFrontImageId(v string)`

SetFrontImageId sets FrontImageId field to given value.

### HasFrontImageId

`func (o *DepositRequest) HasFrontImageId() bool`

HasFrontImageId returns a boolean if a field has been set.

### GetMetadata

`func (o *DepositRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DepositRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DepositRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DepositRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPersonId

`func (o *DepositRequest) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *DepositRequest) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *DepositRequest) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *DepositRequest) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


