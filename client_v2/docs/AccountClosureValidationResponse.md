# AccountClosureValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | **map[string]interface{}** |  | 
**Message** | **string** | Validation message | 
**Name** | **string** | Name | 
**Validated** | **bool** | Validation result | 

## Methods

### NewAccountClosureValidationResponse

`func NewAccountClosureValidationResponse(details map[string]interface{}, message string, name string, validated bool, ) *AccountClosureValidationResponse`

NewAccountClosureValidationResponse instantiates a new AccountClosureValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountClosureValidationResponseWithDefaults

`func NewAccountClosureValidationResponseWithDefaults() *AccountClosureValidationResponse`

NewAccountClosureValidationResponseWithDefaults instantiates a new AccountClosureValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *AccountClosureValidationResponse) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AccountClosureValidationResponse) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AccountClosureValidationResponse) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.


### GetMessage

`func (o *AccountClosureValidationResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AccountClosureValidationResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AccountClosureValidationResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetName

`func (o *AccountClosureValidationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountClosureValidationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountClosureValidationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetValidated

`func (o *AccountClosureValidationResponse) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *AccountClosureValidationResponse) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *AccountClosureValidationResponse) SetValidated(v bool)`

SetValidated sets Validated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


