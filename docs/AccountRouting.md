# AccountRouting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchRoutingNumber** | Pointer to **string** | The routing number used for US ACH payments. Only appears if &#x60;bank_countries&#x60; contains &#x60;US&#x60;. Value may be masked, in which case only the last four digits are returned.  | [optional] 
**BankAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**BankCountries** | **[]string** | The countries that this bank operates the account in | 
**BankName** | **string** | The name of the bank managing the account | 
**CorrespondentBankDetails** | Pointer to [**[]CorrespondentBankDetails**](CorrespondentBankDetails.md) | The details of the correspondent banks for the account.  | [optional] 
**EftCaRoutingNumber** | Pointer to **string** | &gt; 🚧 Alpha &gt; This is an Alpha property. Feedback from the community is welcome. We may make breaking changes to this property. The 9 digit routing number used for EFT CA payments, identifying a Canadian bank. The format is 0xxxyyyyy where xxx is the institution number and yyyyy is the transit number. On write, Synctera will store the entire routing number; on read, we only return the last 4 characters. Value may be masked, in which case only the last four digits are returned.  | [optional] 
**EftRoutingNumber** | Pointer to **string** | The routing number used for EFT payments, identifying a Canadian bank, consisting of the institution number and the branch number. Only appears if &#x60;bank_countries&#x60; contains &#x60;CA&#x60;. Value may be masked, in which case only the last four digits are returned. This attribute is deprecated and will be removed in a future API version. Use eft_ca_routing_number instead.  | [optional] 
**SwiftCode** | Pointer to **string** | The SWIFT code for the bank. Value may be masked, in which case only the last four characters are returned.  | [optional] 
**WireRoutingNumber** | Pointer to **string** | The routing number used for domestic wire payments. Only appears if &#x60;bank_countries&#x60; contains &#x60;US&#x60;. Value may be masked, in which case only the last four digits are returned.  | [optional] 

## Methods

### NewAccountRouting

`func NewAccountRouting(bankCountries []string, bankName string, ) *AccountRouting`

NewAccountRouting instantiates a new AccountRouting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRoutingWithDefaults

`func NewAccountRoutingWithDefaults() *AccountRouting`

NewAccountRoutingWithDefaults instantiates a new AccountRouting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchRoutingNumber

`func (o *AccountRouting) GetAchRoutingNumber() string`

GetAchRoutingNumber returns the AchRoutingNumber field if non-nil, zero value otherwise.

### GetAchRoutingNumberOk

`func (o *AccountRouting) GetAchRoutingNumberOk() (*string, bool)`

GetAchRoutingNumberOk returns a tuple with the AchRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchRoutingNumber

`func (o *AccountRouting) SetAchRoutingNumber(v string)`

SetAchRoutingNumber sets AchRoutingNumber field to given value.

### HasAchRoutingNumber

`func (o *AccountRouting) HasAchRoutingNumber() bool`

HasAchRoutingNumber returns a boolean if a field has been set.

### GetBankAddress

`func (o *AccountRouting) GetBankAddress() Address`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *AccountRouting) GetBankAddressOk() (*Address, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *AccountRouting) SetBankAddress(v Address)`

SetBankAddress sets BankAddress field to given value.

### HasBankAddress

`func (o *AccountRouting) HasBankAddress() bool`

HasBankAddress returns a boolean if a field has been set.

### GetBankCountries

`func (o *AccountRouting) GetBankCountries() []string`

GetBankCountries returns the BankCountries field if non-nil, zero value otherwise.

### GetBankCountriesOk

`func (o *AccountRouting) GetBankCountriesOk() (*[]string, bool)`

GetBankCountriesOk returns a tuple with the BankCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountries

`func (o *AccountRouting) SetBankCountries(v []string)`

SetBankCountries sets BankCountries field to given value.


### GetBankName

`func (o *AccountRouting) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *AccountRouting) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *AccountRouting) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetCorrespondentBankDetails

`func (o *AccountRouting) GetCorrespondentBankDetails() []CorrespondentBankDetails`

GetCorrespondentBankDetails returns the CorrespondentBankDetails field if non-nil, zero value otherwise.

### GetCorrespondentBankDetailsOk

`func (o *AccountRouting) GetCorrespondentBankDetailsOk() (*[]CorrespondentBankDetails, bool)`

GetCorrespondentBankDetailsOk returns a tuple with the CorrespondentBankDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondentBankDetails

`func (o *AccountRouting) SetCorrespondentBankDetails(v []CorrespondentBankDetails)`

SetCorrespondentBankDetails sets CorrespondentBankDetails field to given value.

### HasCorrespondentBankDetails

`func (o *AccountRouting) HasCorrespondentBankDetails() bool`

HasCorrespondentBankDetails returns a boolean if a field has been set.

### SetCorrespondentBankDetailsNil

`func (o *AccountRouting) SetCorrespondentBankDetailsNil(b bool)`

 SetCorrespondentBankDetailsNil sets the value for CorrespondentBankDetails to be an explicit nil

### UnsetCorrespondentBankDetails
`func (o *AccountRouting) UnsetCorrespondentBankDetails()`

UnsetCorrespondentBankDetails ensures that no value is present for CorrespondentBankDetails, not even an explicit nil
### GetEftCaRoutingNumber

`func (o *AccountRouting) GetEftCaRoutingNumber() string`

GetEftCaRoutingNumber returns the EftCaRoutingNumber field if non-nil, zero value otherwise.

### GetEftCaRoutingNumberOk

`func (o *AccountRouting) GetEftCaRoutingNumberOk() (*string, bool)`

GetEftCaRoutingNumberOk returns a tuple with the EftCaRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEftCaRoutingNumber

`func (o *AccountRouting) SetEftCaRoutingNumber(v string)`

SetEftCaRoutingNumber sets EftCaRoutingNumber field to given value.

### HasEftCaRoutingNumber

`func (o *AccountRouting) HasEftCaRoutingNumber() bool`

HasEftCaRoutingNumber returns a boolean if a field has been set.

### GetEftRoutingNumber

`func (o *AccountRouting) GetEftRoutingNumber() string`

GetEftRoutingNumber returns the EftRoutingNumber field if non-nil, zero value otherwise.

### GetEftRoutingNumberOk

`func (o *AccountRouting) GetEftRoutingNumberOk() (*string, bool)`

GetEftRoutingNumberOk returns a tuple with the EftRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEftRoutingNumber

`func (o *AccountRouting) SetEftRoutingNumber(v string)`

SetEftRoutingNumber sets EftRoutingNumber field to given value.

### HasEftRoutingNumber

`func (o *AccountRouting) HasEftRoutingNumber() bool`

HasEftRoutingNumber returns a boolean if a field has been set.

### GetSwiftCode

`func (o *AccountRouting) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountRouting) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountRouting) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *AccountRouting) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetWireRoutingNumber

`func (o *AccountRouting) GetWireRoutingNumber() string`

GetWireRoutingNumber returns the WireRoutingNumber field if non-nil, zero value otherwise.

### GetWireRoutingNumberOk

`func (o *AccountRouting) GetWireRoutingNumberOk() (*string, bool)`

GetWireRoutingNumberOk returns a tuple with the WireRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireRoutingNumber

`func (o *AccountRouting) SetWireRoutingNumber(v string)`

SetWireRoutingNumber sets WireRoutingNumber field to given value.

### HasWireRoutingNumber

`func (o *AccountRouting) HasWireRoutingNumber() bool`

HasWireRoutingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


