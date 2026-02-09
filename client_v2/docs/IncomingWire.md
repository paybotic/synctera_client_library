# IncomingWire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Transfer amount in cents ($100 would be 10000) | 
**BankMessage** | Pointer to **string** | Information from the sending bank to the receiving bank. | [optional] 
**CreationTime** | **time.Time** |  | [readonly] 
**Currency** | Pointer to **string** | 3-character currency code | [optional] 
**CustomerId** | Pointer to **string** | The customer UUID representing the person initiating the Wire transfer | [optional] 
**DeclineReason** | Pointer to **string** | Description of any errors or problems encountered during processing. | [optional] 
**EffectiveDate** | Pointer to **string** | The effective date of the transaction | [optional] 
**EndToEndIdentification** | Pointer to **string** | A mandatory reference assigned by the customer originating the payment, i.e., Debtor or Ultimate Debtor that must be transported unchanged throughout the payment chain and delivered to the end-customer, i.e., Creditor or Ultimate Creditor to facilitate the customer’s reconciliation. | [optional] 
**Id** | **string** | wire ID | [readonly] 
**InputMessageAccountabilityData** | Pointer to **string** | The IMAD of the incoming wire. The field is called Message Identification in the ISO format. | [optional] [readonly] 
**InstructionIdentification** | Pointer to **string** | An optional reference that the Fedwire Funds participant sending the message may include in a customer credit transfers sent across the Fedwire Funds Service and will be delivered unchanged to the Fedwire Funds participant receiving the message. | [optional] 
**IsReturn** | **bool** | Indicates if the wire is a return of an outgoing wire | 
**Network** | Pointer to **string** | The network used to process the wire | [optional] 
**OriginatingAccountNumber** | Pointer to **string** | The account number of the sender, as provided in the wire file. | [optional] 
**Receiver** | [**Party**](Party.md) |  | 
**ReceivingAccountId** | Pointer to **string** | The internal Synctera account uuid representing the recipient of the wire. | [optional] 
**ReceivingAccountNumber** | **string** | The account number of the recipient, as provided in the wire file. | 
**RecipientMessage** | Pointer to **string** | Information from the originator to the beneficiary (recipient). | [optional] 
**ReturnData** | Pointer to [**ReturnData**](ReturnData.md) |  | [optional] 
**ReturnReason** | Pointer to **string** | The reason we are returning this wire. | [optional] 
**Sender** | [**Party**](Party.md) |  | 
**SenderReferenceId** | **string** | Sender&#39;s id associated with fedwire transfer | [readonly] 
**SettlementDate** | Pointer to **string** | The settlement date of the transaction | [optional] 
**Status** | **string** | The current status of the transfer | [readonly] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**TransactionId** | Pointer to **string** | ID of the resulting transaction resource | [optional] [readonly] 
**TransactionIdentification** | Pointer to **string** | An optional reference that the first financial institution in the end-to-end payment chain may assign to the customer credit transfer and that should remain unchanged throughout the end-to-end funds-transfer chain of financial institutions. | [optional] 
**TypeSubtype** | Pointer to **string** | The fedwire label associated with the subtype code | [optional] 
**Uetr** | Pointer to **string** | Unique End-To-End Reference (UETR) is a mandatory reference in line with the Universally Unique Identifier (UUID) version 4 standard. The UETR is used to facilitate end-to-end payments tracking. | [optional] 

## Methods

### NewIncomingWire

`func NewIncomingWire(amount int32, creationTime time.Time, id string, isReturn bool, receiver Party, receivingAccountNumber string, sender Party, senderReferenceId string, status string, tenant string, ) *IncomingWire`

NewIncomingWire instantiates a new IncomingWire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingWireWithDefaults

`func NewIncomingWireWithDefaults() *IncomingWire`

NewIncomingWireWithDefaults instantiates a new IncomingWire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *IncomingWire) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IncomingWire) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IncomingWire) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetBankMessage

`func (o *IncomingWire) GetBankMessage() string`

GetBankMessage returns the BankMessage field if non-nil, zero value otherwise.

### GetBankMessageOk

`func (o *IncomingWire) GetBankMessageOk() (*string, bool)`

GetBankMessageOk returns a tuple with the BankMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankMessage

`func (o *IncomingWire) SetBankMessage(v string)`

SetBankMessage sets BankMessage field to given value.

### HasBankMessage

`func (o *IncomingWire) HasBankMessage() bool`

HasBankMessage returns a boolean if a field has been set.

### GetCreationTime

