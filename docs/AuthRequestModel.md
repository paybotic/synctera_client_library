# AuthRequestModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents) | [default to null]
**CardAcceptor** | [***CardAcceptorModel**](card_acceptor_model.md) |  | [optional] [default to null]
**CardId** | **string** |  | [default to null]
**CardOptions** | [***CardOptions**](card_options.md) |  | [optional] [default to null]
**CashBackAmount** | **int32** |  | [optional] [default to null]
**IsPreAuth** | **bool** |  | [optional] [default to false]
**Mid** | **string** |  | [default to null]
**NetworkFees** | [**[]NetworkFeeModel**](network_fee_model.md) |  | [optional] [default to null]
**Pin** | **string** |  | [optional] [default to null]
**TransactionOptions** | [***TransactionOptions**](transaction_options.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

