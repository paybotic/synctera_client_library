# ChargeUnsecuredStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The unique identifier of the account the statement belongs to | [optional] [readonly] 
**EndDate** | Pointer to **string** | The date indicating the ending of the time interval covered by the statement | [optional] [readonly] 
**Id** | Pointer to **string** | statement ID | [optional] [readonly] 
**IssueDate** | Pointer to **string** | The date when the statement has been issued | [optional] [readonly] 
**StartDate** | Pointer to **string** | The date indicating the beginning of the time interval covered by the statement | [optional] [readonly] 
**AccountSummary** | Pointer to [**AccountSummary**](AccountSummary.md) |  | [optional] 
**AuthorizedSigners** | Pointer to [**[]Person**](Person.md) |  | [optional] [readonly] 
**ClosingBalance** | Pointer to **int64** | The account balance at the end of the statement period, in ISO 4217 minor currency units. | [optional] 
**CustomerServiceDetails** | Pointer to [**CustomerServiceDetails**](CustomerServiceDetails.md) |  | [optional] 
**Disclosure** | Pointer to **string** |  | [optional] 
**ExcludeJitTransactions** | **bool** | Ignore \&quot;JIT funding\&quot; transactions when generating a statement | [default to false]
**IncludeChildTransactions** | **bool** | Include transactions from sub-accounts when generating a statement | [default to false]
**JointAccountHolders** | Pointer to [**[]Person**](Person.md) |  | [optional] [readonly] 
**OpeningBalance** | Pointer to **int64** | The account balance at the start of the statement period, in ISO 4217 minor currency units. | [optional] 
**PrimaryAccountHolderBusiness** | Pointer to [**Business**](Business.md) |  | [optional] 
**PrimaryAccountHolderPersonal** | Pointer to [**Person**](Person.md) |  | [optional] 
**StatementType** | [**StatementType**](StatementType.md) |  | 
**TotalTransactions** | Pointer to **int64** | The total number of transactions for this statement period.  | [optional] 
**CreditSummary** | [**CreditSummary**](CreditSummary.md) |  | 

## Methods

### NewChargeUnsecuredStatement

`func NewChargeUnsecuredStatement(excludeJitTransactions bool, includeChildTransactions bool, statementType StatementType, creditSummary CreditSummary, ) *ChargeUnsecuredStatement`

NewChargeUnsecuredStatement instantiates a new ChargeUnsecuredStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeUnsecuredStatementWithDefaults

`func NewChargeUnsecuredStatementWithDefaults() *ChargeUnsecuredStatement`

NewChargeUnsecuredStatementWithDefaults instantiates a new ChargeUnsecuredStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ChargeUnsecuredStatement) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ChargeUnsecuredStatement) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ChargeUnsecuredStatement) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ChargeUnsecuredStatement) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEndDate

