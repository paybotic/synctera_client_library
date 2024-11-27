# CardProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Indicates whether the Card Product is active | 
**BinCountry** | Pointer to **string** | ISO-3166-1 Alpha-2 country code | [optional] 
**BypassRiskErrors** | Pointer to **bool** | Allow bypassing risk engine errors. | [optional] 
**CardBrand** | Pointer to [**CardBrand**](CardBrand.md) |  | [optional] 
**CardCategory** | Pointer to [**CardCategory**](CardCategory.md) |  | [optional] 
**CardFulfillmentCountry** | Pointer to **string** | ISO-3166-1 Alpha-2 country code | [optional] 
**CardFulfillmentProvider** | Pointer to [**CardFulfillmentProvider**](CardFulfillmentProvider.md) |  | [optional] 
**CardProgramId** | **string** | Card Program ID | 
**CardType** | Pointer to [**CardProgramType**](CardProgramType.md) |  | [optional] 
**Color** | Pointer to **string** | Color code for dynamic card elements such as PAN and card holder name | [optional] 
**CreationTime** | **time.Time** | The timestamp representing when the Card Product was created | [readonly] 
**CrossBorderEnabled** | Pointer to **bool** | Enable/Disable cross border transaction - transaction will be automatically declined when merchant country is different than BIN country when disabled. Cross-Border transaction are disabled by default.  | [optional] [default to false]
**DigitalWalletTokenization** | [**DigitalWalletTokenization**](DigitalWalletTokenization.md) |  | 
**EndDate** | **time.Time** | The time when the Card Product is decommissioned | 
**Id** | **string** | Card Product ID | [readonly] 
**Image** | Pointer to **bool** | Indicates whether or not there is an overlay image of the card product available | [optional] 
**ImageMode** | Pointer to [**CardImageMode**](CardImageMode.md) |  | [optional] 
**IssueWithoutKyc** | Pointer to **bool** | Allow issuing cards on this product without requiring KYC | [optional] 
**L2l3Enabled** | Pointer to **bool** | Enable/Disable l2l3 transaction - L2l3 transaction are disabled by default.  | [optional] [default to false]
**LastModifiedTime** | **time.Time** | The timestamp representing when the Card Product was last modified | [readonly] 
**Name** | **string** | The name of the Card Product | 
**NotificationLanguage** | Pointer to [**NotificationLanguage**](NotificationLanguage.md) |  | [optional] 
**Orientation** | Pointer to **string** | Card orientation | [optional] 
**PackageId** | Pointer to **string** | Card fulfillment provider’s package ID | [optional] 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**ReturnAddress** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 
**StartDate** | **time.Time** | The time when the Card Product goes live | 
**TxnEnhancer** | Pointer to [**TxnEnhancer**](TxnEnhancer.md) |  | [optional] [default to TXNENHANCER_MX]
**Form** | **string** | PHYSICAL or VIRTUAL. | 
**ThreeDsPolicy** | [**ThreeDsPolicy**](ThreeDsPolicy.md) |  | 

## Methods

### NewCardProductResponse

`func NewCardProductResponse(active bool, cardProgramId string, creationTime time.Time, digitalWalletTokenization DigitalWalletTokenization, endDate time.Time, id string, lastModifiedTime time.Time, name string, startDate time.Time, form string, threeDsPolicy ThreeDsPolicy, ) *CardProductResponse`

NewCardProductResponse instantiates a new CardProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductResponseWithDefaults

`func NewCardProductResponseWithDefaults() *CardProductResponse`

NewCardProductResponseWithDefaults instantiates a new CardProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CardProductResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProductResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProductResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetBinCountry

`func (o *CardProductResponse) GetBinCountry() string`

GetBinCountry returns the BinCountry field if non-nil, zero value otherwise.

### GetBinCountryOk

`func (o *CardProductResponse) GetBinCountryOk() (*string, bool)`

GetBinCountryOk returns a tuple with the BinCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCountry

`func (o *CardProductResponse) SetBinCountry(v string)`

SetBinCountry sets BinCountry field to given value.

### HasBinCountry

`func (o *CardProductResponse) HasBinCountry() bool`

HasBinCountry returns a boolean if a field has been set.

### GetBypassRiskErrors

`func (o *CardProductResponse) GetBypassRiskErrors() bool`

GetBypassRiskErrors returns the BypassRiskErrors field if non-nil, zero value otherwise.

### GetBypassRiskErrorsOk

`func (o *CardProductResponse) GetBypassRiskErrorsOk() (*bool, bool)`

GetBypassRiskErrorsOk returns a tuple with the BypassRiskErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassRiskErrors

`func (o *CardProductResponse) SetBypassRiskErrors(v bool)`

SetBypassRiskErrors sets BypassRiskErrors field to given value.

### HasBypassRiskErrors

`func (o *CardProductResponse) HasBypassRiskErrors() bool`

HasBypassRiskErrors returns a boolean if a field has been set.

### GetCardBrand

