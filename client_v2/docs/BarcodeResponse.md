# BarcodeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | ID of the account for which the barcode was generated. | 
**BarcodeNumber** | **string** | The generated barcode number. | 
**CreationTime** | **time.Time** | Timestamp of when the barcode was created. | 
**Currency** | **string** | ISO 4217 Alpha-3 currency code for the cash transaction. | 
**CustomerId** | **string** | ID of the customer for whom the barcode was generated. | 
**CustomerLatitude** | **float32** | Latitude of the customer location. | 
**CustomerLongitude** | **float32** | Longitude of the customer location. | 
**ExternalDeviceId** | Pointer to **string** | ID of the external device used for the barcode generation. | [optional] 
**Id** | **string** | ID of the generated barcode. | 
**LastUpdatedTime** | **time.Time** | Timestamp of the last update to the barcode. | 
**MaxAmount** | **float32** | Maximum amount for the cash transaction. | 
**Metadata** | Pointer to **map[string]string** | Any additional custom metadata related to the barcode.  * Can contain up to 10 key-value pairs with up to 200 characters each.  | [optional] 
**MinAmount** | **float32** | Minimum amount for the cash transaction. | 
**Status** | [**BarcodeStatus**](BarcodeStatus.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**TimestampValidTo** | **time.Time** | Timestamp indicating the expiration time of the barcode. Always set to 15 minutes after the creation time. | 
**Type** | [**BarcodeType**](BarcodeType.md) |  | 

## Methods

### NewBarcodeResponse

`func NewBarcodeResponse(accountId string, barcodeNumber string, creationTime time.Time, currency string, customerId string, customerLatitude float32, customerLongitude float32, id string, lastUpdatedTime time.Time, maxAmount float32, minAmount float32, status BarcodeStatus, tenant string, timestampValidTo time.Time, type_ BarcodeType, ) *BarcodeResponse`

NewBarcodeResponse instantiates a new BarcodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBarcodeResponseWithDefaults

`func NewBarcodeResponseWithDefaults() *BarcodeResponse`

NewBarcodeResponseWithDefaults instantiates a new BarcodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BarcodeResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BarcodeResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BarcodeResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBarcodeNumber

`func (o *BarcodeResponse) GetBarcodeNumber() string`

GetBarcodeNumber returns the BarcodeNumber field if non-nil, zero value otherwise.

### GetBarcodeNumberOk

`func (o *BarcodeResponse) GetBarcodeNumberOk() (*string, bool)`

GetBarcodeNumberOk returns a tuple with the BarcodeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeNumber

`func (o *BarcodeResponse) SetBarcodeNumber(v string)`

SetBarcodeNumber sets BarcodeNumber field to given value.


### GetCreationTime

`func (o *BarcodeResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BarcodeResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BarcodeResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *BarcodeResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BarcodeResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BarcodeResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *BarcodeResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *BarcodeResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *BarcodeResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCustomerLatitude

`func (o *BarcodeResponse) GetCustomerLatitude() float32`

GetCustomerLatitude returns the CustomerLatitude field if non-nil, zero value otherwise.

### GetCustomerLatitudeOk

`func (o *BarcodeResponse) GetCustomerLatitudeOk() (*float32, bool)`

GetCustomerLatitudeOk returns a tuple with the CustomerLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLatitude

`func (o *BarcodeResponse) SetCustomerLatitude(v float32)`

SetCustomerLatitude sets CustomerLatitude field to given value.


### GetCustomerLongitude

`func (o *BarcodeResponse) GetCustomerLongitude() float32`

GetCustomerLongitude returns the CustomerLongitude field if non-nil, zero value otherwise.

### GetCustomerLongitudeOk

`func (o *BarcodeResponse) GetCustomerLongitudeOk() (*float32, bool)`

GetCustomerLongitudeOk returns a tuple with the CustomerLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLongitude

`func (o *BarcodeResponse) SetCustomerLongitude(v float32)`

SetCustomerLongitude sets CustomerLongitude field to given value.


### GetExternalDeviceId

`func (o *BarcodeResponse) GetExternalDeviceId() string`

GetExternalDeviceId returns the ExternalDeviceId field if non-nil, zero value otherwise.

### GetExternalDeviceIdOk

`func (o *BarcodeResponse) GetExternalDeviceIdOk() (*string, bool)`

GetExternalDeviceIdOk returns a tuple with the ExternalDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDeviceId

`func (o *BarcodeResponse) SetExternalDeviceId(v string)`

SetExternalDeviceId sets ExternalDeviceId field to given value.

### HasExternalDeviceId

`func (o *BarcodeResponse) HasExternalDeviceId() bool`

HasExternalDeviceId returns a boolean if a field has been set.

### GetId

`func (o *BarcodeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BarcodeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BarcodeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *BarcodeResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BarcodeResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BarcodeResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMaxAmount

`func (o *BarcodeResponse) GetMaxAmount() float32`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *BarcodeResponse) GetMaxAmountOk() (*float32, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *BarcodeResponse) SetMaxAmount(v float32)`

SetMaxAmount sets MaxAmount field to given value.


### GetMetadata

`func (o *BarcodeResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BarcodeResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BarcodeResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BarcodeResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMinAmount

`func (o *BarcodeResponse) GetMinAmount() float32`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *BarcodeResponse) GetMinAmountOk() (*float32, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *BarcodeResponse) SetMinAmount(v float32)`

SetMinAmount sets MinAmount field to given value.


### GetStatus

`func (o *BarcodeResponse) GetStatus() BarcodeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BarcodeResponse) GetStatusOk() (*BarcodeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BarcodeResponse) SetStatus(v BarcodeStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *BarcodeResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BarcodeResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BarcodeResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTimestampValidTo

`func (o *BarcodeResponse) GetTimestampValidTo() time.Time`

GetTimestampValidTo returns the TimestampValidTo field if non-nil, zero value otherwise.

### GetTimestampValidToOk

`func (o *BarcodeResponse) GetTimestampValidToOk() (*time.Time, bool)`

GetTimestampValidToOk returns a tuple with the TimestampValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampValidTo

`func (o *BarcodeResponse) SetTimestampValidTo(v time.Time)`

SetTimestampValidTo sets TimestampValidTo field to given value.


### GetType

`func (o *BarcodeResponse) GetType() BarcodeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BarcodeResponse) GetTypeOk() (*BarcodeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BarcodeResponse) SetType(v BarcodeType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


