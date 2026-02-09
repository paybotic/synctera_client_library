# BulkIssuanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkIssuancePolicy** | [**BulkIssuancePolicy**](BulkIssuancePolicy.md) |  | [default to BULKISSUANCEPOLICY_AUTO]
**CardProductId** | **string** | The unique identifier of a cards product | 
**Name** | **string** | Name associated with the bulk order configuration. | 
**Shipping** | [**BulkShipping**](BulkShipping.md) |  | 

## Methods

### NewBulkIssuanceRequest

`func NewBulkIssuanceRequest(bulkIssuancePolicy BulkIssuancePolicy, cardProductId string, name string, shipping BulkShipping, ) *BulkIssuanceRequest`

NewBulkIssuanceRequest instantiates a new BulkIssuanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIssuanceRequestWithDefaults

`func NewBulkIssuanceRequestWithDefaults() *BulkIssuanceRequest`

NewBulkIssuanceRequestWithDefaults instantiates a new BulkIssuanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkIssuancePolicy

`func (o *BulkIssuanceRequest) GetBulkIssuancePolicy() BulkIssuancePolicy`

GetBulkIssuancePolicy returns the BulkIssuancePolicy field if non-nil, zero value otherwise.

### GetBulkIssuancePolicyOk

`func (o *BulkIssuanceRequest) GetBulkIssuancePolicyOk() (*BulkIssuancePolicy, bool)`

GetBulkIssuancePolicyOk returns a tuple with the BulkIssuancePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkIssuancePolicy

`func (o *BulkIssuanceRequest) SetBulkIssuancePolicy(v BulkIssuancePolicy)`

SetBulkIssuancePolicy sets BulkIssuancePolicy field to given value.


### GetCardProductId

`func (o *BulkIssuanceRequest) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *BulkIssuanceRequest) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *BulkIssuanceRequest) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.


### GetName

`func (o *BulkIssuanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkIssuanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkIssuanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShipping

`func (o *BulkIssuanceRequest) GetShipping() BulkShipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *BulkIssuanceRequest) GetShippingOk() (*BulkShipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *BulkIssuanceRequest) SetShipping(v BulkShipping)`

SetShipping sets Shipping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


