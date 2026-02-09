# Authenticate3dsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChallengeJwt** | **string** | The JWT recieved from the 3DS challenge | 
**Id** | **string** | The unique identifier of the 3DS authentication | 

## Methods

### NewAuthenticate3dsRequest

`func NewAuthenticate3dsRequest(challengeJwt string, id string, ) *Authenticate3dsRequest`

NewAuthenticate3dsRequest instantiates a new Authenticate3dsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticate3dsRequestWithDefaults

`func NewAuthenticate3dsRequestWithDefaults() *Authenticate3dsRequest`

NewAuthenticate3dsRequestWithDefaults instantiates a new Authenticate3dsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallengeJwt

`func (o *Authenticate3dsRequest) GetChallengeJwt() string`

GetChallengeJwt returns the ChallengeJwt field if non-nil, zero value otherwise.

### GetChallengeJwtOk

`func (o *Authenticate3dsRequest) GetChallengeJwtOk() (*string, bool)`

GetChallengeJwtOk returns a tuple with the ChallengeJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeJwt

`func (o *Authenticate3dsRequest) SetChallengeJwt(v string)`

SetChallengeJwt sets ChallengeJwt field to given value.


### GetId

`func (o *Authenticate3dsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Authenticate3dsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Authenticate3dsRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


