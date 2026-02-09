# GatewayAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier of the account. | 
**AcquirerReferenceId** | Pointer to **string** |  | [optional] 
**Amount** | [**GatewayAuthorizationAmount**](GatewayAuthorizationAmount.md) |  | 
**Balance** | [**GatewayAuthorizationBalance**](GatewayAuthorizationBalance.md) |  | 
**BanknetReferenceId** | Pointer to **string** |  | [optional] 
**Bin** | [**GatewayAuthorizationBin**](GatewayAuthorizationBin.md) |  | 
**CardCategory** | Pointer to **string** |  | [optional] 
**CardFormat** | [**GatewayAuthorizationCardFormat**](GatewayAuthorizationCardFormat.md) |  | 
**CardId** | **string** | Unique identifier of the card. | 
**CustomerId** | **string** | Unique identifier of the customer. | 
**DcSign** | Pointer to **string** |  | [optional] 
**DigitalWalletTokenId** | Pointer to **string** | Unique identifier of the digital wallet token. | [optional] 
**DigitalWalletTokenReferenceId** | Pointer to **string** |  | [optional] 
**EnhancedTransaction** | Pointer to [**GatewayAuthorizationEnhancedTransaction**](GatewayAuthorizationEnhancedTransaction.md) |  | [optional] 
**Fees** | Pointer to [**GatewayAuthorizationFees**](GatewayAuthorizationFees.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier assigned to each individual request. | [optional] 
**IsSurchargeFree** | Pointer to **bool** | Indicates if the transaction is surcharge free. | [optional] 
**LastFour** | **string** | Last four digits of the card PAN. | 
**Merchant** | [**GatewayAuthorizationMerchant**](GatewayAuthorizationMerchant.md) |  | 
**Network** | Pointer to **string** |  | [optional] 
**NetworkFraud** | [**GatewayAuthorizationNetworkFraud**](GatewayAuthorizationNetworkFraud.md) |  | 
**PaymentChannel** | Pointer to **string** |  | [optional] 
**Pos** | Pointer to [**GatewayAuthorizationPos**](GatewayAuthorizationPos.md) |  | [optional] 
**Processor** | [**GatewayAuthorizationProcessor**](GatewayAuthorizationProcessor.md) |  | 
**ProcessorData** | **map[string]interface{}** | An unstructured JSON blob containing raw transaction data from the processor. | 
**SettlementDate** | **time.Time** |  | 
**Subnetwork** | **string** |  | 
**ThreeDsAuthenticationStatus** | Pointer to **string** |  | [optional] 
**TransactionId** | **string** | Unique identifier of the transaction. | 
**Type** | **string** |  | 
**User** | [**GatewayAuthorizationUser**](GatewayAuthorizationUser.md) |  | 
**UserTransactionTime** | **time.Time** | The time when the transaction occurred. | 
**Verification** | Pointer to [**GatewayAuthorizationVerification**](GatewayAuthorizationVerification.md) |  | [optional] 

## Methods

### NewGatewayAuthorization

`func NewGatewayAuthorization(accountId string, amount GatewayAuthorizationAmount, balance GatewayAuthorizationBalance, bin GatewayAuthorizationBin, cardFormat GatewayAuthorizationCardFormat, cardId string, customerId string, lastFour string, merchant GatewayAuthorizationMerchant, networkFraud GatewayAuthorizationNetworkFraud, processor GatewayAuthorizationProcessor, processorData map[string]interface{}, settlementDate time.Time, subnetwork string, transactionId string, type_ string, user GatewayAuthorizationUser, userTransactionTime time.Time, ) *GatewayAuthorization`

NewGatewayAuthorization instantiates a new GatewayAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayAuthorizationWithDefaults

`func NewGatewayAuthorizationWithDefaults() *GatewayAuthorization`

NewGatewayAuthorizationWithDefaults instantiates a new GatewayAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GatewayAuthorization) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GatewayAuthorization) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GatewayAuthorization) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAcquirerReferenceId

`func (o *GatewayAuthorization) GetAcquirerReferenceId() string`

GetAcquirerReferenceId returns the AcquirerReferenceId field if non-nil, zero value otherwise.

### GetAcquirerReferenceIdOk

`func (o *GatewayAuthorization) GetAcquirerReferenceIdOk() (*string, bool)`

GetAcquirerReferenceIdOk returns a tuple with the AcquirerReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReferenceId

