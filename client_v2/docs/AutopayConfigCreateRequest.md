# AutopayConfigCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | The business ID that owns this autopay configuration. Mutually exclusive with person_id - exactly one must be provided.  | [optional] 
**Config** | [**AutopayConfigData**](AutopayConfigData.md) |  | 
**LendingAccountId** | **string** | The lending account ID to configure autopay for | 
**PersonId** | Pointer to **string** | The person ID who owns this autopay configuration. Mutually exclusive with business_id - exactly one must be provided.  | [optional] 

## Methods

### NewAutopayConfigCreateRequest

`func NewAutopayConfigCreateRequest(config AutopayConfigData, lendingAccountId string, ) *AutopayConfigCreateRequest`

NewAutopayConfigCreateRequest instantiates a new AutopayConfigCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayConfigCreateRequestWithDefaults

`func NewAutopayConfigCreateRequestWithDefaults() *AutopayConfigCreateRequest`

NewAutopayConfigCreateRequestWithDefaults instantiates a new AutopayConfigCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *AutopayConfigCreateRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AutopayConfigCreateRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AutopayConfigCreateRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AutopayConfigCreateRequest) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetConfig

`func (o *AutopayConfigCreateRequest) GetConfig() AutopayConfigData`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AutopayConfigCreateRequest) GetConfigOk() (*AutopayConfigData, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AutopayConfigCreateRequest) SetConfig(v AutopayConfigData)`

SetConfig sets Config field to given value.


### GetLendingAccountId

`func (o *AutopayConfigCreateRequest) GetLendingAccountId() string`

GetLendingAccountId returns the LendingAccountId field if non-nil, zero value otherwise.

### GetLendingAccountIdOk

`func (o *AutopayConfigCreateRequest) GetLendingAccountIdOk() (*string, bool)`

GetLendingAccountIdOk returns a tuple with the LendingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLendingAccountId

`func (o *AutopayConfigCreateRequest) SetLendingAccountId(v string)`

SetLendingAccountId sets LendingAccountId field to given value.


### GetPersonId

`func (o *AutopayConfigCreateRequest) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *AutopayConfigCreateRequest) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *AutopayConfigCreateRequest) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *AutopayConfigCreateRequest) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


