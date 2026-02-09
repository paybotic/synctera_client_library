# BatchPaymentTemplateRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableFourEyesReview** | Pointer to **bool** | Whether or not the batch transfer will require four eyes review.  | [optional] 
**HoldTime** | Pointer to **int32** | The number of minutes that the batch transfer will be held before being processed.  | [optional] 
**PaymentRail** | Pointer to **string** | The payment rail that will be used to process the batch transfer.  | [optional] 

## Methods

### NewBatchPaymentTemplateRules

`func NewBatchPaymentTemplateRules() *BatchPaymentTemplateRules`

NewBatchPaymentTemplateRules instantiates a new BatchPaymentTemplateRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchPaymentTemplateRulesWithDefaults

`func NewBatchPaymentTemplateRulesWithDefaults() *BatchPaymentTemplateRules`

NewBatchPaymentTemplateRulesWithDefaults instantiates a new BatchPaymentTemplateRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableFourEyesReview

`func (o *BatchPaymentTemplateRules) GetDisableFourEyesReview() bool`

GetDisableFourEyesReview returns the DisableFourEyesReview field if non-nil, zero value otherwise.

### GetDisableFourEyesReviewOk

`func (o *BatchPaymentTemplateRules) GetDisableFourEyesReviewOk() (*bool, bool)`

GetDisableFourEyesReviewOk returns a tuple with the DisableFourEyesReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFourEyesReview

`func (o *BatchPaymentTemplateRules) SetDisableFourEyesReview(v bool)`

SetDisableFourEyesReview sets DisableFourEyesReview field to given value.

### HasDisableFourEyesReview

`func (o *BatchPaymentTemplateRules) HasDisableFourEyesReview() bool`

HasDisableFourEyesReview returns a boolean if a field has been set.

### GetHoldTime

`func (o *BatchPaymentTemplateRules) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *BatchPaymentTemplateRules) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *BatchPaymentTemplateRules) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *BatchPaymentTemplateRules) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### GetPaymentRail

`func (o *BatchPaymentTemplateRules) GetPaymentRail() string`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *BatchPaymentTemplateRules) GetPaymentRailOk() (*string, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *BatchPaymentTemplateRules) SetPaymentRail(v string)`

SetPaymentRail sets PaymentRail field to given value.

### HasPaymentRail

`func (o *BatchPaymentTemplateRules) HasPaymentRail() bool`

HasPaymentRail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


