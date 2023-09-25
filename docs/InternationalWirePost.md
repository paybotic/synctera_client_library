# InternationalWirePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountInBeneficiaryCurrency** | **int64** | Transfer amount in cents ($100 would be 10000) in the currency of the beneficiary. | 
**BeneficiaryCurrency** | Pointer to **string** | The currency of the beneficiary&#39;s account in ISO4217 format. Optional, the currency will be inferred form the receiving account, validation will be performed if this field is provided, if the receiving account currency and this field do not match the API will return an error. | [optional] 
**CustomerId** | **string** | The customer UUID representing the person initiating the Wire transfer. | 
**Metadata** | Pointer to **map[string]interface{}** | Additional transfer metadata structured as key-value pairs. | [optional] 
**OriginatingAccountId** | **string** | The account uuid representing the sender of the wire. | 
**QuoteId** | Pointer to **string** | The quote id returned from the /international_wire/quote endpoint, contains FX rate information. Required if the international wire is cross currency. | [optional] 
**ReceivingAccountId** | **string** | The external account uuid representing the recipient of the wire. Needs to be configured with international bank routing details (e.g. SWIFT code, refer to external accounts documentation). | 
**SenderCurrency** | Pointer to **string** | The currency of the sender&#39;s account in ISO4217 format. Optional, the currency will be inferred form the originating account, validation will be performed if this field is provided, if the originating account currency and this field do not match the API will return an error. | [optional] 
**WireInstruction** | Pointer to **string** | Optional transaction instructions for the recipient bank if required by the recipient bank. | [optional] 
**WirePurpose** | **string** | Reason for the wire transfer. | 

## Methods

### NewInternationalWirePost

`func NewInternationalWirePost(amountInBeneficiaryCurrency int64, customerId string, originatingAccountId string, receivingAccountId string, wirePurpose string, ) *InternationalWirePost`

NewInternationalWirePost instantiates a new InternationalWirePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalWirePostWithDefaults

`func NewInternationalWirePostWithDefaults() *InternationalWirePost`

NewInternationalWirePostWithDefaults instantiates a new InternationalWirePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountInBeneficiaryCurrency

`func (o *InternationalWirePost) GetAmountInBeneficiaryCurrency() int64`

GetAmountInBeneficiaryCurrency returns the AmountInBeneficiaryCurrency field if non-nil, zero value otherwise.

### GetAmountInBeneficiaryCurrencyOk

`func (o *InternationalWirePost) GetAmountInBeneficiaryCurrencyOk() (*int64, bool)`

GetAmountInBeneficiaryCurrencyOk returns a tuple with the AmountInBeneficiaryCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountInBeneficiaryCurrency

`func (o *InternationalWirePost) SetAmountInBeneficiaryCurrency(v int64)`

SetAmountInBeneficiaryCurrency sets AmountInBeneficiaryCurrency field to given value.


### GetBeneficiaryCurrency

`func (o *InternationalWirePost) GetBeneficiaryCurrency() string`

GetBeneficiaryCurrency returns the BeneficiaryCurrency field if non-nil, zero value otherwise.

### GetBeneficiaryCurrencyOk

`func (o *InternationalWirePost) GetBeneficiaryCurrencyOk() (*string, bool)`

GetBeneficiaryCurrencyOk returns a tuple with the BeneficiaryCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryCurrency

`func (o *InternationalWirePost) SetBeneficiaryCurrency(v string)`

SetBeneficiaryCurrency sets BeneficiaryCurrency field to given value.

### HasBeneficiaryCurrency

`func (o *InternationalWirePost) HasBeneficiaryCurrency() bool`

HasBeneficiaryCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *InternationalWirePost) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InternationalWirePost) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InternationalWirePost) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetMetadata

`func (o *InternationalWirePost) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InternationalWirePost) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InternationalWirePost) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InternationalWirePost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *InternationalWirePost) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *InternationalWirePost) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *InternationalWirePost) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetQuoteId

`func (o *InternationalWirePost) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *InternationalWirePost) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *InternationalWirePost) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *InternationalWirePost) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetReceivingAccountId

`func (o *InternationalWirePost) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *InternationalWirePost) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *InternationalWirePost) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.


### GetSenderCurrency

`func (o *InternationalWirePost) GetSenderCurrency() string`

GetSenderCurrency returns the SenderCurrency field if non-nil, zero value otherwise.

### GetSenderCurrencyOk

`func (o *InternationalWirePost) GetSenderCurrencyOk() (*string, bool)`

GetSenderCurrencyOk returns a tuple with the SenderCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderCurrency

`func (o *InternationalWirePost) SetSenderCurrency(v string)`

SetSenderCurrency sets SenderCurrency field to given value.

### HasSenderCurrency

`func (o *InternationalWirePost) HasSenderCurrency() bool`

HasSenderCurrency returns a boolean if a field has been set.

### GetWireInstruction

`func (o *InternationalWirePost) GetWireInstruction() string`

GetWireInstruction returns the WireInstruction field if non-nil, zero value otherwise.

### GetWireInstructionOk

`func (o *InternationalWirePost) GetWireInstructionOk() (*string, bool)`

GetWireInstructionOk returns a tuple with the WireInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireInstruction

`func (o *InternationalWirePost) SetWireInstruction(v string)`

SetWireInstruction sets WireInstruction field to given value.

### HasWireInstruction

`func (o *InternationalWirePost) HasWireInstruction() bool`

HasWireInstruction returns a boolean if a field has been set.

### GetWirePurpose

`func (o *InternationalWirePost) GetWirePurpose() string`

GetWirePurpose returns the WirePurpose field if non-nil, zero value otherwise.

### GetWirePurposeOk

`func (o *InternationalWirePost) GetWirePurposeOk() (*string, bool)`

GetWirePurposeOk returns a tuple with the WirePurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirePurpose

`func (o *InternationalWirePost) SetWirePurpose(v string)`

SetWirePurpose sets WirePurpose field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


