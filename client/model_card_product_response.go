/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CardProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardProductResponse{}

// CardProductResponse struct for CardProductResponse
type CardProductResponse struct {
	// Indicates whether the Card Product is active
	Active bool `json:"active"`
	// ISO-3166-1 Alpha-2 country code
	BinCountry *string `json:"bin_country,omitempty"`
	// Allow bypassing risk engine errors.
	BypassRiskErrors *bool         `json:"bypass_risk_errors,omitempty"`
	CardBrand        *CardBrand    `json:"card_brand,omitempty"`
	CardCategory     *CardCategory `json:"card_category,omitempty"`
	// ISO-3166-1 Alpha-2 country code
	CardFulfillmentCountry  *string                  `json:"card_fulfillment_country,omitempty"`
	CardFulfillmentProvider *CardFulfillmentProvider `json:"card_fulfillment_provider,omitempty"`
	// Card Program ID
	CardProgramId string           `json:"card_program_id"`
	CardType      *CardProgramType `json:"card_type,omitempty"`
	// Color code for dynamic card elements such as PAN and card holder name
	Color *string `json:"color,omitempty" validate:"regexp=^[0-9A-F]{6}$"`
	// The timestamp representing when the Card Product was created
	CreationTime time.Time `json:"creation_time"`
	// Enable/Disable cross border transaction - transaction will be automatically declined when merchant country is different than BIN country when disabled. Cross-Border transaction are disabled by default.
	CrossBorderEnabled        *bool                     `json:"cross_border_enabled,omitempty"`
	DigitalWalletTokenization DigitalWalletTokenization `json:"digital_wallet_tokenization"`
	// The time when the Card Product is decommissioned
	EndDate time.Time `json:"end_date"`
	// Card Product ID
	Id string `json:"id"`
	// Indicates whether or not there is an overlay image of the card product available
	Image     *bool          `json:"image,omitempty"`
	ImageMode *CardImageMode `json:"image_mode,omitempty"`
	// Allow issuing cards on this product without requiring KYC
	IssueWithoutKyc *bool `json:"issue_without_kyc,omitempty"`
	// Enable/Disable l2l3 transaction - L2l3 transaction are disabled by default.
	L2l3Enabled *bool `json:"l2l3_enabled,omitempty"`
	// The timestamp representing when the Card Product was last modified
	LastModifiedTime time.Time `json:"last_modified_time"`
	// The name of the Card Product
	Name                 string                `json:"name"`
	NotificationLanguage *NotificationLanguage `json:"notification_language,omitempty"`
	// Card orientation
	Orientation *string `json:"orientation,omitempty"`
	// Card fulfillment provider’s package ID
	PackageId          *string             `json:"package_id,omitempty"`
	PhysicalCardFormat *PhysicalCardFormat `json:"physical_card_format,omitempty"`
	ReturnAddress      *Shipping           `json:"return_address,omitempty"`
	// The time when the Card Product goes live
	StartDate   time.Time    `json:"start_date"`
	TxnEnhancer *TxnEnhancer `json:"txn_enhancer,omitempty"`
	// PHYSICAL or VIRTUAL.
	Form                 string        `json:"form"`
	ThreeDsPolicy        ThreeDsPolicy `json:"three_ds_policy"`
	AdditionalProperties map[string]interface{}
}

type _CardProductResponse CardProductResponse

