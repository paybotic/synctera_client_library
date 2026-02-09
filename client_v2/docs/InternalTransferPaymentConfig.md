# InternalTransferPaymentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAccountId** | **string** | The internal Synctera platform account ID to transfer funds from | 
**Subtype** | Pointer to **string** | Transaction subtype for the internal transfer (default \&quot;autopay_payment\&quot;) | [optional] 

## Methods

### NewInternalTransferPaymentConfig

`func NewInternalTransferPaymentConfig(sourceAccountId string, ) *InternalTransferPaymentConfig`

NewInternalTransferPaymentConfig instantiates a new InternalTransferPaymentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalTransferPaymentConfigWithDefaults

`func NewInternalTransferPaymentConfigWithDefaults() *InternalTransferPaymentConfig`

NewInternalTransferPaymentConfigWithDefaults instantiates a new InternalTransferPaymentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAccountId

`func (o *InternalTransferPaymentConfig) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *InternalTransferPaymentConfig) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *InternalTransferPaymentConfig) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.


### GetSubtype

`func (o *InternalTransferPaymentConfig) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InternalTransferPaymentConfig) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InternalTransferPaymentConfig) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *InternalTransferPaymentConfig) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


