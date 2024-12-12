# AccountRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAccountOutOfArea** | Pointer to **bool** | A flag to indicate whether any of the account holders of the account are out of the supported countries. Can only be set by the Synctera platform.  | [optional] 
**IsDelinquent** | Pointer to **bool** | A flag to indicate whether a credit account is greater than 30 days past due on minimum payments. Delinquent accounts will be unable to spend until overdue amount is paid. Can only be set by the Synctera platform.  | [optional] 
**IsPastDue** | Pointer to **bool** | A flag to indicate whether a credit account is past due on minimum payments. Can only be set by the Synctera platform.  | [optional] 
**IsRevoked** | Pointer to **bool** | A flag to indicates whether a credit account has been revoked (greater than 90 days past due). Revoked accounts will be unable to spend and no longer accrue interest. Can only be set by the Synctera platform.  | [optional] 

## Methods

### NewAccountRestrictions

`func NewAccountRestrictions() *AccountRestrictions`

NewAccountRestrictions instantiates a new AccountRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRestrictionsWithDefaults

`func NewAccountRestrictionsWithDefaults() *AccountRestrictions`

NewAccountRestrictionsWithDefaults instantiates a new AccountRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAccountOutOfArea

`func (o *AccountRestrictions) GetIsAccountOutOfArea() bool`

GetIsAccountOutOfArea returns the IsAccountOutOfArea field if non-nil, zero value otherwise.

### GetIsAccountOutOfAreaOk

`func (o *AccountRestrictions) GetIsAccountOutOfAreaOk() (*bool, bool)`

GetIsAccountOutOfAreaOk returns a tuple with the IsAccountOutOfArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountOutOfArea

`func (o *AccountRestrictions) SetIsAccountOutOfArea(v bool)`

SetIsAccountOutOfArea sets IsAccountOutOfArea field to given value.

### HasIsAccountOutOfArea

`func (o *AccountRestrictions) HasIsAccountOutOfArea() bool`

HasIsAccountOutOfArea returns a boolean if a field has been set.

### GetIsDelinquent

`func (o *AccountRestrictions) GetIsDelinquent() bool`

GetIsDelinquent returns the IsDelinquent field if non-nil, zero value otherwise.

### GetIsDelinquentOk

`func (o *AccountRestrictions) GetIsDelinquentOk() (*bool, bool)`

GetIsDelinquentOk returns a tuple with the IsDelinquent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelinquent

`func (o *AccountRestrictions) SetIsDelinquent(v bool)`

SetIsDelinquent sets IsDelinquent field to given value.

### HasIsDelinquent

`func (o *AccountRestrictions) HasIsDelinquent() bool`

HasIsDelinquent returns a boolean if a field has been set.

### GetIsPastDue

`func (o *AccountRestrictions) GetIsPastDue() bool`

GetIsPastDue returns the IsPastDue field if non-nil, zero value otherwise.

### GetIsPastDueOk

`func (o *AccountRestrictions) GetIsPastDueOk() (*bool, bool)`

GetIsPastDueOk returns a tuple with the IsPastDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPastDue

`func (o *AccountRestrictions) SetIsPastDue(v bool)`

SetIsPastDue sets IsPastDue field to given value.

### HasIsPastDue

`func (o *AccountRestrictions) HasIsPastDue() bool`

HasIsPastDue returns a boolean if a field has been set.

### GetIsRevoked

`func (o *AccountRestrictions) GetIsRevoked() bool`

GetIsRevoked returns the IsRevoked field if non-nil, zero value otherwise.

### GetIsRevokedOk

`func (o *AccountRestrictions) GetIsRevokedOk() (*bool, bool)`

GetIsRevokedOk returns a tuple with the IsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRevoked

`func (o *AccountRestrictions) SetIsRevoked(v bool)`

SetIsRevoked sets IsRevoked field to given value.

### HasIsRevoked

`func (o *AccountRestrictions) HasIsRevoked() bool`

HasIsRevoked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


