# CardProduct

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | [default to null]
**Active** | **bool** | Indicates whether the Card Product is active | [default to null]
**CardProgramId** | **string** | Card Program ID | [default to null]
**Color** | **string** | Color code for dynamic card elements such as PAN and card holder name | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | The timestamp representing when the Card Product was created | [optional] [default to null]
**DigitalWalletTokenization** | [***DigitalWalletTokenization**](digital_wallet_tokenization.md) |  | [optional] [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The time when the Card Product is decommissioned | [optional] [default to null]
**GatewayId** | **string** | Gateway ID, used if the Card Product is utilizing FinTech authorization flow | [optional] [default to null]
**Id** | **string** | Card Product ID | [optional] [default to null]
**Image** | **bool** | Indicates whether or not there is an overlay image of the card product available | [optional] [default to null]
**ImageMode** | [***CardImageMode**](card_image_mode.md) |  | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | The timestamp representing when the Card Product was last modified | [optional] [default to null]
**Name** | **string** | The name of the Card Product | [default to null]
**Orientation** | **string** | Card orientation | [optional] [default to null]
**PackageId** | **string** | Card fulfillment provider’s package ID | [optional] [default to null]
**PhysicalCardFormat** | [***PhysicalCardFormat**](physical_card_format.md) |  | [optional] [default to null]
**ReturnAddress** | [***Shipping**](shipping.md) |  | [optional] [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The time when the Card Product goes live | [default to null]
**TxnEnhancer** | [***TxnEnhancer**](txn_enhancer.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

