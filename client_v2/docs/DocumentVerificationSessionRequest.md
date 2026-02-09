# DocumentVerificationSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentType** | Pointer to **string** | If provided, the document_type is used to constrain the type of document that can be collected from the end-customer. | [optional] 
**Language** | Pointer to **string** | The language to be used in the document verification session. If not provided, defaults to &#x60;EN-US&#x60;. One of the following:   * &#x60;AR&#x60; – Arabic   * &#x60;AR-SA&#x60; – Arabic (Saudi Arabia)   * &#x60;BN&#x60; – Bengali   * &#x60;CS-CZ&#x60; – Czech (Czech Republic)   * &#x60;DA-DK&#x60; – Danish (Denmark)   * &#x60;DE-DE&#x60; – German (Germany)   * &#x60;EN&#x60; – English   * &#x60;EN-AU&#x60; – English (Australia)   * &#x60;EN-CA&#x60; – English (Canada)   * &#x60;EN-GB&#x60; – English (United Kingdom)   * &#x60;EN-US&#x60; – English (United States)   * &#x60;ES&#x60; – Spanish   * &#x60;ES-001&#x60; – Spanish (World)   * &#x60;ES-AR&#x60; – Spanish (Argentina)   * &#x60;ES-ES&#x60; – Spanish (Spain)   * &#x60;ES-MX&#x60; – Spanish (Mexico)   * &#x60;ES-US&#x60; – Spanish (United States)   * &#x60;FI-FI&#x60; – Finnish (Finland)   * &#x60;FR&#x60; – French   * &#x60;FR-CA&#x60; – French (Canada)   * &#x60;HE-IL&#x60; – Hebrew (Israel)   * &#x60;HI-IN&#x60; – Hindi (India)   * &#x60;HT&#x60; – Haitian Creole   * &#x60;HU-HU&#x60; – Hungarian (Hungary)   * &#x60;HY&#x60; – Armenian   * &#x60;ID-ID&#x60; – Indonesian (Indonesia)   * &#x60;IT&#x60; – Italian   * &#x60;IT-CH&#x60; – Italian (Switzerland)   * &#x60;JA-JP&#x60; – Japanese (Japan)   * &#x60;KO&#x60; – Korean   * &#x60;MS-MY&#x60; – Malay (Malaysia)   * &#x60;NL-NL&#x60; – Dutch (Netherlands)   * &#x60;NO-NO&#x60; – Norwegian (Norway)   * &#x60;PL-PL&#x60; – Polish (Poland)   * &#x60;PT-BR&#x60; – Portuguese (Brazil)   * &#x60;PT-PT&#x60; – Portuguese (Portugal)   * &#x60;RO-MO&#x60; – Romanian (Moldova)   * &#x60;RO-RO&#x60; – Romanian (Romania)   * &#x60;RU&#x60; – Russian   * &#x60;SK-SK&#x60; – Slovak (Slovakia)   * &#x60;SV-SE&#x60; – Swedish (Sweden)   * &#x60;TH-TH&#x60; – Thai (Thailand)   * &#x60;TL&#x60; – Tagalog   * &#x60;TR-TR&#x60; – Turkish (Turkey)   * &#x60;UR&#x60; – Urdu   * &#x60;VI&#x60; – Vietnamese   * &#x60;ZH-CN&#x60; – Chinese (Simplified, China)   * &#x60;ZH-HK&#x60; – Chinese (Traditional, Hong Kong)   * &#x60;ZH-TW&#x60; – Chinese (Traditional, Taiwan)  | [optional] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 
**SendMessage** | Pointer to **bool** | Send an SMS containing the document verification link to the end-customer using the phone number provided on the customer record. | [optional] 

## Methods

### NewDocumentVerificationSessionRequest

`func NewDocumentVerificationSessionRequest() *DocumentVerificationSessionRequest`

NewDocumentVerificationSessionRequest instantiates a new DocumentVerificationSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentVerificationSessionRequestWithDefaults

`func NewDocumentVerificationSessionRequestWithDefaults() *DocumentVerificationSessionRequest`

NewDocumentVerificationSessionRequestWithDefaults instantiates a new DocumentVerificationSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentType

`func (o *DocumentVerificationSessionRequest) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *DocumentVerificationSessionRequest) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *DocumentVerificationSessionRequest) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *DocumentVerificationSessionRequest) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetLanguage

`func (o *DocumentVerificationSessionRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *DocumentVerificationSessionRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *DocumentVerificationSessionRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *DocumentVerificationSessionRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPersonId

`func (o *DocumentVerificationSessionRequest) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *DocumentVerificationSessionRequest) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *DocumentVerificationSessionRequest) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *DocumentVerificationSessionRequest) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetSendMessage

`func (o *DocumentVerificationSessionRequest) GetSendMessage() bool`

GetSendMessage returns the SendMessage field if non-nil, zero value otherwise.

### GetSendMessageOk

`func (o *DocumentVerificationSessionRequest) GetSendMessageOk() (*bool, bool)`

GetSendMessageOk returns a tuple with the SendMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMessage

`func (o *DocumentVerificationSessionRequest) SetSendMessage(v bool)`

SetSendMessage sets SendMessage field to given value.

### HasSendMessage

`func (o *DocumentVerificationSessionRequest) HasSendMessage() bool`

HasSendMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


