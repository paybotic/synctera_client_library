# ResponseLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**CreationTime** | **time.Time** | The date and time the license resource  was created. | 
**CustomerId** | Pointer to **string** |  | [optional] 
**Id** | **string** | License record unique id | 
**LastUpdatedTime** | **time.Time** | The date and time the license resource was last updated. | 
**LastVerifiedTime** | Pointer to **time.Time** | Timestamp of the last time the license was verified | [optional] 
**LicenseClassification** | Pointer to **string** | The classification of the license | [optional] 
**LicenseExpirationDate** | Pointer to **string** | The date on which the license will expire | [optional] 
**LicenseIssuanceDate** | Pointer to **string** | The date on which the license was issued | [optional] 
**LicenseNumber** | **string** | The entity&#39;s license number | 
**LicenseType** | [**LicenseType**](LicenseType.md) |  | 
**LicenseTypeDescription** | Pointer to **string** | Free-form text describing the type of the license | [optional] 
**LicenseeAddress** | Pointer to **string** | The address of the entity that holds the license as reported by the verifying vendor | [optional] 
**LicenseeName** | Pointer to **string** | The name of the entity that holds the license that&#39;s reported by the verifying vendor | [optional] 
**LicensingAuthority** | Pointer to **string** | The name of the licensing body that granted the license | [optional] 
**Status** | **string** | The status of the license | 
**Tenant** | **string** | The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.  | 

## Methods

### NewResponseLicense

`func NewResponseLicense(creationTime time.Time, id string, lastUpdatedTime time.Time, licenseNumber string, licenseType LicenseType, status string, tenant string, ) *ResponseLicense`

NewResponseLicense instantiates a new ResponseLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseLicenseWithDefaults

`func NewResponseLicenseWithDefaults() *ResponseLicense`

NewResponseLicenseWithDefaults instantiates a new ResponseLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *ResponseLicense) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *ResponseLicense) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *ResponseLicense) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *ResponseLicense) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *ResponseLicense) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ResponseLicense) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ResponseLicense) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomerId

`func (o *ResponseLicense) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ResponseLicense) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ResponseLicense) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ResponseLicense) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *ResponseLicense) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseLicense) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseLicense) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *ResponseLicense) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ResponseLicense) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ResponseLicense) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetLastVerifiedTime

`func (o *ResponseLicense) GetLastVerifiedTime() time.Time`

GetLastVerifiedTime returns the LastVerifiedTime field if non-nil, zero value otherwise.

### GetLastVerifiedTimeOk

`func (o *ResponseLicense) GetLastVerifiedTimeOk() (*time.Time, bool)`

GetLastVerifiedTimeOk returns a tuple with the LastVerifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerifiedTime

`func (o *ResponseLicense) SetLastVerifiedTime(v time.Time)`

SetLastVerifiedTime sets LastVerifiedTime field to given value.

### HasLastVerifiedTime

`func (o *ResponseLicense) HasLastVerifiedTime() bool`

HasLastVerifiedTime returns a boolean if a field has been set.

### GetLicenseClassification

`func (o *ResponseLicense) GetLicenseClassification() string`

GetLicenseClassification returns the LicenseClassification field if non-nil, zero value otherwise.

### GetLicenseClassificationOk

`func (o *ResponseLicense) GetLicenseClassificationOk() (*string, bool)`

GetLicenseClassificationOk returns a tuple with the LicenseClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseClassification

`func (o *ResponseLicense) SetLicenseClassification(v string)`

SetLicenseClassification sets LicenseClassification field to given value.

### HasLicenseClassification

`func (o *ResponseLicense) HasLicenseClassification() bool`

HasLicenseClassification returns a boolean if a field has been set.

### GetLicenseExpirationDate

`func (o *ResponseLicense) GetLicenseExpirationDate() string`

GetLicenseExpirationDate returns the LicenseExpirationDate field if non-nil, zero value otherwise.

### GetLicenseExpirationDateOk

`func (o *ResponseLicense) GetLicenseExpirationDateOk() (*string, bool)`

GetLicenseExpirationDateOk returns a tuple with the LicenseExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpirationDate

`func (o *ResponseLicense) SetLicenseExpirationDate(v string)`

SetLicenseExpirationDate sets LicenseExpirationDate field to given value.

### HasLicenseExpirationDate

`func (o *ResponseLicense) HasLicenseExpirationDate() bool`

HasLicenseExpirationDate returns a boolean if a field has been set.

### GetLicenseIssuanceDate

`func (o *ResponseLicense) GetLicenseIssuanceDate() string`

GetLicenseIssuanceDate returns the LicenseIssuanceDate field if non-nil, zero value otherwise.

### GetLicenseIssuanceDateOk

`func (o *ResponseLicense) GetLicenseIssuanceDateOk() (*string, bool)`

GetLicenseIssuanceDateOk returns a tuple with the LicenseIssuanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseIssuanceDate

`func (o *ResponseLicense) SetLicenseIssuanceDate(v string)`

SetLicenseIssuanceDate sets LicenseIssuanceDate field to given value.

### HasLicenseIssuanceDate

`func (o *ResponseLicense) HasLicenseIssuanceDate() bool`

HasLicenseIssuanceDate returns a boolean if a field has been set.

### GetLicenseNumber

`func (o *ResponseLicense) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *ResponseLicense) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *ResponseLicense) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.


