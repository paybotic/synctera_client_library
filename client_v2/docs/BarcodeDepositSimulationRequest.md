# BarcodeDepositSimulationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount to deposit | 
**BarcodeId** | **string** | Unique identifier of the barcode to be used for the deposit | 
**Status** | **string** |  | 
**StoreId** | **string** | Unique identifier of the barcode simulation store | 
**Type** | **string** |  | 

## Methods

### NewBarcodeDepositSimulationRequest

`func NewBarcodeDepositSimulationRequest(amount float32, barcodeId string, status string, storeId string, type_ string, ) *BarcodeDepositSimulationRequest`

NewBarcodeDepositSimulationRequest instantiates a new BarcodeDepositSimulationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBarcodeDepositSimulationRequestWithDefaults

`func NewBarcodeDepositSimulationRequestWithDefaults() *BarcodeDepositSimulationRequest`

NewBarcodeDepositSimulationRequestWithDefaults instantiates a new BarcodeDepositSimulationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BarcodeDepositSimulationRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BarcodeDepositSimulationRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BarcodeDepositSimulationRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetBarcodeId

`func (o *BarcodeDepositSimulationRequest) GetBarcodeId() string`

GetBarcodeId returns the BarcodeId field if non-nil, zero value otherwise.

### GetBarcodeIdOk

`func (o *BarcodeDepositSimulationRequest) GetBarcodeIdOk() (*string, bool)`

GetBarcodeIdOk returns a tuple with the BarcodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeId

`func (o *BarcodeDepositSimulationRequest) SetBarcodeId(v string)`

SetBarcodeId sets BarcodeId field to given value.


### GetStatus

`func (o *BarcodeDepositSimulationRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BarcodeDepositSimulationRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BarcodeDepositSimulationRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStoreId

`func (o *BarcodeDepositSimulationRequest) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *BarcodeDepositSimulationRequest) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *BarcodeDepositSimulationRequest) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.


### GetType

`func (o *BarcodeDepositSimulationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BarcodeDepositSimulationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BarcodeDepositSimulationRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


