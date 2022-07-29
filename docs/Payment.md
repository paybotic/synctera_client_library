# Payment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | User provided description for the payment schedule | [optional] [default to null]
**ErrorDetails** | [***PaymentErrorDetails**](payment_error_details.md) |  | [optional] [default to null]
**Id** | **string** | Payment ID | [optional] [default to null]
**Metadata** | [***interface{}**](interface{}.md) | User provided JSON format data for the payment schedule | [optional] [default to null]
**PaymentDate** | [***PaymentDate**](payment_date.md) |  | [optional] [default to null]
**PaymentInstruction** | [***PaymentInstruction**](payment_instruction.md) |  | [optional] [default to null]
**PaymentScheduleId** | **string** | ID of the payment schedule that executed this payment | [optional] [default to null]
**Status** | [***PaymentStatus**](payment_status.md) |  | [optional] [default to null]
**TransactionId** | **string** | Transaction ID. It will be included only when status is COMPLETED | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

