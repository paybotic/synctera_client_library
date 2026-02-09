# InventoryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommodityCode** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DiscountAmount** | Pointer to **string** |  | [optional] 
**ItemDetailTaxes** | Pointer to [**[]ItemDetailTaxes**](ItemDetailTaxes.md) |  | [optional] 
**ItemDiscountAmount** | Pointer to **int32** |  | [optional] 
**ItemDiscountAmountIndicator** | Pointer to **string** |  | [optional] 
**ItemDiscountAppliedIndicator** | Pointer to **string** |  | [optional] 
**ItemDiscountRate** | Pointer to **int32** |  | [optional] 
**ItemExtendedAmount** | Pointer to **int32** |  | [optional] 
**ItemExtendedAmountIndicator** | Pointer to **string** |  | [optional] 
**ItemTotalAmount** | Pointer to **int32** |  | [optional] 
**ItemTotalAmountIndicator** | Pointer to **string** |  | [optional] 
**ItemVatAmount** | Pointer to **int32** |  | [optional] 
**ItemVatRate** | Pointer to **int32** |  | [optional] 
**ProductCode** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**TotalAmount** | Pointer to **string** |  | [optional] 
**UnitOfMeasure** | Pointer to **string** |  | [optional] 
**UnitPrice** | Pointer to **int32** |  | [optional] 
**VatTaxAmount** | Pointer to **string** |  | [optional] 
**VatTaxRate** | Pointer to **string** |  | [optional] 

## Methods

### NewInventoryDetails

`func NewInventoryDetails() *InventoryDetails`

NewInventoryDetails instantiates a new InventoryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDetailsWithDefaults

`func NewInventoryDetailsWithDefaults() *InventoryDetails`

NewInventoryDetailsWithDefaults instantiates a new InventoryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommodityCode

`func (o *InventoryDetails) GetCommodityCode() string`

GetCommodityCode returns the CommodityCode field if non-nil, zero value otherwise.

### GetCommodityCodeOk

`func (o *InventoryDetails) GetCommodityCodeOk() (*string, bool)`

GetCommodityCodeOk returns a tuple with the CommodityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommodityCode

`func (o *InventoryDetails) SetCommodityCode(v string)`

SetCommodityCode sets CommodityCode field to given value.

### HasCommodityCode

`func (o *InventoryDetails) HasCommodityCode() bool`

HasCommodityCode returns a boolean if a field has been set.

### GetDescription

`func (o *InventoryDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InventoryDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InventoryDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InventoryDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *InventoryDetails) GetDiscountAmount() string`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *InventoryDetails) GetDiscountAmountOk() (*string, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *InventoryDetails) SetDiscountAmount(v string)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *InventoryDetails) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetItemDetailTaxes

`func (o *InventoryDetails) GetItemDetailTaxes() []ItemDetailTaxes`

GetItemDetailTaxes returns the ItemDetailTaxes field if non-nil, zero value otherwise.

### GetItemDetailTaxesOk

`func (o *InventoryDetails) GetItemDetailTaxesOk() (*[]ItemDetailTaxes, bool)`

GetItemDetailTaxesOk returns a tuple with the ItemDetailTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDetailTaxes

`func (o *InventoryDetails) SetItemDetailTaxes(v []ItemDetailTaxes)`

SetItemDetailTaxes sets ItemDetailTaxes field to given value.

### HasItemDetailTaxes

`func (o *InventoryDetails) HasItemDetailTaxes() bool`

HasItemDetailTaxes returns a boolean if a field has been set.

### GetItemDiscountAmount

`func (o *InventoryDetails) GetItemDiscountAmount() int32`

GetItemDiscountAmount returns the ItemDiscountAmount field if non-nil, zero value otherwise.

### GetItemDiscountAmountOk

`func (o *InventoryDetails) GetItemDiscountAmountOk() (*int32, bool)`

GetItemDiscountAmountOk returns a tuple with the ItemDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDiscountAmount

`func (o *InventoryDetails) SetItemDiscountAmount(v int32)`

SetItemDiscountAmount sets ItemDiscountAmount field to given value.

### HasItemDiscountAmount

`func (o *InventoryDetails) HasItemDiscountAmount() bool`

HasItemDiscountAmount returns a boolean if a field has been set.

### GetItemDiscountAmountIndicator

`func (o *InventoryDetails) GetItemDiscountAmountIndicator() string`

