# CardProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | 
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
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the Card Product was created | [optional] [readonly] 
**CrossBorderEnabled** | Pointer to **bool** | Enable/Disable cross border transaction - transaction will be automatically declined when merchant country is different than BIN country when disabled. Cross-Border transaction are disabled by default.  | [optional] [default to false]
**DigitalWalletTokenization** | Pointer to [**DigitalWalletTokenization**](DigitalWalletTokenization.md) |  | [optional] 
**EndDate** | Pointer to **time.Time** | The time when the Card Product is decommissioned | [optional] 
**Id** | Pointer to **string** | Card Product ID | [optional] [readonly] 
**Image** | Pointer to **bool** | Indicates whether or not there is an overlay image of the card product available | [optional] 
**ImageMode** | Pointer to [**CardImageMode**](CardImageMode.md) |  | [optional] 
**IssueWithoutKyc** | Pointer to **bool** | Allow issuing cards on this product without requiring KYC | [optional] 
**L2l3Enabled** | Pointer to **bool** | Enable/Disable l2l3 transaction - L2l3 transaction are disabled by default.  | [optional] [default to false]
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the Card Product was last modified | [optional] [readonly] 
**Name** | **string** | The name of the Card Product | 
**NotificationLanguage** | Pointer to [**NotificationLanguage**](NotificationLanguage.md) |  | [optional] 
**Orientation** | Pointer to **string** | Card orientation | [optional] 
**PackageId** | Pointer to **string** | Card fulfillment provider’s package ID | [optional] 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**ReturnAddress** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 
**StartDate** | **time.Time** | The time when the Card Product goes live | 
**TxnEnhancer** | Pointer to [**TxnEnhancer**](TxnEnhancer.md) |  | [optional] [default to TXNENHANCER_MX]

## Methods

### NewCardProduct

`func NewCardProduct(form string, active bool, cardProgramId string, name string, startDate time.Time, ) *CardProduct`

NewCardProduct instantiates a new CardProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductWithDefaults

`func NewCardProductWithDefaults() *CardProduct`

NewCardProductWithDefaults instantiates a new CardProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *CardProduct) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CardProduct) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CardProduct) SetForm(v string)`

SetForm sets Form field to given value.


### GetActive

`func (o *CardProduct) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProduct) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProduct) SetActive(v bool)`

SetActive sets Active field to given value.


### GetBinCountry

`func (o *CardProduct) GetBinCountry() string`

GetBinCountry returns the BinCountry field if non-nil, zero value otherwise.

### GetBinCountryOk

`func (o *CardProduct) GetBinCountryOk() (*string, bool)`

GetBinCountryOk returns a tuple with the BinCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCountry

`func (o *CardProduct) SetBinCountry(v string)`

SetBinCountry sets BinCountry field to given value.

### HasBinCountry

`func (o *CardProduct) HasBinCountry() bool`

HasBinCountry returns a boolean if a field has been set.

### GetBypassRiskErrors

`func (o *CardProduct) GetBypassRiskErrors() bool`

GetBypassRiskErrors returns the BypassRiskErrors field if non-nil, zero value otherwise.

### GetBypassRiskErrorsOk

`func (o *CardProduct) GetBypassRiskErrorsOk() (*bool, bool)`

GetBypassRiskErrorsOk returns a tuple with the BypassRiskErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassRiskErrors

`func (o *CardProduct) SetBypassRiskErrors(v bool)`

SetBypassRiskErrors sets BypassRiskErrors field to given value.

### HasBypassRiskErrors

`func (o *CardProduct) HasBypassRiskErrors() bool`

HasBypassRiskErrors returns a boolean if a field has been set.

### GetCardBrand

`func (o *CardProduct) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *CardProduct) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *CardProduct) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *CardProduct) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### GetCardCategory

`func (o *CardProduct) GetCardCategory() CardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *CardProduct) GetCardCategoryOk() (*CardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *CardProduct) SetCardCategory(v CardCategory)`

SetCardCategory sets CardCategory field to given value.

### HasCardCategory

`func (o *CardProduct) HasCardCategory() bool`

HasCardCategory returns a boolean if a field has been set.

### GetCardFulfillmentCountry

`func (o *CardProduct) GetCardFulfillmentCountry() string`

GetCardFulfillmentCountry returns the CardFulfillmentCountry field if non-nil, zero value otherwise.

