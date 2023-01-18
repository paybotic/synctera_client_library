/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.32.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// BinResponse struct for BinResponse
type BinResponse struct {
	// Account range length
	AccountRangeLength int32 `json:"account_range_length"`
	// The bank ID
	BankId int32 `json:"bank_id"`
	// The ICA to which fees will be billed
	BillingIca string `json:"billing_ica"`
	// The bin number
	Bin string `json:"bin"`
	BinStatus BinStatus `json:"bin_status"`
	// The Mastercard or Visa Product Code - 3 alpha-numeric characters
	BrandProductCode string `json:"brand_product_code"`
	CardBrand CardBrand `json:"card_brand"`
	CardCategory CardCategory `json:"card_category"`
	CardProductType CardProductType `json:"card_product_type"`
	// ISO-3166-1 Alpha-2 country code
	Country string `json:"country"`
	// The timestamp representing when the bin was created
	CreationTime time.Time `json:"creation_time"`
	// ISO 4217  Alpha-3 currency code
	Currency string `json:"currency"`
	// Determines if bin supports digital wallet tokenization
	DigitalWalletActive *bool `json:"digital_wallet_active,omitempty"`
	// The time when bin is decommissioned
	EndDate *time.Time `json:"end_date,omitempty"`
	// ICA/BID
	IcaBid string `json:"ica_bid"`
	// Bin ID
	Id string `json:"id"`
	// Controls whether bin allows tokenization
	IsTokenizationEnabled bool `json:"is_tokenization_enabled"`
	// The timestamp representing when the bin was last modified
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Pan utilization
	PanUtilization int32 `json:"pan_utilization"`
	// The partner ID
	PartnerId int32 `json:"partner_id"`
	PhysicalCardFormat *PhysicalCardFormat `json:"physical_card_format,omitempty"`
	// The name of the card processor
	Processor string `json:"processor"`
	// The time when bin goes live
	StartDate *time.Time `json:"start_date,omitempty"`
}

// NewBinResponse instantiates a new BinResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinResponse(accountRangeLength int32, bankId int32, billingIca string, bin string, binStatus BinStatus, brandProductCode string, cardBrand CardBrand, cardCategory CardCategory, cardProductType CardProductType, country string, creationTime time.Time, currency string, icaBid string, id string, isTokenizationEnabled bool, lastModifiedTime time.Time, panUtilization int32, partnerId int32, processor string) *BinResponse {
	this := BinResponse{}
	this.AccountRangeLength = accountRangeLength
	this.BankId = bankId
	this.BillingIca = billingIca
	this.Bin = bin
	this.BinStatus = binStatus
	this.BrandProductCode = brandProductCode
	this.CardBrand = cardBrand
	this.CardCategory = cardCategory
	this.CardProductType = cardProductType
	this.Country = country
	this.CreationTime = creationTime
	this.Currency = currency
	var digitalWalletActive bool = false
	this.DigitalWalletActive = &digitalWalletActive
	this.IcaBid = icaBid
	this.Id = id
	this.IsTokenizationEnabled = isTokenizationEnabled
	this.LastModifiedTime = lastModifiedTime
	this.PanUtilization = panUtilization
	this.PartnerId = partnerId
	this.Processor = processor
	return &this
}

// NewBinResponseWithDefaults instantiates a new BinResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinResponseWithDefaults() *BinResponse {
	this := BinResponse{}
	var digitalWalletActive bool = false
	this.DigitalWalletActive = &digitalWalletActive
	var isTokenizationEnabled bool = false
	this.IsTokenizationEnabled = isTokenizationEnabled
	return &this
}

// GetAccountRangeLength returns the AccountRangeLength field value
func (o *BinResponse) GetAccountRangeLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountRangeLength
}

// GetAccountRangeLengthOk returns a tuple with the AccountRangeLength field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetAccountRangeLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountRangeLength, true
}

// SetAccountRangeLength sets field value
func (o *BinResponse) SetAccountRangeLength(v int32) {
	o.AccountRangeLength = v
}

// GetBankId returns the BankId field value
func (o *BinResponse) GetBankId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetBankIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankId, true
}

// SetBankId sets field value
func (o *BinResponse) SetBankId(v int32) {
	o.BankId = v
}

