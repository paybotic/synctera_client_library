# Deposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The ID of the account | [optional] 
**BackImageId** | Pointer to **string** | ID of the uploaded image of the back of the check | [optional] 
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**CheckAmount** | Pointer to **int32** | Amount on check in ISO 4217 minor currency units | [optional] 
**DepositCurrency** | Pointer to **string** | ISO 4217 currency code for the deposit amount | [optional] 
**FrontImageId** | Pointer to **string** | ID of the uploaded image of the front of the check | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**DateCaptured** | Pointer to **time.Time** | Date the deposit was captured, in RFC 3339 format | [optional] 
**DateProcessed** | Pointer to **time.Time** | Date the deposit was processed, in RFC 3339 format | [optional] 
**DepositAmount** | Pointer to **int32** | Amount deposited in ISO 4217 minor currency units | [optional] 
**Id** | Pointer to **string** | Remote Check Deposit ID | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] 
**OcrAccountNumber** | Pointer to **string** | Account number of the issuer of the check, included if OCR is successful | [optional] 
**OcrCheckNumber** | Pointer to **string** | The unique check number for this check in the checkbook, included if OCR is successful and there is a check number on the check | [optional] 
**OcrRoutingNumber** | Pointer to **string** | Routing number of the issuing bank, included if OCR is successful | [optional] 
**Status** | Pointer to **string** | The status of the deposit | [optional] 
**TransactionId** | Pointer to **string** | The ID of the transaction associated with this deposit | [optional] 
**VendorInfo** | Pointer to [**VendorInfo**](VendorInfo.md) |  | [optional] 

## Methods

### NewDeposit

`func NewDeposit() *Deposit`

NewDeposit instantiates a new Deposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositWithDefaults

`func NewDepositWithDefaults() *Deposit`

NewDepositWithDefaults instantiates a new Deposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Deposit) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Deposit) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Deposit) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Deposit) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBackImageId

`func (o *Deposit) GetBackImageId() string`

GetBackImageId returns the BackImageId field if non-nil, zero value otherwise.

### GetBackImageIdOk

`func (o *Deposit) GetBackImageIdOk() (*string, bool)`

GetBackImageIdOk returns a tuple with the BackImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackImageId

`func (o *Deposit) SetBackImageId(v string)`

SetBackImageId sets BackImageId field to given value.

### HasBackImageId

`func (o *Deposit) HasBackImageId() bool`

HasBackImageId returns a boolean if a field has been set.

### GetBusinessId

`func (o *Deposit) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *Deposit) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *Deposit) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *Deposit) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCheckAmount

`func (o *Deposit) GetCheckAmount() int32`

GetCheckAmount returns the CheckAmount field if non-nil, zero value otherwise.

### GetCheckAmountOk

`func (o *Deposit) GetCheckAmountOk() (*int32, bool)`

GetCheckAmountOk returns a tuple with the CheckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAmount

`func (o *Deposit) SetCheckAmount(v int32)`

SetCheckAmount sets CheckAmount field to given value.

### HasCheckAmount

`func (o *Deposit) HasCheckAmount() bool`

HasCheckAmount returns a boolean if a field has been set.

### GetDepositCurrency

`func (o *Deposit) GetDepositCurrency() string`

GetDepositCurrency returns the DepositCurrency field if non-nil, zero value otherwise.

### GetDepositCurrencyOk

`func (o *Deposit) GetDepositCurrencyOk() (*string, bool)`

GetDepositCurrencyOk returns a tuple with the DepositCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositCurrency

`func (o *Deposit) SetDepositCurrency(v string)`

SetDepositCurrency sets DepositCurrency field to given value.

### HasDepositCurrency

`func (o *Deposit) HasDepositCurrency() bool`

HasDepositCurrency returns a boolean if a field has been set.

### GetFrontImageId

`func (o *Deposit) GetFrontImageId() string`

GetFrontImageId returns the FrontImageId field if non-nil, zero value otherwise.

### GetFrontImageIdOk

`func (o *Deposit) GetFrontImageIdOk() (*string, bool)`

GetFrontImageIdOk returns a tuple with the FrontImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImageId

`func (o *Deposit) SetFrontImageId(v string)`

SetFrontImageId sets FrontImageId field to given value.

### HasFrontImageId

`func (o *Deposit) HasFrontImageId() bool`

HasFrontImageId returns a boolean if a field has been set.

### GetMetadata

`func (o *Deposit) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Deposit) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Deposit) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Deposit) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPersonId

`func (o *Deposit) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Deposit) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Deposit) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Deposit) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetCreationTime

`func (o *Deposit) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Deposit) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Deposit) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Deposit) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDateCaptured

`func (o *Deposit) GetDateCaptured() time.Time`

GetDateCaptured returns the DateCaptured field if non-nil, zero value otherwise.

### GetDateCapturedOk

`func (o *Deposit) GetDateCapturedOk() (*time.Time, bool)`

GetDateCapturedOk returns a tuple with the DateCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCaptured

`func (o *Deposit) SetDateCaptured(v time.Time)`

SetDateCaptured sets DateCaptured field to given value.

### HasDateCaptured

`func (o *Deposit) HasDateCaptured() bool`

HasDateCaptured returns a boolean if a field has been set.

### GetDateProcessed

`func (o *Deposit) GetDateProcessed() time.Time`

GetDateProcessed returns the DateProcessed field if non-nil, zero value otherwise.

### GetDateProcessedOk

`func (o *Deposit) GetDateProcessedOk() (*time.Time, bool)`

GetDateProcessedOk returns a tuple with the DateProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProcessed

`func (o *Deposit) SetDateProcessed(v time.Time)`

SetDateProcessed sets DateProcessed field to given value.

### HasDateProcessed

`func (o *Deposit) HasDateProcessed() bool`

HasDateProcessed returns a boolean if a field has been set.

### GetDepositAmount

`func (o *Deposit) GetDepositAmount() int32`

GetDepositAmount returns the DepositAmount field if non-nil, zero value otherwise.

### GetDepositAmountOk

`func (o *Deposit) GetDepositAmountOk() (*int32, bool)`

GetDepositAmountOk returns a tuple with the DepositAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAmount

`func (o *Deposit) SetDepositAmount(v int32)`

SetDepositAmount sets DepositAmount field to given value.

### HasDepositAmount

`func (o *Deposit) HasDepositAmount() bool`

HasDepositAmount returns a boolean if a field has been set.

### GetId

`func (o *Deposit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deposit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deposit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Deposit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Deposit) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Deposit) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Deposit) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Deposit) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetOcrAccountNumber

`func (o *Deposit) GetOcrAccountNumber() string`

GetOcrAccountNumber returns the OcrAccountNumber field if non-nil, zero value otherwise.

### GetOcrAccountNumberOk

`func (o *Deposit) GetOcrAccountNumberOk() (*string, bool)`

GetOcrAccountNumberOk returns a tuple with the OcrAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcrAccountNumber

`func (o *Deposit) SetOcrAccountNumber(v string)`

SetOcrAccountNumber sets OcrAccountNumber field to given value.

### HasOcrAccountNumber

`func (o *Deposit) HasOcrAccountNumber() bool`

HasOcrAccountNumber returns a boolean if a field has been set.

### GetOcrCheckNumber

`func (o *Deposit) GetOcrCheckNumber() string`

GetOcrCheckNumber returns the OcrCheckNumber field if non-nil, zero value otherwise.

### GetOcrCheckNumberOk

`func (o *Deposit) GetOcrCheckNumberOk() (*string, bool)`

GetOcrCheckNumberOk returns a tuple with the OcrCheckNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcrCheckNumber

`func (o *Deposit) SetOcrCheckNumber(v string)`

SetOcrCheckNumber sets OcrCheckNumber field to given value.

### HasOcrCheckNumber

`func (o *Deposit) HasOcrCheckNumber() bool`

HasOcrCheckNumber returns a boolean if a field has been set.

### GetOcrRoutingNumber

`func (o *Deposit) GetOcrRoutingNumber() string`

GetOcrRoutingNumber returns the OcrRoutingNumber field if non-nil, zero value otherwise.

### GetOcrRoutingNumberOk

`func (o *Deposit) GetOcrRoutingNumberOk() (*string, bool)`

GetOcrRoutingNumberOk returns a tuple with the OcrRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcrRoutingNumber

`func (o *Deposit) SetOcrRoutingNumber(v string)`

SetOcrRoutingNumber sets OcrRoutingNumber field to given value.

### HasOcrRoutingNumber

`func (o *Deposit) HasOcrRoutingNumber() bool`

HasOcrRoutingNumber returns a boolean if a field has been set.

### GetStatus

`func (o *Deposit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deposit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deposit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Deposit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionId

`func (o *Deposit) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Deposit) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Deposit) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Deposit) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetVendorInfo

`func (o *Deposit) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *Deposit) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *Deposit) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *Deposit) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