GetItemDiscountAmountIndicator returns the ItemDiscountAmountIndicator field if non-nil, zero value otherwise.

### GetItemDiscountAmountIndicatorOk

`func (o *InventoryDetails) GetItemDiscountAmountIndicatorOk() (*string, bool)`

GetItemDiscountAmountIndicatorOk returns a tuple with the ItemDiscountAmountIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDiscountAmountIndicator

`func (o *InventoryDetails) SetItemDiscountAmountIndicator(v string)`

SetItemDiscountAmountIndicator sets ItemDiscountAmountIndicator field to given value.

### HasItemDiscountAmountIndicator

`func (o *InventoryDetails) HasItemDiscountAmountIndicator() bool`

HasItemDiscountAmountIndicator returns a boolean if a field has been set.

### GetItemDiscountAppliedIndicator

`func (o *InventoryDetails) GetItemDiscountAppliedIndicator() string`

GetItemDiscountAppliedIndicator returns the ItemDiscountAppliedIndicator field if non-nil, zero value otherwise.

### GetItemDiscountAppliedIndicatorOk

`func (o *InventoryDetails) GetItemDiscountAppliedIndicatorOk() (*string, bool)`

GetItemDiscountAppliedIndicatorOk returns a tuple with the ItemDiscountAppliedIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDiscountAppliedIndicator

`func (o *InventoryDetails) SetItemDiscountAppliedIndicator(v string)`

SetItemDiscountAppliedIndicator sets ItemDiscountAppliedIndicator field to given value.

### HasItemDiscountAppliedIndicator

`func (o *InventoryDetails) HasItemDiscountAppliedIndicator() bool`

HasItemDiscountAppliedIndicator returns a boolean if a field has been set.

### GetItemDiscountRate

`func (o *InventoryDetails) GetItemDiscountRate() int32`

GetItemDiscountRate returns the ItemDiscountRate field if non-nil, zero value otherwise.

### GetItemDiscountRateOk

`func (o *InventoryDetails) GetItemDiscountRateOk() (*int32, bool)`

GetItemDiscountRateOk returns a tuple with the ItemDiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDiscountRate

`func (o *InventoryDetails) SetItemDiscountRate(v int32)`

SetItemDiscountRate sets ItemDiscountRate field to given value.

### HasItemDiscountRate

`func (o *InventoryDetails) HasItemDiscountRate() bool`

HasItemDiscountRate returns a boolean if a field has been set.

### GetItemExtendedAmount

`func (o *InventoryDetails) GetItemExtendedAmount() int32`

GetItemExtendedAmount returns the ItemExtendedAmount field if non-nil, zero value otherwise.

### GetItemExtendedAmountOk

`func (o *InventoryDetails) GetItemExtendedAmountOk() (*int32, bool)`

GetItemExtendedAmountOk returns a tuple with the ItemExtendedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemExtendedAmount

`func (o *InventoryDetails) SetItemExtendedAmount(v int32)`

SetItemExtendedAmount sets ItemExtendedAmount field to given value.

### HasItemExtendedAmount

`func (o *InventoryDetails) HasItemExtendedAmount() bool`

HasItemExtendedAmount returns a boolean if a field has been set.

### GetItemExtendedAmountIndicator

`func (o *InventoryDetails) GetItemExtendedAmountIndicator() string`

GetItemExtendedAmountIndicator returns the ItemExtendedAmountIndicator field if non-nil, zero value otherwise.

### GetItemExtendedAmountIndicatorOk

`func (o *InventoryDetails) GetItemExtendedAmountIndicatorOk() (*string, bool)`

GetItemExtendedAmountIndicatorOk returns a tuple with the ItemExtendedAmountIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemExtendedAmountIndicator

`func (o *InventoryDetails) SetItemExtendedAmountIndicator(v string)`

SetItemExtendedAmountIndicator sets ItemExtendedAmountIndicator field to given value.

### HasItemExtendedAmountIndicator

`func (o *InventoryDetails) HasItemExtendedAmountIndicator() bool`

HasItemExtendedAmountIndicator returns a boolean if a field has been set.

### GetItemTotalAmount

`func (o *InventoryDetails) GetItemTotalAmount() int32`

GetItemTotalAmount returns the ItemTotalAmount field if non-nil, zero value otherwise.

### GetItemTotalAmountOk

`func (o *InventoryDetails) GetItemTotalAmountOk() (*int32, bool)`

GetItemTotalAmountOk returns a tuple with the ItemTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTotalAmount

