# CheckTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckTransaction** | [**CheckTransactionData**](CheckTransactionData.md) |  | 
**Subtype** | [**CheckTransactionSubtypes**](CheckTransactionSubtypes.md) |  | 

## Methods

### NewCheckTransactionAllOf

`func NewCheckTransactionAllOf(checkTransaction CheckTransactionData, subtype CheckTransactionSubtypes, ) *CheckTransactionAllOf`

NewCheckTransactionAllOf instantiates a new CheckTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTransactionAllOfWithDefaults

`func NewCheckTransactionAllOfWithDefaults() *CheckTransactionAllOf`

NewCheckTransactionAllOfWithDefaults instantiates a new CheckTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckTransaction

`func (o *CheckTransactionAllOf) GetCheckTransaction() CheckTransactionData`

GetCheckTransaction returns the CheckTransaction field if non-nil, zero value otherwise.

### GetCheckTransactionOk

`func (o *CheckTransactionAllOf) GetCheckTransactionOk() (*CheckTransactionData, bool)`

GetCheckTransactionOk returns a tuple with the CheckTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTransaction

`func (o *CheckTransactionAllOf) SetCheckTransaction(v CheckTransactionData)`

SetCheckTransaction sets CheckTransaction field to given value.


### GetSubtype

`func (o *CheckTransactionAllOf) GetSubtype() CheckTransactionSubtypes`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CheckTransactionAllOf) GetSubtypeOk() (*CheckTransactionSubtypes, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CheckTransactionAllOf) SetSubtype(v CheckTransactionSubtypes)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


