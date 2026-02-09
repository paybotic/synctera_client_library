# ChargeSecuredStatement

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
**SecurityAccountSummary** | [**SecurityAccountSummary**](SecurityAccountSummary.md) |  | 

## Methods

### NewChargeSecuredStatement

`func NewChargeSecuredStatement(excludeJitTransactions bool, includeChildTransactions bool, statementType StatementType, creditSummary CreditSummary, securityAccountSummary SecurityAccountSummary, ) *ChargeSecuredStatement`

NewChargeSecuredStatement instantiates a new ChargeSecuredStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeSecuredStatementWithDefaults

`func NewChargeSecuredStatementWithDefaults() *ChargeSecuredStatement`

NewChargeSecuredStatementWithDefaults instantiates a new ChargeSecuredStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ChargeSecuredStatement) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ChargeSecuredStatement) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ChargeSecuredStatement) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ChargeSecuredStatement) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEndDate

`func (o *ChargeSecuredStatement) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ChargeSecuredStatement) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ChargeSecuredStatement) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ChargeSecuredStatement) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *ChargeSecuredStatement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeSecuredStatement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeSecuredStatement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChargeSecuredStatement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *ChargeSecuredStatement) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ChargeSecuredStatement) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ChargeSecuredStatement) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *ChargeSecuredStatement) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *ChargeSecuredStatement) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ChargeSecuredStatement) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ChargeSecuredStatement) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ChargeSecuredStatement) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetAccountSummary

`func (o *ChargeSecuredStatement) GetAccountSummary() AccountSummary`

GetAccountSummary returns the AccountSummary field if non-nil, zero value otherwise.

### GetAccountSummaryOk

`func (o *ChargeSecuredStatement) GetAccountSummaryOk() (*AccountSummary, bool)`

GetAccountSummaryOk returns a tuple with the AccountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSummary

`func (o *ChargeSecuredStatement) SetAccountSummary(v AccountSummary)`

SetAccountSummary sets AccountSummary field to given value.

### HasAccountSummary

`func (o *ChargeSecuredStatement) HasAccountSummary() bool`

HasAccountSummary returns a boolean if a field has been set.

### GetAuthorizedSigners

`func (o *ChargeSecuredStatement) GetAuthorizedSigners() []Person`

GetAuthorizedSigners returns the AuthorizedSigners field if non-nil, zero value otherwise.

### GetAuthorizedSignersOk

`func (o *ChargeSecuredStatement) GetAuthorizedSignersOk() (*[]Person, bool)`

GetAuthorizedSignersOk returns a tuple with the AuthorizedSigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedSigners

`func (o *ChargeSecuredStatement) SetAuthorizedSigners(v []Person)`

SetAuthorizedSigners sets AuthorizedSigners field to given value.

### HasAuthorizedSigners

`func (o *ChargeSecuredStatement) HasAuthorizedSigners() bool`

HasAuthorizedSigners returns a boolean if a field has been set.

### GetClosingBalance

`func (o *ChargeSecuredStatement) GetClosingBalance() int64`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *ChargeSecuredStatement) GetClosingBalanceOk() (*int64, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *ChargeSecuredStatement) SetClosingBalance(v int64)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *ChargeSecuredStatement) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetCustomerServiceDetails

`func (o *ChargeSecuredStatement) GetCustomerServiceDetails() CustomerServiceDetails`

GetCustomerServiceDetails returns the CustomerServiceDetails field if non-nil, zero value otherwise.

### GetCustomerServiceDetailsOk

`func (o *ChargeSecuredStatement) GetCustomerServiceDetailsOk() (*CustomerServiceDetails, bool)`

GetCustomerServiceDetailsOk returns a tuple with the CustomerServiceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerServiceDetails

`func (o *ChargeSecuredStatement) SetCustomerServiceDetails(v CustomerServiceDetails)`

SetCustomerServiceDetails sets CustomerServiceDetails field to given value.

### HasCustomerServiceDetails

`func (o *ChargeSecuredStatement) HasCustomerServiceDetails() bool`

HasCustomerServiceDetails returns a boolean if a field has been set.

### GetDisclosure

`func (o *ChargeSecuredStatement) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *ChargeSecuredStatement) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *ChargeSecuredStatement) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *ChargeSecuredStatement) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetExcludeJitTransactions

`func (o *ChargeSecuredStatement) GetExcludeJitTransactions() bool`

GetExcludeJitTransactions returns the ExcludeJitTransactions field if non-nil, zero value otherwise.

### GetExcludeJitTransactionsOk

`func (o *ChargeSecuredStatement) GetExcludeJitTransactionsOk() (*bool, bool)`

GetExcludeJitTransactionsOk returns a tuple with the ExcludeJitTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeJitTransactions

`func (o *ChargeSecuredStatement) SetExcludeJitTransactions(v bool)`

SetExcludeJitTransactions sets ExcludeJitTransactions field to given value.


### GetIncludeChildTransactions

`func (o *ChargeSecuredStatement) GetIncludeChildTransactions() bool`

GetIncludeChildTransactions returns the IncludeChildTransactions field if non-nil, zero value otherwise.

### GetIncludeChildTransactionsOk

`func (o *ChargeSecuredStatement) GetIncludeChildTransactionsOk() (*bool, bool)`

GetIncludeChildTransactionsOk returns a tuple with the IncludeChildTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChildTransactions

