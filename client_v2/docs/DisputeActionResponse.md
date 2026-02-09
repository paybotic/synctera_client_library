# DisputeActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ExternalCardAction**](ExternalCardAction.md) |  | 
**Amount** | Pointer to **int64** | The amount of the action in cents. | [optional] 
**CreationTime** | **time.Time** | The timestamp representing when the object was created | [readonly] 
**ExternalReferenceId** | Pointer to **string** | Reference ID associated with the action on the external network. | [optional] 
**Id** | **string** | The unique identifier of the dispute action | [readonly] 
**Memo** | Pointer to **string** | Memo text related to card dispute action  | [optional] 
**Message** | Pointer to **string** | Message text related to card dispute action | [optional] 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**ReasonCode** | Pointer to [**AchDisputeReasonCodes**](AchDisputeReasonCodes.md) |  | [optional] 
**RepresentmentReasonCode** | Pointer to [**CardDisputeRepresentmentReasonCode**](CardDisputeRepresentmentReasonCode.md) |  | [optional] 
**Status** | [**ExternalCardDisputeActionStatus**](ExternalCardDisputeActionStatus.md) |  | 
**SupportingDocId** | Pointer to **string** | The unique identifier of the supporting document | [optional] [readonly] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**UpdatedReasonCodeMemo** | Pointer to **string** | Memo text describing the reason for updating a reason code.  | [optional] 
**ViolationCode** | Pointer to [**CardActionViolationCode**](CardActionViolationCode.md) |  | [optional] 
**ViolationDate** | Pointer to **time.Time** | The timestamp representing when the violation occurred.  | [optional] 
**ReturnCode** | Pointer to [**AchDisputeReturnCodes**](AchDisputeReturnCodes.md) |  | [optional] 

## Methods

### NewDisputeActionResponse

`func NewDisputeActionResponse(action ExternalCardAction, creationTime time.Time, id string, paymentRail PaymentRail, status ExternalCardDisputeActionStatus, tenant string, ) *DisputeActionResponse`

NewDisputeActionResponse instantiates a new DisputeActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeActionResponseWithDefaults

`func NewDisputeActionResponseWithDefaults() *DisputeActionResponse`

