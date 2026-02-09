# BulkShipping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**Address**](Address.md) | The address to which the bulk order will be shipped. | 
**BusinessName** | Pointer to **string** | The name of the business which the bulk order will be shipped | [optional] 
**IsExpeditedFulfillment** | Pointer to **bool** | Is the shipment expedited | [optional] 
**Method** | [**BulkShippingMethod**](BulkShippingMethod.md) |  | [default to BULKSHIPPINGMETHOD_TWO_DAY]
**PhoneNumber** | Pointer to **string** | The phone number of the recipient | [optional] 
**RecipientName** | [**RecipientName**](RecipientName.md) | The name of the recipient to whom the bulk order will be shipped | 

## Methods

### NewBulkShipping

`func NewBulkShipping(address Address, method BulkShippingMethod, recipientName RecipientName, ) *BulkShipping`

NewBulkShipping instantiates a new BulkShipping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkShippingWithDefaults

`func NewBulkShippingWithDefaults() *BulkShipping`

NewBulkShippingWithDefaults instantiates a new BulkShipping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BulkShipping) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BulkShipping) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BulkShipping) SetAddress(v Address)`

SetAddress sets Address field to given value.


### GetBusinessName

`func (o *BulkShipping) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *BulkShipping) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *BulkShipping) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *BulkShipping) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetIsExpeditedFulfillment

`func (o *BulkShipping) GetIsExpeditedFulfillment() bool`

GetIsExpeditedFulfillment returns the IsExpeditedFulfillment field if non-nil, zero value otherwise.

### GetIsExpeditedFulfillmentOk

`func (o *BulkShipping) GetIsExpeditedFulfillmentOk() (*bool, bool)`

GetIsExpeditedFulfillmentOk returns a tuple with the IsExpeditedFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpeditedFulfillment

`func (o *BulkShipping) SetIsExpeditedFulfillment(v bool)`

SetIsExpeditedFulfillment sets IsExpeditedFulfillment field to given value.

### HasIsExpeditedFulfillment

`func (o *BulkShipping) HasIsExpeditedFulfillment() bool`

HasIsExpeditedFulfillment returns a boolean if a field has been set.

### GetMethod

`func (o *BulkShipping) GetMethod() BulkShippingMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *BulkShipping) GetMethodOk() (*BulkShippingMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *BulkShipping) SetMethod(v BulkShippingMethod)`

SetMethod sets Method field to given value.


### GetPhoneNumber

`func (o *BulkShipping) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BulkShipping) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BulkShipping) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BulkShipping) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRecipientName

`func (o *BulkShipping) GetRecipientName() RecipientName`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *BulkShipping) GetRecipientNameOk() (*RecipientName, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *BulkShipping) SetRecipientName(v RecipientName)`

SetRecipientName sets RecipientName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


