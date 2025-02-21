# CardStatusPendingReasons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsWaitingForBulkShipment** | Pointer to **bool** | The card is to be shipping as part of a bulk shipment but that bulk shipment has not yet been released.  | [optional] 
**IsWaitingForImage** | Pointer to **bool** | The card has a custom image and either that image hasn&#39;t yet been uploaded and approved or the card has not yet been processed by the periodic daily custom card image processing task.  | [optional] 
**IsWaitingForPin** | Pointer to **bool** | The card requires a PIN to be set before it can be issued (refer to the pin_issuance_policy of the related card product). The PIN has not yet been set and not enough time has passed to use a random PIN (if applicable).  | [optional] 

## Methods

### NewCardStatusPendingReasons

`func NewCardStatusPendingReasons() *CardStatusPendingReasons`

NewCardStatusPendingReasons instantiates a new CardStatusPendingReasons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardStatusPendingReasonsWithDefaults

`func NewCardStatusPendingReasonsWithDefaults() *CardStatusPendingReasons`

NewCardStatusPendingReasonsWithDefaults instantiates a new CardStatusPendingReasons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsWaitingForBulkShipment

`func (o *CardStatusPendingReasons) GetIsWaitingForBulkShipment() bool`

GetIsWaitingForBulkShipment returns the IsWaitingForBulkShipment field if non-nil, zero value otherwise.

### GetIsWaitingForBulkShipmentOk

`func (o *CardStatusPendingReasons) GetIsWaitingForBulkShipmentOk() (*bool, bool)`

GetIsWaitingForBulkShipmentOk returns a tuple with the IsWaitingForBulkShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWaitingForBulkShipment

`func (o *CardStatusPendingReasons) SetIsWaitingForBulkShipment(v bool)`

SetIsWaitingForBulkShipment sets IsWaitingForBulkShipment field to given value.

### HasIsWaitingForBulkShipment

`func (o *CardStatusPendingReasons) HasIsWaitingForBulkShipment() bool`

HasIsWaitingForBulkShipment returns a boolean if a field has been set.

### GetIsWaitingForImage

`func (o *CardStatusPendingReasons) GetIsWaitingForImage() bool`

GetIsWaitingForImage returns the IsWaitingForImage field if non-nil, zero value otherwise.

### GetIsWaitingForImageOk

`func (o *CardStatusPendingReasons) GetIsWaitingForImageOk() (*bool, bool)`

GetIsWaitingForImageOk returns a tuple with the IsWaitingForImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWaitingForImage

`func (o *CardStatusPendingReasons) SetIsWaitingForImage(v bool)`

SetIsWaitingForImage sets IsWaitingForImage field to given value.

### HasIsWaitingForImage

`func (o *CardStatusPendingReasons) HasIsWaitingForImage() bool`

HasIsWaitingForImage returns a boolean if a field has been set.

### GetIsWaitingForPin

`func (o *CardStatusPendingReasons) GetIsWaitingForPin() bool`

GetIsWaitingForPin returns the IsWaitingForPin field if non-nil, zero value otherwise.

### GetIsWaitingForPinOk

`func (o *CardStatusPendingReasons) GetIsWaitingForPinOk() (*bool, bool)`

GetIsWaitingForPinOk returns a tuple with the IsWaitingForPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWaitingForPin

`func (o *CardStatusPendingReasons) SetIsWaitingForPin(v bool)`

SetIsWaitingForPin sets IsWaitingForPin field to given value.

### HasIsWaitingForPin

`func (o *CardStatusPendingReasons) HasIsWaitingForPin() bool`

HasIsWaitingForPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


