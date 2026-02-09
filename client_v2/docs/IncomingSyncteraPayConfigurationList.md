# IncomingSyncteraPayConfigurationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Configurations** | [**[]IncomingSyncteraPayConfigurationResponse**](IncomingSyncteraPayConfigurationResponse.md) |  | 

## Methods

### NewIncomingSyncteraPayConfigurationList

`func NewIncomingSyncteraPayConfigurationList(configurations []IncomingSyncteraPayConfigurationResponse, ) *IncomingSyncteraPayConfigurationList`

NewIncomingSyncteraPayConfigurationList instantiates a new IncomingSyncteraPayConfigurationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingSyncteraPayConfigurationListWithDefaults

`func NewIncomingSyncteraPayConfigurationListWithDefaults() *IncomingSyncteraPayConfigurationList`

NewIncomingSyncteraPayConfigurationListWithDefaults instantiates a new IncomingSyncteraPayConfigurationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *IncomingSyncteraPayConfigurationList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *IncomingSyncteraPayConfigurationList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *IncomingSyncteraPayConfigurationList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *IncomingSyncteraPayConfigurationList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetConfigurations

`func (o *IncomingSyncteraPayConfigurationList) GetConfigurations() []IncomingSyncteraPayConfigurationResponse`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *IncomingSyncteraPayConfigurationList) GetConfigurationsOk() (*[]IncomingSyncteraPayConfigurationResponse, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *IncomingSyncteraPayConfigurationList) SetConfigurations(v []IncomingSyncteraPayConfigurationResponse)`

SetConfigurations sets Configurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


