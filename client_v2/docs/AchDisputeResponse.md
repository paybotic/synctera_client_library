# AchDisputeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The customer account related to the dispute, to which dispute-related credits will be posted. | 
**ApplicableRegulation** | Pointer to [**DisputeRegulation**](DisputeRegulation.md) |  | [optional] 
**CreationTime** | **time.Time** | The timestamp representing when the dispute was created | [readonly] 
**Currency** | **string** | ISO 4217 Alpha-3 currency code | 
**CustomerId** | **string** | The customer related to the dispute, to which dispute-related credits will be posted. | 
**DateCustomerReported** | **time.Time** | The timestamp representing when the customer reported the dispute. | 
**Decision** | [**DisputeDecision**](DisputeDecision.md) |  | 
**DisputeDocuments** | [**[]DisputeDocumentResponse**](DisputeDocumentResponse.md) | Documents associated with the dispute. | 
**DisputedAmount** | **int32** | The amount to be disputed in cents. | 
**ExternalReferenceId** | Pointer to **string** | Reference ID associated with the dispute on the external network. | [optional] 
**Id** | **string** | The unique identifier of the dispute | [readonly] 
**LastActionBy** | Pointer to [**DisputeActionBy**](DisputeActionBy.md) |  | [optional] 
**LastUpdatedTime** | **time.Time** | The timestamp representing when the dispute was last modified | [readonly] 
**Memo** | **string** | Brief written message describing the reason for disputing the transaction. | 
**Network** | [**DisputeNetwork**](DisputeNetwork.md) |  | 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**Status** | [**DisputeStatus**](DisputeStatus.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**TimestampFinalDecision** | Pointer to **time.Time** | The time that a final decision was made on the dispute. | [optional] 
**TimestampInvestigationDue** | Pointer to **time.Time** |  | [optional] 
**TimestampProvisionalCreditDue** | Pointer to **time.Time** | The time by which provisional credit should be posted to the customer account in response to the dispute. | [optional] 
**TransactionId** | **string** | The ID of the posted transaction to be disputed. | 
**ActionHistory** | [**[]AchDisputeActionResponse**](AchDisputeActionResponse.md) | History of all the action filed for the dispute. | 
**AvailableActions** | [**[]AvailableAchAction**](AvailableAchAction.md) | List of actions that can be taken on the dispute. | 
**CreditStatus** | [**DisputeCreditStatus**](DisputeCreditStatus.md) |  | 
**LifecycleState** | [**AchLifecycleState**](AchLifecycleState.md) |  | 

## Methods

### NewAchDisputeResponse

`func NewAchDisputeResponse(accountId string, creationTime time.Time, currency string, customerId string, dateCustomerReported time.Time, decision DisputeDecision, disputeDocuments []DisputeDocumentResponse, disputedAmount int32, id string, lastUpdatedTime time.Time, memo string, network DisputeNetwork, paymentRail PaymentRail, status DisputeStatus, tenant string, transactionId string, actionHistory []AchDisputeActionResponse, availableActions []AvailableAchAction, creditStatus DisputeCreditStatus, lifecycleState AchLifecycleState, ) *AchDisputeResponse`

NewAchDisputeResponse instantiates a new AchDisputeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchDisputeResponseWithDefaults

`func NewAchDisputeResponseWithDefaults() *AchDisputeResponse`

NewAchDisputeResponseWithDefaults instantiates a new AchDisputeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AchDisputeResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AchDisputeResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AchDisputeResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetApplicableRegulation

`func (o *AchDisputeResponse) GetApplicableRegulation() DisputeRegulation`

GetApplicableRegulation returns the ApplicableRegulation field if non-nil, zero value otherwise.

### GetApplicableRegulationOk

`func (o *AchDisputeResponse) GetApplicableRegulationOk() (*DisputeRegulation, bool)`

GetApplicableRegulationOk returns a tuple with the ApplicableRegulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableRegulation

`func (o *AchDisputeResponse) SetApplicableRegulation(v DisputeRegulation)`

SetApplicableRegulation sets ApplicableRegulation field to given value.

### HasApplicableRegulation

`func (o *AchDisputeResponse) HasApplicableRegulation() bool`

HasApplicableRegulation returns a boolean if a field has been set.

### GetCreationTime

`func (o *AchDisputeResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AchDisputeResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AchDisputeResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *AchDisputeResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AchDisputeResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AchDisputeResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *AchDisputeResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AchDisputeResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AchDisputeResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDateCustomerReported

`func (o *AchDisputeResponse) GetDateCustomerReported() time.Time`

GetDateCustomerReported returns the DateCustomerReported field if non-nil, zero value otherwise.

### GetDateCustomerReportedOk

`func (o *AchDisputeResponse) GetDateCustomerReportedOk() (*time.Time, bool)`

GetDateCustomerReportedOk returns a tuple with the DateCustomerReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCustomerReported

`func (o *AchDisputeResponse) SetDateCustomerReported(v time.Time)`

SetDateCustomerReported sets DateCustomerReported field to given value.


### GetDecision

`func (o *AchDisputeResponse) GetDecision() DisputeDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *AchDisputeResponse) GetDecisionOk() (*DisputeDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *AchDisputeResponse) SetDecision(v DisputeDecision)`

SetDecision sets Decision field to given value.


### GetDisputeDocuments

`func (o *AchDisputeResponse) GetDisputeDocuments() []DisputeDocumentResponse`

GetDisputeDocuments returns the DisputeDocuments field if non-nil, zero value otherwise.

### GetDisputeDocumentsOk

`func (o *AchDisputeResponse) GetDisputeDocumentsOk() (*[]DisputeDocumentResponse, bool)`

GetDisputeDocumentsOk returns a tuple with the DisputeDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeDocuments

`func (o *AchDisputeResponse) SetDisputeDocuments(v []DisputeDocumentResponse)`

SetDisputeDocuments sets DisputeDocuments field to given value.


### GetDisputedAmount

`func (o *AchDisputeResponse) GetDisputedAmount() int32`

GetDisputedAmount returns the DisputedAmount field if non-nil, zero value otherwise.

### GetDisputedAmountOk

`func (o *AchDisputeResponse) GetDisputedAmountOk() (*int32, bool)`

GetDisputedAmountOk returns a tuple with the DisputedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmount

`func (o *AchDisputeResponse) SetDisputedAmount(v int32)`

SetDisputedAmount sets DisputedAmount field to given value.


### GetExternalReferenceId

`func (o *AchDisputeResponse) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *AchDisputeResponse) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *AchDisputeResponse) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.

