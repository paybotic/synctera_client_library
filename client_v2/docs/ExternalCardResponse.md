# ExternalCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | Pointer to **string** | Bank Identification Number | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Currency** | **string** | ISO 4217  Alpha-3 currency code | 
**CustomerId** | **string** | Customer ID for the application | 
**DeletionTime** | Pointer to **time.Time** |  | [optional] 
**ExpirationMonth** | **string** | Card expiration month | 
**ExpirationYear** | **string** | Card expiration year | 
**Id** | **string** | Unique identifier | 
**Issuer** | Pointer to **string** | Name of the issuing financial institution | [optional] 
**LastFour** | **string** | The last four digits of the card PAN | 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** | The cardholder name | 
**PaymentAccountReference** | Pointer to **string** | A unique identifier associated with a specific cardholder PAN and its affiliated tokens | [optional] 
**Status** | [**ExternalCardStatus**](ExternalCardStatus.md) |  | 
**Tenant** | **string** | The id of the tenant containing the resource.  | 
**Verifications** | Pointer to [**ExternalCardVerifications**](ExternalCardVerifications.md) |  | [optional] 

## Methods

### NewExternalCardResponse

`func NewExternalCardResponse(currency string, customerId string, expirationMonth string, expirationYear string, id string, lastFour string, name string, status ExternalCardStatus, tenant string, ) *ExternalCardResponse`

NewExternalCardResponse instantiates a new ExternalCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardResponseWithDefaults

`func NewExternalCardResponseWithDefaults() *ExternalCardResponse`

NewExternalCardResponseWithDefaults instantiates a new ExternalCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *ExternalCardResponse) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *ExternalCardResponse) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *ExternalCardResponse) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *ExternalCardResponse) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCreationTime

`func (o *ExternalCardResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ExternalCardResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ExternalCardResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ExternalCardResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *ExternalCardResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExternalCardResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExternalCardResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *ExternalCardResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExternalCardResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExternalCardResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDeletionTime

`func (o *ExternalCardResponse) GetDeletionTime() time.Time`

GetDeletionTime returns the DeletionTime field if non-nil, zero value otherwise.

### GetDeletionTimeOk

`func (o *ExternalCardResponse) GetDeletionTimeOk() (*time.Time, bool)`

GetDeletionTimeOk returns a tuple with the DeletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTime

`func (o *ExternalCardResponse) SetDeletionTime(v time.Time)`

SetDeletionTime sets DeletionTime field to given value.

### HasDeletionTime

`func (o *ExternalCardResponse) HasDeletionTime() bool`

HasDeletionTime returns a boolean if a field has been set.

### GetExpirationMonth

`func (o *ExternalCardResponse) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *ExternalCardResponse) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *ExternalCardResponse) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.


### GetExpirationYear

`func (o *ExternalCardResponse) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *ExternalCardResponse) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *ExternalCardResponse) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.


### GetId

`func (o *ExternalCardResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalCardResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalCardResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIssuer

`func (o *ExternalCardResponse) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ExternalCardResponse) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ExternalCardResponse) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ExternalCardResponse) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLastFour

`func (o *ExternalCardResponse) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *ExternalCardResponse) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *ExternalCardResponse) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.


### GetLastUpdatedTime

`func (o *ExternalCardResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ExternalCardResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ExternalCardResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *ExternalCardResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetName

`func (o *ExternalCardResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalCardResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalCardResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPaymentAccountReference

`func (o *ExternalCardResponse) GetPaymentAccountReference() string`

GetPaymentAccountReference returns the PaymentAccountReference field if non-nil, zero value otherwise.

### GetPaymentAccountReferenceOk

`func (o *ExternalCardResponse) GetPaymentAccountReferenceOk() (*string, bool)`

GetPaymentAccountReferenceOk returns a tuple with the PaymentAccountReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountReference

`func (o *ExternalCardResponse) SetPaymentAccountReference(v string)`

SetPaymentAccountReference sets PaymentAccountReference field to given value.

### HasPaymentAccountReference

`func (o *ExternalCardResponse) HasPaymentAccountReference() bool`

HasPaymentAccountReference returns a boolean if a field has been set.

### GetStatus

`func (o *ExternalCardResponse) GetStatus() ExternalCardStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalCardResponse) GetStatusOk() (*ExternalCardStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalCardResponse) SetStatus(v ExternalCardStatus)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *ExternalCardResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ExternalCardResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ExternalCardResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetVerifications

`func (o *ExternalCardResponse) GetVerifications() ExternalCardVerifications`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *ExternalCardResponse) GetVerificationsOk() (*ExternalCardVerifications, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *ExternalCardResponse) SetVerifications(v ExternalCardVerifications)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *ExternalCardResponse) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


