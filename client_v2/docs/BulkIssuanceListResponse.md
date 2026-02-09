# BulkIssuanceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**BulkOrderConfigs** | [**[]BulkIssuanceResponse**](BulkIssuanceResponse.md) | Array of bulk order configurations | 

## Methods

### NewBulkIssuanceListResponse

`func NewBulkIssuanceListResponse(bulkOrderConfigs []BulkIssuanceResponse, ) *BulkIssuanceListResponse`

NewBulkIssuanceListResponse instantiates a new BulkIssuanceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIssuanceListResponseWithDefaults

`func NewBulkIssuanceListResponseWithDefaults() *BulkIssuanceListResponse`

NewBulkIssuanceListResponseWithDefaults instantiates a new BulkIssuanceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *BulkIssuanceListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *BulkIssuanceListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *BulkIssuanceListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *BulkIssuanceListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetBulkOrderConfigs

`func (o *BulkIssuanceListResponse) GetBulkOrderConfigs() []BulkIssuanceResponse`

GetBulkOrderConfigs returns the BulkOrderConfigs field if non-nil, zero value otherwise.

### GetBulkOrderConfigsOk

`func (o *BulkIssuanceListResponse) GetBulkOrderConfigsOk() (*[]BulkIssuanceResponse, bool)`

GetBulkOrderConfigsOk returns a tuple with the BulkOrderConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkOrderConfigs

`func (o *BulkIssuanceListResponse) SetBulkOrderConfigs(v []BulkIssuanceResponse)`

SetBulkOrderConfigs sets BulkOrderConfigs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


