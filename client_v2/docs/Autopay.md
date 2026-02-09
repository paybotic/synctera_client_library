# Autopay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptHistory** | Pointer to [**AutopayAttemptHistory**](AutopayAttemptHistory.md) |  | [optional] 
**AutopayConfigId** | **string** | ID of the autopay configuration that created this autopay | 
**BillingPeriodId** | **string** | The billing period ID this autopay is for | 
**ConfigSnapshot** | [**AutopayConfigData**](AutopayConfigData.md) |  | 
**CreationTime** | **time.Time** | Timestamp when the autopay was created | 
**CurrentAmount** | Pointer to **int64** | Current payment amount in cents based on the config snapshot and current balances. This value reflects what would be paid if the autopay executed right now. It may change before actual execution as account balances change.  | [optional] 
**Id** | **string** | Unique identifier for the autopay | 
**LastUpdatedTime** | **time.Time** | Timestamp when the autopay was last updated | 
**LendingAccountId** | **string** | The lending account ID this autopay belongs to | 
**PaymentAttributes** | Pointer to [**AutopayPaymentAttributes**](AutopayPaymentAttributes.md) |  | [optional] 
**RenderedDescription** | **string** | The payment description rendered from the description template at autopay creation time. This is used for ACH addenda and internal transfer memo fields.  | 
**ScheduledDate** | **string** | Date when the autopay is scheduled to execute | 
**StatementId** | Pointer to **string** | The statement ID that triggered this autopay | [optional] 
**Status** | [**AutopayStatus**](AutopayStatus.md) |  | 
**Tenant** | Pointer to [**Tenant**](Tenant.md) |  | [optional] 

## Methods

### NewAutopay

`func NewAutopay(autopayConfigId string, billingPeriodId string, configSnapshot AutopayConfigData, creationTime time.Time, id string, lastUpdatedTime time.Time, lendingAccountId string, renderedDescription string, scheduledDate string, status AutopayStatus, ) *Autopay`

NewAutopay instantiates a new Autopay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayWithDefaults

`func NewAutopayWithDefaults() *Autopay`

NewAutopayWithDefaults instantiates a new Autopay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptHistory

`func (o *Autopay) GetAttemptHistory() AutopayAttemptHistory`

GetAttemptHistory returns the AttemptHistory field if non-nil, zero value otherwise.

### GetAttemptHistoryOk

`func (o *Autopay) GetAttemptHistoryOk() (*AutopayAttemptHistory, bool)`

GetAttemptHistoryOk returns a tuple with the AttemptHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptHistory

`func (o *Autopay) SetAttemptHistory(v AutopayAttemptHistory)`

SetAttemptHistory sets AttemptHistory field to given value.

### HasAttemptHistory

`func (o *Autopay) HasAttemptHistory() bool`

HasAttemptHistory returns a boolean if a field has been set.

### GetAutopayConfigId

`func (o *Autopay) GetAutopayConfigId() string`

GetAutopayConfigId returns the AutopayConfigId field if non-nil, zero value otherwise.

### GetAutopayConfigIdOk

`func (o *Autopay) GetAutopayConfigIdOk() (*string, bool)`

GetAutopayConfigIdOk returns a tuple with the AutopayConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutopayConfigId

`func (o *Autopay) SetAutopayConfigId(v string)`

SetAutopayConfigId sets AutopayConfigId field to given value.


### GetBillingPeriodId

`func (o *Autopay) GetBillingPeriodId() string`

GetBillingPeriodId returns the BillingPeriodId field if non-nil, zero value otherwise.

### GetBillingPeriodIdOk

`func (o *Autopay) GetBillingPeriodIdOk() (*string, bool)`

GetBillingPeriodIdOk returns a tuple with the BillingPeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodId

`func (o *Autopay) SetBillingPeriodId(v string)`

SetBillingPeriodId sets BillingPeriodId field to given value.


### GetConfigSnapshot

`func (o *Autopay) GetConfigSnapshot() AutopayConfigData`

