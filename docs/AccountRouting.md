# AccountRouting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchRoutingNumber** | **string** | The routing number used for US ACH payments. Only appears if &#x60;bank_countries&#x60; contains &#x60;US&#x60;. Value may be masked, in which case only the last four digits are returned.  | [optional] [default to null]
**BankCountries** | **[]string** | The countries that this bank operates the account in | [default to null]
**BankName** | **string** | The name of the bank managing the account | [default to null]
**EftRoutingNumber** | **string** | The routing number used for EFT payments, identifying a Canadian bank, consisting of the institution number and the branch number. Only appears if &#x60;bank_countries&#x60; contains &#x60;CA&#x60;. Value may be masked, in which case only the last four digits are returned.  | [optional] [default to null]
**SwiftCode** | **string** | The SWIFT code for the bank. Value may be masked, in which case only the last four characters are returned.  | [optional] [default to null]
**WireRoutingNumber** | **string** | The routing number used for domestic wire payments. Only appears if &#x60;bank_countries&#x60; contains &#x60;US&#x60;. Value may be masked, in which case only the last four digits are returned.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

