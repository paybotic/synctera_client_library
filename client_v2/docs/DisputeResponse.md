# DisputeResponse

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
**AcquirerReferenceNumber** | Pointer to **string** | The acquirer reference number of the original transaction associated with the dispute. | [optional] 
**ActionHistory** | [**[]AchDisputeActionResponse**](AchDisputeActionResponse.md) | History of all the action filed for the dispute. | 
**AvailableActions** | [**[]AvailableAchAction**](AvailableAchAction.md) | List of actions that can be taken on the dispute. | 
**CreditStatus** | [**DisputeCreditStatus**](DisputeCreditStatus.md) |  | 
**LifecycleState** | [**AchLifecycleState**](AchLifecycleState.md) |  | 
**ManagedBy** | [**CardDisputeManagedBy**](CardDisputeManagedBy.md) |  | 
**NetworkEligibility** | [**ExternalCardNetworkEligibility**](ExternalCardNetworkEligibility.md) |  | 
**ReasonCode** | [**ExternalCardDisputeReasonCode**](ExternalCardDisputeReasonCode.md) |  | 
**SwitchSerialNumber** | Pointer to **string** | The switch serial number of the original transaction associated with the dispute. | [optional] 
**DebitStatus** | [**DisputeDebitStatus**](DisputeDebitStatus.md) |  | 

## Methods

### NewDisputeResponse

`func NewDisputeResponse(accountId string, creationTime time.Time, currency string, customerId string, dateCustomerReported time.Time, decision DisputeDecision, disputeDocuments []DisputeDocumentResponse, disputedAmount int32, id string, lastUpdatedTime time.Time, memo string, network DisputeNetwork, paymentRail PaymentRail, status DisputeStatus, tenant string, transactionId string, actionHistory []AchDisputeActionResponse, availableActions []AvailableAchAction, creditStatus DisputeCreditStatus, lifecycleState AchLifecycleState, managedBy CardDisputeManagedBy, networkEligibility ExternalCardNetworkEligibility, reasonCode ExternalCardDisputeReasonCode, debitStatus DisputeDebitStatus, ) *DisputeResponse`

NewDisputeResponse instantiates a new DisputeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeResponseWithDefaults

`func NewDisputeResponseWithDefaults() *DisputeResponse`

NewDisputeResponseWithDefaults instantiates a new DisputeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DisputeResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DisputeResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DisputeResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetApplicableRegulation

`func (o *DisputeResponse) GetApplicableRegulation() DisputeRegulation`

GetApplicableRegulation returns the ApplicableRegulation field if non-nil, zero value otherwise.

### GetApplicableRegulationOk

`func (o *DisputeResponse) GetApplicableRegulationOk() (*DisputeRegulation, bool)`

GetApplicableRegulationOk returns a tuple with the ApplicableRegulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableRegulation

`func (o *DisputeResponse) SetApplicableRegulation(v DisputeRegulation)`

SetApplicableRegulation sets ApplicableRegulation field to given value.

### HasApplicableRegulation

`func (o *DisputeResponse) HasApplicableRegulation() bool`

HasApplicableRegulation returns a boolean if a field has been set.

### GetCreationTime

`func (o *DisputeResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DisputeResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DisputeResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *DisputeResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DisputeResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DisputeResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *DisputeResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DisputeResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DisputeResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDateCustomerReported

`func (o *DisputeResponse) GetDateCustomerReported() time.Time`

GetDateCustomerReported returns the DateCustomerReported field if non-nil, zero value otherwise.

### GetDateCustomerReportedOk

`func (o *DisputeResponse) GetDateCustomerReportedOk() (*time.Time, bool)`

GetDateCustomerReportedOk returns a tuple with the DateCustomerReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCustomerReported

`func (o *DisputeResponse) SetDateCustomerReported(v time.Time)`

SetDateCustomerReported sets DateCustomerReported field to given value.


### GetDecision

`func (o *DisputeResponse) GetDecision() DisputeDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *DisputeResponse) GetDecisionOk() (*DisputeDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *DisputeResponse) SetDecision(v DisputeDecision)`

SetDecision sets Decision field to given value.


### GetDisputeDocuments

`func (o *DisputeResponse) GetDisputeDocuments() []DisputeDocumentResponse`

GetDisputeDocuments returns the DisputeDocuments field if non-nil, zero value otherwise.

### GetDisputeDocumentsOk

`func (o *DisputeResponse) GetDisputeDocumentsOk() (*[]DisputeDocumentResponse, bool)`

GetDisputeDocumentsOk returns a tuple with the DisputeDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeDocuments

`func (o *DisputeResponse) SetDisputeDocuments(v []DisputeDocumentResponse)`

SetDisputeDocuments sets DisputeDocuments field to given value.


### GetDisputedAmount

`func (o *DisputeResponse) GetDisputedAmount() int32`

GetDisputedAmount returns the DisputedAmount field if non-nil, zero value otherwise.

### GetDisputedAmountOk

`func (o *DisputeResponse) GetDisputedAmountOk() (*int32, bool)`

GetDisputedAmountOk returns a tuple with the DisputedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmount

`func (o *DisputeResponse) SetDisputedAmount(v int32)`

SetDisputedAmount sets DisputedAmount field to given value.


### GetExternalReferenceId

`func (o *DisputeResponse) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *DisputeResponse) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *DisputeResponse) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.

