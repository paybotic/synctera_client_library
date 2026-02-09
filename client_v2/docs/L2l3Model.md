# L2l3Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnhancedDataId** | Pointer to **string** |  | [optional] 
**Financial** | Pointer to [**Financial**](Financial.md) |  | [optional] 
**FleetEmv** | Pointer to [**FleetsEmv**](FleetsEmv.md) |  | [optional] 
**Fleets** | Pointer to [**Fleets**](Fleets.md) |  | [optional] 
**InventoryDetails** | Pointer to [**[]InventoryDetails**](InventoryDetails.md) |  | [optional] 
**OriginalTransactionId** | **string** | Value of the &#x60;transaction.token&#x60; returned in the simulated clearing response. | 
**PurchaseOrder** | Pointer to **string** |  | [optional] 

## Methods

### NewL2l3Model

`func NewL2l3Model(originalTransactionId string, ) *L2l3Model`

NewL2l3Model instantiates a new L2l3Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewL2l3ModelWithDefaults

`func NewL2l3ModelWithDefaults() *L2l3Model`

NewL2l3ModelWithDefaults instantiates a new L2l3Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnhancedDataId

`func (o *L2l3Model) GetEnhancedDataId() string`

GetEnhancedDataId returns the EnhancedDataId field if non-nil, zero value otherwise.

### GetEnhancedDataIdOk

`func (o *L2l3Model) GetEnhancedDataIdOk() (*string, bool)`

GetEnhancedDataIdOk returns a tuple with the EnhancedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedDataId

`func (o *L2l3Model) SetEnhancedDataId(v string)`

SetEnhancedDataId sets EnhancedDataId field to given value.

### HasEnhancedDataId

`func (o *L2l3Model) HasEnhancedDataId() bool`

HasEnhancedDataId returns a boolean if a field has been set.

### GetFinancial

`func (o *L2l3Model) GetFinancial() Financial`

GetFinancial returns the Financial field if non-nil, zero value otherwise.

### GetFinancialOk

`func (o *L2l3Model) GetFinancialOk() (*Financial, bool)`

GetFinancialOk returns a tuple with the Financial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancial

`func (o *L2l3Model) SetFinancial(v Financial)`

SetFinancial sets Financial field to given value.

### HasFinancial

`func (o *L2l3Model) HasFinancial() bool`

HasFinancial returns a boolean if a field has been set.

### GetFleetEmv

`func (o *L2l3Model) GetFleetEmv() FleetsEmv`

GetFleetEmv returns the FleetEmv field if non-nil, zero value otherwise.

### GetFleetEmvOk

`func (o *L2l3Model) GetFleetEmvOk() (*FleetsEmv, bool)`

GetFleetEmvOk returns a tuple with the FleetEmv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetEmv

`func (o *L2l3Model) SetFleetEmv(v FleetsEmv)`

SetFleetEmv sets FleetEmv field to given value.

### HasFleetEmv

`func (o *L2l3Model) HasFleetEmv() bool`

HasFleetEmv returns a boolean if a field has been set.

### GetFleets

`func (o *L2l3Model) GetFleets() Fleets`

GetFleets returns the Fleets field if non-nil, zero value otherwise.

### GetFleetsOk

`func (o *L2l3Model) GetFleetsOk() (*Fleets, bool)`

GetFleetsOk returns a tuple with the Fleets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleets

`func (o *L2l3Model) SetFleets(v Fleets)`

SetFleets sets Fleets field to given value.

### HasFleets

`func (o *L2l3Model) HasFleets() bool`

HasFleets returns a boolean if a field has been set.

### GetInventoryDetails

`func (o *L2l3Model) GetInventoryDetails() []InventoryDetails`

GetInventoryDetails returns the InventoryDetails field if non-nil, zero value otherwise.

### GetInventoryDetailsOk

`func (o *L2l3Model) GetInventoryDetailsOk() (*[]InventoryDetails, bool)`

GetInventoryDetailsOk returns a tuple with the InventoryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDetails

`func (o *L2l3Model) SetInventoryDetails(v []InventoryDetails)`

SetInventoryDetails sets InventoryDetails field to given value.

### HasInventoryDetails

`func (o *L2l3Model) HasInventoryDetails() bool`

HasInventoryDetails returns a boolean if a field has been set.

### GetOriginalTransactionId

`func (o *L2l3Model) GetOriginalTransactionId() string`

GetOriginalTransactionId returns the OriginalTransactionId field if non-nil, zero value otherwise.

### GetOriginalTransactionIdOk

`func (o *L2l3Model) GetOriginalTransactionIdOk() (*string, bool)`

GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTransactionId

`func (o *L2l3Model) SetOriginalTransactionId(v string)`

SetOriginalTransactionId sets OriginalTransactionId field to given value.


### GetPurchaseOrder

`func (o *L2l3Model) GetPurchaseOrder() string`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *L2l3Model) GetPurchaseOrderOk() (*string, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *L2l3Model) SetPurchaseOrder(v string)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *L2l3Model) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


