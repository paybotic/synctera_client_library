# Business

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

## Methods

### NewBusiness

`func NewBusiness() *Business`

NewBusiness instantiates a new Business object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessWithDefaults

`func NewBusinessWithDefaults() *Business`

NewBusinessWithDefaults instantiates a new Business object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *Business) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Business) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Business) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Business) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEin

`func (o *Business) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *Business) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *Business) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *Business) HasEin() bool`

HasEin returns a boolean if a field has been set.

### GetEmail

`func (o *Business) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Business) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Business) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Business) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEntityName

`func (o *Business) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *Business) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *Business) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *Business) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetFormationDate

`func (o *Business) GetFormationDate() string`

GetFormationDate returns the FormationDate field if non-nil, zero value otherwise.

### GetFormationDateOk

`func (o *Business) GetFormationDateOk() (*string, bool)`

GetFormationDateOk returns a tuple with the FormationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormationDate

`func (o *Business) SetFormationDate(v string)`

SetFormationDate sets FormationDate field to given value.

### HasFormationDate

`func (o *Business) HasFormationDate() bool`

HasFormationDate returns a boolean if a field has been set.

### GetFormationState

`func (o *Business) GetFormationState() string`

GetFormationState returns the FormationState field if non-nil, zero value otherwise.

### GetFormationStateOk

`func (o *Business) GetFormationStateOk() (*string, bool)`

GetFormationStateOk returns a tuple with the FormationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormationState

`func (o *Business) SetFormationState(v string)`

SetFormationState sets FormationState field to given value.

### HasFormationState

`func (o *Business) HasFormationState() bool`

HasFormationState returns a boolean if a field has been set.

### GetId

`func (o *Business) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Business) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Business) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Business) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Business) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Business) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Business) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Business) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Business) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Business) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Business) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Business) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *Business) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Business) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Business) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Business) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStructure

`func (o *Business) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *Business) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *Business) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *Business) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetTradeNames

`func (o *Business) GetTradeNames() []string`

GetTradeNames returns the TradeNames field if non-nil, zero value otherwise.

### GetTradeNamesOk

`func (o *Business) GetTradeNamesOk() (*[]string, bool)`

GetTradeNamesOk returns a tuple with the TradeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNames

`func (o *Business) SetTradeNames(v []string)`

SetTradeNames sets TradeNames field to given value.

### HasTradeNames

`func (o *Business) HasTradeNames() bool`

HasTradeNames returns a boolean if a field has been set.

### GetVerificationLastRun

`func (o *Business) GetVerificationLastRun() time.Time`

GetVerificationLastRun returns the VerificationLastRun field if non-nil, zero value otherwise.

### GetVerificationLastRunOk

`func (o *Business) GetVerificationLastRunOk() (*time.Time, bool)`

GetVerificationLastRunOk returns a tuple with the VerificationLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationLastRun

`func (o *Business) SetVerificationLastRun(v time.Time)`

SetVerificationLastRun sets VerificationLastRun field to given value.

### HasVerificationLastRun

`func (o *Business) HasVerificationLastRun() bool`

HasVerificationLastRun returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *Business) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *Business) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *Business) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *Business) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


