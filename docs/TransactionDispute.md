# TransactionDispute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | The creation time of the dispute | 
**ExternalCaseReference** | Pointer to **string** | The external case number or id for the dispute (eg: from a vendor such as Marqeta), if one exists. | [optional] 
**Id** | **string** | The unique identifier of the dispute. | 
**InternalCaseReference** | Pointer to **string** | The internal case number or id for the dispute in the Synctera platform, if one exists. | [optional] 
**Status** | **string** |  | 
**Updated** | **time.Time** | The time the dispute was last updated | 

## Methods

### NewTransactionDispute

`func NewTransactionDispute(created time.Time, id string, status string, updated time.Time, ) *TransactionDispute`

NewTransactionDispute instantiates a new TransactionDispute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDisputeWithDefaults

`func NewTransactionDisputeWithDefaults() *TransactionDispute`

NewTransactionDisputeWithDefaults instantiates a new TransactionDispute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *TransactionDispute) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TransactionDispute) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TransactionDispute) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExternalCaseReference

`func (o *TransactionDispute) GetExternalCaseReference() string`

GetExternalCaseReference returns the ExternalCaseReference field if non-nil, zero value otherwise.

### GetExternalCaseReferenceOk

`func (o *TransactionDispute) GetExternalCaseReferenceOk() (*string, bool)`

GetExternalCaseReferenceOk returns a tuple with the ExternalCaseReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCaseReference

`func (o *TransactionDispute) SetExternalCaseReference(v string)`

SetExternalCaseReference sets ExternalCaseReference field to given value.

### HasExternalCaseReference

`func (o *TransactionDispute) HasExternalCaseReference() bool`

HasExternalCaseReference returns a boolean if a field has been set.

### GetId

`func (o *TransactionDispute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionDispute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionDispute) SetId(v string)`

SetId sets Id field to given value.


### GetInternalCaseReference

`func (o *TransactionDispute) GetInternalCaseReference() string`

GetInternalCaseReference returns the InternalCaseReference field if non-nil, zero value otherwise.

### GetInternalCaseReferenceOk

`func (o *TransactionDispute) GetInternalCaseReferenceOk() (*string, bool)`

GetInternalCaseReferenceOk returns a tuple with the InternalCaseReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalCaseReference

`func (o *TransactionDispute) SetInternalCaseReference(v string)`

SetInternalCaseReference sets InternalCaseReference field to given value.

### HasInternalCaseReference

`func (o *TransactionDispute) HasInternalCaseReference() bool`

HasInternalCaseReference returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionDispute) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionDispute) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionDispute) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdated

`func (o *TransactionDispute) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TransactionDispute) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TransactionDispute) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


