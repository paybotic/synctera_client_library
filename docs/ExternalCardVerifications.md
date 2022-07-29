# ExternalCardVerifications

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressVerificationResults** | **string** | Address verification results  Status | Description --- | --- VERIFIED | Address verified NOT_VERIFIED | Address not verified ADDRESS_NO_MATCH | Postal/ZIP match, street addresses do not match or street address not included in request  | [optional] [default to null]
**CardType** | **string** | Indicates whether the external card is credit, debit, prepaid, deferred debit, or charge. | [optional] [default to null]
**Cvv2Result** | **string** | Card Verification Value results  Status | Description --- | --- VERIFIED | CVV and expiration dates verified INCORRECT | Either CVV or expiration date is incorrect NOT_SUPPORTED |  Issuer does not participate in CVV2 service  | [optional] [default to null]
**FastFundsIndicator** | **bool** | Indicates if card is Fast Funds eligible (i.e. if the funds will settle in 30 mins or less). If not eligible, typically funds will settle within 2 business days. | [optional] [default to null]
**OnlineGamblingBlockIndicator** | **bool** | Indicates if the card can receive push-payments for online gambling payouts. | [optional] [default to null]
**Processor** | **string** | The name of the processor | [optional] [default to null]
**PushFundsBlockIndicator** | **bool** | Indicates if the associated card can receive push-to-card disbursements. | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