// NewCardProductResponse instantiates a new CardProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProductResponse(active bool, cardProgramId string, creationTime time.Time, digitalWalletTokenization DigitalWalletTokenization, endDate time.Time, id string, lastModifiedTime time.Time, name string, startDate time.Time, form string, threeDsPolicy ThreeDsPolicy) *CardProductResponse {
	this := CardProductResponse{}
	this.Active = active
	this.CardProgramId = cardProgramId
	this.CreationTime = creationTime
	var crossBorderEnabled bool = false
	this.CrossBorderEnabled = &crossBorderEnabled
	this.DigitalWalletTokenization = digitalWalletTokenization
	this.EndDate = endDate
	this.Id = id
	var l2l3Enabled bool = false
	this.L2l3Enabled = &l2l3Enabled
	this.LastModifiedTime = lastModifiedTime
	this.Name = name
	this.StartDate = startDate
	var txnEnhancer TxnEnhancer = TXNENHANCER_MX
	this.TxnEnhancer = &txnEnhancer
	this.Form = form
	this.ThreeDsPolicy = threeDsPolicy
	return &this
}

// NewCardProductResponseWithDefaults instantiates a new CardProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProductResponseWithDefaults() *CardProductResponse {
	this := CardProductResponse{}
	var crossBorderEnabled bool = false
	this.CrossBorderEnabled = &crossBorderEnabled
	var l2l3Enabled bool = false
	this.L2l3Enabled = &l2l3Enabled
	var txnEnhancer TxnEnhancer = TXNENHANCER_MX
	this.TxnEnhancer = &txnEnhancer
	return &this
}

// GetActive returns the Active field value
func (o *CardProductResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *CardProductResponse) SetActive(v bool) {
	o.Active = v
}

// GetBinCountry returns the BinCountry field value if set, zero value otherwise.
func (o *CardProductResponse) GetBinCountry() string {
	if o == nil || IsNil(o.BinCountry) {
		var ret string
		return ret
	}
	return *o.BinCountry
}

// GetBinCountryOk returns a tuple with the BinCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetBinCountryOk() (*string, bool) {
	if o == nil || IsNil(o.BinCountry) {
		return nil, false
	}
	return o.BinCountry, true
}

// HasBinCountry returns a boolean if a field has been set.
func (o *CardProductResponse) HasBinCountry() bool {
	if o != nil && !IsNil(o.BinCountry) {
		return true
	}

	return false
}

// SetBinCountry gets a reference to the given string and assigns it to the BinCountry field.
func (o *CardProductResponse) SetBinCountry(v string) {
	o.BinCountry = &v
}

// GetBypassRiskErrors returns the BypassRiskErrors field value if set, zero value otherwise.
func (o *CardProductResponse) GetBypassRiskErrors() bool {
	if o == nil || IsNil(o.BypassRiskErrors) {
		var ret bool
		return ret
	}
	return *o.BypassRiskErrors
}

// GetBypassRiskErrorsOk returns a tuple with the BypassRiskErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetBypassRiskErrorsOk() (*bool, bool) {
	if o == nil || IsNil(o.BypassRiskErrors) {
		return nil, false
	}
	return o.BypassRiskErrors, true
}

// HasBypassRiskErrors returns a boolean if a field has been set.
func (o *CardProductResponse) HasBypassRiskErrors() bool {
	if o != nil && !IsNil(o.BypassRiskErrors) {
		return true
	}

	return false
}

