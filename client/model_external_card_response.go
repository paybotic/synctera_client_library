/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.153.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the ExternalCardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalCardResponse{}

// ExternalCardResponse struct for ExternalCardResponse
type ExternalCardResponse struct {
	// Bank Identification Number
	Bin         *string    `json:"bin,omitempty"`
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// ISO 4217  Alpha-3 currency code
	Currency string `json:"currency"`
	// The unique identifier of a customer
	CustomerId      string `json:"customer_id"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
	// External card ID
	Id string `json:"id"`
	// Name of issuing financial institution
	Issuer *string `json:"issuer,omitempty"`
	// The last 4 digits of the card PAN
	LastFour         string     `json:"last_four"`
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// The cardholder name
	Name          string                     `json:"name"`
	Status        ExternalCardStatus         `json:"status"`
	Verifications *ExternalCardVerifications `json:"verifications,omitempty"`
}

type _ExternalCardResponse ExternalCardResponse

// NewExternalCardResponse instantiates a new ExternalCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalCardResponse(currency string, customerId string, expirationMonth string, expirationYear string, id string, lastFour string, name string, status ExternalCardStatus) *ExternalCardResponse {
	this := ExternalCardResponse{}
	this.Currency = currency
	this.CustomerId = customerId
	this.ExpirationMonth = expirationMonth
	this.ExpirationYear = expirationYear
	this.Id = id
	this.LastFour = lastFour
	this.Name = name
	this.Status = status
	return &this
}

// NewExternalCardResponseWithDefaults instantiates a new ExternalCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalCardResponseWithDefaults() *ExternalCardResponse {
	this := ExternalCardResponse{}
	return &this
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *ExternalCardResponse) GetBin() string {
	if o == nil || IsNil(o.Bin) {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetBinOk() (*string, bool) {
	if o == nil || IsNil(o.Bin) {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *ExternalCardResponse) HasBin() bool {
	if o != nil && !IsNil(o.Bin) {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *ExternalCardResponse) SetBin(v string) {
	o.Bin = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ExternalCardResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ExternalCardResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *ExternalCardResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCurrency returns the Currency field value
func (o *ExternalCardResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ExternalCardResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value
func (o *ExternalCardResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ExternalCardResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetExpirationMonth returns the ExpirationMonth field value
func (o *ExternalCardResponse) GetExpirationMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetExpirationMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationMonth, true
}

// SetExpirationMonth sets field value
func (o *ExternalCardResponse) SetExpirationMonth(v string) {
	o.ExpirationMonth = v
}

// GetExpirationYear returns the ExpirationYear field value
func (o *ExternalCardResponse) GetExpirationYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetExpirationYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationYear, true
}

// SetExpirationYear sets field value
func (o *ExternalCardResponse) SetExpirationYear(v string) {
	o.ExpirationYear = v
}

// GetId returns the Id field value
func (o *ExternalCardResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalCardResponse) SetId(v string) {
	o.Id = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *ExternalCardResponse) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *ExternalCardResponse) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *ExternalCardResponse) SetIssuer(v string) {
	o.Issuer = &v
}

// GetLastFour returns the LastFour field value
func (o *ExternalCardResponse) GetLastFour() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetLastFourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastFour, true
}

// SetLastFour sets field value
func (o *ExternalCardResponse) SetLastFour(v string) {
	o.LastFour = v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *ExternalCardResponse) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *ExternalCardResponse) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *ExternalCardResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetName returns the Name field value
func (o *ExternalCardResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalCardResponse) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *ExternalCardResponse) GetStatus() ExternalCardStatus {
	if o == nil {
		var ret ExternalCardStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetStatusOk() (*ExternalCardStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExternalCardResponse) SetStatus(v ExternalCardStatus) {
	o.Status = v
}

// GetVerifications returns the Verifications field value if set, zero value otherwise.
func (o *ExternalCardResponse) GetVerifications() ExternalCardVerifications {
	if o == nil || IsNil(o.Verifications) {
		var ret ExternalCardVerifications
		return ret
	}
	return *o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalCardResponse) GetVerificationsOk() (*ExternalCardVerifications, bool) {
	if o == nil || IsNil(o.Verifications) {
		return nil, false
	}
	return o.Verifications, true
}

// HasVerifications returns a boolean if a field has been set.
func (o *ExternalCardResponse) HasVerifications() bool {
	if o != nil && !IsNil(o.Verifications) {
		return true
	}

	return false
}

// SetVerifications gets a reference to the given ExternalCardVerifications and assigns it to the Verifications field.
func (o *ExternalCardResponse) SetVerifications(v ExternalCardVerifications) {
	o.Verifications = &v
}

func (o ExternalCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalCardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bin) {
		toSerialize["bin"] = o.Bin
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	toSerialize["currency"] = o.Currency
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["expiration_month"] = o.ExpirationMonth
	toSerialize["expiration_year"] = o.ExpirationYear
	toSerialize["id"] = o.Id
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	toSerialize["last_four"] = o.LastFour
	if !IsNil(o.LastModifiedTime) {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	if !IsNil(o.Verifications) {
		toSerialize["verifications"] = o.Verifications
	}
	return toSerialize, nil
}

func (o *ExternalCardResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"customer_id",
		"expiration_month",
		"expiration_year",
		"id",
		"last_four",
		"name",
		"status",
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

	varExternalCardResponse := _ExternalCardResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalCardResponse)

	if err != nil {
		return err
	}

	*o = ExternalCardResponse(varExternalCardResponse)

	return err
}

type NullableExternalCardResponse struct {
	value *ExternalCardResponse
	isSet bool
}

func (v NullableExternalCardResponse) Get() *ExternalCardResponse {
	return v.value
}

func (v *NullableExternalCardResponse) Set(val *ExternalCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalCardResponse(val *ExternalCardResponse) *NullableExternalCardResponse {
	return &NullableExternalCardResponse{value: val, isSet: true}
}

func (v NullableExternalCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
