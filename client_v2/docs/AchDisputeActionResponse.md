# AchDisputeActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**AchAction**](AchAction.md) |  | 
**CreationTime** | **time.Time** | The date and time the resource was created. | [readonly] 
**Id** | **string** | The unique identifier of the dispute action | [readonly] 
**Message** | Pointer to **string** | Message text related to card dispute action | [optional] 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**ReasonCode** | Pointer to [**AchDisputeReasonCodes**](AchDisputeReasonCodes.md) |  | [optional] 
**ReturnCode** | Pointer to [**AchDisputeReturnCodes**](AchDisputeReturnCodes.md) |  | [optional] 
**Status** | [**AchDisputeActionStatus**](AchDisputeActionStatus.md) |  | 
**SupportingDocId** | Pointer to **string** | The unique identifier of the supporting document | [optional] [readonly] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewAchDisputeActionResponse

`func NewAchDisputeActionResponse(action AchAction, creationTime time.Time, id string, paymentRail PaymentRail, status AchDisputeActionStatus, tenant string, ) *AchDisputeActionResponse`

NewAchDisputeActionResponse instantiates a new AchDisputeActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchDisputeActionResponseWithDefaults

`func NewAchDisputeActionResponseWithDefaults() *AchDisputeActionResponse`

NewAchDisputeActionResponseWithDefaults instantiates a new AchDisputeActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AchDisputeActionResponse) GetAction() AchAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AchDisputeActionResponse) GetActionOk() (*AchAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AchDisputeActionResponse) SetAction(v AchAction)`

SetAction sets Action field to given value.


### GetCreationTime

`func (o *AchDisputeActionResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AchDisputeActionResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AchDisputeActionResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *AchDisputeActionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AchDisputeActionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AchDisputeActionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *AchDisputeActionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AchDisputeActionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AchDisputeActionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AchDisputeActionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPaymentRail

`func (o *AchDisputeActionResponse) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *AchDisputeActionResponse) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *AchDisputeActionResponse) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetReasonCode

`func (o *AchDisputeActionResponse) GetReasonCode() AchDisputeReasonCodes`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *AchDisputeActionResponse) GetReasonCodeOk() (*AchDisputeReasonCodes, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *AchDisputeActionResponse) SetReasonCode(v AchDisputeReasonCodes)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *AchDisputeActionResponse) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReturnCode

`func (o *AchDisputeActionResponse) GetReturnCode() AchDisputeReturnCodes`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *AchDisputeActionResponse) GetReturnCodeOk() (*AchDisputeReturnCodes, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *AchDisputeActionResponse) SetReturnCode(v AchDisputeReturnCodes)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *AchDisputeActionResponse) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetStatus

`func (o *AchDisputeActionResponse) GetStatus() AchDisputeActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchDisputeActionResponse) GetStatusOk() (*AchDisputeActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AchDisputeActionResponse) SetStatus(v AchDisputeActionStatus)`

SetStatus sets Status field to given value.


### GetSupportingDocId

`func (o *AchDisputeActionResponse) GetSupportingDocId() string`

GetSupportingDocId returns the SupportingDocId field if non-nil, zero value otherwise.

### GetSupportingDocIdOk

`func (o *AchDisputeActionResponse) GetSupportingDocIdOk() (*string, bool)`

GetSupportingDocIdOk returns a tuple with the SupportingDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingDocId

`func (o *AchDisputeActionResponse) SetSupportingDocId(v string)`

SetSupportingDocId sets SupportingDocId field to given value.

### HasSupportingDocId

`func (o *AchDisputeActionResponse) HasSupportingDocId() bool`

HasSupportingDocId returns a boolean if a field has been set.

### GetTenant

`func (o *AchDisputeActionResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AchDisputeActionResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AchDisputeActionResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