NewDisputeActionResponseWithDefaults instantiates a new DisputeActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DisputeActionResponse) GetAction() ExternalCardAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DisputeActionResponse) GetActionOk() (*ExternalCardAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DisputeActionResponse) SetAction(v ExternalCardAction)`

SetAction sets Action field to given value.


### GetAmount

`func (o *DisputeActionResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DisputeActionResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DisputeActionResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DisputeActionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreationTime

`func (o *DisputeActionResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DisputeActionResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DisputeActionResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetExternalReferenceId

`func (o *DisputeActionResponse) GetExternalReferenceId() string`

GetExternalReferenceId returns the ExternalReferenceId field if non-nil, zero value otherwise.

### GetExternalReferenceIdOk

`func (o *DisputeActionResponse) GetExternalReferenceIdOk() (*string, bool)`

GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferenceId

`func (o *DisputeActionResponse) SetExternalReferenceId(v string)`

SetExternalReferenceId sets ExternalReferenceId field to given value.

### HasExternalReferenceId

`func (o *DisputeActionResponse) HasExternalReferenceId() bool`

HasExternalReferenceId returns a boolean if a field has been set.

### GetId

`func (o *DisputeActionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisputeActionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisputeActionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMemo

`func (o *DisputeActionResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *DisputeActionResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *DisputeActionResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *DisputeActionResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMessage

`func (o *DisputeActionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DisputeActionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DisputeActionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DisputeActionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPaymentRail

`func (o *DisputeActionResponse) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *DisputeActionResponse) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *DisputeActionResponse) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetReasonCode

`func (o *DisputeActionResponse) GetReasonCode() AchDisputeReasonCodes`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *DisputeActionResponse) GetReasonCodeOk() (*AchDisputeReasonCodes, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *DisputeActionResponse) SetReasonCode(v AchDisputeReasonCodes)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *DisputeActionResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetRepresentmentReasonCode

`func (o *DisputeActionResponse) GetRepresentmentReasonCode() CardDisputeRepresentmentReasonCode`

GetRepresentmentReasonCode returns the RepresentmentReasonCode field if non-nil, zero value otherwise.

### GetRepresentmentReasonCodeOk

`func (o *DisputeActionResponse) GetRepresentmentReasonCodeOk() (*CardDisputeRepresentmentReasonCode, bool)`

GetRepresentmentReasonCodeOk returns a tuple with the RepresentmentReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentmentReasonCode

`func (o *DisputeActionResponse) SetRepresentmentReasonCode(v CardDisputeRepresentmentReasonCode)`

SetRepresentmentReasonCode sets RepresentmentReasonCode field to given value.

### HasRepresentmentReasonCode

`func (o *DisputeActionResponse) HasRepresentmentReasonCode() bool`

HasRepresentmentReasonCode returns a boolean if a field has been set.

### GetStatus

`func (o *DisputeActionResponse) GetStatus() ExternalCardDisputeActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisputeActionResponse) GetStatusOk() (*ExternalCardDisputeActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisputeActionResponse) SetStatus(v ExternalCardDisputeActionStatus)`

SetStatus sets Status field to given value.


### GetSupportingDocId

`func (o *DisputeActionResponse) GetSupportingDocId() string`

GetSupportingDocId returns the SupportingDocId field if non-nil, zero value otherwise.

### GetSupportingDocIdOk

`func (o *DisputeActionResponse) GetSupportingDocIdOk() (*string, bool)`

GetSupportingDocIdOk returns a tuple with the SupportingDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingDocId

`func (o *DisputeActionResponse) SetSupportingDocId(v string)`

SetSupportingDocId sets SupportingDocId field to given value.

### HasSupportingDocId

`func (o *DisputeActionResponse) HasSupportingDocId() bool`

HasSupportingDocId returns a boolean if a field has been set.

### GetTenant

`func (o *DisputeActionResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DisputeActionResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DisputeActionResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetUpdatedReasonCodeMemo

`func (o *DisputeActionResponse) GetUpdatedReasonCodeMemo() string`

GetUpdatedReasonCodeMemo returns the UpdatedReasonCodeMemo field if non-nil, zero value otherwise.

### GetUpdatedReasonCodeMemoOk

`func (o *DisputeActionResponse) GetUpdatedReasonCodeMemoOk() (*string, bool)`

GetUpdatedReasonCodeMemoOk returns a tuple with the UpdatedReasonCodeMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReasonCodeMemo

`func (o *DisputeActionResponse) SetUpdatedReasonCodeMemo(v string)`

SetUpdatedReasonCodeMemo sets UpdatedReasonCodeMemo field to given value.

### HasUpdatedReasonCodeMemo

`func (o *DisputeActionResponse) HasUpdatedReasonCodeMemo() bool`

HasUpdatedReasonCodeMemo returns a boolean if a field has been set.

### GetViolationCode

`func (o *DisputeActionResponse) GetViolationCode() CardActionViolationCode`

GetViolationCode returns the ViolationCode field if non-nil, zero value otherwise.

### GetViolationCodeOk

`func (o *DisputeActionResponse) GetViolationCodeOk() (*CardActionViolationCode, bool)`

GetViolationCodeOk returns a tuple with the ViolationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCode

`func (o *DisputeActionResponse) SetViolationCode(v CardActionViolationCode)`

SetViolationCode sets ViolationCode field to given value.

### HasViolationCode

`func (o *DisputeActionResponse) HasViolationCode() bool`

HasViolationCode returns a boolean if a field has been set.

### GetViolationDate

`func (o *DisputeActionResponse) GetViolationDate() time.Time`

GetViolationDate returns the ViolationDate field if non-nil, zero value otherwise.

### GetViolationDateOk

`func (o *DisputeActionResponse) GetViolationDateOk() (*time.Time, bool)`

GetViolationDateOk returns a tuple with the ViolationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationDate

`func (o *DisputeActionResponse) SetViolationDate(v time.Time)`

SetViolationDate sets ViolationDate field to given value.

### HasViolationDate

`func (o *DisputeActionResponse) HasViolationDate() bool`

HasViolationDate returns a boolean if a field has been set.

### GetReturnCode

`func (o *DisputeActionResponse) GetReturnCode() AchDisputeReturnCodes`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *DisputeActionResponse) GetReturnCodeOk() (*AchDisputeReturnCodes, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *DisputeActionResponse) SetReturnCode(v AchDisputeReturnCodes)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *DisputeActionResponse) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