`func (o *ChargeUnsecuredStatement) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ChargeUnsecuredStatement) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ChargeUnsecuredStatement) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ChargeUnsecuredStatement) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *ChargeUnsecuredStatement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeUnsecuredStatement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeUnsecuredStatement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChargeUnsecuredStatement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *ChargeUnsecuredStatement) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ChargeUnsecuredStatement) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ChargeUnsecuredStatement) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *ChargeUnsecuredStatement) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *ChargeUnsecuredStatement) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ChargeUnsecuredStatement) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ChargeUnsecuredStatement) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ChargeUnsecuredStatement) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetAccountSummary

`func (o *ChargeUnsecuredStatement) GetAccountSummary() AccountSummary`

GetAccountSummary returns the AccountSummary field if non-nil, zero value otherwise.

### GetAccountSummaryOk

`func (o *ChargeUnsecuredStatement) GetAccountSummaryOk() (*AccountSummary, bool)`

GetAccountSummaryOk returns a tuple with the AccountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSummary

`func (o *ChargeUnsecuredStatement) SetAccountSummary(v AccountSummary)`

SetAccountSummary sets AccountSummary field to given value.

### HasAccountSummary

`func (o *ChargeUnsecuredStatement) HasAccountSummary() bool`

HasAccountSummary returns a boolean if a field has been set.

### GetAuthorizedSigners

`func (o *ChargeUnsecuredStatement) GetAuthorizedSigners() []Person`

GetAuthorizedSigners returns the AuthorizedSigners field if non-nil, zero value otherwise.

### GetAuthorizedSignersOk

`func (o *ChargeUnsecuredStatement) GetAuthorizedSignersOk() (*[]Person, bool)`

GetAuthorizedSignersOk returns a tuple with the AuthorizedSigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedSigners

`func (o *ChargeUnsecuredStatement) SetAuthorizedSigners(v []Person)`

SetAuthorizedSigners sets AuthorizedSigners field to given value.

### HasAuthorizedSigners

`func (o *ChargeUnsecuredStatement) HasAuthorizedSigners() bool`

HasAuthorizedSigners returns a boolean if a field has been set.

### GetClosingBalance

`func (o *ChargeUnsecuredStatement) GetClosingBalance() int64`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *ChargeUnsecuredStatement) GetClosingBalanceOk() (*int64, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *ChargeUnsecuredStatement) SetClosingBalance(v int64)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *ChargeUnsecuredStatement) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetCustomerServiceDetails

`func (o *ChargeUnsecuredStatement) GetCustomerServiceDetails() CustomerServiceDetails`

GetCustomerServiceDetails returns the CustomerServiceDetails field if non-nil, zero value otherwise.

### GetCustomerServiceDetailsOk

`func (o *ChargeUnsecuredStatement) GetCustomerServiceDetailsOk() (*CustomerServiceDetails, bool)`

GetCustomerServiceDetailsOk returns a tuple with the CustomerServiceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerServiceDetails

`func (o *ChargeUnsecuredStatement) SetCustomerServiceDetails(v CustomerServiceDetails)`

SetCustomerServiceDetails sets CustomerServiceDetails field to given value.

### HasCustomerServiceDetails

`func (o *ChargeUnsecuredStatement) HasCustomerServiceDetails() bool`

HasCustomerServiceDetails returns a boolean if a field has been set.

### GetDisclosure

`func (o *ChargeUnsecuredStatement) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *ChargeUnsecuredStatement) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *ChargeUnsecuredStatement) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *ChargeUnsecuredStatement) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetExcludeJitTransactions

`func (o *ChargeUnsecuredStatement) GetExcludeJitTransactions() bool`

GetExcludeJitTransactions returns the ExcludeJitTransactions field if non-nil, zero value otherwise.

### GetExcludeJitTransactionsOk

`func (o *ChargeUnsecuredStatement) GetExcludeJitTransactionsOk() (*bool, bool)`

GetExcludeJitTransactionsOk returns a tuple with the ExcludeJitTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeJitTransactions

`func (o *ChargeUnsecuredStatement) SetExcludeJitTransactions(v bool)`

SetExcludeJitTransactions sets ExcludeJitTransactions field to given value.


### GetIncludeChildTransactions

`func (o *ChargeUnsecuredStatement) GetIncludeChildTransactions() bool`

GetIncludeChildTransactions returns the IncludeChildTransactions field if non-nil, zero value otherwise.

### GetIncludeChildTransactionsOk

`func (o *ChargeUnsecuredStatement) GetIncludeChildTransactionsOk() (*bool, bool)`

GetIncludeChildTransactionsOk returns a tuple with the IncludeChildTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChildTransactions

`func (o *ChargeUnsecuredStatement) SetIncludeChildTransactions(v bool)`

SetIncludeChildTransactions sets IncludeChildTransactions field to given value.


### GetJointAccountHolders

`func (o *ChargeUnsecuredStatement) GetJointAccountHolders() []Person`

GetJointAccountHolders returns the JointAccountHolders field if non-nil, zero value otherwise.

### GetJointAccountHoldersOk

`func (o *ChargeUnsecuredStatement) GetJointAccountHoldersOk() (*[]Person, bool)`

GetJointAccountHoldersOk returns a tuple with the JointAccountHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJointAccountHolders

`func (o *ChargeUnsecuredStatement) SetJointAccountHolders(v []Person)`

SetJointAccountHolders sets JointAccountHolders field to given value.

