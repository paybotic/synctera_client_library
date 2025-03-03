# OutgoingAch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Receiving account uuid. This links to the external account resource (represented by &#39;receiving_account_id&#39; in the request to send outgoing ACH). | [optional] 
**AccountName** | Pointer to **string** | Receiving account name | [optional] 
**AccountNo** | **string** | Receiving account number | 
**Amount** | **int32** | Transaction amount in cents | 
**BankId** | **int32** |  | 
**CompanyEntryDescription** | Pointer to **string** | Company Entry Description field in ACH batch header. Originator inserts this field&#39;s value to provide the Receiver with a description of the entry&#39;s purpose. | [optional] 
**CompanyName** | Pointer to **string** | Company Name field in ACH batch header. | [optional] 
**EffectiveDate** | **string** | Effective date of the transaction | 
**ExternalId** | **string** | Transaction ID in the ledger | 
**Hold** | Pointer to [**HoldData**](HoldData.md) |  | [optional] 
**Id** | **string** |  | 
**IncomingAchId** | Pointer to **string** | ID of the linked incoming ACH entry. This is filled only for outgoing ACH entries that are returns and links to the originally received incoming entry. | [optional] 
**IsSameDay** | **bool** | Was initiated as same-day ACH transaction | 
**Memo** | **string** |  | 
**PartnerId** | **int32** |  | 
**ReferenceInfo** | Pointer to **string** | Transaction reference info | [optional] 
**ReturnData** | Pointer to [**ReturnData**](ReturnData.md) |  | [optional] 
**SourceAccountId** | Pointer to **string** | Originating account uuid. This links to the customer account on Synctera platform (represented by &#39;originating_account_id&#39; in the request to send outgoing ACH). | [optional] 
**SourceAccountName** | Pointer to **string** | Originating account name | [optional] 
**SourceAccountNo** | **string** | Originating account number | 
**Status** | **string** |  | 
**TraceNo** | **string** | Trace number of the transaction | 

## Methods

### NewOutgoingAch

`func NewOutgoingAch(accountNo string, amount int32, bankId int32, effectiveDate string, externalId string, id string, isSameDay bool, memo string, partnerId int32, sourceAccountNo string, status string, traceNo string, ) *OutgoingAch`

NewOutgoingAch instantiates a new OutgoingAch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingAchWithDefaults

`func NewOutgoingAchWithDefaults() *OutgoingAch`

NewOutgoingAchWithDefaults instantiates a new OutgoingAch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *OutgoingAch) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OutgoingAch) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OutgoingAch) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *OutgoingAch) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *OutgoingAch) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *OutgoingAch) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *OutgoingAch) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *OutgoingAch) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNo

`func (o *OutgoingAch) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *OutgoingAch) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *OutgoingAch) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.


### GetAmount

`func (o *OutgoingAch) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OutgoingAch) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OutgoingAch) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetBankId

`func (o *OutgoingAch) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *OutgoingAch) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *OutgoingAch) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetCompanyEntryDescription

`func (o *OutgoingAch) GetCompanyEntryDescription() string`

GetCompanyEntryDescription returns the CompanyEntryDescription field if non-nil, zero value otherwise.

### GetCompanyEntryDescriptionOk

`func (o *OutgoingAch) GetCompanyEntryDescriptionOk() (*string, bool)`

GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntryDescription

`func (o *OutgoingAch) SetCompanyEntryDescription(v string)`

SetCompanyEntryDescription sets CompanyEntryDescription field to given value.

### HasCompanyEntryDescription

`func (o *OutgoingAch) HasCompanyEntryDescription() bool`

HasCompanyEntryDescription returns a boolean if a field has been set.

### GetCompanyName

`func (o *OutgoingAch) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OutgoingAch) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OutgoingAch) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OutgoingAch) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *OutgoingAch) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *OutgoingAch) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *OutgoingAch) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetExternalId

`func (o *OutgoingAch) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *OutgoingAch) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *OutgoingAch) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetHold

`func (o *OutgoingAch) GetHold() HoldData`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *OutgoingAch) GetHoldOk() (*HoldData, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *OutgoingAch) SetHold(v HoldData)`

SetHold sets Hold field to given value.

### HasHold

`func (o *OutgoingAch) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetId

`func (o *OutgoingAch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingAch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutgoingAch) SetId(v string)`

SetId sets Id field to given value.


### GetIncomingAchId

`func (o *OutgoingAch) GetIncomingAchId() string`

GetIncomingAchId returns the IncomingAchId field if non-nil, zero value otherwise.

### GetIncomingAchIdOk

`func (o *OutgoingAch) GetIncomingAchIdOk() (*string, bool)`

GetIncomingAchIdOk returns a tuple with the IncomingAchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingAchId

`func (o *OutgoingAch) SetIncomingAchId(v string)`

SetIncomingAchId sets IncomingAchId field to given value.

### HasIncomingAchId

`func (o *OutgoingAch) HasIncomingAchId() bool`

HasIncomingAchId returns a boolean if a field has been set.

### GetIsSameDay

`func (o *OutgoingAch) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *OutgoingAch) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *OutgoingAch) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.


### GetMemo

`func (o *OutgoingAch) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *OutgoingAch) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *OutgoingAch) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetPartnerId

`func (o *OutgoingAch) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *OutgoingAch) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *OutgoingAch) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetReferenceInfo

`func (o *OutgoingAch) GetReferenceInfo() string`

GetReferenceInfo returns the ReferenceInfo field if non-nil, zero value otherwise.

### GetReferenceInfoOk

`func (o *OutgoingAch) GetReferenceInfoOk() (*string, bool)`

GetReferenceInfoOk returns a tuple with the ReferenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInfo

`func (o *OutgoingAch) SetReferenceInfo(v string)`

SetReferenceInfo sets ReferenceInfo field to given value.

### HasReferenceInfo

`func (o *OutgoingAch) HasReferenceInfo() bool`

HasReferenceInfo returns a boolean if a field has been set.

### GetReturnData

`func (o *OutgoingAch) GetReturnData() ReturnData`

GetReturnData returns the ReturnData field if non-nil, zero value otherwise.

### GetReturnDataOk

`func (o *OutgoingAch) GetReturnDataOk() (*ReturnData, bool)`

GetReturnDataOk returns a tuple with the ReturnData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnData

`func (o *OutgoingAch) SetReturnData(v ReturnData)`

SetReturnData sets ReturnData field to given value.

### HasReturnData

`func (o *OutgoingAch) HasReturnData() bool`

HasReturnData returns a boolean if a field has been set.

### GetSourceAccountId

`func (o *OutgoingAch) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *OutgoingAch) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *OutgoingAch) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.

### HasSourceAccountId

`func (o *OutgoingAch) HasSourceAccountId() bool`

HasSourceAccountId returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *OutgoingAch) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *OutgoingAch) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *OutgoingAch) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *OutgoingAch) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetSourceAccountNo

`func (o *OutgoingAch) GetSourceAccountNo() string`

GetSourceAccountNo returns the SourceAccountNo field if non-nil, zero value otherwise.

### GetSourceAccountNoOk

`func (o *OutgoingAch) GetSourceAccountNoOk() (*string, bool)`

GetSourceAccountNoOk returns a tuple with the SourceAccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountNo

`func (o *OutgoingAch) SetSourceAccountNo(v string)`

SetSourceAccountNo sets SourceAccountNo field to given value.


### GetStatus

`func (o *OutgoingAch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutgoingAch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutgoingAch) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTraceNo

`func (o *OutgoingAch) GetTraceNo() string`

GetTraceNo returns the TraceNo field if non-nil, zero value otherwise.

### GetTraceNoOk

`func (o *OutgoingAch) GetTraceNoOk() (*string, bool)`

GetTraceNoOk returns a tuple with the TraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceNo

`func (o *OutgoingAch) SetTraceNo(v string)`

SetTraceNo sets TraceNo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


