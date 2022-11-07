# VirtualCardResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | Pointer to **string** | The bin number | [optional] 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 

## Methods

### NewVirtualCardResponseAllOf

`func NewVirtualCardResponseAllOf(cardBrand CardBrand, ) *VirtualCardResponseAllOf`

NewVirtualCardResponseAllOf instantiates a new VirtualCardResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCardResponseAllOfWithDefaults

`func NewVirtualCardResponseAllOfWithDefaults() *VirtualCardResponseAllOf`

NewVirtualCardResponseAllOfWithDefaults instantiates a new VirtualCardResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *VirtualCardResponseAllOf) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *VirtualCardResponseAllOf) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *VirtualCardResponseAllOf) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *VirtualCardResponseAllOf) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCardBrand

`func (o *VirtualCardResponseAllOf) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *VirtualCardResponseAllOf) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *VirtualCardResponseAllOf) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


