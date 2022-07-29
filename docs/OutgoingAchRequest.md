# OutgoingAchRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount to transfer in ISO 4217 minor currency units | [default to null]
**CompanyEntryDescription** | **string** | Company entry description ACH field. Originator inserts this field&#x27;s value to provide the Receiver with a description of the entry&#x27;s purpose. | [optional] [default to null]
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | [default to null]
**CustomerId** | **string** | The customer&#x27;s unique identifier | [default to null]
**DcSign** | **string** | The type of transaction (debit or credit). A debit is a transfer in and a credit is a transfer out of the originating account | [default to null]
**EffectiveDate** | **string** | Effective date transaction proccesses (is_same_day needs to be false or not present at all) | [optional] [default to null]
**ExternalData** | [***interface{}**](interface{}.md) | Additional transfer metadata structured as key-value pairs | [optional] [default to null]
**FinalCustomerId** | **string** | ID of the international customer that receives the final remittance transfer (required for OFAC enabled payments) | [optional] [default to null]
**Hold** | [***AchRequestHoldData**](ach_request_hold_data.md) |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**IsSameDay** | **bool** | Send as same day ACH transaction (use only is_same_day without specific effective_date) | [optional] [default to null]
**Memo** | **string** | Memo for the payment | [optional] [default to null]
**OriginatingAccountId** | **string** | The unique identifier for an originating account | [default to null]
**ReceivingAccountId** | **string** | The unique identifier for an receiving account | [default to null]
**ReferenceInfo** | **string** | Will be sent to the ACH network and maps to Addenda record 05 - the recipient bank will receive this info | [optional] [default to null]
**Risk** | [***RiskData**](risk_data.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

