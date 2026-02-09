# CardDisputeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**CardAction**](CardAction.md) |  | 
**Amount** | Pointer to **int64** | Optional amount for the action in cents. If not provided, defaults to the dispute amount. Amount is only applicable when creating actions (state: &#x60;CREATE&#x60;) and is only supported for specific action types (see table below); amounts provided for other action types or states will be ignored.  | Action Type          | Maximum Allowed Amount                  | | -------------------- | --------------------------------------- | | &#x60;CHARGEBACK&#x60;         | Eligible chargeback amount from network | | &#x60;PROVISIONAL_CREDIT&#x60; | Dispute amount                          | | &#x60;WRITE_OFF&#x60;          | Dispute amount                          |  | [optional] 
**Memo** | Pointer to **string** | Memo text related to card dispute action * Required for PRE_ARBITRATION, ARBITRATION, PRE_COMPLIANCE, COMPLIANCE  | [optional] 
**Message** | Pointer to **string** | Message text related to card dispute action  | [optional] 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**ReasonCode** | Pointer to [**CardDisputeReasonCodes**](CardDisputeReasonCodes.md) |  | [optional] 
**State** | [**CardActionState**](CardActionState.md) |  | 
**SupportingDocId** | Pointer to **string** | The unique identifier of the supporting document | [optional] 
**UpdatedReasonCodeMemo** | Pointer to **string** | Memo text describing the reason for updating a reason code. * Required for PRE_ARBITRATION and ARBITRATION if a reason code is provided in the request. If no reason code is present in the request, the original CHARGEBACK reason code will be used.  | [optional] 
**ViolationCode** | Pointer to [**CardActionViolationCode**](CardActionViolationCode.md) |  | [optional] 
**ViolationDate** | Pointer to **time.Time** | The timestamp representing when the violation occurred. * Required for &#x60;PRE-COMPLIANCE&#x60;  | [optional] 

## Methods

### NewCardDisputeAction

`func NewCardDisputeAction(action CardAction, paymentRail PaymentRail, state CardActionState, ) *CardDisputeAction`

NewCardDisputeAction instantiates a new CardDisputeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDisputeActionWithDefaults

`func NewCardDisputeActionWithDefaults() *CardDisputeAction`

NewCardDisputeActionWithDefaults instantiates a new CardDisputeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CardDisputeAction) GetAction() CardAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CardDisputeAction) GetActionOk() (*CardAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CardDisputeAction) SetAction(v CardAction)`

SetAction sets Action field to given value.


### GetAmount

`func (o *CardDisputeAction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardDisputeAction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardDisputeAction) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CardDisputeAction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMemo

`func (o *CardDisputeAction) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CardDisputeAction) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CardDisputeAction) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CardDisputeAction) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMessage

`func (o *CardDisputeAction) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CardDisputeAction) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CardDisputeAction) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CardDisputeAction) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPaymentRail

`func (o *CardDisputeAction) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *CardDisputeAction) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *CardDisputeAction) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetReasonCode

`func (o *CardDisputeAction) GetReasonCode() CardDisputeReasonCodes`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CardDisputeAction) GetReasonCodeOk() (*CardDisputeReasonCodes, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CardDisputeAction) SetReasonCode(v CardDisputeReasonCodes)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *CardDisputeAction) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetState

`func (o *CardDisputeAction) GetState() CardActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardDisputeAction) GetStateOk() (*CardActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardDisputeAction) SetState(v CardActionState)`

SetState sets State field to given value.


### GetSupportingDocId

`func (o *CardDisputeAction) GetSupportingDocId() string`

GetSupportingDocId returns the SupportingDocId field if non-nil, zero value otherwise.

### GetSupportingDocIdOk

`func (o *CardDisputeAction) GetSupportingDocIdOk() (*string, bool)`

GetSupportingDocIdOk returns a tuple with the SupportingDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingDocId

`func (o *CardDisputeAction) SetSupportingDocId(v string)`

SetSupportingDocId sets SupportingDocId field to given value.

### HasSupportingDocId

`func (o *CardDisputeAction) HasSupportingDocId() bool`

HasSupportingDocId returns a boolean if a field has been set.

### GetUpdatedReasonCodeMemo

`func (o *CardDisputeAction) GetUpdatedReasonCodeMemo() string`

GetUpdatedReasonCodeMemo returns the UpdatedReasonCodeMemo field if non-nil, zero value otherwise.

### GetUpdatedReasonCodeMemoOk

`func (o *CardDisputeAction) GetUpdatedReasonCodeMemoOk() (*string, bool)`

GetUpdatedReasonCodeMemoOk returns a tuple with the UpdatedReasonCodeMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReasonCodeMemo

`func (o *CardDisputeAction) SetUpdatedReasonCodeMemo(v string)`

SetUpdatedReasonCodeMemo sets UpdatedReasonCodeMemo field to given value.

### HasUpdatedReasonCodeMemo

`func (o *CardDisputeAction) HasUpdatedReasonCodeMemo() bool`

HasUpdatedReasonCodeMemo returns a boolean if a field has been set.

### GetViolationCode

`func (o *CardDisputeAction) GetViolationCode() CardActionViolationCode`

GetViolationCode returns the ViolationCode field if non-nil, zero value otherwise.

### GetViolationCodeOk

`func (o *CardDisputeAction) GetViolationCodeOk() (*CardActionViolationCode, bool)`

GetViolationCodeOk returns a tuple with the ViolationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCode

`func (o *CardDisputeAction) SetViolationCode(v CardActionViolationCode)`

SetViolationCode sets ViolationCode field to given value.

### HasViolationCode

`func (o *CardDisputeAction) HasViolationCode() bool`

HasViolationCode returns a boolean if a field has been set.

### GetViolationDate

`func (o *CardDisputeAction) GetViolationDate() time.Time`

GetViolationDate returns the ViolationDate field if non-nil, zero value otherwise.

### GetViolationDateOk

`func (o *CardDisputeAction) GetViolationDateOk() (*time.Time, bool)`

GetViolationDateOk returns a tuple with the ViolationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationDate

`func (o *CardDisputeAction) SetViolationDate(v time.Time)`

SetViolationDate sets ViolationDate field to given value.

### HasViolationDate

`func (o *CardDisputeAction) HasViolationDate() bool`

HasViolationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


