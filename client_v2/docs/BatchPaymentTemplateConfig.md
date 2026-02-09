# BatchPaymentTemplateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalAccountId** | **string** | The ID of the external account that will be to match the batch transfer the appropriate template.  | 
**MaxIndividualTransactionAmount** | Pointer to **int64** | The maximum amount that a single transaction can be for the batch transfer.  | [optional] 
**MaxTotalTransactionAmount** | Pointer to **int64** | The maximum amount that the total transactions can be for the batch transfer.  | [optional] 
**MaxTransactionCount** | Pointer to **int32** | The maximum number of transactions that can be in the batch transfer.  | [optional] 
**MinIndividualTransactionAmount** | Pointer to **int64** | The minimum amount that a single transaction can be for the batch transfer.  | [optional] 
**MinTotalTransactionAmount** | Pointer to **int64** | The minimum amount that the total transactions can be for the batch transfer.  | [optional] 
**MinTransactionCount** | Pointer to **int32** | The minimum number of transactions that can be in the batch transfer.  | [optional] 
**SettlementAccountId** | **string** | The ID of the settlement account that will be to match the batch transfer the appropriate template.  | 
**SettlementCustomerId** | **string** | The customer ID of the settlement account.  | 
**Subtypes** | **[]string** | The transaction subtypes that will be to match the batch transfer the appropriate template.  | 
**Type** | **string** | The transaction type that will be to match the batch transfer the appropriate template.  | 

## Methods

### NewBatchPaymentTemplateConfig

`func NewBatchPaymentTemplateConfig(externalAccountId string, settlementAccountId string, settlementCustomerId string, subtypes []string, type_ string, ) *BatchPaymentTemplateConfig`

NewBatchPaymentTemplateConfig instantiates a new BatchPaymentTemplateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchPaymentTemplateConfigWithDefaults

`func NewBatchPaymentTemplateConfigWithDefaults() *BatchPaymentTemplateConfig`

NewBatchPaymentTemplateConfigWithDefaults instantiates a new BatchPaymentTemplateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalAccountId

`func (o *BatchPaymentTemplateConfig) GetExternalAccountId() string`

GetExternalAccountId returns the ExternalAccountId field if non-nil, zero value otherwise.

### GetExternalAccountIdOk

`func (o *BatchPaymentTemplateConfig) GetExternalAccountIdOk() (*string, bool)`

GetExternalAccountIdOk returns a tuple with the ExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountId

`func (o *BatchPaymentTemplateConfig) SetExternalAccountId(v string)`

SetExternalAccountId sets ExternalAccountId field to given value.


### GetMaxIndividualTransactionAmount

`func (o *BatchPaymentTemplateConfig) GetMaxIndividualTransactionAmount() int64`

GetMaxIndividualTransactionAmount returns the MaxIndividualTransactionAmount field if non-nil, zero value otherwise.

### GetMaxIndividualTransactionAmountOk

`func (o *BatchPaymentTemplateConfig) GetMaxIndividualTransactionAmountOk() (*int64, bool)`

GetMaxIndividualTransactionAmountOk returns a tuple with the MaxIndividualTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIndividualTransactionAmount

`func (o *BatchPaymentTemplateConfig) SetMaxIndividualTransactionAmount(v int64)`

SetMaxIndividualTransactionAmount sets MaxIndividualTransactionAmount field to given value.

### HasMaxIndividualTransactionAmount

`func (o *BatchPaymentTemplateConfig) HasMaxIndividualTransactionAmount() bool`

HasMaxIndividualTransactionAmount returns a boolean if a field has been set.

### GetMaxTotalTransactionAmount

`func (o *BatchPaymentTemplateConfig) GetMaxTotalTransactionAmount() int64`

GetMaxTotalTransactionAmount returns the MaxTotalTransactionAmount field if non-nil, zero value otherwise.

### GetMaxTotalTransactionAmountOk

`func (o *BatchPaymentTemplateConfig) GetMaxTotalTransactionAmountOk() (*int64, bool)`

GetMaxTotalTransactionAmountOk returns a tuple with the MaxTotalTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalTransactionAmount

`func (o *BatchPaymentTemplateConfig) SetMaxTotalTransactionAmount(v int64)`

SetMaxTotalTransactionAmount sets MaxTotalTransactionAmount field to given value.

### HasMaxTotalTransactionAmount

`func (o *BatchPaymentTemplateConfig) HasMaxTotalTransactionAmount() bool`

HasMaxTotalTransactionAmount returns a boolean if a field has been set.

### GetMaxTransactionCount

`func (o *BatchPaymentTemplateConfig) GetMaxTransactionCount() int32`

GetMaxTransactionCount returns the MaxTransactionCount field if non-nil, zero value otherwise.

### GetMaxTransactionCountOk

`func (o *BatchPaymentTemplateConfig) GetMaxTransactionCountOk() (*int32, bool)`

