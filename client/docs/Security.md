# Security

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkedAccountId** | **string** | ID of linked backing account for use as a security, e.g. for use in a Smart Charge Card offering. Must be of type CHECKING or SAVING.  | 

## Methods

### NewSecurity

`func NewSecurity(linkedAccountId string, ) *Security`

NewSecurity instantiates a new Security object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityWithDefaults

`func NewSecurityWithDefaults() *Security`

NewSecurityWithDefaults instantiates a new Security object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkedAccountId

`func (o *Security) GetLinkedAccountId() string`

GetLinkedAccountId returns the LinkedAccountId field if non-nil, zero value otherwise.

### GetLinkedAccountIdOk

`func (o *Security) GetLinkedAccountIdOk() (*string, bool)`

GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountId

`func (o *Security) SetLinkedAccountId(v string)`

SetLinkedAccountId sets LinkedAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


