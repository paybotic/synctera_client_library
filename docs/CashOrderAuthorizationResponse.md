# CashOrderAuthorizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transfer amount in cents | 
**AuthorizationType** | [**CashAuthorizationType**](CashAuthorizationType.md) |  | 
**DestinationAccountId** | **string** | The UUID of the Synctera account resource that is the destination of the transfer.  | 
**OrderDate** | **string** | The date the cash order was placed with cash distribution provider | 
**AccountNumber** | **string** | The account number of the client business the cash order is for. Known as \&quot;Location ID\&quot; in the cash order CSV files.  | 
**ClientName** | **string** | The name of the client business the cash order is for. | 
**HoldId** | **string** | ID of the hold created for this authorization | 
**Id** | **string** | ID of the transfer | 
**OriginatingAccountId** | **string** | The UUID of the Synctera account resource originating the transfer.  | 
**Status** | **string** |  | 
**Tenant** | **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | 

## Methods

### NewCashOrderAuthorizationResponse

`func NewCashOrderAuthorizationResponse(amount int64, authorizationType CashAuthorizationType, destinationAccountId string, orderDate string, accountNumber string, clientName string, holdId string, id string, originatingAccountId string, status string, tenant string, ) *CashOrderAuthorizationResponse`

NewCashOrderAuthorizationResponse instantiates a new CashOrderAuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashOrderAuthorizationResponseWithDefaults

`func NewCashOrderAuthorizationResponseWithDefaults() *CashOrderAuthorizationResponse`

NewCashOrderAuthorizationResponseWithDefaults instantiates a new CashOrderAuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CashOrderAuthorizationResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashOrderAuthorizationResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashOrderAuthorizationResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetAuthorizationType

`func (o *CashOrderAuthorizationResponse) GetAuthorizationType() CashAuthorizationType`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *CashOrderAuthorizationResponse) GetAuthorizationTypeOk() (*CashAuthorizationType, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *CashOrderAuthorizationResponse) SetAuthorizationType(v CashAuthorizationType)`

SetAuthorizationType sets AuthorizationType field to given value.


### GetDestinationAccountId

`func (o *CashOrderAuthorizationResponse) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *CashOrderAuthorizationResponse) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *CashOrderAuthorizationResponse) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.


### GetOrderDate

`func (o *CashOrderAuthorizationResponse) GetOrderDate() string`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *CashOrderAuthorizationResponse) GetOrderDateOk() (*string, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *CashOrderAuthorizationResponse) SetOrderDate(v string)`

SetOrderDate sets OrderDate field to given value.


### GetAccountNumber

`func (o *CashOrderAuthorizationResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CashOrderAuthorizationResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CashOrderAuthorizationResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetClientName

`func (o *CashOrderAuthorizationResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *CashOrderAuthorizationResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *CashOrderAuthorizationResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetHoldId

`func (o *CashOrderAuthorizationResponse) GetHoldId() string`

GetHoldId returns the HoldId field if non-nil, zero value otherwise.

### GetHoldIdOk

`func (o *CashOrderAuthorizationResponse) GetHoldIdOk() (*string, bool)`

GetHoldIdOk returns a tuple with the HoldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldId

`func (o *CashOrderAuthorizationResponse) SetHoldId(v string)`

SetHoldId sets HoldId field to given value.


### GetId

`func (o *CashOrderAuthorizationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CashOrderAuthorizationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CashOrderAuthorizationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetOriginatingAccountId

`func (o *CashOrderAuthorizationResponse) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *CashOrderAuthorizationResponse) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *CashOrderAuthorizationResponse) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetStatus

`func (o *CashOrderAuthorizationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CashOrderAuthorizationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CashOrderAuthorizationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *CashOrderAuthorizationResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CashOrderAuthorizationResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CashOrderAuthorizationResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