`func (o *CardProductResponse) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *CardProductResponse) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *CardProductResponse) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *CardProductResponse) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### GetCardCategory

`func (o *CardProductResponse) GetCardCategory() CardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *CardProductResponse) GetCardCategoryOk() (*CardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *CardProductResponse) SetCardCategory(v CardCategory)`

SetCardCategory sets CardCategory field to given value.

### HasCardCategory

`func (o *CardProductResponse) HasCardCategory() bool`

HasCardCategory returns a boolean if a field has been set.

### GetCardFulfillmentCountry

`func (o *CardProductResponse) GetCardFulfillmentCountry() string`

GetCardFulfillmentCountry returns the CardFulfillmentCountry field if non-nil, zero value otherwise.

### GetCardFulfillmentCountryOk

`func (o *CardProductResponse) GetCardFulfillmentCountryOk() (*string, bool)`

GetCardFulfillmentCountryOk returns a tuple with the CardFulfillmentCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentCountry

`func (o *CardProductResponse) SetCardFulfillmentCountry(v string)`

SetCardFulfillmentCountry sets CardFulfillmentCountry field to given value.

### HasCardFulfillmentCountry

`func (o *CardProductResponse) HasCardFulfillmentCountry() bool`

HasCardFulfillmentCountry returns a boolean if a field has been set.

### GetCardFulfillmentProvider

`func (o *CardProductResponse) GetCardFulfillmentProvider() CardFulfillmentProvider`

GetCardFulfillmentProvider returns the CardFulfillmentProvider field if non-nil, zero value otherwise.

### GetCardFulfillmentProviderOk

`func (o *CardProductResponse) GetCardFulfillmentProviderOk() (*CardFulfillmentProvider, bool)`

GetCardFulfillmentProviderOk returns a tuple with the CardFulfillmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentProvider

`func (o *CardProductResponse) SetCardFulfillmentProvider(v CardFulfillmentProvider)`

SetCardFulfillmentProvider sets CardFulfillmentProvider field to given value.

### HasCardFulfillmentProvider

`func (o *CardProductResponse) HasCardFulfillmentProvider() bool`

HasCardFulfillmentProvider returns a boolean if a field has been set.

### GetCardProgramId

`func (o *CardProductResponse) GetCardProgramId() string`

GetCardProgramId returns the CardProgramId field if non-nil, zero value otherwise.

### GetCardProgramIdOk

`func (o *CardProductResponse) GetCardProgramIdOk() (*string, bool)`

GetCardProgramIdOk returns a tuple with the CardProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProgramId

`func (o *CardProductResponse) SetCardProgramId(v string)`

SetCardProgramId sets CardProgramId field to given value.


### GetCardType

`func (o *CardProductResponse) GetCardType() CardProgramType`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CardProductResponse) GetCardTypeOk() (*CardProgramType, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CardProductResponse) SetCardType(v CardProgramType)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CardProductResponse) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetColor

`func (o *CardProductResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CardProductResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CardProductResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CardProductResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCreationTime

`func (o *CardProductResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardProductResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardProductResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCrossBorderEnabled

`func (o *CardProductResponse) GetCrossBorderEnabled() bool`

GetCrossBorderEnabled returns the CrossBorderEnabled field if non-nil, zero value otherwise.

### GetCrossBorderEnabledOk

`func (o *CardProductResponse) GetCrossBorderEnabledOk() (*bool, bool)`

GetCrossBorderEnabledOk returns a tuple with the CrossBorderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossBorderEnabled

`func (o *CardProductResponse) SetCrossBorderEnabled(v bool)`

SetCrossBorderEnabled sets CrossBorderEnabled field to given value.

### HasCrossBorderEnabled

`func (o *CardProductResponse) HasCrossBorderEnabled() bool`

HasCrossBorderEnabled returns a boolean if a field has been set.

### GetDigitalWalletTokenization

`func (o *CardProductResponse) GetDigitalWalletTokenization() DigitalWalletTokenization`

GetDigitalWalletTokenization returns the DigitalWalletTokenization field if non-nil, zero value otherwise.

### GetDigitalWalletTokenizationOk

`func (o *CardProductResponse) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool)`

GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokenization

`func (o *CardProductResponse) SetDigitalWalletTokenization(v DigitalWalletTokenization)`

SetDigitalWalletTokenization sets DigitalWalletTokenization field to given value.


### GetEndDate

`func (o *CardProductResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProductResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProductResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetId

`func (o *CardProductResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardProductResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardProductResponse) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *CardProductResponse) GetImage() bool`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CardProductResponse) GetImageOk() (*bool, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CardProductResponse) SetImage(v bool)`

SetImage sets Image field to given value.

### HasImage

`func (o *CardProductResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageMode

`func (o *CardProductResponse) GetImageMode() CardImageMode`

GetImageMode returns the ImageMode field if non-nil, zero value otherwise.

### GetImageModeOk

`func (o *CardProductResponse) GetImageModeOk() (*CardImageMode, bool)`

GetImageModeOk returns a tuple with the ImageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMode

`func (o *CardProductResponse) SetImageMode(v CardImageMode)`

SetImageMode sets ImageMode field to given value.

### HasImageMode

`func (o *CardProductResponse) HasImageMode() bool`

HasImageMode returns a boolean if a field has been set.

### GetIssueWithoutKyc

`func (o *CardProductResponse) GetIssueWithoutKyc() bool`

GetIssueWithoutKyc returns the IssueWithoutKyc field if non-nil, zero value otherwise.

### GetIssueWithoutKycOk

`func (o *CardProductResponse) GetIssueWithoutKycOk() (*bool, bool)`

GetIssueWithoutKycOk returns a tuple with the IssueWithoutKyc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueWithoutKyc

`func (o *CardProductResponse) SetIssueWithoutKyc(v bool)`

SetIssueWithoutKyc sets IssueWithoutKyc field to given value.

### HasIssueWithoutKyc

`func (o *CardProductResponse) HasIssueWithoutKyc() bool`

HasIssueWithoutKyc returns a boolean if a field has been set.

### GetL2l3Enabled

`func (o *CardProductResponse) GetL2l3Enabled() bool`

GetL2l3Enabled returns the L2l3Enabled field if non-nil, zero value otherwise.

### GetL2l3EnabledOk

`func (o *CardProductResponse) GetL2l3EnabledOk() (*bool, bool)`

GetL2l3EnabledOk returns a tuple with the L2l3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2l3Enabled

`func (o *CardProductResponse) SetL2l3Enabled(v bool)`

SetL2l3Enabled sets L2l3Enabled field to given value.

### HasL2l3Enabled

`func (o *CardProductResponse) HasL2l3Enabled() bool`

HasL2l3Enabled returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CardProductResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProductResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProductResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetName

`func (o *CardProductResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProductResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProductResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationLanguage

`func (o *CardProductResponse) GetNotificationLanguage() NotificationLanguage`

GetNotificationLanguage returns the NotificationLanguage field if non-nil, zero value otherwise.

### GetNotificationLanguageOk

`func (o *CardProductResponse) GetNotificationLanguageOk() (*NotificationLanguage, bool)`

GetNotificationLanguageOk returns a tuple with the NotificationLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationLanguage

`func (o *CardProductResponse) SetNotificationLanguage(v NotificationLanguage)`

SetNotificationLanguage sets NotificationLanguage field to given value.

### HasNotificationLanguage

`func (o *CardProductResponse) HasNotificationLanguage() bool`

HasNotificationLanguage returns a boolean if a field has been set.

### GetOrientation

`func (o *CardProductResponse) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *CardProductResponse) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *CardProductResponse) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *CardProductResponse) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetPackageId

