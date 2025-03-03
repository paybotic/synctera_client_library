# Wire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Transfer amount in cents ($100 would be 10000) | 
**BankMessage** | Pointer to **string** | Instructions intended for the financial institutions that are processing the wire. | [optional] 
**BatchId** | Pointer to **string** | The batch ID associated with the wire if it was created via the batch payment API. | [optional] 
**CaseId** | Pointer to **int32** | The case id associated with the wire. | [optional] 
**CreationTime** | **time.Time** |  | [readonly] 
**Currency** | **string** | 3-character currency code | 
**CustomerId** | Pointer to **string** | The customer UUID representing the person initiating the Wire transfer | [optional] 
**EffectiveDate** | **string** | The effective date of the transaction once it gets posted | 
**Id** | **string** | wire ID | [readonly] 
**InputMessageAccountabilityData** | Pointer to **string** | The input message accountability data consists of a 8 character cycle date (CCYYMMDD) an 8 character source and a 6 character sequence number. | [optional] [readonly] 
**IsBulk** | **bool** | Whether or not the wire is a \&quot;bulk\&quot; wire created via the batch payment API. | 
**LastUpdatedTime** | **time.Time** |  | [readonly] 
**Network** | Pointer to **string** | The network used to process the wire | [optional] 
**OriginatingAccountId** | Pointer to **string** | Sender account ID | [optional] 
**OriginatingAccountNumber** | **string** | The account number representing the sender account. If the outgoing wire is a return, it refers to the sender of the initial wire not the sender of the return. | 
**ReceivingAccountId** | Pointer to **string** | The external account uuid representing the recipient of the wire. | [optional] 
**ReceivingAccountNumber** | **string** | The account number representing the recipient account. If the outgoing wire is a return, it refers to the recipient of the initial wire not the destination of the return. | 
**RecipientMessage** | Pointer to **string** | Information from the originator to the beneficiary (recipient). | [optional] 
**ReturnData** | Pointer to [**ReturnData1**](ReturnData1.md) |  | [optional] 
**SenderReferenceId** | **string** | Sender&#39;s id associated with fedwire transfer | [readonly] 
**SettlementDate** | Pointer to **string** | The settlement date of the transaction once it gets posted | [optional] 
**Status** | **string** | The current status of the transfer | [readonly] 
**StatusDetails** | Pointer to **string** | Additional details about the status of the transfer | [optional] 
**TransactionId** | **string** | ID of the resulting transaction resource | [readonly] 
**TransactionInId** | Pointer to **string** | The transaction uuid of the incoming wire that triggered an outgoing return. This is only used if the outgoing wire is a return. | [optional] 

## Methods

### NewWire

`func NewWire(amount int32, creationTime time.Time, currency string, effectiveDate string, id string, isBulk bool, lastUpdatedTime time.Time, originatingAccountNumber string, receivingAccountNumber string, senderReferenceId string, status string, transactionId string, ) *Wire`

NewWire instantiates a new Wire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireWithDefaults

`func NewWireWithDefaults() *Wire`

NewWireWithDefaults instantiates a new Wire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Wire) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Wire) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Wire) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetBankMessage

`func (o *Wire) GetBankMessage() string`

GetBankMessage returns the BankMessage field if non-nil, zero value otherwise.

### GetBankMessageOk

`func (o *Wire) GetBankMessageOk() (*string, bool)`

GetBankMessageOk returns a tuple with the BankMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankMessage

`func (o *Wire) SetBankMessage(v string)`

SetBankMessage sets BankMessage field to given value.

### HasBankMessage

`func (o *Wire) HasBankMessage() bool`

HasBankMessage returns a boolean if a field has been set.

### GetBatchId

`func (o *Wire) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *Wire) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *Wire) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *Wire) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetCaseId

`func (o *Wire) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *Wire) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *Wire) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *Wire) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### GetCreationTime

`func (o *Wire) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Wire) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Wire) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *Wire) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Wire) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Wire) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *Wire) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Wire) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Wire) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Wire) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *Wire) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *Wire) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *Wire) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetId

`func (o *Wire) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Wire) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Wire) SetId(v string)`

SetId sets Id field to given value.


### GetInputMessageAccountabilityData

`func (o *Wire) GetInputMessageAccountabilityData() string`

GetInputMessageAccountabilityData returns the InputMessageAccountabilityData field if non-nil, zero value otherwise.

### GetInputMessageAccountabilityDataOk

`func (o *Wire) GetInputMessageAccountabilityDataOk() (*string, bool)`

GetInputMessageAccountabilityDataOk returns a tuple with the InputMessageAccountabilityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageAccountabilityData

`func (o *Wire) SetInputMessageAccountabilityData(v string)`

SetInputMessageAccountabilityData sets InputMessageAccountabilityData field to given value.

### HasInputMessageAccountabilityData

`func (o *Wire) HasInputMessageAccountabilityData() bool`

HasInputMessageAccountabilityData returns a boolean if a field has been set.

### GetIsBulk

`func (o *Wire) GetIsBulk() bool`

GetIsBulk returns the IsBulk field if non-nil, zero value otherwise.

### GetIsBulkOk

`func (o *Wire) GetIsBulkOk() (*bool, bool)`

GetIsBulkOk returns a tuple with the IsBulk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBulk

`func (o *Wire) SetIsBulk(v bool)`

SetIsBulk sets IsBulk field to given value.


### GetLastUpdatedTime

`func (o *Wire) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Wire) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Wire) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetNetwork

`func (o *Wire) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Wire) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Wire) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Wire) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *Wire) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *Wire) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *Wire) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.

### HasOriginatingAccountId

`func (o *Wire) HasOriginatingAccountId() bool`

HasOriginatingAccountId returns a boolean if a field has been set.

### GetOriginatingAccountNumber

`func (o *Wire) GetOriginatingAccountNumber() string`

GetOriginatingAccountNumber returns the OriginatingAccountNumber field if non-nil, zero value otherwise.

### GetOriginatingAccountNumberOk

`func (o *Wire) GetOriginatingAccountNumberOk() (*string, bool)`

GetOriginatingAccountNumberOk returns a tuple with the OriginatingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountNumber

`func (o *Wire) SetOriginatingAccountNumber(v string)`

SetOriginatingAccountNumber sets OriginatingAccountNumber field to given value.


### GetReceivingAccountId

`func (o *Wire) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *Wire) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *Wire) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.

### HasReceivingAccountId

`func (o *Wire) HasReceivingAccountId() bool`

HasReceivingAccountId returns a boolean if a field has been set.

### GetReceivingAccountNumber

`func (o *Wire) GetReceivingAccountNumber() string`

GetReceivingAccountNumber returns the ReceivingAccountNumber field if non-nil, zero value otherwise.

### GetReceivingAccountNumberOk

`func (o *Wire) GetReceivingAccountNumberOk() (*string, bool)`

GetReceivingAccountNumberOk returns a tuple with the ReceivingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountNumber

`func (o *Wire) SetReceivingAccountNumber(v string)`

SetReceivingAccountNumber sets ReceivingAccountNumber field to given value.


### GetRecipientMessage

`func (o *Wire) GetRecipientMessage() string`

GetRecipientMessage returns the RecipientMessage field if non-nil, zero value otherwise.

### GetRecipientMessageOk

`func (o *Wire) GetRecipientMessageOk() (*string, bool)`

GetRecipientMessageOk returns a tuple with the RecipientMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientMessage

`func (o *Wire) SetRecipientMessage(v string)`

SetRecipientMessage sets RecipientMessage field to given value.

### HasRecipientMessage

`func (o *Wire) HasRecipientMessage() bool`

HasRecipientMessage returns a boolean if a field has been set.

### GetReturnData

`func (o *Wire) GetReturnData() ReturnData1`

GetReturnData returns the ReturnData field if non-nil, zero value otherwise.

### GetReturnDataOk

`func (o *Wire) GetReturnDataOk() (*ReturnData1, bool)`

GetReturnDataOk returns a tuple with the ReturnData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnData

`func (o *Wire) SetReturnData(v ReturnData1)`

SetReturnData sets ReturnData field to given value.

### HasReturnData

`func (o *Wire) HasReturnData() bool`

HasReturnData returns a boolean if a field has been set.

### GetSenderReferenceId

`func (o *Wire) GetSenderReferenceId() string`

GetSenderReferenceId returns the SenderReferenceId field if non-nil, zero value otherwise.

### GetSenderReferenceIdOk

`func (o *Wire) GetSenderReferenceIdOk() (*string, bool)`

GetSenderReferenceIdOk returns a tuple with the SenderReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderReferenceId

`func (o *Wire) SetSenderReferenceId(v string)`

SetSenderReferenceId sets SenderReferenceId field to given value.


### GetSettlementDate

`func (o *Wire) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *Wire) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *Wire) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *Wire) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetStatus

`func (o *Wire) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Wire) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Wire) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusDetails

`func (o *Wire) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *Wire) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *Wire) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *Wire) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetTransactionId

`func (o *Wire) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Wire) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Wire) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTransactionInId

`func (o *Wire) GetTransactionInId() string`

GetTransactionInId returns the TransactionInId field if non-nil, zero value otherwise.

### GetTransactionInIdOk

`func (o *Wire) GetTransactionInIdOk() (*string, bool)`

GetTransactionInIdOk returns a tuple with the TransactionInId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionInId

`func (o *Wire) SetTransactionInId(v string)`

SetTransactionInId sets TransactionInId field to given value.

### HasTransactionInId

`func (o *Wire) HasTransactionInId() bool`

HasTransactionInId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


