# CardDisputeResponse

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
**ActionHistory** | [**[]CardDisputeActionResponse**](CardDisputeActionResponse.md) | History of all the action filed for the dispute. | 
**AvailableActions** | [**[]AvailableCardAction**](AvailableCardAction.md) | List of actions that can be taken on the dispute. | 
**CreditStatus** | [**DisputeCreditStatus**](DisputeCreditStatus.md) |  | 
**LifecycleState** | [**CardLifecycleState**](CardLifecycleState.md) |  | 
**ManagedBy** | [**CardDisputeManagedBy**](CardDisputeManagedBy.md) |  | 
**NetworkEligibility** | [**CardNetworkEligibility**](CardNetworkEligibility.md) |  | 
**ReasonCode** | [**CardDisputeReasonCodes**](CardDisputeReasonCodes.md) |  | 
**SwitchSerialNumber** | Pointer to **string** | The switch serial number of the original transaction associated with the dispute. | [optional] 

## Methods

### NewCardDisputeResponse

`func NewCardDisputeResponse(accountId string, creationTime time.Time, currency string, customerId string, dateCustomerReported time.Time, decision DisputeDecision, disputeDocuments []DisputeDocumentResponse, disputedAmount int32, id string, lastUpdatedTime time.Time, memo string, network DisputeNetwork, paymentRail PaymentRail, status DisputeStatus, tenant string, transactionId string, actionHistory []CardDisputeActionResponse, availableActions []AvailableCardAction, creditStatus DisputeCreditStatus, lifecycleState CardLifecycleState, managedBy CardDisputeManagedBy, networkEligibility CardNetworkEligibility, reasonCode CardDisputeReasonCodes, ) *CardDisputeResponse`

NewCardDisputeResponse instantiates a new CardDisputeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDisputeResponseWithDefaults

`func NewCardDisputeResponseWithDefaults() *CardDisputeResponse`

NewCardDisputeResponseWithDefaults instantiates a new CardDisputeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CardDisputeResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CardDisputeResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CardDisputeResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetApplicableRegulation

`func (o *CardDisputeResponse) GetApplicableRegulation() DisputeRegulation`

GetApplicableRegulation returns the ApplicableRegulation field if non-nil, zero value otherwise.

### GetApplicableRegulationOk

`func (o *CardDisputeResponse) GetApplicableRegulationOk() (*DisputeRegulation, bool)`

GetApplicableRegulationOk returns a tuple with the ApplicableRegulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableRegulation

`func (o *CardDisputeResponse) SetApplicableRegulation(v DisputeRegulation)`

SetApplicableRegulation sets ApplicableRegulation field to given value.

### HasApplicableRegulation

`func (o *CardDisputeResponse) HasApplicableRegulation() bool`

HasApplicableRegulation returns a boolean if a field has been set.

### GetCreationTime

`func (o *CardDisputeResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardDisputeResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardDisputeResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *CardDisputeResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CardDisputeResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CardDisputeResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *CardDisputeResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CardDisputeResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CardDisputeResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDateCustomerReported

`func (o *CardDisputeResponse) GetDateCustomerReported() time.Time`

GetDateCustomerReported returns the DateCustomerReported field if non-nil, zero value otherwise.

### GetDateCustomerReportedOk

`func (o *CardDisputeResponse) GetDateCustomerReportedOk() (*time.Time, bool)`

GetDateCustomerReportedOk returns a tuple with the DateCustomerReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCustomerReported

`func (o *CardDisputeResponse) SetDateCustomerReported(v time.Time)`

SetDateCustomerReported sets DateCustomerReported field to given value.


### GetDecision

`func (o *CardDisputeResponse) GetDecision() DisputeDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *CardDisputeResponse) GetDecisionOk() (*DisputeDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *CardDisputeResponse) SetDecision(v DisputeDecision)`

SetDecision sets Decision field to given value.


### GetDisputeDocuments

`func (o *CardDisputeResponse) GetDisputeDocuments() []DisputeDocumentResponse`

GetDisputeDocuments returns the DisputeDocuments field if non-nil, zero value otherwise.

### GetDisputeDocumentsOk

`func (o *CardDisputeResponse) GetDisputeDocumentsOk() (*[]DisputeDocumentResponse, bool)`

GetDisputeDocumentsOk returns a tuple with the DisputeDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeDocuments

`func (o *CardDisputeResponse) SetDisputeDocuments(v []DisputeDocumentResponse)`

SetDisputeDocuments sets DisputeDocuments field to given value.


### GetDisputedAmount

`func (o *CardDisputeResponse) GetDisputedAmount() int32`

GetDisputedAmount returns the DisputedAmount field if non-nil, zero value otherwise.

### GetDisputedAmountOk

`func (o *CardDisputeResponse) GetDisputedAmountOk() (*int32, bool)`

GetDisputedAmountOk returns a tuple with the DisputedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAmount

`func (o *CardDisputeResponse) SetDisputedAmount(v int32)`

SetDisputedAmount sets DisputedAmount field to given value.


### GetExternalReferenceId

`func (o *CardDisputeResponse) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *CardDisputeResponse) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *CardDisputeResponse) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.