// GetBillingIca returns the BillingIca field value
func (o *BinResponse) GetBillingIca() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingIca
}

// GetBillingIcaOk returns a tuple with the BillingIca field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetBillingIcaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingIca, true
}

// SetBillingIca sets field value
func (o *BinResponse) SetBillingIca(v string) {
	o.BillingIca = v
}

// GetBin returns the Bin field value
func (o *BinResponse) GetBin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bin
}

// GetBinOk returns a tuple with the Bin field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetBinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bin, true
}

// SetBin sets field value
func (o *BinResponse) SetBin(v string) {
	o.Bin = v
}

// GetBinStatus returns the BinStatus field value
func (o *BinResponse) GetBinStatus() BinStatus {
	if o == nil {
		var ret BinStatus
		return ret
	}

	return o.BinStatus
}

// GetBinStatusOk returns a tuple with the BinStatus field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetBinStatusOk() (*BinStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinStatus, true
}

// SetBinStatus sets field value
func (o *BinResponse) SetBinStatus(v BinStatus) {
	o.BinStatus = v
}

// GetBrandProductCode returns the BrandProductCode field value
func (o *BinResponse) GetBrandProductCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandProductCode
}

// GetBrandProductCodeOk returns a tuple with the BrandProductCode field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetBrandProductCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandProductCode, true
}

// SetBrandProductCode sets field value
func (o *BinResponse) SetBrandProductCode(v string) {
	o.BrandProductCode = v
}

// GetCardBrand returns the CardBrand field value
func (o *BinResponse) GetCardBrand() CardBrand {
	if o == nil {
		var ret CardBrand
		return ret
	}

	return o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetCardBrandOk() (*CardBrand, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardBrand, true
}

// SetCardBrand sets field value
func (o *BinResponse) SetCardBrand(v CardBrand) {
	o.CardBrand = v
}

// GetCardCategory returns the CardCategory field value
func (o *BinResponse) GetCardCategory() CardCategory {
	if o == nil {
		var ret CardCategory
		return ret
	}

	return o.CardCategory
}

// GetCardCategoryOk returns a tuple with the CardCategory field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetCardCategoryOk() (*CardCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardCategory, true
}

// SetCardCategory sets field value
func (o *BinResponse) SetCardCategory(v CardCategory) {
	o.CardCategory = v
}

// GetCardProductType returns the CardProductType field value
func (o *BinResponse) GetCardProductType() CardProductType {
	if o == nil {
		var ret CardProductType
		return ret
	}

	return o.CardProductType
}

// GetCardProductTypeOk returns a tuple with the CardProductType field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetCardProductTypeOk() (*CardProductType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProductType, true
}

// SetCardProductType sets field value
func (o *BinResponse) SetCardProductType(v CardProductType) {
	o.CardProductType = v
}

// GetCountry returns the Country field value
func (o *BinResponse) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *BinResponse) SetCountry(v string) {
	o.Country = v
}

// GetCreationTime returns the CreationTime field value
func (o *BinResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *BinResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCurrency returns the Currency field value
func (o *BinResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *BinResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetDigitalWalletActive returns the DigitalWalletActive field value if set, zero value otherwise.
func (o *BinResponse) GetDigitalWalletActive() bool {
	if o == nil || o.DigitalWalletActive == nil {
		var ret bool
		return ret
	}
	return *o.DigitalWalletActive
}

// GetDigitalWalletActiveOk returns a tuple with the DigitalWalletActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinResponse) GetDigitalWalletActiveOk() (*bool, bool) {
	if o == nil || o.DigitalWalletActive == nil {
		return nil, false
	}
	return o.DigitalWalletActive, true
}

// HasDigitalWalletActive returns a boolean if a field has been set.
func (o *BinResponse) HasDigitalWalletActive() bool {
	if o != nil && o.DigitalWalletActive != nil {
		return true
	}

	return false
}

// SetDigitalWalletActive gets a reference to the given bool and assigns it to the DigitalWalletActive field.
func (o *BinResponse) SetDigitalWalletActive(v bool) {
	o.DigitalWalletActive = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BinResponse) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BinResponse) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BinResponse) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetIcaBid returns the IcaBid field value
func (o *BinResponse) GetIcaBid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IcaBid
}

// GetIcaBidOk returns a tuple with the IcaBid field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetIcaBidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IcaBid, true
}

