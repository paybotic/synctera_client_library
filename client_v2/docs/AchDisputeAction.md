# AchDisputeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**AchAction**](AchAction.md) |  | 
**Message** | Pointer to **string** | Message text related to ACH dispute action * Max length for debit transactions is 38 characters  | [optional] 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**ReasonCode** | Pointer to [**AchDisputeReasonCodes**](AchDisputeReasonCodes.md) |  | [optional] 
**ReturnCode** | Pointer to [**AchDisputeReturnCodes**](AchDisputeReturnCodes.md) |  | [optional] 
**State** | [**AchActionState**](AchActionState.md) |  | 
**SupportingDocId** | Pointer to **string** | The unique identifier of the supporting document | [optional] 

## Methods

### NewAchDisputeAction

`func NewAchDisputeAction(action AchAction, paymentRail PaymentRail, state AchActionState, ) *AchDisputeAction`

NewAchDisputeAction instantiates a new AchDisputeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchDisputeActionWithDefaults

`func NewAchDisputeActionWithDefaults() *AchDisputeAction`

NewAchDisputeActionWithDefaults instantiates a new AchDisputeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AchDisputeAction) GetAction() AchAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AchDisputeAction) GetActionOk() (*AchAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AchDisputeAction) SetAction(v AchAction)`

SetAction sets Action field to given value.


### GetMessage

`func (o *AchDisputeAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AchDisputeAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AchDisputeAction) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AchDisputeAction) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPaymentRail

`func (o *AchDisputeAction) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *AchDisputeAction) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *AchDisputeAction) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetReasonCode

`func (o *AchDisputeAction) GetReasonCode() AchDisputeReasonCodes`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *AchDisputeAction) GetReasonCodeOk() (*AchDisputeReasonCodes, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *AchDisputeAction) SetReasonCode(v AchDisputeReasonCodes)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *AchDisputeAction) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReturnCode

`func (o *AchDisputeAction) GetReturnCode() AchDisputeReturnCodes`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *AchDisputeAction) GetReturnCodeOk() (*AchDisputeReturnCodes, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *AchDisputeAction) SetReturnCode(v AchDisputeReturnCodes)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *AchDisputeAction) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetState

`func (o *AchDisputeAction) GetState() AchActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AchDisputeAction) GetStateOk() (*AchActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AchDisputeAction) SetState(v AchActionState)`

SetState sets State field to given value.


### GetSupportingDocId

`func (o *AchDisputeAction) GetSupportingDocId() string`

GetSupportingDocId returns the SupportingDocId field if non-nil, zero value otherwise.

### GetSupportingDocIdOk

`func (o *AchDisputeAction) GetSupportingDocIdOk() (*string, bool)`

GetSupportingDocIdOk returns a tuple with the SupportingDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingDocId

`func (o *AchDisputeAction) SetSupportingDocId(v string)`

SetSupportingDocId sets SupportingDocId field to given value.

### HasSupportingDocId

`func (o *AchDisputeAction) HasSupportingDocId() bool`

HasSupportingDocId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


