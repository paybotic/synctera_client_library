# PhysicalDebitCardResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**PendingReasons** | Pointer to [**CardStatusPendingReasons**](CardStatusPendingReasons.md) |  | [optional] 
**StatusReason** | [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | 
**CardFulfillmentStatus** | [**CardFulfillmentStatus**](CardFulfillmentStatus.md) |  | 
**FulfillmentDetails** | Pointer to [**FulfillmentDetails**](FulfillmentDetails.md) |  | [optional] 

## Methods

### NewPhysicalDebitCardResponseStatus

`func NewPhysicalDebitCardResponseStatus(cardStatus CardStatus, statusReason CardStatusReasonCode, cardFulfillmentStatus CardFulfillmentStatus, ) *PhysicalDebitCardResponseStatus`

NewPhysicalDebitCardResponseStatus instantiates a new PhysicalDebitCardResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalDebitCardResponseStatusWithDefaults

`func NewPhysicalDebitCardResponseStatusWithDefaults() *PhysicalDebitCardResponseStatus`

NewPhysicalDebitCardResponseStatusWithDefaults instantiates a new PhysicalDebitCardResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardStatus

`func (o *PhysicalDebitCardResponseStatus) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *PhysicalDebitCardResponseStatus) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *PhysicalDebitCardResponseStatus) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetMemo

`func (o *PhysicalDebitCardResponseStatus) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PhysicalDebitCardResponseStatus) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PhysicalDebitCardResponseStatus) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *PhysicalDebitCardResponseStatus) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPendingReasons

`func (o *PhysicalDebitCardResponseStatus) GetPendingReasons() CardStatusPendingReasons`

GetPendingReasons returns the PendingReasons field if non-nil, zero value otherwise.

### GetPendingReasonsOk

`func (o *PhysicalDebitCardResponseStatus) GetPendingReasonsOk() (*CardStatusPendingReasons, bool)`

GetPendingReasonsOk returns a tuple with the PendingReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReasons

`func (o *PhysicalDebitCardResponseStatus) SetPendingReasons(v CardStatusPendingReasons)`

SetPendingReasons sets PendingReasons field to given value.

### HasPendingReasons

`func (o *PhysicalDebitCardResponseStatus) HasPendingReasons() bool`

HasPendingReasons returns a boolean if a field has been set.

### GetStatusReason

`func (o *PhysicalDebitCardResponseStatus) GetStatusReason() CardStatusReasonCode`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PhysicalDebitCardResponseStatus) GetStatusReasonOk() (*CardStatusReasonCode, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PhysicalDebitCardResponseStatus) SetStatusReason(v CardStatusReasonCode)`

SetStatusReason sets StatusReason field to given value.


### GetCardFulfillmentStatus

`func (o *PhysicalDebitCardResponseStatus) GetCardFulfillmentStatus() CardFulfillmentStatus`

GetCardFulfillmentStatus returns the CardFulfillmentStatus field if non-nil, zero value otherwise.

### GetCardFulfillmentStatusOk

`func (o *PhysicalDebitCardResponseStatus) GetCardFulfillmentStatusOk() (*CardFulfillmentStatus, bool)`

GetCardFulfillmentStatusOk returns a tuple with the CardFulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentStatus

`func (o *PhysicalDebitCardResponseStatus) SetCardFulfillmentStatus(v CardFulfillmentStatus)`

SetCardFulfillmentStatus sets CardFulfillmentStatus field to given value.


### GetFulfillmentDetails

`func (o *PhysicalDebitCardResponseStatus) GetFulfillmentDetails() FulfillmentDetails`

GetFulfillmentDetails returns the FulfillmentDetails field if non-nil, zero value otherwise.

### GetFulfillmentDetailsOk

`func (o *PhysicalDebitCardResponseStatus) GetFulfillmentDetailsOk() (*FulfillmentDetails, bool)`

GetFulfillmentDetailsOk returns a tuple with the FulfillmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentDetails

`func (o *PhysicalDebitCardResponseStatus) SetFulfillmentDetails(v FulfillmentDetails)`

SetFulfillmentDetails sets FulfillmentDetails field to given value.

### HasFulfillmentDetails

`func (o *PhysicalDebitCardResponseStatus) HasFulfillmentDetails() bool`

HasFulfillmentDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


