# TransactionLoanproVendorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketId** | Pointer to **int64** | LoanPro bucket ID | [optional] 
**LocId** | Pointer to **int64** | LoanPro line of credit account ID | [optional] 
**SwipeId** | Pointer to **string** | LoanPro Secure Payments swipe UUID | [optional] 
**TransactionId** | Pointer to **int64** | LoanPro LMS transaction ID | [optional] 
**TransactionType** | Pointer to [**LoanproTransactionType**](LoanproTransactionType.md) |  | [optional] 

## Methods

### NewTransactionLoanproVendorData

`func NewTransactionLoanproVendorData() *TransactionLoanproVendorData`

NewTransactionLoanproVendorData instantiates a new TransactionLoanproVendorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionLoanproVendorDataWithDefaults

`func NewTransactionLoanproVendorDataWithDefaults() *TransactionLoanproVendorData`

NewTransactionLoanproVendorDataWithDefaults instantiates a new TransactionLoanproVendorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketId

`func (o *TransactionLoanproVendorData) GetBucketId() int64`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *TransactionLoanproVendorData) GetBucketIdOk() (*int64, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *TransactionLoanproVendorData) SetBucketId(v int64)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *TransactionLoanproVendorData) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetLocId

`func (o *TransactionLoanproVendorData) GetLocId() int64`

GetLocId returns the LocId field if non-nil, zero value otherwise.

### GetLocIdOk

`func (o *TransactionLoanproVendorData) GetLocIdOk() (*int64, bool)`

GetLocIdOk returns a tuple with the LocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocId

`func (o *TransactionLoanproVendorData) SetLocId(v int64)`

SetLocId sets LocId field to given value.

### HasLocId

`func (o *TransactionLoanproVendorData) HasLocId() bool`

HasLocId returns a boolean if a field has been set.

### GetSwipeId

`func (o *TransactionLoanproVendorData) GetSwipeId() string`

GetSwipeId returns the SwipeId field if non-nil, zero value otherwise.

### GetSwipeIdOk

`func (o *TransactionLoanproVendorData) GetSwipeIdOk() (*string, bool)`

GetSwipeIdOk returns a tuple with the SwipeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwipeId

`func (o *TransactionLoanproVendorData) SetSwipeId(v string)`

SetSwipeId sets SwipeId field to given value.

### HasSwipeId

`func (o *TransactionLoanproVendorData) HasSwipeId() bool`

HasSwipeId returns a boolean if a field has been set.

### GetTransactionId

`func (o *TransactionLoanproVendorData) GetTransactionId() int64`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionLoanproVendorData) GetTransactionIdOk() (*int64, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionLoanproVendorData) SetTransactionId(v int64)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *TransactionLoanproVendorData) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionType

`func (o *TransactionLoanproVendorData) GetTransactionType() LoanproTransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *TransactionLoanproVendorData) GetTransactionTypeOk() (*LoanproTransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *TransactionLoanproVendorData) SetTransactionType(v LoanproTransactionType)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *TransactionLoanproVendorData) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


