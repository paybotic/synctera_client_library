# Lookup3dsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChallengePayload** | Pointer to **string** | 3DS challenge payload, returned if challenge is required | [optional] 
**ChallengeUrl** | Pointer to **string** | 3DS challenge URL, returned if challenge is required | [optional] 
**Id** | **string** | The unique identifier of the 3DS authentication | 
**ProcessorTransactionId** | Pointer to **string** | Processor Transaction ID, returned if challenge is required | [optional] 
**Status** | **string** | Status of the 3DS authentication | 

## Methods

### NewLookup3dsResponse

`func NewLookup3dsResponse(id string, status string, ) *Lookup3dsResponse`

NewLookup3dsResponse instantiates a new Lookup3dsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookup3dsResponseWithDefaults

`func NewLookup3dsResponseWithDefaults() *Lookup3dsResponse`

NewLookup3dsResponseWithDefaults instantiates a new Lookup3dsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallengePayload

`func (o *Lookup3dsResponse) GetChallengePayload() string`

GetChallengePayload returns the ChallengePayload field if non-nil, zero value otherwise.

### GetChallengePayloadOk

`func (o *Lookup3dsResponse) GetChallengePayloadOk() (*string, bool)`

GetChallengePayloadOk returns a tuple with the ChallengePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengePayload

`func (o *Lookup3dsResponse) SetChallengePayload(v string)`

SetChallengePayload sets ChallengePayload field to given value.

### HasChallengePayload

`func (o *Lookup3dsResponse) HasChallengePayload() bool`

HasChallengePayload returns a boolean if a field has been set.

### GetChallengeUrl

`func (o *Lookup3dsResponse) GetChallengeUrl() string`

GetChallengeUrl returns the ChallengeUrl field if non-nil, zero value otherwise.

### GetChallengeUrlOk

`func (o *Lookup3dsResponse) GetChallengeUrlOk() (*string, bool)`

GetChallengeUrlOk returns a tuple with the ChallengeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeUrl

`func (o *Lookup3dsResponse) SetChallengeUrl(v string)`

SetChallengeUrl sets ChallengeUrl field to given value.

### HasChallengeUrl

`func (o *Lookup3dsResponse) HasChallengeUrl() bool`

HasChallengeUrl returns a boolean if a field has been set.

### GetId

`func (o *Lookup3dsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Lookup3dsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Lookup3dsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetProcessorTransactionId

`func (o *Lookup3dsResponse) GetProcessorTransactionId() string`

GetProcessorTransactionId returns the ProcessorTransactionId field if non-nil, zero value otherwise.

### GetProcessorTransactionIdOk

`func (o *Lookup3dsResponse) GetProcessorTransactionIdOk() (*string, bool)`

GetProcessorTransactionIdOk returns a tuple with the ProcessorTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorTransactionId

`func (o *Lookup3dsResponse) SetProcessorTransactionId(v string)`

SetProcessorTransactionId sets ProcessorTransactionId field to given value.

### HasProcessorTransactionId

`func (o *Lookup3dsResponse) HasProcessorTransactionId() bool`

HasProcessorTransactionId returns a boolean if a field has been set.

### GetStatus

`func (o *Lookup3dsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Lookup3dsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Lookup3dsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