### GetCardFulfillmentCountryOk

`func (o *CardProduct) GetCardFulfillmentCountryOk() (*string, bool)`

GetCardFulfillmentCountryOk returns a tuple with the CardFulfillmentCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentCountry

`func (o *CardProduct) SetCardFulfillmentCountry(v string)`

SetCardFulfillmentCountry sets CardFulfillmentCountry field to given value.

### HasCardFulfillmentCountry

`func (o *CardProduct) HasCardFulfillmentCountry() bool`

HasCardFulfillmentCountry returns a boolean if a field has been set.

### GetCardFulfillmentProvider

`func (o *CardProduct) GetCardFulfillmentProvider() CardFulfillmentProvider`

GetCardFulfillmentProvider returns the CardFulfillmentProvider field if non-nil, zero value otherwise.

### GetCardFulfillmentProviderOk

`func (o *CardProduct) GetCardFulfillmentProviderOk() (*CardFulfillmentProvider, bool)`

GetCardFulfillmentProviderOk returns a tuple with the CardFulfillmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentProvider

`func (o *CardProduct) SetCardFulfillmentProvider(v CardFulfillmentProvider)`

SetCardFulfillmentProvider sets CardFulfillmentProvider field to given value.

### HasCardFulfillmentProvider

`func (o *CardProduct) HasCardFulfillmentProvider() bool`

HasCardFulfillmentProvider returns a boolean if a field has been set.

### GetCardProgramId

`func (o *CardProduct) GetCardProgramId() string`

GetCardProgramId returns the CardProgramId field if non-nil, zero value otherwise.

### GetCardProgramIdOk

`func (o *CardProduct) GetCardProgramIdOk() (*string, bool)`

GetCardProgramIdOk returns a tuple with the CardProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProgramId

`func (o *CardProduct) SetCardProgramId(v string)`

SetCardProgramId sets CardProgramId field to given value.


### GetCardType

`func (o *CardProduct) GetCardType() CardProgramType`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CardProduct) GetCardTypeOk() (*CardProgramType, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CardProduct) SetCardType(v CardProgramType)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CardProduct) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetColor

`func (o *CardProduct) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CardProduct) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CardProduct) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CardProduct) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCreationTime

`func (o *CardProduct) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardProduct) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardProduct) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CardProduct) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCrossBorderEnabled

`func (o *CardProduct) GetCrossBorderEnabled() bool`

GetCrossBorderEnabled returns the CrossBorderEnabled field if non-nil, zero value otherwise.

### GetCrossBorderEnabledOk

`func (o *CardProduct) GetCrossBorderEnabledOk() (*bool, bool)`

GetCrossBorderEnabledOk returns a tuple with the CrossBorderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossBorderEnabled

`func (o *CardProduct) SetCrossBorderEnabled(v bool)`

SetCrossBorderEnabled sets CrossBorderEnabled field to given value.

### HasCrossBorderEnabled

`func (o *CardProduct) HasCrossBorderEnabled() bool`

HasCrossBorderEnabled returns a boolean if a field has been set.

### GetDigitalWalletTokenization

`func (o *CardProduct) GetDigitalWalletTokenization() DigitalWalletTokenization`

GetDigitalWalletTokenization returns the DigitalWalletTokenization field if non-nil, zero value otherwise.

### GetDigitalWalletTokenizationOk

`func (o *CardProduct) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool)`

GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokenization

`func (o *CardProduct) SetDigitalWalletTokenization(v DigitalWalletTokenization)`

SetDigitalWalletTokenization sets DigitalWalletTokenization field to given value.

### HasDigitalWalletTokenization

`func (o *CardProduct) HasDigitalWalletTokenization() bool`

HasDigitalWalletTokenization returns a boolean if a field has been set.

### GetEndDate

`func (o *CardProduct) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProduct) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProduct) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProduct) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *CardProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *CardProduct) GetImage() bool`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CardProduct) GetImageOk() (*bool, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CardProduct) SetImage(v bool)`

SetImage sets Image field to given value.

### HasImage

`func (o *CardProduct) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageMode

`func (o *CardProduct) GetImageMode() CardImageMode`

GetImageMode returns the ImageMode field if non-nil, zero value otherwise.

### GetImageModeOk

