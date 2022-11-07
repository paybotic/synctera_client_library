/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CardProduct Card Product
type CardProduct struct {
	// PHYSICAL or VIRTUAL.
	Form string `json:"form"`
	// Indicates whether the Card Product is active
	Active bool `json:"active"`
	// Card Program ID
	CardProgramId string `json:"card_program_id"`
	// Color code for dynamic card elements such as PAN and card holder name
	Color *string `json:"color,omitempty"`
	// The timestamp representing when the Card Product was created
	CreationTime *time.Time `json:"creation_time,omitempty"`
	DigitalWalletTokenization *DigitalWalletTokenization `json:"digital_wallet_tokenization,omitempty"`
	// The time when the Card Product is decommissioned
	EndDate *time.Time `json:"end_date,omitempty"`
	// Card Product ID
	Id *string `json:"id,omitempty"`
	// Indicates whether or not there is an overlay image of the card product available
	Image *bool `json:"image,omitempty"`
	ImageMode *CardImageMode `json:"image_mode,omitempty"`
	// Allow issuing cards on this product without requiring KYC
	IssueWithoutKyc *bool `json:"issue_without_kyc,omitempty"`
	// The timestamp representing when the Card Product was last modified
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// The name of the Card Product
	Name string `json:"name"`
	// Card orientation
	Orientation *string `json:"orientation,omitempty"`
	// Card fulfillment provider’s package ID
	PackageId *string `json:"package_id,omitempty"`
	PhysicalCardFormat *PhysicalCardFormat `json:"physical_card_format,omitempty"`
	ReturnAddress *Shipping `json:"return_address,omitempty"`
	// The time when the Card Product goes live
	StartDate time.Time `json:"start_date"`
	TxnEnhancer *TxnEnhancer `json:"txn_enhancer,omitempty"`
}

// NewCardProduct instantiates a new CardProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProduct(form string, active bool, cardProgramId string, name string, startDate time.Time) *CardProduct {
	this := CardProduct{}
	this.Form = form
	this.Active = active
	this.CardProgramId = cardProgramId
	this.Name = name
	this.StartDate = startDate
	var txnEnhancer TxnEnhancer = TXNENHANCER_NONE
	this.TxnEnhancer = &txnEnhancer
	return &this
}

// NewCardProductWithDefaults instantiates a new CardProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProductWithDefaults() *CardProduct {
	this := CardProduct{}
	var txnEnhancer TxnEnhancer = TXNENHANCER_NONE
	this.TxnEnhancer = &txnEnhancer
	return &this
}

// GetForm returns the Form field value
func (o *CardProduct) GetForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *CardProduct) GetFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *CardProduct) SetForm(v string) {
	o.Form = v
}

// GetActive returns the Active field value
func (o *CardProduct) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *CardProduct) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *CardProduct) SetActive(v bool) {
	o.Active = v
}

// GetCardProgramId returns the CardProgramId field value
func (o *CardProduct) GetCardProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardProgramId
}

// GetCardProgramIdOk returns a tuple with the CardProgramId field value
// and a boolean to check if the value has been set.
func (o *CardProduct) GetCardProgramIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProgramId, true
}

// SetCardProgramId sets field value
func (o *CardProduct) SetCardProgramId(v string) {
	o.CardProgramId = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *CardProduct) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *CardProduct) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *CardProduct) SetColor(v string) {
	o.Color = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *CardProduct) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *CardProduct) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *CardProduct) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDigitalWalletTokenization returns the DigitalWalletTokenization field value if set, zero value otherwise.
func (o *CardProduct) GetDigitalWalletTokenization() DigitalWalletTokenization {
	if o == nil || o.DigitalWalletTokenization == nil {
		var ret DigitalWalletTokenization
		return ret
	}
	return *o.DigitalWalletTokenization
}

// GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool) {
	if o == nil || o.DigitalWalletTokenization == nil {
		return nil, false
	}
	return o.DigitalWalletTokenization, true
}

