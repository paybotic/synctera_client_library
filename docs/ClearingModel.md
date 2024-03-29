# ClearingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents) | 
**CardAcceptor** | Pointer to [**CardAcceptorModel**](CardAcceptorModel.md) |  | [optional] 
**ForcePost** | Pointer to **bool** |  | [optional] [default to false]
**IsRefund** | Pointer to **bool** |  | [optional] [default to false]
**Mid** | Pointer to **string** |  | [optional] 
**NetworkFees** | Pointer to [**[]NetworkFeeModel**](NetworkFeeModel.md) |  | [optional] 
**OriginalTransactionId** | **string** | When you simulated a clearing the synchronous response would have included a &#x60;transaction.token&#x60;. That value should be used here. | 

## Methods

### NewClearingModel

`func NewClearingModel(amount int32, originalTransactionId string, ) *ClearingModel`

NewClearingModel instantiates a new ClearingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearingModelWithDefaults

`func NewClearingModelWithDefaults() *ClearingModel`

NewClearingModelWithDefaults instantiates a new ClearingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ClearingModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ClearingModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ClearingModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *ClearingModel) GetCardAcceptor() CardAcceptorModel`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *ClearingModel) GetCardAcceptorOk() (*CardAcceptorModel, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *ClearingModel) SetCardAcceptor(v CardAcceptorModel)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *ClearingModel) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.

### GetForcePost

`func (o *ClearingModel) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *ClearingModel) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *ClearingModel) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.

### HasForcePost

`func (o *ClearingModel) HasForcePost() bool`

HasForcePost returns a boolean if a field has been set.

### GetIsRefund

`func (o *ClearingModel) GetIsRefund() bool`

GetIsRefund returns the IsRefund field if non-nil, zero value otherwise.

### GetIsRefundOk

`func (o *ClearingModel) GetIsRefundOk() (*bool, bool)`

GetIsRefundOk returns a tuple with the IsRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRefund

`func (o *ClearingModel) SetIsRefund(v bool)`

SetIsRefund sets IsRefund field to given value.

### HasIsRefund

`func (o *ClearingModel) HasIsRefund() bool`

HasIsRefund returns a boolean if a field has been set.

### GetMid

`func (o *ClearingModel) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *ClearingModel) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *ClearingModel) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *ClearingModel) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetNetworkFees

`func (o *ClearingModel) GetNetworkFees() []NetworkFeeModel`

GetNetworkFees returns the NetworkFees field if non-nil, zero value otherwise.

### GetNetworkFeesOk

`func (o *ClearingModel) GetNetworkFeesOk() (*[]NetworkFeeModel, bool)`

GetNetworkFeesOk returns a tuple with the NetworkFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFees

`func (o *ClearingModel) SetNetworkFees(v []NetworkFeeModel)`

SetNetworkFees sets NetworkFees field to given value.

### HasNetworkFees

`func (o *ClearingModel) HasNetworkFees() bool`

HasNetworkFees returns a boolean if a field has been set.

### GetOriginalTransactionId

`func (o *ClearingModel) GetOriginalTransactionId() string`

GetOriginalTransactionId returns the OriginalTransactionId field if non-nil, zero value otherwise.

### GetOriginalTransactionIdOk

`func (o *ClearingModel) GetOriginalTransactionIdOk() (*string, bool)`

GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTransactionId

`func (o *ClearingModel) SetOriginalTransactionId(v string)`

SetOriginalTransactionId sets OriginalTransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


