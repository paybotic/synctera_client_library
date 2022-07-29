# ResponseHistoryItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | Response code from the request | [optional] [default to null]
**ResponseBody** | **string** | Response body from the request(Length more than 1024 will be trimmed) | [optional] [default to null]
**ResponseTime** | [**time.Time**](time.Time.md) | Timestamp that the response is received | [optional] [default to null]
**SentTime** | [**time.Time**](time.Time.md) | Timestamp that the request is sent | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