// HasDigitalWalletTokenization returns a boolean if a field has been set.
func (o *CardProduct) HasDigitalWalletTokenization() bool {
	if o != nil && o.DigitalWalletTokenization != nil {
		return true
	}

	return false
}

// SetDigitalWalletTokenization gets a reference to the given DigitalWalletTokenization and assigns it to the DigitalWalletTokenization field.
func (o *CardProduct) SetDigitalWalletTokenization(v DigitalWalletTokenization) {
	o.DigitalWalletTokenization = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CardProduct) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CardProduct) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *CardProduct) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CardProduct) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CardProduct) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CardProduct) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CardProduct) GetImage() bool {
	if o == nil || o.Image == nil {
		var ret bool
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetImageOk() (*bool, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CardProduct) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given bool and assigns it to the Image field.
func (o *CardProduct) SetImage(v bool) {
	o.Image = &v
}

// GetImageMode returns the ImageMode field value if set, zero value otherwise.
func (o *CardProduct) GetImageMode() CardImageMode {
	if o == nil || o.ImageMode == nil {
		var ret CardImageMode
		return ret
	}
	return *o.ImageMode
}

// GetImageModeOk returns a tuple with the ImageMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetImageModeOk() (*CardImageMode, bool) {
	if o == nil || o.ImageMode == nil {
		return nil, false
	}
	return o.ImageMode, true
}

// HasImageMode returns a boolean if a field has been set.
func (o *CardProduct) HasImageMode() bool {
	if o != nil && o.ImageMode != nil {
		return true
	}

	return false
}

// SetImageMode gets a reference to the given CardImageMode and assigns it to the ImageMode field.
func (o *CardProduct) SetImageMode(v CardImageMode) {
	o.ImageMode = &v
}

// GetIssueWithoutKyc returns the IssueWithoutKyc field value if set, zero value otherwise.
func (o *CardProduct) GetIssueWithoutKyc() bool {
	if o == nil || o.IssueWithoutKyc == nil {
		var ret bool
		return ret
	}
	return *o.IssueWithoutKyc
}

// GetIssueWithoutKycOk returns a tuple with the IssueWithoutKyc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetIssueWithoutKycOk() (*bool, bool) {
	if o == nil || o.IssueWithoutKyc == nil {
		return nil, false
	}
	return o.IssueWithoutKyc, true
}

// HasIssueWithoutKyc returns a boolean if a field has been set.
func (o *CardProduct) HasIssueWithoutKyc() bool {
	if o != nil && o.IssueWithoutKyc != nil {
		return true
	}

	return false
}

// SetIssueWithoutKyc gets a reference to the given bool and assigns it to the IssueWithoutKyc field.
func (o *CardProduct) SetIssueWithoutKyc(v bool) {
	o.IssueWithoutKyc = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *CardProduct) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *CardProduct) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *CardProduct) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetName returns the Name field value
