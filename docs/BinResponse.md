# BinResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRangeLength** | **int32** | Account range length | [default to null]
**BankId** | **int32** | The bank ID | [default to null]
**BillingIca** | **string** | The ICA to which fees will be billed | [default to null]
**Bin** | **string** | The bin number | [default to null]
**BinStatus** | [***BinStatus**](bin_status.md) |  | [default to null]
**BrandProductCode** | **string** |  | [default to null]
**CardBrand** | [***CardBrand**](card_brand.md) |  | [default to null]
**CardCategory** | [***CardCategory**](card_category.md) |  | [default to null]
**CardProductType** | [***CardProductType**](card_product_type.md) |  | [default to null]
**Country** | **string** |  | [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | The timestamp representing when the bin was created | [default to null]
**Currency** | **string** |  | [default to null]
**DigitalWalletActive** | **bool** | Determines if bin supports digital wallet tokenization | [optional] [default to false]
**EndDate** | [**time.Time**](time.Time.md) | The time when bin is decommissioned | [optional] [default to null]
**IcaBid** | **string** | ICA/BID | [default to null]
**Id** | **string** | Bin ID | [default to null]
**IsTokenizationEnabled** | **bool** | Controls whether bin allows tokenization | [default to false]
**LastModifiedTime** | [**time.Time**](time.Time.md) | The timestamp representing when the bin was last modified | [default to null]
**PanUtilization** | **int32** | Pan utilization | [default to null]
**PartnerId** | **int32** | The partner ID | [default to null]
**PhysicalCardFormat** | [***PhysicalCardFormat**](physical_card_format.md) |  | [optional] [default to null]
**Processor** | **string** | The name of the card processor | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The time when bin goes live | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

