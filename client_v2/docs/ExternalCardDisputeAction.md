# ExternalCardDisputeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ExternalCardAction**](ExternalCardAction.md) |  | 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**State** | [**ExternalCardActionState**](ExternalCardActionState.md) |  | 
**SupportingDocId** | Pointer to **string** | The unique identifier of the supporting document * Required for &#x60;REPRESENTMENT.CREATE&#x60;  | [optional] 

## Methods

### NewExternalCardDisputeAction

`func NewExternalCardDisputeAction(action ExternalCardAction, paymentRail PaymentRail, state ExternalCardActionState, ) *ExternalCardDisputeAction`

NewExternalCardDisputeAction instantiates a new ExternalCardDisputeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardDisputeActionWithDefaults

`func NewExternalCardDisputeActionWithDefaults() *ExternalCardDisputeAction`

NewExternalCardDisputeActionWithDefaults instantiates a new ExternalCardDisputeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ExternalCardDisputeAction) GetAction() ExternalCardAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ExternalCardDisputeAction) GetActionOk() (*ExternalCardAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ExternalCardDisputeAction) SetAction(v ExternalCardAction)`

SetAction sets Action field to given value.


### GetPaymentRail

`func (o *ExternalCardDisputeAction) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *ExternalCardDisputeAction) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *ExternalCardDisputeAction) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetState

`func (o *ExternalCardDisputeAction) GetState() ExternalCardActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExternalCardDisputeAction) GetStateOk() (*ExternalCardActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExternalCardDisputeAction) SetState(v ExternalCardActionState)`

SetState sets State field to given value.


### GetSupportingDocId

`func (o *ExternalCardDisputeAction) GetSupportingDocId() string`

GetSupportingDocId returns the SupportingDocId field if non-nil, zero value otherwise.

### GetSupportingDocIdOk

`func (o *ExternalCardDisputeAction) GetSupportingDocIdOk() (*string, bool)`

GetSupportingDocIdOk returns a tuple with the SupportingDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingDocId

`func (o *ExternalCardDisputeAction) SetSupportingDocId(v string)`

SetSupportingDocId sets SupportingDocId field to given value.

### HasSupportingDocId

`func (o *ExternalCardDisputeAction) HasSupportingDocId() bool`

HasSupportingDocId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


