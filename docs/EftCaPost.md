# EftCaPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Transfer amount in cents | 
**CustomerId** | **string** | The UUID of the Synctera customer resource that is the originator of the transfer.  | 
**DcSign** | **string** | Debit or credit sign | 
**SourceData** | Pointer to **map[string]interface{}** | Additional information to be added to the transfer | [optional] 
**TransactionCode** | **string** | The three digit transaction code that identifies the type of transaction. More information can be found here: https://www.payments.ca/sites/default/files/standard007eng.pdf.  | 
**DestinationAccountId** | **string** | The UUID of the Synctera external account resource that is the destination of the transfer. This external account represents the account on the destination bank&#39;s platform.  | 
**DestinationAccountOwnerName** | **string** | The official name of the account owner of the destination account.  | 
**EffectiveDate** | Pointer to **string** | The effective date of the transaction once it gets posted | [optional] 
**IsSameDay** | Pointer to **bool** | Send the same day (use only is_same_day without specific effective_date). | [optional] 
**OriginatingAccountId** | **string** | The UUID of the Synctera account resource originating the transfer.  | 
**OriginatingAccountOwnerName** | **string** | The official name of the account owner of the originating account.  | 

## Methods

### NewEftCaPost

`func NewEftCaPost(amount int64, customerId string, dcSign string, transactionCode string, destinationAccountId string, destinationAccountOwnerName string, originatingAccountId string, originatingAccountOwnerName string, ) *EftCaPost`

NewEftCaPost instantiates a new EftCaPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEftCaPostWithDefaults

`func NewEftCaPostWithDefaults() *EftCaPost`

NewEftCaPostWithDefaults instantiates a new EftCaPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EftCaPost) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EftCaPost) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EftCaPost) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCustomerId

`func (o *EftCaPost) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *EftCaPost) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *EftCaPost) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDcSign

`func (o *EftCaPost) GetDcSign() string`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *EftCaPost) GetDcSignOk() (*string, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *EftCaPost) SetDcSign(v string)`

SetDcSign sets DcSign field to given value.


### GetSourceData

`func (o *EftCaPost) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *EftCaPost) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *EftCaPost) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *EftCaPost) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetTransactionCode

`func (o *EftCaPost) GetTransactionCode() string`

GetTransactionCode returns the TransactionCode field if non-nil, zero value otherwise.

### GetTransactionCodeOk

`func (o *EftCaPost) GetTransactionCodeOk() (*string, bool)`

GetTransactionCodeOk returns a tuple with the TransactionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCode

`func (o *EftCaPost) SetTransactionCode(v string)`

SetTransactionCode sets TransactionCode field to given value.


### GetDestinationAccountId

`func (o *EftCaPost) GetDestinationAccountId() string`

GetDestinationAccountId returns the DestinationAccountId field if non-nil, zero value otherwise.

### GetDestinationAccountIdOk

`func (o *EftCaPost) GetDestinationAccountIdOk() (*string, bool)`

GetDestinationAccountIdOk returns a tuple with the DestinationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountId

`func (o *EftCaPost) SetDestinationAccountId(v string)`

SetDestinationAccountId sets DestinationAccountId field to given value.


### GetDestinationAccountOwnerName

`func (o *EftCaPost) GetDestinationAccountOwnerName() string`

GetDestinationAccountOwnerName returns the DestinationAccountOwnerName field if non-nil, zero value otherwise.

### GetDestinationAccountOwnerNameOk

`func (o *EftCaPost) GetDestinationAccountOwnerNameOk() (*string, bool)`

GetDestinationAccountOwnerNameOk returns a tuple with the DestinationAccountOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAccountOwnerName

`func (o *EftCaPost) SetDestinationAccountOwnerName(v string)`

SetDestinationAccountOwnerName sets DestinationAccountOwnerName field to given value.


### GetEffectiveDate

`func (o *EftCaPost) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *EftCaPost) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *EftCaPost) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *EftCaPost) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetIsSameDay

`func (o *EftCaPost) GetIsSameDay() bool`

GetIsSameDay returns the IsSameDay field if non-nil, zero value otherwise.

### GetIsSameDayOk

`func (o *EftCaPost) GetIsSameDayOk() (*bool, bool)`

GetIsSameDayOk returns a tuple with the IsSameDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSameDay

`func (o *EftCaPost) SetIsSameDay(v bool)`

SetIsSameDay sets IsSameDay field to given value.

### HasIsSameDay

`func (o *EftCaPost) HasIsSameDay() bool`

HasIsSameDay returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *EftCaPost) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *EftCaPost) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *EftCaPost) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetOriginatingAccountOwnerName

`func (o *EftCaPost) GetOriginatingAccountOwnerName() string`

GetOriginatingAccountOwnerName returns the OriginatingAccountOwnerName field if non-nil, zero value otherwise.

### GetOriginatingAccountOwnerNameOk

`func (o *EftCaPost) GetOriginatingAccountOwnerNameOk() (*string, bool)`

GetOriginatingAccountOwnerNameOk returns a tuple with the OriginatingAccountOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountOwnerName

`func (o *EftCaPost) SetOriginatingAccountOwnerName(v string)`

SetOriginatingAccountOwnerName sets OriginatingAccountOwnerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