### HasJointAccountHolders

`func (o *ChargeUnsecuredStatement) HasJointAccountHolders() bool`

HasJointAccountHolders returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *ChargeUnsecuredStatement) GetOpeningBalance() int64`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *ChargeUnsecuredStatement) GetOpeningBalanceOk() (*int64, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *ChargeUnsecuredStatement) SetOpeningBalance(v int64)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *ChargeUnsecuredStatement) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetPrimaryAccountHolderBusiness

`func (o *ChargeUnsecuredStatement) GetPrimaryAccountHolderBusiness() Business`

GetPrimaryAccountHolderBusiness returns the PrimaryAccountHolderBusiness field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderBusinessOk

`func (o *ChargeUnsecuredStatement) GetPrimaryAccountHolderBusinessOk() (*Business, bool)`

GetPrimaryAccountHolderBusinessOk returns a tuple with the PrimaryAccountHolderBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderBusiness

`func (o *ChargeUnsecuredStatement) SetPrimaryAccountHolderBusiness(v Business)`

SetPrimaryAccountHolderBusiness sets PrimaryAccountHolderBusiness field to given value.

### HasPrimaryAccountHolderBusiness

`func (o *ChargeUnsecuredStatement) HasPrimaryAccountHolderBusiness() bool`

HasPrimaryAccountHolderBusiness returns a boolean if a field has been set.

### GetPrimaryAccountHolderPersonal

`func (o *ChargeUnsecuredStatement) GetPrimaryAccountHolderPersonal() Person`

GetPrimaryAccountHolderPersonal returns the PrimaryAccountHolderPersonal field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderPersonalOk

`func (o *ChargeUnsecuredStatement) GetPrimaryAccountHolderPersonalOk() (*Person, bool)`

GetPrimaryAccountHolderPersonalOk returns a tuple with the PrimaryAccountHolderPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderPersonal

`func (o *ChargeUnsecuredStatement) SetPrimaryAccountHolderPersonal(v Person)`

SetPrimaryAccountHolderPersonal sets PrimaryAccountHolderPersonal field to given value.

### HasPrimaryAccountHolderPersonal

`func (o *ChargeUnsecuredStatement) HasPrimaryAccountHolderPersonal() bool`

HasPrimaryAccountHolderPersonal returns a boolean if a field has been set.

### GetStatementType

`func (o *ChargeUnsecuredStatement) GetStatementType() StatementType`

GetStatementType returns the StatementType field if non-nil, zero value otherwise.

### GetStatementTypeOk

`func (o *ChargeUnsecuredStatement) GetStatementTypeOk() (*StatementType, bool)`

GetStatementTypeOk returns a tuple with the StatementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementType

`func (o *ChargeUnsecuredStatement) SetStatementType(v StatementType)`

SetStatementType sets StatementType field to given value.


### GetTotalTransactions

`func (o *ChargeUnsecuredStatement) GetTotalTransactions() int64`

GetTotalTransactions returns the TotalTransactions field if non-nil, zero value otherwise.

### GetTotalTransactionsOk

`func (o *ChargeUnsecuredStatement) GetTotalTransactionsOk() (*int64, bool)`

GetTotalTransactionsOk returns a tuple with the TotalTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactions

`func (o *ChargeUnsecuredStatement) SetTotalTransactions(v int64)`

SetTotalTransactions sets TotalTransactions field to given value.

### HasTotalTransactions

`func (o *ChargeUnsecuredStatement) HasTotalTransactions() bool`

HasTotalTransactions returns a boolean if a field has been set.

### GetCreditSummary

`func (o *ChargeUnsecuredStatement) GetCreditSummary() CreditSummary`

GetCreditSummary returns the CreditSummary field if non-nil, zero value otherwise.

### GetCreditSummaryOk

`func (o *ChargeUnsecuredStatement) GetCreditSummaryOk() (*CreditSummary, bool)`

GetCreditSummaryOk returns a tuple with the CreditSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditSummary

`func (o *ChargeUnsecuredStatement) SetCreditSummary(v CreditSummary)`

SetCreditSummary sets CreditSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


