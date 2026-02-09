# CashBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transfer amount in cents | 
**DcSign** | **string** | Debit or credit sign | 
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 

## Methods

### NewCashBase

`func NewCashBase(amount int64, dcSign string, ) *CashBase`

NewCashBase instantiates a new CashBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashBaseWithDefaults

`func NewCashBaseWithDefaults() *CashBase`

NewCashBaseWithDefaults instantiates a new CashBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CashBase) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashBase) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashBase) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetDcSign

`func (o *CashBase) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *CashBase) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *CashBase) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetSourceData

`func (o *CashBase) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *CashBase) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *CashBase) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *CashBase) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


