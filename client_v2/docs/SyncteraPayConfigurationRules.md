# SyncteraPayConfigurationRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforcePayee** | Pointer to **bool** | If true, the customer matches the payee if and only if the peer_to_peer flag is set.  | [optional] 
**HoldTime** | Pointer to **int32** | The number of minutes that the Outgoing Synctera Pay transfer will be held before being processed.  | [optional] 
**PayeeAccountRequired** | Pointer to **bool** | If true, external account creation is mandatory and customer must pass final_external_account_id in the payload.  | [optional] 
**PeerToPeer** | Pointer to **bool** | If true, peer-to-peer payments are allowed. Has no effect if enforce_payee is false.  | [optional] 
**RequiredDetails** | Pointer to [**[]SyncteraPayConfigurationRequiredDetails**](SyncteraPayConfigurationRequiredDetails.md) | The required objects on a transfer that must be present in order to post the transfer. Note that a transfer can be initiated without these objects, but they must be present in order to post the transfer.  | [optional] 

## Methods

### NewSyncteraPayConfigurationRules

`func NewSyncteraPayConfigurationRules() *SyncteraPayConfigurationRules`

NewSyncteraPayConfigurationRules instantiates a new SyncteraPayConfigurationRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayConfigurationRulesWithDefaults

`func NewSyncteraPayConfigurationRulesWithDefaults() *SyncteraPayConfigurationRules`

NewSyncteraPayConfigurationRulesWithDefaults instantiates a new SyncteraPayConfigurationRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforcePayee

`func (o *SyncteraPayConfigurationRules) GetEnforcePayee() bool`

GetEnforcePayee returns the EnforcePayee field if non-nil, zero value otherwise.

### GetEnforcePayeeOk

`func (o *SyncteraPayConfigurationRules) GetEnforcePayeeOk() (*bool, bool)`

GetEnforcePayeeOk returns a tuple with the EnforcePayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcePayee

`func (o *SyncteraPayConfigurationRules) SetEnforcePayee(v bool)`

SetEnforcePayee sets EnforcePayee field to given value.

### HasEnforcePayee

`func (o *SyncteraPayConfigurationRules) HasEnforcePayee() bool`

HasEnforcePayee returns a boolean if a field has been set.

### GetHoldTime

`func (o *SyncteraPayConfigurationRules) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *SyncteraPayConfigurationRules) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *SyncteraPayConfigurationRules) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *SyncteraPayConfigurationRules) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### GetPayeeAccountRequired

`func (o *SyncteraPayConfigurationRules) GetPayeeAccountRequired() bool`

GetPayeeAccountRequired returns the PayeeAccountRequired field if non-nil, zero value otherwise.

### GetPayeeAccountRequiredOk

`func (o *SyncteraPayConfigurationRules) GetPayeeAccountRequiredOk() (*bool, bool)`

GetPayeeAccountRequiredOk returns a tuple with the PayeeAccountRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeAccountRequired

`func (o *SyncteraPayConfigurationRules) SetPayeeAccountRequired(v bool)`

SetPayeeAccountRequired sets PayeeAccountRequired field to given value.

### HasPayeeAccountRequired

`func (o *SyncteraPayConfigurationRules) HasPayeeAccountRequired() bool`

HasPayeeAccountRequired returns a boolean if a field has been set.

### GetPeerToPeer

`func (o *SyncteraPayConfigurationRules) GetPeerToPeer() bool`

GetPeerToPeer returns the PeerToPeer field if non-nil, zero value otherwise.

### GetPeerToPeerOk

`func (o *SyncteraPayConfigurationRules) GetPeerToPeerOk() (*bool, bool)`

GetPeerToPeerOk returns a tuple with the PeerToPeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerToPeer

`func (o *SyncteraPayConfigurationRules) SetPeerToPeer(v bool)`

SetPeerToPeer sets PeerToPeer field to given value.

### HasPeerToPeer

`func (o *SyncteraPayConfigurationRules) HasPeerToPeer() bool`

HasPeerToPeer returns a boolean if a field has been set.

### GetRequiredDetails

`func (o *SyncteraPayConfigurationRules) GetRequiredDetails() []SyncteraPayConfigurationRequiredDetails`

GetRequiredDetails returns the RequiredDetails field if non-nil, zero value otherwise.

### GetRequiredDetailsOk

`func (o *SyncteraPayConfigurationRules) GetRequiredDetailsOk() (*[]SyncteraPayConfigurationRequiredDetails, bool)`

GetRequiredDetailsOk returns a tuple with the RequiredDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDetails

`func (o *SyncteraPayConfigurationRules) SetRequiredDetails(v []SyncteraPayConfigurationRequiredDetails)`

SetRequiredDetails sets RequiredDetails field to given value.

### HasRequiredDetails

`func (o *SyncteraPayConfigurationRules) HasRequiredDetails() bool`

HasRequiredDetails returns a boolean if a field has been set.

### SetRequiredDetailsNil

`func (o *SyncteraPayConfigurationRules) SetRequiredDetailsNil(b bool)`

 SetRequiredDetailsNil sets the value for RequiredDetails to be an explicit nil

### UnsetRequiredDetails
`func (o *SyncteraPayConfigurationRules) UnsetRequiredDetails()`

UnsetRequiredDetails ensures that no value is present for RequiredDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


