# CashOrderAuthorizationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | The account number of the client business the cash order is for. Known as \&quot;Location ID\&quot; in the cash order CSV files.  | [optional] 
**AuthorizationType** | [**CashAuthorizationType**](CashAuthorizationType.md) |  | 
**ClientName** | Pointer to **string** | The name of the client business the cash order is for. | [optional] 
**OrderDate** | Pointer to **string** | The date the cash order was placed with Empyreal | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCashOrderAuthorizationPatch

`func NewCashOrderAuthorizationPatch(authorizationType CashAuthorizationType, ) *CashOrderAuthorizationPatch`

NewCashOrderAuthorizationPatch instantiates a new CashOrderAuthorizationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashOrderAuthorizationPatchWithDefaults

`func NewCashOrderAuthorizationPatchWithDefaults() *CashOrderAuthorizationPatch`

NewCashOrderAuthorizationPatchWithDefaults instantiates a new CashOrderAuthorizationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *CashOrderAuthorizationPatch) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CashOrderAuthorizationPatch) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CashOrderAuthorizationPatch) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CashOrderAuthorizationPatch) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAuthorizationType

`func (o *CashOrderAuthorizationPatch) GetAuthorizationType() CashAuthorizationType`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *CashOrderAuthorizationPatch) GetAuthorizationTypeOk() (*CashAuthorizationType, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *CashOrderAuthorizationPatch) SetAuthorizationType(v CashAuthorizationType)`

SetAuthorizationType sets AuthorizationType field to given value.


### GetClientName

`func (o *CashOrderAuthorizationPatch) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *CashOrderAuthorizationPatch) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *CashOrderAuthorizationPatch) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *CashOrderAuthorizationPatch) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetOrderDate

`func (o *CashOrderAuthorizationPatch) GetOrderDate() string`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *CashOrderAuthorizationPatch) GetOrderDateOk() (*string, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *CashOrderAuthorizationPatch) SetOrderDate(v string)`

SetOrderDate sets OrderDate field to given value.

### HasOrderDate

`func (o *CashOrderAuthorizationPatch) HasOrderDate() bool`

HasOrderDate returns a boolean if a field has been set.

### GetStatus

`func (o *CashOrderAuthorizationPatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CashOrderAuthorizationPatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CashOrderAuthorizationPatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CashOrderAuthorizationPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


