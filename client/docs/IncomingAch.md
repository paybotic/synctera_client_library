# IncomingAch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The unique identifier for a receiving account. If the entry is in suspense (status &#x3D; \&quot;IN_SUSPENSE\&quot;), this represents the intended customer account, i.e. the account that &#39;account_no&#39; from the ACH entry refers to (if it exists in the Synctera system). If the entry has been posted (status &#x3D; \&quot;POSTED\&quot;), it is the account the transaction was actually posted to. In that case it does not necessarily correspond to the account number in &#39;account_no&#39;.  | [optional] 
**AccountNo** | **string** | Receiving account number. This is the account number from the ACH entry, actual account the transaction gets posted to after processing may be different (e.g. suspense account). | 
**Amount** | **int32** | Transfer amount in cents ($100 would be 10000) | 
**CompanyEntryDescription** | **string** | Company Entry Description field in ACH batch header. | 
**CompanyName** | **string** | Company Name field in ACH batch header. | 
**DcSign** | **string** | The type of transaction (debit or credit). An incoming debit pulls money out of the receiving account, a credit is a transfer in. | 
**DeclineReason** | Pointer to **string** | Free-form text describing the reason why this entry did not post to the intended account when first processed. | [optional] 
**EffectiveDate** | **string** | Effective date of the transaction | 
**ExternalId** | Pointer to **string** | Transaction ID in the ledger. The transaction may not exist yet in case the entry is a future-dated ACH. | [optional] 
**IatInfo** | Pointer to [**IatData**](IatData.md) |  | [optional] 
**Id** | **string** |  | 
**IdentificationNumber** | **string** | Value in this field varies depending on the SEC code. Can contain check serial number, identification number or a name of the originator. | 
**IsFutureDated** | **bool** | Was the effective date in the future when the entry was received? | 
**NotificationOfChange** | Pointer to [**NocData**](NocData.md) |  | [optional] 
**OriginatingRoutingNumber** | **string** | The routing number of the DFI that originated the entry, with check digit included (9 digits in total). | 
**OutgoingAchId** | Pointer to **string** | ID of the linked outgoing ACH entry. This is filled only for incoming ACH entries that are returns and links to the original outgoing entry that is now being returned. | [optional] 
**ReferenceInfo** | Pointer to **[]string** | Contents of all attached records with addenda 05 (payment related information). Some SEC codes allow multiple instances of addenda 05. | [optional] 
**ReturnData** | Pointer to [**ReturnData**](ReturnData.md) |  | [optional] 
**SecCode** | **string** | SEC (Standard Entry Class) code of the ACH entry | 
**SettlementDate** | **string** | Settlement date of the transaction | 
**Status** | **string** | Processing status of the incoming entry | 
**Tenant** | **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | 
**TraceNo** | **string** | Trace number of the ACH entry | 

## Methods

### NewIncomingAch

`func NewIncomingAch(accountNo string, amount int32, companyEntryDescription string, companyName string, dcSign string, effectiveDate string, id string, identificationNumber string, isFutureDated bool, originatingRoutingNumber string, secCode string, settlementDate string, status string, tenant string, traceNo string, ) *IncomingAch`

NewIncomingAch instantiates a new IncomingAch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomingAchWithDefaults

`func NewIncomingAchWithDefaults() *IncomingAch`

NewIncomingAchWithDefaults instantiates a new IncomingAch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *IncomingAch) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *IncomingAch) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *IncomingAch) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *IncomingAch) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountNo

`func (o *IncomingAch) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *IncomingAch) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *IncomingAch) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.


### GetAmount

`func (o *IncomingAch) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IncomingAch) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IncomingAch) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCompanyEntryDescription

`func (o *IncomingAch) GetCompanyEntryDescription() string`

GetCompanyEntryDescription returns the CompanyEntryDescription field if non-nil, zero value otherwise.

### GetCompanyEntryDescriptionOk

`func (o *IncomingAch) GetCompanyEntryDescriptionOk() (*string, bool)`

GetCompanyEntryDescriptionOk returns a tuple with the CompanyEntryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyEntryDescription

`func (o *IncomingAch) SetCompanyEntryDescription(v string)`

SetCompanyEntryDescription sets CompanyEntryDescription field to given value.


### GetCompanyName

`func (o *IncomingAch) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *IncomingAch) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *IncomingAch) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetDcSign

`func (o *IncomingAch) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *IncomingAch) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *IncomingAch) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetDeclineReason

`func (o *IncomingAch) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *IncomingAch) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *IncomingAch) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *IncomingAch) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *IncomingAch) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *IncomingAch) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *IncomingAch) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetExternalId

`func (o *IncomingAch) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IncomingAch) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IncomingAch) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *IncomingAch) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIatInfo

`func (o *IncomingAch) GetIatInfo() IatData`

GetIatInfo returns the IatInfo field if non-nil, zero value otherwise.

### GetIatInfoOk

`func (o *IncomingAch) GetIatInfoOk() (*IatData, bool)`

GetIatInfoOk returns a tuple with the IatInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIatInfo

`func (o *IncomingAch) SetIatInfo(v IatData)`

SetIatInfo sets IatInfo field to given value.

### HasIatInfo

`func (o *IncomingAch) HasIatInfo() bool`

