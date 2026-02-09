# RestrictedApplicationQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | **string** | Responses to the question | 
**FintechItemId** | Pointer to **string** | Client supplied item id | [optional] 
**Question** | **string** | Question for the applicant | 
**Type** | [**RestrictedApplicationItemType**](RestrictedApplicationItemType.md) |  | 

## Methods

### NewRestrictedApplicationQuestion

`func NewRestrictedApplicationQuestion(answer string, question string, type_ RestrictedApplicationItemType, ) *RestrictedApplicationQuestion`

NewRestrictedApplicationQuestion instantiates a new RestrictedApplicationQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedApplicationQuestionWithDefaults

`func NewRestrictedApplicationQuestionWithDefaults() *RestrictedApplicationQuestion`

NewRestrictedApplicationQuestionWithDefaults instantiates a new RestrictedApplicationQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *RestrictedApplicationQuestion) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *RestrictedApplicationQuestion) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *RestrictedApplicationQuestion) SetAnswer(v string)`

SetAnswer sets Answer field to given value.


### GetFintechItemId

`func (o *RestrictedApplicationQuestion) GetFintechItemId() string`

GetFintechItemId returns the FintechItemId field if non-nil, zero value otherwise.

### GetFintechItemIdOk

`func (o *RestrictedApplicationQuestion) GetFintechItemIdOk() (*string, bool)`

GetFintechItemIdOk returns a tuple with the FintechItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFintechItemId

`func (o *RestrictedApplicationQuestion) SetFintechItemId(v string)`

SetFintechItemId sets FintechItemId field to given value.

### HasFintechItemId

`func (o *RestrictedApplicationQuestion) HasFintechItemId() bool`

HasFintechItemId returns a boolean if a field has been set.

### GetQuestion

`func (o *RestrictedApplicationQuestion) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *RestrictedApplicationQuestion) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *RestrictedApplicationQuestion) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetType

`func (o *RestrictedApplicationQuestion) GetType() RestrictedApplicationItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestrictedApplicationQuestion) GetTypeOk() (*RestrictedApplicationItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestrictedApplicationQuestion) SetType(v RestrictedApplicationItemType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


