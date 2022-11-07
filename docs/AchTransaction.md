# AchTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account uuid associated with the transaction. &#x60;account_id&#x60; and &#x60;internal_account_id&#x60; are mutually exclusive | [optional] 
**Amount** | **int64** | The total amount of the transaction including both pending and already posted amounts. The value is represented as the smallest denomination of the applicable currency. | 
**CreationTime** | **time.Time** | The exact time the transaction was recorded in the ledger | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**CustomerId** | Pointer to **string** | The uuid of the customer that initiated the transaction (if any)  | [optional] 
**DcSign** | **string** | The &#x60;dc_sign&#x60; represents the direction money was moved. A value of &#x60;DEBIT&#x60; is money moving out of an account, a value of &#x60;CREDIT&#x60; is money moving into an account | 
**Decline** | Pointer to [**BaseTransactionDecline**](BaseTransactionDecline.md) |  | [optional] 
**Description** | Pointer to **string** | A human-friendly description of the transaction, provided by the integrator | [optional] 
**EffectiveDate** | **string** | The effective date of the transaction. This usually aligns with network settlement date, which differs between transaction types. The effective date is also used to determine effective daily balances for the purposes of interest calculation. | 
**EnhancedTransaction** | Pointer to [**EnhancedTransactionData**](EnhancedTransactionData.md) |  | [optional] 
**ForcePost** | **bool** | Determines whether or not a transaction or auth was \&quot;forced\&quot; or not. A forced transaction skips any account balance checks | 
**GroupId** | **string** | The group id of the transaction. Every transaction in the ledger is one entry in a double-entry system and the primary and offset transactions share the same &#x60;group_id&#x60; | 
**HoldExpirationTime** | Pointer to **time.Time** | The date and time any pending amount is expected to be released back to the account. | [optional] 
**Id** | **string** | The unique identifier of the transaction | 
**InternalAccountId** | Pointer to **string** | The internal account uuid associated with the transaction. &#x60;account_id&#x60; and &#x60;internal_account_id&#x60; are mutually exclusive | [optional] 
**LastUpdatedTime** | **time.Time** | The date and time the transaction was last modified | 
**PartialHold** | **bool** | Determines whether or not the funds on hold were the result of a partial auth or not. If &#x60;true&#x60; the &#x60;pending_amount&#x60; of the transaction will be less than the requested amount. This is primarily used for certain types of card transactions. | 
**PendingAmount** | **int64** | The amount amount of the transaction currently authorized or on hold | 
**PostedAmount** | **int64** | The amount of the transaction that has been fully posted to the account | 
**PostedDate** | Pointer to **string** | The date the transaction was posted (based on the bank calendar and end-of-day). For transaction with multiple postings, this is the date of the earliest posting. This will be omitted for transactions with a &#x60;posted_amount&#x60; of &#x60;0&#x60;. | [optional] 
**ReferenceId** | Pointer to **string** | An external ID provided by the payment network to represent this transaction. This is not guaranteed to be globally unique. This will always be omitted for internal transfers. | [optional] 
**Status** | **string** | The status of the transaction | 
**SystemDescription** | Pointer to **string** | A human-friendly description of the transaction, provided by the Synctera platform | [optional] 
**TransactionTime** | **time.Time** | The time the transaction occurred. In most cases this will be roughly identical to creation_time, but it can differ in some situations if the payment doesn&#39;t appear in the ledger in real-time. | 
**Type** | **string** | The type of the transaction. This typically represents the \&quot;payment\&quot; rail that is used. For example, for ACH payments this will be &#x60;ach&#x60;, while debit card transactions will use &#x60;card&#x60;. | 
**AchTransaction** | [**AchTransactionData**](AchTransactionData.md) |  | 
**Subtype** | [**AchTransactionSubtypes**](AchTransactionSubtypes.md) |  | 

## Methods

### NewAchTransaction

`func NewAchTransaction(amount int64, creationTime time.Time, currency string, dcSign string, effectiveDate string, forcePost bool, groupId string, id string, lastUpdatedTime time.Time, partialHold bool, pendingAmount int64, postedAmount int64, status string, transactionTime time.Time, type_ string, achTransaction AchTransactionData, subtype AchTransactionSubtypes, ) *AchTransaction`

NewAchTransaction instantiates a new AchTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchTransactionWithDefaults

`func NewAchTransactionWithDefaults() *AchTransaction`

NewAchTransactionWithDefaults instantiates a new AchTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AchTransaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AchTransaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AchTransaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AchTransaction) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *AchTransaction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AchTransaction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AchTransaction) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCreationTime

`func (o *AchTransaction) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AchTransaction) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AchTransaction) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *AchTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AchTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AchTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *AchTransaction) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AchTransaction) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AchTransaction) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AchTransaction) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDcSign

`func (o *AchTransaction) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *AchTransaction) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *AchTransaction) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetDecline

`func (o *AchTransaction) GetDecline() BaseTransactionDecline`

GetDecline returns the Decline field if non-nil, zero value otherwise.

### GetDeclineOk

`func (o *AchTransaction) GetDeclineOk() (*BaseTransactionDecline, bool)`

GetDeclineOk returns a tuple with the Decline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecline

`func (o *AchTransaction) SetDecline(v BaseTransactionDecline)`

SetDecline sets Decline field to given value.

### HasDecline

`func (o *AchTransaction) HasDecline() bool`

HasDecline returns a boolean if a field has been set.

### GetDescription

