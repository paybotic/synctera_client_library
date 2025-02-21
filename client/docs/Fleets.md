# Fleets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverId** | Pointer to **string** |  | [optional] 
**ExpandedFuelType** | Pointer to **string** |  | [optional] 
**FleetNumber** | Pointer to **string** |  | [optional] 
**FuelGrossAmount** | Pointer to **int64** |  | [optional] 
**FuelNetAmount** | Pointer to **int64** |  | [optional] 
**FuelProductQualifier** | Pointer to **string** |  | [optional] 
**FuelPurchaseType** | Pointer to **string** |  | [optional] 
**FuelQuantity** | Pointer to **float32** |  | [optional] 
**FuelServiceType** | Pointer to **string** |  | [optional] 
**FuelTaxAmount** | Pointer to **int64** |  | [optional] 
**FuelTaxExemptionStatus** | Pointer to **string** |  | [optional] 
**FuelType** | Pointer to **string** |  | [optional] 
**FuelUnitOfMeasure** | Pointer to **string** |  | [optional] 
**FuelUnitPrice** | Pointer to **float32** |  | [optional] 
**JobNumber** | Pointer to **string** |  | [optional] 
**NonFuelGrossAmount** | Pointer to **int64** |  | [optional] 
**NonFuelItemDetails** | Pointer to [**[]NonFuelItemDetails**](NonFuelItemDetails.md) |  | [optional] 
**NonFuelNetAmount** | Pointer to **int64** |  | [optional] 
**NonFuelTaxAmount** | Pointer to **int64** |  | [optional] 
**NonFuelTaxExemptionStatus** | Pointer to **string** |  | [optional] 
**OdometerReading** | Pointer to **string** |  | [optional] 
**SalesTaxAmount** | Pointer to **int64** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**TypeOfPurchase** | Pointer to **string** |  | [optional] 
**VatTaxRate** | Pointer to **float32** |  | [optional] 
**VehicleId** | Pointer to **string** |  | [optional] 

## Methods

### NewFleets

`func NewFleets() *Fleets`

NewFleets instantiates a new Fleets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetsWithDefaults

`func NewFleetsWithDefaults() *Fleets`

NewFleetsWithDefaults instantiates a new Fleets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverId

