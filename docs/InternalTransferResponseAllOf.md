# InternalTransferResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The transaction id associated with the transfer | 
**Status** | **string** | The status of the internal transfer auth. A value of &#x60;PENDING&#x60; indicates that the funds have been reserved and the transaction is ready to be either completed or canceled. A value of &#x60;COMPLETE&#x60; indicates the funds have been successfully moved and no more action can be performed. A value of &#x60;CANCELED&#x60; or &#x60;EXPIRED&#x60; means that the transaction has rolled back and the funds have been returned to the originating account, either by explicitly canceling via the API, or due to the expiry time having passed. | 

## Methods

### NewInternalTransferResponseAllOf

`func NewInternalTransferResponseAllOf(id string, status string, ) *InternalTransferResponseAllOf`

NewInternalTransferResponseAllOf instantiates a new InternalTransferResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalTransferResponseAllOfWithDefaults

`func NewInternalTransferResponseAllOfWithDefaults() *InternalTransferResponseAllOf`

NewInternalTransferResponseAllOfWithDefaults instantiates a new InternalTransferResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalTransferResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalTransferResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalTransferResponseAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *InternalTransferResponseAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalTransferResponseAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalTransferResponseAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


