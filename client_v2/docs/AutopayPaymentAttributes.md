# AutopayPaymentAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchId** | Pointer to **string** | ACH transaction ID when payment method is \&quot;ach\&quot; | [optional] 
**TransactionId** | Pointer to **string** | Internal transaction UUID when payment method is \&quot;internal_transfer\&quot; | [optional] 

## Methods

### NewAutopayPaymentAttributes

`func NewAutopayPaymentAttributes() *AutopayPaymentAttributes`

NewAutopayPaymentAttributes instantiates a new AutopayPaymentAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayPaymentAttributesWithDefaults

`func NewAutopayPaymentAttributesWithDefaults() *AutopayPaymentAttributes`

NewAutopayPaymentAttributesWithDefaults instantiates a new AutopayPaymentAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchId

`func (o *AutopayPaymentAttributes) GetAchId() string`

GetAchId returns the AchId field if non-nil, zero value otherwise.

### GetAchIdOk

`func (o *AutopayPaymentAttributes) GetAchIdOk() (*string, bool)`

GetAchIdOk returns a tuple with the AchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchId

`func (o *AutopayPaymentAttributes) SetAchId(v string)`

SetAchId sets AchId field to given value.

### HasAchId

`func (o *AutopayPaymentAttributes) HasAchId() bool`

HasAchId returns a boolean if a field has been set.

### GetTransactionId

`func (o *AutopayPaymentAttributes) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *AutopayPaymentAttributes) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *AutopayPaymentAttributes) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *AutopayPaymentAttributes) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