// SetBypassRiskErrors gets a reference to the given bool and assigns it to the BypassRiskErrors field.
func (o *CardProductResponse) SetBypassRiskErrors(v bool) {
	o.BypassRiskErrors = &v
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise.
func (o *CardProductResponse) GetCardBrand() CardBrand {
	if o == nil || IsNil(o.CardBrand) {
		var ret CardBrand
		return ret
	}
	return *o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetCardBrandOk() (*CardBrand, bool) {
	if o == nil || IsNil(o.CardBrand) {
		return nil, false
	}
	return o.CardBrand, true
}

// HasCardBrand returns a boolean if a field has been set.
func (o *CardProductResponse) HasCardBrand() bool {
	if o != nil && !IsNil(o.CardBrand) {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given CardBrand and assigns it to the CardBrand field.
func (o *CardProductResponse) SetCardBrand(v CardBrand) {
	o.CardBrand = &v
}

// GetCardCategory returns the CardCategory field value if set, zero value otherwise.
func (o *CardProductResponse) GetCardCategory() CardCategory {
	if o == nil || IsNil(o.CardCategory) {
		var ret CardCategory
		return ret
	}
	return *o.CardCategory
}

// GetCardCategoryOk returns a tuple with the CardCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetCardCategoryOk() (*CardCategory, bool) {
	if o == nil || IsNil(o.CardCategory) {
		return nil, false
	}
	return o.CardCategory, true
}

// HasCardCategory returns a boolean if a field has been set.
func (o *CardProductResponse) HasCardCategory() bool {
	if o != nil && !IsNil(o.CardCategory) {
		return true
	}

	return false
}

// SetCardCategory gets a reference to the given CardCategory and assigns it to the CardCategory field.
func (o *CardProductResponse) SetCardCategory(v CardCategory) {
	o.CardCategory = &v
}

// GetCardFulfillmentCountry returns the CardFulfillmentCountry field value if set, zero value otherwise.
func (o *CardProductResponse) GetCardFulfillmentCountry() string {
	if o == nil || IsNil(o.CardFulfillmentCountry) {
		var ret string
		return ret
	}
	return *o.CardFulfillmentCountry
}

// GetCardFulfillmentCountryOk returns a tuple with the CardFulfillmentCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetCardFulfillmentCountryOk() (*string, bool) {
	if o == nil || IsNil(o.CardFulfillmentCountry) {
		return nil, false
	}
	return o.CardFulfillmentCountry, true
}

// HasCardFulfillmentCountry returns a boolean if a field has been set.
func (o *CardProductResponse) HasCardFulfillmentCountry() bool {
	if o != nil && !IsNil(o.CardFulfillmentCountry) {
		return true
	}

	return false
}

// SetCardFulfillmentCountry gets a reference to the given string and assigns it to the CardFulfillmentCountry field.
func (o *CardProductResponse) SetCardFulfillmentCountry(v string) {
	o.CardFulfillmentCountry = &v
}

// GetCardFulfillmentProvider returns the CardFulfillmentProvider field value if set, zero value otherwise.
func (o *CardProductResponse) GetCardFulfillmentProvider() CardFulfillmentProvider {
	if o == nil || IsNil(o.CardFulfillmentProvider) {
		var ret CardFulfillmentProvider
		return ret
	}
	return *o.CardFulfillmentProvider
}

// GetCardFulfillmentProviderOk returns a tuple with the CardFulfillmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetCardFulfillmentProviderOk() (*CardFulfillmentProvider, bool) {
	if o == nil || IsNil(o.CardFulfillmentProvider) {
		return nil, false
	}
	return o.CardFulfillmentProvider, true
}

// HasCardFulfillmentProvider returns a boolean if a field has been set.
func (o *CardProductResponse) HasCardFulfillmentProvider() bool {
	if o != nil && !IsNil(o.CardFulfillmentProvider) {
		return true
	}

	return false
}

// SetCardFulfillmentProvider gets a reference to the given CardFulfillmentProvider and assigns it to the CardFulfillmentProvider field.
func (o *CardProductResponse) SetCardFulfillmentProvider(v CardFulfillmentProvider) {
	o.CardFulfillmentProvider = &v
}

// GetCardProgramId returns the CardProgramId field value
func (o *CardProductResponse) GetCardProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardProgramId
}

// GetCardProgramIdOk returns a tuple with the CardProgramId field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetCardProgramIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProgramId, true
}

// SetCardProgramId sets field value
func (o *CardProductResponse) SetCardProgramId(v string) {
	o.CardProgramId = v
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *CardProductResponse) GetCardType() CardProgramType {
	if o == nil || IsNil(o.CardType) {
		var ret CardProgramType
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetCardTypeOk() (*CardProgramType, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *CardProductResponse) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given CardProgramType and assigns it to the CardType field.
func (o *CardProductResponse) SetCardType(v CardProgramType) {
	o.CardType = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *CardProductResponse) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *CardProductResponse) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *CardProductResponse) SetColor(v string) {
	o.Color = &v
}