HasIatInfo returns a boolean if a field has been set.

### GetId

`func (o *IncomingAch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncomingAch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncomingAch) SetId(v string)`

SetId sets Id field to given value.


### GetIdentificationNumber

`func (o *IncomingAch) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *IncomingAch) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *IncomingAch) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.


### GetIsFutureDated

`func (o *IncomingAch) GetIsFutureDated() bool`

GetIsFutureDated returns the IsFutureDated field if non-nil, zero value otherwise.

### GetIsFutureDatedOk

`func (o *IncomingAch) GetIsFutureDatedOk() (*bool, bool)`

GetIsFutureDatedOk returns a tuple with the IsFutureDated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFutureDated

`func (o *IncomingAch) SetIsFutureDated(v bool)`

SetIsFutureDated sets IsFutureDated field to given value.


### GetNotificationOfChange

`func (o *IncomingAch) GetNotificationOfChange() NocData`

GetNotificationOfChange returns the NotificationOfChange field if non-nil, zero value otherwise.

### GetNotificationOfChangeOk

`func (o *IncomingAch) GetNotificationOfChangeOk() (*NocData, bool)`

GetNotificationOfChangeOk returns a tuple with the NotificationOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationOfChange

`func (o *IncomingAch) SetNotificationOfChange(v NocData)`

SetNotificationOfChange sets NotificationOfChange field to given value.

### HasNotificationOfChange

`func (o *IncomingAch) HasNotificationOfChange() bool`

HasNotificationOfChange returns a boolean if a field has been set.

### GetOriginatingRoutingNumber

`func (o *IncomingAch) GetOriginatingRoutingNumber() string`

GetOriginatingRoutingNumber returns the OriginatingRoutingNumber field if non-nil, zero value otherwise.

### GetOriginatingRoutingNumberOk

`func (o *IncomingAch) GetOriginatingRoutingNumberOk() (*string, bool)`

GetOriginatingRoutingNumberOk returns a tuple with the OriginatingRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingRoutingNumber

`func (o *IncomingAch) SetOriginatingRoutingNumber(v string)`

SetOriginatingRoutingNumber sets OriginatingRoutingNumber field to given value.


### GetOutgoingAchId

`func (o *IncomingAch) GetOutgoingAchId() string`

GetOutgoingAchId returns the OutgoingAchId field if non-nil, zero value otherwise.

### GetOutgoingAchIdOk

`func (o *IncomingAch) GetOutgoingAchIdOk() (*string, bool)`

GetOutgoingAchIdOk returns a tuple with the OutgoingAchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingAchId

`func (o *IncomingAch) SetOutgoingAchId(v string)`

SetOutgoingAchId sets OutgoingAchId field to given value.

### HasOutgoingAchId

`func (o *IncomingAch) HasOutgoingAchId() bool`

HasOutgoingAchId returns a boolean if a field has been set.

### GetReferenceInfo

`func (o *IncomingAch) GetReferenceInfo() []string`

GetReferenceInfo returns the ReferenceInfo field if non-nil, zero value otherwise.

### GetReferenceInfoOk

`func (o *IncomingAch) GetReferenceInfoOk() (*[]string, bool)`

GetReferenceInfoOk returns a tuple with the ReferenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInfo

`func (o *IncomingAch) SetReferenceInfo(v []string)`

SetReferenceInfo sets ReferenceInfo field to given value.

### HasReferenceInfo

`func (o *IncomingAch) HasReferenceInfo() bool`

HasReferenceInfo returns a boolean if a field has been set.

### GetReturnData

`func (o *IncomingAch) GetReturnData() ReturnData`

GetReturnData returns the ReturnData field if non-nil, zero value otherwise.

### GetReturnDataOk

`func (o *IncomingAch) GetReturnDataOk() (*ReturnData, bool)`

GetReturnDataOk returns a tuple with the ReturnData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnData

`func (o *IncomingAch) SetReturnData(v ReturnData)`

SetReturnData sets ReturnData field to given value.

### HasReturnData

`func (o *IncomingAch) HasReturnData() bool`

HasReturnData returns a boolean if a field has been set.

### GetSecCode

`func (o *IncomingAch) GetSecCode() string`

GetSecCode returns the SecCode field if non-nil, zero value otherwise.

### GetSecCodeOk

`func (o *IncomingAch) GetSecCodeOk() (*string, bool)`

GetSecCodeOk returns a tuple with the SecCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecCode

`func (o *IncomingAch) SetSecCode(v string)`

SetSecCode sets SecCode field to given value.


### GetSettlementDate

`func (o *IncomingAch) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *IncomingAch) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *IncomingAch) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.


### GetStatus

`func (o *IncomingAch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncomingAch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncomingAch) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *IncomingAch) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IncomingAch) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IncomingAch) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTraceNo

`func (o *IncomingAch) GetTraceNo() string`

GetTraceNo returns the TraceNo field if non-nil, zero value otherwise.

### GetTraceNoOk

`func (o *IncomingAch) GetTraceNoOk() (*string, bool)`

GetTraceNoOk returns a tuple with the TraceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceNo

`func (o *IncomingAch) SetTraceNo(v string)`

SetTraceNo sets TraceNo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


