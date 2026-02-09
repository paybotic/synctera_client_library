# LoanproVendorDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CashAdvanceBucketId** | **int64** | LoanPro cash advance bucket ID | 
**LocId** | **int64** | LoanPro line of credit account ID | 
**LocProductId** | **int64** | LoanPro line of credit product ID | 
**PurchasesBucketId** | **int64** | LoanPro purchases bucket ID | 
**TenantId** | **string** | LoanPro tenant ID | 
**DailyArchive** | Pointer to [**LoanproDailyArchive**](LoanproDailyArchive.md) |  | [optional] 

## Methods

### NewLoanproVendorDataResponse

`func NewLoanproVendorDataResponse(cashAdvanceBucketId int64, locId int64, locProductId int64, purchasesBucketId int64, tenantId string, ) *LoanproVendorDataResponse`

NewLoanproVendorDataResponse instantiates a new LoanproVendorDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoanproVendorDataResponseWithDefaults

`func NewLoanproVendorDataResponseWithDefaults() *LoanproVendorDataResponse`

NewLoanproVendorDataResponseWithDefaults instantiates a new LoanproVendorDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCashAdvanceBucketId

`func (o *LoanproVendorDataResponse) GetCashAdvanceBucketId() int64`

GetCashAdvanceBucketId returns the CashAdvanceBucketId field if non-nil, zero value otherwise.

### GetCashAdvanceBucketIdOk

`func (o *LoanproVendorDataResponse) GetCashAdvanceBucketIdOk() (*int64, bool)`

GetCashAdvanceBucketIdOk returns a tuple with the CashAdvanceBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAdvanceBucketId

`func (o *LoanproVendorDataResponse) SetCashAdvanceBucketId(v int64)`

SetCashAdvanceBucketId sets CashAdvanceBucketId field to given value.


### GetLocId

`func (o *LoanproVendorDataResponse) GetLocId() int64`

GetLocId returns the LocId field if non-nil, zero value otherwise.

### GetLocIdOk

`func (o *LoanproVendorDataResponse) GetLocIdOk() (*int64, bool)`

GetLocIdOk returns a tuple with the LocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocId

`func (o *LoanproVendorDataResponse) SetLocId(v int64)`

SetLocId sets LocId field to given value.


### GetLocProductId

`func (o *LoanproVendorDataResponse) GetLocProductId() int64`

GetLocProductId returns the LocProductId field if non-nil, zero value otherwise.

### GetLocProductIdOk

`func (o *LoanproVendorDataResponse) GetLocProductIdOk() (*int64, bool)`

GetLocProductIdOk returns a tuple with the LocProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocProductId

`func (o *LoanproVendorDataResponse) SetLocProductId(v int64)`

SetLocProductId sets LocProductId field to given value.


### GetPurchasesBucketId

`func (o *LoanproVendorDataResponse) GetPurchasesBucketId() int64`

GetPurchasesBucketId returns the PurchasesBucketId field if non-nil, zero value otherwise.

### GetPurchasesBucketIdOk

`func (o *LoanproVendorDataResponse) GetPurchasesBucketIdOk() (*int64, bool)`

GetPurchasesBucketIdOk returns a tuple with the PurchasesBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasesBucketId

`func (o *LoanproVendorDataResponse) SetPurchasesBucketId(v int64)`

SetPurchasesBucketId sets PurchasesBucketId field to given value.


### GetTenantId

`func (o *LoanproVendorDataResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LoanproVendorDataResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LoanproVendorDataResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetDailyArchive

`func (o *LoanproVendorDataResponse) GetDailyArchive() LoanproDailyArchive`

GetDailyArchive returns the DailyArchive field if non-nil, zero value otherwise.

### GetDailyArchiveOk

`func (o *LoanproVendorDataResponse) GetDailyArchiveOk() (*LoanproDailyArchive, bool)`

GetDailyArchiveOk returns a tuple with the DailyArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyArchive

`func (o *LoanproVendorDataResponse) SetDailyArchive(v LoanproDailyArchive)`

SetDailyArchive sets DailyArchive field to given value.

### HasDailyArchive

`func (o *LoanproVendorDataResponse) HasDailyArchive() bool`

HasDailyArchive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


