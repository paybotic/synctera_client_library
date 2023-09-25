# Lookup3dsRequestBrowser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationIndicator** | **string** |  | 
**DeviceChannel** | **string** | Channel through which Device Data Collection was performed  Channel | Description --- | --- &#x60;BROWSER&#x60; | Internet browser &#x60;SDK&#x60; | Mobile app  | 
**Id** | **string** | The unique identifier of the 3DS authentication | 
**TransactionMode** | **string** |  | 
**DeviceDetails** | Pointer to [**DeviceDetails**](DeviceDetails.md) |  | [optional] 

## Methods

### NewLookup3dsRequestBrowser

`func NewLookup3dsRequestBrowser(authenticationIndicator string, deviceChannel string, id string, transactionMode string, ) *Lookup3dsRequestBrowser`

NewLookup3dsRequestBrowser instantiates a new Lookup3dsRequestBrowser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookup3dsRequestBrowserWithDefaults

`func NewLookup3dsRequestBrowserWithDefaults() *Lookup3dsRequestBrowser`

NewLookup3dsRequestBrowserWithDefaults instantiates a new Lookup3dsRequestBrowser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationIndicator

`func (o *Lookup3dsRequestBrowser) GetAuthenticationIndicator() string`

GetAuthenticationIndicator returns the AuthenticationIndicator field if non-nil, zero value otherwise.

### GetAuthenticationIndicatorOk

`func (o *Lookup3dsRequestBrowser) GetAuthenticationIndicatorOk() (*string, bool)`

GetAuthenticationIndicatorOk returns a tuple with the AuthenticationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationIndicator

`func (o *Lookup3dsRequestBrowser) SetAuthenticationIndicator(v string)`

SetAuthenticationIndicator sets AuthenticationIndicator field to given value.


### GetDeviceChannel

`func (o *Lookup3dsRequestBrowser) GetDeviceChannel() string`

GetDeviceChannel returns the DeviceChannel field if non-nil, zero value otherwise.

### GetDeviceChannelOk

`func (o *Lookup3dsRequestBrowser) GetDeviceChannelOk() (*string, bool)`

GetDeviceChannelOk returns a tuple with the DeviceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceChannel

`func (o *Lookup3dsRequestBrowser) SetDeviceChannel(v string)`

SetDeviceChannel sets DeviceChannel field to given value.


### GetId

`func (o *Lookup3dsRequestBrowser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Lookup3dsRequestBrowser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Lookup3dsRequestBrowser) SetId(v string)`

SetId sets Id field to given value.


### GetTransactionMode

`func (o *Lookup3dsRequestBrowser) GetTransactionMode() string`

GetTransactionMode returns the TransactionMode field if non-nil, zero value otherwise.

### GetTransactionModeOk

`func (o *Lookup3dsRequestBrowser) GetTransactionModeOk() (*string, bool)`

GetTransactionModeOk returns a tuple with the TransactionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMode

`func (o *Lookup3dsRequestBrowser) SetTransactionMode(v string)`

SetTransactionMode sets TransactionMode field to given value.


### GetDeviceDetails

`func (o *Lookup3dsRequestBrowser) GetDeviceDetails() DeviceDetails`

GetDeviceDetails returns the DeviceDetails field if non-nil, zero value otherwise.

### GetDeviceDetailsOk

`func (o *Lookup3dsRequestBrowser) GetDeviceDetailsOk() (*DeviceDetails, bool)`

GetDeviceDetailsOk returns a tuple with the DeviceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetails

`func (o *Lookup3dsRequestBrowser) SetDeviceDetails(v DeviceDetails)`

SetDeviceDetails sets DeviceDetails field to given value.

### HasDeviceDetails

`func (o *Lookup3dsRequestBrowser) HasDeviceDetails() bool`

HasDeviceDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


