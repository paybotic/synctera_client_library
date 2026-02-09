# Applicant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdverseActionId** | Pointer to **string** | Adverse Action ID for the applicant. | [optional] 
**BusinessId** | Pointer to **string** | Business ID for the application. Only one of customer_id or business_id can be provided.  | [optional] 
**CreditScoreIds** | Pointer to **[]string** | List of credit score IDs for the applicant of the credit application | [optional] 
**CustomerId** | Pointer to **string** | Customer ID for the application. Only one of customer_id or business_id can be provided.  | [optional] 
**IsPrimary** | **bool** | Whether this applicant is the primary applicant | 
**UnderwritingData** | Pointer to [**[]UnderwritingData**](UnderwritingData.md) |  | [optional] 

## Methods

### NewApplicant

`func NewApplicant(isPrimary bool, ) *Applicant`

NewApplicant instantiates a new Applicant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicantWithDefaults

`func NewApplicantWithDefaults() *Applicant`

NewApplicantWithDefaults instantiates a new Applicant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdverseActionId

`func (o *Applicant) GetAdverseActionId() string`

GetAdverseActionId returns the AdverseActionId field if non-nil, zero value otherwise.

### GetAdverseActionIdOk

`func (o *Applicant) GetAdverseActionIdOk() (*string, bool)`

GetAdverseActionIdOk returns a tuple with the AdverseActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdverseActionId

`func (o *Applicant) SetAdverseActionId(v string)`

SetAdverseActionId sets AdverseActionId field to given value.

### HasAdverseActionId

`func (o *Applicant) HasAdverseActionId() bool`

HasAdverseActionId returns a boolean if a field has been set.

### GetBusinessId

`func (o *Applicant) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *Applicant) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *Applicant) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *Applicant) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreditScoreIds

`func (o *Applicant) GetCreditScoreIds() []string`

GetCreditScoreIds returns the CreditScoreIds field if non-nil, zero value otherwise.

### GetCreditScoreIdsOk

`func (o *Applicant) GetCreditScoreIdsOk() (*[]string, bool)`

GetCreditScoreIdsOk returns a tuple with the CreditScoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditScoreIds

`func (o *Applicant) SetCreditScoreIds(v []string)`

SetCreditScoreIds sets CreditScoreIds field to given value.

### HasCreditScoreIds

`func (o *Applicant) HasCreditScoreIds() bool`

HasCreditScoreIds returns a boolean if a field has been set.

### GetCustomerId

`func (o *Applicant) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Applicant) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Applicant) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Applicant) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetIsPrimary

`func (o *Applicant) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *Applicant) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *Applicant) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### GetUnderwritingData

`func (o *Applicant) GetUnderwritingData() []UnderwritingData`

GetUnderwritingData returns the UnderwritingData field if non-nil, zero value otherwise.

### GetUnderwritingDataOk

`func (o *Applicant) GetUnderwritingDataOk() (*[]UnderwritingData, bool)`

GetUnderwritingDataOk returns a tuple with the UnderwritingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderwritingData

`func (o *Applicant) SetUnderwritingData(v []UnderwritingData)`

SetUnderwritingData sets UnderwritingData field to given value.

### HasUnderwritingData

`func (o *Applicant) HasUnderwritingData() bool`

HasUnderwritingData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


