# LicenseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Licenses** | Pointer to [**[]ResponseLicense**](ResponseLicense.md) | Array of license records | [optional] 

## Methods

### NewLicenseList

`func NewLicenseList() *LicenseList`

NewLicenseList instantiates a new LicenseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseListWithDefaults

`func NewLicenseListWithDefaults() *LicenseList`

NewLicenseListWithDefaults instantiates a new LicenseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *LicenseList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *LicenseList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *LicenseList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *LicenseList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetLicenses

`func (o *LicenseList) GetLicenses() []ResponseLicense`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *LicenseList) GetLicensesOk() (*[]ResponseLicense, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *LicenseList) SetLicenses(v []ResponseLicense)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *LicenseList) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


