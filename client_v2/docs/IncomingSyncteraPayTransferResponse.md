# IncomingSyncteraPayTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount in the source currency&#39;s minor unit. For example, 10000 would be $100 for USD. | 
**ConfigurationId** | **string** | The ID of the Incoming Synctera Pay configuration used for this transfer | 
**CreationTime** | **time.Time** | The timestamp representing when the dispute was created | 
**DestinationAccountId** | **string** | The ID of the account in the Synctera Platform that will receive the funds for this transfer. | 
**ExternalData** | Pointer to **map[string]interface{}** | External Data for FinTech use, copied to ledger transactions related to this transfer. | [optional] 
**Id** | **string** | The unique ID of the Incoming Synctera Pay transfer | 
**LastUpdatedTime** | **time.Time** | The timestamp representing when the dispute was last modified | 
**PayeeId** | **string** | The ID of the person or business customer in the Synctera Platform that is the receiver of the this transfer. | 
**PayerId** | **string** | The ID of the person or business in the Synctera Platform that is the originator of the this transfer. Depending on the configuration being used, this does not necessarily need to be a customer. | 
**ReferenceId** | **string** | The unique network reference ID assigned to the transfer by the FinTech. | 
**SettlementDate** | **string** |  | 
**SourceExternalAccountId** | **string** | The ID of the external account in the Synctera Platform that is the ultimate source of funds for this transfer. | 
**Status** | [**IncomingSyncteraPayStatus**](IncomingSyncteraPayStatus.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**TransactionId** | Pointer to **string** | The ID of the transaction in the Synctera Ledger related to this transfer. | [optional] 

## Methods

### NewIncomingSyncteraPayTransferResponse

`func NewIncomingSyncteraPayTransferResponse(amount int64, configurationId string, creationTime time.Time, destinationAccountId string, id string, lastUpdatedTime time.Time, payeeId string, payerId string, referenceId string, settlementDate string, sourceExternalAccountId string, status IncomingSyncteraPayStatus, tenant string, ) *IncomingSyncteraPayTransferResponse`

NewIncomingSyncteraPayTransferResponse instantiates a new IncomingSyncteraPayTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingSyncteraPayTransferResponseWithDefaults

`func NewIncomingSyncteraPayTransferResponseWithDefaults() *IncomingSyncteraPayTransferResponse`

NewIncomingSyncteraPayTransferResponseWithDefaults instantiates a new IncomingSyncteraPayTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *IncomingSyncteraPayTransferResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IncomingSyncteraPayTransferResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IncomingSyncteraPayTransferResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetConfigurationId

`func (o *IncomingSyncteraPayTransferResponse) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *IncomingSyncteraPayTransferResponse) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *IncomingSyncteraPayTransferResponse) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetCreationTime

`func (o *IncomingSyncteraPayTransferResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *IncomingSyncteraPayTransferResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *IncomingSyncteraPayTransferResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetDestinationAccountId

`func (o *IncomingSyncteraPayTransferResponse) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *IncomingSyncteraPayTransferResponse) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *IncomingSyncteraPayTransferResponse) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.


### GetExternalData

`func (o *IncomingSyncteraPayTransferResponse) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *IncomingSyncteraPayTransferResponse) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *IncomingSyncteraPayTransferResponse) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.

### HasExternalData

`func (o *IncomingSyncteraPayTransferResponse) HasExternalData() bool`

HasExternalData returns a boolean if a field has been set.

### GetId

`func (o *IncomingSyncteraPayTransferResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncomingSyncteraPayTransferResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncomingSyncteraPayTransferResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *IncomingSyncteraPayTransferResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *IncomingSyncteraPayTransferResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *IncomingSyncteraPayTransferResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetPayeeId

`func (o *IncomingSyncteraPayTransferResponse) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *IncomingSyncteraPayTransferResponse) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *IncomingSyncteraPayTransferResponse) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetPayerId

`func (o *IncomingSyncteraPayTransferResponse) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *IncomingSyncteraPayTransferResponse) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *IncomingSyncteraPayTransferResponse) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetReferenceId

`func (o *IncomingSyncteraPayTransferResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *IncomingSyncteraPayTransferResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *IncomingSyncteraPayTransferResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetSettlementDate

`func (o *IncomingSyncteraPayTransferResponse) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *IncomingSyncteraPayTransferResponse) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *IncomingSyncteraPayTransferResponse) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.


### GetSourceExternalAccountId

`func (o *IncomingSyncteraPayTransferResponse) GetSourceExternalAccountId() string`

GetSourceExternalAccountId returns the SourceExternalAccountId field if non-nil, zero value otherwise.

### GetSourceExternalAccountIdOk

`func (o *IncomingSyncteraPayTransferResponse) GetSourceExternalAccountIdOk() (*string, bool)`

GetSourceExternalAccountIdOk returns a tuple with the SourceExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceExternalAccountId

`func (o *IncomingSyncteraPayTransferResponse) SetSourceExternalAccountId(v string)`

SetSourceExternalAccountId sets SourceExternalAccountId field to given value.


### GetStatus

`func (o *IncomingSyncteraPayTransferResponse) GetStatus() IncomingSyncteraPayStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncomingSyncteraPayTransferResponse) GetStatusOk() (*IncomingSyncteraPayStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncomingSyncteraPayTransferResponse) SetStatus(v IncomingSyncteraPayStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *IncomingSyncteraPayTransferResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IncomingSyncteraPayTransferResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IncomingSyncteraPayTransferResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionId

`func (o *IncomingSyncteraPayTransferResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *IncomingSyncteraPayTransferResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *IncomingSyncteraPayTransferResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *IncomingSyncteraPayTransferResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