### HasExternalReferenceId

`func (o *CardDisputeResponse) HasExternalReferenceId() bool`

HasExternalReferenceId returns a boolean if a field has been set.

### GetId

`func (o *CardDisputeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardDisputeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardDisputeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastActionBy

`func (o *CardDisputeResponse) GetLastActionBy() DisputeActionBy`

GetLastActionBy returns the LastActionBy field if non-nil, zero value otherwise.

### GetLastActionByOk

`func (o *CardDisputeResponse) GetLastActionByOk() (*DisputeActionBy, bool)`

GetLastActionByOk returns a tuple with the LastActionBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActionBy

`func (o *CardDisputeResponse) SetLastActionBy(v DisputeActionBy)`

SetLastActionBy sets LastActionBy field to given value.

### HasLastActionBy

`func (o *CardDisputeResponse) HasLastActionBy() bool`

HasLastActionBy returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *CardDisputeResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CardDisputeResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CardDisputeResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMemo

`func (o *CardDisputeResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CardDisputeResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CardDisputeResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetNetwork

`func (o *CardDisputeResponse) GetNetwork() DisputeNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CardDisputeResponse) GetNetworkOk() (*DisputeNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CardDisputeResponse) SetNetwork(v DisputeNetwork)`

SetNetwork sets Network field to given value.


### GetPaymentRail

`func (o *CardDisputeResponse) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *CardDisputeResponse) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *CardDisputeResponse) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetStatus

`func (o *CardDisputeResponse) GetStatus() DisputeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardDisputeResponse) GetStatusOk() (*DisputeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardDisputeResponse) SetStatus(v DisputeStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *CardDisputeResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CardDisputeResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CardDisputeResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTimestampFinalDecision

`func (o *CardDisputeResponse) GetTimestampFinalDecision() time.Time`

GetTimestampFinalDecision returns the TimestampFinalDecision field if non-nil, zero value otherwise.

### GetTimestampFinalDecisionOk

`func (o *CardDisputeResponse) GetTimestampFinalDecisionOk() (*time.Time, bool)`

GetTimestampFinalDecisionOk returns a tuple with the TimestampFinalDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFinalDecision

`func (o *CardDisputeResponse) SetTimestampFinalDecision(v time.Time)`

SetTimestampFinalDecision sets TimestampFinalDecision field to given value.

### HasTimestampFinalDecision

`func (o *CardDisputeResponse) HasTimestampFinalDecision() bool`

HasTimestampFinalDecision returns a boolean if a field has been set.

### GetTimestampInvestigationDue

`func (o *CardDisputeResponse) GetTimestampInvestigationDue() time.Time`

GetTimestampInvestigationDue returns the TimestampInvestigationDue field if non-nil, zero value otherwise.

### GetTimestampInvestigationDueOk

`func (o *CardDisputeResponse) GetTimestampInvestigationDueOk() (*time.Time, bool)`

GetTimestampInvestigationDueOk returns a tuple with the TimestampInvestigationDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampInvestigationDue

`func (o *CardDisputeResponse) SetTimestampInvestigationDue(v time.Time)`

SetTimestampInvestigationDue sets TimestampInvestigationDue field to given value.

### HasTimestampInvestigationDue

`func (o *CardDisputeResponse) HasTimestampInvestigationDue() bool`

HasTimestampInvestigationDue returns a boolean if a field has been set.

### GetTimestampProvisionalCreditDue

`func (o *CardDisputeResponse) GetTimestampProvisionalCreditDue() time.Time`

GetTimestampProvisionalCreditDue returns the TimestampProvisionalCreditDue field if non-nil, zero value otherwise.

### GetTimestampProvisionalCreditDueOk

`func (o *CardDisputeResponse) GetTimestampProvisionalCreditDueOk() (*time.Time, bool)`

GetTimestampProvisionalCreditDueOk returns a tuple with the TimestampProvisionalCreditDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampProvisionalCreditDue

`func (o *CardDisputeResponse) SetTimestampProvisionalCreditDue(v time.Time)`

SetTimestampProvisionalCreditDue sets TimestampProvisionalCreditDue field to given value.

### HasTimestampProvisionalCreditDue

`func (o *CardDisputeResponse) HasTimestampProvisionalCreditDue() bool`

HasTimestampProvisionalCreditDue returns a boolean if a field has been set.

### GetTransactionId

`func (o *CardDisputeResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CardDisputeResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CardDisputeResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAcquirerReferenceNumber

`func (o *CardDisputeResponse) GetAcquirerReferenceNumber() string`

GetAcquirerReferenceNumber returns the AcquirerReferenceNumber field if non-nil, zero value otherwise.

### GetAcquirerReferenceNumberOk

`func (o *CardDisputeResponse) GetAcquirerReferenceNumberOk() (*string, bool)`

GetAcquirerReferenceNumberOk returns a tuple with the AcquirerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReferenceNumber

`func (o *CardDisputeResponse) SetAcquirerReferenceNumber(v string)`

SetAcquirerReferenceNumber sets AcquirerReferenceNumber field to given value.

### HasAcquirerReferenceNumber

`func (o *CardDisputeResponse) HasAcquirerReferenceNumber() bool`

HasAcquirerReferenceNumber returns a boolean if a field has been set.

### GetActionHistory

`func (o *CardDisputeResponse) GetActionHistory() []CardDisputeActionResponse`

GetActionHistory returns the ActionHistory field if non-nil, zero value otherwise.

### GetActionHistoryOk

`func (o *CardDisputeResponse) GetActionHistoryOk() (*[]CardDisputeActionResponse, bool)`

GetActionHistoryOk returns a tuple with the ActionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionHistory

`func (o *CardDisputeResponse) SetActionHistory(v []CardDisputeActionResponse)`

SetActionHistory sets ActionHistory field to given value.


### GetAvailableActions

`func (o *CardDisputeResponse) GetAvailableActions() []AvailableCardAction`

GetAvailableActions returns the AvailableActions field if non-nil, zero value otherwise.

### GetAvailableActionsOk

`func (o *CardDisputeResponse) GetAvailableActionsOk() (*[]AvailableCardAction, bool)`

GetAvailableActionsOk returns a tuple with the AvailableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableActions

`func (o *CardDisputeResponse) SetAvailableActions(v []AvailableCardAction)`

SetAvailableActions sets AvailableActions field to given value.


### GetCreditStatus

`func (o *CardDisputeResponse) GetCreditStatus() DisputeCreditStatus`

GetCreditStatus returns the CreditStatus field if non-nil, zero value otherwise.

### GetCreditStatusOk

`func (o *CardDisputeResponse) GetCreditStatusOk() (*DisputeCreditStatus, bool)`

GetCreditStatusOk returns a tuple with the CreditStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditStatus

`func (o *CardDisputeResponse) SetCreditStatus(v DisputeCreditStatus)`

SetCreditStatus sets CreditStatus field to given value.


### GetLifecycleState

`func (o *CardDisputeResponse) GetLifecycleState() CardLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *CardDisputeResponse) GetLifecycleStateOk() (*CardLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *CardDisputeResponse) SetLifecycleState(v CardLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.


### GetManagedBy

`func (o *CardDisputeResponse) GetManagedBy() CardDisputeManagedBy`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *CardDisputeResponse) GetManagedByOk() (*CardDisputeManagedBy, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *CardDisputeResponse) SetManagedBy(v CardDisputeManagedBy)`

SetManagedBy sets ManagedBy field to given value.


### GetNetworkEligibility

`func (o *CardDisputeResponse) GetNetworkEligibility() CardNetworkEligibility`

GetNetworkEligibility returns the NetworkEligibility field if non-nil, zero value otherwise.

### GetNetworkEligibilityOk

`func (o *CardDisputeResponse) GetNetworkEligibilityOk() (*CardNetworkEligibility, bool)`

GetNetworkEligibilityOk returns a tuple with the NetworkEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEligibility

`func (o *CardDisputeResponse) SetNetworkEligibility(v CardNetworkEligibility)`

SetNetworkEligibility sets NetworkEligibility field to given value.


### GetReasonCode

`func (o *CardDisputeResponse) GetReasonCode() CardDisputeReasonCodes`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CardDisputeResponse) GetReasonCodeOk() (*CardDisputeReasonCodes, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CardDisputeResponse) SetReasonCode(v CardDisputeReasonCodes)`

SetReasonCode sets ReasonCode field to given value.


### GetSwitchSerialNumber

`func (o *CardDisputeResponse) GetSwitchSerialNumber() string`

GetSwitchSerialNumber returns the SwitchSerialNumber field if non-nil, zero value otherwise.

### GetSwitchSerialNumberOk

`func (o *CardDisputeResponse) GetSwitchSerialNumberOk() (*string, bool)`

GetSwitchSerialNumberOk returns a tuple with the SwitchSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchSerialNumber

`func (o *CardDisputeResponse) SetSwitchSerialNumber(v string)`

SetSwitchSerialNumber sets SwitchSerialNumber field to given value.

### HasSwitchSerialNumber

`func (o *CardDisputeResponse) HasSwitchSerialNumber() bool`

HasSwitchSerialNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


