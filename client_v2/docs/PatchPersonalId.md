# PatchPersonalId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | The updated ISO 3166 Alpha-2 country code for the country that issued the personal identifier. The country code cannot be modified for personal identifier types that have an implicit country, e.g. SSN.  | [optional] 
**ExpiryDate** | Pointer to **string** | The date the associated document is set to expire on set by the governing authority. | [optional] 
**IdType** | [**PersonalIdType**](PersonalIdType.md) |  | 
**Identifier** | Pointer to **string** | The updated personal identifier | [optional] 
**IssueDate** | Pointer to **string** | The date the associated document was issued by the governing authority. | [optional] 
**Tenant** | Pointer to **string** | The id of the tenant containing the resource.  | [optional] 

## Methods

### NewPatchPersonalId

`func NewPatchPersonalId(idType PersonalIdType, ) *PatchPersonalId`

NewPatchPersonalId instantiates a new PatchPersonalId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPersonalIdWithDefaults

`func NewPatchPersonalIdWithDefaults() *PatchPersonalId`

NewPatchPersonalIdWithDefaults instantiates a new PatchPersonalId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *PatchPersonalId) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PatchPersonalId) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PatchPersonalId) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PatchPersonalId) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetExpiryDate

`func (o *PatchPersonalId) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *PatchPersonalId) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *PatchPersonalId) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *PatchPersonalId) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetIdType

`func (o *PatchPersonalId) GetIdType() PersonalIdType`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *PatchPersonalId) GetIdTypeOk() (*PersonalIdType, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *PatchPersonalId) SetIdType(v PersonalIdType)`

SetIdType sets IdType field to given value.


### GetIdentifier

`func (o *PatchPersonalId) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PatchPersonalId) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PatchPersonalId) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PatchPersonalId) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIssueDate

`func (o *PatchPersonalId) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *PatchPersonalId) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *PatchPersonalId) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *PatchPersonalId) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetTenant

`func (o *PatchPersonalId) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchPersonalId) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchPersonalId) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchPersonalId) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


