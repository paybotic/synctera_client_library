# BaseTransactionDecline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeclineDetails** | **string** | Additional detail about the decline. | 
**Reason** | **string** | The reason the transaction was declined | 

## Methods

### NewBaseTransactionDecline

`func NewBaseTransactionDecline(declineDetails string, reason string, ) *BaseTransactionDecline`

NewBaseTransactionDecline instantiates a new BaseTransactionDecline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTransactionDeclineWithDefaults

`func NewBaseTransactionDeclineWithDefaults() *BaseTransactionDecline`

NewBaseTransactionDeclineWithDefaults instantiates a new BaseTransactionDecline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeclineDetails

`func (o *BaseTransactionDecline) GetDeclineDetails() string`

GetDeclineDetails returns the DeclineDetails field if non-nil, zero value otherwise.

### GetDeclineDetailsOk

`func (o *BaseTransactionDecline) GetDeclineDetailsOk() (*string, bool)`

GetDeclineDetailsOk returns a tuple with the DeclineDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineDetails

`func (o *BaseTransactionDecline) SetDeclineDetails(v string)`

SetDeclineDetails sets DeclineDetails field to given value.


### GetReason

`func (o *BaseTransactionDecline) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BaseTransactionDecline) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BaseTransactionDecline) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


