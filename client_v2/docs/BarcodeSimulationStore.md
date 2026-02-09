# BarcodeSimulationStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**BarcodeSimulationStoreAddress**](BarcodeSimulationStoreAddress.md) |  | 
**BusinessName** | **string** | Name of the store | 
**Coordinates** | [**BarcodeSimulationStoreCoordinates**](BarcodeSimulationStoreCoordinates.md) |  | 
**Distance** | **float32** | Distance from the provided location in miles | 
**Id** | **string** | Unique identifier for the store | 

## Methods

### NewBarcodeSimulationStore

`func NewBarcodeSimulationStore(address BarcodeSimulationStoreAddress, businessName string, coordinates BarcodeSimulationStoreCoordinates, distance float32, id string, ) *BarcodeSimulationStore`

NewBarcodeSimulationStore instantiates a new BarcodeSimulationStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBarcodeSimulationStoreWithDefaults

`func NewBarcodeSimulationStoreWithDefaults() *BarcodeSimulationStore`

NewBarcodeSimulationStoreWithDefaults instantiates a new BarcodeSimulationStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BarcodeSimulationStore) GetAddress() BarcodeSimulationStoreAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BarcodeSimulationStore) GetAddressOk() (*BarcodeSimulationStoreAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BarcodeSimulationStore) SetAddress(v BarcodeSimulationStoreAddress)`

SetAddress sets Address field to given value.


### GetBusinessName

`func (o *BarcodeSimulationStore) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *BarcodeSimulationStore) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *BarcodeSimulationStore) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.


### GetCoordinates

`func (o *BarcodeSimulationStore) GetCoordinates() BarcodeSimulationStoreCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *BarcodeSimulationStore) GetCoordinatesOk() (*BarcodeSimulationStoreCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *BarcodeSimulationStore) SetCoordinates(v BarcodeSimulationStoreCoordinates)`

SetCoordinates sets Coordinates field to given value.


### GetDistance

`func (o *BarcodeSimulationStore) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *BarcodeSimulationStore) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *BarcodeSimulationStore) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetId

`func (o *BarcodeSimulationStore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BarcodeSimulationStore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BarcodeSimulationStore) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


