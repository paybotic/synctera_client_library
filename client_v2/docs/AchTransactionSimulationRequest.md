# AchTransactionSimulationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Number of the receiving account | 
**Amount** | **int32** | Amount to transfer in cents (e.g. 100 &#x3D; $1). Generates a prenote if set to 0. | 
**DcSign** | **string** | The type of transaction (debit or credit) in relation to the receiving account. A credit is a transfer in and a debit is a transfer pulling money out of the receiving account. | 
**EffectiveDate** | **string** | Effective date of the transaction. Transactions with the current date or date in the past are posted immediately. Future-dated transactions are scheduled to be posted on the chosen date. | 

## Methods

### NewAchTransactionSimulationRequest

`func NewAchTransactionSimulationRequest(accountNumber string, amount int32, dcSign string, effectiveDate string, ) *AchTransactionSimulationRequest`

NewAchTransactionSimulationRequest instantiates a new AchTransactionSimulationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchTransactionSimulationRequestWithDefaults

`func NewAchTransactionSimulationRequestWithDefaults() *AchTransactionSimulationRequest`

NewAchTransactionSimulationRequestWithDefaults instantiates a new AchTransactionSimulationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *AchTransactionSimulationRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AchTransactionSimulationRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AchTransactionSimulationRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAmount

`func (o *AchTransactionSimulationRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AchTransactionSimulationRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AchTransactionSimulationRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetDcSign

`func (o *AchTransactionSimulationRequest) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *AchTransactionSimulationRequest) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *AchTransactionSimulationRequest) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetEffectiveDate

`func (o *AchTransactionSimulationRequest) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AchTransactionSimulationRequest) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AchTransactionSimulationRequest) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


