# CardTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardTransaction** | [**CardTransactionData**](CardTransactionData.md) |  | 
**Subtype** | [**CardTransactionSubtypes**](CardTransactionSubtypes.md) |  | 

## Methods

### NewCardTransactionAllOf

`func NewCardTransactionAllOf(cardTransaction CardTransactionData, subtype CardTransactionSubtypes, ) *CardTransactionAllOf`

NewCardTransactionAllOf instantiates a new CardTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTransactionAllOfWithDefaults

`func NewCardTransactionAllOfWithDefaults() *CardTransactionAllOf`

NewCardTransactionAllOfWithDefaults instantiates a new CardTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardTransaction

`func (o *CardTransactionAllOf) GetCardTransaction() CardTransactionData`

GetCardTransaction returns the CardTransaction field if non-nil, zero value otherwise.

### GetCardTransactionOk

`func (o *CardTransactionAllOf) GetCardTransactionOk() (*CardTransactionData, bool)`

GetCardTransactionOk returns a tuple with the CardTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardTransaction

`func (o *CardTransactionAllOf) SetCardTransaction(v CardTransactionData)`

SetCardTransaction sets CardTransaction field to given value.


### GetSubtype

`func (o *CardTransactionAllOf) GetSubtype() CardTransactionSubtypes`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CardTransactionAllOf) GetSubtypeOk() (*CardTransactionSubtypes, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CardTransactionAllOf) SetSubtype(v CardTransactionSubtypes)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


