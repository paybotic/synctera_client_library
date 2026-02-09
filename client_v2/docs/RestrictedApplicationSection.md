# RestrictedApplicationSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FintechSectionId** | Pointer to **string** | Client supplied section ID | [optional] 
**Pages** | [**[]RestrictedApplicationPage**](RestrictedApplicationPage.md) | Pages within the section | 
**Title** | **string** | Section title | 

## Methods

### NewRestrictedApplicationSection

`func NewRestrictedApplicationSection(pages []RestrictedApplicationPage, title string, ) *RestrictedApplicationSection`

NewRestrictedApplicationSection instantiates a new RestrictedApplicationSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedApplicationSectionWithDefaults

`func NewRestrictedApplicationSectionWithDefaults() *RestrictedApplicationSection`

NewRestrictedApplicationSectionWithDefaults instantiates a new RestrictedApplicationSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFintechSectionId

`func (o *RestrictedApplicationSection) GetFintechSectionId() string`

GetFintechSectionId returns the FintechSectionId field if non-nil, zero value otherwise.

### GetFintechSectionIdOk

`func (o *RestrictedApplicationSection) GetFintechSectionIdOk() (*string, bool)`

GetFintechSectionIdOk returns a tuple with the FintechSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFintechSectionId

`func (o *RestrictedApplicationSection) SetFintechSectionId(v string)`

SetFintechSectionId sets FintechSectionId field to given value.

### HasFintechSectionId

`func (o *RestrictedApplicationSection) HasFintechSectionId() bool`

HasFintechSectionId returns a boolean if a field has been set.

### GetPages

`func (o *RestrictedApplicationSection) GetPages() []RestrictedApplicationPage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *RestrictedApplicationSection) GetPagesOk() (*[]RestrictedApplicationPage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *RestrictedApplicationSection) SetPages(v []RestrictedApplicationPage)`

SetPages sets Pages field to given value.


### GetTitle

`func (o *RestrictedApplicationSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RestrictedApplicationSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RestrictedApplicationSection) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


