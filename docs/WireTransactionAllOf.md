# WireTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subtype** | [**WireTransactionSubtypes**](WireTransactionSubtypes.md) |  | 
**WireTransaction** | [**WireTransactionData**](WireTransactionData.md) |  | 

## Methods

### NewWireTransactionAllOf

`func NewWireTransactionAllOf(subtype WireTransactionSubtypes, wireTransaction WireTransactionData, ) *WireTransactionAllOf`

NewWireTransactionAllOf instantiates a new WireTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireTransactionAllOfWithDefaults

`func NewWireTransactionAllOfWithDefaults() *WireTransactionAllOf`

NewWireTransactionAllOfWithDefaults instantiates a new WireTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubtype

`func (o *WireTransactionAllOf) GetSubtype() WireTransactionSubtypes`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *WireTransactionAllOf) GetSubtypeOk() (*WireTransactionSubtypes, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *WireTransactionAllOf) SetSubtype(v WireTransactionSubtypes)`

SetSubtype sets Subtype field to given value.


### GetWireTransaction

`func (o *WireTransactionAllOf) GetWireTransaction() WireTransactionData`

GetWireTransaction returns the WireTransaction field if non-nil, zero value otherwise.

### GetWireTransactionOk

`func (o *WireTransactionAllOf) GetWireTransactionOk() (*WireTransactionData, bool)`

GetWireTransactionOk returns a tuple with the WireTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireTransaction

`func (o *WireTransactionAllOf) SetWireTransaction(v WireTransactionData)`

SetWireTransaction sets WireTransaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


