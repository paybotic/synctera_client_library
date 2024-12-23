# AddAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifiers** | [**AddAccountsRequestAccountIdentifiers**](AddAccountsRequestAccountIdentifiers.md) |  | 
**AccountOwnerNames** | **[]string** | The names of the account owners. | 
**BusinessId** | Pointer to **string** | The identifier for the business customer associated with this external account. Exactly one of &#x60;business_id&#x60; or &#x60;customer_id&#x60; must be specified.  | [optional] 
**Currency** | Pointer to **string** | ISO 4217 alphabetic currency code | [optional] 
**CustomerId** | Pointer to **string** | The identifier for the personal customer associated with this external account. Exactly one of &#x60;customer_id&#x60; or &#x60;business_id&#x60; must be specified.  | [optional] 
**CustomerType** | [**ExtAccountCustomerType**](ExtAccountCustomerType.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** | User-supplied metadata | [optional] 
**Nickname** | Pointer to **string** | A user-meaningful name for the account | [optional] 
**RoutingIdentifiers** | [**AddAccountsRequestRoutingIdentifiers**](AddAccountsRequestRoutingIdentifiers.md) |  | 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | [optional] 
**Type** | **string** | The type of the account | 
**VendorAccountId** | Pointer to **string** | The ID of the vendor account, will be empty for MANUAL vendor | [optional] 
**Verification** | Pointer to [**NullableAccountVerification**](AccountVerification.md) |  | [optional] 

## Methods

### NewAddAccountsRequest

`func NewAddAccountsRequest(accountIdentifiers AddAccountsRequestAccountIdentifiers, accountOwnerNames []string, customerType ExtAccountCustomerType, routingIdentifiers AddAccountsRequestRoutingIdentifiers, type_ string, ) *AddAccountsRequest`

NewAddAccountsRequest instantiates a new AddAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccountsRequestWithDefaults

`func NewAddAccountsRequestWithDefaults() *AddAccountsRequest`

NewAddAccountsRequestWithDefaults instantiates a new AddAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIdentifiers

`func (o *AddAccountsRequest) GetAccountIdentifiers() AddAccountsRequestAccountIdentifiers`

GetAccountIdentifiers returns the AccountIdentifiers field if non-nil, zero value otherwise.

### GetAccountIdentifiersOk

`func (o *AddAccountsRequest) GetAccountIdentifiersOk() (*AddAccountsRequestAccountIdentifiers, bool)`

GetAccountIdentifiersOk returns a tuple with the AccountIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifiers

`func (o *AddAccountsRequest) SetAccountIdentifiers(v AddAccountsRequestAccountIdentifiers)`

SetAccountIdentifiers sets AccountIdentifiers field to given value.


### GetAccountOwnerNames

`func (o *AddAccountsRequest) GetAccountOwnerNames() []string`

GetAccountOwnerNames returns the AccountOwnerNames field if non-nil, zero value otherwise.

### GetAccountOwnerNamesOk

`func (o *AddAccountsRequest) GetAccountOwnerNamesOk() (*[]string, bool)`

GetAccountOwnerNamesOk returns a tuple with the AccountOwnerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwnerNames

`func (o *AddAccountsRequest) SetAccountOwnerNames(v []string)`

SetAccountOwnerNames sets AccountOwnerNames field to given value.


### GetBusinessId

`func (o *AddAccountsRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AddAccountsRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AddAccountsRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AddAccountsRequest) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCurrency

`func (o *AddAccountsRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AddAccountsRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AddAccountsRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AddAccountsRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *AddAccountsRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AddAccountsRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AddAccountsRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AddAccountsRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerType

`func (o *AddAccountsRequest) GetCustomerType() ExtAccountCustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AddAccountsRequest) GetCustomerTypeOk() (*ExtAccountCustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AddAccountsRequest) SetCustomerType(v ExtAccountCustomerType)`

SetCustomerType sets CustomerType field to given value.


### GetMetadata

`func (o *AddAccountsRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddAccountsRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddAccountsRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddAccountsRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNickname

`func (o *AddAccountsRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AddAccountsRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AddAccountsRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AddAccountsRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetRoutingIdentifiers

`func (o *AddAccountsRequest) GetRoutingIdentifiers() AddAccountsRequestRoutingIdentifiers`

GetRoutingIdentifiers returns the RoutingIdentifiers field if non-nil, zero value otherwise.

### GetRoutingIdentifiersOk

`func (o *AddAccountsRequest) GetRoutingIdentifiersOk() (*AddAccountsRequestRoutingIdentifiers, bool)`

GetRoutingIdentifiersOk returns a tuple with the RoutingIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifiers

`func (o *AddAccountsRequest) SetRoutingIdentifiers(v AddAccountsRequestRoutingIdentifiers)`

SetRoutingIdentifiers sets RoutingIdentifiers field to given value.


### GetTenant

`func (o *AddAccountsRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AddAccountsRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AddAccountsRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AddAccountsRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetType

`func (o *AddAccountsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddAccountsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddAccountsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetVendorAccountId

`func (o *AddAccountsRequest) GetVendorAccountId() string`

GetVendorAccountId returns the VendorAccountId field if non-nil, zero value otherwise.

### GetVendorAccountIdOk

`func (o *AddAccountsRequest) GetVendorAccountIdOk() (*string, bool)`

GetVendorAccountIdOk returns a tuple with the VendorAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccountId

`func (o *AddAccountsRequest) SetVendorAccountId(v string)`

SetVendorAccountId sets VendorAccountId field to given value.

### HasVendorAccountId

`func (o *AddAccountsRequest) HasVendorAccountId() bool`

HasVendorAccountId returns a boolean if a field has been set.

### GetVerification

`func (o *AddAccountsRequest) GetVerification() AccountVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *AddAccountsRequest) GetVerificationOk() (*AccountVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *AddAccountsRequest) SetVerification(v AccountVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *AddAccountsRequest) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### SetVerificationNil

`func (o *AddAccountsRequest) SetVerificationNil(b bool)`

 SetVerificationNil sets the value for Verification to be an explicit nil

### UnsetVerification
`func (o *AddAccountsRequest) UnsetVerification()`

UnsetVerification ensures that no value is present for Verification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