// GetCreationTime returns the CreationTime field value
func (o *CardProductResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *CardProductResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCrossBorderEnabled returns the CrossBorderEnabled field value if set, zero value otherwise.
func (o *CardProductResponse) GetCrossBorderEnabled() bool {
	if o == nil || IsNil(o.CrossBorderEnabled) {
		var ret bool
		return ret
	}
	return *o.CrossBorderEnabled
}

// GetCrossBorderEnabledOk returns a tuple with the CrossBorderEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetCrossBorderEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CrossBorderEnabled) {
		return nil, false
	}
	return o.CrossBorderEnabled, true
}

// HasCrossBorderEnabled returns a boolean if a field has been set.
func (o *CardProductResponse) HasCrossBorderEnabled() bool {
	if o != nil && !IsNil(o.CrossBorderEnabled) {
		return true
	}

	return false
}

// SetCrossBorderEnabled gets a reference to the given bool and assigns it to the CrossBorderEnabled field.
func (o *CardProductResponse) SetCrossBorderEnabled(v bool) {
	o.CrossBorderEnabled = &v
}

// GetDigitalWalletTokenization returns the DigitalWalletTokenization field value
func (o *CardProductResponse) GetDigitalWalletTokenization() DigitalWalletTokenization {
	if o == nil {
		var ret DigitalWalletTokenization
		return ret
	}

	return o.DigitalWalletTokenization
}

// GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DigitalWalletTokenization, true
}

// SetDigitalWalletTokenization sets field value
func (o *CardProductResponse) SetDigitalWalletTokenization(v DigitalWalletTokenization) {
	o.DigitalWalletTokenization = v
}

// GetEndDate returns the EndDate field value
func (o *CardProductResponse) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *CardProductResponse) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetId returns the Id field value
func (o *CardProductResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CardProductResponse) SetId(v string) {
	o.Id = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CardProductResponse) GetImage() bool {
	if o == nil || IsNil(o.Image) {
		var ret bool
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetImageOk() (*bool, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CardProductResponse) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given bool and assigns it to the Image field.
func (o *CardProductResponse) SetImage(v bool) {
	o.Image = &v
}

// GetImageMode returns the ImageMode field value if set, zero value otherwise.
func (o *CardProductResponse) GetImageMode() CardImageMode {
	if o == nil || IsNil(o.ImageMode) {
		var ret CardImageMode
		return ret
	}
	return *o.ImageMode
}

// GetImageModeOk returns a tuple with the ImageMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetImageModeOk() (*CardImageMode, bool) {
	if o == nil || IsNil(o.ImageMode) {
		return nil, false
	}
	return o.ImageMode, true
}

// HasImageMode returns a boolean if a field has been set.
func (o *CardProductResponse) HasImageMode() bool {
	if o != nil && !IsNil(o.ImageMode) {
		return true
	}

	return false
}

// SetImageMode gets a reference to the given CardImageMode and assigns it to the ImageMode field.
func (o *CardProductResponse) SetImageMode(v CardImageMode) {
	o.ImageMode = &v
}

// GetIssueWithoutKyc returns the IssueWithoutKyc field value if set, zero value otherwise.
func (o *CardProductResponse) GetIssueWithoutKyc() bool {
	if o == nil || IsNil(o.IssueWithoutKyc) {
		var ret bool
		return ret
	}
	return *o.IssueWithoutKyc
}

// GetIssueWithoutKycOk returns a tuple with the IssueWithoutKyc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetIssueWithoutKycOk() (*bool, bool) {
	if o == nil || IsNil(o.IssueWithoutKyc) {
		return nil, false
	}
	return o.IssueWithoutKyc, true
}

