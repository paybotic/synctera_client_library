# DocumentVersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created | [optional] [readonly] 
**FileName** | Pointer to **string** | The file name of the document | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated | [optional] [readonly] 
**Version** | Pointer to **int32** | Positive integer representing the version of the document | [optional] 

## Methods

### NewDocumentVersionInfo

`func NewDocumentVersionInfo() *DocumentVersionInfo`

NewDocumentVersionInfo instantiates a new DocumentVersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentVersionInfoWithDefaults

`func NewDocumentVersionInfoWithDefaults() *DocumentVersionInfo`

NewDocumentVersionInfoWithDefaults instantiates a new DocumentVersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *DocumentVersionInfo) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DocumentVersionInfo) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DocumentVersionInfo) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *DocumentVersionInfo) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetFileName

`func (o *DocumentVersionInfo) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DocumentVersionInfo) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DocumentVersionInfo) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DocumentVersionInfo) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *DocumentVersionInfo) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *DocumentVersionInfo) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *DocumentVersionInfo) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *DocumentVersionInfo) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetVersion

`func (o *DocumentVersionInfo) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DocumentVersionInfo) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DocumentVersionInfo) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DocumentVersionInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


