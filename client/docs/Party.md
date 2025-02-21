# Party

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | account number of the person | [optional] [readonly] 
**Address** | Pointer to [**Address3**](Address3.md) |  | [optional] 
**AddressLines** | Pointer to [**AddressLines**](AddressLines.md) |  | [optional] 
**AlternateIdentifier** | Pointer to **string** | alternate identifier of the party, used in place of account number | [optional] [readonly] 
**BankName** | Pointer to **string** | name of the bank this party is a member of | [optional] [readonly] 
**IdentifierType** | Pointer to **string** | type of value used to identify the party | [optional] [readonly] 
**Name** | Pointer to **string** | name of the person | [optional] [readonly] 
**RoutingNumber** | Pointer to **string** | routing number of the bank this person is a member of | [optional] [readonly] 

## Methods

### NewParty

`func NewParty() *Party`

NewParty instantiates a new Party object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyWithDefaults

`func NewPartyWithDefaults() *Party`

NewPartyWithDefaults instantiates a new Party object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *Party) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Party) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Party) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Party) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAddress

`func (o *Party) GetAddress() Address3`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Party) GetAddressOk() (*Address3, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Party) SetAddress(v Address3)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Party) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressLines

`func (o *Party) GetAddressLines() AddressLines`

GetAddressLines returns the AddressLines field if non-nil, zero value otherwise.

### GetAddressLinesOk

`func (o *Party) GetAddressLinesOk() (*AddressLines, bool)`

GetAddressLinesOk returns a tuple with the AddressLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLines

`func (o *Party) SetAddressLines(v AddressLines)`

SetAddressLines sets AddressLines field to given value.

### HasAddressLines

`func (o *Party) HasAddressLines() bool`

HasAddressLines returns a boolean if a field has been set.

### GetAlternateIdentifier

`func (o *Party) GetAlternateIdentifier() string`

GetAlternateIdentifier returns the AlternateIdentifier field if non-nil, zero value otherwise.

### GetAlternateIdentifierOk

`func (o *Party) GetAlternateIdentifierOk() (*string, bool)`

GetAlternateIdentifierOk returns a tuple with the AlternateIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateIdentifier

`func (o *Party) SetAlternateIdentifier(v string)`

SetAlternateIdentifier sets AlternateIdentifier field to given value.

### HasAlternateIdentifier

`func (o *Party) HasAlternateIdentifier() bool`

HasAlternateIdentifier returns a boolean if a field has been set.

### GetBankName

`func (o *Party) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *Party) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *Party) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *Party) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetIdentifierType

`func (o *Party) GetIdentifierType() string`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *Party) GetIdentifierTypeOk() (*string, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *Party) SetIdentifierType(v string)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *Party) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetName

`func (o *Party) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Party) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Party) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Party) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *Party) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *Party) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *Party) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *Party) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


