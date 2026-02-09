# AutopayRuleConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentBalance** | Pointer to **map[string]interface{}** | Configuration for current balance amount rule | [optional] 
**DaysBeforeDue** | Pointer to [**DaysBeforeDueConfig**](DaysBeforeDueConfig.md) |  | [optional] 
**FixedAmount** | Pointer to [**FixedAmountConfig**](FixedAmountConfig.md) |  | [optional] 
**MinimumDue** | Pointer to **map[string]interface{}** | Configuration for minimum due amount rule | [optional] 
**StatementBalance** | Pointer to **map[string]interface{}** | Configuration for statement balance amount rule | [optional] 

## Methods

### NewAutopayRuleConfigs

`func NewAutopayRuleConfigs() *AutopayRuleConfigs`

NewAutopayRuleConfigs instantiates a new AutopayRuleConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayRuleConfigsWithDefaults

`func NewAutopayRuleConfigsWithDefaults() *AutopayRuleConfigs`

NewAutopayRuleConfigsWithDefaults instantiates a new AutopayRuleConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentBalance

`func (o *AutopayRuleConfigs) GetCurrentBalance() map[string]interface{}`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *AutopayRuleConfigs) GetCurrentBalanceOk() (*map[string]interface{}, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *AutopayRuleConfigs) SetCurrentBalance(v map[string]interface{})`

SetCurrentBalance sets CurrentBalance field to given value.

### HasCurrentBalance

`func (o *AutopayRuleConfigs) HasCurrentBalance() bool`

HasCurrentBalance returns a boolean if a field has been set.

### GetDaysBeforeDue

`func (o *AutopayRuleConfigs) GetDaysBeforeDue() DaysBeforeDueConfig`

GetDaysBeforeDue returns the DaysBeforeDue field if non-nil, zero value otherwise.

### GetDaysBeforeDueOk

`func (o *AutopayRuleConfigs) GetDaysBeforeDueOk() (*DaysBeforeDueConfig, bool)`

GetDaysBeforeDueOk returns a tuple with the DaysBeforeDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysBeforeDue

`func (o *AutopayRuleConfigs) SetDaysBeforeDue(v DaysBeforeDueConfig)`

SetDaysBeforeDue sets DaysBeforeDue field to given value.

### HasDaysBeforeDue

`func (o *AutopayRuleConfigs) HasDaysBeforeDue() bool`

HasDaysBeforeDue returns a boolean if a field has been set.

### GetFixedAmount

`func (o *AutopayRuleConfigs) GetFixedAmount() FixedAmountConfig`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *AutopayRuleConfigs) GetFixedAmountOk() (*FixedAmountConfig, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *AutopayRuleConfigs) SetFixedAmount(v FixedAmountConfig)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *AutopayRuleConfigs) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetMinimumDue

`func (o *AutopayRuleConfigs) GetMinimumDue() map[string]interface{}`

GetMinimumDue returns the MinimumDue field if non-nil, zero value otherwise.

### GetMinimumDueOk

`func (o *AutopayRuleConfigs) GetMinimumDueOk() (*map[string]interface{}, bool)`

GetMinimumDueOk returns a tuple with the MinimumDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumDue

`func (o *AutopayRuleConfigs) SetMinimumDue(v map[string]interface{})`

SetMinimumDue sets MinimumDue field to given value.

### HasMinimumDue

`func (o *AutopayRuleConfigs) HasMinimumDue() bool`

HasMinimumDue returns a boolean if a field has been set.

### GetStatementBalance

`func (o *AutopayRuleConfigs) GetStatementBalance() map[string]interface{}`

GetStatementBalance returns the StatementBalance field if non-nil, zero value otherwise.

### GetStatementBalanceOk

`func (o *AutopayRuleConfigs) GetStatementBalanceOk() (*map[string]interface{}, bool)`

GetStatementBalanceOk returns a tuple with the StatementBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementBalance

`func (o *AutopayRuleConfigs) SetStatementBalance(v map[string]interface{})`

SetStatementBalance sets StatementBalance field to given value.

### HasStatementBalance

`func (o *AutopayRuleConfigs) HasStatementBalance() bool`

HasStatementBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


