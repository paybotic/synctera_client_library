# PatchInternationalWireDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**CorrespondentBanksDetails** | Pointer to [**[]CorrespondentBankDetails**](CorrespondentBankDetails.md) | Correspondent banks details used for international payments. Note that in a patch request,  the entirity of the correspondent_banks_details array will be updated.  | [optional] 
**SwiftCode** | Pointer to **string** | The SWIFT code (also known as BIC code) used for international payments.  | [optional] 

## Methods

### NewPatchInternationalWireDetails

`func NewPatchInternationalWireDetails() *PatchInternationalWireDetails`

NewPatchInternationalWireDetails instantiates a new PatchInternationalWireDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchInternationalWireDetailsWithDefaults

`func NewPatchInternationalWireDetailsWithDefaults() *PatchInternationalWireDetails`

NewPatchInternationalWireDetailsWithDefaults instantiates a new PatchInternationalWireDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAddress

`func (o *PatchInternationalWireDetails) GetBankAddress() Address`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *PatchInternationalWireDetails) GetBankAddressOk() (*Address, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *PatchInternationalWireDetails) SetBankAddress(v Address)`

SetBankAddress sets BankAddress field to given value.

### HasBankAddress

`func (o *PatchInternationalWireDetails) HasBankAddress() bool`

HasBankAddress returns a boolean if a field has been set.

### GetCorrespondentBanksDetails

`func (o *PatchInternationalWireDetails) GetCorrespondentBanksDetails() []CorrespondentBankDetails`

GetCorrespondentBanksDetails returns the CorrespondentBanksDetails field if non-nil, zero value otherwise.

### GetCorrespondentBanksDetailsOk

`func (o *PatchInternationalWireDetails) GetCorrespondentBanksDetailsOk() (*[]CorrespondentBankDetails, bool)`

GetCorrespondentBanksDetailsOk returns a tuple with the CorrespondentBanksDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondentBanksDetails

`func (o *PatchInternationalWireDetails) SetCorrespondentBanksDetails(v []CorrespondentBankDetails)`

SetCorrespondentBanksDetails sets CorrespondentBanksDetails field to given value.

### HasCorrespondentBanksDetails

`func (o *PatchInternationalWireDetails) HasCorrespondentBanksDetails() bool`

HasCorrespondentBanksDetails returns a boolean if a field has been set.

### GetSwiftCode

`func (o *PatchInternationalWireDetails) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *PatchInternationalWireDetails) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *PatchInternationalWireDetails) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *PatchInternationalWireDetails) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


