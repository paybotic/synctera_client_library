# ExternalCardDisputeResponse

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
**ActionHistory** | [**[]ExternalCardDisputeActionResponse**](ExternalCardDisputeActionResponse.md) | History of all actions taken on the dispute | 
**AvailableActions** | [**[]AvailableExternalCardAction**](AvailableExternalCardAction.md) | List of actions that can be taken on the dispute | 
**DebitStatus** | [**DisputeDebitStatus**](DisputeDebitStatus.md) |  | 
**LifecycleState** | [**ExternalCardLifecycleState**](ExternalCardLifecycleState.md) |  | 
**ManagedBy** | [**CardDisputeManagedBy**](CardDisputeManagedBy.md) |  | 
**NetworkEligibility** | [**ExternalCardNetworkEligibility**](ExternalCardNetworkEligibility.md) |  | 
**ReasonCode** | [**ExternalCardDisputeReasonCode**](ExternalCardDisputeReasonCode.md) |  | 

## Methods

### NewExternalCardDisputeResponse

`func NewExternalCardDisputeResponse(accountId string, creationTime time.Time, currency string, customerId string, dateCustomerReported time.Time, decision DisputeDecision, disputeDocuments []DisputeDocumentResponse, disputedAmount int32, id string, lastUpdatedTime time.Time, memo string, network DisputeNetwork, paymentRail PaymentRail, status DisputeStatus, tenant string, transactionId string, actionHistory []ExternalCardDisputeActionResponse, availableActions []AvailableExternalCardAction, debitStatus DisputeDebitStatus, lifecycleState ExternalCardLifecycleState, managedBy CardDisputeManagedBy, networkEligibility ExternalCardNetworkEligibility, reasonCode ExternalCardDisputeReasonCode, ) *ExternalCardDisputeResponse`

NewExternalCardDisputeResponse instantiates a new ExternalCardDisputeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardDisputeResponseWithDefaults

`func NewExternalCardDisputeResponseWithDefaults() *ExternalCardDisputeResponse`

NewExternalCardDisputeResponseWithDefaults instantiates a new ExternalCardDisputeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ExternalCardDisputeResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ExternalCardDisputeResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ExternalCardDisputeResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetApplicableRegulation

`func (o *ExternalCardDisputeResponse) GetApplicableRegulation() DisputeRegulation`

GetApplicableRegulation returns the ApplicableRegulation field if non-nil, zero value otherwise.

### GetApplicableRegulationOk

`func (o *ExternalCardDisputeResponse) GetApplicableRegulationOk() (*DisputeRegulation, bool)`

GetApplicableRegulationOk returns a tuple with the ApplicableRegulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableRegulation

`func (o *ExternalCardDisputeResponse) SetApplicableRegulation(v DisputeRegulation)`

SetApplicableRegulation sets ApplicableRegulation field to given value.

### HasApplicableRegulation

`func (o *ExternalCardDisputeResponse) HasApplicableRegulation() bool`

HasApplicableRegulation returns a boolean if a field has been set.

### GetCreationTime

`func (o *ExternalCardDisputeResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ExternalCardDisputeResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ExternalCardDisputeResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *ExternalCardDisputeResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExternalCardDisputeResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExternalCardDisputeResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *ExternalCardDisputeResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExternalCardDisputeResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExternalCardDisputeResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDateCustomerReported

`func (o *ExternalCardDisputeResponse) GetDateCustomerReported() time.Time`

GetDateCustomerReported returns the DateCustomerReported field if non-nil, zero value otherwise.

### GetDateCustomerReportedOk

`func (o *ExternalCardDisputeResponse) GetDateCustomerReportedOk() (*time.Time, bool)`

GetDateCustomerReportedOk returns a tuple with the DateCustomerReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCustomerReported

`func (o *ExternalCardDisputeResponse) SetDateCustomerReported(v time.Time)`

SetDateCustomerReported sets DateCustomerReported field to given value.


### GetDecision

`func (o *ExternalCardDisputeResponse) GetDecision() DisputeDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ExternalCardDisputeResponse) GetDecisionOk() (*DisputeDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ExternalCardDisputeResponse) SetDecision(v DisputeDecision)`

SetDecision sets Decision field to given value.


### GetDisputeDocuments

`func (o *ExternalCardDisputeResponse) GetDisputeDocuments() []DisputeDocumentResponse`

GetDisputeDocuments returns the DisputeDocuments field if non-nil, zero value otherwise.

### GetDisputeDocumentsOk

`func (o *ExternalCardDisputeResponse) GetDisputeDocumentsOk() (*[]DisputeDocumentResponse, bool)`

GetDisputeDocumentsOk returns a tuple with the DisputeDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeDocuments

`func (o *ExternalCardDisputeResponse) SetDisputeDocuments(v []DisputeDocumentResponse)`

SetDisputeDocuments sets DisputeDocuments field to given value.


### GetDisputedAmount

`func (o *ExternalCardDisputeResponse) GetDisputedAmount() int32`

GetDisputedAmount returns the DisputedAmount field if non-nil, zero value otherwise.

### GetDisputedAmountOk

`func (o *ExternalCardDisputeResponse) GetDisputedAmountOk() (*int32, bool)`

GetDisputedAmountOk returns a tuple with the DisputedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmount

`func (o *ExternalCardDisputeResponse) SetDisputedAmount(v int32)`

SetDisputedAmount sets DisputedAmount field to given value.


### GetExternalReferenceId

`func (o *ExternalCardDisputeResponse) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *ExternalCardDisputeResponse) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *ExternalCardDisputeResponse) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.