### HasExternalReferenceId

`func (o *DisputeResponse) HasExternalReferenceId() bool`

HasExternalReferenceId returns a boolean if a field has been set.

### GetId

`func (o *DisputeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisputeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisputeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastActionBy

`func (o *DisputeResponse) GetLastActionBy() DisputeActionBy`

GetLastActionBy returns the LastActionBy field if non-nil, zero value otherwise.

### GetLastActionByOk

`func (o *DisputeResponse) GetLastActionByOk() (*DisputeActionBy, bool)`

GetLastActionByOk returns a tuple with the LastActionBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActionBy

`func (o *DisputeResponse) SetLastActionBy(v DisputeActionBy)`

SetLastActionBy sets LastActionBy field to given value.

### HasLastActionBy

`func (o *DisputeResponse) HasLastActionBy() bool`

HasLastActionBy returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *DisputeResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *DisputeResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *DisputeResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMemo

`func (o *DisputeResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *DisputeResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *DisputeResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetNetwork

`func (o *DisputeResponse) GetNetwork() DisputeNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DisputeResponse) GetNetworkOk() (*DisputeNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DisputeResponse) SetNetwork(v DisputeNetwork)`

SetNetwork sets Network field to given value.


### GetPaymentRail

`func (o *DisputeResponse) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *DisputeResponse) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *DisputeResponse) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetStatus

`func (o *DisputeResponse) GetStatus() DisputeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisputeResponse) GetStatusOk() (*DisputeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisputeResponse) SetStatus(v DisputeStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *DisputeResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DisputeResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DisputeResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTimestampFinalDecision

`func (o *DisputeResponse) GetTimestampFinalDecision() time.Time`

GetTimestampFinalDecision returns the TimestampFinalDecision field if non-nil, zero value otherwise.

### GetTimestampFinalDecisionOk

`func (o *DisputeResponse) GetTimestampFinalDecisionOk() (*time.Time, bool)`

GetTimestampFinalDecisionOk returns a tuple with the TimestampFinalDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFinalDecision

`func (o *DisputeResponse) SetTimestampFinalDecision(v time.Time)`

SetTimestampFinalDecision sets TimestampFinalDecision field to given value.

### HasTimestampFinalDecision

`func (o *DisputeResponse) HasTimestampFinalDecision() bool`

HasTimestampFinalDecision returns a boolean if a field has been set.

### GetTimestampInvestigationDue

`func (o *DisputeResponse) GetTimestampInvestigationDue() time.Time`

GetTimestampInvestigationDue returns the TimestampInvestigationDue field if non-nil, zero value otherwise.

### GetTimestampInvestigationDueOk

`func (o *DisputeResponse) GetTimestampInvestigationDueOk() (*time.Time, bool)`

GetTimestampInvestigationDueOk returns a tuple with the TimestampInvestigationDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampInvestigationDue

`func (o *DisputeResponse) SetTimestampInvestigationDue(v time.Time)`

SetTimestampInvestigationDue sets TimestampInvestigationDue field to given value.

### HasTimestampInvestigationDue

`func (o *DisputeResponse) HasTimestampInvestigationDue() bool`

HasTimestampInvestigationDue returns a boolean if a field has been set.

### GetTimestampProvisionalCreditDue

`func (o *DisputeResponse) GetTimestampProvisionalCreditDue() time.Time`

GetTimestampProvisionalCreditDue returns the TimestampProvisionalCreditDue field if non-nil, zero value otherwise.

### GetTimestampProvisionalCreditDueOk

`func (o *DisputeResponse) GetTimestampProvisionalCreditDueOk() (*time.Time, bool)`

GetTimestampProvisionalCreditDueOk returns a tuple with the TimestampProvisionalCreditDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampProvisionalCreditDue

`func (o *DisputeResponse) SetTimestampProvisionalCreditDue(v time.Time)`

SetTimestampProvisionalCreditDue sets TimestampProvisionalCreditDue field to given value.

### HasTimestampProvisionalCreditDue

`func (o *DisputeResponse) HasTimestampProvisionalCreditDue() bool`

HasTimestampProvisionalCreditDue returns a boolean if a field has been set.

### GetTransactionId

`func (o *DisputeResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DisputeResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DisputeResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAcquirerReferenceNumber

`func (o *DisputeResponse) GetAcquirerReferenceNumber() string`

GetAcquirerReferenceNumber returns the AcquirerReferenceNumber field if non-nil, zero value otherwise.

### GetAcquirerReferenceNumberOk

`func (o *DisputeResponse) GetAcquirerReferenceNumberOk() (*string, bool)`

GetAcquirerReferenceNumberOk returns a tuple with the AcquirerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReferenceNumber

`func (o *DisputeResponse) SetAcquirerReferenceNumber(v string)`

SetAcquirerReferenceNumber sets AcquirerReferenceNumber field to given value.

### HasAcquirerReferenceNumber

`func (o *DisputeResponse) HasAcquirerReferenceNumber() bool`

HasAcquirerReferenceNumber returns a boolean if a field has been set.

### GetActionHistory

`func (o *DisputeResponse) GetActionHistory() []AchDisputeActionResponse`

GetActionHistory returns the ActionHistory field if non-nil, zero value otherwise.

### GetActionHistoryOk

`func (o *DisputeResponse) GetActionHistoryOk() (*[]AchDisputeActionResponse, bool)`

GetActionHistoryOk returns a tuple with the ActionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionHistory

`func (o *DisputeResponse) SetActionHistory(v []AchDisputeActionResponse)`

SetActionHistory sets ActionHistory field to given value.


### GetAvailableActions

`func (o *DisputeResponse) GetAvailableActions() []AvailableAchAction`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *DisputeResponse) GetAvailableActionsOk() (*[]AvailableAchAction, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *DisputeResponse) SetAvailableActions(v []AvailableAchAction)`

SetAvailableActions sets AvailableActions field to given value.


### GetCreditStatus

`func (o *DisputeResponse) GetCreditStatus() DisputeCreditStatus`

GetCreditStatus returns the CreditStatus field if non-nil, zero value otherwise.

### GetCreditStatusOk

`func (o *DisputeResponse) GetCreditStatusOk() (*DisputeCreditStatus, bool)`

GetCreditStatusOk returns a tuple with the CreditStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditStatus

`func (o *DisputeResponse) SetCreditStatus(v DisputeCreditStatus)`

SetCreditStatus sets CreditStatus field to given value.


### GetLifecycleState

`func (o *DisputeResponse) GetLifecycleState() AchLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *DisputeResponse) GetLifecycleStateOk() (*AchLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *DisputeResponse) SetLifecycleState(v AchLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.


### GetManagedBy

`func (o *DisputeResponse) GetManagedBy() CardDisputeManagedBy`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *DisputeResponse) GetManagedByOk() (*CardDisputeManagedBy, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *DisputeResponse) SetManagedBy(v CardDisputeManagedBy)`

SetManagedBy sets ManagedBy field to given value.


### GetNetworkEligibility

`func (o *DisputeResponse) GetNetworkEligibility() ExternalCardNetworkEligibility`

GetNetworkEligibility returns the NetworkEligibility field if non-nil, zero value otherwise.

### GetNetworkEligibilityOk

`func (o *DisputeResponse) GetNetworkEligibilityOk() (*ExternalCardNetworkEligibility, bool)`

GetNetworkEligibilityOk returns a tuple with the NetworkEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEligibility

`func (o *DisputeResponse) SetNetworkEligibility(v ExternalCardNetworkEligibility)`

SetNetworkEligibility sets NetworkEligibility field to given value.


### GetReasonCode

`func (o *DisputeResponse) GetReasonCode() ExternalCardDisputeReasonCode`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DisputeResponse) GetReasonCodeOk() (*ExternalCardDisputeReasonCode, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DisputeResponse) SetReasonCode(v ExternalCardDisputeReasonCode)`

SetReasonCode sets ReasonCode field to given value.


### GetSwitchSerialNumber

`func (o *DisputeResponse) GetSwitchSerialNumber() string`

GetSwitchSerialNumber returns the SwitchSerialNumber field if non-nil, zero value otherwise.

### GetSwitchSerialNumberOk

`func (o *DisputeResponse) GetSwitchSerialNumberOk() (*string, bool)`

GetSwitchSerialNumberOk returns a tuple with the SwitchSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchSerialNumber

`func (o *DisputeResponse) SetSwitchSerialNumber(v string)`

SetSwitchSerialNumber sets SwitchSerialNumber field to given value.

### HasSwitchSerialNumber

`func (o *DisputeResponse) HasSwitchSerialNumber() bool`

HasSwitchSerialNumber returns a boolean if a field has been set.

### GetDebitStatus

`func (o *DisputeResponse) GetDebitStatus() DisputeDebitStatus`

GetDebitStatus returns the DebitStatus field if non-nil, zero value otherwise.

### GetDebitStatusOk

`func (o *DisputeResponse) GetDebitStatusOk() (*DisputeDebitStatus, bool)`

GetDebitStatusOk returns a tuple with the DebitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitStatus

`func (o *DisputeResponse) SetDebitStatus(v DisputeDebitStatus)`

SetDebitStatus sets DebitStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