`func (o *InventoryDetails) SetItemTotalAmount(v int32)`

SetItemTotalAmount sets ItemTotalAmount field to given value.

### HasItemTotalAmount

`func (o *InventoryDetails) HasItemTotalAmount() bool`

HasItemTotalAmount returns a boolean if a field has been set.

### GetItemTotalAmountIndicator

`func (o *InventoryDetails) GetItemTotalAmountIndicator() string`

GetItemTotalAmountIndicator returns the ItemTotalAmountIndicator field if non-nil, zero value otherwise.

### GetItemTotalAmountIndicatorOk

`func (o *InventoryDetails) GetItemTotalAmountIndicatorOk() (*string, bool)`

GetItemTotalAmountIndicatorOk returns a tuple with the ItemTotalAmountIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTotalAmountIndicator

`func (o *InventoryDetails) SetItemTotalAmountIndicator(v string)`

SetItemTotalAmountIndicator sets ItemTotalAmountIndicator field to given value.

### HasItemTotalAmountIndicator

`func (o *InventoryDetails) HasItemTotalAmountIndicator() bool`

HasItemTotalAmountIndicator returns a boolean if a field has been set.

### GetItemVatAmount

`func (o *InventoryDetails) GetItemVatAmount() int32`

GetItemVatAmount returns the ItemVatAmount field if non-nil, zero value otherwise.

### GetItemVatAmountOk

`func (o *InventoryDetails) GetItemVatAmountOk() (*int32, bool)`

GetItemVatAmountOk returns a tuple with the ItemVatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVatAmount

`func (o *InventoryDetails) SetItemVatAmount(v int32)`

SetItemVatAmount sets ItemVatAmount field to given value.

### HasItemVatAmount

`func (o *InventoryDetails) HasItemVatAmount() bool`

HasItemVatAmount returns a boolean if a field has been set.

### GetItemVatRate

`func (o *InventoryDetails) GetItemVatRate() int32`

GetItemVatRate returns the ItemVatRate field if non-nil, zero value otherwise.

### GetItemVatRateOk

`func (o *InventoryDetails) GetItemVatRateOk() (*int32, bool)`

GetItemVatRateOk returns a tuple with the ItemVatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVatRate

`func (o *InventoryDetails) SetItemVatRate(v int32)`

SetItemVatRate sets ItemVatRate field to given value.

### HasItemVatRate

`func (o *InventoryDetails) HasItemVatRate() bool`

HasItemVatRate returns a boolean if a field has been set.

### GetProductCode

`func (o *InventoryDetails) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *InventoryDetails) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *InventoryDetails) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *InventoryDetails) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetQuantity

`func (o *InventoryDetails) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InventoryDetails) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InventoryDetails) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InventoryDetails) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTotalAmount

`func (o *InventoryDetails) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *InventoryDetails) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *InventoryDetails) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *InventoryDetails) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUnitOfMeasure

`func (o *InventoryDetails) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *InventoryDetails) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *InventoryDetails) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *InventoryDetails) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetUnitPrice

`func (o *InventoryDetails) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InventoryDetails) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InventoryDetails) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *InventoryDetails) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetVatTaxAmount

`func (o *InventoryDetails) GetVatTaxAmount() string`

GetVatTaxAmount returns the VatTaxAmount field if non-nil, zero value otherwise.

### GetVatTaxAmountOk

`func (o *InventoryDetails) GetVatTaxAmountOk() (*string, bool)`

GetVatTaxAmountOk returns a tuple with the VatTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatTaxAmount

`func (o *InventoryDetails) SetVatTaxAmount(v string)`

SetVatTaxAmount sets VatTaxAmount field to given value.

### HasVatTaxAmount

`func (o *InventoryDetails) HasVatTaxAmount() bool`

HasVatTaxAmount returns a boolean if a field has been set.

### GetVatTaxRate

`func (o *InventoryDetails) GetVatTaxRate() string`

GetVatTaxRate returns the VatTaxRate field if non-nil, zero value otherwise.

### GetVatTaxRateOk

`func (o *InventoryDetails) GetVatTaxRateOk() (*string, bool)`

GetVatTaxRateOk returns a tuple with the VatTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatTaxRate

`func (o *InventoryDetails) SetVatTaxRate(v string)`

SetVatTaxRate sets VatTaxRate field to given value.

### HasVatTaxRate

`func (o *InventoryDetails) HasVatTaxRate() bool`

HasVatTaxRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


