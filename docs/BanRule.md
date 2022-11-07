# BanRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | What action this rule performs | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Email** | Pointer to **string** | This rule will match requests for customers with this email address. May contain a \&quot;*\&quot; as a wildcard.  | [optional] 
**Id** | Pointer to **string** | ban rule ID | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** | This rule will match requests for customers with this phone number. Use E.164 format, with a leading \&quot;+\&quot; and a country code.  | [optional] 
**Reason** | Pointer to **string** | The reason why the rule was added | [optional] 
**ResourceType** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] 
**Source** | Pointer to **string** | The source of the rule | [optional] 
**SsnHash** | Pointer to **string** | This rule will match requests for customers with an SSN that matches this value when hashed.  | [optional] 
**Status** | Pointer to [**BanRuleStatus**](BanRuleStatus.md) |  | [optional] 
**Tenant** | Pointer to **string** | The tenant the rule applies to, with a format of \&quot;b_p\&quot; where \&quot;b\&quot; is the bank_id and \&quot;p\&quot; is the fintech partner_id. Bank requesters can specify a the fintech partner id as \&quot;*\&quot; to indicate the rule applies to all their fintechs. Synctera requesters can specify \&quot;*_*\&quot; to indicate the rule applies globally. The requester must have access to the specified tenant. If not supplied, the tenant defaults to the requester&#39;s tenant.  | [optional] 

## Methods

### NewBanRule

`func NewBanRule() *BanRule`

NewBanRule instantiates a new BanRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanRuleWithDefaults

`func NewBanRuleWithDefaults() *BanRule`

NewBanRuleWithDefaults instantiates a new BanRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BanRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BanRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BanRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BanRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreationTime

`func (o *BanRule) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BanRule) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BanRule) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BanRule) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEmail

`func (o *BanRule) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BanRule) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BanRule) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BanRule) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *BanRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BanRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BanRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BanRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *BanRule) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BanRule) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BanRule) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *BanRule) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *BanRule) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BanRule) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BanRule) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BanRule) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetReason

`func (o *BanRule) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BanRule) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BanRule) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BanRule) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResourceType

`func (o *BanRule) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BanRule) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BanRule) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BanRule) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSource

`func (o *BanRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BanRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BanRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BanRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSsnHash

`func (o *BanRule) GetSsnHash() string`

GetSsnHash returns the SsnHash field if non-nil, zero value otherwise.

### GetSsnHashOk

`func (o *BanRule) GetSsnHashOk() (*string, bool)`

GetSsnHashOk returns a tuple with the SsnHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnHash

`func (o *BanRule) SetSsnHash(v string)`

SetSsnHash sets SsnHash field to given value.

### HasSsnHash

`func (o *BanRule) HasSsnHash() bool`

HasSsnHash returns a boolean if a field has been set.

### GetStatus

`func (o *BanRule) GetStatus() BanRuleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BanRule) GetStatusOk() (*BanRuleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BanRule) SetStatus(v BanRuleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BanRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *BanRule) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BanRule) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BanRule) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BanRule) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


