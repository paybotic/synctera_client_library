# AutopayConfigData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountRule** | [**AutopayAmountRule**](AutopayAmountRule.md) |  | 
**DescriptionTemplate** | Pointer to **string** | Go template for payment description. Used for ACH addenda and internal transfer memo. Supports template variables: - {{.Last4AccountId}}: Last 4 characters of the lending account UUID - {{.BillingPeriodEndDate}}: End date of the billing period (YYYY-MM-DD) - {{.PaymentAmount}}: Payment amount formatted as currency (e.g., \&quot;$50.00\&quot;) Default: \&quot;Autopay\&quot;  | [optional] 
**FailurePolicy** | [**AutopayFailurePolicy**](AutopayFailurePolicy.md) |  | 
**PaymentConfigs** | Pointer to [**AutopayPaymentConfigs**](AutopayPaymentConfigs.md) |  | [optional] 
**PaymentMethod** | [**AutopayPaymentMethod**](AutopayPaymentMethod.md) |  | 
**RuleConfigs** | Pointer to [**AutopayRuleConfigs**](AutopayRuleConfigs.md) |  | [optional] 
**TimingRule** | [**AutopayTimingRule**](AutopayTimingRule.md) |  | 

## Methods

### NewAutopayConfigData

`func NewAutopayConfigData(amountRule AutopayAmountRule, failurePolicy AutopayFailurePolicy, paymentMethod AutopayPaymentMethod, timingRule AutopayTimingRule, ) *AutopayConfigData`

NewAutopayConfigData instantiates a new AutopayConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayConfigDataWithDefaults

`func NewAutopayConfigDataWithDefaults() *AutopayConfigData`

NewAutopayConfigDataWithDefaults instantiates a new AutopayConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountRule

`func (o *AutopayConfigData) GetAmountRule() AutopayAmountRule`

GetAmountRule returns the AmountRule field if non-nil, zero value otherwise.

### GetAmountRuleOk

`func (o *AutopayConfigData) GetAmountRuleOk() (*AutopayAmountRule, bool)`

GetAmountRuleOk returns a tuple with the AmountRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRule

`func (o *AutopayConfigData) SetAmountRule(v AutopayAmountRule)`

SetAmountRule sets AmountRule field to given value.


### GetDescriptionTemplate

`func (o *AutopayConfigData) GetDescriptionTemplate() string`

GetDescriptionTemplate returns the DescriptionTemplate field if non-nil, zero value otherwise.

### GetDescriptionTemplateOk

`func (o *AutopayConfigData) GetDescriptionTemplateOk() (*string, bool)`

GetDescriptionTemplateOk returns a tuple with the DescriptionTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionTemplate

`func (o *AutopayConfigData) SetDescriptionTemplate(v string)`

SetDescriptionTemplate sets DescriptionTemplate field to given value.

### HasDescriptionTemplate

`func (o *AutopayConfigData) HasDescriptionTemplate() bool`

HasDescriptionTemplate returns a boolean if a field has been set.

### GetFailurePolicy

`func (o *AutopayConfigData) GetFailurePolicy() AutopayFailurePolicy`

GetFailurePolicy returns the FailurePolicy field if non-nil, zero value otherwise.

### GetFailurePolicyOk

`func (o *AutopayConfigData) GetFailurePolicyOk() (*AutopayFailurePolicy, bool)`

GetFailurePolicyOk returns a tuple with the FailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurePolicy

`func (o *AutopayConfigData) SetFailurePolicy(v AutopayFailurePolicy)`

SetFailurePolicy sets FailurePolicy field to given value.


### GetPaymentConfigs

`func (o *AutopayConfigData) GetPaymentConfigs() AutopayPaymentConfigs`

GetPaymentConfigs returns the PaymentConfigs field if non-nil, zero value otherwise.

### GetPaymentConfigsOk

`func (o *AutopayConfigData) GetPaymentConfigsOk() (*AutopayPaymentConfigs, bool)`

GetPaymentConfigsOk returns a tuple with the PaymentConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfigs

`func (o *AutopayConfigData) SetPaymentConfigs(v AutopayPaymentConfigs)`

SetPaymentConfigs sets PaymentConfigs field to given value.

### HasPaymentConfigs

`func (o *AutopayConfigData) HasPaymentConfigs() bool`

HasPaymentConfigs returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *AutopayConfigData) GetPaymentMethod() AutopayPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *AutopayConfigData) GetPaymentMethodOk() (*AutopayPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *AutopayConfigData) SetPaymentMethod(v AutopayPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetRuleConfigs

`func (o *AutopayConfigData) GetRuleConfigs() AutopayRuleConfigs`

GetRuleConfigs returns the RuleConfigs field if non-nil, zero value otherwise.

### GetRuleConfigsOk

`func (o *AutopayConfigData) GetRuleConfigsOk() (*AutopayRuleConfigs, bool)`

GetRuleConfigsOk returns a tuple with the RuleConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleConfigs

`func (o *AutopayConfigData) SetRuleConfigs(v AutopayRuleConfigs)`

SetRuleConfigs sets RuleConfigs field to given value.

### HasRuleConfigs

`func (o *AutopayConfigData) HasRuleConfigs() bool`

HasRuleConfigs returns a boolean if a field has been set.

### GetTimingRule

`func (o *AutopayConfigData) GetTimingRule() AutopayTimingRule`

GetTimingRule returns the TimingRule field if non-nil, zero value otherwise.

### GetTimingRuleOk

`func (o *AutopayConfigData) GetTimingRuleOk() (*AutopayTimingRule, bool)`

GetTimingRuleOk returns a tuple with the TimingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimingRule

`func (o *AutopayConfigData) SetTimingRule(v AutopayTimingRule)`

SetTimingRule sets TimingRule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


