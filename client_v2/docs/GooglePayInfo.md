# GooglePayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssuranceDetails** | [**GooglePayAssuranceDetails**](GooglePayAssuranceDetails.md) |  | 
**BillingAddress** | Pointer to [**GooglePayAddress**](GooglePayAddress.md) |  | [optional] 
**CardDetails** | **string** | Details about the card. This value is commonly the last four digits of the selected payment account number. | 
**CardNetwork** | **string** | Payment card network of the selected payment | 

## Methods

### NewGooglePayInfo

`func NewGooglePayInfo(assuranceDetails GooglePayAssuranceDetails, cardDetails string, cardNetwork string, ) *GooglePayInfo`

NewGooglePayInfo instantiates a new GooglePayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayInfoWithDefaults

`func NewGooglePayInfoWithDefaults() *GooglePayInfo`

NewGooglePayInfoWithDefaults instantiates a new GooglePayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssuranceDetails

`func (o *GooglePayInfo) GetAssuranceDetails() GooglePayAssuranceDetails`

GetAssuranceDetails returns the AssuranceDetails field if non-nil, zero value otherwise.

### GetAssuranceDetailsOk

`func (o *GooglePayInfo) GetAssuranceDetailsOk() (*GooglePayAssuranceDetails, bool)`

GetAssuranceDetailsOk returns a tuple with the AssuranceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceDetails

`func (o *GooglePayInfo) SetAssuranceDetails(v GooglePayAssuranceDetails)`

SetAssuranceDetails sets AssuranceDetails field to given value.


### GetBillingAddress

`func (o *GooglePayInfo) GetBillingAddress() GooglePayAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *GooglePayInfo) GetBillingAddressOk() (*GooglePayAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *GooglePayInfo) SetBillingAddress(v GooglePayAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *GooglePayInfo) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCardDetails

`func (o *GooglePayInfo) GetCardDetails() string`

GetCardDetails returns the CardDetails field if non-nil, zero value otherwise.

### GetCardDetailsOk

`func (o *GooglePayInfo) GetCardDetailsOk() (*string, bool)`

GetCardDetailsOk returns a tuple with the CardDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardDetails

`func (o *GooglePayInfo) SetCardDetails(v string)`

SetCardDetails sets CardDetails field to given value.


### GetCardNetwork

`func (o *GooglePayInfo) GetCardNetwork() string`

GetCardNetwork returns the CardNetwork field if non-nil, zero value otherwise.

### GetCardNetworkOk

`func (o *GooglePayInfo) GetCardNetworkOk() (*string, bool)`

GetCardNetworkOk returns a tuple with the CardNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNetwork

`func (o *GooglePayInfo) SetCardNetwork(v string)`

SetCardNetwork sets CardNetwork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


