# FleetsEmv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeNumber** | Pointer to **string** |  | [optional] 
**ExpandedFuelType** | Pointer to **string** |  | [optional] 
**FuelGrossAmount** | Pointer to **int64** |  | [optional] 
**FuelNetAmount** | Pointer to **int64** |  | [optional] 
**FuelQuantity** | Pointer to **float32** |  | [optional] 
**FuelUnitOfMeasure** | Pointer to **string** |  | [optional] 
**FuelUnitPrice** | Pointer to **float32** |  | [optional] 
**NonFuelGrossAmount** | Pointer to **int64** |  | [optional] 
**NonFuelItemDetails** | Pointer to [**[]NonFuelItemDetails**](NonFuelItemDetails.md) |  | [optional] 
**NonFuelNetAmount** | Pointer to **int64** |  | [optional] 
**OdometerReading** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**TrailerNumber** | Pointer to **string** |  | [optional] 
**TypeOfPurchase** | Pointer to **string** |  | [optional] 
**VatTaxRate** | Pointer to **float32** |  | [optional] 

## Methods

### NewFleetsEmv

`func NewFleetsEmv() *FleetsEmv`

NewFleetsEmv instantiates a new FleetsEmv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetsEmvWithDefaults

`func NewFleetsEmvWithDefaults() *FleetsEmv`

NewFleetsEmvWithDefaults instantiates a new FleetsEmv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeNumber

`func (o *FleetsEmv) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *FleetsEmv) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *FleetsEmv) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *FleetsEmv) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetExpandedFuelType

`func (o *FleetsEmv) GetExpandedFuelType() string`

GetExpandedFuelType returns the ExpandedFuelType field if non-nil, zero value otherwise.

### GetExpandedFuelTypeOk

`func (o *FleetsEmv) GetExpandedFuelTypeOk() (*string, bool)`

GetExpandedFuelTypeOk returns a tuple with the ExpandedFuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandedFuelType

`func (o *FleetsEmv) SetExpandedFuelType(v string)`

SetExpandedFuelType sets ExpandedFuelType field to given value.

### HasExpandedFuelType

`func (o *FleetsEmv) HasExpandedFuelType() bool`

HasExpandedFuelType returns a boolean if a field has been set.

### GetFuelGrossAmount

`func (o *FleetsEmv) GetFuelGrossAmount() int64`

GetFuelGrossAmount returns the FuelGrossAmount field if non-nil, zero value otherwise.

### GetFuelGrossAmountOk

`func (o *FleetsEmv) GetFuelGrossAmountOk() (*int64, bool)`

GetFuelGrossAmountOk returns a tuple with the FuelGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelGrossAmount

`func (o *FleetsEmv) SetFuelGrossAmount(v int64)`

SetFuelGrossAmount sets FuelGrossAmount field to given value.

### HasFuelGrossAmount

`func (o *FleetsEmv) HasFuelGrossAmount() bool`

HasFuelGrossAmount returns a boolean if a field has been set.

### GetFuelNetAmount

`func (o *FleetsEmv) GetFuelNetAmount() int64`

GetFuelNetAmount returns the FuelNetAmount field if non-nil, zero value otherwise.

### GetFuelNetAmountOk

`func (o *FleetsEmv) GetFuelNetAmountOk() (*int64, bool)`

GetFuelNetAmountOk returns a tuple with the FuelNetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelNetAmount

`func (o *FleetsEmv) SetFuelNetAmount(v int64)`

SetFuelNetAmount sets FuelNetAmount field to given value.

### HasFuelNetAmount

`func (o *FleetsEmv) HasFuelNetAmount() bool`

HasFuelNetAmount returns a boolean if a field has been set.

### GetFuelQuantity

`func (o *FleetsEmv) GetFuelQuantity() float32`

GetFuelQuantity returns the FuelQuantity field if non-nil, zero value otherwise.

### GetFuelQuantityOk

`func (o *FleetsEmv) GetFuelQuantityOk() (*float32, bool)`

GetFuelQuantityOk returns a tuple with the FuelQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelQuantity

`func (o *FleetsEmv) SetFuelQuantity(v float32)`

SetFuelQuantity sets FuelQuantity field to given value.

### HasFuelQuantity

`func (o *FleetsEmv) HasFuelQuantity() bool`

HasFuelQuantity returns a boolean if a field has been set.

### GetFuelUnitOfMeasure

`func (o *FleetsEmv) GetFuelUnitOfMeasure() string`

GetFuelUnitOfMeasure returns the FuelUnitOfMeasure field if non-nil, zero value otherwise.

### GetFuelUnitOfMeasureOk

`func (o *FleetsEmv) GetFuelUnitOfMeasureOk() (*string, bool)`

GetFuelUnitOfMeasureOk returns a tuple with the FuelUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelUnitOfMeasure

`func (o *FleetsEmv) SetFuelUnitOfMeasure(v string)`

SetFuelUnitOfMeasure sets FuelUnitOfMeasure field to given value.

### HasFuelUnitOfMeasure

`func (o *FleetsEmv) HasFuelUnitOfMeasure() bool`

HasFuelUnitOfMeasure returns a boolean if a field has been set.

### GetFuelUnitPrice

`func (o *FleetsEmv) GetFuelUnitPrice() float32`

GetFuelUnitPrice returns the FuelUnitPrice field if non-nil, zero value otherwise.

### GetFuelUnitPriceOk

`func (o *FleetsEmv) GetFuelUnitPriceOk() (*float32, bool)`

GetFuelUnitPriceOk returns a tuple with the FuelUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelUnitPrice

`func (o *FleetsEmv) SetFuelUnitPrice(v float32)`

SetFuelUnitPrice sets FuelUnitPrice field to given value.

### HasFuelUnitPrice

`func (o *FleetsEmv) HasFuelUnitPrice() bool`

HasFuelUnitPrice returns a boolean if a field has been set.

### GetNonFuelGrossAmount

`func (o *FleetsEmv) GetNonFuelGrossAmount() int64`

GetNonFuelGrossAmount returns the NonFuelGrossAmount field if non-nil, zero value otherwise.

### GetNonFuelGrossAmountOk

`func (o *FleetsEmv) GetNonFuelGrossAmountOk() (*int64, bool)`

GetNonFuelGrossAmountOk returns a tuple with the NonFuelGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFuelGrossAmount

`func (o *FleetsEmv) SetNonFuelGrossAmount(v int64)`

SetNonFuelGrossAmount sets NonFuelGrossAmount field to given value.

### HasNonFuelGrossAmount

`func (o *FleetsEmv) HasNonFuelGrossAmount() bool`

HasNonFuelGrossAmount returns a boolean if a field has been set.

### GetNonFuelItemDetails

`func (o *FleetsEmv) GetNonFuelItemDetails() []NonFuelItemDetails`

GetNonFuelItemDetails returns the NonFuelItemDetails field if non-nil, zero value otherwise.

### GetNonFuelItemDetailsOk

`func (o *FleetsEmv) GetNonFuelItemDetailsOk() (*[]NonFuelItemDetails, bool)`

GetNonFuelItemDetailsOk returns a tuple with the NonFuelItemDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFuelItemDetails

`func (o *FleetsEmv) SetNonFuelItemDetails(v []NonFuelItemDetails)`

SetNonFuelItemDetails sets NonFuelItemDetails field to given value.

### HasNonFuelItemDetails

`func (o *FleetsEmv) HasNonFuelItemDetails() bool`

HasNonFuelItemDetails returns a boolean if a field has been set.

### GetNonFuelNetAmount

`func (o *FleetsEmv) GetNonFuelNetAmount() int64`

GetNonFuelNetAmount returns the NonFuelNetAmount field if non-nil, zero value otherwise.

### GetNonFuelNetAmountOk

`func (o *FleetsEmv) GetNonFuelNetAmountOk() (*int64, bool)`

GetNonFuelNetAmountOk returns a tuple with the NonFuelNetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFuelNetAmount

`func (o *FleetsEmv) SetNonFuelNetAmount(v int64)`

SetNonFuelNetAmount sets NonFuelNetAmount field to given value.

### HasNonFuelNetAmount

`func (o *FleetsEmv) HasNonFuelNetAmount() bool`

HasNonFuelNetAmount returns a boolean if a field has been set.

### GetOdometerReading

`func (o *FleetsEmv) GetOdometerReading() string`

GetOdometerReading returns the OdometerReading field if non-nil, zero value otherwise.

### GetOdometerReadingOk

`func (o *FleetsEmv) GetOdometerReadingOk() (*string, bool)`

GetOdometerReadingOk returns a tuple with the OdometerReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerReading

`func (o *FleetsEmv) SetOdometerReading(v string)`

SetOdometerReading sets OdometerReading field to given value.

### HasOdometerReading

`func (o *FleetsEmv) HasOdometerReading() bool`

HasOdometerReading returns a boolean if a field has been set.

### GetServiceType

`func (o *FleetsEmv) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *FleetsEmv) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *FleetsEmv) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *FleetsEmv) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTrailerNumber

