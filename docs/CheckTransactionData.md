# CheckTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackImageId** | Pointer to **string** | The back image id of the captured check | [optional] 
**FrontImageId** | Pointer to **string** | The front image id of the captured check | [optional] 
**Id** | Pointer to **string** | The Synctera check deposit ID | [optional] 

## Methods

### NewCheckTransactionData

`func NewCheckTransactionData() *CheckTransactionData`

NewCheckTransactionData instantiates a new CheckTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTransactionDataWithDefaults

`func NewCheckTransactionDataWithDefaults() *CheckTransactionData`

NewCheckTransactionDataWithDefaults instantiates a new CheckTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackImageId

`func (o *CheckTransactionData) GetBackImageId() string`

GetBackImageId returns the BackImageId field if non-nil, zero value otherwise.

### GetBackImageIdOk

`func (o *CheckTransactionData) GetBackImageIdOk() (*string, bool)`

GetBackImageIdOk returns a tuple with the BackImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackImageId

`func (o *CheckTransactionData) SetBackImageId(v string)`

SetBackImageId sets BackImageId field to given value.

### HasBackImageId

`func (o *CheckTransactionData) HasBackImageId() bool`

HasBackImageId returns a boolean if a field has been set.

### GetFrontImageId

`func (o *CheckTransactionData) GetFrontImageId() string`

GetFrontImageId returns the FrontImageId field if non-nil, zero value otherwise.

### GetFrontImageIdOk

`func (o *CheckTransactionData) GetFrontImageIdOk() (*string, bool)`

GetFrontImageIdOk returns a tuple with the FrontImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImageId

`func (o *CheckTransactionData) SetFrontImageId(v string)`

SetFrontImageId sets FrontImageId field to given value.

### HasFrontImageId

`func (o *CheckTransactionData) HasFrontImageId() bool`

HasFrontImageId returns a boolean if a field has been set.

### GetId

`func (o *CheckTransactionData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckTransactionData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckTransactionData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckTransactionData) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