`func (o *IncomingWire) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *IncomingWire) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *IncomingWire) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *IncomingWire) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IncomingWire) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IncomingWire) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IncomingWire) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *IncomingWire) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *IncomingWire) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *IncomingWire) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *IncomingWire) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDeclineReason

`func (o *IncomingWire) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *IncomingWire) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *IncomingWire) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *IncomingWire) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *IncomingWire) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *IncomingWire) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *IncomingWire) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *IncomingWire) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetEndToEndIdentification

`func (o *IncomingWire) GetEndToEndIdentification() string`

GetEndToEndIdentification returns the EndToEndIdentification field if non-nil, zero value otherwise.

### GetEndToEndIdentificationOk

`func (o *IncomingWire) GetEndToEndIdentificationOk() (*string, bool)`

GetEndToEndIdentificationOk returns a tuple with the EndToEndIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndIdentification

`func (o *IncomingWire) SetEndToEndIdentification(v string)`

SetEndToEndIdentification sets EndToEndIdentification field to given value.

### HasEndToEndIdentification

`func (o *IncomingWire) HasEndToEndIdentification() bool`

HasEndToEndIdentification returns a boolean if a field has been set.

### GetId

`func (o *IncomingWire) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncomingWire) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncomingWire) SetId(v string)`

SetId sets Id field to given value.


### GetInputMessageAccountabilityData

`func (o *IncomingWire) GetInputMessageAccountabilityData() string`

GetInputMessageAccountabilityData returns the InputMessageAccountabilityData field if non-nil, zero value otherwise.

### GetInputMessageAccountabilityDataOk

`func (o *IncomingWire) GetInputMessageAccountabilityDataOk() (*string, bool)`

GetInputMessageAccountabilityDataOk returns a tuple with the InputMessageAccountabilityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputMessageAccountabilityData

`func (o *IncomingWire) SetInputMessageAccountabilityData(v string)`

SetInputMessageAccountabilityData sets InputMessageAccountabilityData field to given value.

### HasInputMessageAccountabilityData

`func (o *IncomingWire) HasInputMessageAccountabilityData() bool`

HasInputMessageAccountabilityData returns a boolean if a field has been set.

### GetInstructionIdentification

`func (o *IncomingWire) GetInstructionIdentification() string`

GetInstructionIdentification returns the InstructionIdentification field if non-nil, zero value otherwise.

### GetInstructionIdentificationOk

`func (o *IncomingWire) GetInstructionIdentificationOk() (*string, bool)`

GetInstructionIdentificationOk returns a tuple with the InstructionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionIdentification

`func (o *IncomingWire) SetInstructionIdentification(v string)`

SetInstructionIdentification sets InstructionIdentification field to given value.

### HasInstructionIdentification

`func (o *IncomingWire) HasInstructionIdentification() bool`

HasInstructionIdentification returns a boolean if a field has been set.

### GetIsReturn

`func (o *IncomingWire) GetIsReturn() bool`

GetIsReturn returns the IsReturn field if non-nil, zero value otherwise.

### GetIsReturnOk

`func (o *IncomingWire) GetIsReturnOk() (*bool, bool)`

GetIsReturnOk returns a tuple with the IsReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReturn

`func (o *IncomingWire) SetIsReturn(v bool)`

SetIsReturn sets IsReturn field to given value.


### GetNetwork

`func (o *IncomingWire) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IncomingWire) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IncomingWire) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IncomingWire) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOriginatingAccountNumber

`func (o *IncomingWire) GetOriginatingAccountNumber() string`

GetOriginatingAccountNumber returns the OriginatingAccountNumber field if non-nil, zero value otherwise.

### GetOriginatingAccountNumberOk

`func (o *IncomingWire) GetOriginatingAccountNumberOk() (*string, bool)`

GetOriginatingAccountNumberOk returns a tuple with the OriginatingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountNumber

`func (o *IncomingWire) SetOriginatingAccountNumber(v string)`

SetOriginatingAccountNumber sets OriginatingAccountNumber field to given value.

### HasOriginatingAccountNumber

`func (o *IncomingWire) HasOriginatingAccountNumber() bool`

HasOriginatingAccountNumber returns a boolean if a field has been set.

### GetReceiver

`func (o *IncomingWire) GetReceiver() Party`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *IncomingWire) GetReceiverOk() (*Party, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *IncomingWire) SetReceiver(v Party)`

SetReceiver sets Receiver field to given value.


### GetReceivingAccountId

