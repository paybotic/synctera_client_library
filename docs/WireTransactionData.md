# WireTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | The UUID if the customer associated with the wire payment | [optional] 
**FileName** | Pointer to **string** | The file name the wire was sent in or received from | [optional] 
**Id** | Pointer to **string** | The Synctera Wire payment UUID | [optional] 
**OriginatingAccountId** | Pointer to **string** | The UUID of the account that the wire payment is being sent from | [optional] 
**ReceivingAccountId** | Pointer to **string** | The UUID if the account that is receiving the wire | [optional] 
**RecipientMessage** | Pointer to **string** | The message to the recipient | [optional] 
**WireReferenceNumber** | Pointer to **string** | The wire Sender Reference Number for the transfer | [optional] 

## Methods

### NewWireTransactionData

`func NewWireTransactionData() *WireTransactionData`

NewWireTransactionData instantiates a new WireTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireTransactionDataWithDefaults

`func NewWireTransactionDataWithDefaults() *WireTransactionData`

NewWireTransactionDataWithDefaults instantiates a new WireTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *WireTransactionData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *WireTransactionData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *WireTransactionData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *WireTransactionData) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetFileName

`func (o *WireTransactionData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *WireTransactionData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *WireTransactionData) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *WireTransactionData) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetId

`func (o *WireTransactionData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireTransactionData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireTransactionData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WireTransactionData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *WireTransactionData) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *WireTransactionData) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *WireTransactionData) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.

### HasOriginatingAccountId

`func (o *WireTransactionData) HasOriginatingAccountId() bool`

HasOriginatingAccountId returns a boolean if a field has been set.

### GetReceivingAccountId

`func (o *WireTransactionData) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *WireTransactionData) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *WireTransactionData) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.

### HasReceivingAccountId

`func (o *WireTransactionData) HasReceivingAccountId() bool`

HasReceivingAccountId returns a boolean if a field has been set.

### GetRecipientMessage

`func (o *WireTransactionData) GetRecipientMessage() string`

GetRecipientMessage returns the RecipientMessage field if non-nil, zero value otherwise.

### GetRecipientMessageOk

`func (o *WireTransactionData) GetRecipientMessageOk() (*string, bool)`

GetRecipientMessageOk returns a tuple with the RecipientMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientMessage

`func (o *WireTransactionData) SetRecipientMessage(v string)`

SetRecipientMessage sets RecipientMessage field to given value.

### HasRecipientMessage

`func (o *WireTransactionData) HasRecipientMessage() bool`

HasRecipientMessage returns a boolean if a field has been set.

### GetWireReferenceNumber

`func (o *WireTransactionData) GetWireReferenceNumber() string`

GetWireReferenceNumber returns the WireReferenceNumber field if non-nil, zero value otherwise.

### GetWireReferenceNumberOk

`func (o *WireTransactionData) GetWireReferenceNumberOk() (*string, bool)`

GetWireReferenceNumberOk returns a tuple with the WireReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireReferenceNumber

`func (o *WireTransactionData) SetWireReferenceNumber(v string)`

SetWireReferenceNumber sets WireReferenceNumber field to given value.

### HasWireReferenceNumber

`func (o *WireTransactionData) HasWireReferenceNumber() bool`

HasWireReferenceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


