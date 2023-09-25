# PatchAccountsRequestRoutingIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchRoutingNumber** | Pointer to **string** | The routing number used for US ACH payments. On write, Synctera will store the entire routing number; on read, we only return the last 4 characters.  | [optional] 
**BankName** | Pointer to **string** | The name of the bank managing the account | [optional] 
**EftCaRoutingNumber** | Pointer to **string** | &gt; 🚧 Alpha &gt; This is an Alpha property. Feedback from the community is welcome. We may make breaking changes to this property. The 9 digit routing number used for EFT CA payments, identifying a Canadian bank.  The format is 0xxxyyyyy where xxx is the institution number and yyyyy is the transit number. On write, Synctera will store the entire routing number; on read, we only return the last 4 characters.  | [optional] 
**InternationalWireDetails** | Pointer to [**PatchInternationalWireDetails**](PatchInternationalWireDetails.md) |  | [optional] 
**WireRoutingNumber** | Pointer to **string** | The routing number used for domestic wire payments. On write, Synctera will store the entire routing number; on read, we only return the last 4 characters.  | [optional] 

## Methods

### NewPatchAccountsRequestRoutingIdentifiers

`func NewPatchAccountsRequestRoutingIdentifiers() *PatchAccountsRequestRoutingIdentifiers`

NewPatchAccountsRequestRoutingIdentifiers instantiates a new PatchAccountsRequestRoutingIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchAccountsRequestRoutingIdentifiersWithDefaults

`func NewPatchAccountsRequestRoutingIdentifiersWithDefaults() *PatchAccountsRequestRoutingIdentifiers`

NewPatchAccountsRequestRoutingIdentifiersWithDefaults instantiates a new PatchAccountsRequestRoutingIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) GetAchRoutingNumber() string`

GetAchRoutingNumber returns the AchRoutingNumber field if non-nil, zero value otherwise.

### GetAchRoutingNumberOk

`func (o *PatchAccountsRequestRoutingIdentifiers) GetAchRoutingNumberOk() (*string, bool)`

GetAchRoutingNumberOk returns a tuple with the AchRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) SetAchRoutingNumber(v string)`

SetAchRoutingNumber sets AchRoutingNumber field to given value.

### HasAchRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) HasAchRoutingNumber() bool`

HasAchRoutingNumber returns a boolean if a field has been set.

### GetBankName

`func (o *PatchAccountsRequestRoutingIdentifiers) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PatchAccountsRequestRoutingIdentifiers) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PatchAccountsRequestRoutingIdentifiers) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PatchAccountsRequestRoutingIdentifiers) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetEftCaRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) GetEftCaRoutingNumber() string`

GetEftCaRoutingNumber returns the EftCaRoutingNumber field if non-nil, zero value otherwise.

### GetEftCaRoutingNumberOk

`func (o *PatchAccountsRequestRoutingIdentifiers) GetEftCaRoutingNumberOk() (*string, bool)`

GetEftCaRoutingNumberOk returns a tuple with the EftCaRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEftCaRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) SetEftCaRoutingNumber(v string)`

SetEftCaRoutingNumber sets EftCaRoutingNumber field to given value.

### HasEftCaRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) HasEftCaRoutingNumber() bool`

HasEftCaRoutingNumber returns a boolean if a field has been set.

### GetInternationalWireDetails

`func (o *PatchAccountsRequestRoutingIdentifiers) GetInternationalWireDetails() PatchInternationalWireDetails`

GetInternationalWireDetails returns the InternationalWireDetails field if non-nil, zero value otherwise.

### GetInternationalWireDetailsOk

`func (o *PatchAccountsRequestRoutingIdentifiers) GetInternationalWireDetailsOk() (*PatchInternationalWireDetails, bool)`

GetInternationalWireDetailsOk returns a tuple with the InternationalWireDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalWireDetails

`func (o *PatchAccountsRequestRoutingIdentifiers) SetInternationalWireDetails(v PatchInternationalWireDetails)`

SetInternationalWireDetails sets InternationalWireDetails field to given value.

### HasInternationalWireDetails

`func (o *PatchAccountsRequestRoutingIdentifiers) HasInternationalWireDetails() bool`

HasInternationalWireDetails returns a boolean if a field has been set.

### GetWireRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) GetWireRoutingNumber() string`

GetWireRoutingNumber returns the WireRoutingNumber field if non-nil, zero value otherwise.

### GetWireRoutingNumberOk

`func (o *PatchAccountsRequestRoutingIdentifiers) GetWireRoutingNumberOk() (*string, bool)`

GetWireRoutingNumberOk returns a tuple with the WireRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) SetWireRoutingNumber(v string)`

SetWireRoutingNumber sets WireRoutingNumber field to given value.

### HasWireRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) HasWireRoutingNumber() bool`

HasWireRoutingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


