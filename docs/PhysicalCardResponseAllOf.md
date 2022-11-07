# PhysicalCardResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | Pointer to **string** | The bin number | [optional] 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 
**PhysicalCardFormat** | [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | 

## Methods

### NewPhysicalCardResponseAllOf

`func NewPhysicalCardResponseAllOf(cardBrand CardBrand, physicalCardFormat PhysicalCardFormat, ) *PhysicalCardResponseAllOf`

NewPhysicalCardResponseAllOf instantiates a new PhysicalCardResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalCardResponseAllOfWithDefaults

`func NewPhysicalCardResponseAllOfWithDefaults() *PhysicalCardResponseAllOf`

NewPhysicalCardResponseAllOfWithDefaults instantiates a new PhysicalCardResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *PhysicalCardResponseAllOf) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *PhysicalCardResponseAllOf) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *PhysicalCardResponseAllOf) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *PhysicalCardResponseAllOf) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCardBrand

`func (o *PhysicalCardResponseAllOf) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *PhysicalCardResponseAllOf) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *PhysicalCardResponseAllOf) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetPhysicalCardFormat

`func (o *PhysicalCardResponseAllOf) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *PhysicalCardResponseAllOf) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *PhysicalCardResponseAllOf) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


