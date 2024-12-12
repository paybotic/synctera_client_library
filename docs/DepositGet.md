# DepositGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**DateCaptured** | **time.Time** | Date the deposit was captured, in RFC 3339 format | 
**DateProcessed** | **time.Time** | Date the deposit was processed, in RFC 3339 format | 
**DepositAmount** | **int32** | Amount deposited in ISO 4217 minor currency units | 
**Id** | **string** | Remote Check Deposit ID | 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] 
**OcrAccountNumber** | Pointer to **string** | Account number of the issuer of the check, included if OCR is successful | [optional] 
**OcrCheckNumber** | Pointer to **string** | The unique check number for this check in the checkbook, included if OCR is successful and there is a check number on the check | [optional] 
**OcrRoutingNumber** | Pointer to **string** | Routing number of the issuing bank, included if OCR is successful | [optional] 
**Status** | **string** | The status of the deposit | 
**TransactionId** | Pointer to **string** | The ID of the transaction associated with this deposit | [optional] 
**VendorInfo** | Pointer to [**VendorInfo**](VendorInfo.md) |  | [optional] 
**AccountId** | **string** | The ID of the account | 
**BackImageId** | **string** | ID of the uploaded image of the back of the check | 
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**CheckAmount** | **int32** | Amount on check in ISO 4217 minor currency units | 
**DepositCurrency** | **string** | ISO 4217 currency code for the deposit amount | 
**FrontImageId** | **string** | ID of the uploaded image of the front of the check | 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 

## Methods

### NewDepositGet

`func NewDepositGet(dateCaptured time.Time, dateProcessed time.Time, depositAmount int32, id string, status string, accountId string, backImageId string, checkAmount int32, depositCurrency string, frontImageId string, ) *DepositGet`

NewDepositGet instantiates a new DepositGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositGetWithDefaults

`func NewDepositGetWithDefaults() *DepositGet`

NewDepositGetWithDefaults instantiates a new DepositGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *DepositGet) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *DepositGet) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *DepositGet) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *DepositGet) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDateCaptured

`func (o *DepositGet) GetDateCaptured() time.Time`

GetDateCaptured returns the DateCaptured field if non-nil, zero value otherwise.

### GetDateCapturedOk

`func (o *DepositGet) GetDateCapturedOk() (*time.Time, bool)`

GetDateCapturedOk returns a tuple with the DateCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCaptured

`func (o *DepositGet) SetDateCaptured(v time.Time)`

SetDateCaptured sets DateCaptured field to given value.


### GetDateProcessed

`func (o *DepositGet) GetDateProcessed() time.Time`

GetDateProcessed returns the DateProcessed field if non-nil, zero value otherwise.

### GetDateProcessedOk

`func (o *DepositGet) GetDateProcessedOk() (*time.Time, bool)`

GetDateProcessedOk returns a tuple with the DateProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateProcessed

`func (o *DepositGet) SetDateProcessed(v time.Time)`

SetDateProcessed sets DateProcessed field to given value.


### GetDepositAmount

`func (o *DepositGet) GetDepositAmount() int32`

GetDepositAmount returns the DepositAmount field if non-nil, zero value otherwise.

### GetDepositAmountOk

`func (o *DepositGet) GetDepositAmountOk() (*int32, bool)`

GetDepositAmountOk returns a tuple with the DepositAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAmount

`func (o *DepositGet) SetDepositAmount(v int32)`

SetDepositAmount sets DepositAmount field to given value.


### GetId

