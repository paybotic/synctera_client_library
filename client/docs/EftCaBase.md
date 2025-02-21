# EftCaBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transfer amount in cents | 
**CustomerId** | **string** | The UUID of the Synctera customer resource that is the originator of the transfer.  | 
**DcSign** | **string** | Debit or credit sign | 
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 
**TransactionCode** | **string** | The three digit transaction code that identifies the type of transaction. More information can be found here: https://www.payments.ca/sites/default/files/standard007eng.pdf.  | 

## Methods

### NewEftCaBase

`func NewEftCaBase(amount int64, customerId string, dcSign string, transactionCode string, ) *EftCaBase`

NewEftCaBase instantiates a new EftCaBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEftCaBaseWithDefaults

`func NewEftCaBaseWithDefaults() *EftCaBase`

NewEftCaBaseWithDefaults instantiates a new EftCaBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EftCaBase) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EftCaBase) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EftCaBase) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCustomerId

`func (o *EftCaBase) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *EftCaBase) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *EftCaBase) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDcSign

`func (o *EftCaBase) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *EftCaBase) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *EftCaBase) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetSourceData

`func (o *EftCaBase) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *EftCaBase) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *EftCaBase) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *EftCaBase) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetTransactionCode

`func (o *EftCaBase) GetTransactionCode() string`

GetTransactionCode returns the TransactionCode field if non-nil, zero value otherwise.

### GetTransactionCodeOk

`func (o *EftCaBase) GetTransactionCodeOk() (*string, bool)`

GetTransactionCodeOk returns a tuple with the TransactionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCode

`func (o *EftCaBase) SetTransactionCode(v string)`

SetTransactionCode sets TransactionCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