### HasExternalReferenceId

`func (o *AchDisputeResponse) HasExternalReferenceId() bool`

HasExternalReferenceId returns a boolean if a field has been set.

### GetId

`func (o *AchDisputeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AchDisputeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AchDisputeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastActionBy

`func (o *AchDisputeResponse) GetLastActionBy() DisputeActionBy`

GetLastActionBy returns the LastActionBy field if non-nil, zero value otherwise.

### GetLastActionByOk

`func (o *AchDisputeResponse) GetLastActionByOk() (*DisputeActionBy, bool)`

GetLastActionByOk returns a tuple with the LastActionBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActionBy

`func (o *AchDisputeResponse) SetLastActionBy(v DisputeActionBy)`

SetLastActionBy sets LastActionBy field to given value.

### HasLastActionBy

`func (o *AchDisputeResponse) HasLastActionBy() bool`

HasLastActionBy returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AchDisputeResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AchDisputeResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AchDisputeResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMemo

`func (o *AchDisputeResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *AchDisputeResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *AchDisputeResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetNetwork

`func (o *AchDisputeResponse) GetNetwork() DisputeNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AchDisputeResponse) GetNetworkOk() (*DisputeNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AchDisputeResponse) SetNetwork(v DisputeNetwork)`

SetNetwork sets Network field to given value.


### GetPaymentRail

`func (o *AchDisputeResponse) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *AchDisputeResponse) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *AchDisputeResponse) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetStatus

`func (o *AchDisputeResponse) GetStatus() DisputeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchDisputeResponse) GetStatusOk() (*DisputeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AchDisputeResponse) SetStatus(v DisputeStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *AchDisputeResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AchDisputeResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AchDisputeResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTimestampFinalDecision

`func (o *AchDisputeResponse) GetTimestampFinalDecision() time.Time`

GetTimestampFinalDecision returns the TimestampFinalDecision field if non-nil, zero value otherwise.

### GetTimestampFinalDecisionOk

`func (o *AchDisputeResponse) GetTimestampFinalDecisionOk() (*time.Time, bool)`

GetTimestampFinalDecisionOk returns a tuple with the TimestampFinalDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFinalDecision

`func (o *AchDisputeResponse) SetTimestampFinalDecision(v time.Time)`

SetTimestampFinalDecision sets TimestampFinalDecision field to given value.

### HasTimestampFinalDecision

`func (o *AchDisputeResponse) HasTimestampFinalDecision() bool`

HasTimestampFinalDecision returns a boolean if a field has been set.

### GetTimestampInvestigationDue

`func (o *AchDisputeResponse) GetTimestampInvestigationDue() time.Time`

GetTimestampInvestigationDue returns the TimestampInvestigationDue field if non-nil, zero value otherwise.

### GetTimestampInvestigationDueOk

`func (o *AchDisputeResponse) GetTimestampInvestigationDueOk() (*time.Time, bool)`

GetTimestampInvestigationDueOk returns a tuple with the TimestampInvestigationDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampInvestigationDue

`func (o *AchDisputeResponse) SetTimestampInvestigationDue(v time.Time)`

SetTimestampInvestigationDue sets TimestampInvestigationDue field to given value.

### HasTimestampInvestigationDue

`func (o *AchDisputeResponse) HasTimestampInvestigationDue() bool`

HasTimestampInvestigationDue returns a boolean if a field has been set.

### GetTimestampProvisionalCreditDue

`func (o *AchDisputeResponse) GetTimestampProvisionalCreditDue() time.Time`

GetTimestampProvisionalCreditDue returns the TimestampProvisionalCreditDue field if non-nil, zero value otherwise.

### GetTimestampProvisionalCreditDueOk

`func (o *AchDisputeResponse) GetTimestampProvisionalCreditDueOk() (*time.Time, bool)`

GetTimestampProvisionalCreditDueOk returns a tuple with the TimestampProvisionalCreditDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampProvisionalCreditDue

`func (o *AchDisputeResponse) SetTimestampProvisionalCreditDue(v time.Time)`

SetTimestampProvisionalCreditDue sets TimestampProvisionalCreditDue field to given value.

### HasTimestampProvisionalCreditDue

`func (o *AchDisputeResponse) HasTimestampProvisionalCreditDue() bool`

HasTimestampProvisionalCreditDue returns a boolean if a field has been set.

### GetTransactionId

`func (o *AchDisputeResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AchDisputeResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AchDisputeResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetActionHistory

`func (o *AchDisputeResponse) GetActionHistory() []AchDisputeActionResponse`

GetActionHistory returns the ActionHistory field if non-nil, zero value otherwise.

### GetActionHistoryOk

`func (o *AchDisputeResponse) GetActionHistoryOk() (*[]AchDisputeActionResponse, bool)`

GetActionHistoryOk returns a tuple with the ActionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionHistory

`func (o *AchDisputeResponse) SetActionHistory(v []AchDisputeActionResponse)`

SetActionHistory sets ActionHistory field to given value.


### GetAvailableActions

`func (o *AchDisputeResponse) GetAvailableActions() []AvailableAchAction`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *AchDisputeResponse) GetAvailableActionsOk() (*[]AvailableAchAction, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *AchDisputeResponse) SetAvailableActions(v []AvailableAchAction)`

SetAvailableActions sets AvailableActions field to given value.


### GetCreditStatus

`func (o *AchDisputeResponse) GetCreditStatus() DisputeCreditStatus`

GetCreditStatus returns the CreditStatus field if non-nil, zero value otherwise.

### GetCreditStatusOk

`func (o *AchDisputeResponse) GetCreditStatusOk() (*DisputeCreditStatus, bool)`

GetCreditStatusOk returns a tuple with the CreditStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditStatus

`func (o *AchDisputeResponse) SetCreditStatus(v DisputeCreditStatus)`

SetCreditStatus sets CreditStatus field to given value.


### GetLifecycleState

`func (o *AchDisputeResponse) GetLifecycleState() AchLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *AchDisputeResponse) GetLifecycleStateOk() (*AchLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *AchDisputeResponse) SetLifecycleState(v AchLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


