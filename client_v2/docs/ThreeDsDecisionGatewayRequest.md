# ThreeDsDecisionGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsTransactionId** | **string** | Universally unique transaction identifier assigned by the ACS to identify a single transaction. | 
**AuthenticationRequestType** | Pointer to [**ThreeDsAuthenticationRequestType**](ThreeDsAuthenticationRequestType.md) |  | [optional] 
**CardId** | **string** |  | 
**CardProductId** | **string** |  | 
**ClientIpAddress** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** | ISO 4217  Alpha-3 currency code | [optional] 
**DeviceChannel** | Pointer to [**ThreeDsDeviceChannel**](ThreeDsDeviceChannel.md) |  | [optional] 
**Merchant** | Pointer to [**ThreeDsDecisionGatewayRequestMerchant**](ThreeDsDecisionGatewayRequestMerchant.md) |  | [optional] 
**TransactionAmount** | Pointer to **int32** |  | [optional] 
**TransactionSubType** | Pointer to [**ThreeDsTransactionSubtype**](ThreeDsTransactionSubtype.md) |  | [optional] 
**TransactionType** | Pointer to [**ThreeDsTransactionType**](ThreeDsTransactionType.md) |  | [optional] 

## Methods

### NewThreeDsDecisionGatewayRequest

`func NewThreeDsDecisionGatewayRequest(acsTransactionId string, cardId string, cardProductId string, ) *ThreeDsDecisionGatewayRequest`

NewThreeDsDecisionGatewayRequest instantiates a new ThreeDsDecisionGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDsDecisionGatewayRequestWithDefaults

`func NewThreeDsDecisionGatewayRequestWithDefaults() *ThreeDsDecisionGatewayRequest`

NewThreeDsDecisionGatewayRequestWithDefaults instantiates a new ThreeDsDecisionGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsTransactionId

`func (o *ThreeDsDecisionGatewayRequest) GetAcsTransactionId() string`

GetAcsTransactionId returns the AcsTransactionId field if non-nil, zero value otherwise.

### GetAcsTransactionIdOk

`func (o *ThreeDsDecisionGatewayRequest) GetAcsTransactionIdOk() (*string, bool)`

GetAcsTransactionIdOk returns a tuple with the AcsTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsTransactionId

`func (o *ThreeDsDecisionGatewayRequest) SetAcsTransactionId(v string)`

SetAcsTransactionId sets AcsTransactionId field to given value.


### GetAuthenticationRequestType

`func (o *ThreeDsDecisionGatewayRequest) GetAuthenticationRequestType() ThreeDsAuthenticationRequestType`

GetAuthenticationRequestType returns the AuthenticationRequestType field if non-nil, zero value otherwise.

### GetAuthenticationRequestTypeOk

`func (o *ThreeDsDecisionGatewayRequest) GetAuthenticationRequestTypeOk() (*ThreeDsAuthenticationRequestType, bool)`

GetAuthenticationRequestTypeOk returns a tuple with the AuthenticationRequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationRequestType

`func (o *ThreeDsDecisionGatewayRequest) SetAuthenticationRequestType(v ThreeDsAuthenticationRequestType)`

SetAuthenticationRequestType sets AuthenticationRequestType field to given value.

### HasAuthenticationRequestType

`func (o *ThreeDsDecisionGatewayRequest) HasAuthenticationRequestType() bool`

HasAuthenticationRequestType returns a boolean if a field has been set.

### GetCardId

`func (o *ThreeDsDecisionGatewayRequest) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *ThreeDsDecisionGatewayRequest) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *ThreeDsDecisionGatewayRequest) SetCardId(v string)`

SetCardId sets CardId field to given value.


### GetCardProductId

`func (o *ThreeDsDecisionGatewayRequest) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *ThreeDsDecisionGatewayRequest) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *ThreeDsDecisionGatewayRequest) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.


### GetClientIpAddress

`func (o *ThreeDsDecisionGatewayRequest) GetClientIpAddress() string`

GetClientIpAddress returns the ClientIpAddress field if non-nil, zero value otherwise.

### GetClientIpAddressOk

`func (o *ThreeDsDecisionGatewayRequest) GetClientIpAddressOk() (*string, bool)`

GetClientIpAddressOk returns a tuple with the ClientIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpAddress

`func (o *ThreeDsDecisionGatewayRequest) SetClientIpAddress(v string)`

SetClientIpAddress sets ClientIpAddress field to given value.

### HasClientIpAddress

`func (o *ThreeDsDecisionGatewayRequest) HasClientIpAddress() bool`

HasClientIpAddress returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ThreeDsDecisionGatewayRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ThreeDsDecisionGatewayRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ThreeDsDecisionGatewayRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ThreeDsDecisionGatewayRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDeviceChannel

`func (o *ThreeDsDecisionGatewayRequest) GetDeviceChannel() ThreeDsDeviceChannel`

GetDeviceChannel returns the DeviceChannel field if non-nil, zero value otherwise.

### GetDeviceChannelOk

`func (o *ThreeDsDecisionGatewayRequest) GetDeviceChannelOk() (*ThreeDsDeviceChannel, bool)`

GetDeviceChannelOk returns a tuple with the DeviceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceChannel

`func (o *ThreeDsDecisionGatewayRequest) SetDeviceChannel(v ThreeDsDeviceChannel)`

SetDeviceChannel sets DeviceChannel field to given value.

### HasDeviceChannel

`func (o *ThreeDsDecisionGatewayRequest) HasDeviceChannel() bool`

HasDeviceChannel returns a boolean if a field has been set.

### GetMerchant

`func (o *ThreeDsDecisionGatewayRequest) GetMerchant() ThreeDsDecisionGatewayRequestMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *ThreeDsDecisionGatewayRequest) GetMerchantOk() (*ThreeDsDecisionGatewayRequestMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *ThreeDsDecisionGatewayRequest) SetMerchant(v ThreeDsDecisionGatewayRequestMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *ThreeDsDecisionGatewayRequest) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *ThreeDsDecisionGatewayRequest) GetTransactionAmount() int32`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *ThreeDsDecisionGatewayRequest) GetTransactionAmountOk() (*int32, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *ThreeDsDecisionGatewayRequest) SetTransactionAmount(v int32)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *ThreeDsDecisionGatewayRequest) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetTransactionSubType

`func (o *ThreeDsDecisionGatewayRequest) GetTransactionSubType() ThreeDsTransactionSubtype`

GetTransactionSubType returns the TransactionSubType field if non-nil, zero value otherwise.

### GetTransactionSubTypeOk

`func (o *ThreeDsDecisionGatewayRequest) GetTransactionSubTypeOk() (*ThreeDsTransactionSubtype, bool)`

GetTransactionSubTypeOk returns a tuple with the TransactionSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionSubType

`func (o *ThreeDsDecisionGatewayRequest) SetTransactionSubType(v ThreeDsTransactionSubtype)`

SetTransactionSubType sets TransactionSubType field to given value.

### HasTransactionSubType

`func (o *ThreeDsDecisionGatewayRequest) HasTransactionSubType() bool`

HasTransactionSubType returns a boolean if a field has been set.

### GetTransactionType

`func (o *ThreeDsDecisionGatewayRequest) GetTransactionType() ThreeDsTransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ThreeDsDecisionGatewayRequest) GetTransactionTypeOk() (*ThreeDsTransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ThreeDsDecisionGatewayRequest) SetTransactionType(v ThreeDsTransactionType)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ThreeDsDecisionGatewayRequest) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


