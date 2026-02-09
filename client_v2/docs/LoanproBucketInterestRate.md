# LoanproBucketInterestRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketDisplayName** | Pointer to **string** | Human-readable display name for the bucket | [optional] 
**BucketLabel** | Pointer to **string** | Bucket label identifier. Common values include purchase, cash_advances, and balance_transfer, but custom bucket labels may be configured in LoanPro.  | [optional] 
**RateBps** | Pointer to **int32** | Interest rate in basis points (1800 &#x3D; 18.00% APR) | [optional] 

## Methods

### NewLoanproBucketInterestRate

`func NewLoanproBucketInterestRate() *LoanproBucketInterestRate`

NewLoanproBucketInterestRate instantiates a new LoanproBucketInterestRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoanproBucketInterestRateWithDefaults

`func NewLoanproBucketInterestRateWithDefaults() *LoanproBucketInterestRate`

NewLoanproBucketInterestRateWithDefaults instantiates a new LoanproBucketInterestRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketDisplayName

`func (o *LoanproBucketInterestRate) GetBucketDisplayName() string`

GetBucketDisplayName returns the BucketDisplayName field if non-nil, zero value otherwise.

### GetBucketDisplayNameOk

`func (o *LoanproBucketInterestRate) GetBucketDisplayNameOk() (*string, bool)`

GetBucketDisplayNameOk returns a tuple with the BucketDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketDisplayName

`func (o *LoanproBucketInterestRate) SetBucketDisplayName(v string)`

SetBucketDisplayName sets BucketDisplayName field to given value.

### HasBucketDisplayName

`func (o *LoanproBucketInterestRate) HasBucketDisplayName() bool`

HasBucketDisplayName returns a boolean if a field has been set.

### GetBucketLabel

`func (o *LoanproBucketInterestRate) GetBucketLabel() string`

GetBucketLabel returns the BucketLabel field if non-nil, zero value otherwise.

### GetBucketLabelOk

`func (o *LoanproBucketInterestRate) GetBucketLabelOk() (*string, bool)`

GetBucketLabelOk returns a tuple with the BucketLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketLabel

`func (o *LoanproBucketInterestRate) SetBucketLabel(v string)`

SetBucketLabel sets BucketLabel field to given value.

### HasBucketLabel

`func (o *LoanproBucketInterestRate) HasBucketLabel() bool`

HasBucketLabel returns a boolean if a field has been set.

### GetRateBps

`func (o *LoanproBucketInterestRate) GetRateBps() int32`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *LoanproBucketInterestRate) GetRateBpsOk() (*int32, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *LoanproBucketInterestRate) SetRateBps(v int32)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *LoanproBucketInterestRate) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


