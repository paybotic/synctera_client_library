# Institution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCodes** | **[]string** | The countries of operation of the financial institution | 
**Logo** | Pointer to **string** | Base64 encoded representation of the institution&#39;s logo, returned as a base64 encoded 152x152 PNG | [optional] 
**Name** | **string** | The name of the financial institution | 
**RoutingIdentifiers** | [**[]RoutingIdentifier**](RoutingIdentifier.md) | Array of routing identifier objects | 

## Methods

### NewInstitution

`func NewInstitution(countryCodes []string, name string, routingIdentifiers []RoutingIdentifier, ) *Institution`

NewInstitution instantiates a new Institution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionWithDefaults

`func NewInstitutionWithDefaults() *Institution`

NewInstitutionWithDefaults instantiates a new Institution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCodes

`func (o *Institution) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *Institution) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *Institution) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.


### GetLogo

`func (o *Institution) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Institution) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Institution) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Institution) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *Institution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Institution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Institution) SetName(v string)`

SetName sets Name field to given value.


### GetRoutingIdentifiers

`func (o *Institution) GetRoutingIdentifiers() []RoutingIdentifier`

GetRoutingIdentifiers returns the RoutingIdentifiers field if non-nil, zero value otherwise.

### GetRoutingIdentifiersOk

`func (o *Institution) GetRoutingIdentifiersOk() (*[]RoutingIdentifier, bool)`

GetRoutingIdentifiersOk returns a tuple with the RoutingIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifiers

`func (o *Institution) SetRoutingIdentifiers(v []RoutingIdentifier)`

SetRoutingIdentifiers sets RoutingIdentifiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


