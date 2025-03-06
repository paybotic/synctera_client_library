# InternationalWireDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAddress** | [**Address**](Address.md) |  | 
**CorrespondentBanksDetails** | Pointer to [**[]CorrespondentBankDetails**](CorrespondentBankDetails.md) | Correspondent banks details used for international payments.  | [optional] 
**SwiftCode** | **string** | The SWIFT code (also known as BIC code) used for international payments.  | 

## Methods

### NewInternationalWireDetails

`func NewInternationalWireDetails(bankAddress Address, swiftCode string, ) *InternationalWireDetails`

NewInternationalWireDetails instantiates a new InternationalWireDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalWireDetailsWithDefaults

`func NewInternationalWireDetailsWithDefaults() *InternationalWireDetails`

NewInternationalWireDetailsWithDefaults instantiates a new InternationalWireDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAddress

`func (o *InternationalWireDetails) GetBankAddress() Address`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *InternationalWireDetails) GetBankAddressOk() (*Address, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *InternationalWireDetails) SetBankAddress(v Address)`

SetBankAddress sets BankAddress field to given value.


### GetCorrespondentBanksDetails

`func (o *InternationalWireDetails) GetCorrespondentBanksDetails() []CorrespondentBankDetails`

GetCorrespondentBanksDetails returns the CorrespondentBanksDetails field if non-nil, zero value otherwise.

### GetCorrespondentBanksDetailsOk

`func (o *InternationalWireDetails) GetCorrespondentBanksDetailsOk() (*[]CorrespondentBankDetails, bool)`

GetCorrespondentBanksDetailsOk returns a tuple with the CorrespondentBanksDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondentBanksDetails

`func (o *InternationalWireDetails) SetCorrespondentBanksDetails(v []CorrespondentBankDetails)`

SetCorrespondentBanksDetails sets CorrespondentBanksDetails field to given value.

### HasCorrespondentBanksDetails

`func (o *InternationalWireDetails) HasCorrespondentBanksDetails() bool`

HasCorrespondentBanksDetails returns a boolean if a field has been set.

### GetSwiftCode

`func (o *InternationalWireDetails) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *InternationalWireDetails) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *InternationalWireDetails) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


