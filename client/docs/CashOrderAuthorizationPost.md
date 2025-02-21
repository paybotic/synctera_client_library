# CashOrderAuthorizationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transfer amount in cents | 
**AuthorizationType** | [**CashAuthorizationType**](CashAuthorizationType.md) |  | 
**DestinationAccountId** | **string** | The UUID of the Synctera account resource that is the destination of the transfer.  | 
**OrderDate** | **string** | The date the cash order was placed with cash distribution provider | 

## Methods

### NewCashOrderAuthorizationPost

`func NewCashOrderAuthorizationPost(amount int64, authorizationType CashAuthorizationType, destinationAccountId string, orderDate string, ) *CashOrderAuthorizationPost`

NewCashOrderAuthorizationPost instantiates a new CashOrderAuthorizationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashOrderAuthorizationPostWithDefaults

`func NewCashOrderAuthorizationPostWithDefaults() *CashOrderAuthorizationPost`

NewCashOrderAuthorizationPostWithDefaults instantiates a new CashOrderAuthorizationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CashOrderAuthorizationPost) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashOrderAuthorizationPost) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashOrderAuthorizationPost) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetAuthorizationType

`func (o *CashOrderAuthorizationPost) GetAuthorizationType() CashAuthorizationType`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *CashOrderAuthorizationPost) GetAuthorizationTypeOk() (*CashAuthorizationType, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *CashOrderAuthorizationPost) SetAuthorizationType(v CashAuthorizationType)`

SetAuthorizationType sets AuthorizationType field to given value.


### GetDestinationAccountId

`func (o *CashOrderAuthorizationPost) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *CashOrderAuthorizationPost) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *CashOrderAuthorizationPost) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.


### GetOrderDate

`func (o *CashOrderAuthorizationPost) GetOrderDate() string`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *CashOrderAuthorizationPost) GetOrderDateOk() (*string, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *CashOrderAuthorizationPost) SetOrderDate(v string)`

SetOrderDate sets OrderDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