`func (o *Fleets) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *Fleets) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *Fleets) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *Fleets) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetExpandedFuelType

`func (o *Fleets) GetExpandedFuelType() string`

GetExpandedFuelType returns the ExpandedFuelType field if non-nil, zero value otherwise.

### GetExpandedFuelTypeOk

`func (o *Fleets) GetExpandedFuelTypeOk() (*string, bool)`

GetExpandedFuelTypeOk returns a tuple with the ExpandedFuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandedFuelType

`func (o *Fleets) SetExpandedFuelType(v string)`

SetExpandedFuelType sets ExpandedFuelType field to given value.

### HasExpandedFuelType

`func (o *Fleets) HasExpandedFuelType() bool`

HasExpandedFuelType returns a boolean if a field has been set.

### GetFleetNumber

`func (o *Fleets) GetFleetNumber() string`

GetFleetNumber returns the FleetNumber field if non-nil, zero value otherwise.

### GetFleetNumberOk

`func (o *Fleets) GetFleetNumberOk() (*string, bool)`

GetFleetNumberOk returns a tuple with the FleetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetNumber

`func (o *Fleets) SetFleetNumber(v string)`

SetFleetNumber sets FleetNumber field to given value.

### HasFleetNumber

`func (o *Fleets) HasFleetNumber() bool`

HasFleetNumber returns a boolean if a field has been set.

### GetFuelGrossAmount

`func (o *Fleets) GetFuelGrossAmount() int64`

GetFuelGrossAmount returns the FuelGrossAmount field if non-nil, zero value otherwise.

### GetFuelGrossAmountOk

`func (o *Fleets) GetFuelGrossAmountOk() (*int64, bool)`

GetFuelGrossAmountOk returns a tuple with the FuelGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelGrossAmount

`func (o *Fleets) SetFuelGrossAmount(v int64)`

SetFuelGrossAmount sets FuelGrossAmount field to given value.

### HasFuelGrossAmount

`func (o *Fleets) HasFuelGrossAmount() bool`

HasFuelGrossAmount returns a boolean if a field has been set.

### GetFuelNetAmount

`func (o *Fleets) GetFuelNetAmount() int64`

GetFuelNetAmount returns the FuelNetAmount field if non-nil, zero value otherwise.

### GetFuelNetAmountOk

`func (o *Fleets) GetFuelNetAmountOk() (*int64, bool)`

GetFuelNetAmountOk returns a tuple with the FuelNetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelNetAmount

`func (o *Fleets) SetFuelNetAmount(v int64)`

SetFuelNetAmount sets FuelNetAmount field to given value.

### HasFuelNetAmount

`func (o *Fleets) HasFuelNetAmount() bool`

HasFuelNetAmount returns a boolean if a field has been set.

### GetFuelProductQualifier

`func (o *Fleets) GetFuelProductQualifier() string`

GetFuelProductQualifier returns the FuelProductQualifier field if non-nil, zero value otherwise.

### GetFuelProductQualifierOk

`func (o *Fleets) GetFuelProductQualifierOk() (*string, bool)`

GetFuelProductQualifierOk returns a tuple with the FuelProductQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelProductQualifier

`func (o *Fleets) SetFuelProductQualifier(v string)`

SetFuelProductQualifier sets FuelProductQualifier field to given value.

### HasFuelProductQualifier

`func (o *Fleets) HasFuelProductQualifier() bool`

HasFuelProductQualifier returns a boolean if a field has been set.

### GetFuelPurchaseType

`func (o *Fleets) GetFuelPurchaseType() string`

GetFuelPurchaseType returns the FuelPurchaseType field if non-nil, zero value otherwise.

### GetFuelPurchaseTypeOk

`func (o *Fleets) GetFuelPurchaseTypeOk() (*string, bool)`

GetFuelPurchaseTypeOk returns a tuple with the FuelPurchaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelPurchaseType

`func (o *Fleets) SetFuelPurchaseType(v string)`

SetFuelPurchaseType sets FuelPurchaseType field to given value.

### HasFuelPurchaseType

`func (o *Fleets) HasFuelPurchaseType() bool`

HasFuelPurchaseType returns a boolean if a field has been set.

### GetFuelQuantity

`func (o *Fleets) GetFuelQuantity() float32`

GetFuelQuantity returns the FuelQuantity field if non-nil, zero value otherwise.

### GetFuelQuantityOk

`func (o *Fleets) GetFuelQuantityOk() (*float32, bool)`

GetFuelQuantityOk returns a tuple with the FuelQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelQuantity

`func (o *Fleets) SetFuelQuantity(v float32)`

SetFuelQuantity sets FuelQuantity field to given value.

### HasFuelQuantity

`func (o *Fleets) HasFuelQuantity() bool`

HasFuelQuantity returns a boolean if a field has been set.

### GetFuelServiceType

`func (o *Fleets) GetFuelServiceType() string`

GetFuelServiceType returns the FuelServiceType field if non-nil, zero value otherwise.

### GetFuelServiceTypeOk

`func (o *Fleets) GetFuelServiceTypeOk() (*string, bool)`

GetFuelServiceTypeOk returns a tuple with the FuelServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelServiceType

`func (o *Fleets) SetFuelServiceType(v string)`

SetFuelServiceType sets FuelServiceType field to given value.

### HasFuelServiceType

`func (o *Fleets) HasFuelServiceType() bool`

HasFuelServiceType returns a boolean if a field has been set.

### GetFuelTaxAmount

`func (o *Fleets) GetFuelTaxAmount() int64`

GetFuelTaxAmount returns the FuelTaxAmount field if non-nil, zero value otherwise.

### GetFuelTaxAmountOk

`func (o *Fleets) GetFuelTaxAmountOk() (*int64, bool)`

GetFuelTaxAmountOk returns a tuple with the FuelTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelTaxAmount

`func (o *Fleets) SetFuelTaxAmount(v int64)`

SetFuelTaxAmount sets FuelTaxAmount field to given value.

### HasFuelTaxAmount

`func (o *Fleets) HasFuelTaxAmount() bool`

HasFuelTaxAmount returns a boolean if a field has been set.

### GetFuelTaxExemptionStatus

`func (o *Fleets) GetFuelTaxExemptionStatus() string`

GetFuelTaxExemptionStatus returns the FuelTaxExemptionStatus field if non-nil, zero value otherwise.

### GetFuelTaxExemptionStatusOk

`func (o *Fleets) GetFuelTaxExemptionStatusOk() (*string, bool)`

GetFuelTaxExemptionStatusOk returns a tuple with the FuelTaxExemptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelTaxExemptionStatus

`func (o *Fleets) SetFuelTaxExemptionStatus(v string)`

SetFuelTaxExemptionStatus sets FuelTaxExemptionStatus field to given value.

### HasFuelTaxExemptionStatus

`func (o *Fleets) HasFuelTaxExemptionStatus() bool`

HasFuelTaxExemptionStatus returns a boolean if a field has been set.

### GetFuelType

`func (o *Fleets) GetFuelType() string`

GetFuelType returns the FuelType field if non-nil, zero value otherwise.

### GetFuelTypeOk

`func (o *Fleets) GetFuelTypeOk() (*string, bool)`

GetFuelTypeOk returns a tuple with the FuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelType

`func (o *Fleets) SetFuelType(v string)`

SetFuelType sets FuelType field to given value.

### HasFuelType

`func (o *Fleets) HasFuelType() bool`

HasFuelType returns a boolean if a field has been set.

### GetFuelUnitOfMeasure

`func (o *Fleets) GetFuelUnitOfMeasure() string`

GetFuelUnitOfMeasure returns the FuelUnitOfMeasure field if non-nil, zero value otherwise.

### GetFuelUnitOfMeasureOk

`func (o *Fleets) GetFuelUnitOfMeasureOk() (*string, bool)`

GetFuelUnitOfMeasureOk returns a tuple with the FuelUnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelUnitOfMeasure

`func (o *Fleets) SetFuelUnitOfMeasure(v string)`

SetFuelUnitOfMeasure sets FuelUnitOfMeasure field to given value.

### HasFuelUnitOfMeasure

`func (o *Fleets) HasFuelUnitOfMeasure() bool`

HasFuelUnitOfMeasure returns a boolean if a field has been set.

### GetFuelUnitPrice

`func (o *Fleets) GetFuelUnitPrice() float32`

GetFuelUnitPrice returns the FuelUnitPrice field if non-nil, zero value otherwise.

### GetFuelUnitPriceOk

`func (o *Fleets) GetFuelUnitPriceOk() (*float32, bool)`

GetFuelUnitPriceOk returns a tuple with the FuelUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelUnitPrice

`func (o *Fleets) SetFuelUnitPrice(v float32)`

SetFuelUnitPrice sets FuelUnitPrice field to given value.

### HasFuelUnitPrice

`func (o *Fleets) HasFuelUnitPrice() bool`

HasFuelUnitPrice returns a boolean if a field has been set.

### GetJobNumber

`func (o *Fleets) GetJobNumber() string`

GetJobNumber returns the JobNumber field if non-nil, zero value otherwise.

### GetJobNumberOk

`func (o *Fleets) GetJobNumberOk() (*string, bool)`

GetJobNumberOk returns a tuple with the JobNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobNumber

`func (o *Fleets) SetJobNumber(v string)`

SetJobNumber sets JobNumber field to given value.

### HasJobNumber

`func (o *Fleets) HasJobNumber() bool`

HasJobNumber returns a boolean if a field has been set.

### GetNonFuelGrossAmount

`func (o *Fleets) GetNonFuelGrossAmount() int64`

GetNonFuelGrossAmount returns the NonFuelGrossAmount field if non-nil, zero value otherwise.

### GetNonFuelGrossAmountOk

`func (o *Fleets) GetNonFuelGrossAmountOk() (*int64, bool)`

GetNonFuelGrossAmountOk returns a tuple with the NonFuelGrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFuelGrossAmount

`func (o *Fleets) SetNonFuelGrossAmount(v int64)`

SetNonFuelGrossAmount sets NonFuelGrossAmount field to given value.

### HasNonFuelGrossAmount

`func (o *Fleets) HasNonFuelGrossAmount() bool`

HasNonFuelGrossAmount returns a boolean if a field has been set.

### GetNonFuelItemDetails

`func (o *Fleets) GetNonFuelItemDetails() []NonFuelItemDetails`

GetNonFuelItemDetails returns the NonFuelItemDetails field if non-nil, zero value otherwise.

### GetNonFuelItemDetailsOk

`func (o *Fleets) GetNonFuelItemDetailsOk() (*[]NonFuelItemDetails, bool)`

GetNonFuelItemDetailsOk returns a tuple with the NonFuelItemDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFuelItemDetails

`func (o *Fleets) SetNonFuelItemDetails(v []NonFuelItemDetails)`

SetNonFuelItemDetails sets NonFuelItemDetails field to given value.

### HasNonFuelItemDetails

`func (o *Fleets) HasNonFuelItemDetails() bool`

HasNonFuelItemDetails returns a boolean if a field has been set.

### GetNonFuelNetAmount

`func (o *Fleets) GetNonFuelNetAmount() int64`

GetNonFuelNetAmount returns the NonFuelNetAmount field if non-nil, zero value otherwise.

### GetNonFuelNetAmountOk

`func (o *Fleets) GetNonFuelNetAmountOk() (*int64, bool)`

GetNonFuelNetAmountOk returns a tuple with the NonFuelNetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFuelNetAmount

`func (o *Fleets) SetNonFuelNetAmount(v int64)`

SetNonFuelNetAmount sets NonFuelNetAmount field to given value.

### HasNonFuelNetAmount

`func (o *Fleets) HasNonFuelNetAmount() bool`

HasNonFuelNetAmount returns a boolean if a field has been set.

### GetNonFuelTaxAmount

`func (o *Fleets) GetNonFuelTaxAmount() int64`

GetNonFuelTaxAmount returns the NonFuelTaxAmount field if non-nil, zero value otherwise.

### GetNonFuelTaxAmountOk

`func (o *Fleets) GetNonFuelTaxAmountOk() (*int64, bool)`

GetNonFuelTaxAmountOk returns a tuple with the NonFuelTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFuelTaxAmount

`func (o *Fleets) SetNonFuelTaxAmount(v int64)`

SetNonFuelTaxAmount sets NonFuelTaxAmount field to given value.

### HasNonFuelTaxAmount

`func (o *Fleets) HasNonFuelTaxAmount() bool`

HasNonFuelTaxAmount returns a boolean if a field has been set.

### GetNonFuelTaxExemptionStatus

`func (o *Fleets) GetNonFuelTaxExemptionStatus() string`

GetNonFuelTaxExemptionStatus returns the NonFuelTaxExemptionStatus field if non-nil, zero value otherwise.

### GetNonFuelTaxExemptionStatusOk

`func (o *Fleets) GetNonFuelTaxExemptionStatusOk() (*string, bool)`

GetNonFuelTaxExemptionStatusOk returns a tuple with the NonFuelTaxExemptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFuelTaxExemptionStatus

`func (o *Fleets) SetNonFuelTaxExemptionStatus(v string)`

SetNonFuelTaxExemptionStatus sets NonFuelTaxExemptionStatus field to given value.

### HasNonFuelTaxExemptionStatus

`func (o *Fleets) HasNonFuelTaxExemptionStatus() bool`

HasNonFuelTaxExemptionStatus returns a boolean if a field has been set.

### GetOdometerReading

`func (o *Fleets) GetOdometerReading() string`

GetOdometerReading returns the OdometerReading field if non-nil, zero value otherwise.

### GetOdometerReadingOk

`func (o *Fleets) GetOdometerReadingOk() (*string, bool)`

GetOdometerReadingOk returns a tuple with the OdometerReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerReading

`func (o *Fleets) SetOdometerReading(v string)`

SetOdometerReading sets OdometerReading field to given value.

### HasOdometerReading

`func (o *Fleets) HasOdometerReading() bool`

HasOdometerReading returns a boolean if a field has been set.

### GetSalesTaxAmount

`func (o *Fleets) GetSalesTaxAmount() int64`

GetSalesTaxAmount returns the SalesTaxAmount field if non-nil, zero value otherwise.

### GetSalesTaxAmountOk

`func (o *Fleets) GetSalesTaxAmountOk() (*int64, bool)`

GetSalesTaxAmountOk returns a tuple with the SalesTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxAmount

`func (o *Fleets) SetSalesTaxAmount(v int64)`

SetSalesTaxAmount sets SalesTaxAmount field to given value.

### HasSalesTaxAmount

`func (o *Fleets) HasSalesTaxAmount() bool`

HasSalesTaxAmount returns a boolean if a field has been set.

### GetServiceType

`func (o *Fleets) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Fleets) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Fleets) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Fleets) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetTypeOfPurchase

`func (o *Fleets) GetTypeOfPurchase() string`

GetTypeOfPurchase returns the TypeOfPurchase field if non-nil, zero value otherwise.

### GetTypeOfPurchaseOk

`func (o *Fleets) GetTypeOfPurchaseOk() (*string, bool)`

GetTypeOfPurchaseOk returns a tuple with the TypeOfPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfPurchase

`func (o *Fleets) SetTypeOfPurchase(v string)`

SetTypeOfPurchase sets TypeOfPurchase field to given value.

### HasTypeOfPurchase

`func (o *Fleets) HasTypeOfPurchase() bool`

HasTypeOfPurchase returns a boolean if a field has been set.

### GetVatTaxRate

`func (o *Fleets) GetVatTaxRate() float32`

GetVatTaxRate returns the VatTaxRate field if non-nil, zero value otherwise.

### GetVatTaxRateOk

`func (o *Fleets) GetVatTaxRateOk() (*float32, bool)`

GetVatTaxRateOk returns a tuple with the VatTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatTaxRate

`func (o *Fleets) SetVatTaxRate(v float32)`

SetVatTaxRate sets VatTaxRate field to given value.

### HasVatTaxRate

`func (o *Fleets) HasVatTaxRate() bool`

HasVatTaxRate returns a boolean if a field has been set.

### GetVehicleId

`func (o *Fleets) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *Fleets) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *Fleets) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *Fleets) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


