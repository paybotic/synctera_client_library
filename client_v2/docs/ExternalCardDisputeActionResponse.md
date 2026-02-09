# ExternalCardDisputeActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ExternalCardAction**](ExternalCardAction.md) |  | 
**CreationTime** | **time.Time** | The timestamp representing when the object was created | [readonly] 
**Id** | **string** | The unique identifier of the dispute action | [readonly] 
**PaymentRail** | [**PaymentRail**](PaymentRail.md) |  | 
**Status** | [**ExternalCardDisputeActionStatus**](ExternalCardDisputeActionStatus.md) |  | 
**SupportingDocId** | Pointer to **string** | The unique identifier of the supporting document | [optional] [readonly] 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewExternalCardDisputeActionResponse

`func NewExternalCardDisputeActionResponse(action ExternalCardAction, creationTime time.Time, id string, paymentRail PaymentRail, status ExternalCardDisputeActionStatus, tenant string, ) *ExternalCardDisputeActionResponse`

NewExternalCardDisputeActionResponse instantiates a new ExternalCardDisputeActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardDisputeActionResponseWithDefaults

`func NewExternalCardDisputeActionResponseWithDefaults() *ExternalCardDisputeActionResponse`

NewExternalCardDisputeActionResponseWithDefaults instantiates a new ExternalCardDisputeActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ExternalCardDisputeActionResponse) GetAction() ExternalCardAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ExternalCardDisputeActionResponse) GetActionOk() (*ExternalCardAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ExternalCardDisputeActionResponse) SetAction(v ExternalCardAction)`

SetAction sets Action field to given value.


### GetCreationTime

`func (o *ExternalCardDisputeActionResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ExternalCardDisputeActionResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ExternalCardDisputeActionResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *ExternalCardDisputeActionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalCardDisputeActionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalCardDisputeActionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPaymentRail

`func (o *ExternalCardDisputeActionResponse) GetPaymentRail() PaymentRail`

GetPaymentRail returns the PaymentRail field if non-nil, zero value otherwise.

### GetPaymentRailOk

`func (o *ExternalCardDisputeActionResponse) GetPaymentRailOk() (*PaymentRail, bool)`

GetPaymentRailOk returns a tuple with the PaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRail

`func (o *ExternalCardDisputeActionResponse) SetPaymentRail(v PaymentRail)`

SetPaymentRail sets PaymentRail field to given value.


### GetStatus

`func (o *ExternalCardDisputeActionResponse) GetStatus() ExternalCardDisputeActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalCardDisputeActionResponse) GetStatusOk() (*ExternalCardDisputeActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalCardDisputeActionResponse) SetStatus(v ExternalCardDisputeActionStatus)`

SetStatus sets Status field to given value.


### GetSupportingDocId

`func (o *ExternalCardDisputeActionResponse) GetSupportingDocId() string`

GetSupportingDocId returns the SupportingDocId field if non-nil, zero value otherwise.

### GetSupportingDocIdOk

`func (o *ExternalCardDisputeActionResponse) GetSupportingDocIdOk() (*string, bool)`

GetSupportingDocIdOk returns a tuple with the SupportingDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingDocId

`func (o *ExternalCardDisputeActionResponse) SetSupportingDocId(v string)`

SetSupportingDocId sets SupportingDocId field to given value.

### HasSupportingDocId

`func (o *ExternalCardDisputeActionResponse) HasSupportingDocId() bool`

HasSupportingDocId returns a boolean if a field has been set.

### GetTenant

`func (o *ExternalCardDisputeActionResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ExternalCardDisputeActionResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ExternalCardDisputeActionResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


