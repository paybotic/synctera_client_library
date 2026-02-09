# IncomingSyncteraPayConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | **time.Time** | The timestamp representing when the dispute was created | 
**Currency** | **string** | ISO 4217 Alpha-3 code representing the currency to be credited to the Synctera account. | 
**Enabled** | **bool** | Determines if this Incoming Synctera Pay Configuration can be used to create new transfers. | 
**Id** | **string** | The unique ID of the Incoming Synctera Pay configuration | 
**LastUpdatedTime** | **time.Time** | The timestamp representing when the dispute was last modified | 
**Name** | **string** | User-supplied name for the configuration | 
**PayerMustBePayee** | **bool** | If true, transfers using this config must have the same payer and payee (\&quot;me-to-me\&quot;) and peer-to-peer transfers can not use this configuration. | 
**SettlementAccountId** | **string** | Account used for settlement of all transfers that use this configuration. | 
**SettlementAccountOwnerId** | **string** | Business or Person that owns the settlement account for all transfers that use this configuration. | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**TransactionSubtype** | [**IncomingSyncteraPaySubtype**](IncomingSyncteraPaySubtype.md) |  | 

## Methods

### NewIncomingSyncteraPayConfigurationResponse

`func NewIncomingSyncteraPayConfigurationResponse(creationTime time.Time, currency string, enabled bool, id string, lastUpdatedTime time.Time, name string, payerMustBePayee bool, settlementAccountId string, settlementAccountOwnerId string, tenant string, transactionSubtype IncomingSyncteraPaySubtype, ) *IncomingSyncteraPayConfigurationResponse`

NewIncomingSyncteraPayConfigurationResponse instantiates a new IncomingSyncteraPayConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingSyncteraPayConfigurationResponseWithDefaults

`func NewIncomingSyncteraPayConfigurationResponseWithDefaults() *IncomingSyncteraPayConfigurationResponse`

NewIncomingSyncteraPayConfigurationResponseWithDefaults instantiates a new IncomingSyncteraPayConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *IncomingSyncteraPayConfigurationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *IncomingSyncteraPayConfigurationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *IncomingSyncteraPayConfigurationResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IncomingSyncteraPayConfigurationResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetEnabled

`func (o *IncomingSyncteraPayConfigurationResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IncomingSyncteraPayConfigurationResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *IncomingSyncteraPayConfigurationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncomingSyncteraPayConfigurationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *IncomingSyncteraPayConfigurationResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *IncomingSyncteraPayConfigurationResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetName

`func (o *IncomingSyncteraPayConfigurationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncomingSyncteraPayConfigurationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPayerMustBePayee

`func (o *IncomingSyncteraPayConfigurationResponse) GetPayerMustBePayee() bool`

GetPayerMustBePayee returns the PayerMustBePayee field if non-nil, zero value otherwise.

### GetPayerMustBePayeeOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetPayerMustBePayeeOk() (*bool, bool)`

GetPayerMustBePayeeOk returns a tuple with the PayerMustBePayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerMustBePayee

`func (o *IncomingSyncteraPayConfigurationResponse) SetPayerMustBePayee(v bool)`

SetPayerMustBePayee sets PayerMustBePayee field to given value.


### GetSettlementAccountId

`func (o *IncomingSyncteraPayConfigurationResponse) GetSettlementAccountId() string`

GetSettlementAccountId returns the SettlementAccountId field if non-nil, zero value otherwise.

### GetSettlementAccountIdOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetSettlementAccountIdOk() (*string, bool)`

GetSettlementAccountIdOk returns a tuple with the SettlementAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAccountId

`func (o *IncomingSyncteraPayConfigurationResponse) SetSettlementAccountId(v string)`

SetSettlementAccountId sets SettlementAccountId field to given value.


### GetSettlementAccountOwnerId

`func (o *IncomingSyncteraPayConfigurationResponse) GetSettlementAccountOwnerId() string`

GetSettlementAccountOwnerId returns the SettlementAccountOwnerId field if non-nil, zero value otherwise.

### GetSettlementAccountOwnerIdOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetSettlementAccountOwnerIdOk() (*string, bool)`

GetSettlementAccountOwnerIdOk returns a tuple with the SettlementAccountOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAccountOwnerId

`func (o *IncomingSyncteraPayConfigurationResponse) SetSettlementAccountOwnerId(v string)`

SetSettlementAccountOwnerId sets SettlementAccountOwnerId field to given value.


### GetTenant

`func (o *IncomingSyncteraPayConfigurationResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IncomingSyncteraPayConfigurationResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionSubtype

`func (o *IncomingSyncteraPayConfigurationResponse) GetTransactionSubtype() IncomingSyncteraPaySubtype`

GetTransactionSubtype returns the TransactionSubtype field if non-nil, zero value otherwise.

### GetTransactionSubtypeOk

`func (o *IncomingSyncteraPayConfigurationResponse) GetTransactionSubtypeOk() (*IncomingSyncteraPaySubtype, bool)`

GetTransactionSubtypeOk returns a tuple with the TransactionSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSubtype

`func (o *IncomingSyncteraPayConfigurationResponse) SetTransactionSubtype(v IncomingSyncteraPaySubtype)`

SetTransactionSubtype sets TransactionSubtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


