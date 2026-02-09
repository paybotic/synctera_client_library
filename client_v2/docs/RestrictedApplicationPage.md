# RestrictedApplicationPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FintechPageId** | Pointer to **string** | Client supplied page id | [optional] 
**Items** | [**[]RestrictedApplicationItem**](RestrictedApplicationItem.md) | Items (questions and files) within the page | 
**Title** | **string** | Page title | 

## Methods

### NewRestrictedApplicationPage

`func NewRestrictedApplicationPage(items []RestrictedApplicationItem, title string, ) *RestrictedApplicationPage`

NewRestrictedApplicationPage instantiates a new RestrictedApplicationPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedApplicationPageWithDefaults

`func NewRestrictedApplicationPageWithDefaults() *RestrictedApplicationPage`

NewRestrictedApplicationPageWithDefaults instantiates a new RestrictedApplicationPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFintechPageId

`func (o *RestrictedApplicationPage) GetFintechPageId() string`

GetFintechPageId returns the FintechPageId field if non-nil, zero value otherwise.

### GetFintechPageIdOk

`func (o *RestrictedApplicationPage) GetFintechPageIdOk() (*string, bool)`

GetFintechPageIdOk returns a tuple with the FintechPageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFintechPageId

`func (o *RestrictedApplicationPage) SetFintechPageId(v string)`

SetFintechPageId sets FintechPageId field to given value.

### HasFintechPageId

`func (o *RestrictedApplicationPage) HasFintechPageId() bool`

HasFintechPageId returns a boolean if a field has been set.

### GetItems

`func (o *RestrictedApplicationPage) GetItems() []RestrictedApplicationItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RestrictedApplicationPage) GetItemsOk() (*[]RestrictedApplicationItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RestrictedApplicationPage) SetItems(v []RestrictedApplicationItem)`

SetItems sets Items field to given value.


### GetTitle

`func (o *RestrictedApplicationPage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RestrictedApplicationPage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RestrictedApplicationPage) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