GetMaxTransactionCountOk returns a tuple with the MaxTransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactionCount

`func (o *BatchPaymentTemplateConfig) SetMaxTransactionCount(v int32)`

SetMaxTransactionCount sets MaxTransactionCount field to given value.

### HasMaxTransactionCount

`func (o *BatchPaymentTemplateConfig) HasMaxTransactionCount() bool`

HasMaxTransactionCount returns a boolean if a field has been set.

### GetMinIndividualTransactionAmount

`func (o *BatchPaymentTemplateConfig) GetMinIndividualTransactionAmount() int64`

GetMinIndividualTransactionAmount returns the MinIndividualTransactionAmount field if non-nil, zero value otherwise.

### GetMinIndividualTransactionAmountOk

`func (o *BatchPaymentTemplateConfig) GetMinIndividualTransactionAmountOk() (*int64, bool)`

GetMinIndividualTransactionAmountOk returns a tuple with the MinIndividualTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIndividualTransactionAmount

`func (o *BatchPaymentTemplateConfig) SetMinIndividualTransactionAmount(v int64)`

SetMinIndividualTransactionAmount sets MinIndividualTransactionAmount field to given value.

### HasMinIndividualTransactionAmount

`func (o *BatchPaymentTemplateConfig) HasMinIndividualTransactionAmount() bool`

HasMinIndividualTransactionAmount returns a boolean if a field has been set.

### GetMinTotalTransactionAmount

`func (o *BatchPaymentTemplateConfig) GetMinTotalTransactionAmount() int64`

GetMinTotalTransactionAmount returns the MinTotalTransactionAmount field if non-nil, zero value otherwise.

### GetMinTotalTransactionAmountOk

`func (o *BatchPaymentTemplateConfig) GetMinTotalTransactionAmountOk() (*int64, bool)`

GetMinTotalTransactionAmountOk returns a tuple with the MinTotalTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTotalTransactionAmount

`func (o *BatchPaymentTemplateConfig) SetMinTotalTransactionAmount(v int64)`

SetMinTotalTransactionAmount sets MinTotalTransactionAmount field to given value.

### HasMinTotalTransactionAmount

`func (o *BatchPaymentTemplateConfig) HasMinTotalTransactionAmount() bool`

HasMinTotalTransactionAmount returns a boolean if a field has been set.

### GetMinTransactionCount

`func (o *BatchPaymentTemplateConfig) GetMinTransactionCount() int32`

GetMinTransactionCount returns the MinTransactionCount field if non-nil, zero value otherwise.

### GetMinTransactionCountOk

`func (o *BatchPaymentTemplateConfig) GetMinTransactionCountOk() (*int32, bool)`

GetMinTransactionCountOk returns a tuple with the MinTransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTransactionCount

`func (o *BatchPaymentTemplateConfig) SetMinTransactionCount(v int32)`

SetMinTransactionCount sets MinTransactionCount field to given value.

### HasMinTransactionCount

`func (o *BatchPaymentTemplateConfig) HasMinTransactionCount() bool`

HasMinTransactionCount returns a boolean if a field has been set.

### GetSettlementAccountId

`func (o *BatchPaymentTemplateConfig) GetSettlementAccountId() string`

GetSettlementAccountId returns the SettlementAccountId field if non-nil, zero value otherwise.

### GetSettlementAccountIdOk

`func (o *BatchPaymentTemplateConfig) GetSettlementAccountIdOk() (*string, bool)`

GetSettlementAccountIdOk returns a tuple with the SettlementAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAccountId

`func (o *BatchPaymentTemplateConfig) SetSettlementAccountId(v string)`

SetSettlementAccountId sets SettlementAccountId field to given value.


### GetSettlementCustomerId

`func (o *BatchPaymentTemplateConfig) GetSettlementCustomerId() string`

GetSettlementCustomerId returns the SettlementCustomerId field if non-nil, zero value otherwise.

### GetSettlementCustomerIdOk

`func (o *BatchPaymentTemplateConfig) GetSettlementCustomerIdOk() (*string, bool)`

GetSettlementCustomerIdOk returns a tuple with the SettlementCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCustomerId

`func (o *BatchPaymentTemplateConfig) SetSettlementCustomerId(v string)`

SetSettlementCustomerId sets SettlementCustomerId field to given value.


### GetSubtypes

`func (o *BatchPaymentTemplateConfig) GetSubtypes() []string`

GetSubtypes returns the Subtypes field if non-nil, zero value otherwise.

### GetSubtypesOk

`func (o *BatchPaymentTemplateConfig) GetSubtypesOk() (*[]string, bool)`

GetSubtypesOk returns a tuple with the Subtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtypes

`func (o *BatchPaymentTemplateConfig) SetSubtypes(v []string)`

SetSubtypes sets Subtypes field to given value.


### GetType

`func (o *BatchPaymentTemplateConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BatchPaymentTemplateConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BatchPaymentTemplateConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


