# AchTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The ACH company entry description | [optional] 
**FileName** | Pointer to **string** | For outgoing ACH, the name of the file the entry went out in. For incoming ACH, the name of the file the entry came from. The value will be omitted for outgoing payments that have not yet been written into a file. | [optional] 
**Id** | **string** | The ACH payment uuid (used in &#x60;/v0/ach&#x60; endpoint) | 
**OriginatingAccountId** | Pointer to **string** | The uuid of the account originating the ACH. This will be a customer account uuid for outgoing ACH, and omitted for incoming ACH. | [optional] 
**OriginatorName** | Pointer to **string** | The name of the originator according to the ACH entry. This should map to the ACH &#x60;Identification Number&#x60; field, if provided. | [optional] 
**ReceivingAccountId** | Pointer to **string** | The uuid of the account receiving the ACH entry. In the case of an outgoing ACH, this will be an external_account uuid. For incoming ACH, this will be an account uuid. | [optional] 
**RecipientName** | Pointer to **string** | The name of the recipient according to the ACH entry. This should map to the ACH &#x60;Individual Name&#x60; field. | [optional] 
**ReturnCode** | Pointer to **string** | The ACH return code, if this transaction was a return | [optional] 
**TraceNumber** | Pointer to **string** | The ACH trace number associated with the transaction | [optional] 

## Methods

### NewAchTransactionData

`func NewAchTransactionData(id string, ) *AchTransactionData`

NewAchTransactionData instantiates a new AchTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchTransactionDataWithDefaults

`func NewAchTransactionDataWithDefaults() *AchTransactionData`

NewAchTransactionDataWithDefaults instantiates a new AchTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AchTransactionData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AchTransactionData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AchTransactionData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AchTransactionData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileName

`func (o *AchTransactionData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AchTransactionData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AchTransactionData) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AchTransactionData) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetId

`func (o *AchTransactionData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AchTransactionData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AchTransactionData) SetId(v string)`

SetId sets Id field to given value.


### GetOriginatingAccountId

`func (o *AchTransactionData) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *AchTransactionData) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *AchTransactionData) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.

### HasOriginatingAccountId

`func (o *AchTransactionData) HasOriginatingAccountId() bool`

HasOriginatingAccountId returns a boolean if a field has been set.

### GetOriginatorName

`func (o *AchTransactionData) GetOriginatorName() string`

GetOriginatorName returns the OriginatorName field if non-nil, zero value otherwise.

### GetOriginatorNameOk

`func (o *AchTransactionData) GetOriginatorNameOk() (*string, bool)`

GetOriginatorNameOk returns a tuple with the OriginatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorName

`func (o *AchTransactionData) SetOriginatorName(v string)`

SetOriginatorName sets OriginatorName field to given value.

### HasOriginatorName

`func (o *AchTransactionData) HasOriginatorName() bool`

HasOriginatorName returns a boolean if a field has been set.

### GetReceivingAccountId

`func (o *AchTransactionData) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *AchTransactionData) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *AchTransactionData) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.

### HasReceivingAccountId

`func (o *AchTransactionData) HasReceivingAccountId() bool`

HasReceivingAccountId returns a boolean if a field has been set.

### GetRecipientName

`func (o *AchTransactionData) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *AchTransactionData) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *AchTransactionData) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *AchTransactionData) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### GetReturnCode

`func (o *AchTransactionData) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *AchTransactionData) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *AchTransactionData) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *AchTransactionData) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetTraceNumber

`func (o *AchTransactionData) GetTraceNumber() string`

GetTraceNumber returns the TraceNumber field if non-nil, zero value otherwise.

### GetTraceNumberOk

`func (o *AchTransactionData) GetTraceNumberOk() (*string, bool)`

GetTraceNumberOk returns a tuple with the TraceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceNumber

`func (o *AchTransactionData) SetTraceNumber(v string)`

SetTraceNumber sets TraceNumber field to given value.

### HasTraceNumber

`func (o *AchTransactionData) HasTraceNumber() bool`

HasTraceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


