# BulkIssuanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkIssuancePolicy** | [**BulkIssuancePolicy**](BulkIssuancePolicy.md) |  | [default to BULKISSUANCEPOLICY_AUTO]
**CardProductId** | **string** | The unique identifier of a cards product | 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the bulk order config was created | [optional] [readonly] 
**Id** | **string** | The unique identifier of a bulk order configuration | 
**Name** | **string** | Name associated with the bulk order configuration. | 
**Shipping** | [**BulkShipping**](BulkShipping.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 

## Methods

### NewBulkIssuanceResponse

`func NewBulkIssuanceResponse(bulkIssuancePolicy BulkIssuancePolicy, cardProductId string, id string, name string, shipping BulkShipping, tenant string, ) *BulkIssuanceResponse`

NewBulkIssuanceResponse instantiates a new BulkIssuanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIssuanceResponseWithDefaults

`func NewBulkIssuanceResponseWithDefaults() *BulkIssuanceResponse`

NewBulkIssuanceResponseWithDefaults instantiates a new BulkIssuanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkIssuancePolicy

`func (o *BulkIssuanceResponse) GetBulkIssuancePolicy() BulkIssuancePolicy`

GetBulkIssuancePolicy returns the BulkIssuancePolicy field if non-nil, zero value otherwise.

### GetBulkIssuancePolicyOk

`func (o *BulkIssuanceResponse) GetBulkIssuancePolicyOk() (*BulkIssuancePolicy, bool)`

GetBulkIssuancePolicyOk returns a tuple with the BulkIssuancePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkIssuancePolicy

`func (o *BulkIssuanceResponse) SetBulkIssuancePolicy(v BulkIssuancePolicy)`

SetBulkIssuancePolicy sets BulkIssuancePolicy field to given value.


### GetCardProductId

`func (o *BulkIssuanceResponse) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *BulkIssuanceResponse) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *BulkIssuanceResponse) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.


### GetCreationTime

`func (o *BulkIssuanceResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BulkIssuanceResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BulkIssuanceResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BulkIssuanceResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *BulkIssuanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkIssuanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkIssuanceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkIssuanceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkIssuanceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkIssuanceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetShipping

`func (o *BulkIssuanceResponse) GetShipping() BulkShipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *BulkIssuanceResponse) GetShippingOk() (*BulkShipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *BulkIssuanceResponse) SetShipping(v BulkShipping)`

SetShipping sets Shipping field to given value.


### GetTenant

`func (o *BulkIssuanceResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkIssuanceResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkIssuanceResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


