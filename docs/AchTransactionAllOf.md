# AchTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchTransaction** | [**AchTransactionData**](AchTransactionData.md) |  | 
**Subtype** | [**AchTransactionSubtypes**](AchTransactionSubtypes.md) |  | 

## Methods

### NewAchTransactionAllOf

`func NewAchTransactionAllOf(achTransaction AchTransactionData, subtype AchTransactionSubtypes, ) *AchTransactionAllOf`

NewAchTransactionAllOf instantiates a new AchTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchTransactionAllOfWithDefaults

`func NewAchTransactionAllOfWithDefaults() *AchTransactionAllOf`

NewAchTransactionAllOfWithDefaults instantiates a new AchTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchTransaction

`func (o *AchTransactionAllOf) GetAchTransaction() AchTransactionData`

GetAchTransaction returns the AchTransaction field if non-nil, zero value otherwise.

### GetAchTransactionOk

`func (o *AchTransactionAllOf) GetAchTransactionOk() (*AchTransactionData, bool)`

GetAchTransactionOk returns a tuple with the AchTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchTransaction

`func (o *AchTransactionAllOf) SetAchTransaction(v AchTransactionData)`

SetAchTransaction sets AchTransaction field to given value.


### GetSubtype

`func (o *AchTransactionAllOf) GetSubtype() AchTransactionSubtypes`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AchTransactionAllOf) GetSubtypeOk() (*AchTransactionSubtypes, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AchTransactionAllOf) SetSubtype(v AchTransactionSubtypes)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


