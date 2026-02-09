# Lookup3dsRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationIndicator** | **string** |  | 
**DeviceChannel** | **string** | Channel through which Device Data Collection was performed  Channel | Description --- | --- &#x60;BROWSER&#x60; | Internet browser &#x60;SDK&#x60; | Mobile app  | 
**Id** | **string** | The unique identifier of the 3DS authentication | 
**TransactionMode** | **string** |  | 

## Methods

### NewLookup3dsRequestBase

`func NewLookup3dsRequestBase(authenticationIndicator string, deviceChannel string, id string, transactionMode string, ) *Lookup3dsRequestBase`

NewLookup3dsRequestBase instantiates a new Lookup3dsRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookup3dsRequestBaseWithDefaults

`func NewLookup3dsRequestBaseWithDefaults() *Lookup3dsRequestBase`

NewLookup3dsRequestBaseWithDefaults instantiates a new Lookup3dsRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationIndicator

`func (o *Lookup3dsRequestBase) GetAuthenticationIndicator() string`

GetAuthenticationIndicator returns the AuthenticationIndicator field if non-nil, zero value otherwise.

### GetAuthenticationIndicatorOk

`func (o *Lookup3dsRequestBase) GetAuthenticationIndicatorOk() (*string, bool)`

GetAuthenticationIndicatorOk returns a tuple with the AuthenticationIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationIndicator

`func (o *Lookup3dsRequestBase) SetAuthenticationIndicator(v string)`

SetAuthenticationIndicator sets AuthenticationIndicator field to given value.


### GetDeviceChannel

`func (o *Lookup3dsRequestBase) GetDeviceChannel() string`

GetDeviceChannel returns the DeviceChannel field if non-nil, zero value otherwise.

### GetDeviceChannelOk

`func (o *Lookup3dsRequestBase) GetDeviceChannelOk() (*string, bool)`

GetDeviceChannelOk returns a tuple with the DeviceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceChannel

`func (o *Lookup3dsRequestBase) SetDeviceChannel(v string)`

SetDeviceChannel sets DeviceChannel field to given value.


### GetId

`func (o *Lookup3dsRequestBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Lookup3dsRequestBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Lookup3dsRequestBase) SetId(v string)`

SetId sets Id field to given value.


### GetTransactionMode

`func (o *Lookup3dsRequestBase) GetTransactionMode() string`

GetTransactionMode returns the TransactionMode field if non-nil, zero value otherwise.

### GetTransactionModeOk

`func (o *Lookup3dsRequestBase) GetTransactionModeOk() (*string, bool)`

GetTransactionModeOk returns a tuple with the TransactionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMode

`func (o *Lookup3dsRequestBase) SetTransactionMode(v string)`

SetTransactionMode sets TransactionMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