// HasIssueWithoutKyc returns a boolean if a field has been set.
func (o *CardProductResponse) HasIssueWithoutKyc() bool {
	if o != nil && !IsNil(o.IssueWithoutKyc) {
		return true
	}

	return false
}

// SetIssueWithoutKyc gets a reference to the given bool and assigns it to the IssueWithoutKyc field.
func (o *CardProductResponse) SetIssueWithoutKyc(v bool) {
	o.IssueWithoutKyc = &v
}

// GetL2l3Enabled returns the L2l3Enabled field value if set, zero value otherwise.
func (o *CardProductResponse) GetL2l3Enabled() bool {
	if o == nil || IsNil(o.L2l3Enabled) {
		var ret bool
		return ret
	}
	return *o.L2l3Enabled
}

// GetL2l3EnabledOk returns a tuple with the L2l3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetL2l3EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.L2l3Enabled) {
		return nil, false
	}
	return o.L2l3Enabled, true
}

// HasL2l3Enabled returns a boolean if a field has been set.
func (o *CardProductResponse) HasL2l3Enabled() bool {
	if o != nil && !IsNil(o.L2l3Enabled) {
		return true
	}

	return false
}

// SetL2l3Enabled gets a reference to the given bool and assigns it to the L2l3Enabled field.
func (o *CardProductResponse) SetL2l3Enabled(v bool) {
	o.L2l3Enabled = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *CardProductResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *CardProductResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetName returns the Name field value
func (o *CardProductResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CardProductResponse) SetName(v string) {
	o.Name = v
}

// GetNotificationLanguage returns the NotificationLanguage field value if set, zero value otherwise.
func (o *CardProductResponse) GetNotificationLanguage() NotificationLanguage {
	if o == nil || IsNil(o.NotificationLanguage) {
		var ret NotificationLanguage
		return ret
	}
	return *o.NotificationLanguage
}

// GetNotificationLanguageOk returns a tuple with the NotificationLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetNotificationLanguageOk() (*NotificationLanguage, bool) {
	if o == nil || IsNil(o.NotificationLanguage) {
		return nil, false
	}
	return o.NotificationLanguage, true
}

// HasNotificationLanguage returns a boolean if a field has been set.
func (o *CardProductResponse) HasNotificationLanguage() bool {
	if o != nil && !IsNil(o.NotificationLanguage) {
		return true
	}

	return false
}

