# CorrespondentBankDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAddress** | [**Address**](Address.md) |  | 
**BankName** | **string** | The name of the correspondent bank.  | 
**SwiftCode** | **string** | The SWIFT code (also known as BIC code) used for international payments.  | 

## Methods

### NewCorrespondentBankDetails

`func NewCorrespondentBankDetails(bankAddress Address, bankName string, swiftCode string, ) *CorrespondentBankDetails`

NewCorrespondentBankDetails instantiates a new CorrespondentBankDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrespondentBankDetailsWithDefaults

`func NewCorrespondentBankDetailsWithDefaults() *CorrespondentBankDetails`

NewCorrespondentBankDetailsWithDefaults instantiates a new CorrespondentBankDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAddress

`func (o *CorrespondentBankDetails) GetBankAddress() Address`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *CorrespondentBankDetails) GetBankAddressOk() (*Address, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *CorrespondentBankDetails) SetBankAddress(v Address)`

SetBankAddress sets BankAddress field to given value.


### GetBankName

`func (o *CorrespondentBankDetails) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *CorrespondentBankDetails) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *CorrespondentBankDetails) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetSwiftCode

`func (o *CorrespondentBankDetails) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *CorrespondentBankDetails) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *CorrespondentBankDetails) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


