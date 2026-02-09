# Initialize3dsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceDataCollectionJwt** | **string** | JWT used for device data collection | 
**DeviceDataCollectionUrl** | **string** | URL used for device data collection | 
**Id** | **string** | The unique identifier of the 3DS authentication | 

## Methods

### NewInitialize3dsResponse

`func NewInitialize3dsResponse(deviceDataCollectionJwt string, deviceDataCollectionUrl string, id string, ) *Initialize3dsResponse`

NewInitialize3dsResponse instantiates a new Initialize3dsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialize3dsResponseWithDefaults

`func NewInitialize3dsResponseWithDefaults() *Initialize3dsResponse`

NewInitialize3dsResponseWithDefaults instantiates a new Initialize3dsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceDataCollectionJwt

`func (o *Initialize3dsResponse) GetDeviceDataCollectionJwt() string`

GetDeviceDataCollectionJwt returns the DeviceDataCollectionJwt field if non-nil, zero value otherwise.

### GetDeviceDataCollectionJwtOk

`func (o *Initialize3dsResponse) GetDeviceDataCollectionJwtOk() (*string, bool)`

GetDeviceDataCollectionJwtOk returns a tuple with the DeviceDataCollectionJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDataCollectionJwt

`func (o *Initialize3dsResponse) SetDeviceDataCollectionJwt(v string)`

SetDeviceDataCollectionJwt sets DeviceDataCollectionJwt field to given value.


### GetDeviceDataCollectionUrl

`func (o *Initialize3dsResponse) GetDeviceDataCollectionUrl() string`

GetDeviceDataCollectionUrl returns the DeviceDataCollectionUrl field if non-nil, zero value otherwise.

### GetDeviceDataCollectionUrlOk

`func (o *Initialize3dsResponse) GetDeviceDataCollectionUrlOk() (*string, bool)`

GetDeviceDataCollectionUrlOk returns a tuple with the DeviceDataCollectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDataCollectionUrl

`func (o *Initialize3dsResponse) SetDeviceDataCollectionUrl(v string)`

SetDeviceDataCollectionUrl sets DeviceDataCollectionUrl field to given value.


### GetId

`func (o *Initialize3dsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Initialize3dsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Initialize3dsResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


