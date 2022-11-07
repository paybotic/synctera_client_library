# InternalTransferTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalTransferTransaction** | Pointer to [**InternalTransferTransactionData**](InternalTransferTransactionData.md) |  | [optional] 
**Subtype** | [**InternalTransferTransactionSubtypes**](InternalTransferTransactionSubtypes.md) |  | 

## Methods

### NewInternalTransferTransactionAllOf

`func NewInternalTransferTransactionAllOf(subtype InternalTransferTransactionSubtypes, ) *InternalTransferTransactionAllOf`

NewInternalTransferTransactionAllOf instantiates a new InternalTransferTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalTransferTransactionAllOfWithDefaults

`func NewInternalTransferTransactionAllOfWithDefaults() *InternalTransferTransactionAllOf`

NewInternalTransferTransactionAllOfWithDefaults instantiates a new InternalTransferTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalTransferTransaction

`func (o *InternalTransferTransactionAllOf) GetInternalTransferTransaction() InternalTransferTransactionData`

GetInternalTransferTransaction returns the InternalTransferTransaction field if non-nil, zero value otherwise.

### GetInternalTransferTransactionOk

`func (o *InternalTransferTransactionAllOf) GetInternalTransferTransactionOk() (*InternalTransferTransactionData, bool)`

GetInternalTransferTransactionOk returns a tuple with the InternalTransferTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTransferTransaction

`func (o *InternalTransferTransactionAllOf) SetInternalTransferTransaction(v InternalTransferTransactionData)`

SetInternalTransferTransaction sets InternalTransferTransaction field to given value.

### HasInternalTransferTransaction

`func (o *InternalTransferTransactionAllOf) HasInternalTransferTransaction() bool`

HasInternalTransferTransaction returns a boolean if a field has been set.

### GetSubtype

`func (o *InternalTransferTransactionAllOf) GetSubtype() InternalTransferTransactionSubtypes`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InternalTransferTransactionAllOf) GetSubtypeOk() (*InternalTransferTransactionSubtypes, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InternalTransferTransactionAllOf) SetSubtype(v InternalTransferTransactionSubtypes)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


