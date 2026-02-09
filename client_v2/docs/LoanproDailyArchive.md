# LoanproDailyArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveDate** | Pointer to **string** | Date of this snapshot | [optional] 
**BucketInterestRates** | Pointer to [**[]LoanproBucketInterestRate**](LoanproBucketInterestRate.md) | Interest rates for each transaction bucket | [optional] 

## Methods

### NewLoanproDailyArchive

`func NewLoanproDailyArchive() *LoanproDailyArchive`

NewLoanproDailyArchive instantiates a new LoanproDailyArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoanproDailyArchiveWithDefaults

`func NewLoanproDailyArchiveWithDefaults() *LoanproDailyArchive`

NewLoanproDailyArchiveWithDefaults instantiates a new LoanproDailyArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveDate

`func (o *LoanproDailyArchive) GetArchiveDate() string`

GetArchiveDate returns the ArchiveDate field if non-nil, zero value otherwise.

### GetArchiveDateOk

`func (o *LoanproDailyArchive) GetArchiveDateOk() (*string, bool)`

GetArchiveDateOk returns a tuple with the ArchiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveDate

`func (o *LoanproDailyArchive) SetArchiveDate(v string)`

SetArchiveDate sets ArchiveDate field to given value.

### HasArchiveDate

`func (o *LoanproDailyArchive) HasArchiveDate() bool`

HasArchiveDate returns a boolean if a field has been set.

### GetBucketInterestRates

`func (o *LoanproDailyArchive) GetBucketInterestRates() []LoanproBucketInterestRate`

GetBucketInterestRates returns the BucketInterestRates field if non-nil, zero value otherwise.

### GetBucketInterestRatesOk

`func (o *LoanproDailyArchive) GetBucketInterestRatesOk() (*[]LoanproBucketInterestRate, bool)`

GetBucketInterestRatesOk returns a tuple with the BucketInterestRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketInterestRates

`func (o *LoanproDailyArchive) SetBucketInterestRates(v []LoanproBucketInterestRate)`

SetBucketInterestRates sets BucketInterestRates field to given value.

### HasBucketInterestRates

`func (o *LoanproDailyArchive) HasBucketInterestRates() bool`

HasBucketInterestRates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


