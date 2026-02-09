# BarcodeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 
**Barcodes** | [**[]BarcodeResponse**](BarcodeResponse.md) |  | 

## Methods

### NewBarcodeListResponse

`func NewBarcodeListResponse(barcodes []BarcodeResponse, ) *BarcodeListResponse`

NewBarcodeListResponse instantiates a new BarcodeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBarcodeListResponseWithDefaults

`func NewBarcodeListResponseWithDefaults() *BarcodeListResponse`

NewBarcodeListResponseWithDefaults instantiates a new BarcodeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *BarcodeListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *BarcodeListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *BarcodeListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *BarcodeListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetBarcodes

`func (o *BarcodeListResponse) GetBarcodes() []BarcodeResponse`

GetBarcodes returns the Barcodes field if non-nil, zero value otherwise.

### GetBarcodesOk

`func (o *BarcodeListResponse) GetBarcodesOk() (*[]BarcodeResponse, bool)`

GetBarcodesOk returns a tuple with the Barcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodes

`func (o *BarcodeListResponse) SetBarcodes(v []BarcodeResponse)`

SetBarcodes sets Barcodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