### GetLicenseType

`func (o *ResponseLicense) GetLicenseType() LicenseType`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ResponseLicense) GetLicenseTypeOk() (*LicenseType, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ResponseLicense) SetLicenseType(v LicenseType)`

SetLicenseType sets LicenseType field to given value.


### GetLicenseTypeDescription

`func (o *ResponseLicense) GetLicenseTypeDescription() string`

GetLicenseTypeDescription returns the LicenseTypeDescription field if non-nil, zero value otherwise.

### GetLicenseTypeDescriptionOk

`func (o *ResponseLicense) GetLicenseTypeDescriptionOk() (*string, bool)`

GetLicenseTypeDescriptionOk returns a tuple with the LicenseTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTypeDescription

`func (o *ResponseLicense) SetLicenseTypeDescription(v string)`

SetLicenseTypeDescription sets LicenseTypeDescription field to given value.

### HasLicenseTypeDescription

`func (o *ResponseLicense) HasLicenseTypeDescription() bool`

HasLicenseTypeDescription returns a boolean if a field has been set.

### GetLicenseeAddress

`func (o *ResponseLicense) GetLicenseeAddress() string`

GetLicenseeAddress returns the LicenseeAddress field if non-nil, zero value otherwise.

### GetLicenseeAddressOk

`func (o *ResponseLicense) GetLicenseeAddressOk() (*string, bool)`

GetLicenseeAddressOk returns a tuple with the LicenseeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseeAddress

`func (o *ResponseLicense) SetLicenseeAddress(v string)`

SetLicenseeAddress sets LicenseeAddress field to given value.

### HasLicenseeAddress

`func (o *ResponseLicense) HasLicenseeAddress() bool`

HasLicenseeAddress returns a boolean if a field has been set.

### GetLicenseeName

`func (o *ResponseLicense) GetLicenseeName() string`

GetLicenseeName returns the LicenseeName field if non-nil, zero value otherwise.

### GetLicenseeNameOk

`func (o *ResponseLicense) GetLicenseeNameOk() (*string, bool)`

GetLicenseeNameOk returns a tuple with the LicenseeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseeName

`func (o *ResponseLicense) SetLicenseeName(v string)`

SetLicenseeName sets LicenseeName field to given value.

### HasLicenseeName

`func (o *ResponseLicense) HasLicenseeName() bool`

HasLicenseeName returns a boolean if a field has been set.

### GetLicensingAuthority

`func (o *ResponseLicense) GetLicensingAuthority() string`

GetLicensingAuthority returns the LicensingAuthority field if non-nil, zero value otherwise.

### GetLicensingAuthorityOk

`func (o *ResponseLicense) GetLicensingAuthorityOk() (*string, bool)`

GetLicensingAuthorityOk returns a tuple with the LicensingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingAuthority

`func (o *ResponseLicense) SetLicensingAuthority(v string)`

SetLicensingAuthority sets LicensingAuthority field to given value.

### HasLicensingAuthority

`func (o *ResponseLicense) HasLicensingAuthority() bool`

HasLicensingAuthority returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseLicense) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseLicense) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseLicense) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *ResponseLicense) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ResponseLicense) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ResponseLicense) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