`func (o *DepositGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepositGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepositGet) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *DepositGet) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *DepositGet) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *DepositGet) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *DepositGet) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetOcrAccountNumber

`func (o *DepositGet) GetOcrAccountNumber() string`

GetOcrAccountNumber returns the OcrAccountNumber field if non-nil, zero value otherwise.

### GetOcrAccountNumberOk

`func (o *DepositGet) GetOcrAccountNumberOk() (*string, bool)`

GetOcrAccountNumberOk returns a tuple with the OcrAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcrAccountNumber

`func (o *DepositGet) SetOcrAccountNumber(v string)`

SetOcrAccountNumber sets OcrAccountNumber field to given value.

### HasOcrAccountNumber

`func (o *DepositGet) HasOcrAccountNumber() bool`

HasOcrAccountNumber returns a boolean if a field has been set.

### GetOcrCheckNumber

`func (o *DepositGet) GetOcrCheckNumber() string`

GetOcrCheckNumber returns the OcrCheckNumber field if non-nil, zero value otherwise.

### GetOcrCheckNumberOk

`func (o *DepositGet) GetOcrCheckNumberOk() (*string, bool)`

GetOcrCheckNumberOk returns a tuple with the OcrCheckNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcrCheckNumber

`func (o *DepositGet) SetOcrCheckNumber(v string)`

SetOcrCheckNumber sets OcrCheckNumber field to given value.

### HasOcrCheckNumber

`func (o *DepositGet) HasOcrCheckNumber() bool`

HasOcrCheckNumber returns a boolean if a field has been set.

### GetOcrRoutingNumber

`func (o *DepositGet) GetOcrRoutingNumber() string`

GetOcrRoutingNumber returns the OcrRoutingNumber field if non-nil, zero value otherwise.

### GetOcrRoutingNumberOk

`func (o *DepositGet) GetOcrRoutingNumberOk() (*string, bool)`

GetOcrRoutingNumberOk returns a tuple with the OcrRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcrRoutingNumber

`func (o *DepositGet) SetOcrRoutingNumber(v string)`

SetOcrRoutingNumber sets OcrRoutingNumber field to given value.

### HasOcrRoutingNumber

`func (o *DepositGet) HasOcrRoutingNumber() bool`

HasOcrRoutingNumber returns a boolean if a field has been set.

### GetStatus

`func (o *DepositGet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DepositGet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DepositGet) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *DepositGet) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *DepositGet) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *DepositGet) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *DepositGet) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetVendorInfo

`func (o *DepositGet) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *DepositGet) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *DepositGet) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *DepositGet) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.

### GetAccountId

`func (o *DepositGet) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DepositGet) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DepositGet) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBackImageId

`func (o *DepositGet) GetBackImageId() string`

GetBackImageId returns the BackImageId field if non-nil, zero value otherwise.

### GetBackImageIdOk

`func (o *DepositGet) GetBackImageIdOk() (*string, bool)`

GetBackImageIdOk returns a tuple with the BackImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackImageId

`func (o *DepositGet) SetBackImageId(v string)`

SetBackImageId sets BackImageId field to given value.


### GetBusinessId

`func (o *DepositGet) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *DepositGet) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *DepositGet) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *DepositGet) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCheckAmount

`func (o *DepositGet) GetCheckAmount() int32`

GetCheckAmount returns the CheckAmount field if non-nil, zero value otherwise.

### GetCheckAmountOk

`func (o *DepositGet) GetCheckAmountOk() (*int32, bool)`

GetCheckAmountOk returns a tuple with the CheckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAmount

`func (o *DepositGet) SetCheckAmount(v int32)`

SetCheckAmount sets CheckAmount field to given value.


### GetDepositCurrency

`func (o *DepositGet) GetDepositCurrency() string`

GetDepositCurrency returns the DepositCurrency field if non-nil, zero value otherwise.

### GetDepositCurrencyOk

`func (o *DepositGet) GetDepositCurrencyOk() (*string, bool)`

GetDepositCurrencyOk returns a tuple with the DepositCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositCurrency

`func (o *DepositGet) SetDepositCurrency(v string)`

SetDepositCurrency sets DepositCurrency field to given value.


### GetFrontImageId

`func (o *DepositGet) GetFrontImageId() string`

GetFrontImageId returns the FrontImageId field if non-nil, zero value otherwise.

### GetFrontImageIdOk

`func (o *DepositGet) GetFrontImageIdOk() (*string, bool)`

GetFrontImageIdOk returns a tuple with the FrontImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontImageId

`func (o *DepositGet) SetFrontImageId(v string)`

SetFrontImageId sets FrontImageId field to given value.


### GetMetadata

`func (o *DepositGet) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DepositGet) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DepositGet) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DepositGet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPersonId

`func (o *DepositGet) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *DepositGet) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *DepositGet) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *DepositGet) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


