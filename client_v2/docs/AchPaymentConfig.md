# AchPaymentConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalAccountId** | **string** | The external account ID to debit via ACH | 
**IsSameDay** | Pointer to **bool** | Whether to send as same-day ACH (default false) | [optional] 
**SecCode** | Pointer to **string** | Standard Entry Class Code (WEB, CCD, or PPD). Default is WEB. | [optional] 

## Methods

### NewAchPaymentConfig

`func NewAchPaymentConfig(externalAccountId string, ) *AchPaymentConfig`

NewAchPaymentConfig instantiates a new AchPaymentConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchPaymentConfigWithDefaults

`func NewAchPaymentConfigWithDefaults() *AchPaymentConfig`

NewAchPaymentConfigWithDefaults instantiates a new AchPaymentConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalAccountId

`func (o *AchPaymentConfig) GetExternalAccountId() string`

GetExternalAccountId returns the ExternalAccountId field if non-nil, zero value otherwise.

### GetExternalAccountIdOk

`func (o *AchPaymentConfig) GetExternalAccountIdOk() (*string, bool)`

GetExternalAccountIdOk returns a tuple with the ExternalAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAccountId

`func (o *AchPaymentConfig) SetExternalAccountId(v string)`

SetExternalAccountId sets ExternalAccountId field to given value.


### GetIsSameDay

`func (o *AchPaymentConfig) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *AchPaymentConfig) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *AchPaymentConfig) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.

### HasIsSameDay

`func (o *AchPaymentConfig) HasIsSameDay() bool`

HasIsSameDay returns a boolean if a field has been set.

### GetSecCode

`func (o *AchPaymentConfig) GetSecCode() string`

GetSecCode returns the SecCode field if non-nil, zero value otherwise.

### GetSecCodeOk

`func (o *AchPaymentConfig) GetSecCodeOk() (*string, bool)`

GetSecCodeOk returns a tuple with the SecCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecCode

`func (o *AchPaymentConfig) SetSecCode(v string)`

SetSecCode sets SecCode field to given value.

### HasSecCode

`func (o *AchPaymentConfig) HasSecCode() bool`

HasSecCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


