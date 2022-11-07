# CheckTransaction

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
**CheckTransaction** | [**CheckTransactionData**](CheckTransactionData.md) |  | 
**Subtype** | [**CheckTransactionSubtypes**](CheckTransactionSubtypes.md) |  | 

## Methods

### NewCheckTransaction

`func NewCheckTransaction(amount int64, creationTime time.Time, currency string, dcSign string, effectiveDate string, forcePost bool, groupId string, id string, lastUpdatedTime time.Time, partialHold bool, pendingAmount int64, postedAmount int64, status string, transactionTime time.Time, type_ string, checkTransaction CheckTransactionData, subtype CheckTransactionSubtypes, ) *CheckTransaction`

NewCheckTransaction instantiates a new CheckTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTransactionWithDefaults

`func NewCheckTransactionWithDefaults() *CheckTransaction`

NewCheckTransactionWithDefaults instantiates a new CheckTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CheckTransaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CheckTransaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CheckTransaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CheckTransaction) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *CheckTransaction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckTransaction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckTransaction) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCreationTime

`func (o *CheckTransaction) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CheckTransaction) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CheckTransaction) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *CheckTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *CheckTransaction) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CheckTransaction) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CheckTransaction) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CheckTransaction) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDcSign

`func (o *CheckTransaction) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *CheckTransaction) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *CheckTransaction) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetDecline

`func (o *CheckTransaction) GetDecline() BaseTransactionDecline`

GetDecline returns the Decline field if non-nil, zero value otherwise.

### GetDeclineOk

`func (o *CheckTransaction) GetDeclineOk() (*BaseTransactionDecline, bool)`

GetDeclineOk returns a tuple with the Decline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecline

`func (o *CheckTransaction) SetDecline(v BaseTransactionDecline)`

SetDecline sets Decline field to given value.

### HasDecline

`func (o *CheckTransaction) HasDecline() bool`

HasDecline returns a boolean if a field has been set.

### GetDescription

`func (o *CheckTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CheckTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CheckTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CheckTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *CheckTransaction) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *CheckTransaction) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *CheckTransaction) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetEnhancedTransaction

`func (o *CheckTransaction) GetEnhancedTransaction() EnhancedTransactionData`

GetEnhancedTransaction returns the EnhancedTransaction field if non-nil, zero value otherwise.

### GetEnhancedTransactionOk

`func (o *CheckTransaction) GetEnhancedTransactionOk() (*EnhancedTransactionData, bool)`

GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedTransaction

`func (o *CheckTransaction) SetEnhancedTransaction(v EnhancedTransactionData)`

SetEnhancedTransaction sets EnhancedTransaction field to given value.

### HasEnhancedTransaction

`func (o *CheckTransaction) HasEnhancedTransaction() bool`

HasEnhancedTransaction returns a boolean if a field has been set.

### GetForcePost

`func (o *CheckTransaction) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *CheckTransaction) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *CheckTransaction) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.


### GetGroupId

`func (o *CheckTransaction) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CheckTransaction) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CheckTransaction) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetHoldExpirationTime

`func (o *CheckTransaction) GetHoldExpirationTime() time.Time`

GetHoldExpirationTime returns the HoldExpirationTime field if non-nil, zero value otherwise.

### GetHoldExpirationTimeOk

`func (o *CheckTransaction) GetHoldExpirationTimeOk() (*time.Time, bool)`

GetHoldExpirationTimeOk returns a tuple with the HoldExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldExpirationTime

`func (o *CheckTransaction) SetHoldExpirationTime(v time.Time)`

SetHoldExpirationTime sets HoldExpirationTime field to given value.

### HasHoldExpirationTime

`func (o *CheckTransaction) HasHoldExpirationTime() bool`

HasHoldExpirationTime returns a boolean if a field has been set.

### GetId

`func (o *CheckTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetInternalAccountId

`func (o *CheckTransaction) GetInternalAccountId() string`

GetInternalAccountId returns the InternalAccountId field if non-nil, zero value otherwise.

### GetInternalAccountIdOk

`func (o *CheckTransaction) GetInternalAccountIdOk() (*string, bool)`

GetInternalAccountIdOk returns a tuple with the InternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAccountId

`func (o *CheckTransaction) SetInternalAccountId(v string)`

SetInternalAccountId sets InternalAccountId field to given value.

### HasInternalAccountId

`func (o *CheckTransaction) HasInternalAccountId() bool`

HasInternalAccountId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *CheckTransaction) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CheckTransaction) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CheckTransaction) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetPartialHold

