# InternationalWireResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountInBeneficiaryCurrency** | **int64** | Transfer amount in cents ($100 would be 10000) in the currency of the beneficiary. | 
**BeneficiaryCurrency** | **string** | The currency of the beneficiary&#39;s account in ISO4217 format. | 
**CaseId** | Pointer to **int32** | The ID of the associated case. | [optional] 
**CreationTime** | **time.Time** |  | [readonly] 
**CustomerId** | **string** | The customer UUID representing the person initiating the Wire transfer. | 
**DeclineReason** | Pointer to **string** | If the wire is declined after being successfully submitted, this field will contain a message explaining the reason for the decline. | [optional] 
**EffectiveDate** | Pointer to **string** | The effective date of the transaction once it gets posted. | [optional] 
**FxExchangeRate** | Pointer to **float32** | The exchange rate used for the wire transfer. | [optional] 
**Id** | **string** | The ID of the international wire transfer. | 
**LastUpdatedTime** | **time.Time** |  | [readonly] 
**MaxAcceptedSenderCost** | Pointer to **int32** | The maximum amount the sender is willing to pay for the wire transfer in case FX fluctuates, required if the quote type is &#39;estimate&#39;. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Additional transfer metadata structured as key-value pairs. | [optional] 
**NetworkReferenceId** | Pointer to **string** | The network reference ID of the wire transfer. | [optional] 
**OriginatingAccountId** | **string** | The account uuid representing the sender of the wire. | 
**QuoteId** | Pointer to **string** | The quote id returned from the /international_wire/quote endpoint, contains FX rate information. Required if the international wire is cross currency. | [optional] 
**ReceivingAccountId** | **string** | The external account uuid representing the recipient of the wire. Needs to be configured with international bank routing details (e.g. SWIFT code, refer to external accounts documentation). | 
**SenderAmount** | **int64** | Debit amount to the wire sender, in cents ($100 would be 10000). | 
**SenderCurrency** | **string** | The currency of the sender&#39;s account in ISO4217 format. | 
**Status** | [**InternationalWireStatus**](InternationalWireStatus.md) |  | 
**TransactionId** | Pointer to **string** | ID of the resulting transaction resource. | [optional] 
**WireInstruction** | Pointer to **string** | Optional transaction instructions for the recipient bank if required by the recipient bank. | [optional] 
**WirePurpose** | **string** | Reason for the wire transfer. | 

## Methods

### NewInternationalWireResponse

`func NewInternationalWireResponse(amountInBeneficiaryCurrency int64, beneficiaryCurrency string, creationTime time.Time, customerId string, id string, lastUpdatedTime time.Time, originatingAccountId string, receivingAccountId string, senderAmount int64, senderCurrency string, status InternationalWireStatus, wirePurpose string, ) *InternationalWireResponse`

NewInternationalWireResponse instantiates a new InternationalWireResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalWireResponseWithDefaults

`func NewInternationalWireResponseWithDefaults() *InternationalWireResponse`

NewInternationalWireResponseWithDefaults instantiates a new InternationalWireResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountInBeneficiaryCurrency

`func (o *InternationalWireResponse) GetAmountInBeneficiaryCurrency() int64`

GetAmountInBeneficiaryCurrency returns the AmountInBeneficiaryCurrency field if non-nil, zero value otherwise.

### GetAmountInBeneficiaryCurrencyOk

`func (o *InternationalWireResponse) GetAmountInBeneficiaryCurrencyOk() (*int64, bool)`

GetAmountInBeneficiaryCurrencyOk returns a tuple with the AmountInBeneficiaryCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInBeneficiaryCurrency

`func (o *InternationalWireResponse) SetAmountInBeneficiaryCurrency(v int64)`

SetAmountInBeneficiaryCurrency sets AmountInBeneficiaryCurrency field to given value.


### GetBeneficiaryCurrency

`func (o *InternationalWireResponse) GetBeneficiaryCurrency() string`

GetBeneficiaryCurrency returns the BeneficiaryCurrency field if non-nil, zero value otherwise.

### GetBeneficiaryCurrencyOk

`func (o *InternationalWireResponse) GetBeneficiaryCurrencyOk() (*string, bool)`

GetBeneficiaryCurrencyOk returns a tuple with the BeneficiaryCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryCurrency

`func (o *InternationalWireResponse) SetBeneficiaryCurrency(v string)`

SetBeneficiaryCurrency sets BeneficiaryCurrency field to given value.


### GetCaseId

`func (o *InternationalWireResponse) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *InternationalWireResponse) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *InternationalWireResponse) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *InternationalWireResponse) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetCreationTime

`func (o *InternationalWireResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *InternationalWireResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *InternationalWireResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomerId

`func (o *InternationalWireResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InternationalWireResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InternationalWireResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDeclineReason

`func (o *InternationalWireResponse) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *InternationalWireResponse) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *InternationalWireResponse) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *InternationalWireResponse) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *InternationalWireResponse) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *InternationalWireResponse) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *InternationalWireResponse) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *InternationalWireResponse) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetFxExchangeRate

`func (o *InternationalWireResponse) GetFxExchangeRate() float32`

GetFxExchangeRate returns the FxExchangeRate field if non-nil, zero value otherwise.

