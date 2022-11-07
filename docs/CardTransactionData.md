# CardTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | Pointer to **string** | Debit Network ID | [optional] [readonly] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**CurrencyConversion** | Pointer to [**CardTransactionDataCurrencyConversion**](CardTransactionDataCurrencyConversion.md) |  | [optional] 
**Merchant** | Pointer to [**CardTransactionDataMerchant**](CardTransactionDataMerchant.md) |  | [optional] 
**Network** | Pointer to **string** | The network used for the transaction | [optional] 
**NetworkReferenceId** | Pointer to **string** | The ID provided by he network to represent this transaction | [optional] 
**Pos** | Pointer to [**CardTransactionDataPos**](CardTransactionDataPos.md) |  | [optional] 

## Methods

### NewCardTransactionData

`func NewCardTransactionData() *CardTransactionData`

NewCardTransactionData instantiates a new CardTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTransactionDataWithDefaults

`func NewCardTransactionDataWithDefaults() *CardTransactionData`

NewCardTransactionDataWithDefaults instantiates a new CardTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardId

`func (o *CardTransactionData) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *CardTransactionData) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *CardTransactionData) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *CardTransactionData) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CardTransactionData) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CardTransactionData) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CardTransactionData) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CardTransactionData) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencyConversion

`func (o *CardTransactionData) GetCurrencyConversion() CardTransactionDataCurrencyConversion`

GetCurrencyConversion returns the CurrencyConversion field if non-nil, zero value otherwise.

### GetCurrencyConversionOk

`func (o *CardTransactionData) GetCurrencyConversionOk() (*CardTransactionDataCurrencyConversion, bool)`

GetCurrencyConversionOk returns a tuple with the CurrencyConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyConversion

`func (o *CardTransactionData) SetCurrencyConversion(v CardTransactionDataCurrencyConversion)`

SetCurrencyConversion sets CurrencyConversion field to given value.

### HasCurrencyConversion

`func (o *CardTransactionData) HasCurrencyConversion() bool`

HasCurrencyConversion returns a boolean if a field has been set.

### GetMerchant

`func (o *CardTransactionData) GetMerchant() CardTransactionDataMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *CardTransactionData) GetMerchantOk() (*CardTransactionDataMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *CardTransactionData) SetMerchant(v CardTransactionDataMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *CardTransactionData) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetNetwork

`func (o *CardTransactionData) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CardTransactionData) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CardTransactionData) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CardTransactionData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkReferenceId

`func (o *CardTransactionData) GetNetworkReferenceId() string`

GetNetworkReferenceId returns the NetworkReferenceId field if non-nil, zero value otherwise.

### GetNetworkReferenceIdOk

`func (o *CardTransactionData) GetNetworkReferenceIdOk() (*string, bool)`

GetNetworkReferenceIdOk returns a tuple with the NetworkReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkReferenceId

`func (o *CardTransactionData) SetNetworkReferenceId(v string)`

SetNetworkReferenceId sets NetworkReferenceId field to given value.

### HasNetworkReferenceId

`func (o *CardTransactionData) HasNetworkReferenceId() bool`

HasNetworkReferenceId returns a boolean if a field has been set.

### GetPos

`func (o *CardTransactionData) GetPos() CardTransactionDataPos`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *CardTransactionData) GetPosOk() (*CardTransactionDataPos, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *CardTransactionData) SetPos(v CardTransactionDataPos)`

SetPos sets Pos field to given value.

### HasPos

`func (o *CardTransactionData) HasPos() bool`

HasPos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