func (o *CardProduct) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CardProduct) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CardProduct) SetName(v string) {
	o.Name = v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *CardProduct) GetOrientation() string {
	if o == nil || o.Orientation == nil {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetOrientationOk() (*string, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *CardProduct) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *CardProduct) SetOrientation(v string) {
	o.Orientation = &v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *CardProduct) GetPackageId() string {
	if o == nil || o.PackageId == nil {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetPackageIdOk() (*string, bool) {
	if o == nil || o.PackageId == nil {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *CardProduct) HasPackageId() bool {
	if o != nil && o.PackageId != nil {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *CardProduct) SetPackageId(v string) {
	o.PackageId = &v
}

// GetPhysicalCardFormat returns the PhysicalCardFormat field value if set, zero value otherwise.
func (o *CardProduct) GetPhysicalCardFormat() PhysicalCardFormat {
	if o == nil || o.PhysicalCardFormat == nil {
		var ret PhysicalCardFormat
		return ret
	}
	return *o.PhysicalCardFormat
}

// GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool) {
	if o == nil || o.PhysicalCardFormat == nil {
		return nil, false
	}
	return o.PhysicalCardFormat, true
}

// HasPhysicalCardFormat returns a boolean if a field has been set.
func (o *CardProduct) HasPhysicalCardFormat() bool {
	if o != nil && o.PhysicalCardFormat != nil {
		return true
	}

	return false
}

// SetPhysicalCardFormat gets a reference to the given PhysicalCardFormat and assigns it to the PhysicalCardFormat field.
func (o *CardProduct) SetPhysicalCardFormat(v PhysicalCardFormat) {
	o.PhysicalCardFormat = &v
}

// GetReturnAddress returns the ReturnAddress field value if set, zero value otherwise.
func (o *CardProduct) GetReturnAddress() Shipping {
	if o == nil || o.ReturnAddress == nil {
		var ret Shipping
		return ret
	}
	return *o.ReturnAddress
}

// GetReturnAddressOk returns a tuple with the ReturnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetReturnAddressOk() (*Shipping, bool) {
	if o == nil || o.ReturnAddress == nil {
		return nil, false
	}
	return o.ReturnAddress, true
}

// HasReturnAddress returns a boolean if a field has been set.
func (o *CardProduct) HasReturnAddress() bool {
	if o != nil && o.ReturnAddress != nil {
		return true
	}

	return false
}

// SetReturnAddress gets a reference to the given Shipping and assigns it to the ReturnAddress field.
func (o *CardProduct) SetReturnAddress(v Shipping) {
	o.ReturnAddress = &v
}

// GetStartDate returns the StartDate field value
func (o *CardProduct) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CardProduct) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *CardProduct) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetTxnEnhancer returns the TxnEnhancer field value if set, zero value otherwise.
func (o *CardProduct) GetTxnEnhancer() TxnEnhancer {
	if o == nil || o.TxnEnhancer == nil {
		var ret TxnEnhancer
		return ret
	}
	return *o.TxnEnhancer
}

// GetTxnEnhancerOk returns a tuple with the TxnEnhancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProduct) GetTxnEnhancerOk() (*TxnEnhancer, bool) {
	if o == nil || o.TxnEnhancer == nil {
		return nil, false
	}
	return o.TxnEnhancer, true
}

// HasTxnEnhancer returns a boolean if a field has been set.
func (o *CardProduct) HasTxnEnhancer() bool {
	if o != nil && o.TxnEnhancer != nil {
		return true
	}

	return false
}

// SetTxnEnhancer gets a reference to the given TxnEnhancer and assigns it to the TxnEnhancer field.
func (o *CardProduct) SetTxnEnhancer(v TxnEnhancer) {
	o.TxnEnhancer = &v
}

func (o CardProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["form"] = o.Form
	}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["card_program_id"] = o.CardProgramId
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.DigitalWalletTokenization != nil {
		toSerialize["digital_wallet_tokenization"] = o.DigitalWalletTokenization
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.ImageMode != nil {
		toSerialize["image_mode"] = o.ImageMode
	}
	if o.IssueWithoutKyc != nil {
		toSerialize["issue_without_kyc"] = o.IssueWithoutKyc
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Orientation != nil {
		toSerialize["orientation"] = o.Orientation
	}
	if o.PackageId != nil {
		toSerialize["package_id"] = o.PackageId
	}
	if o.PhysicalCardFormat != nil {
		toSerialize["physical_card_format"] = o.PhysicalCardFormat
	}
	if o.ReturnAddress != nil {
		toSerialize["return_address"] = o.ReturnAddress
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if o.TxnEnhancer != nil {
		toSerialize["txn_enhancer"] = o.TxnEnhancer
	}
	return json.Marshal(toSerialize)
}

type NullableCardProduct struct {
	value *CardProduct
	isSet bool
}

func (v NullableCardProduct) Get() *CardProduct {
	return v.value
}

func (v *NullableCardProduct) Set(val *CardProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProduct(val *CardProduct) *NullableCardProduct {
	return &NullableCardProduct{value: val, isSet: true}
}

func (v NullableCardProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


