# BarcodeSimulationStoreAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | **string** | First line of the store address | 
**Address2** | Pointer to **string** | Second line of the store address (optional) | [optional] 
**City** | **string** | City where the store is located | 
**Country** | **string** | Country where the store is located | 
**County** | Pointer to **string** | County where the store is located | [optional] 
**State** | **string** | State where the store is located | 
**ZipCode** | **string** | ZIP code of the store location | 

## Methods

### NewBarcodeSimulationStoreAddress

`func NewBarcodeSimulationStoreAddress(address1 string, city string, country string, state string, zipCode string, ) *BarcodeSimulationStoreAddress`

NewBarcodeSimulationStoreAddress instantiates a new BarcodeSimulationStoreAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBarcodeSimulationStoreAddressWithDefaults

`func NewBarcodeSimulationStoreAddressWithDefaults() *BarcodeSimulationStoreAddress`

NewBarcodeSimulationStoreAddressWithDefaults instantiates a new BarcodeSimulationStoreAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *BarcodeSimulationStoreAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *BarcodeSimulationStoreAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *BarcodeSimulationStoreAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *BarcodeSimulationStoreAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *BarcodeSimulationStoreAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *BarcodeSimulationStoreAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *BarcodeSimulationStoreAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *BarcodeSimulationStoreAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BarcodeSimulationStoreAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BarcodeSimulationStoreAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *BarcodeSimulationStoreAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BarcodeSimulationStoreAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BarcodeSimulationStoreAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCounty

`func (o *BarcodeSimulationStoreAddress) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *BarcodeSimulationStoreAddress) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *BarcodeSimulationStoreAddress) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *BarcodeSimulationStoreAddress) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetState

`func (o *BarcodeSimulationStoreAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BarcodeSimulationStoreAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BarcodeSimulationStoreAddress) SetState(v string)`

SetState sets State field to given value.


### GetZipCode

`func (o *BarcodeSimulationStoreAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *BarcodeSimulationStoreAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *BarcodeSimulationStoreAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


