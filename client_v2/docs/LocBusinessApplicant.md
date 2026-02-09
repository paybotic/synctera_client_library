# LocBusinessApplicant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnualRevenue** | Pointer to **int32** | Annual revenue in cents (e.g., 500000000 &#x3D; $5,000,000) | [optional] 
**BusinessId** | **string** | Business ID for the applicant | 

## Methods

### NewLocBusinessApplicant

`func NewLocBusinessApplicant(businessId string, ) *LocBusinessApplicant`

NewLocBusinessApplicant instantiates a new LocBusinessApplicant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocBusinessApplicantWithDefaults

`func NewLocBusinessApplicantWithDefaults() *LocBusinessApplicant`

NewLocBusinessApplicantWithDefaults instantiates a new LocBusinessApplicant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnualRevenue

`func (o *LocBusinessApplicant) GetAnnualRevenue() int32`

GetAnnualRevenue returns the AnnualRevenue field if non-nil, zero value otherwise.

### GetAnnualRevenueOk

`func (o *LocBusinessApplicant) GetAnnualRevenueOk() (*int32, bool)`

GetAnnualRevenueOk returns a tuple with the AnnualRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualRevenue

`func (o *LocBusinessApplicant) SetAnnualRevenue(v int32)`

SetAnnualRevenue sets AnnualRevenue field to given value.

### HasAnnualRevenue

`func (o *LocBusinessApplicant) HasAnnualRevenue() bool`

HasAnnualRevenue returns a boolean if a field has been set.

### GetBusinessId

`func (o *LocBusinessApplicant) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *LocBusinessApplicant) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *LocBusinessApplicant) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