`func (o *ChargeSecuredStatement) SetIncludeChildTransactions(v bool)`

SetIncludeChildTransactions sets IncludeChildTransactions field to given value.


### GetJointAccountHolders

`func (o *ChargeSecuredStatement) GetJointAccountHolders() []Person`

GetJointAccountHolders returns the JointAccountHolders field if non-nil, zero value otherwise.

### GetJointAccountHoldersOk

`func (o *ChargeSecuredStatement) GetJointAccountHoldersOk() (*[]Person, bool)`

GetJointAccountHoldersOk returns a tuple with the JointAccountHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJointAccountHolders

`func (o *ChargeSecuredStatement) SetJointAccountHolders(v []Person)`

SetJointAccountHolders sets JointAccountHolders field to given value.

### HasJointAccountHolders

`func (o *ChargeSecuredStatement) HasJointAccountHolders() bool`

HasJointAccountHolders returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *ChargeSecuredStatement) GetOpeningBalance() int64`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *ChargeSecuredStatement) GetOpeningBalanceOk() (*int64, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *ChargeSecuredStatement) SetOpeningBalance(v int64)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *ChargeSecuredStatement) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetPrimaryAccountHolderBusiness

`func (o *ChargeSecuredStatement) GetPrimaryAccountHolderBusiness() Business`

GetPrimaryAccountHolderBusiness returns the PrimaryAccountHolderBusiness field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderBusinessOk

`func (o *ChargeSecuredStatement) GetPrimaryAccountHolderBusinessOk() (*Business, bool)`

GetPrimaryAccountHolderBusinessOk returns a tuple with the PrimaryAccountHolderBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderBusiness

`func (o *ChargeSecuredStatement) SetPrimaryAccountHolderBusiness(v Business)`

SetPrimaryAccountHolderBusiness sets PrimaryAccountHolderBusiness field to given value.

### HasPrimaryAccountHolderBusiness

`func (o *ChargeSecuredStatement) HasPrimaryAccountHolderBusiness() bool`

HasPrimaryAccountHolderBusiness returns a boolean if a field has been set.

### GetPrimaryAccountHolderPersonal

`func (o *ChargeSecuredStatement) GetPrimaryAccountHolderPersonal() Person`

GetPrimaryAccountHolderPersonal returns the PrimaryAccountHolderPersonal field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderPersonalOk

`func (o *ChargeSecuredStatement) GetPrimaryAccountHolderPersonalOk() (*Person, bool)`

GetPrimaryAccountHolderPersonalOk returns a tuple with the PrimaryAccountHolderPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderPersonal

`func (o *ChargeSecuredStatement) SetPrimaryAccountHolderPersonal(v Person)`

SetPrimaryAccountHolderPersonal sets PrimaryAccountHolderPersonal field to given value.

### HasPrimaryAccountHolderPersonal

`func (o *ChargeSecuredStatement) HasPrimaryAccountHolderPersonal() bool`

HasPrimaryAccountHolderPersonal returns a boolean if a field has been set.

### GetStatementType

`func (o *ChargeSecuredStatement) GetStatementType() StatementType`

GetStatementType returns the StatementType field if non-nil, zero value otherwise.

### GetStatementTypeOk

`func (o *ChargeSecuredStatement) GetStatementTypeOk() (*StatementType, bool)`

GetStatementTypeOk returns a tuple with the StatementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementType

`func (o *ChargeSecuredStatement) SetStatementType(v StatementType)`

SetStatementType sets StatementType field to given value.


### GetTotalTransactions

`func (o *ChargeSecuredStatement) GetTotalTransactions() int64`

GetTotalTransactions returns the TotalTransactions field if non-nil, zero value otherwise.

### GetTotalTransactionsOk

`func (o *ChargeSecuredStatement) GetTotalTransactionsOk() (*int64, bool)`

GetTotalTransactionsOk returns a tuple with the TotalTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactions

`func (o *ChargeSecuredStatement) SetTotalTransactions(v int64)`

SetTotalTransactions sets TotalTransactions field to given value.

### HasTotalTransactions

`func (o *ChargeSecuredStatement) HasTotalTransactions() bool`

HasTotalTransactions returns a boolean if a field has been set.

### GetCreditSummary

`func (o *ChargeSecuredStatement) GetCreditSummary() CreditSummary`

GetCreditSummary returns the CreditSummary field if non-nil, zero value otherwise.

### GetCreditSummaryOk

`func (o *ChargeSecuredStatement) GetCreditSummaryOk() (*CreditSummary, bool)`

GetCreditSummaryOk returns a tuple with the CreditSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditSummary

`func (o *ChargeSecuredStatement) SetCreditSummary(v CreditSummary)`

SetCreditSummary sets CreditSummary field to given value.


### GetSecurityAccountSummary

`func (o *ChargeSecuredStatement) GetSecurityAccountSummary() SecurityAccountSummary`

GetSecurityAccountSummary returns the SecurityAccountSummary field if non-nil, zero value otherwise.

### GetSecurityAccountSummaryOk

`func (o *ChargeSecuredStatement) GetSecurityAccountSummaryOk() (*SecurityAccountSummary, bool)`

GetSecurityAccountSummaryOk returns a tuple with the SecurityAccountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAccountSummary

`func (o *ChargeSecuredStatement) SetSecurityAccountSummary(v SecurityAccountSummary)`

SetSecurityAccountSummary sets SecurityAccountSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