`func (o *FleetsEmv) GetTrailerNumber() string`

GetTrailerNumber returns the TrailerNumber field if non-nil, zero value otherwise.

### GetTrailerNumberOk

`func (o *FleetsEmv) GetTrailerNumberOk() (*string, bool)`

GetTrailerNumberOk returns a tuple with the TrailerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerNumber

`func (o *FleetsEmv) SetTrailerNumber(v string)`

SetTrailerNumber sets TrailerNumber field to given value.

### HasTrailerNumber

`func (o *FleetsEmv) HasTrailerNumber() bool`

HasTrailerNumber returns a boolean if a field has been set.

### GetTypeOfPurchase

`func (o *FleetsEmv) GetTypeOfPurchase() string`

GetTypeOfPurchase returns the TypeOfPurchase field if non-nil, zero value otherwise.

### GetTypeOfPurchaseOk

`func (o *FleetsEmv) GetTypeOfPurchaseOk() (*string, bool)`

GetTypeOfPurchaseOk returns a tuple with the TypeOfPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfPurchase

`func (o *FleetsEmv) SetTypeOfPurchase(v string)`

SetTypeOfPurchase sets TypeOfPurchase field to given value.

### HasTypeOfPurchase

`func (o *FleetsEmv) HasTypeOfPurchase() bool`

HasTypeOfPurchase returns a boolean if a field has been set.

### GetVatTaxRate

`func (o *FleetsEmv) GetVatTaxRate() float32`

GetVatTaxRate returns the VatTaxRate field if non-nil, zero value otherwise.

### GetVatTaxRateOk

`func (o *FleetsEmv) GetVatTaxRateOk() (*float32, bool)`

GetVatTaxRateOk returns a tuple with the VatTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatTaxRate

`func (o *FleetsEmv) SetVatTaxRate(v float32)`

SetVatTaxRate sets VatTaxRate field to given value.

### HasVatTaxRate

`func (o *FleetsEmv) HasVatTaxRate() bool`

HasVatTaxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


