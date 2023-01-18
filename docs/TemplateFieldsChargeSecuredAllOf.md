# TemplateFieldsChargeSecuredAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeoffPeriod** | Pointer to **int32** | The number of days an account can stay delinquent before marking an account as charged-off.  | [optional] [default to 90]
**DelinquencyPeriod** | Pointer to **int32** | The number of days past the due date to wait for a minimum payment before marking an account as delinquent.  | [optional] [default to 30]
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment.  | [optional] [default to 30]
**InterestProductId** | **string** | An interest account product that the current account associates with. The account product must have its calculation_method set to COMPOUNDED_DAILY, and its rates set to 0%.  | 
**MinimumPayment** | [**MinimumPaymentFull**](MinimumPaymentFull.md) |  | 

## Methods

### NewTemplateFieldsChargeSecuredAllOf

`func NewTemplateFieldsChargeSecuredAllOf(interestProductId string, minimumPayment MinimumPaymentFull, ) *TemplateFieldsChargeSecuredAllOf`

NewTemplateFieldsChargeSecuredAllOf instantiates a new TemplateFieldsChargeSecuredAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldsChargeSecuredAllOfWithDefaults

`func NewTemplateFieldsChargeSecuredAllOfWithDefaults() *TemplateFieldsChargeSecuredAllOf`

NewTemplateFieldsChargeSecuredAllOfWithDefaults instantiates a new TemplateFieldsChargeSecuredAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeoffPeriod

`func (o *TemplateFieldsChargeSecuredAllOf) GetChargeoffPeriod() int32`

GetChargeoffPeriod returns the ChargeoffPeriod field if non-nil, zero value otherwise.

### GetChargeoffPeriodOk

`func (o *TemplateFieldsChargeSecuredAllOf) GetChargeoffPeriodOk() (*int32, bool)`

GetChargeoffPeriodOk returns a tuple with the ChargeoffPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeoffPeriod

`func (o *TemplateFieldsChargeSecuredAllOf) SetChargeoffPeriod(v int32)`

SetChargeoffPeriod sets ChargeoffPeriod field to given value.

### HasChargeoffPeriod

`func (o *TemplateFieldsChargeSecuredAllOf) HasChargeoffPeriod() bool`

HasChargeoffPeriod returns a boolean if a field has been set.

### GetDelinquencyPeriod

`func (o *TemplateFieldsChargeSecuredAllOf) GetDelinquencyPeriod() int32`

GetDelinquencyPeriod returns the DelinquencyPeriod field if non-nil, zero value otherwise.

### GetDelinquencyPeriodOk

`func (o *TemplateFieldsChargeSecuredAllOf) GetDelinquencyPeriodOk() (*int32, bool)`

GetDelinquencyPeriodOk returns a tuple with the DelinquencyPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelinquencyPeriod

`func (o *TemplateFieldsChargeSecuredAllOf) SetDelinquencyPeriod(v int32)`

SetDelinquencyPeriod sets DelinquencyPeriod field to given value.

### HasDelinquencyPeriod

`func (o *TemplateFieldsChargeSecuredAllOf) HasDelinquencyPeriod() bool`

HasDelinquencyPeriod returns a boolean if a field has been set.

### GetGracePeriod

`func (o *TemplateFieldsChargeSecuredAllOf) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *TemplateFieldsChargeSecuredAllOf) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *TemplateFieldsChargeSecuredAllOf) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *TemplateFieldsChargeSecuredAllOf) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetInterestProductId

`func (o *TemplateFieldsChargeSecuredAllOf) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *TemplateFieldsChargeSecuredAllOf) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *TemplateFieldsChargeSecuredAllOf) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.


### GetMinimumPayment

`func (o *TemplateFieldsChargeSecuredAllOf) GetMinimumPayment() MinimumPaymentFull`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *TemplateFieldsChargeSecuredAllOf) GetMinimumPaymentOk() (*MinimumPaymentFull, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *TemplateFieldsChargeSecuredAllOf) SetMinimumPayment(v MinimumPaymentFull)`

SetMinimumPayment sets MinimumPayment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