`func (o *IncomingWire) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *IncomingWire) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *IncomingWire) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.

### HasReceivingAccountId

`func (o *IncomingWire) HasReceivingAccountId() bool`

HasReceivingAccountId returns a boolean if a field has been set.

### GetReceivingAccountNumber

`func (o *IncomingWire) GetReceivingAccountNumber() string`

GetReceivingAccountNumber returns the ReceivingAccountNumber field if non-nil, zero value otherwise.

### GetReceivingAccountNumberOk

`func (o *IncomingWire) GetReceivingAccountNumberOk() (*string, bool)`

GetReceivingAccountNumberOk returns a tuple with the ReceivingAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountNumber

`func (o *IncomingWire) SetReceivingAccountNumber(v string)`

SetReceivingAccountNumber sets ReceivingAccountNumber field to given value.


### GetRecipientMessage

`func (o *IncomingWire) GetRecipientMessage() string`

GetRecipientMessage returns the RecipientMessage field if non-nil, zero value otherwise.

### GetRecipientMessageOk

`func (o *IncomingWire) GetRecipientMessageOk() (*string, bool)`

GetRecipientMessageOk returns a tuple with the RecipientMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientMessage

`func (o *IncomingWire) SetRecipientMessage(v string)`

SetRecipientMessage sets RecipientMessage field to given value.

### HasRecipientMessage

`func (o *IncomingWire) HasRecipientMessage() bool`

HasRecipientMessage returns a boolean if a field has been set.

### GetReturnData

`func (o *IncomingWire) GetReturnData() ReturnData`

GetReturnData returns the ReturnData field if non-nil, zero value otherwise.

### GetReturnDataOk

`func (o *IncomingWire) GetReturnDataOk() (*ReturnData, bool)`

GetReturnDataOk returns a tuple with the ReturnData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnData

`func (o *IncomingWire) SetReturnData(v ReturnData)`

SetReturnData sets ReturnData field to given value.

### HasReturnData

`func (o *IncomingWire) HasReturnData() bool`

HasReturnData returns a boolean if a field has been set.

### GetReturnReason

`func (o *IncomingWire) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *IncomingWire) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *IncomingWire) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.

### HasReturnReason

`func (o *IncomingWire) HasReturnReason() bool`

HasReturnReason returns a boolean if a field has been set.

### GetSender

`func (o *IncomingWire) GetSender() Party`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *IncomingWire) GetSenderOk() (*Party, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *IncomingWire) SetSender(v Party)`

SetSender sets Sender field to given value.


### GetSenderReferenceId

`func (o *IncomingWire) GetSenderReferenceId() string`

GetSenderReferenceId returns the SenderReferenceId field if non-nil, zero value otherwise.

### GetSenderReferenceIdOk

`func (o *IncomingWire) GetSenderReferenceIdOk() (*string, bool)`

GetSenderReferenceIdOk returns a tuple with the SenderReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderReferenceId

`func (o *IncomingWire) SetSenderReferenceId(v string)`

SetSenderReferenceId sets SenderReferenceId field to given value.


### GetSettlementDate

`func (o *IncomingWire) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *IncomingWire) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *IncomingWire) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *IncomingWire) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetStatus

`func (o *IncomingWire) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncomingWire) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncomingWire) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *IncomingWire) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IncomingWire) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IncomingWire) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionId

`func (o *IncomingWire) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *IncomingWire) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *IncomingWire) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *IncomingWire) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionIdentification

`func (o *IncomingWire) GetTransactionIdentification() string`

GetTransactionIdentification returns the TransactionIdentification field if non-nil, zero value otherwise.

### GetTransactionIdentificationOk

`func (o *IncomingWire) GetTransactionIdentificationOk() (*string, bool)`

GetTransactionIdentificationOk returns a tuple with the TransactionIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentification

`func (o *IncomingWire) SetTransactionIdentification(v string)`

SetTransactionIdentification sets TransactionIdentification field to given value.

### HasTransactionIdentification

`func (o *IncomingWire) HasTransactionIdentification() bool`

HasTransactionIdentification returns a boolean if a field has been set.

### GetTypeSubtype

`func (o *IncomingWire) GetTypeSubtype() string`

GetTypeSubtype returns the TypeSubtype field if non-nil, zero value otherwise.

### GetTypeSubtypeOk

`func (o *IncomingWire) GetTypeSubtypeOk() (*string, bool)`

GetTypeSubtypeOk returns a tuple with the TypeSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSubtype

`func (o *IncomingWire) SetTypeSubtype(v string)`

SetTypeSubtype sets TypeSubtype field to given value.

### HasTypeSubtype

`func (o *IncomingWire) HasTypeSubtype() bool`

HasTypeSubtype returns a boolean if a field has been set.

### GetUetr

`func (o *IncomingWire) GetUetr() string`

GetUetr returns the Uetr field if non-nil, zero value otherwise.

### GetUetrOk

`func (o *IncomingWire) GetUetrOk() (*string, bool)`

GetUetrOk returns a tuple with the Uetr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUetr

`func (o *IncomingWire) SetUetr(v string)`

SetUetr sets Uetr field to given value.

### HasUetr

`func (o *IncomingWire) HasUetr() bool`

HasUetr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


