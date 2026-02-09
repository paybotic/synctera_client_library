# BusinessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Ein** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EntityName** | Pointer to **string** |  | [optional] 
**FormationDate** | Pointer to **string** |  | [optional] 
**FormationState** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Structure** | Pointer to **string** |  | [optional] 
**TradeNames** | Pointer to **[]string** |  | [optional] 
**VerificationLastRun** | Pointer to **time.Time** |  | [optional] 
**VerificationStatus** | Pointer to **string** |  | [optional] 
**VendorInfo** | Pointer to [**PartyVendorInfo**](PartyVendorInfo.md) |  | [optional] 

## Methods

### NewBusinessResponse

`func NewBusinessResponse() *BusinessResponse`

NewBusinessResponse instantiates a new BusinessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessResponseWithDefaults

`func NewBusinessResponseWithDefaults() *BusinessResponse`

NewBusinessResponseWithDefaults instantiates a new BusinessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *BusinessResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BusinessResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BusinessResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BusinessResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEin

`func (o *BusinessResponse) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *BusinessResponse) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *BusinessResponse) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *BusinessResponse) HasEin() bool`

HasEin returns a boolean if a field has been set.

### GetEmail

`func (o *BusinessResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BusinessResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BusinessResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BusinessResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEntityName

`func (o *BusinessResponse) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *BusinessResponse) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *BusinessResponse) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *BusinessResponse) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetFormationDate

`func (o *BusinessResponse) GetFormationDate() string`

GetFormationDate returns the FormationDate field if non-nil, zero value otherwise.

### GetFormationDateOk

`func (o *BusinessResponse) GetFormationDateOk() (*string, bool)`

GetFormationDateOk returns a tuple with the FormationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormationDate

`func (o *BusinessResponse) SetFormationDate(v string)`

SetFormationDate sets FormationDate field to given value.

### HasFormationDate

`func (o *BusinessResponse) HasFormationDate() bool`

HasFormationDate returns a boolean if a field has been set.

### GetFormationState

`func (o *BusinessResponse) GetFormationState() string`

GetFormationState returns the FormationState field if non-nil, zero value otherwise.

### GetFormationStateOk

`func (o *BusinessResponse) GetFormationStateOk() (*string, bool)`

GetFormationStateOk returns a tuple with the FormationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormationState

`func (o *BusinessResponse) SetFormationState(v string)`

SetFormationState sets FormationState field to given value.

### HasFormationState

`func (o *BusinessResponse) HasFormationState() bool`

HasFormationState returns a boolean if a field has been set.

### GetId

`func (o *BusinessResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BusinessResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *BusinessResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BusinessResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BusinessResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *BusinessResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *BusinessResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BusinessResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BusinessResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BusinessResponse) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *BusinessResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BusinessResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BusinessResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BusinessResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStructure

`func (o *BusinessResponse) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *BusinessResponse) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *BusinessResponse) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *BusinessResponse) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetTradeNames

`func (o *BusinessResponse) GetTradeNames() []string`

GetTradeNames returns the TradeNames field if non-nil, zero value otherwise.

### GetTradeNamesOk

`func (o *BusinessResponse) GetTradeNamesOk() (*[]string, bool)`

GetTradeNamesOk returns a tuple with the TradeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNames

`func (o *BusinessResponse) SetTradeNames(v []string)`

SetTradeNames sets TradeNames field to given value.

### HasTradeNames

`func (o *BusinessResponse) HasTradeNames() bool`

HasTradeNames returns a boolean if a field has been set.

### GetVerificationLastRun

`func (o *BusinessResponse) GetVerificationLastRun() time.Time`

GetVerificationLastRun returns the VerificationLastRun field if non-nil, zero value otherwise.

### GetVerificationLastRunOk

`func (o *BusinessResponse) GetVerificationLastRunOk() (*time.Time, bool)`

GetVerificationLastRunOk returns a tuple with the VerificationLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLastRun

`func (o *BusinessResponse) SetVerificationLastRun(v time.Time)`

SetVerificationLastRun sets VerificationLastRun field to given value.

### HasVerificationLastRun

`func (o *BusinessResponse) HasVerificationLastRun() bool`

HasVerificationLastRun returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *BusinessResponse) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *BusinessResponse) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *BusinessResponse) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *BusinessResponse) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetVendorInfo

`func (o *BusinessResponse) GetVendorInfo() PartyVendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *BusinessResponse) GetVendorInfoOk() (*PartyVendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *BusinessResponse) SetVendorInfo(v PartyVendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *BusinessResponse) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