`func (o *CardProduct) GetImageModeOk() (*CardImageMode, bool)`

GetImageModeOk returns a tuple with the ImageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMode

`func (o *CardProduct) SetImageMode(v CardImageMode)`

SetImageMode sets ImageMode field to given value.

### HasImageMode

`func (o *CardProduct) HasImageMode() bool`

HasImageMode returns a boolean if a field has been set.

### GetIssueWithoutKyc

`func (o *CardProduct) GetIssueWithoutKyc() bool`

GetIssueWithoutKyc returns the IssueWithoutKyc field if non-nil, zero value otherwise.

### GetIssueWithoutKycOk

`func (o *CardProduct) GetIssueWithoutKycOk() (*bool, bool)`

GetIssueWithoutKycOk returns a tuple with the IssueWithoutKyc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueWithoutKyc

`func (o *CardProduct) SetIssueWithoutKyc(v bool)`

SetIssueWithoutKyc sets IssueWithoutKyc field to given value.

### HasIssueWithoutKyc

`func (o *CardProduct) HasIssueWithoutKyc() bool`

HasIssueWithoutKyc returns a boolean if a field has been set.

### GetL2l3Enabled

`func (o *CardProduct) GetL2l3Enabled() bool`

GetL2l3Enabled returns the L2l3Enabled field if non-nil, zero value otherwise.

### GetL2l3EnabledOk

`func (o *CardProduct) GetL2l3EnabledOk() (*bool, bool)`

GetL2l3EnabledOk returns a tuple with the L2l3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2l3Enabled

`func (o *CardProduct) SetL2l3Enabled(v bool)`

SetL2l3Enabled sets L2l3Enabled field to given value.

### HasL2l3Enabled

`func (o *CardProduct) HasL2l3Enabled() bool`

HasL2l3Enabled returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CardProduct) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProduct) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProduct) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *CardProduct) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *CardProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProduct) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationLanguage

`func (o *CardProduct) GetNotificationLanguage() NotificationLanguage`

GetNotificationLanguage returns the NotificationLanguage field if non-nil, zero value otherwise.

### GetNotificationLanguageOk

`func (o *CardProduct) GetNotificationLanguageOk() (*NotificationLanguage, bool)`

GetNotificationLanguageOk returns a tuple with the NotificationLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationLanguage

`func (o *CardProduct) SetNotificationLanguage(v NotificationLanguage)`

SetNotificationLanguage sets NotificationLanguage field to given value.

### HasNotificationLanguage

`func (o *CardProduct) HasNotificationLanguage() bool`

HasNotificationLanguage returns a boolean if a field has been set.

### GetOrientation

`func (o *CardProduct) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *CardProduct) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *CardProduct) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *CardProduct) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetPackageId

`func (o *CardProduct) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CardProduct) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CardProduct) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CardProduct) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPhysicalCardFormat

`func (o *CardProduct) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *CardProduct) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *CardProduct) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *CardProduct) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.

### GetReturnAddress

`func (o *CardProduct) GetReturnAddress() Shipping`

GetReturnAddress returns the ReturnAddress field if non-nil, zero value otherwise.

### GetReturnAddressOk

`func (o *CardProduct) GetReturnAddressOk() (*Shipping, bool)`

GetReturnAddressOk returns a tuple with the ReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAddress

`func (o *CardProduct) SetReturnAddress(v Shipping)`

SetReturnAddress sets ReturnAddress field to given value.

### HasReturnAddress

`func (o *CardProduct) HasReturnAddress() bool`

HasReturnAddress returns a boolean if a field has been set.

### GetStartDate

`func (o *CardProduct) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProduct) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProduct) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetTxnEnhancer

`func (o *CardProduct) GetTxnEnhancer() TxnEnhancer`

GetTxnEnhancer returns the TxnEnhancer field if non-nil, zero value otherwise.

### GetTxnEnhancerOk

`func (o *CardProduct) GetTxnEnhancerOk() (*TxnEnhancer, bool)`

GetTxnEnhancerOk returns a tuple with the TxnEnhancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnEnhancer

`func (o *CardProduct) SetTxnEnhancer(v TxnEnhancer)`

SetTxnEnhancer sets TxnEnhancer field to given value.

### HasTxnEnhancer

`func (o *CardProduct) HasTxnEnhancer() bool`

HasTxnEnhancer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


