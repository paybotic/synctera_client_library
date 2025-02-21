# ReturnData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Return reason code, i.e. why is the entry being returned. | 
**DishonoredReturnCode** | Pointer to **string** | Code of the original dishonored return (without \&quot;R\&quot;). Filled only if this is a contested return. | [optional] 
**DishonoredReturnSettlementDate** | Pointer to **string** | Settlement date of the original dishonored return. Filled only if this is a contested return. Formatted as an ordinal date, a single day-of-year number between 1-366. | [optional] 
**DishonoredReturnTrace** | Pointer to **string** | Trace number of the original dishonored return. Filled only if this is a contested return. | [optional] 
**FieldErrors** | Pointer to **string** | Required for return reason code R69. Contains the code(s) to indicate the field(s) in which erroneous information in the original return is located. | [optional] 
**OriginalDfiNo** | **string** | Receiving financial institution of the original entry. | 
**OriginalTrace** | **string** | Trace number of the original entry that is being returned. | 
**ReturnCode** | Pointer to **string** | Return reason code of the original return (just the number). Filled only if this is a dishonored return. | [optional] 
**ReturnSettlementDate** | Pointer to **string** | Settlement date of the original return. Filled only if this is a dishonored return. Formatted as an ordinal date, a single day-of-year number between 1-366. | [optional] 
**ReturnTrace** | Pointer to **string** | Trace number of the original return. Filled only if this is a dishonored return. | [optional] 

## Methods

### NewReturnData

`func NewReturnData(code string, originalDfiNo string, originalTrace string, ) *ReturnData`

NewReturnData instantiates a new ReturnData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnDataWithDefaults

`func NewReturnDataWithDefaults() *ReturnData`

NewReturnDataWithDefaults instantiates a new ReturnData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ReturnData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReturnData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReturnData) SetCode(v string)`

SetCode sets Code field to given value.


### GetDishonoredReturnCode

`func (o *ReturnData) GetDishonoredReturnCode() string`

GetDishonoredReturnCode returns the DishonoredReturnCode field if non-nil, zero value otherwise.

### GetDishonoredReturnCodeOk

`func (o *ReturnData) GetDishonoredReturnCodeOk() (*string, bool)`

GetDishonoredReturnCodeOk returns a tuple with the DishonoredReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDishonoredReturnCode

`func (o *ReturnData) SetDishonoredReturnCode(v string)`

SetDishonoredReturnCode sets DishonoredReturnCode field to given value.

### HasDishonoredReturnCode

`func (o *ReturnData) HasDishonoredReturnCode() bool`

HasDishonoredReturnCode returns a boolean if a field has been set.

### GetDishonoredReturnSettlementDate

`func (o *ReturnData) GetDishonoredReturnSettlementDate() string`

GetDishonoredReturnSettlementDate returns the DishonoredReturnSettlementDate field if non-nil, zero value otherwise.

### GetDishonoredReturnSettlementDateOk

`func (o *ReturnData) GetDishonoredReturnSettlementDateOk() (*string, bool)`

GetDishonoredReturnSettlementDateOk returns a tuple with the DishonoredReturnSettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDishonoredReturnSettlementDate

`func (o *ReturnData) SetDishonoredReturnSettlementDate(v string)`

SetDishonoredReturnSettlementDate sets DishonoredReturnSettlementDate field to given value.

### HasDishonoredReturnSettlementDate

`func (o *ReturnData) HasDishonoredReturnSettlementDate() bool`

HasDishonoredReturnSettlementDate returns a boolean if a field has been set.

### GetDishonoredReturnTrace

`func (o *ReturnData) GetDishonoredReturnTrace() string`

GetDishonoredReturnTrace returns the DishonoredReturnTrace field if non-nil, zero value otherwise.

### GetDishonoredReturnTraceOk

`func (o *ReturnData) GetDishonoredReturnTraceOk() (*string, bool)`

GetDishonoredReturnTraceOk returns a tuple with the DishonoredReturnTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDishonoredReturnTrace

`func (o *ReturnData) SetDishonoredReturnTrace(v string)`

SetDishonoredReturnTrace sets DishonoredReturnTrace field to given value.

### HasDishonoredReturnTrace

`func (o *ReturnData) HasDishonoredReturnTrace() bool`

HasDishonoredReturnTrace returns a boolean if a field has been set.

### GetFieldErrors

`func (o *ReturnData) GetFieldErrors() string`

GetFieldErrors returns the FieldErrors field if non-nil, zero value otherwise.

### GetFieldErrorsOk

`func (o *ReturnData) GetFieldErrorsOk() (*string, bool)`

GetFieldErrorsOk returns a tuple with the FieldErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldErrors

`func (o *ReturnData) SetFieldErrors(v string)`

SetFieldErrors sets FieldErrors field to given value.

### HasFieldErrors

`func (o *ReturnData) HasFieldErrors() bool`

HasFieldErrors returns a boolean if a field has been set.

### GetOriginalDfiNo

`func (o *ReturnData) GetOriginalDfiNo() string`

GetOriginalDfiNo returns the OriginalDfiNo field if non-nil, zero value otherwise.

### GetOriginalDfiNoOk

`func (o *ReturnData) GetOriginalDfiNoOk() (*string, bool)`

GetOriginalDfiNoOk returns a tuple with the OriginalDfiNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDfiNo

`func (o *ReturnData) SetOriginalDfiNo(v string)`

SetOriginalDfiNo sets OriginalDfiNo field to given value.


### GetOriginalTrace

`func (o *ReturnData) GetOriginalTrace() string`

GetOriginalTrace returns the OriginalTrace field if non-nil, zero value otherwise.

### GetOriginalTraceOk

`func (o *ReturnData) GetOriginalTraceOk() (*string, bool)`

GetOriginalTraceOk returns a tuple with the OriginalTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTrace

`func (o *ReturnData) SetOriginalTrace(v string)`

SetOriginalTrace sets OriginalTrace field to given value.


### GetReturnCode

`func (o *ReturnData) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ReturnData) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ReturnData) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *ReturnData) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetReturnSettlementDate

`func (o *ReturnData) GetReturnSettlementDate() string`

GetReturnSettlementDate returns the ReturnSettlementDate field if non-nil, zero value otherwise.

### GetReturnSettlementDateOk

`func (o *ReturnData) GetReturnSettlementDateOk() (*string, bool)`

GetReturnSettlementDateOk returns a tuple with the ReturnSettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSettlementDate

`func (o *ReturnData) SetReturnSettlementDate(v string)`

SetReturnSettlementDate sets ReturnSettlementDate field to given value.

### HasReturnSettlementDate

`func (o *ReturnData) HasReturnSettlementDate() bool`

HasReturnSettlementDate returns a boolean if a field has been set.

### GetReturnTrace

`func (o *ReturnData) GetReturnTrace() string`

GetReturnTrace returns the ReturnTrace field if non-nil, zero value otherwise.

### GetReturnTraceOk

`func (o *ReturnData) GetReturnTraceOk() (*string, bool)`

GetReturnTraceOk returns a tuple with the ReturnTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTrace

`func (o *ReturnData) SetReturnTrace(v string)`

SetReturnTrace sets ReturnTrace field to given value.

### HasReturnTrace

`func (o *ReturnData) HasReturnTrace() bool`

HasReturnTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


