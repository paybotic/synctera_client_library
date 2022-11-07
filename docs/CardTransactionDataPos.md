# CardTransactionDataPos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardDataInputCapability** | Pointer to **string** | Terminal Card data acceptance method | [optional] 
**CardHolderPresence** | Pointer to **bool** | Cardholder presence | [optional] 
**CardPresence** | Pointer to **bool** | Card presence | [optional] 
**CardholderAuthenticationMethod** | Pointer to **string** | Cardholder authentication method | [optional] 
**CountryCode** | Pointer to **string** | Terminal country code | [optional] 
**IsInstallment** | Pointer to **bool** | Transaction is an installment payment | [optional] 
**IsRecurring** | Pointer to **bool** | Transaction is recurring | [optional] 
**PanEntryMode** | Pointer to **string** | Card pan capture method | [optional] 
**PartialApprovalCapable** | Pointer to **bool** | Terminal partial approval capability | [optional] 
**PinEntryMode** | Pointer to **string** | Card pin capture method | [optional] 
**PinPresent** | Pointer to **bool** | Pin presence | [optional] 
**PurchaseAmountOnly** | Pointer to **bool** | Terminal purchase amount only | [optional] 
**TerminalAttendance** | Pointer to **string** | Terminal attendance | [optional] 
**TerminalId** | Pointer to **string** | Terminal identifier | [optional] 
**TerminalLocation** | Pointer to **string** | Terminal location | [optional] 
**TerminalType** | Pointer to **string** | Terminal type | [optional] 
**Zip** | Pointer to **string** | Terminal zip code | [optional] 

## Methods

### NewCardTransactionDataPos

`func NewCardTransactionDataPos() *CardTransactionDataPos`

NewCardTransactionDataPos instantiates a new CardTransactionDataPos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTransactionDataPosWithDefaults

`func NewCardTransactionDataPosWithDefaults() *CardTransactionDataPos`

NewCardTransactionDataPosWithDefaults instantiates a new CardTransactionDataPos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardDataInputCapability

`func (o *CardTransactionDataPos) GetCardDataInputCapability() string`

GetCardDataInputCapability returns the CardDataInputCapability field if non-nil, zero value otherwise.

### GetCardDataInputCapabilityOk

`func (o *CardTransactionDataPos) GetCardDataInputCapabilityOk() (*string, bool)`

GetCardDataInputCapabilityOk returns a tuple with the CardDataInputCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDataInputCapability

`func (o *CardTransactionDataPos) SetCardDataInputCapability(v string)`

SetCardDataInputCapability sets CardDataInputCapability field to given value.

### HasCardDataInputCapability

`func (o *CardTransactionDataPos) HasCardDataInputCapability() bool`

HasCardDataInputCapability returns a boolean if a field has been set.

### GetCardHolderPresence

`func (o *CardTransactionDataPos) GetCardHolderPresence() bool`

GetCardHolderPresence returns the CardHolderPresence field if non-nil, zero value otherwise.

### GetCardHolderPresenceOk

`func (o *CardTransactionDataPos) GetCardHolderPresenceOk() (*bool, bool)`

GetCardHolderPresenceOk returns a tuple with the CardHolderPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderPresence

`func (o *CardTransactionDataPos) SetCardHolderPresence(v bool)`

SetCardHolderPresence sets CardHolderPresence field to given value.

### HasCardHolderPresence

`func (o *CardTransactionDataPos) HasCardHolderPresence() bool`

HasCardHolderPresence returns a boolean if a field has been set.

### GetCardPresence

`func (o *CardTransactionDataPos) GetCardPresence() bool`

GetCardPresence returns the CardPresence field if non-nil, zero value otherwise.

### GetCardPresenceOk

`func (o *CardTransactionDataPos) GetCardPresenceOk() (*bool, bool)`

GetCardPresenceOk returns a tuple with the CardPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPresence

`func (o *CardTransactionDataPos) SetCardPresence(v bool)`

SetCardPresence sets CardPresence field to given value.

### HasCardPresence

`func (o *CardTransactionDataPos) HasCardPresence() bool`

