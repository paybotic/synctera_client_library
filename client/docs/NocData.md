# NocData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeCode** | **string** | Change code, indicating which part of the original entry is to be corrected. | 
**CorrectedData** | **string** | Corrected information from the original entry (e.g. correct account number). | 
**OriginalDfiNo** | **string** | Receiving financial institution of the original entry. | 
**OriginalTrace** | **string** | Trace number of the original entry that is being corrected. | 

## Methods

### NewNocData

`func NewNocData(changeCode string, correctedData string, originalDfiNo string, originalTrace string, ) *NocData`

NewNocData instantiates a new NocData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNocDataWithDefaults

`func NewNocDataWithDefaults() *NocData`

NewNocDataWithDefaults instantiates a new NocData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeCode

`func (o *NocData) GetChangeCode() string`

GetChangeCode returns the ChangeCode field if non-nil, zero value otherwise.

### GetChangeCodeOk

`func (o *NocData) GetChangeCodeOk() (*string, bool)`

GetChangeCodeOk returns a tuple with the ChangeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeCode

`func (o *NocData) SetChangeCode(v string)`

SetChangeCode sets ChangeCode field to given value.


### GetCorrectedData

`func (o *NocData) GetCorrectedData() string`

GetCorrectedData returns the CorrectedData field if non-nil, zero value otherwise.

### GetCorrectedDataOk

`func (o *NocData) GetCorrectedDataOk() (*string, bool)`

GetCorrectedDataOk returns a tuple with the CorrectedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectedData

`func (o *NocData) SetCorrectedData(v string)`

SetCorrectedData sets CorrectedData field to given value.


### GetOriginalDfiNo

`func (o *NocData) GetOriginalDfiNo() string`

GetOriginalDfiNo returns the OriginalDfiNo field if non-nil, zero value otherwise.

### GetOriginalDfiNoOk

`func (o *NocData) GetOriginalDfiNoOk() (*string, bool)`

GetOriginalDfiNoOk returns a tuple with the OriginalDfiNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDfiNo

`func (o *NocData) SetOriginalDfiNo(v string)`

SetOriginalDfiNo sets OriginalDfiNo field to given value.


### GetOriginalTrace

`func (o *NocData) GetOriginalTrace() string`

GetOriginalTrace returns the OriginalTrace field if non-nil, zero value otherwise.

### GetOriginalTraceOk

`func (o *NocData) GetOriginalTraceOk() (*string, bool)`

GetOriginalTraceOk returns a tuple with the OriginalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTrace

`func (o *NocData) SetOriginalTrace(v string)`

SetOriginalTrace sets OriginalTrace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