`func (o *GatewayAuthorization) SetAcquirerReferenceId(v string)`

SetAcquirerReferenceId sets AcquirerReferenceId field to given value.

### HasAcquirerReferenceId

`func (o *GatewayAuthorization) HasAcquirerReferenceId() bool`

HasAcquirerReferenceId returns a boolean if a field has been set.

### GetAmount

`func (o *GatewayAuthorization) GetAmount() GatewayAuthorizationAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GatewayAuthorization) GetAmountOk() (*GatewayAuthorizationAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GatewayAuthorization) SetAmount(v GatewayAuthorizationAmount)`

SetAmount sets Amount field to given value.


### GetBalance

`func (o *GatewayAuthorization) GetBalance() GatewayAuthorizationBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GatewayAuthorization) GetBalanceOk() (*GatewayAuthorizationBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GatewayAuthorization) SetBalance(v GatewayAuthorizationBalance)`

SetBalance sets Balance field to given value.


### GetBanknetReferenceId

`func (o *GatewayAuthorization) GetBanknetReferenceId() string`

GetBanknetReferenceId returns the BanknetReferenceId field if non-nil, zero value otherwise.

### GetBanknetReferenceIdOk

`func (o *GatewayAuthorization) GetBanknetReferenceIdOk() (*string, bool)`

GetBanknetReferenceIdOk returns a tuple with the BanknetReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanknetReferenceId

`func (o *GatewayAuthorization) SetBanknetReferenceId(v string)`

SetBanknetReferenceId sets BanknetReferenceId field to given value.

### HasBanknetReferenceId

`func (o *GatewayAuthorization) HasBanknetReferenceId() bool`

HasBanknetReferenceId returns a boolean if a field has been set.

### GetBin

`func (o *GatewayAuthorization) GetBin() GatewayAuthorizationBin`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *GatewayAuthorization) GetBinOk() (*GatewayAuthorizationBin, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *GatewayAuthorization) SetBin(v GatewayAuthorizationBin)`

SetBin sets Bin field to given value.


### GetCardCategory

`func (o *GatewayAuthorization) GetCardCategory() string`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *GatewayAuthorization) GetCardCategoryOk() (*string, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *GatewayAuthorization) SetCardCategory(v string)`

SetCardCategory sets CardCategory field to given value.

### HasCardCategory

`func (o *GatewayAuthorization) HasCardCategory() bool`

HasCardCategory returns a boolean if a field has been set.

### GetCardFormat

`func (o *GatewayAuthorization) GetCardFormat() GatewayAuthorizationCardFormat`

GetCardFormat returns the CardFormat field if non-nil, zero value otherwise.

### GetCardFormatOk

`func (o *GatewayAuthorization) GetCardFormatOk() (*GatewayAuthorizationCardFormat, bool)`

GetCardFormatOk returns a tuple with the CardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFormat

`func (o *GatewayAuthorization) SetCardFormat(v GatewayAuthorizationCardFormat)`

SetCardFormat sets CardFormat field to given value.


### GetCardId

`func (o *GatewayAuthorization) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *GatewayAuthorization) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *GatewayAuthorization) SetCardId(v string)`

SetCardId sets CardId field to given value.


### GetCustomerId

`func (o *GatewayAuthorization) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GatewayAuthorization) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GatewayAuthorization) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDcSign

`func (o *GatewayAuthorization) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *GatewayAuthorization) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *GatewayAuthorization) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.

### HasDcSign

`func (o *GatewayAuthorization) HasDcSign() bool`

HasDcSign returns a boolean if a field has been set.

### GetDigitalWalletTokenId

`func (o *GatewayAuthorization) GetDigitalWalletTokenId() string`

GetDigitalWalletTokenId returns the DigitalWalletTokenId field if non-nil, zero value otherwise.

### GetDigitalWalletTokenIdOk

`func (o *GatewayAuthorization) GetDigitalWalletTokenIdOk() (*string, bool)`

GetDigitalWalletTokenIdOk returns a tuple with the DigitalWalletTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokenId

`func (o *GatewayAuthorization) SetDigitalWalletTokenId(v string)`

SetDigitalWalletTokenId sets DigitalWalletTokenId field to given value.

### HasDigitalWalletTokenId

`func (o *GatewayAuthorization) HasDigitalWalletTokenId() bool`

HasDigitalWalletTokenId returns a boolean if a field has been set.

### GetDigitalWalletTokenReferenceId

