# RoutingIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** | The address of the institution with the routing number specified in the routing_idetifier object | [optional] 
**PaymentRails** | Pointer to **[]string** | The supported payment types by the routing number specified in routing_identifier | [optional] 
**RoutingNumber** | **string** | A sequence of digits used to identify specific financial institution | 

## Methods

### NewRoutingIdentifier

`func NewRoutingIdentifier(routingNumber string, ) *RoutingIdentifier`

NewRoutingIdentifier instantiates a new RoutingIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingIdentifierWithDefaults

`func NewRoutingIdentifierWithDefaults() *RoutingIdentifier`

NewRoutingIdentifierWithDefaults instantiates a new RoutingIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *RoutingIdentifier) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RoutingIdentifier) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RoutingIdentifier) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *RoutingIdentifier) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPaymentRails

`func (o *RoutingIdentifier) GetPaymentRails() []string`

GetPaymentRails returns the PaymentRails field if non-nil, zero value otherwise.

### GetPaymentRailsOk

`func (o *RoutingIdentifier) GetPaymentRailsOk() (*[]string, bool)`

GetPaymentRailsOk returns a tuple with the PaymentRails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRails

`func (o *RoutingIdentifier) SetPaymentRails(v []string)`

SetPaymentRails sets PaymentRails field to given value.

### HasPaymentRails

`func (o *RoutingIdentifier) HasPaymentRails() bool`

HasPaymentRails returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *RoutingIdentifier) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *RoutingIdentifier) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *RoutingIdentifier) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