`func (o *CheckTransaction) GetPartialHold() bool`

GetPartialHold returns the PartialHold field if non-nil, zero value otherwise.

### GetPartialHoldOk

`func (o *CheckTransaction) GetPartialHoldOk() (*bool, bool)`

GetPartialHoldOk returns a tuple with the PartialHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialHold

`func (o *CheckTransaction) SetPartialHold(v bool)`

SetPartialHold sets PartialHold field to given value.


### GetPendingAmount

`func (o *CheckTransaction) GetPendingAmount() int64`

GetPendingAmount returns the PendingAmount field if non-nil, zero value otherwise.

### GetPendingAmountOk

`func (o *CheckTransaction) GetPendingAmountOk() (*int64, bool)`

GetPendingAmountOk returns a tuple with the PendingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmount

`func (o *CheckTransaction) SetPendingAmount(v int64)`

SetPendingAmount sets PendingAmount field to given value.


### GetPostedAmount

`func (o *CheckTransaction) GetPostedAmount() int64`

GetPostedAmount returns the PostedAmount field if non-nil, zero value otherwise.

### GetPostedAmountOk

`func (o *CheckTransaction) GetPostedAmountOk() (*int64, bool)`

GetPostedAmountOk returns a tuple with the PostedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAmount

`func (o *CheckTransaction) SetPostedAmount(v int64)`

SetPostedAmount sets PostedAmount field to given value.


### GetPostedDate

`func (o *CheckTransaction) GetPostedDate() string`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *CheckTransaction) GetPostedDateOk() (*string, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *CheckTransaction) SetPostedDate(v string)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *CheckTransaction) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetReferenceId

`func (o *CheckTransaction) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CheckTransaction) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CheckTransaction) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *CheckTransaction) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetStatus

`func (o *CheckTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSystemDescription

`func (o *CheckTransaction) GetSystemDescription() string`

GetSystemDescription returns the SystemDescription field if non-nil, zero value otherwise.

### GetSystemDescriptionOk

`func (o *CheckTransaction) GetSystemDescriptionOk() (*string, bool)`

GetSystemDescriptionOk returns a tuple with the SystemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDescription

`func (o *CheckTransaction) SetSystemDescription(v string)`

SetSystemDescription sets SystemDescription field to given value.

### HasSystemDescription

`func (o *CheckTransaction) HasSystemDescription() bool`

HasSystemDescription returns a boolean if a field has been set.

### GetTransactionTime

`func (o *CheckTransaction) GetTransactionTime() time.Time`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *CheckTransaction) GetTransactionTimeOk() (*time.Time, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *CheckTransaction) SetTransactionTime(v time.Time)`

SetTransactionTime sets TransactionTime field to given value.


### GetType

`func (o *CheckTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetCheckTransaction

`func (o *CheckTransaction) GetCheckTransaction() CheckTransactionData`

GetCheckTransaction returns the CheckTransaction field if non-nil, zero value otherwise.

### GetCheckTransactionOk

`func (o *CheckTransaction) GetCheckTransactionOk() (*CheckTransactionData, bool)`

GetCheckTransactionOk returns a tuple with the CheckTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTransaction

`func (o *CheckTransaction) SetCheckTransaction(v CheckTransactionData)`

SetCheckTransaction sets CheckTransaction field to given value.


### GetSubtype

`func (o *CheckTransaction) GetSubtype() CheckTransactionSubtypes`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *CheckTransaction) GetSubtypeOk() (*CheckTransactionSubtypes, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *CheckTransaction) SetSubtype(v CheckTransactionSubtypes)`

SetSubtype sets Subtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