HasCardPresence returns a boolean if a field has been set.

### GetCardholderAuthenticationMethod

`func (o *CardTransactionDataPos) GetCardholderAuthenticationMethod() string`

GetCardholderAuthenticationMethod returns the CardholderAuthenticationMethod field if non-nil, zero value otherwise.

### GetCardholderAuthenticationMethodOk

`func (o *CardTransactionDataPos) GetCardholderAuthenticationMethodOk() (*string, bool)`

GetCardholderAuthenticationMethodOk returns a tuple with the CardholderAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderAuthenticationMethod

`func (o *CardTransactionDataPos) SetCardholderAuthenticationMethod(v string)`

SetCardholderAuthenticationMethod sets CardholderAuthenticationMethod field to given value.

### HasCardholderAuthenticationMethod

`func (o *CardTransactionDataPos) HasCardholderAuthenticationMethod() bool`

HasCardholderAuthenticationMethod returns a boolean if a field has been set.

### GetCountryCode

`func (o *CardTransactionDataPos) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CardTransactionDataPos) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CardTransactionDataPos) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CardTransactionDataPos) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetIsInstallment

`func (o *CardTransactionDataPos) GetIsInstallment() bool`

GetIsInstallment returns the IsInstallment field if non-nil, zero value otherwise.

### GetIsInstallmentOk

`func (o *CardTransactionDataPos) GetIsInstallmentOk() (*bool, bool)`

GetIsInstallmentOk returns a tuple with the IsInstallment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstallment

`func (o *CardTransactionDataPos) SetIsInstallment(v bool)`

SetIsInstallment sets IsInstallment field to given value.

### HasIsInstallment

`func (o *CardTransactionDataPos) HasIsInstallment() bool`

HasIsInstallment returns a boolean if a field has been set.

### GetIsRecurring

`func (o *CardTransactionDataPos) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *CardTransactionDataPos) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *CardTransactionDataPos) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.

### HasIsRecurring

`func (o *CardTransactionDataPos) HasIsRecurring() bool`

HasIsRecurring returns a boolean if a field has been set.

### GetPanEntryMode

`func (o *CardTransactionDataPos) GetPanEntryMode() string`

GetPanEntryMode returns the PanEntryMode field if non-nil, zero value otherwise.

### GetPanEntryModeOk

`func (o *CardTransactionDataPos) GetPanEntryModeOk() (*string, bool)`

GetPanEntryModeOk returns a tuple with the PanEntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanEntryMode

`func (o *CardTransactionDataPos) SetPanEntryMode(v string)`

SetPanEntryMode sets PanEntryMode field to given value.

### HasPanEntryMode

`func (o *CardTransactionDataPos) HasPanEntryMode() bool`

HasPanEntryMode returns a boolean if a field has been set.

### GetPartialApprovalCapable

`func (o *CardTransactionDataPos) GetPartialApprovalCapable() bool`

GetPartialApprovalCapable returns the PartialApprovalCapable field if non-nil, zero value otherwise.

### GetPartialApprovalCapableOk

`func (o *CardTransactionDataPos) GetPartialApprovalCapableOk() (*bool, bool)`

GetPartialApprovalCapableOk returns a tuple with the PartialApprovalCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialApprovalCapable

`func (o *CardTransactionDataPos) SetPartialApprovalCapable(v bool)`

SetPartialApprovalCapable sets PartialApprovalCapable field to given value.

### HasPartialApprovalCapable

`func (o *CardTransactionDataPos) HasPartialApprovalCapable() bool`

HasPartialApprovalCapable returns a boolean if a field has been set.

### GetPinEntryMode

`func (o *CardTransactionDataPos) GetPinEntryMode() string`

GetPinEntryMode returns the PinEntryMode field if non-nil, zero value otherwise.

### GetPinEntryModeOk

`func (o *CardTransactionDataPos) GetPinEntryModeOk() (*string, bool)`

GetPinEntryModeOk returns a tuple with the PinEntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinEntryMode

`func (o *CardTransactionDataPos) SetPinEntryMode(v string)`

