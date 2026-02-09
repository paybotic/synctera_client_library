# CardDisputeActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**CardAction**](CardAction.md) |  | 
**Amount** | Pointer to **int64** | The amount of the action in cents. | [optional] 
**CreationTime** | **time.Time** | The timestamp representing when the object was created | [readonly] 
**ExternalReferenceId** | Pointer to **string** | Reference ID associated with the action on the external network. | [optional] 
**Id** | **string** | The unique identifier of the dispute action | [readonly] 
**Memo** | Pointer to **string** | Memo text related to card dispute action  | [optional] 
**Message** | Pointer to **string** | Message text related to card dispute action | [optional] 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**ReasonCode** | Pointer to [**CardDisputeReasonCodes**](CardDisputeReasonCodes.md) |  | [optional] 
**RepresentmentReasonCode** | Pointer to [**CardDisputeRepresentmentReasonCode**](CardDisputeRepresentmentReasonCode.md) |  | [optional] 
**Status** | [**CardDisputeActionStatus**](CardDisputeActionStatus.md) |  | 
**SupportingDocId** | Pointer to **string** | The unique identifier of the supporting document | [optional] [readonly] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**UpdatedReasonCodeMemo** | Pointer to **string** | Memo text describing the reason for updating a reason code.  | [optional] 
**ViolationCode** | Pointer to [**CardActionViolationCode**](CardActionViolationCode.md) |  | [optional] 
**ViolationDate** | Pointer to **time.Time** | The timestamp representing when the violation occurred.  | [optional] 

## Methods

### NewCardDisputeActionResponse

`func NewCardDisputeActionResponse(action CardAction, creationTime time.Time, id string, paymentRail PaymentRail, status CardDisputeActionStatus, tenant string, ) *CardDisputeActionResponse`

NewCardDisputeActionResponse instantiates a new CardDisputeActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDisputeActionResponseWithDefaults

`func NewCardDisputeActionResponseWithDefaults() *CardDisputeActionResponse`

NewCardDisputeActionResponseWithDefaults instantiates a new CardDisputeActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CardDisputeActionResponse) GetAction() CardAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CardDisputeActionResponse) GetActionOk() (*CardAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CardDisputeActionResponse) SetAction(v CardAction)`

SetAction sets Action field to given value.


### GetAmount

`func (o *CardDisputeActionResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CardDisputeActionResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CardDisputeActionResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CardDisputeActionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreationTime

`func (o *CardDisputeActionResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardDisputeActionResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardDisputeActionResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetExternalReferenceId

`func (o *CardDisputeActionResponse) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *CardDisputeActionResponse) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *CardDisputeActionResponse) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.

### HasExternalReferenceId

`func (o *CardDisputeActionResponse) HasExternalReferenceId() bool`

HasExternalReferenceId returns a boolean if a field has been set.

### GetId

`func (o *CardDisputeActionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardDisputeActionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardDisputeActionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMemo

`func (o *CardDisputeActionResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CardDisputeActionResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CardDisputeActionResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CardDisputeActionResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMessage

`func (o *CardDisputeActionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CardDisputeActionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CardDisputeActionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CardDisputeActionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPaymentRail

`func (o *CardDisputeActionResponse) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *CardDisputeActionResponse) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *CardDisputeActionResponse) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetReasonCode

`func (o *CardDisputeActionResponse) GetReasonCode() CardDisputeReasonCodes`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CardDisputeActionResponse) GetReasonCodeOk() (*CardDisputeReasonCodes, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CardDisputeActionResponse) SetReasonCode(v CardDisputeReasonCodes)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *CardDisputeActionResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetRepresentmentReasonCode

`func (o *CardDisputeActionResponse) GetRepresentmentReasonCode() CardDisputeRepresentmentReasonCode`

GetRepresentmentReasonCode returns the RepresentmentReasonCode field if non-nil, zero value otherwise.

### GetRepresentmentReasonCodeOk

