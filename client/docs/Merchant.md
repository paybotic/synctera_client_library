# Merchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**Address1**](Address1.md) |  | 
**Email** | Pointer to **string** | Merchant&#39;s email | [optional] 
**Name** | **string** | Merchant&#39;s name | 
**PhoneNumber** | Pointer to **string** | Merchant&#39;s phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated. | [optional] 

## Methods

### NewMerchant

`func NewMerchant(address Address1, name string, ) *Merchant`

NewMerchant instantiates a new Merchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWithDefaults

`func NewMerchantWithDefaults() *Merchant`

NewMerchantWithDefaults instantiates a new Merchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Merchant) GetAddress() Address1`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Merchant) GetAddressOk() (*Address1, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Merchant) SetAddress(v Address1)`

SetAddress sets Address field to given value.


### GetEmail

`func (o *Merchant) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Merchant) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Merchant) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Merchant) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Merchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Merchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Merchant) SetName(v string)`

SetName sets Name field to given value.


### GetPhoneNumber

`func (o *Merchant) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Merchant) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Merchant) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Merchant) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


