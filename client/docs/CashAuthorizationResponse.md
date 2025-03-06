# CashAuthorizationResponse

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

### NewCashAuthorizationResponse

`func NewCashAuthorizationResponse(amount int64, authorizationType CashAuthorizationType, destinationAccountId string, orderDate string, accountNumber string, clientName string, holdId string, id string, originatingAccountId string, status string, tenant string, ) *CashAuthorizationResponse`

NewCashAuthorizationResponse instantiates a new CashAuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashAuthorizationResponseWithDefaults

`func NewCashAuthorizationResponseWithDefaults() *CashAuthorizationResponse`

NewCashAuthorizationResponseWithDefaults instantiates a new CashAuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CashAuthorizationResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CashAuthorizationResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CashAuthorizationResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetAuthorizationType

`func (o *CashAuthorizationResponse) GetAuthorizationType() CashAuthorizationType`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *CashAuthorizationResponse) GetAuthorizationTypeOk() (*CashAuthorizationType, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *CashAuthorizationResponse) SetAuthorizationType(v CashAuthorizationType)`

SetAuthorizationType sets AuthorizationType field to given value.


### GetDestinationAccountId

`func (o *CashAuthorizationResponse) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *CashAuthorizationResponse) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *CashAuthorizationResponse) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.


### GetOrderDate

`func (o *CashAuthorizationResponse) GetOrderDate() string`

GetOrderDate returns the OrderDate field if non-nil, zero value otherwise.

### GetOrderDateOk

`func (o *CashAuthorizationResponse) GetOrderDateOk() (*string, bool)`

GetOrderDateOk returns a tuple with the OrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDate

`func (o *CashAuthorizationResponse) SetOrderDate(v string)`

SetOrderDate sets OrderDate field to given value.


### GetAccountNumber

`func (o *CashAuthorizationResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CashAuthorizationResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CashAuthorizationResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetClientName

`func (o *CashAuthorizationResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *CashAuthorizationResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *CashAuthorizationResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetHoldId

`func (o *CashAuthorizationResponse) GetHoldId() string`

GetHoldId returns the HoldId field if non-nil, zero value otherwise.

### GetHoldIdOk

`func (o *CashAuthorizationResponse) GetHoldIdOk() (*string, bool)`

GetHoldIdOk returns a tuple with the HoldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldId

`func (o *CashAuthorizationResponse) SetHoldId(v string)`

SetHoldId sets HoldId field to given value.


### GetId

`func (o *CashAuthorizationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CashAuthorizationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CashAuthorizationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetOriginatingAccountId

`func (o *CashAuthorizationResponse) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *CashAuthorizationResponse) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *CashAuthorizationResponse) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetStatus

`func (o *CashAuthorizationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CashAuthorizationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CashAuthorizationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *CashAuthorizationResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CashAuthorizationResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CashAuthorizationResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


