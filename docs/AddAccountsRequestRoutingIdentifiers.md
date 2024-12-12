# AddAccountsRequestRoutingIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchRoutingNumber** | Pointer to **string** | The routing number used for US ACH payments.  | [optional] 
**BankCountries** | **[]string** | The countries that this bank operates the account in | 
**BankName** | **string** | The name of the bank managing the account | 
**EftCaRoutingNumber** | Pointer to **string** | &gt; 🚧 Alpha &gt; This is an Alpha property. Feedback from the community is welcome. We may make breaking changes to this property. The 9 digit routing number used for EFT CA payments, identifying a Canadian bank. The format is 0xxxyyyyy where xxx is the institution number and yyyyy is the transit number. On write, Synctera will store the entire routing number; on read, we only return the last 4 characters.  | [optional] 
**InternationalWireDetails** | Pointer to [**InternationalWireDetails**](InternationalWireDetails.md) |  | [optional] 
**WireRoutingNumber** | Pointer to **string** | The routing number used for US domestic wire payments.  | [optional] 

## Methods

### NewAddAccountsRequestRoutingIdentifiers

`func NewAddAccountsRequestRoutingIdentifiers(bankCountries []string, bankName string, ) *AddAccountsRequestRoutingIdentifiers`

NewAddAccountsRequestRoutingIdentifiers instantiates a new AddAccountsRequestRoutingIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccountsRequestRoutingIdentifiersWithDefaults

`func NewAddAccountsRequestRoutingIdentifiersWithDefaults() *AddAccountsRequestRoutingIdentifiers`

NewAddAccountsRequestRoutingIdentifiersWithDefaults instantiates a new AddAccountsRequestRoutingIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) GetAchRoutingNumber() string`

GetAchRoutingNumber returns the AchRoutingNumber field if non-nil, zero value otherwise.

### GetAchRoutingNumberOk

`func (o *AddAccountsRequestRoutingIdentifiers) GetAchRoutingNumberOk() (*string, bool)`

GetAchRoutingNumberOk returns a tuple with the AchRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) SetAchRoutingNumber(v string)`

SetAchRoutingNumber sets AchRoutingNumber field to given value.

### HasAchRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) HasAchRoutingNumber() bool`

HasAchRoutingNumber returns a boolean if a field has been set.

### GetBankCountries

`func (o *AddAccountsRequestRoutingIdentifiers) GetBankCountries() []string`

GetBankCountries returns the BankCountries field if non-nil, zero value otherwise.

### GetBankCountriesOk

`func (o *AddAccountsRequestRoutingIdentifiers) GetBankCountriesOk() (*[]string, bool)`

GetBankCountriesOk returns a tuple with the BankCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountries

`func (o *AddAccountsRequestRoutingIdentifiers) SetBankCountries(v []string)`

SetBankCountries sets BankCountries field to given value.


### GetBankName

`func (o *AddAccountsRequestRoutingIdentifiers) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *AddAccountsRequestRoutingIdentifiers) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *AddAccountsRequestRoutingIdentifiers) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetEftCaRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) GetEftCaRoutingNumber() string`

GetEftCaRoutingNumber returns the EftCaRoutingNumber field if non-nil, zero value otherwise.

### GetEftCaRoutingNumberOk

`func (o *AddAccountsRequestRoutingIdentifiers) GetEftCaRoutingNumberOk() (*string, bool)`

GetEftCaRoutingNumberOk returns a tuple with the EftCaRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEftCaRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) SetEftCaRoutingNumber(v string)`

SetEftCaRoutingNumber sets EftCaRoutingNumber field to given value.

### HasEftCaRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) HasEftCaRoutingNumber() bool`

HasEftCaRoutingNumber returns a boolean if a field has been set.

### GetInternationalWireDetails

`func (o *AddAccountsRequestRoutingIdentifiers) GetInternationalWireDetails() InternationalWireDetails`

GetInternationalWireDetails returns the InternationalWireDetails field if non-nil, zero value otherwise.

### GetInternationalWireDetailsOk

`func (o *AddAccountsRequestRoutingIdentifiers) GetInternationalWireDetailsOk() (*InternationalWireDetails, bool)`

GetInternationalWireDetailsOk returns a tuple with the InternationalWireDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalWireDetails

`func (o *AddAccountsRequestRoutingIdentifiers) SetInternationalWireDetails(v InternationalWireDetails)`

SetInternationalWireDetails sets InternationalWireDetails field to given value.

### HasInternationalWireDetails

`func (o *AddAccountsRequestRoutingIdentifiers) HasInternationalWireDetails() bool`

HasInternationalWireDetails returns a boolean if a field has been set.

### GetWireRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) GetWireRoutingNumber() string`

GetWireRoutingNumber returns the WireRoutingNumber field if non-nil, zero value otherwise.

### GetWireRoutingNumberOk

`func (o *AddAccountsRequestRoutingIdentifiers) GetWireRoutingNumberOk() (*string, bool)`

GetWireRoutingNumberOk returns a tuple with the WireRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) SetWireRoutingNumber(v string)`

SetWireRoutingNumber sets WireRoutingNumber field to given value.

### HasWireRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) HasWireRoutingNumber() bool`

HasWireRoutingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


