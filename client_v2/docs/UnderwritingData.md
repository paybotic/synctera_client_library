# UnderwritingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestTime** | Pointer to **time.Time** | Timestamp of the request in RFC3359 format | [optional] 
**Vendor** | Pointer to **string** | Vendor name | [optional] 
**VendorInfo** | Pointer to **map[string]interface{}** | Information about the vendor provided info | [optional] 

## Methods

### NewUnderwritingData

`func NewUnderwritingData() *UnderwritingData`

NewUnderwritingData instantiates a new UnderwritingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderwritingDataWithDefaults

`func NewUnderwritingDataWithDefaults() *UnderwritingData`

NewUnderwritingDataWithDefaults instantiates a new UnderwritingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestTime

`func (o *UnderwritingData) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *UnderwritingData) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *UnderwritingData) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *UnderwritingData) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetVendor

`func (o *UnderwritingData) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *UnderwritingData) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *UnderwritingData) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *UnderwritingData) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorInfo

`func (o *UnderwritingData) GetVendorInfo() map[string]interface{}`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *UnderwritingData) GetVendorInfoOk() (*map[string]interface{}, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *UnderwritingData) SetVendorInfo(v map[string]interface{})`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *UnderwritingData) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