`func (o *GatewayAuthorization) GetDigitalWalletTokenReferenceId() string`

GetDigitalWalletTokenReferenceId returns the DigitalWalletTokenReferenceId field if non-nil, zero value otherwise.

### GetDigitalWalletTokenReferenceIdOk

`func (o *GatewayAuthorization) GetDigitalWalletTokenReferenceIdOk() (*string, bool)`

GetDigitalWalletTokenReferenceIdOk returns a tuple with the DigitalWalletTokenReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokenReferenceId

`func (o *GatewayAuthorization) SetDigitalWalletTokenReferenceId(v string)`

SetDigitalWalletTokenReferenceId sets DigitalWalletTokenReferenceId field to given value.

### HasDigitalWalletTokenReferenceId

`func (o *GatewayAuthorization) HasDigitalWalletTokenReferenceId() bool`

HasDigitalWalletTokenReferenceId returns a boolean if a field has been set.

### GetEnhancedTransaction

`func (o *GatewayAuthorization) GetEnhancedTransaction() GatewayAuthorizationEnhancedTransaction`

GetEnhancedTransaction returns the EnhancedTransaction field if non-nil, zero value otherwise.

### GetEnhancedTransactionOk

`func (o *GatewayAuthorization) GetEnhancedTransactionOk() (*GatewayAuthorizationEnhancedTransaction, bool)`

GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedTransaction

`func (o *GatewayAuthorization) SetEnhancedTransaction(v GatewayAuthorizationEnhancedTransaction)`

SetEnhancedTransaction sets EnhancedTransaction field to given value.

### HasEnhancedTransaction

`func (o *GatewayAuthorization) HasEnhancedTransaction() bool`

HasEnhancedTransaction returns a boolean if a field has been set.

### GetFees

`func (o *GatewayAuthorization) GetFees() GatewayAuthorizationFees`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *GatewayAuthorization) GetFeesOk() (*GatewayAuthorizationFees, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *GatewayAuthorization) SetFees(v GatewayAuthorizationFees)`

SetFees sets Fees field to given value.

### HasFees

`func (o *GatewayAuthorization) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetId

`func (o *GatewayAuthorization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayAuthorization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayAuthorization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayAuthorization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSurchargeFree

`func (o *GatewayAuthorization) GetIsSurchargeFree() bool`

GetIsSurchargeFree returns the IsSurchargeFree field if non-nil, zero value otherwise.

### GetIsSurchargeFreeOk

`func (o *GatewayAuthorization) GetIsSurchargeFreeOk() (*bool, bool)`

GetIsSurchargeFreeOk returns a tuple with the IsSurchargeFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSurchargeFree

`func (o *GatewayAuthorization) SetIsSurchargeFree(v bool)`

SetIsSurchargeFree sets IsSurchargeFree field to given value.

### HasIsSurchargeFree

`func (o *GatewayAuthorization) HasIsSurchargeFree() bool`

HasIsSurchargeFree returns a boolean if a field has been set.

### GetLastFour

`func (o *GatewayAuthorization) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *GatewayAuthorization) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *GatewayAuthorization) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.


### GetMerchant

`func (o *GatewayAuthorization) GetMerchant() GatewayAuthorizationMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *GatewayAuthorization) GetMerchantOk() (*GatewayAuthorizationMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *GatewayAuthorization) SetMerchant(v GatewayAuthorizationMerchant)`

SetMerchant sets Merchant field to given value.


### GetNetwork

`func (o *GatewayAuthorization) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GatewayAuthorization) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GatewayAuthorization) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GatewayAuthorization) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkFraud

`func (o *GatewayAuthorization) GetNetworkFraud() GatewayAuthorizationNetworkFraud`

GetNetworkFraud returns the NetworkFraud field if non-nil, zero value otherwise.

### GetNetworkFraudOk

`func (o *GatewayAuthorization) GetNetworkFraudOk() (*GatewayAuthorizationNetworkFraud, bool)`

GetNetworkFraudOk returns a tuple with the NetworkFraud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFraud

`func (o *GatewayAuthorization) SetNetworkFraud(v GatewayAuthorizationNetworkFraud)`

SetNetworkFraud sets NetworkFraud field to given value.


### GetPaymentChannel

`func (o *GatewayAuthorization) GetPaymentChannel() string`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *GatewayAuthorization) GetPaymentChannelOk() (*string, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *GatewayAuthorization) SetPaymentChannel(v string)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *GatewayAuthorization) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetPos

`func (o *GatewayAuthorization) GetPos() GatewayAuthorizationPos`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *GatewayAuthorization) GetPosOk() (*GatewayAuthorizationPos, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *GatewayAuthorization) SetPos(v GatewayAuthorizationPos)`

