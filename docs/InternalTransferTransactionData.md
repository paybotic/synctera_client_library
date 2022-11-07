# InternalTransferTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginatingAccountId** | Pointer to **string** | The uuid of the account being debited | [optional] 
**ReceivingAccountId** | Pointer to **string** | The uuid of the account being credited | [optional] 

## Methods

### NewInternalTransferTransactionData

`func NewInternalTransferTransactionData() *InternalTransferTransactionData`

NewInternalTransferTransactionData instantiates a new InternalTransferTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalTransferTransactionDataWithDefaults

`func NewInternalTransferTransactionDataWithDefaults() *InternalTransferTransactionData`

NewInternalTransferTransactionDataWithDefaults instantiates a new InternalTransferTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginatingAccountId

`func (o *InternalTransferTransactionData) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *InternalTransferTransactionData) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *InternalTransferTransactionData) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.

### HasOriginatingAccountId

`func (o *InternalTransferTransactionData) HasOriginatingAccountId() bool`

HasOriginatingAccountId returns a boolean if a field has been set.

### GetReceivingAccountId

`func (o *InternalTransferTransactionData) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *InternalTransferTransactionData) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *InternalTransferTransactionData) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.

### HasReceivingAccountId

`func (o *InternalTransferTransactionData) HasReceivingAccountId() bool`

HasReceivingAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