### GetFxExchangeRateOk

`func (o *InternationalWireResponse) GetFxExchangeRateOk() (*float32, bool)`

GetFxExchangeRateOk returns a tuple with the FxExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxExchangeRate

`func (o *InternationalWireResponse) SetFxExchangeRate(v float32)`

SetFxExchangeRate sets FxExchangeRate field to given value.

### HasFxExchangeRate

`func (o *InternationalWireResponse) HasFxExchangeRate() bool`

HasFxExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *InternationalWireResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternationalWireResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternationalWireResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *InternationalWireResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *InternationalWireResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *InternationalWireResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMaxAcceptedSenderCost

`func (o *InternationalWireResponse) GetMaxAcceptedSenderCost() int32`

GetMaxAcceptedSenderCost returns the MaxAcceptedSenderCost field if non-nil, zero value otherwise.

### GetMaxAcceptedSenderCostOk

`func (o *InternationalWireResponse) GetMaxAcceptedSenderCostOk() (*int32, bool)`

GetMaxAcceptedSenderCostOk returns a tuple with the MaxAcceptedSenderCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAcceptedSenderCost

`func (o *InternationalWireResponse) SetMaxAcceptedSenderCost(v int32)`

SetMaxAcceptedSenderCost sets MaxAcceptedSenderCost field to given value.

### HasMaxAcceptedSenderCost

`func (o *InternationalWireResponse) HasMaxAcceptedSenderCost() bool`

HasMaxAcceptedSenderCost returns a boolean if a field has been set.

### GetMetadata

`func (o *InternationalWireResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternationalWireResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternationalWireResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternationalWireResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNetworkReferenceId

`func (o *InternationalWireResponse) GetNetworkReferenceId() string`

GetNetworkReferenceId returns the NetworkReferenceId field if non-nil, zero value otherwise.

### GetNetworkReferenceIdOk

`func (o *InternationalWireResponse) GetNetworkReferenceIdOk() (*string, bool)`

GetNetworkReferenceIdOk returns a tuple with the NetworkReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkReferenceId

`func (o *InternationalWireResponse) SetNetworkReferenceId(v string)`

SetNetworkReferenceId sets NetworkReferenceId field to given value.

### HasNetworkReferenceId

`func (o *InternationalWireResponse) HasNetworkReferenceId() bool`

HasNetworkReferenceId returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *InternationalWireResponse) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *InternationalWireResponse) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *InternationalWireResponse) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetQuoteId

`func (o *InternationalWireResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *InternationalWireResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *InternationalWireResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *InternationalWireResponse) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetReceivingAccountId

`func (o *InternationalWireResponse) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *InternationalWireResponse) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *InternationalWireResponse) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.


### GetSenderAmount

`func (o *InternationalWireResponse) GetSenderAmount() int64`

GetSenderAmount returns the SenderAmount field if non-nil, zero value otherwise.

### GetSenderAmountOk

`func (o *InternationalWireResponse) GetSenderAmountOk() (*int64, bool)`

GetSenderAmountOk returns a tuple with the SenderAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAmount

`func (o *InternationalWireResponse) SetSenderAmount(v int64)`

SetSenderAmount sets SenderAmount field to given value.


### GetSenderCurrency

`func (o *InternationalWireResponse) GetSenderCurrency() string`

GetSenderCurrency returns the SenderCurrency field if non-nil, zero value otherwise.

### GetSenderCurrencyOk

`func (o *InternationalWireResponse) GetSenderCurrencyOk() (*string, bool)`

GetSenderCurrencyOk returns a tuple with the SenderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderCurrency

`func (o *InternationalWireResponse) SetSenderCurrency(v string)`

SetSenderCurrency sets SenderCurrency field to given value.


### GetStatus

`func (o *InternationalWireResponse) GetStatus() InternationalWireStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternationalWireResponse) GetStatusOk() (*InternationalWireStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternationalWireResponse) SetStatus(v InternationalWireStatus)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *InternationalWireResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *InternationalWireResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *InternationalWireResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *InternationalWireResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetWireInstruction

`func (o *InternationalWireResponse) GetWireInstruction() string`

GetWireInstruction returns the WireInstruction field if non-nil, zero value otherwise.

### GetWireInstructionOk

`func (o *InternationalWireResponse) GetWireInstructionOk() (*string, bool)`

GetWireInstructionOk returns a tuple with the WireInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireInstruction

`func (o *InternationalWireResponse) SetWireInstruction(v string)`

SetWireInstruction sets WireInstruction field to given value.

### HasWireInstruction

`func (o *InternationalWireResponse) HasWireInstruction() bool`

HasWireInstruction returns a boolean if a field has been set.

### GetWirePurpose

`func (o *InternationalWireResponse) GetWirePurpose() string`

GetWirePurpose returns the WirePurpose field if non-nil, zero value otherwise.

### GetWirePurposeOk

`func (o *InternationalWireResponse) GetWirePurposeOk() (*string, bool)`

GetWirePurposeOk returns a tuple with the WirePurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirePurpose

`func (o *InternationalWireResponse) SetWirePurpose(v string)`

SetWirePurpose sets WirePurpose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