SetPinEntryMode sets PinEntryMode field to given value.

### HasPinEntryMode

`func (o *CardTransactionDataPos) HasPinEntryMode() bool`

HasPinEntryMode returns a boolean if a field has been set.

### GetPinPresent

`func (o *CardTransactionDataPos) GetPinPresent() bool`

GetPinPresent returns the PinPresent field if non-nil, zero value otherwise.

### GetPinPresentOk

`func (o *CardTransactionDataPos) GetPinPresentOk() (*bool, bool)`

GetPinPresentOk returns a tuple with the PinPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinPresent

`func (o *CardTransactionDataPos) SetPinPresent(v bool)`

SetPinPresent sets PinPresent field to given value.

### HasPinPresent

`func (o *CardTransactionDataPos) HasPinPresent() bool`

HasPinPresent returns a boolean if a field has been set.

### GetPurchaseAmountOnly

`func (o *CardTransactionDataPos) GetPurchaseAmountOnly() bool`

GetPurchaseAmountOnly returns the PurchaseAmountOnly field if non-nil, zero value otherwise.

### GetPurchaseAmountOnlyOk

`func (o *CardTransactionDataPos) GetPurchaseAmountOnlyOk() (*bool, bool)`

GetPurchaseAmountOnlyOk returns a tuple with the PurchaseAmountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseAmountOnly

`func (o *CardTransactionDataPos) SetPurchaseAmountOnly(v bool)`

SetPurchaseAmountOnly sets PurchaseAmountOnly field to given value.

### HasPurchaseAmountOnly

`func (o *CardTransactionDataPos) HasPurchaseAmountOnly() bool`

HasPurchaseAmountOnly returns a boolean if a field has been set.

### GetTerminalAttendance

`func (o *CardTransactionDataPos) GetTerminalAttendance() string`

GetTerminalAttendance returns the TerminalAttendance field if non-nil, zero value otherwise.

### GetTerminalAttendanceOk

`func (o *CardTransactionDataPos) GetTerminalAttendanceOk() (*string, bool)`

GetTerminalAttendanceOk returns a tuple with the TerminalAttendance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalAttendance

`func (o *CardTransactionDataPos) SetTerminalAttendance(v string)`

SetTerminalAttendance sets TerminalAttendance field to given value.

### HasTerminalAttendance

`func (o *CardTransactionDataPos) HasTerminalAttendance() bool`

HasTerminalAttendance returns a boolean if a field has been set.

### GetTerminalId

`func (o *CardTransactionDataPos) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *CardTransactionDataPos) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *CardTransactionDataPos) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *CardTransactionDataPos) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### GetTerminalLocation

`func (o *CardTransactionDataPos) GetTerminalLocation() string`

GetTerminalLocation returns the TerminalLocation field if non-nil, zero value otherwise.

### GetTerminalLocationOk

`func (o *CardTransactionDataPos) GetTerminalLocationOk() (*string, bool)`

GetTerminalLocationOk returns a tuple with the TerminalLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalLocation

`func (o *CardTransactionDataPos) SetTerminalLocation(v string)`

SetTerminalLocation sets TerminalLocation field to given value.

### HasTerminalLocation

`func (o *CardTransactionDataPos) HasTerminalLocation() bool`

HasTerminalLocation returns a boolean if a field has been set.

### GetTerminalType

`func (o *CardTransactionDataPos) GetTerminalType() string`

GetTerminalType returns the TerminalType field if non-nil, zero value otherwise.

### GetTerminalTypeOk

`func (o *CardTransactionDataPos) GetTerminalTypeOk() (*string, bool)`

GetTerminalTypeOk returns a tuple with the TerminalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalType

`func (o *CardTransactionDataPos) SetTerminalType(v string)`

SetTerminalType sets TerminalType field to given value.

### HasTerminalType

`func (o *CardTransactionDataPos) HasTerminalType() bool`

HasTerminalType returns a boolean if a field has been set.

### GetZip

`func (o *CardTransactionDataPos) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CardTransactionDataPos) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CardTransactionDataPos) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *CardTransactionDataPos) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