GetConfigSnapshot returns the ConfigSnapshot field if non-nil, zero value otherwise.

### GetConfigSnapshotOk

`func (o *Autopay) GetConfigSnapshotOk() (*AutopayConfigData, bool)`

GetConfigSnapshotOk returns a tuple with the ConfigSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSnapshot

`func (o *Autopay) SetConfigSnapshot(v AutopayConfigData)`

SetConfigSnapshot sets ConfigSnapshot field to given value.


### GetCreationTime

`func (o *Autopay) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Autopay) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Autopay) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrentAmount

`func (o *Autopay) GetCurrentAmount() int64`

GetCurrentAmount returns the CurrentAmount field if non-nil, zero value otherwise.

### GetCurrentAmountOk

`func (o *Autopay) GetCurrentAmountOk() (*int64, bool)`

GetCurrentAmountOk returns a tuple with the CurrentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmount

`func (o *Autopay) SetCurrentAmount(v int64)`

SetCurrentAmount sets CurrentAmount field to given value.

### HasCurrentAmount

`func (o *Autopay) HasCurrentAmount() bool`

HasCurrentAmount returns a boolean if a field has been set.

### GetId

`func (o *Autopay) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Autopay) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Autopay) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *Autopay) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Autopay) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Autopay) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetLendingAccountId

`func (o *Autopay) GetLendingAccountId() string`

GetLendingAccountId returns the LendingAccountId field if non-nil, zero value otherwise.

### GetLendingAccountIdOk

`func (o *Autopay) GetLendingAccountIdOk() (*string, bool)`

GetLendingAccountIdOk returns a tuple with the LendingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLendingAccountId

`func (o *Autopay) SetLendingAccountId(v string)`

SetLendingAccountId sets LendingAccountId field to given value.


### GetPaymentAttributes

`func (o *Autopay) GetPaymentAttributes() AutopayPaymentAttributes`

GetPaymentAttributes returns the PaymentAttributes field if non-nil, zero value otherwise.

### GetPaymentAttributesOk

`func (o *Autopay) GetPaymentAttributesOk() (*AutopayPaymentAttributes, bool)`

GetPaymentAttributesOk returns a tuple with the PaymentAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAttributes

`func (o *Autopay) SetPaymentAttributes(v AutopayPaymentAttributes)`

SetPaymentAttributes sets PaymentAttributes field to given value.

### HasPaymentAttributes

`func (o *Autopay) HasPaymentAttributes() bool`

HasPaymentAttributes returns a boolean if a field has been set.

### GetRenderedDescription

`func (o *Autopay) GetRenderedDescription() string`

GetRenderedDescription returns the RenderedDescription field if non-nil, zero value otherwise.

### GetRenderedDescriptionOk

`func (o *Autopay) GetRenderedDescriptionOk() (*string, bool)`

GetRenderedDescriptionOk returns a tuple with the RenderedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderedDescription

`func (o *Autopay) SetRenderedDescription(v string)`

SetRenderedDescription sets RenderedDescription field to given value.


### GetScheduledDate

`func (o *Autopay) GetScheduledDate() string`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *Autopay) GetScheduledDateOk() (*string, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *Autopay) SetScheduledDate(v string)`

SetScheduledDate sets ScheduledDate field to given value.


### GetStatementId

`func (o *Autopay) GetStatementId() string`

GetStatementId returns the StatementId field if non-nil, zero value otherwise.

### GetStatementIdOk

`func (o *Autopay) GetStatementIdOk() (*string, bool)`

GetStatementIdOk returns a tuple with the StatementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementId

`func (o *Autopay) SetStatementId(v string)`

SetStatementId sets StatementId field to given value.

### HasStatementId

`func (o *Autopay) HasStatementId() bool`

HasStatementId returns a boolean if a field has been set.

### GetStatus

`func (o *Autopay) GetStatus() AutopayStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Autopay) GetStatusOk() (*AutopayStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Autopay) SetStatus(v AutopayStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *Autopay) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Autopay) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Autopay) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Autopay) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


