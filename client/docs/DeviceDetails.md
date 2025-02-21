# DeviceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorDepth** | **int32** | The device&#39;s color palette bit depth, in bits per pixel | 
**Header** | **string** | The exact contents of the browser&#39;s HTTP accept headers | 
**IpAddress** | **string** | The device&#39;s IP address (either IPv4 or IPv6 formats are acceptable) | 
**JavaEnabled** | **bool** | If the browser has the ability to execute Java (value returned from &#x60;navigator.javaEnabled&#x60; property) | 
**JavascriptEnabled** | **bool** | If the browser has the ability to execute JavaScript (value returned from &#x60;navigator.javaScriptEnabled&#x60; property) | 
**Language** | **string** | The browser&#39;s language as defined in IETF BCP47 | 
**ScreenHeight** | **int32** | The total height of the device&#39;s screen, in pixels | 
**ScreenWidth** | **int32** | The total width of the device&#39;s screen, in pixels | 
**Timezone** | **int32** | The offset from UTC of the device&#39;s local timezone, in minutes | 
**UserAgent** | **string** | The exact contents of the HTTP user agent header | 

## Methods

### NewDeviceDetails

`func NewDeviceDetails(colorDepth int32, header string, ipAddress string, javaEnabled bool, javascriptEnabled bool, language string, screenHeight int32, screenWidth int32, timezone int32, userAgent string, ) *DeviceDetails`

NewDeviceDetails instantiates a new DeviceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceDetailsWithDefaults

`func NewDeviceDetailsWithDefaults() *DeviceDetails`

NewDeviceDetailsWithDefaults instantiates a new DeviceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColorDepth

`func (o *DeviceDetails) GetColorDepth() int32`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *DeviceDetails) GetColorDepthOk() (*int32, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *DeviceDetails) SetColorDepth(v int32)`

SetColorDepth sets ColorDepth field to given value.


### GetHeader

`func (o *DeviceDetails) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *DeviceDetails) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *DeviceDetails) SetHeader(v string)`

SetHeader sets Header field to given value.


### GetIpAddress

`func (o *DeviceDetails) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DeviceDetails) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DeviceDetails) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetJavaEnabled

`func (o *DeviceDetails) GetJavaEnabled() bool`

GetJavaEnabled returns the JavaEnabled field if non-nil, zero value otherwise.

### GetJavaEnabledOk

`func (o *DeviceDetails) GetJavaEnabledOk() (*bool, bool)`

GetJavaEnabledOk returns a tuple with the JavaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaEnabled

`func (o *DeviceDetails) SetJavaEnabled(v bool)`

SetJavaEnabled sets JavaEnabled field to given value.


### GetJavascriptEnabled

`func (o *DeviceDetails) GetJavascriptEnabled() bool`

GetJavascriptEnabled returns the JavascriptEnabled field if non-nil, zero value otherwise.

### GetJavascriptEnabledOk

`func (o *DeviceDetails) GetJavascriptEnabledOk() (*bool, bool)`

GetJavascriptEnabledOk returns a tuple with the JavascriptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptEnabled

`func (o *DeviceDetails) SetJavascriptEnabled(v bool)`

SetJavascriptEnabled sets JavascriptEnabled field to given value.


### GetLanguage

`func (o *DeviceDetails) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *DeviceDetails) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *DeviceDetails) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetScreenHeight

`func (o *DeviceDetails) GetScreenHeight() int32`

GetScreenHeight returns the ScreenHeight field if non-nil, zero value otherwise.

### GetScreenHeightOk

`func (o *DeviceDetails) GetScreenHeightOk() (*int32, bool)`

GetScreenHeightOk returns a tuple with the ScreenHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenHeight

`func (o *DeviceDetails) SetScreenHeight(v int32)`

SetScreenHeight sets ScreenHeight field to given value.


### GetScreenWidth

`func (o *DeviceDetails) GetScreenWidth() int32`

GetScreenWidth returns the ScreenWidth field if non-nil, zero value otherwise.

### GetScreenWidthOk

`func (o *DeviceDetails) GetScreenWidthOk() (*int32, bool)`

GetScreenWidthOk returns a tuple with the ScreenWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenWidth

`func (o *DeviceDetails) SetScreenWidth(v int32)`

SetScreenWidth sets ScreenWidth field to given value.


### GetTimezone

`func (o *DeviceDetails) GetTimezone() int32`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DeviceDetails) GetTimezoneOk() (*int32, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DeviceDetails) SetTimezone(v int32)`

SetTimezone sets Timezone field to given value.


### GetUserAgent

`func (o *DeviceDetails) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *DeviceDetails) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *DeviceDetails) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


