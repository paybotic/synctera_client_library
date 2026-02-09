# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account uuid associated with the transaction. &#x60;account_id&#x60; and &#x60;internal_account_id&#x60; are mutually exclusive | [optional] 
**Amount** | **int64** | The total amount of the transaction including both pending and already posted amounts. The value is represented as the smallest denomination of the applicable currency. | 
**CreationTime** | **time.Time** | The exact time the transaction was recorded in the ledger | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**CustomerId** | Pointer to **string** | The uuid of the customer that initiated the transaction (if any)  | [optional] 
**DcSign** | **string** | The &#x60;dc_sign&#x60; represents the direction money was moved. A value of &#x60;DEBIT&#x60; is money moving out of an account, a value of &#x60;CREDIT&#x60; is money moving into an account | 
**Decline** | Pointer to [**DeclineInfo**](DeclineInfo.md) |  | [optional] 
**Description** | Pointer to **string** | A human-friendly description of the transaction, provided by the integrator | [optional] 
**EffectiveDate** | **string** | The effective date of the transaction. This usually aligns with network settlement date, which differs between transaction types. The effective date is also used to determine effective daily balances for the purposes of interest calculation. | 
**EnhancedTransaction** | Pointer to [**EnhancedTransactionData**](EnhancedTransactionData.md) |  | [optional] 
**ForcePost** | **bool** | Determines whether or not a transaction or auth was \&quot;forced\&quot; or not. A forced transaction skips any account balance checks | 
**FraudRisk** | Pointer to [**FraudRiskData**](FraudRiskData.md) |  | [optional] 
**GroupId** | **string** | The group id of the transaction. Every transaction in the ledger is one entry in a double-entry system and the primary and offset transactions share the same &#x60;group_id&#x60; | 
**HoldExpirationTime** | Pointer to **time.Time** | The date and time any pending amount is expected to be released back to the account. | [optional] 
**Id** | **string** | The unique identifier of the transaction | 
**InternalAccountId** | Pointer to **string** | The internal account uuid associated with the transaction. &#x60;account_id&#x60; and &#x60;internal_account_id&#x60; are mutually exclusive | [optional] 
**LastUpdatedTime** | **time.Time** | The date and time the transaction was last modified | 
**Metadata** | Pointer to **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | [optional] 
**PartialHold** | **bool** | Determines whether or not the funds on hold were the result of a partial auth or not. If &#x60;true&#x60; the &#x60;pending_amount&#x60; of the transaction will be less than the requested amount. This is primarily used for certain types of card transactions. | 
**PendingAmount** | **int64** | The amount amount of the transaction currently authorized or on hold | 
**PostedAmount** | **int64** | The amount of the transaction that has been fully posted to the account | 
**PostedDate** | Pointer to **string** | The date the transaction was posted (based on the bank calendar and end-of-day). For transaction with multiple postings, this is the date of the earliest posting. This will be omitted for transactions with a &#x60;posted_amount&#x60; of &#x60;0&#x60;. | [optional] 
**ReferenceId** | Pointer to **string** | An external ID provided by the payment network to represent this transaction. This is not guaranteed to be globally unique. This will always be omitted for internal transfers. | [optional] 
**Status** | [**TransactionStatuses**](TransactionStatuses.md) |  | 
**Subtype** | **string** | The subtype of the transaction. For the full list of available values, see the spec for &#x60;GET /v1/transactions&#x60;. | 
**SystemDescription** | Pointer to **string** | A human-friendly description of the transaction, provided by the Synctera platform | [optional] 
**TransactionTime** | **time.Time** | The time the transaction occurred. In most cases this will be roughly identical to creation_time, but it can differ in some situations if the payment doesn&#39;t appear in the ledger in real-time. | 
**Type** | **string** | The type of the transaction. For the full list of available values, see the spec for &#x60;GET /v1/transactions&#x60;. | 

## Methods

### NewTransaction

`func NewTransaction(amount int64, creationTime time.Time, currency string, dcSign string, effectiveDate string, forcePost bool, groupId string, id string, lastUpdatedTime time.Time, partialHold bool, pendingAmount int64, postedAmount int64, status TransactionStatuses, subtype string, transactionTime time.Time, type_ string, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Transaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Transaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Transaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Transaction) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *Transaction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCreationTime

`func (o *Transaction) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Transaction) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Transaction) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *Transaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *Transaction) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Transaction) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Transaction) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Transaction) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDcSign

