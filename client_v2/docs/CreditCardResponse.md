# CreditCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardDetails** | [**CreditCardResponseDetails**](CreditCardResponseDetails.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**Type** | [**CardType**](CardType.md) |  | 
**VendorData** | Pointer to [**CardResponseVendorData**](CardResponseVendorData.md) |  | [optional] 

## Methods

### NewCreditCardResponse

`func NewCreditCardResponse(cardDetails CreditCardResponseDetails, tenant string, type_ CardType, ) *CreditCardResponse`

NewCreditCardResponse instantiates a new CreditCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardResponseWithDefaults

`func NewCreditCardResponseWithDefaults() *CreditCardResponse`

NewCreditCardResponseWithDefaults instantiates a new CreditCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardDetails

`func (o *CreditCardResponse) GetCardDetails() CreditCardResponseDetails`

GetCardDetails returns the CardDetails field if non-nil, zero value otherwise.

### GetCardDetailsOk

`func (o *CreditCardResponse) GetCardDetailsOk() (*CreditCardResponseDetails, bool)`

GetCardDetailsOk returns a tuple with the CardDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDetails

`func (o *CreditCardResponse) SetCardDetails(v CreditCardResponseDetails)`

SetCardDetails sets CardDetails field to given value.


### GetTenant

`func (o *CreditCardResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CreditCardResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CreditCardResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetType

`func (o *CreditCardResponse) GetType() CardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditCardResponse) GetTypeOk() (*CardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditCardResponse) SetType(v CardType)`

SetType sets Type field to given value.


### GetVendorData

`func (o *CreditCardResponse) GetVendorData() CardResponseVendorData`

GetVendorData returns the VendorData field if non-nil, zero value otherwise.

### GetVendorDataOk

`func (o *CreditCardResponse) GetVendorDataOk() (*CardResponseVendorData, bool)`

GetVendorDataOk returns a tuple with the VendorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorData

`func (o *CreditCardResponse) SetVendorData(v CardResponseVendorData)`

SetVendorData sets VendorData field to given value.

### HasVendorData

`func (o *CreditCardResponse) HasVendorData() bool`

HasVendorData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


