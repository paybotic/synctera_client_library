# TransactionVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The amount of the transactions in ISO 4217 minor currency units. For example, a transaction of $100 USD will be displayed as 10000. | [optional] 
**Channel** | Pointer to **string** | The channel of the transaction volume. | [optional] 
**ChannelCoverage** | Pointer to **string** | The channel coverage of the transaction volume. | [optional] 
**Currency** | Pointer to **string** | The currency in ISO 4217 format. | [optional] 
**Frequency** | Pointer to [**NullableFrequency**](Frequency.md) |  | [optional] 
**OnSynctera** | Pointer to **bool** | Whether the transaction volume is on Synctera. | [optional] 
**TransactionCount** | Pointer to **int32** | The number of transactions. | [optional] 

## Methods

### NewTransactionVolume

`func NewTransactionVolume() *TransactionVolume`

NewTransactionVolume instantiates a new TransactionVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionVolumeWithDefaults

`func NewTransactionVolumeWithDefaults() *TransactionVolume`

NewTransactionVolumeWithDefaults instantiates a new TransactionVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionVolume) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionVolume) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionVolume) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionVolume) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChannel

`func (o *TransactionVolume) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TransactionVolume) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TransactionVolume) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TransactionVolume) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelCoverage

`func (o *TransactionVolume) GetChannelCoverage() string`

GetChannelCoverage returns the ChannelCoverage field if non-nil, zero value otherwise.

### GetChannelCoverageOk

`func (o *TransactionVolume) GetChannelCoverageOk() (*string, bool)`

GetChannelCoverageOk returns a tuple with the ChannelCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCoverage

`func (o *TransactionVolume) SetChannelCoverage(v string)`

SetChannelCoverage sets ChannelCoverage field to given value.

### HasChannelCoverage

`func (o *TransactionVolume) HasChannelCoverage() bool`

HasChannelCoverage returns a boolean if a field has been set.

### GetCurrency

`func (o *TransactionVolume) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionVolume) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionVolume) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionVolume) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFrequency

`func (o *TransactionVolume) GetFrequency() Frequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *TransactionVolume) GetFrequencyOk() (*Frequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *TransactionVolume) SetFrequency(v Frequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *TransactionVolume) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *TransactionVolume) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *TransactionVolume) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetOnSynctera

`func (o *TransactionVolume) GetOnSynctera() bool`

GetOnSynctera returns the OnSynctera field if non-nil, zero value otherwise.

### GetOnSyncteraOk

`func (o *TransactionVolume) GetOnSyncteraOk() (*bool, bool)`

GetOnSyncteraOk returns a tuple with the OnSynctera field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSynctera

`func (o *TransactionVolume) SetOnSynctera(v bool)`

SetOnSynctera sets OnSynctera field to given value.

### HasOnSynctera

`func (o *TransactionVolume) HasOnSynctera() bool`

HasOnSynctera returns a boolean if a field has been set.

### GetTransactionCount

`func (o *TransactionVolume) GetTransactionCount() int32`

GetTransactionCount returns the TransactionCount field if non-nil, zero value otherwise.

### GetTransactionCountOk

`func (o *TransactionVolume) GetTransactionCountOk() (*int32, bool)`

GetTransactionCountOk returns a tuple with the TransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCount

`func (o *TransactionVolume) SetTransactionCount(v int32)`

SetTransactionCount sets TransactionCount field to given value.

### HasTransactionCount

`func (o *TransactionVolume) HasTransactionCount() bool`

HasTransactionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


