# ExternalCardTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalCardTransaction** | [**CardTransactionData**](CardTransactionData.md) |  | 
**Subtype** | [**ExternalCardTransactionSubtypes**](ExternalCardTransactionSubtypes.md) |  | 

## Methods

### NewExternalCardTransactionAllOf

`func NewExternalCardTransactionAllOf(externalCardTransaction CardTransactionData, subtype ExternalCardTransactionSubtypes, ) *ExternalCardTransactionAllOf`

NewExternalCardTransactionAllOf instantiates a new ExternalCardTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardTransactionAllOfWithDefaults

`func NewExternalCardTransactionAllOfWithDefaults() *ExternalCardTransactionAllOf`

NewExternalCardTransactionAllOfWithDefaults instantiates a new ExternalCardTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalCardTransaction

`func (o *ExternalCardTransactionAllOf) GetExternalCardTransaction() CardTransactionData`

GetExternalCardTransaction returns the ExternalCardTransaction field if non-nil, zero value otherwise.

### GetExternalCardTransactionOk

`func (o *ExternalCardTransactionAllOf) GetExternalCardTransactionOk() (*CardTransactionData, bool)`

GetExternalCardTransactionOk returns a tuple with the ExternalCardTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCardTransaction

`func (o *ExternalCardTransactionAllOf) SetExternalCardTransaction(v CardTransactionData)`

SetExternalCardTransaction sets ExternalCardTransaction field to given value.


### GetSubtype

`func (o *ExternalCardTransactionAllOf) GetSubtype() ExternalCardTransactionSubtypes`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ExternalCardTransactionAllOf) GetSubtypeOk() (*ExternalCardTransactionSubtypes, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ExternalCardTransactionAllOf) SetSubtype(v ExternalCardTransactionSubtypes)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


