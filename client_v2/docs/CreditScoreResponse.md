# CreditScoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The unique identifier of a customer | 
**Score** | Pointer to **int32** | The credit score value | [optional] 
**ScoreRequestedTime** | Pointer to **time.Time** | The time the credit score was requested | [optional] 
**SourceOfScore** | [**SourceOfScore**](SourceOfScore.md) |  | 
**Type** | Pointer to **string** | The type of the credit score | [optional] 
**VendorName** | Pointer to [**VendorName**](VendorName.md) |  | [optional] 
**Version** | Pointer to **string** | The version of the credit score | [optional] 
**CreationTime** | **time.Time** | The date and time the note was created. | [readonly] 
**Id** | **string** | credit score ID | [readonly] 
**LastUpdatedTime** | **time.Time** | The date and time the note was last updated. | [readonly] 

## Methods

### NewCreditScoreResponse

`func NewCreditScoreResponse(customerId string, sourceOfScore SourceOfScore, creationTime time.Time, id string, lastUpdatedTime time.Time, ) *CreditScoreResponse`

NewCreditScoreResponse instantiates a new CreditScoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditScoreResponseWithDefaults

`func NewCreditScoreResponseWithDefaults() *CreditScoreResponse`

NewCreditScoreResponseWithDefaults instantiates a new CreditScoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CreditScoreResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreditScoreResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreditScoreResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetScore

`func (o *CreditScoreResponse) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreditScoreResponse) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreditScoreResponse) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *CreditScoreResponse) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetScoreRequestedTime

`func (o *CreditScoreResponse) GetScoreRequestedTime() time.Time`

GetScoreRequestedTime returns the ScoreRequestedTime field if non-nil, zero value otherwise.

### GetScoreRequestedTimeOk

`func (o *CreditScoreResponse) GetScoreRequestedTimeOk() (*time.Time, bool)`

GetScoreRequestedTimeOk returns a tuple with the ScoreRequestedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreRequestedTime

`func (o *CreditScoreResponse) SetScoreRequestedTime(v time.Time)`

SetScoreRequestedTime sets ScoreRequestedTime field to given value.

### HasScoreRequestedTime

`func (o *CreditScoreResponse) HasScoreRequestedTime() bool`

HasScoreRequestedTime returns a boolean if a field has been set.

### GetSourceOfScore

`func (o *CreditScoreResponse) GetSourceOfScore() SourceOfScore`

GetSourceOfScore returns the SourceOfScore field if non-nil, zero value otherwise.

### GetSourceOfScoreOk

`func (o *CreditScoreResponse) GetSourceOfScoreOk() (*SourceOfScore, bool)`

GetSourceOfScoreOk returns a tuple with the SourceOfScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfScore

`func (o *CreditScoreResponse) SetSourceOfScore(v SourceOfScore)`

SetSourceOfScore sets SourceOfScore field to given value.


### GetType

`func (o *CreditScoreResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditScoreResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditScoreResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreditScoreResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendorName

`func (o *CreditScoreResponse) GetVendorName() VendorName`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *CreditScoreResponse) GetVendorNameOk() (*VendorName, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *CreditScoreResponse) SetVendorName(v VendorName)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *CreditScoreResponse) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetVersion

`func (o *CreditScoreResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreditScoreResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreditScoreResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreditScoreResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreationTime

`func (o *CreditScoreResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CreditScoreResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CreditScoreResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetId

`func (o *CreditScoreResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditScoreResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditScoreResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *CreditScoreResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CreditScoreResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CreditScoreResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


