# Lookup3dsRequestSdk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationIndicator** | **string** |  | 
**DeviceChannel** | **string** | Channel through which Device Data Collection was performed  Channel | Description --- | --- &#x60;BROWSER&#x60; | Internet browser &#x60;SDK&#x60; | Mobile app  | 
**Id** | **string** | The unique identifier of the 3DS authentication | 
**TransactionMode** | **string** |  | 

## Methods

### NewLookup3dsRequestSdk

`func NewLookup3dsRequestSdk(authenticationIndicator string, deviceChannel string, id string, transactionMode string, ) *Lookup3dsRequestSdk`

NewLookup3dsRequestSdk instantiates a new Lookup3dsRequestSdk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookup3dsRequestSdkWithDefaults

`func NewLookup3dsRequestSdkWithDefaults() *Lookup3dsRequestSdk`

NewLookup3dsRequestSdkWithDefaults instantiates a new Lookup3dsRequestSdk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationIndicator

`func (o *Lookup3dsRequestSdk) GetAuthenticationIndicator() string`

GetAuthenticationIndicator returns the AuthenticationIndicator field if non-nil, zero value otherwise.

### GetAuthenticationIndicatorOk

`func (o *Lookup3dsRequestSdk) GetAuthenticationIndicatorOk() (*string, bool)`

GetAuthenticationIndicatorOk returns a tuple with the AuthenticationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationIndicator

`func (o *Lookup3dsRequestSdk) SetAuthenticationIndicator(v string)`

SetAuthenticationIndicator sets AuthenticationIndicator field to given value.


### GetDeviceChannel

`func (o *Lookup3dsRequestSdk) GetDeviceChannel() string`

GetDeviceChannel returns the DeviceChannel field if non-nil, zero value otherwise.

### GetDeviceChannelOk

`func (o *Lookup3dsRequestSdk) GetDeviceChannelOk() (*string, bool)`

GetDeviceChannelOk returns a tuple with the DeviceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceChannel

`func (o *Lookup3dsRequestSdk) SetDeviceChannel(v string)`

SetDeviceChannel sets DeviceChannel field to given value.


### GetId

`func (o *Lookup3dsRequestSdk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Lookup3dsRequestSdk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Lookup3dsRequestSdk) SetId(v string)`

SetId sets Id field to given value.


### GetTransactionMode

`func (o *Lookup3dsRequestSdk) GetTransactionMode() string`

GetTransactionMode returns the TransactionMode field if non-nil, zero value otherwise.

### GetTransactionModeOk

`func (o *Lookup3dsRequestSdk) GetTransactionModeOk() (*string, bool)`

GetTransactionModeOk returns a tuple with the TransactionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMode

`func (o *Lookup3dsRequestSdk) SetTransactionMode(v string)`

SetTransactionMode sets TransactionMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