SetPos sets Pos field to given value.

### HasPos

`func (o *GatewayAuthorization) HasPos() bool`

HasPos returns a boolean if a field has been set.

### GetProcessor

`func (o *GatewayAuthorization) GetProcessor() GatewayAuthorizationProcessor`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *GatewayAuthorization) GetProcessorOk() (*GatewayAuthorizationProcessor, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *GatewayAuthorization) SetProcessor(v GatewayAuthorizationProcessor)`

SetProcessor sets Processor field to given value.


### GetProcessorData

`func (o *GatewayAuthorization) GetProcessorData() map[string]interface{}`

GetProcessorData returns the ProcessorData field if non-nil, zero value otherwise.

### GetProcessorDataOk

`func (o *GatewayAuthorization) GetProcessorDataOk() (*map[string]interface{}, bool)`

GetProcessorDataOk returns a tuple with the ProcessorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorData

`func (o *GatewayAuthorization) SetProcessorData(v map[string]interface{})`

SetProcessorData sets ProcessorData field to given value.


### GetSettlementDate

`func (o *GatewayAuthorization) GetSettlementDate() time.Time`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *GatewayAuthorization) GetSettlementDateOk() (*time.Time, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *GatewayAuthorization) SetSettlementDate(v time.Time)`

SetSettlementDate sets SettlementDate field to given value.


### GetSubnetwork

`func (o *GatewayAuthorization) GetSubnetwork() string`

GetSubnetwork returns the Subnetwork field if non-nil, zero value otherwise.

### GetSubnetworkOk

`func (o *GatewayAuthorization) GetSubnetworkOk() (*string, bool)`

GetSubnetworkOk returns a tuple with the Subnetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetwork

`func (o *GatewayAuthorization) SetSubnetwork(v string)`

SetSubnetwork sets Subnetwork field to given value.


### GetThreeDsAuthenticationStatus

`func (o *GatewayAuthorization) GetThreeDsAuthenticationStatus() string`

GetThreeDsAuthenticationStatus returns the ThreeDsAuthenticationStatus field if non-nil, zero value otherwise.

### GetThreeDsAuthenticationStatusOk

`func (o *GatewayAuthorization) GetThreeDsAuthenticationStatusOk() (*string, bool)`

GetThreeDsAuthenticationStatusOk returns a tuple with the ThreeDsAuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsAuthenticationStatus

`func (o *GatewayAuthorization) SetThreeDsAuthenticationStatus(v string)`

SetThreeDsAuthenticationStatus sets ThreeDsAuthenticationStatus field to given value.

### HasThreeDsAuthenticationStatus

`func (o *GatewayAuthorization) HasThreeDsAuthenticationStatus() bool`

HasThreeDsAuthenticationStatus returns a boolean if a field has been set.

### GetTransactionId

`func (o *GatewayAuthorization) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GatewayAuthorization) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GatewayAuthorization) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetType

`func (o *GatewayAuthorization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayAuthorization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayAuthorization) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *GatewayAuthorization) GetUser() GatewayAuthorizationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GatewayAuthorization) GetUserOk() (*GatewayAuthorizationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GatewayAuthorization) SetUser(v GatewayAuthorizationUser)`

SetUser sets User field to given value.


### GetUserTransactionTime

`func (o *GatewayAuthorization) GetUserTransactionTime() time.Time`

GetUserTransactionTime returns the UserTransactionTime field if non-nil, zero value otherwise.

### GetUserTransactionTimeOk

`func (o *GatewayAuthorization) GetUserTransactionTimeOk() (*time.Time, bool)`

GetUserTransactionTimeOk returns a tuple with the UserTransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTransactionTime

`func (o *GatewayAuthorization) SetUserTransactionTime(v time.Time)`

SetUserTransactionTime sets UserTransactionTime field to given value.


### GetVerification

`func (o *GatewayAuthorization) GetVerification() GatewayAuthorizationVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *GatewayAuthorization) GetVerificationOk() (*GatewayAuthorizationVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *GatewayAuthorization) SetVerification(v GatewayAuthorizationVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *GatewayAuthorization) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


