# AutopayAttemptHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | [**[]AutopayAttempt**](AutopayAttempt.md) | List of execution attempts, ordered chronologically | 

## Methods

### NewAutopayAttemptHistory

`func NewAutopayAttemptHistory(attempts []AutopayAttempt, ) *AutopayAttemptHistory`

NewAutopayAttemptHistory instantiates a new AutopayAttemptHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopayAttemptHistoryWithDefaults

`func NewAutopayAttemptHistoryWithDefaults() *AutopayAttemptHistory`

NewAutopayAttemptHistoryWithDefaults instantiates a new AutopayAttemptHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *AutopayAttemptHistory) GetAttempts() []AutopayAttempt`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *AutopayAttemptHistory) GetAttemptsOk() (*[]AutopayAttempt, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *AutopayAttemptHistory) SetAttempts(v []AutopayAttempt)`

SetAttempts sets Attempts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