`func (o *AchTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AchTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AchTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AchTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *AchTransaction) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AchTransaction) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AchTransaction) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetEnhancedTransaction

`func (o *AchTransaction) GetEnhancedTransaction() EnhancedTransactionData`

GetEnhancedTransaction returns the EnhancedTransaction field if non-nil, zero value otherwise.

### GetEnhancedTransactionOk

`func (o *AchTransaction) GetEnhancedTransactionOk() (*EnhancedTransactionData, bool)`

GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedTransaction

`func (o *AchTransaction) SetEnhancedTransaction(v EnhancedTransactionData)`

SetEnhancedTransaction sets EnhancedTransaction field to given value.

### HasEnhancedTransaction

`func (o *AchTransaction) HasEnhancedTransaction() bool`

HasEnhancedTransaction returns a boolean if a field has been set.

### GetForcePost

`func (o *AchTransaction) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *AchTransaction) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *AchTransaction) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.


### GetGroupId

`func (o *AchTransaction) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AchTransaction) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AchTransaction) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetHoldExpirationTime

`func (o *AchTransaction) GetHoldExpirationTime() time.Time`

GetHoldExpirationTime returns the HoldExpirationTime field if non-nil, zero value otherwise.

### GetHoldExpirationTimeOk

`func (o *AchTransaction) GetHoldExpirationTimeOk() (*time.Time, bool)`

GetHoldExpirationTimeOk returns a tuple with the HoldExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldExpirationTime

`func (o *AchTransaction) SetHoldExpirationTime(v time.Time)`

SetHoldExpirationTime sets HoldExpirationTime field to given value.

### HasHoldExpirationTime

`func (o *AchTransaction) HasHoldExpirationTime() bool`

HasHoldExpirationTime returns a boolean if a field has been set.

### GetId

`func (o *AchTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AchTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AchTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetInternalAccountId

`func (o *AchTransaction) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *AchTransaction) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *AchTransaction) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.

### HasInternalAccountId

`func (o *AchTransaction) HasInternalAccountId() bool`

HasInternalAccountId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AchTransaction) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AchTransaction) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AchTransaction) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetPartialHold

`func (o *AchTransaction) GetPartialHold() bool`

GetPartialHold returns the PartialHold field if non-nil, zero value otherwise.

### GetPartialHoldOk

`func (o *AchTransaction) GetPartialHoldOk() (*bool, bool)`

GetPartialHoldOk returns a tuple with the PartialHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialHold

`func (o *AchTransaction) SetPartialHold(v bool)`

SetPartialHold sets PartialHold field to given value.


### GetPendingAmount

`func (o *AchTransaction) GetPendingAmount() int64`

GetPendingAmount returns the PendingAmount field if non-nil, zero value otherwise.

### GetPendingAmountOk

`func (o *AchTransaction) GetPendingAmountOk() (*int64, bool)`

GetPendingAmountOk returns a tuple with the PendingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmount

`func (o *AchTransaction) SetPendingAmount(v int64)`

SetPendingAmount sets PendingAmount field to given value.


### GetPostedAmount

`func (o *AchTransaction) GetPostedAmount() int64`

GetPostedAmount returns the PostedAmount field if non-nil, zero value otherwise.

### GetPostedAmountOk

`func (o *AchTransaction) GetPostedAmountOk() (*int64, bool)`

GetPostedAmountOk returns a tuple with the PostedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAmount

`func (o *AchTransaction) SetPostedAmount(v int64)`

SetPostedAmount sets PostedAmount field to given value.


### GetPostedDate

`func (o *AchTransaction) GetPostedDate() string`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *AchTransaction) GetPostedDateOk() (*string, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *AchTransaction) SetPostedDate(v string)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *AchTransaction) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *AchTransaction) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *AchTransaction) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *AchTransaction) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *AchTransaction) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetStatus

`func (o *AchTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AchTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSystemDescription

`func (o *AchTransaction) GetSystemDescription() string`

GetSystemDescription returns the SystemDescription field if non-nil, zero value otherwise.

### GetSystemDescriptionOk

`func (o *AchTransaction) GetSystemDescriptionOk() (*string, bool)`

GetSystemDescriptionOk returns a tuple with the SystemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDescription

`func (o *AchTransaction) SetSystemDescription(v string)`

SetSystemDescription sets SystemDescription field to given value.

### HasSystemDescription

`func (o *AchTransaction) HasSystemDescription() bool`

HasSystemDescription returns a boolean if a field has been set.

### GetTransactionTime

`func (o *AchTransaction) GetTransactionTime() time.Time`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *AchTransaction) GetTransactionTimeOk() (*time.Time, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *AchTransaction) SetTransactionTime(v time.Time)`

SetTransactionTime sets TransactionTime field to given value.


### GetType

`func (o *AchTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AchTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AchTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetAchTransaction

`func (o *AchTransaction) GetAchTransaction() AchTransactionData`

GetAchTransaction returns the AchTransaction field if non-nil, zero value otherwise.

### GetAchTransactionOk

`func (o *AchTransaction) GetAchTransactionOk() (*AchTransactionData, bool)`

GetAchTransactionOk returns a tuple with the AchTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchTransaction

`func (o *AchTransaction) SetAchTransaction(v AchTransactionData)`

SetAchTransaction sets AchTransaction field to given value.


### GetSubtype

`func (o *AchTransaction) GetSubtype() AchTransactionSubtypes`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AchTransaction) GetSubtypeOk() (*AchTransactionSubtypes, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AchTransaction) SetSubtype(v AchTransactionSubtypes)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


