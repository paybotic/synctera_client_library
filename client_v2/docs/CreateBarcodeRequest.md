# CreateBarcodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | ID of the account for which the barcode is generated. | 
**CustomerId** | **string** | ID of the customer for whom the barcode is generated. | 
**CustomerLatitude** | **float32** | Latitude of the customer location. | 
**CustomerLongitude** | **float32** | Longitude of the customer location. | 
**ExternalDeviceId** | Pointer to **string** | ID of the external device used for the barcode generation. | [optional] 
**Metadata** | Pointer to **map[string]string** | Any additional custom metadata related to the barcode.  * Can contain up to 10 key-value pairs with up to 200 characters each.  | [optional] 
**Type** | [**BarcodeType**](BarcodeType.md) |  | 

## Methods

### NewCreateBarcodeRequest

`func NewCreateBarcodeRequest(accountId string, customerId string, customerLatitude float32, customerLongitude float32, type_ BarcodeType, ) *CreateBarcodeRequest`

NewCreateBarcodeRequest instantiates a new CreateBarcodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBarcodeRequestWithDefaults

`func NewCreateBarcodeRequestWithDefaults() *CreateBarcodeRequest`

NewCreateBarcodeRequestWithDefaults instantiates a new CreateBarcodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateBarcodeRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateBarcodeRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateBarcodeRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomerId

`func (o *CreateBarcodeRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateBarcodeRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateBarcodeRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCustomerLatitude

`func (o *CreateBarcodeRequest) GetCustomerLatitude() float32`

GetCustomerLatitude returns the CustomerLatitude field if non-nil, zero value otherwise.

### GetCustomerLatitudeOk

`func (o *CreateBarcodeRequest) GetCustomerLatitudeOk() (*float32, bool)`

GetCustomerLatitudeOk returns a tuple with the CustomerLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLatitude

`func (o *CreateBarcodeRequest) SetCustomerLatitude(v float32)`

SetCustomerLatitude sets CustomerLatitude field to given value.


### GetCustomerLongitude

`func (o *CreateBarcodeRequest) GetCustomerLongitude() float32`

GetCustomerLongitude returns the CustomerLongitude field if non-nil, zero value otherwise.

### GetCustomerLongitudeOk

`func (o *CreateBarcodeRequest) GetCustomerLongitudeOk() (*float32, bool)`

GetCustomerLongitudeOk returns a tuple with the CustomerLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLongitude

`func (o *CreateBarcodeRequest) SetCustomerLongitude(v float32)`

SetCustomerLongitude sets CustomerLongitude field to given value.


### GetExternalDeviceId

`func (o *CreateBarcodeRequest) GetExternalDeviceId() string`

GetExternalDeviceId returns the ExternalDeviceId field if non-nil, zero value otherwise.

### GetExternalDeviceIdOk

`func (o *CreateBarcodeRequest) GetExternalDeviceIdOk() (*string, bool)`

GetExternalDeviceIdOk returns a tuple with the ExternalDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDeviceId

`func (o *CreateBarcodeRequest) SetExternalDeviceId(v string)`

SetExternalDeviceId sets ExternalDeviceId field to given value.

### HasExternalDeviceId

`func (o *CreateBarcodeRequest) HasExternalDeviceId() bool`

HasExternalDeviceId returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateBarcodeRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateBarcodeRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateBarcodeRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateBarcodeRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *CreateBarcodeRequest) GetType() BarcodeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBarcodeRequest) GetTypeOk() (*BarcodeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBarcodeRequest) SetType(v BarcodeType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