// SetNotificationLanguage gets a reference to the given NotificationLanguage and assigns it to the NotificationLanguage field.
func (o *CardProductResponse) SetNotificationLanguage(v NotificationLanguage) {
	o.NotificationLanguage = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *CardProductResponse) GetOrientation() string {
	if o == nil || IsNil(o.Orientation) {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetOrientationOk() (*string, bool) {
	if o == nil || IsNil(o.Orientation) {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *CardProductResponse) HasOrientation() bool {
	if o != nil && !IsNil(o.Orientation) {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *CardProductResponse) SetOrientation(v string) {
	o.Orientation = &v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *CardProductResponse) GetPackageId() string {
	if o == nil || IsNil(o.PackageId) {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PackageId) {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *CardProductResponse) HasPackageId() bool {
	if o != nil && !IsNil(o.PackageId) {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *CardProductResponse) SetPackageId(v string) {
	o.PackageId = &v
}

// GetPhysicalCardFormat returns the PhysicalCardFormat field value if set, zero value otherwise.
func (o *CardProductResponse) GetPhysicalCardFormat() PhysicalCardFormat {
	if o == nil || IsNil(o.PhysicalCardFormat) {
		var ret PhysicalCardFormat
		return ret
	}
	return *o.PhysicalCardFormat
}

// GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool) {
	if o == nil || IsNil(o.PhysicalCardFormat) {
		return nil, false
	}
	return o.PhysicalCardFormat, true
}

// HasPhysicalCardFormat returns a boolean if a field has been set.
func (o *CardProductResponse) HasPhysicalCardFormat() bool {
	if o != nil && !IsNil(o.PhysicalCardFormat) {
		return true
	}

	return false
}

// SetPhysicalCardFormat gets a reference to the given PhysicalCardFormat and assigns it to the PhysicalCardFormat field.
func (o *CardProductResponse) SetPhysicalCardFormat(v PhysicalCardFormat) {
	o.PhysicalCardFormat = &v
}

// GetReturnAddress returns the ReturnAddress field value if set, zero value otherwise.
func (o *CardProductResponse) GetReturnAddress() Shipping {
	if o == nil || IsNil(o.ReturnAddress) {
		var ret Shipping
		return ret
	}
	return *o.ReturnAddress
}

// GetReturnAddressOk returns a tuple with the ReturnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetReturnAddressOk() (*Shipping, bool) {
	if o == nil || IsNil(o.ReturnAddress) {
		return nil, false
	}
	return o.ReturnAddress, true
}

// HasReturnAddress returns a boolean if a field has been set.
func (o *CardProductResponse) HasReturnAddress() bool {
	if o != nil && !IsNil(o.ReturnAddress) {
		return true
	}

	return false
}

// SetReturnAddress gets a reference to the given Shipping and assigns it to the ReturnAddress field.
func (o *CardProductResponse) SetReturnAddress(v Shipping) {
	o.ReturnAddress = &v
}

// GetStartDate returns the StartDate field value
func (o *CardProductResponse) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *CardProductResponse) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetTxnEnhancer returns the TxnEnhancer field value if set, zero value otherwise.
func (o *CardProductResponse) GetTxnEnhancer() TxnEnhancer {
	if o == nil || IsNil(o.TxnEnhancer) {
		var ret TxnEnhancer
		return ret
	}
	return *o.TxnEnhancer
}

// GetTxnEnhancerOk returns a tuple with the TxnEnhancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetTxnEnhancerOk() (*TxnEnhancer, bool) {
	if o == nil || IsNil(o.TxnEnhancer) {
		return nil, false
	}
	return o.TxnEnhancer, true
}

// HasTxnEnhancer returns a boolean if a field has been set.
func (o *CardProductResponse) HasTxnEnhancer() bool {
	if o != nil && !IsNil(o.TxnEnhancer) {
		return true
	}

	return false
}

// SetTxnEnhancer gets a reference to the given TxnEnhancer and assigns it to the TxnEnhancer field.
func (o *CardProductResponse) SetTxnEnhancer(v TxnEnhancer) {
	o.TxnEnhancer = &v
}

// GetForm returns the Form field value
func (o *CardProductResponse) GetForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *CardProductResponse) SetForm(v string) {
	o.Form = v
}

// GetThreeDsPolicy returns the ThreeDsPolicy field value
func (o *CardProductResponse) GetThreeDsPolicy() ThreeDsPolicy {
	if o == nil {
		var ret ThreeDsPolicy
		return ret
	}

	return o.ThreeDsPolicy
}

// GetThreeDsPolicyOk returns a tuple with the ThreeDsPolicy field value
// and a boolean to check if the value has been set.
func (o *CardProductResponse) GetThreeDsPolicyOk() (*ThreeDsPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreeDsPolicy, true
}

// SetThreeDsPolicy sets field value
func (o *CardProductResponse) SetThreeDsPolicy(v ThreeDsPolicy) {
	o.ThreeDsPolicy = v
}

func (o CardProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	if !IsNil(o.BinCountry) {
		toSerialize["bin_country"] = o.BinCountry
	}
	if !IsNil(o.BypassRiskErrors) {
		toSerialize["bypass_risk_errors"] = o.BypassRiskErrors
	}
	if !IsNil(o.CardBrand) {
		toSerialize["card_brand"] = o.CardBrand
	}
	if !IsNil(o.CardCategory) {
		toSerialize["card_category"] = o.CardCategory
	}
	if !IsNil(o.CardFulfillmentCountry) {
		toSerialize["card_fulfillment_country"] = o.CardFulfillmentCountry
	}
	if !IsNil(o.CardFulfillmentProvider) {
		toSerialize["card_fulfillment_provider"] = o.CardFulfillmentProvider
	}
	toSerialize["card_program_id"] = o.CardProgramId
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	toSerialize["creation_time"] = o.CreationTime
	if !IsNil(o.CrossBorderEnabled) {
		toSerialize["cross_border_enabled"] = o.CrossBorderEnabled
	}
	toSerialize["digital_wallet_tokenization"] = o.DigitalWalletTokenization
	toSerialize["end_date"] = o.EndDate
	toSerialize["id"] = o.Id
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.ImageMode) {
		toSerialize["image_mode"] = o.ImageMode
	}
	if !IsNil(o.IssueWithoutKyc) {
		toSerialize["issue_without_kyc"] = o.IssueWithoutKyc
	}
	if !IsNil(o.L2l3Enabled) {
		toSerialize["l2l3_enabled"] = o.L2l3Enabled
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	toSerialize["name"] = o.Name
	if !IsNil(o.NotificationLanguage) {
		toSerialize["notification_language"] = o.NotificationLanguage
	}
	if !IsNil(o.Orientation) {
		toSerialize["orientation"] = o.Orientation
	}
	if !IsNil(o.PackageId) {
		toSerialize["package_id"] = o.PackageId
	}
	if !IsNil(o.PhysicalCardFormat) {
		toSerialize["physical_card_format"] = o.PhysicalCardFormat
	}
	if !IsNil(o.ReturnAddress) {
		toSerialize["return_address"] = o.ReturnAddress
	}
	toSerialize["start_date"] = o.StartDate
	if !IsNil(o.TxnEnhancer) {
		toSerialize["txn_enhancer"] = o.TxnEnhancer
	}
	toSerialize["form"] = o.Form
	toSerialize["three_ds_policy"] = o.ThreeDsPolicy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardProductResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"card_program_id",
		"creation_time",
		"digital_wallet_tokenization",
		"end_date",
		"id",
		"last_modified_time",
		"name",
		"start_date",
		"form",
		"three_ds_policy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCardProductResponse := _CardProductResponse{}

	err = json.Unmarshal(data, &varCardProductResponse)

	if err != nil {
		return err
	}

	*o = CardProductResponse(varCardProductResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "bin_country")
		delete(additionalProperties, "bypass_risk_errors")
		delete(additionalProperties, "card_brand")
		delete(additionalProperties, "card_category")
		delete(additionalProperties, "card_fulfillment_country")
		delete(additionalProperties, "card_fulfillment_provider")
		delete(additionalProperties, "card_program_id")
		delete(additionalProperties, "card_type")
		delete(additionalProperties, "color")
		delete(additionalProperties, "creation_time")
		delete(additionalProperties, "cross_border_enabled")
		delete(additionalProperties, "digital_wallet_tokenization")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "id")
		delete(additionalProperties, "image")
		delete(additionalProperties, "image_mode")
		delete(additionalProperties, "issue_without_kyc")
		delete(additionalProperties, "l2l3_enabled")
		delete(additionalProperties, "last_modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "notification_language")
		delete(additionalProperties, "orientation")
		delete(additionalProperties, "package_id")
		delete(additionalProperties, "physical_card_format")
		delete(additionalProperties, "return_address")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "txn_enhancer")
		delete(additionalProperties, "form")
		delete(additionalProperties, "three_ds_policy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardProductResponse struct {
	value *CardProductResponse
	isSet bool
}

func (v NullableCardProductResponse) Get() *CardProductResponse {
	return v.value
}

func (v *NullableCardProductResponse) Set(val *CardProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProductResponse(val *CardProductResponse) *NullableCardProductResponse {
	return &NullableCardProductResponse{value: val, isSet: true}
}

func (v NullableCardProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
