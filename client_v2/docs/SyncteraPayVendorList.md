# SyncteraPayVendorList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Vendors** | [**[]SyncteraPayVendorResponse**](SyncteraPayVendorResponse.md) | Array of vendors | 

## Methods

### NewSyncteraPayVendorList

`func NewSyncteraPayVendorList(vendors []SyncteraPayVendorResponse, ) *SyncteraPayVendorList`

NewSyncteraPayVendorList instantiates a new SyncteraPayVendorList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayVendorListWithDefaults

`func NewSyncteraPayVendorListWithDefaults() *SyncteraPayVendorList`

NewSyncteraPayVendorListWithDefaults instantiates a new SyncteraPayVendorList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *SyncteraPayVendorList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *SyncteraPayVendorList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *SyncteraPayVendorList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *SyncteraPayVendorList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetVendors

`func (o *SyncteraPayVendorList) GetVendors() []SyncteraPayVendorResponse`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *SyncteraPayVendorList) GetVendorsOk() (*[]SyncteraPayVendorResponse, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *SyncteraPayVendorList) SetVendors(v []SyncteraPayVendorResponse)`

SetVendors sets Vendors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