`func (o *CardProductResponse) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CardProductResponse) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CardProductResponse) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CardProductResponse) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPhysicalCardFormat

`func (o *CardProductResponse) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *CardProductResponse) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *CardProductResponse) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *CardProductResponse) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.

### GetReturnAddress

`func (o *CardProductResponse) GetReturnAddress() Shipping`

GetReturnAddress returns the ReturnAddress field if non-nil, zero value otherwise.

### GetReturnAddressOk

`func (o *CardProductResponse) GetReturnAddressOk() (*Shipping, bool)`

GetReturnAddressOk returns a tuple with the ReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAddress

`func (o *CardProductResponse) SetReturnAddress(v Shipping)`

SetReturnAddress sets ReturnAddress field to given value.

### HasReturnAddress

`func (o *CardProductResponse) HasReturnAddress() bool`

HasReturnAddress returns a boolean if a field has been set.

### GetStartDate

`func (o *CardProductResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProductResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProductResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetTxnEnhancer

`func (o *CardProductResponse) GetTxnEnhancer() TxnEnhancer`

GetTxnEnhancer returns the TxnEnhancer field if non-nil, zero value otherwise.

### GetTxnEnhancerOk

`func (o *CardProductResponse) GetTxnEnhancerOk() (*TxnEnhancer, bool)`

GetTxnEnhancerOk returns a tuple with the TxnEnhancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnEnhancer

`func (o *CardProductResponse) SetTxnEnhancer(v TxnEnhancer)`

SetTxnEnhancer sets TxnEnhancer field to given value.

### HasTxnEnhancer

`func (o *CardProductResponse) HasTxnEnhancer() bool`

HasTxnEnhancer returns a boolean if a field has been set.

### GetForm

`func (o *CardProductResponse) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CardProductResponse) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CardProductResponse) SetForm(v string)`

SetForm sets Form field to given value.


### GetThreeDsPolicy

`func (o *CardProductResponse) GetThreeDsPolicy() ThreeDsPolicy`

GetThreeDsPolicy returns the ThreeDsPolicy field if non-nil, zero value otherwise.

### GetThreeDsPolicyOk

`func (o *CardProductResponse) GetThreeDsPolicyOk() (*ThreeDsPolicy, bool)`

GetThreeDsPolicyOk returns a tuple with the ThreeDsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsPolicy

`func (o *CardProductResponse) SetThreeDsPolicy(v ThreeDsPolicy)`

SetThreeDsPolicy sets ThreeDsPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


