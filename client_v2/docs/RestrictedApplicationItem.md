# RestrictedApplicationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | **string** | Responses to the question | 
**FintechItemId** | Pointer to **string** | Client supplied item id | [optional] 
**Question** | **string** | Question for the applicant | 
**Type** | [**RestrictedApplicationItemType**](RestrictedApplicationItemType.md) |  | 
**Description** | **string** | File description | 
**DocumentId** | **string** | ID of the document within the Synctera platform | 

## Methods

### NewRestrictedApplicationItem

`func NewRestrictedApplicationItem(answer string, question string, type_ RestrictedApplicationItemType, description string, documentId string, ) *RestrictedApplicationItem`

NewRestrictedApplicationItem instantiates a new RestrictedApplicationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedApplicationItemWithDefaults

`func NewRestrictedApplicationItemWithDefaults() *RestrictedApplicationItem`

NewRestrictedApplicationItemWithDefaults instantiates a new RestrictedApplicationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *RestrictedApplicationItem) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *RestrictedApplicationItem) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *RestrictedApplicationItem) SetAnswer(v string)`

SetAnswer sets Answer field to given value.


### GetFintechItemId

`func (o *RestrictedApplicationItem) GetFintechItemId() string`

GetFintechItemId returns the FintechItemId field if non-nil, zero value otherwise.

### GetFintechItemIdOk

`func (o *RestrictedApplicationItem) GetFintechItemIdOk() (*string, bool)`

GetFintechItemIdOk returns a tuple with the FintechItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFintechItemId

`func (o *RestrictedApplicationItem) SetFintechItemId(v string)`

SetFintechItemId sets FintechItemId field to given value.

### HasFintechItemId

`func (o *RestrictedApplicationItem) HasFintechItemId() bool`

HasFintechItemId returns a boolean if a field has been set.

### GetQuestion

`func (o *RestrictedApplicationItem) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *RestrictedApplicationItem) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *RestrictedApplicationItem) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetType

`func (o *RestrictedApplicationItem) GetType() RestrictedApplicationItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestrictedApplicationItem) GetTypeOk() (*RestrictedApplicationItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestrictedApplicationItem) SetType(v RestrictedApplicationItemType)`

SetType sets Type field to given value.


### GetDescription

`func (o *RestrictedApplicationItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestrictedApplicationItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestrictedApplicationItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentId

`func (o *RestrictedApplicationItem) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *RestrictedApplicationItem) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *RestrictedApplicationItem) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


