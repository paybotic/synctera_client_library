# CustomerServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address2**](Address2.md) |  | [optional] 
**Email** | Pointer to **string** | The customer service email address | [optional] 
**PhoneNumber** | Pointer to **string** | The customer service phone number | [optional] 

## Methods

### NewCustomerServiceDetails

`func NewCustomerServiceDetails() *CustomerServiceDetails`

NewCustomerServiceDetails instantiates a new CustomerServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceDetailsWithDefaults

`func NewCustomerServiceDetailsWithDefaults() *CustomerServiceDetails`

NewCustomerServiceDetailsWithDefaults instantiates a new CustomerServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CustomerServiceDetails) GetAddress() Address2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerServiceDetails) GetAddressOk() (*Address2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerServiceDetails) SetAddress(v Address2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerServiceDetails) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerServiceDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerServiceDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerServiceDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerServiceDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CustomerServiceDetails) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerServiceDetails) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerServiceDetails) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerServiceDetails) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


