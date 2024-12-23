# TransactionLine1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account uuid associated with this transaction line | 
**AccountNo** | **string** | The account number associated with this transaction line | 
**Amount** | **int64** | The amount (in cents) of the transaction | 
**AvailBalance** | **int64** | The account \&quot;available balance\&quot; at the point in time this (to be deprecated) transaction was posted | 
**AvailableBalance** | **int64** | The account \&quot;available balance\&quot; at the point in time this transaction was posted | 
**Balance** | **int64** | The account balance at the point in time this transaction was posted | 
**Created** | **time.Time** | The creation date of the transaction | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**DcSign** | [**DcSign**](DcSign.md) |  | 
**IsFee** | **bool** | Whether or not this line is considered a fee | 
**IsGlAcc** | **bool** | Whether or not this line represents a GL account | 
**IsOffset** | **bool** | Whether or not this line is considered the \&quot;offset\&quot; line | 
**IsPrimary** | **bool** | Whether or not this line is considered the \&quot;primary\&quot; line | 
**Meta** | **map[string]interface{}** |  | 
**Network** | **string** | The network this transaction is associated with | 
**RelatedLine** | **int32** |  | 
**Seq** | **int32** |  | 
**Tenant** | **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | 
**Updated** | **time.Time** | The date the transaction was last updated | 
**Uuid** | **string** |  | 

## Methods

### NewTransactionLine1

`func NewTransactionLine1(accountId string, accountNo string, amount int64, availBalance int64, availableBalance int64, balance int64, created time.Time, currency string, dcSign DcSign, isFee bool, isGlAcc bool, isOffset bool, isPrimary bool, meta map[string]interface{}, network string, relatedLine int32, seq int32, tenant string, updated time.Time, uuid string, ) *TransactionLine1`

NewTransactionLine1 instantiates a new TransactionLine1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionLine1WithDefaults

`func NewTransactionLine1WithDefaults() *TransactionLine1`

NewTransactionLine1WithDefaults instantiates a new TransactionLine1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransactionLine1) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionLine1) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionLine1) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountNo

`func (o *TransactionLine1) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *TransactionLine1) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *TransactionLine1) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.


### GetAmount

`func (o *TransactionLine1) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionLine1) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionLine1) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetAvailBalance

`func (o *TransactionLine1) GetAvailBalance() int64`

GetAvailBalance returns the AvailBalance field if non-nil, zero value otherwise.

### GetAvailBalanceOk

`func (o *TransactionLine1) GetAvailBalanceOk() (*int64, bool)`

GetAvailBalanceOk returns a tuple with the AvailBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBalance

`func (o *TransactionLine1) SetAvailBalance(v int64)`

SetAvailBalance sets AvailBalance field to given value.


### GetAvailableBalance

`func (o *TransactionLine1) GetAvailableBalance() int64`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *TransactionLine1) GetAvailableBalanceOk() (*int64, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *TransactionLine1) SetAvailableBalance(v int64)`

SetAvailableBalance sets AvailableBalance field to given value.


### GetBalance

`func (o *TransactionLine1) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TransactionLine1) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TransactionLine1) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetCreated

`func (o *TransactionLine1) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TransactionLine1) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TransactionLine1) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCurrency

`func (o *TransactionLine1) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionLine1) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionLine1) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDcSign

`func (o *TransactionLine1) GetDcSign() DcSign`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *TransactionLine1) GetDcSignOk() (*DcSign, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *TransactionLine1) SetDcSign(v DcSign)`

SetDcSign sets DcSign field to given value.


### GetIsFee

`func (o *TransactionLine1) GetIsFee() bool`

GetIsFee returns the IsFee field if non-nil, zero value otherwise.

### GetIsFeeOk

`func (o *TransactionLine1) GetIsFeeOk() (*bool, bool)`

GetIsFeeOk returns a tuple with the IsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFee

`func (o *TransactionLine1) SetIsFee(v bool)`

SetIsFee sets IsFee field to given value.


### GetIsGlAcc

`func (o *TransactionLine1) GetIsGlAcc() bool`

GetIsGlAcc returns the IsGlAcc field if non-nil, zero value otherwise.

### GetIsGlAccOk

`func (o *TransactionLine1) GetIsGlAccOk() (*bool, bool)`

GetIsGlAccOk returns a tuple with the IsGlAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlAcc

`func (o *TransactionLine1) SetIsGlAcc(v bool)`

SetIsGlAcc sets IsGlAcc field to given value.


### GetIsOffset

`func (o *TransactionLine1) GetIsOffset() bool`

GetIsOffset returns the IsOffset field if non-nil, zero value otherwise.

### GetIsOffsetOk

`func (o *TransactionLine1) GetIsOffsetOk() (*bool, bool)`

GetIsOffsetOk returns a tuple with the IsOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOffset

`func (o *TransactionLine1) SetIsOffset(v bool)`

SetIsOffset sets IsOffset field to given value.


### GetIsPrimary

`func (o *TransactionLine1) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *TransactionLine1) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *TransactionLine1) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### GetMeta

`func (o *TransactionLine1) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TransactionLine1) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TransactionLine1) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### SetMetaNil

`func (o *TransactionLine1) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *TransactionLine1) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetNetwork

`func (o *TransactionLine1) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionLine1) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionLine1) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRelatedLine

`func (o *TransactionLine1) GetRelatedLine() int32`

GetRelatedLine returns the RelatedLine field if non-nil, zero value otherwise.

### GetRelatedLineOk

`func (o *TransactionLine1) GetRelatedLineOk() (*int32, bool)`

GetRelatedLineOk returns a tuple with the RelatedLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedLine

`func (o *TransactionLine1) SetRelatedLine(v int32)`

SetRelatedLine sets RelatedLine field to given value.


### GetSeq

`func (o *TransactionLine1) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *TransactionLine1) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *TransactionLine1) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetTenant

`func (o *TransactionLine1) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransactionLine1) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransactionLine1) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetUpdated

`func (o *TransactionLine1) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TransactionLine1) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TransactionLine1) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetUuid

`func (o *TransactionLine1) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TransactionLine1) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TransactionLine1) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


