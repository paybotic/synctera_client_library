# RestrictedApplicationDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | File description | 
**DocumentId** | **string** | ID of the document within the Synctera platform | 
**FintechItemId** | Pointer to **string** | Client supplied item id | [optional] 
**Type** | [**RestrictedApplicationItemType**](RestrictedApplicationItemType.md) |  | 

## Methods

### NewRestrictedApplicationDocument

`func NewRestrictedApplicationDocument(description string, documentId string, type_ RestrictedApplicationItemType, ) *RestrictedApplicationDocument`

NewRestrictedApplicationDocument instantiates a new RestrictedApplicationDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedApplicationDocumentWithDefaults

`func NewRestrictedApplicationDocumentWithDefaults() *RestrictedApplicationDocument`

NewRestrictedApplicationDocumentWithDefaults instantiates a new RestrictedApplicationDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RestrictedApplicationDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestrictedApplicationDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestrictedApplicationDocument) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentId

`func (o *RestrictedApplicationDocument) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *RestrictedApplicationDocument) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *RestrictedApplicationDocument) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetFintechItemId

`func (o *RestrictedApplicationDocument) GetFintechItemId() string`

GetFintechItemId returns the FintechItemId field if non-nil, zero value otherwise.

### GetFintechItemIdOk

`func (o *RestrictedApplicationDocument) GetFintechItemIdOk() (*string, bool)`

GetFintechItemIdOk returns a tuple with the FintechItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFintechItemId

`func (o *RestrictedApplicationDocument) SetFintechItemId(v string)`

SetFintechItemId sets FintechItemId field to given value.

### HasFintechItemId

`func (o *RestrictedApplicationDocument) HasFintechItemId() bool`

HasFintechItemId returns a boolean if a field has been set.

### GetType

`func (o *RestrictedApplicationDocument) GetType() RestrictedApplicationItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestrictedApplicationDocument) GetTypeOk() (*RestrictedApplicationItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestrictedApplicationDocument) SetType(v RestrictedApplicationItemType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


