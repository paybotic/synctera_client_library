# Detail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Machine-readable identifier to allow grouping details into subsets.  Checks for the following categories should typcially be provided when onboarding personal customers: * &#x60;CIP&#x60; – Checks relating to a Customer Identification Program (CIP) including name, ID verification, and date of birth. * &#x60;ADDRESS&#x60; – Address check done either through a bureau or document verification. * &#x60;PHONE&#x60; – Phone checks, including verifying number and other correlation methods. * &#x60;EMAIL&#x60; – Email checks, including verifying email address and other correlation methods. * &#x60;FRAUD&#x60; – Checks relating to elements of fraud. * &#x60;SYNTHETIC&#x60; – Check relating to elements of synthetic identity. * &#x60;DEVICE&#x60; – Checks relating to elements of device risk. Can include SDK or IP screening. * &#x60;DOC_VERIFICATION&#x60; – Document Verification checks, where ID documents are submitted for verification, including liveness test, or other anti-forgery verification. * &#x60;SELFIE_CAPTURE&#x60; – Selfie capture check, including liveness test, or other anti-fake verification. * &#x60;SELFIE_DOCUMENT&#x60; – Check if selfie matches with photo on document. * &#x60;DOC_DETAILS&#x60; – Check if details on ID documents matches with supplied information. * &#x60;WATCHLIST&#x60; – Outcome from watchlist screenings (especially OFAC/SDN and other sanctions lists).  Checks for the following categories should typcially be provided when onboarding businesses customers: * &#x60;BUSINESS_NAME&#x60; – Business name check done against a bureau. * &#x60;OFFICE_ADDRESS&#x60; – Address check done either through a bureau or document verification. * &#x60;SOS_FILINGS&#x60; – Findings from a Secretary of State (SOS) filings check. * &#x60;WEBSITE&#x60; – Business website check. * &#x60;TIN_MATCH&#x60; – Taxpayer Identification Number (TIN) check from tax bureau. * &#x60;BANKRUPTCIES&#x60; – Check for any bankruptcy filings. * &#x60;SOS_DOMESTIC&#x60; – Outcome from a Domestic Secretary of State (SOS) filings check * &#x60;LICENSE&#x60; –  * &#x60;WATCHLIST&#x60; – Outcome from watchlist screenings (especially OFAC/SDN and other sanctions lists).  If you would like to provide additional information that doesn&#39;t fit with an existing category you may provide: * &#x60;OTHER&#x60; – A category for miscellaneous or uncategorized checks.  | [optional] 
**Description** | Pointer to **string** | Human-readable description explaining the individual check. | [optional] 
**Label** | Pointer to **string** | Human-readable grouping describing the aspect of the customer&#39;s identity examined by this check. | [optional] [readonly] 
**Result** | Pointer to **string** | The result of the individual check. One of the following: * &#x60;PASS&#x60; – the check passed contributing to a positive outcome (or accepted verification result). * &#x60;INFO&#x60; – the check returned neutral information which may or may not explain a negative result. * &#x60;WARN&#x60; – the check was inconclusive and might require review. * &#x60;FAIL&#x60; – the check failed and might result in a failing outcome (or rejected verification_result).  | [optional] 
**Score** | Pointer to **float64** | An arbitrary floating point score value which may be used to contextualize the human readable description. | [optional] 
**Url** | Pointer to **string** | A URL containing supporting information for this individual check. | [optional] 
**VendorCode** | Pointer to **string** | Machine-readable description of the individual check. This field contains vendor-specific terms and may not be populated in all cases. | [optional] 

## Methods

### NewDetail

`func NewDetail() *Detail`

NewDetail instantiates a new Detail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailWithDefaults

`func NewDetailWithDefaults() *Detail`

NewDetailWithDefaults instantiates a new Detail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Detail) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Detail) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Detail) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Detail) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *Detail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Detail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Detail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Detail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *Detail) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Detail) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Detail) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Detail) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetResult

`func (o *Detail) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Detail) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Detail) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Detail) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetScore

`func (o *Detail) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Detail) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Detail) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *Detail) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetUrl

`func (o *Detail) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Detail) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Detail) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Detail) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVendorCode

`func (o *Detail) GetVendorCode() string`

GetVendorCode returns the VendorCode field if non-nil, zero value otherwise.

### GetVendorCodeOk

`func (o *Detail) GetVendorCodeOk() (*string, bool)`

GetVendorCodeOk returns a tuple with the VendorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCode

`func (o *Detail) SetVendorCode(v string)`

SetVendorCode sets VendorCode field to given value.

### HasVendorCode

`func (o *Detail) HasVendorCode() bool`

HasVendorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


