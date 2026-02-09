# DebitCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardDetails** | [**DebitCardResponseDetails**](DebitCardResponseDetails.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**Type** | [**CardType**](CardType.md) |  | 

## Methods

### NewDebitCardResponse

`func NewDebitCardResponse(cardDetails DebitCardResponseDetails, tenant string, type_ CardType, ) *DebitCardResponse`

NewDebitCardResponse instantiates a new DebitCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebitCardResponseWithDefaults

`func NewDebitCardResponseWithDefaults() *DebitCardResponse`

NewDebitCardResponseWithDefaults instantiates a new DebitCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardDetails

`func (o *DebitCardResponse) GetCardDetails() DebitCardResponseDetails`

GetCardDetails returns the CardDetails field if non-nil, zero value otherwise.

### GetCardDetailsOk

`func (o *DebitCardResponse) GetCardDetailsOk() (*DebitCardResponseDetails, bool)`

GetCardDetailsOk returns a tuple with the CardDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDetails

`func (o *DebitCardResponse) SetCardDetails(v DebitCardResponseDetails)`

SetCardDetails sets CardDetails field to given value.


### GetTenant

`func (o *DebitCardResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DebitCardResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DebitCardResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetType

`func (o *DebitCardResponse) GetType() CardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DebitCardResponse) GetTypeOk() (*CardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DebitCardResponse) SetType(v CardType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


