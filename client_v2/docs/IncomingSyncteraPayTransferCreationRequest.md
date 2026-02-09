# IncomingSyncteraPayTransferCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount in the source currency&#39;s minor unit. For example, 10000 would be $100 for USD. | 
**ConfigurationId** | **string** | The ID of the Incoming Synctera Pay configuration to be used for this transfer | 
**DestinationAccountId** | **string** | The ID of the account in the Synctera Platform that will receive the funds for this transfer. | 
**ExternalData** | Pointer to **map[string]interface{}** | External Data for FinTech use, copied to ledger transactions related to this transfer. | [optional] 
**PayeeId** | **string** | The ID of the person or business customer in the Synctera Platform that is the receiver of the this transfer. | 
**PayerId** | **string** | The ID of the person or business in the Synctera Platform that is the originator of the this transfer. Depending on the configuration being used, this does not necessarily need to be a customer. | 
**ReferenceId** | **string** | The unique network reference ID assigned to the transfer by the FinTech. This will be copied to the reference_id of the related ledger transaction. | 
**SettlementDate** | **string** |  | 
**SourceExternalAccountId** | **string** | The ID of the external account in the Synctera Platform that is the ultimate source of funds for this transfer. | 

## Methods

### NewIncomingSyncteraPayTransferCreationRequest

`func NewIncomingSyncteraPayTransferCreationRequest(amount int64, configurationId string, destinationAccountId string, payeeId string, payerId string, referenceId string, settlementDate string, sourceExternalAccountId string, ) *IncomingSyncteraPayTransferCreationRequest`

NewIncomingSyncteraPayTransferCreationRequest instantiates a new IncomingSyncteraPayTransferCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingSyncteraPayTransferCreationRequestWithDefaults

`func NewIncomingSyncteraPayTransferCreationRequestWithDefaults() *IncomingSyncteraPayTransferCreationRequest`

NewIncomingSyncteraPayTransferCreationRequestWithDefaults instantiates a new IncomingSyncteraPayTransferCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *IncomingSyncteraPayTransferCreationRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IncomingSyncteraPayTransferCreationRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IncomingSyncteraPayTransferCreationRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetConfigurationId

`func (o *IncomingSyncteraPayTransferCreationRequest) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *IncomingSyncteraPayTransferCreationRequest) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *IncomingSyncteraPayTransferCreationRequest) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.


### GetDestinationAccountId

`func (o *IncomingSyncteraPayTransferCreationRequest) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *IncomingSyncteraPayTransferCreationRequest) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *IncomingSyncteraPayTransferCreationRequest) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.


### GetExternalData

`func (o *IncomingSyncteraPayTransferCreationRequest) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *IncomingSyncteraPayTransferCreationRequest) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *IncomingSyncteraPayTransferCreationRequest) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.

### HasExternalData

`func (o *IncomingSyncteraPayTransferCreationRequest) HasExternalData() bool`

HasExternalData returns a boolean if a field has been set.

### GetPayeeId

`func (o *IncomingSyncteraPayTransferCreationRequest) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *IncomingSyncteraPayTransferCreationRequest) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *IncomingSyncteraPayTransferCreationRequest) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetPayerId

`func (o *IncomingSyncteraPayTransferCreationRequest) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *IncomingSyncteraPayTransferCreationRequest) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *IncomingSyncteraPayTransferCreationRequest) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetReferenceId

`func (o *IncomingSyncteraPayTransferCreationRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *IncomingSyncteraPayTransferCreationRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *IncomingSyncteraPayTransferCreationRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetSettlementDate

`func (o *IncomingSyncteraPayTransferCreationRequest) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *IncomingSyncteraPayTransferCreationRequest) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *IncomingSyncteraPayTransferCreationRequest) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.


### GetSourceExternalAccountId

`func (o *IncomingSyncteraPayTransferCreationRequest) GetSourceExternalAccountId() string`

GetSourceExternalAccountId returns the SourceExternalAccountId field if non-nil, zero value otherwise.

### GetSourceExternalAccountIdOk

`func (o *IncomingSyncteraPayTransferCreationRequest) GetSourceExternalAccountIdOk() (*string, bool)`

GetSourceExternalAccountIdOk returns a tuple with the SourceExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceExternalAccountId

`func (o *IncomingSyncteraPayTransferCreationRequest) SetSourceExternalAccountId(v string)`

SetSourceExternalAccountId sets SourceExternalAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


