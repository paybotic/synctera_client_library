# WireTransactionSimulationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Number of the receiving account | 
**Amount** | **int32** | Amount to transfer in cents (e.g. 100 &#x3D; $1). | 
**Network** | Pointer to **string** | Network to use for the Wire transfer | [optional] [default to "FEDWIRE"]

## Methods

### NewWireTransactionSimulationRequest

`func NewWireTransactionSimulationRequest(accountNumber string, amount int32, ) *WireTransactionSimulationRequest`

NewWireTransactionSimulationRequest instantiates a new WireTransactionSimulationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireTransactionSimulationRequestWithDefaults

`func NewWireTransactionSimulationRequestWithDefaults() *WireTransactionSimulationRequest`

NewWireTransactionSimulationRequestWithDefaults instantiates a new WireTransactionSimulationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *WireTransactionSimulationRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *WireTransactionSimulationRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *WireTransactionSimulationRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAmount

`func (o *WireTransactionSimulationRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WireTransactionSimulationRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WireTransactionSimulationRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetNetwork

`func (o *WireTransactionSimulationRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WireTransactionSimulationRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WireTransactionSimulationRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *WireTransactionSimulationRequest) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