`func (o *CardDisputeActionResponse) GetRepresentmentReasonCodeOk() (*CardDisputeRepresentmentReasonCode, bool)`

GetRepresentmentReasonCodeOk returns a tuple with the RepresentmentReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentmentReasonCode

`func (o *CardDisputeActionResponse) SetRepresentmentReasonCode(v CardDisputeRepresentmentReasonCode)`

SetRepresentmentReasonCode sets RepresentmentReasonCode field to given value.

### HasRepresentmentReasonCode

`func (o *CardDisputeActionResponse) HasRepresentmentReasonCode() bool`

HasRepresentmentReasonCode returns a boolean if a field has been set.

### GetStatus

`func (o *CardDisputeActionResponse) GetStatus() CardDisputeActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardDisputeActionResponse) GetStatusOk() (*CardDisputeActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardDisputeActionResponse) SetStatus(v CardDisputeActionStatus)`

SetStatus sets Status field to given value.


### GetSupportingDocId

`func (o *CardDisputeActionResponse) GetSupportingDocId() string`

GetSupportingDocId returns the SupportingDocId field if non-nil, zero value otherwise.

### GetSupportingDocIdOk

`func (o *CardDisputeActionResponse) GetSupportingDocIdOk() (*string, bool)`

GetSupportingDocIdOk returns a tuple with the SupportingDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingDocId

`func (o *CardDisputeActionResponse) SetSupportingDocId(v string)`

SetSupportingDocId sets SupportingDocId field to given value.

### HasSupportingDocId

`func (o *CardDisputeActionResponse) HasSupportingDocId() bool`

HasSupportingDocId returns a boolean if a field has been set.

### GetTenant

`func (o *CardDisputeActionResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CardDisputeActionResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CardDisputeActionResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetUpdatedReasonCodeMemo

`func (o *CardDisputeActionResponse) GetUpdatedReasonCodeMemo() string`

GetUpdatedReasonCodeMemo returns the UpdatedReasonCodeMemo field if non-nil, zero value otherwise.

### GetUpdatedReasonCodeMemoOk

`func (o *CardDisputeActionResponse) GetUpdatedReasonCodeMemoOk() (*string, bool)`

GetUpdatedReasonCodeMemoOk returns a tuple with the UpdatedReasonCodeMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReasonCodeMemo

`func (o *CardDisputeActionResponse) SetUpdatedReasonCodeMemo(v string)`

SetUpdatedReasonCodeMemo sets UpdatedReasonCodeMemo field to given value.

### HasUpdatedReasonCodeMemo

`func (o *CardDisputeActionResponse) HasUpdatedReasonCodeMemo() bool`

HasUpdatedReasonCodeMemo returns a boolean if a field has been set.

### GetViolationCode

`func (o *CardDisputeActionResponse) GetViolationCode() CardActionViolationCode`

GetViolationCode returns the ViolationCode field if non-nil, zero value otherwise.

### GetViolationCodeOk

`func (o *CardDisputeActionResponse) GetViolationCodeOk() (*CardActionViolationCode, bool)`

GetViolationCodeOk returns a tuple with the ViolationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCode

`func (o *CardDisputeActionResponse) SetViolationCode(v CardActionViolationCode)`

SetViolationCode sets ViolationCode field to given value.

### HasViolationCode

`func (o *CardDisputeActionResponse) HasViolationCode() bool`

HasViolationCode returns a boolean if a field has been set.

### GetViolationDate

`func (o *CardDisputeActionResponse) GetViolationDate() time.Time`

GetViolationDate returns the ViolationDate field if non-nil, zero value otherwise.

### GetViolationDateOk

`func (o *CardDisputeActionResponse) GetViolationDateOk() (*time.Time, bool)`

GetViolationDateOk returns a tuple with the ViolationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationDate

`func (o *CardDisputeActionResponse) SetViolationDate(v time.Time)`

SetViolationDate sets ViolationDate field to given value.

### HasViolationDate

`func (o *CardDisputeActionResponse) HasViolationDate() bool`

HasViolationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


