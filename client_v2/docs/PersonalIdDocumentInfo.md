# PersonalIdDocumentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryDate** | Pointer to **string** | The date the associated document is set to expire on set by the governing authority. | [optional] 
**IssueDate** | Pointer to **string** | The date the associated document was issued by the governing authority. | [optional] 

## Methods

### NewPersonalIdDocumentInfo

`func NewPersonalIdDocumentInfo() *PersonalIdDocumentInfo`

NewPersonalIdDocumentInfo instantiates a new PersonalIdDocumentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalIdDocumentInfoWithDefaults

`func NewPersonalIdDocumentInfoWithDefaults() *PersonalIdDocumentInfo`

NewPersonalIdDocumentInfoWithDefaults instantiates a new PersonalIdDocumentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryDate

`func (o *PersonalIdDocumentInfo) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *PersonalIdDocumentInfo) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *PersonalIdDocumentInfo) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *PersonalIdDocumentInfo) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetIssueDate

`func (o *PersonalIdDocumentInfo) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *PersonalIdDocumentInfo) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *PersonalIdDocumentInfo) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *PersonalIdDocumentInfo) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


