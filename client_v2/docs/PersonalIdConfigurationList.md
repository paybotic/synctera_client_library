# PersonalIdConfigurationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**PersonalIdConfigurations** | [**[]PersonalIdConfigurationResponse**](PersonalIdConfigurationResponse.md) | Array of personal ID configurations | 

## Methods

### NewPersonalIdConfigurationList

`func NewPersonalIdConfigurationList(personalIdConfigurations []PersonalIdConfigurationResponse, ) *PersonalIdConfigurationList`

NewPersonalIdConfigurationList instantiates a new PersonalIdConfigurationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalIdConfigurationListWithDefaults

`func NewPersonalIdConfigurationListWithDefaults() *PersonalIdConfigurationList`

NewPersonalIdConfigurationListWithDefaults instantiates a new PersonalIdConfigurationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *PersonalIdConfigurationList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *PersonalIdConfigurationList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *PersonalIdConfigurationList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *PersonalIdConfigurationList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPersonalIdConfigurations

`func (o *PersonalIdConfigurationList) GetPersonalIdConfigurations() []PersonalIdConfigurationResponse`

GetPersonalIdConfigurations returns the PersonalIdConfigurations field if non-nil, zero value otherwise.

### GetPersonalIdConfigurationsOk

`func (o *PersonalIdConfigurationList) GetPersonalIdConfigurationsOk() (*[]PersonalIdConfigurationResponse, bool)`

GetPersonalIdConfigurationsOk returns a tuple with the PersonalIdConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalIdConfigurations

`func (o *PersonalIdConfigurationList) SetPersonalIdConfigurations(v []PersonalIdConfigurationResponse)`

SetPersonalIdConfigurations sets PersonalIdConfigurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


