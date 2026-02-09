# SyncteraPayConfigurationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The currency that will be used to match the Outgoing Synctera Pay transfer to the appropriate configuration.  | 
**Domestic** | Pointer to **bool** | Whether the transfer is domestic or international.  | [optional] 
**ExternalPaymentRail** | **string** | The external payment rail that will be used for AML checks.  | 
**SettlementAccountId** | **string** | The ID of the settlement account that will be used to match the Outgoing Synctera Pay transfer to the appropriate configuration.  | 
**SettlementAccountType** | **string** | The type of the settlement account that will be used to match the Outgoing Synctera Pay transfer to the appropriate configuration.  | 
**Subtypes** | [**[]SyncteraPaySubtype**](SyncteraPaySubtype.md) | The subtypes that will be used to match the Outgoing Synctera Pay transfer to the appropriate configuration.  | 

## Methods

### NewSyncteraPayConfigurationDetails

`func NewSyncteraPayConfigurationDetails(currency string, externalPaymentRail string, settlementAccountId string, settlementAccountType string, subtypes []SyncteraPaySubtype, ) *SyncteraPayConfigurationDetails`

NewSyncteraPayConfigurationDetails instantiates a new SyncteraPayConfigurationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncteraPayConfigurationDetailsWithDefaults

`func NewSyncteraPayConfigurationDetailsWithDefaults() *SyncteraPayConfigurationDetails`

NewSyncteraPayConfigurationDetailsWithDefaults instantiates a new SyncteraPayConfigurationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *SyncteraPayConfigurationDetails) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SyncteraPayConfigurationDetails) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SyncteraPayConfigurationDetails) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDomestic

`func (o *SyncteraPayConfigurationDetails) GetDomestic() bool`

GetDomestic returns the Domestic field if non-nil, zero value otherwise.

### GetDomesticOk

`func (o *SyncteraPayConfigurationDetails) GetDomesticOk() (*bool, bool)`

GetDomesticOk returns a tuple with the Domestic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomestic

`func (o *SyncteraPayConfigurationDetails) SetDomestic(v bool)`

SetDomestic sets Domestic field to given value.

### HasDomestic

`func (o *SyncteraPayConfigurationDetails) HasDomestic() bool`

HasDomestic returns a boolean if a field has been set.

### GetExternalPaymentRail

`func (o *SyncteraPayConfigurationDetails) GetExternalPaymentRail() string`

GetExternalPaymentRail returns the ExternalPaymentRail field if non-nil, zero value otherwise.

### GetExternalPaymentRailOk

`func (o *SyncteraPayConfigurationDetails) GetExternalPaymentRailOk() (*string, bool)`

GetExternalPaymentRailOk returns a tuple with the ExternalPaymentRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentRail

`func (o *SyncteraPayConfigurationDetails) SetExternalPaymentRail(v string)`

SetExternalPaymentRail sets ExternalPaymentRail field to given value.


### GetSettlementAccountId

`func (o *SyncteraPayConfigurationDetails) GetSettlementAccountId() string`

GetSettlementAccountId returns the SettlementAccountId field if non-nil, zero value otherwise.

### GetSettlementAccountIdOk

`func (o *SyncteraPayConfigurationDetails) GetSettlementAccountIdOk() (*string, bool)`

GetSettlementAccountIdOk returns a tuple with the SettlementAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAccountId

`func (o *SyncteraPayConfigurationDetails) SetSettlementAccountId(v string)`

SetSettlementAccountId sets SettlementAccountId field to given value.


### GetSettlementAccountType

`func (o *SyncteraPayConfigurationDetails) GetSettlementAccountType() string`

GetSettlementAccountType returns the SettlementAccountType field if non-nil, zero value otherwise.

### GetSettlementAccountTypeOk

`func (o *SyncteraPayConfigurationDetails) GetSettlementAccountTypeOk() (*string, bool)`

GetSettlementAccountTypeOk returns a tuple with the SettlementAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAccountType

`func (o *SyncteraPayConfigurationDetails) SetSettlementAccountType(v string)`

SetSettlementAccountType sets SettlementAccountType field to given value.


### GetSubtypes

`func (o *SyncteraPayConfigurationDetails) GetSubtypes() []SyncteraPaySubtype`

GetSubtypes returns the Subtypes field if non-nil, zero value otherwise.

### GetSubtypesOk

`func (o *SyncteraPayConfigurationDetails) GetSubtypesOk() (*[]SyncteraPaySubtype, bool)`

GetSubtypesOk returns a tuple with the Subtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtypes

`func (o *SyncteraPayConfigurationDetails) SetSubtypes(v []SyncteraPaySubtype)`

SetSubtypes sets Subtypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


