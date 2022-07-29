# RiskRating

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationId** | **string** | The risk configuration id used in risk score calculation | [optional] [default to null]
**Id** | **string** | Risk rating ID | [optional] [default to null]
**NextReview** | [**time.Time**](time.Time.md) | The next review date where customer risk will be calculated | [optional] [default to null]
**RiskLevel** | **string** | A textual representation of the customer risk score | [optional] [default to null]
**RiskReview** | [**time.Time**](time.Time.md) | The date the customer risk rating was calculated | [optional] [default to null]
**RiskScore** | **float32** | The cumulative score of all risk rating fields | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