### HasExternalReferenceId

`func (o *ExternalCardDisputeResponse) HasExternalReferenceId() bool`

HasExternalReferenceId returns a boolean if a field has been set.

### GetId

`func (o *ExternalCardDisputeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalCardDisputeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalCardDisputeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastActionBy

`func (o *ExternalCardDisputeResponse) GetLastActionBy() DisputeActionBy`

GetLastActionBy returns the LastActionBy field if non-nil, zero value otherwise.

### GetLastActionByOk

`func (o *ExternalCardDisputeResponse) GetLastActionByOk() (*DisputeActionBy, bool)`

GetLastActionByOk returns a tuple with the LastActionBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActionBy

`func (o *ExternalCardDisputeResponse) SetLastActionBy(v DisputeActionBy)`

SetLastActionBy sets LastActionBy field to given value.

### HasLastActionBy

`func (o *ExternalCardDisputeResponse) HasLastActionBy() bool`

HasLastActionBy returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *ExternalCardDisputeResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ExternalCardDisputeResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ExternalCardDisputeResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMemo

`func (o *ExternalCardDisputeResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ExternalCardDisputeResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ExternalCardDisputeResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetNetwork

`func (o *ExternalCardDisputeResponse) GetNetwork() DisputeNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ExternalCardDisputeResponse) GetNetworkOk() (*DisputeNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ExternalCardDisputeResponse) SetNetwork(v DisputeNetwork)`

SetNetwork sets Network field to given value.


### GetPaymentRail

`func (o *ExternalCardDisputeResponse) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *ExternalCardDisputeResponse) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *ExternalCardDisputeResponse) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetStatus

`func (o *ExternalCardDisputeResponse) GetStatus() DisputeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalCardDisputeResponse) GetStatusOk() (*DisputeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalCardDisputeResponse) SetStatus(v DisputeStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *ExternalCardDisputeResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ExternalCardDisputeResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ExternalCardDisputeResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTimestampFinalDecision

`func (o *ExternalCardDisputeResponse) GetTimestampFinalDecision() time.Time`

GetTimestampFinalDecision returns the TimestampFinalDecision field if non-nil, zero value otherwise.

### GetTimestampFinalDecisionOk

`func (o *ExternalCardDisputeResponse) GetTimestampFinalDecisionOk() (*time.Time, bool)`

GetTimestampFinalDecisionOk returns a tuple with the TimestampFinalDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFinalDecision

`func (o *ExternalCardDisputeResponse) SetTimestampFinalDecision(v time.Time)`

SetTimestampFinalDecision sets TimestampFinalDecision field to given value.

### HasTimestampFinalDecision

`func (o *ExternalCardDisputeResponse) HasTimestampFinalDecision() bool`

HasTimestampFinalDecision returns a boolean if a field has been set.

### GetTimestampInvestigationDue

`func (o *ExternalCardDisputeResponse) GetTimestampInvestigationDue() time.Time`

GetTimestampInvestigationDue returns the TimestampInvestigationDue field if non-nil, zero value otherwise.

### GetTimestampInvestigationDueOk

`func (o *ExternalCardDisputeResponse) GetTimestampInvestigationDueOk() (*time.Time, bool)`

GetTimestampInvestigationDueOk returns a tuple with the TimestampInvestigationDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampInvestigationDue

`func (o *ExternalCardDisputeResponse) SetTimestampInvestigationDue(v time.Time)`

SetTimestampInvestigationDue sets TimestampInvestigationDue field to given value.

### HasTimestampInvestigationDue

`func (o *ExternalCardDisputeResponse) HasTimestampInvestigationDue() bool`

HasTimestampInvestigationDue returns a boolean if a field has been set.

### GetTimestampProvisionalCreditDue

`func (o *ExternalCardDisputeResponse) GetTimestampProvisionalCreditDue() time.Time`

GetTimestampProvisionalCreditDue returns the TimestampProvisionalCreditDue field if non-nil, zero value otherwise.

### GetTimestampProvisionalCreditDueOk

`func (o *ExternalCardDisputeResponse) GetTimestampProvisionalCreditDueOk() (*time.Time, bool)`

GetTimestampProvisionalCreditDueOk returns a tuple with the TimestampProvisionalCreditDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampProvisionalCreditDue

`func (o *ExternalCardDisputeResponse) SetTimestampProvisionalCreditDue(v time.Time)`

SetTimestampProvisionalCreditDue sets TimestampProvisionalCreditDue field to given value.

### HasTimestampProvisionalCreditDue

`func (o *ExternalCardDisputeResponse) HasTimestampProvisionalCreditDue() bool`

HasTimestampProvisionalCreditDue returns a boolean if a field has been set.

### GetTransactionId

`func (o *ExternalCardDisputeResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ExternalCardDisputeResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ExternalCardDisputeResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetActionHistory

`func (o *ExternalCardDisputeResponse) GetActionHistory() []ExternalCardDisputeActionResponse`

GetActionHistory returns the ActionHistory field if non-nil, zero value otherwise.

### GetActionHistoryOk

`func (o *ExternalCardDisputeResponse) GetActionHistoryOk() (*[]ExternalCardDisputeActionResponse, bool)`

GetActionHistoryOk returns a tuple with the ActionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionHistory

`func (o *ExternalCardDisputeResponse) SetActionHistory(v []ExternalCardDisputeActionResponse)`

SetActionHistory sets ActionHistory field to given value.


### GetAvailableActions

`func (o *ExternalCardDisputeResponse) GetAvailableActions() []AvailableExternalCardAction`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *ExternalCardDisputeResponse) GetAvailableActionsOk() (*[]AvailableExternalCardAction, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *ExternalCardDisputeResponse) SetAvailableActions(v []AvailableExternalCardAction)`

SetAvailableActions sets AvailableActions field to given value.


### GetDebitStatus

`func (o *ExternalCardDisputeResponse) GetDebitStatus() DisputeDebitStatus`

GetDebitStatus returns the DebitStatus field if non-nil, zero value otherwise.

### GetDebitStatusOk

`func (o *ExternalCardDisputeResponse) GetDebitStatusOk() (*DisputeDebitStatus, bool)`

GetDebitStatusOk returns a tuple with the DebitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitStatus

`func (o *ExternalCardDisputeResponse) SetDebitStatus(v DisputeDebitStatus)`

SetDebitStatus sets DebitStatus field to given value.


### GetLifecycleState

`func (o *ExternalCardDisputeResponse) GetLifecycleState() ExternalCardLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *ExternalCardDisputeResponse) GetLifecycleStateOk() (*ExternalCardLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *ExternalCardDisputeResponse) SetLifecycleState(v ExternalCardLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.


### GetManagedBy

`func (o *ExternalCardDisputeResponse) GetManagedBy() CardDisputeManagedBy`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *ExternalCardDisputeResponse) GetManagedByOk() (*CardDisputeManagedBy, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *ExternalCardDisputeResponse) SetManagedBy(v CardDisputeManagedBy)`

SetManagedBy sets ManagedBy field to given value.


### GetNetworkEligibility

`func (o *ExternalCardDisputeResponse) GetNetworkEligibility() ExternalCardNetworkEligibility`

GetNetworkEligibility returns the NetworkEligibility field if non-nil, zero value otherwise.

### GetNetworkEligibilityOk

`func (o *ExternalCardDisputeResponse) GetNetworkEligibilityOk() (*ExternalCardNetworkEligibility, bool)`

GetNetworkEligibilityOk returns a tuple with the NetworkEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEligibility

`func (o *ExternalCardDisputeResponse) SetNetworkEligibility(v ExternalCardNetworkEligibility)`

SetNetworkEligibility sets NetworkEligibility field to given value.


### GetReasonCode

`func (o *ExternalCardDisputeResponse) GetReasonCode() ExternalCardDisputeReasonCode`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ExternalCardDisputeResponse) GetReasonCodeOk() (*ExternalCardDisputeReasonCode, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ExternalCardDisputeResponse) SetReasonCode(v ExternalCardDisputeReasonCode)`

SetReasonCode sets ReasonCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


