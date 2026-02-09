# CreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardDetails** | [**CreditCardRequestDetails**](CreditCardRequestDetails.md) |  | 
**Type** | [**CardType**](CardType.md) |  | 

## Methods

### NewCreditCard

`func NewCreditCard(cardDetails CreditCardRequestDetails, type_ CardType, ) *CreditCard`

NewCreditCard instantiates a new CreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardWithDefaults

`func NewCreditCardWithDefaults() *CreditCard`

NewCreditCardWithDefaults instantiates a new CreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardDetails

`func (o *CreditCard) GetCardDetails() CreditCardRequestDetails`

GetCardDetails returns the CardDetails field if non-nil, zero value otherwise.

### GetCardDetailsOk

`func (o *CreditCard) GetCardDetailsOk() (*CreditCardRequestDetails, bool)`

GetCardDetailsOk returns a tuple with the CardDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDetails

`func (o *CreditCard) SetCardDetails(v CreditCardRequestDetails)`

SetCardDetails sets CardDetails field to given value.


### GetType

`func (o *CreditCard) GetType() CardType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditCard) GetTypeOk() (*CardType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditCard) SetType(v CardType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