// SetIcaBid sets field value
func (o *BinResponse) SetIcaBid(v string) {
	o.IcaBid = v
}

// GetId returns the Id field value
func (o *BinResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BinResponse) SetId(v string) {
	o.Id = v
}

// GetIsTokenizationEnabled returns the IsTokenizationEnabled field value
func (o *BinResponse) GetIsTokenizationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTokenizationEnabled
}

// GetIsTokenizationEnabledOk returns a tuple with the IsTokenizationEnabled field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetIsTokenizationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTokenizationEnabled, true
}

// SetIsTokenizationEnabled sets field value
func (o *BinResponse) SetIsTokenizationEnabled(v bool) {
	o.IsTokenizationEnabled = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *BinResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *BinResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetPanUtilization returns the PanUtilization field value
func (o *BinResponse) GetPanUtilization() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PanUtilization
}

// GetPanUtilizationOk returns a tuple with the PanUtilization field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetPanUtilizationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PanUtilization, true
}

// SetPanUtilization sets field value
func (o *BinResponse) SetPanUtilization(v int32) {
	o.PanUtilization = v
}

// GetPartnerId returns the PartnerId field value
func (o *BinResponse) GetPartnerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetPartnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerId, true
}

// SetPartnerId sets field value
func (o *BinResponse) SetPartnerId(v int32) {
	o.PartnerId = v
}

// GetPhysicalCardFormat returns the PhysicalCardFormat field value if set, zero value otherwise.
func (o *BinResponse) GetPhysicalCardFormat() PhysicalCardFormat {
	if o == nil || o.PhysicalCardFormat == nil {
		var ret PhysicalCardFormat
		return ret
	}
	return *o.PhysicalCardFormat
}

// GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinResponse) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool) {
	if o == nil || o.PhysicalCardFormat == nil {
		return nil, false
	}
	return o.PhysicalCardFormat, true
}

// HasPhysicalCardFormat returns a boolean if a field has been set.
func (o *BinResponse) HasPhysicalCardFormat() bool {
	if o != nil && o.PhysicalCardFormat != nil {
		return true
	}

	return false
}

// SetPhysicalCardFormat gets a reference to the given PhysicalCardFormat and assigns it to the PhysicalCardFormat field.
func (o *BinResponse) SetPhysicalCardFormat(v PhysicalCardFormat) {
	o.PhysicalCardFormat = &v
}

// GetProcessor returns the Processor field value
func (o *BinResponse) GetProcessor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value
// and a boolean to check if the value has been set.
func (o *BinResponse) GetProcessorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Processor, true
}

// SetProcessor sets field value
func (o *BinResponse) SetProcessor(v string) {
	o.Processor = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BinResponse) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BinResponse) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BinResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o BinResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_range_length"] = o.AccountRangeLength
	}
	if true {
		toSerialize["bank_id"] = o.BankId
	}
	if true {
		toSerialize["billing_ica"] = o.BillingIca
	}
	if true {
		toSerialize["bin"] = o.Bin
	}
	if true {
		toSerialize["bin_status"] = o.BinStatus
	}
	if true {
		toSerialize["brand_product_code"] = o.BrandProductCode
	}
	if true {
		toSerialize["card_brand"] = o.CardBrand
	}
	if true {
		toSerialize["card_category"] = o.CardCategory
	}
	if true {
		toSerialize["card_product_type"] = o.CardProductType
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.DigitalWalletActive != nil {
		toSerialize["digital_wallet_active"] = o.DigitalWalletActive
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["ica_bid"] = o.IcaBid
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["is_tokenization_enabled"] = o.IsTokenizationEnabled
	}
	if true {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["pan_utilization"] = o.PanUtilization
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId
	}
	if o.PhysicalCardFormat != nil {
		toSerialize["physical_card_format"] = o.PhysicalCardFormat
	}
	if true {
		toSerialize["processor"] = o.Processor
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableBinResponse struct {
	value *BinResponse
	isSet bool
}

func (v NullableBinResponse) Get() *BinResponse {
	return v.value
}

func (v *NullableBinResponse) Set(val *BinResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBinResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBinResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinResponse(val *BinResponse) *NullableBinResponse {
	return &NullableBinResponse{value: val, isSet: true}
}

func (v NullableBinResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