`func (o *Transaction) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *Transaction) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *Transaction) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetDecline

`func (o *Transaction) GetDecline() DeclineInfo`

GetDecline returns the Decline field if non-nil, zero value otherwise.

### GetDeclineOk

`func (o *Transaction) GetDeclineOk() (*DeclineInfo, bool)`

GetDeclineOk returns a tuple with the Decline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecline

`func (o *Transaction) SetDecline(v DeclineInfo)`

SetDecline sets Decline field to given value.

### HasDecline

`func (o *Transaction) HasDecline() bool`

HasDecline returns a boolean if a field has been set.

### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *Transaction) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *Transaction) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *Transaction) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetEnhancedTransaction

`func (o *Transaction) GetEnhancedTransaction() EnhancedTransactionData`

GetEnhancedTransaction returns the EnhancedTransaction field if non-nil, zero value otherwise.

### GetEnhancedTransactionOk

`func (o *Transaction) GetEnhancedTransactionOk() (*EnhancedTransactionData, bool)`

GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedTransaction

`func (o *Transaction) SetEnhancedTransaction(v EnhancedTransactionData)`

SetEnhancedTransaction sets EnhancedTransaction field to given value.

### HasEnhancedTransaction

`func (o *Transaction) HasEnhancedTransaction() bool`

HasEnhancedTransaction returns a boolean if a field has been set.

### GetForcePost

`func (o *Transaction) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *Transaction) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *Transaction) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.


### GetFraudRisk

`func (o *Transaction) GetFraudRisk() FraudRiskData`

GetFraudRisk returns the FraudRisk field if non-nil, zero value otherwise.

### GetFraudRiskOk

`func (o *Transaction) GetFraudRiskOk() (*FraudRiskData, bool)`

GetFraudRiskOk returns a tuple with the FraudRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudRisk

`func (o *Transaction) SetFraudRisk(v FraudRiskData)`

SetFraudRisk sets FraudRisk field to given value.

### HasFraudRisk

`func (o *Transaction) HasFraudRisk() bool`

HasFraudRisk returns a boolean if a field has been set.

### GetGroupId

`func (o *Transaction) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Transaction) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Transaction) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetHoldExpirationTime

`func (o *Transaction) GetHoldExpirationTime() time.Time`

GetHoldExpirationTime returns the HoldExpirationTime field if non-nil, zero value otherwise.

### GetHoldExpirationTimeOk

`func (o *Transaction) GetHoldExpirationTimeOk() (*time.Time, bool)`

GetHoldExpirationTimeOk returns a tuple with the HoldExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldExpirationTime

`func (o *Transaction) SetHoldExpirationTime(v time.Time)`

SetHoldExpirationTime sets HoldExpirationTime field to given value.

### HasHoldExpirationTime

`func (o *Transaction) HasHoldExpirationTime() bool`

HasHoldExpirationTime returns a boolean if a field has been set.

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.


### GetInternalAccountId

`func (o *Transaction) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *Transaction) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *Transaction) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.

### HasInternalAccountId

`func (o *Transaction) HasInternalAccountId() bool`

HasInternalAccountId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Transaction) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Transaction) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Transaction) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMetadata

`func (o *Transaction) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Transaction) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Transaction) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Transaction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Transaction) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Transaction) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPartialHold

`func (o *Transaction) GetPartialHold() bool`

GetPartialHold returns the PartialHold field if non-nil, zero value otherwise.

### GetPartialHoldOk

`func (o *Transaction) GetPartialHoldOk() (*bool, bool)`

GetPartialHoldOk returns a tuple with the PartialHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialHold

`func (o *Transaction) SetPartialHold(v bool)`

SetPartialHold sets PartialHold field to given value.


### GetPendingAmount

`func (o *Transaction) GetPendingAmount() int64`

GetPendingAmount returns the PendingAmount field if non-nil, zero value otherwise.

### GetPendingAmountOk

`func (o *Transaction) GetPendingAmountOk() (*int64, bool)`

GetPendingAmountOk returns a tuple with the PendingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmount

`func (o *Transaction) SetPendingAmount(v int64)`

SetPendingAmount sets PendingAmount field to given value.


### GetPostedAmount

`func (o *Transaction) GetPostedAmount() int64`

GetPostedAmount returns the PostedAmount field if non-nil, zero value otherwise.

### GetPostedAmountOk

`func (o *Transaction) GetPostedAmountOk() (*int64, bool)`

GetPostedAmountOk returns a tuple with the PostedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAmount

`func (o *Transaction) SetPostedAmount(v int64)`

SetPostedAmount sets PostedAmount field to given value.


### GetPostedDate

`func (o *Transaction) GetPostedDate() string`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *Transaction) GetPostedDateOk() (*string, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *Transaction) SetPostedDate(v string)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *Transaction) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *Transaction) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Transaction) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Transaction) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Transaction) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() TransactionStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*TransactionStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v TransactionStatuses)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *Transaction) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *Transaction) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *Transaction) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetSystemDescription

`func (o *Transaction) GetSystemDescription() string`

GetSystemDescription returns the SystemDescription field if non-nil, zero value otherwise.

### GetSystemDescriptionOk

`func (o *Transaction) GetSystemDescriptionOk() (*string, bool)`

GetSystemDescriptionOk returns a tuple with the SystemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDescription

`func (o *Transaction) SetSystemDescription(v string)`

SetSystemDescription sets SystemDescription field to given value.

### HasSystemDescription

`func (o *Transaction) HasSystemDescription() bool`

HasSystemDescription returns a boolean if a field has been set.

### GetTransactionTime

`func (o *Transaction) GetTransactionTime() time.Time`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *Transaction) GetTransactionTimeOk() (*time.Time, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *Transaction) SetTransactionTime(v time.Time)`

SetTransactionTime sets TransactionTime field to given value.


### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


