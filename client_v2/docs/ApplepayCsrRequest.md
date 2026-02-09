# ApplepayCsrRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The merchant ID for which the CSR is being created. This is the same merchant ID used in the Apple Developer Portal and will be the one registered to accept payments in your iOS app.  The merchant ID must be in the format &#x60;merchant.com.[your-company-name]&#x60;.  | 
**OrganizationName** | **string** | The name of the organization that owns the merchant ID. This is typically the legal name of your company or organization.  | 

## Methods

### NewApplepayCsrRequest

`func NewApplepayCsrRequest(merchantId string, organizationName string, ) *ApplepayCsrRequest`

NewApplepayCsrRequest instantiates a new ApplepayCsrRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplepayCsrRequestWithDefaults

`func NewApplepayCsrRequestWithDefaults() *ApplepayCsrRequest`

NewApplepayCsrRequestWithDefaults instantiates a new ApplepayCsrRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *ApplepayCsrRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *ApplepayCsrRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *ApplepayCsrRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetOrganizationName

`func (o *ApplepayCsrRequest) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ApplepayCsrRequest) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ApplepayCsrRequest) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


