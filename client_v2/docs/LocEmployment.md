# LocEmployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerName** | Pointer to **string** |  | [optional] 
**EmploymentStartDate** | Pointer to **string** |  | [optional] 
**Income** | Pointer to **int32** | Annual income in cents (e.g., 12000000 &#x3D; $120,000) | [optional] 
**Position** | Pointer to **string** |  | [optional] 

## Methods

### NewLocEmployment

`func NewLocEmployment() *LocEmployment`

NewLocEmployment instantiates a new LocEmployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocEmploymentWithDefaults

`func NewLocEmploymentWithDefaults() *LocEmployment`

NewLocEmploymentWithDefaults instantiates a new LocEmployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerName

`func (o *LocEmployment) GetEmployerName() string`

GetEmployerName returns the EmployerName field if non-nil, zero value otherwise.

### GetEmployerNameOk

`func (o *LocEmployment) GetEmployerNameOk() (*string, bool)`

GetEmployerNameOk returns a tuple with the EmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerName

`func (o *LocEmployment) SetEmployerName(v string)`

SetEmployerName sets EmployerName field to given value.

### HasEmployerName

`func (o *LocEmployment) HasEmployerName() bool`

HasEmployerName returns a boolean if a field has been set.

### GetEmploymentStartDate

`func (o *LocEmployment) GetEmploymentStartDate() string`

GetEmploymentStartDate returns the EmploymentStartDate field if non-nil, zero value otherwise.

### GetEmploymentStartDateOk

`func (o *LocEmployment) GetEmploymentStartDateOk() (*string, bool)`

GetEmploymentStartDateOk returns a tuple with the EmploymentStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmploymentStartDate

`func (o *LocEmployment) SetEmploymentStartDate(v string)`

SetEmploymentStartDate sets EmploymentStartDate field to given value.

### HasEmploymentStartDate

`func (o *LocEmployment) HasEmploymentStartDate() bool`

HasEmploymentStartDate returns a boolean if a field has been set.

### GetIncome

`func (o *LocEmployment) GetIncome() int32`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *LocEmployment) GetIncomeOk() (*int32, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *LocEmployment) SetIncome(v int32)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *LocEmployment) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetPosition

`func (o *LocEmployment) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *LocEmployment) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *LocEmployment) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *LocEmployment) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


