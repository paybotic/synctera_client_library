# PersonalIdRevealResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | **string** | JWE compact serialization of the encrypted personal id data as defined in  [RFC 7516 Section 7.1](https://datatracker.ietf.org/doc/html/rfc7516#section-7.1).   Format: &#x60;header.encryptedKey.iv.ciphertext.authTag&#x60;   All segments are base64url-encoded and concatenated with periods.   - header: Contains algorithm information and key identifier    - \&quot;alg\&quot;: \&quot;ECDH-ES+A256KW\&quot; - Key management algorithm    - \&quot;enc\&quot;: \&quot;A256GCM\&quot; - Content encryption algorithm    - \&quot;kid\&quot;: base64url-encoded JWK thumbprint. The thumbprint is calculated using the JWK as per [RFC 7638](https://datatracker.ietf.org/doc/html/rfc7638#section-3.1) and is used to identify the public key used for encryption.     The algorithm identifiers follow [RFC 7518](https://datatracker.ietf.org/doc/html/rfc7518)   - encryptedKey: The content encryption key encrypted using ECDH-ES  - iv: Initialization vector for AES-GCM  - ciphertext: The actual encrypted data  - authTag: Authentication tag for AES-GCM integrity verification  | 

## Methods

### NewPersonalIdRevealResponse

`func NewPersonalIdRevealResponse(ciphertext string, ) *PersonalIdRevealResponse`

NewPersonalIdRevealResponse instantiates a new PersonalIdRevealResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalIdRevealResponseWithDefaults

`func NewPersonalIdRevealResponseWithDefaults() *PersonalIdRevealResponse`

NewPersonalIdRevealResponseWithDefaults instantiates a new PersonalIdRevealResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *PersonalIdRevealResponse) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *PersonalIdRevealResponse) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *PersonalIdRevealResponse) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


