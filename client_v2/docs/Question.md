# Question

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | **string** | The answer | 
**Question** | **string** | The question | 
**Section** | Pointer to **string** | The section of the question | [optional] 

## Methods

### NewQuestion

`func NewQuestion(answer string, question string, ) *Question`

NewQuestion instantiates a new Question object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionWithDefaults

`func NewQuestionWithDefaults() *Question`

NewQuestionWithDefaults instantiates a new Question object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *Question) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *Question) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *Question) SetAnswer(v string)`

SetAnswer sets Answer field to given value.


### GetQuestion

`func (o *Question) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *Question) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *Question) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetSection

`func (o *Question) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *Question) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *Question) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *Question) HasSection() bool`

HasSection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


