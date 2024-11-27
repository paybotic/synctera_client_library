/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.140.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Business1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Business1{}

// Business1 struct for Business1
type Business1 struct {
	CreationTime        *time.Time `json:"creation_time,omitempty"`
	Ein                 *string    `json:"ein,omitempty"`
	Email               *string    `json:"email,omitempty"`
	EntityName          *string    `json:"entity_name,omitempty"`
	FormationDate       *string    `json:"formation_date,omitempty"`
	FormationState      *string    `json:"formation_state,omitempty"`
	Id                  *string    `json:"id,omitempty"`
	LastUpdatedTime     *time.Time `json:"last_updated_time,omitempty"`
	PhoneNumber         *string    `json:"phone_number,omitempty"`
	Status              *string    `json:"status,omitempty"`
	Structure           *string    `json:"structure,omitempty"`
	TradeNames          []string   `json:"trade_names,omitempty"`
	VerificationLastRun *time.Time `json:"verification_last_run,omitempty"`
	VerificationStatus  *string    `json:"verification_status,omitempty"`
}

// NewBusiness1 instantiates a new Business1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusiness1() *Business1 {
	this := Business1{}
	return &this
}

// NewBusiness1WithDefaults instantiates a new Business1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusiness1WithDefaults() *Business1 {
	this := Business1{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Business1) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Business1) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Business1) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEin returns the Ein field value if set, zero value otherwise.
func (o *Business1) GetEin() string {
	if o == nil || IsNil(o.Ein) {
		var ret string
		return ret
	}
	return *o.Ein
}

// GetEinOk returns a tuple with the Ein field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetEinOk() (*string, bool) {
	if o == nil || IsNil(o.Ein) {
		return nil, false
	}
	return o.Ein, true
}

// HasEin returns a boolean if a field has been set.
func (o *Business1) HasEin() bool {
	if o != nil && !IsNil(o.Ein) {
		return true
	}

	return false
}

// SetEin gets a reference to the given string and assigns it to the Ein field.
func (o *Business1) SetEin(v string) {
	o.Ein = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Business1) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Business1) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Business1) SetEmail(v string) {
	o.Email = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *Business1) GetEntityName() string {
	if o == nil || IsNil(o.EntityName) {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetEntityNameOk() (*string, bool) {
	if o == nil || IsNil(o.EntityName) {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *Business1) HasEntityName() bool {
	if o != nil && !IsNil(o.EntityName) {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *Business1) SetEntityName(v string) {
	o.EntityName = &v
}

// GetFormationDate returns the FormationDate field value if set, zero value otherwise.
func (o *Business1) GetFormationDate() string {
	if o == nil || IsNil(o.FormationDate) {
		var ret string
		return ret
	}
	return *o.FormationDate
}

// GetFormationDateOk returns a tuple with the FormationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetFormationDateOk() (*string, bool) {
	if o == nil || IsNil(o.FormationDate) {
		return nil, false
	}
	return o.FormationDate, true
}

// HasFormationDate returns a boolean if a field has been set.
func (o *Business1) HasFormationDate() bool {
	if o != nil && !IsNil(o.FormationDate) {
		return true
	}

	return false
}

// SetFormationDate gets a reference to the given string and assigns it to the FormationDate field.
func (o *Business1) SetFormationDate(v string) {
	o.FormationDate = &v
}

// GetFormationState returns the FormationState field value if set, zero value otherwise.
func (o *Business1) GetFormationState() string {
	if o == nil || IsNil(o.FormationState) {
		var ret string
		return ret
	}
	return *o.FormationState
}

// GetFormationStateOk returns a tuple with the FormationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetFormationStateOk() (*string, bool) {
	if o == nil || IsNil(o.FormationState) {
		return nil, false
	}
	return o.FormationState, true
}

// HasFormationState returns a boolean if a field has been set.
func (o *Business1) HasFormationState() bool {
	if o != nil && !IsNil(o.FormationState) {
		return true
	}

	return false
}

// SetFormationState gets a reference to the given string and assigns it to the FormationState field.
func (o *Business1) SetFormationState(v string) {
	o.FormationState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Business1) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Business1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Business1) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Business1) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Business1) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Business1) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Business1) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Business1) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Business1) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Business1) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Business1) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Business1) SetStatus(v string) {
	o.Status = &v
}

// GetStructure returns the Structure field value if set, zero value otherwise.
func (o *Business1) GetStructure() string {
	if o == nil || IsNil(o.Structure) {
		var ret string
		return ret
	}
	return *o.Structure
}

// GetStructureOk returns a tuple with the Structure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetStructureOk() (*string, bool) {
	if o == nil || IsNil(o.Structure) {
		return nil, false
	}
	return o.Structure, true
}

// HasStructure returns a boolean if a field has been set.
func (o *Business1) HasStructure() bool {
	if o != nil && !IsNil(o.Structure) {
		return true
	}

	return false
}

// SetStructure gets a reference to the given string and assigns it to the Structure field.
func (o *Business1) SetStructure(v string) {
	o.Structure = &v
}

// GetTradeNames returns the TradeNames field value if set, zero value otherwise.
func (o *Business1) GetTradeNames() []string {
	if o == nil || IsNil(o.TradeNames) {
		var ret []string
		return ret
	}
	return o.TradeNames
}

// GetTradeNamesOk returns a tuple with the TradeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetTradeNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.TradeNames) {
		return nil, false
	}
	return o.TradeNames, true
}

// HasTradeNames returns a boolean if a field has been set.
func (o *Business1) HasTradeNames() bool {
	if o != nil && !IsNil(o.TradeNames) {
		return true
	}

	return false
}

// SetTradeNames gets a reference to the given []string and assigns it to the TradeNames field.
func (o *Business1) SetTradeNames(v []string) {
	o.TradeNames = v
}

// GetVerificationLastRun returns the VerificationLastRun field value if set, zero value otherwise.
func (o *Business1) GetVerificationLastRun() time.Time {
	if o == nil || IsNil(o.VerificationLastRun) {
		var ret time.Time
		return ret
	}
	return *o.VerificationLastRun
}

// GetVerificationLastRunOk returns a tuple with the VerificationLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetVerificationLastRunOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VerificationLastRun) {
		return nil, false
	}
	return o.VerificationLastRun, true
}

// HasVerificationLastRun returns a boolean if a field has been set.
func (o *Business1) HasVerificationLastRun() bool {
	if o != nil && !IsNil(o.VerificationLastRun) {
		return true
	}

	return false
}

// SetVerificationLastRun gets a reference to the given time.Time and assigns it to the VerificationLastRun field.
func (o *Business1) SetVerificationLastRun(v time.Time) {
	o.VerificationLastRun = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *Business1) GetVerificationStatus() string {
	if o == nil || IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business1) GetVerificationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *Business1) HasVerificationStatus() bool {
	if o != nil && !IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *Business1) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o Business1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Business1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.Ein) {
		toSerialize["ein"] = o.Ein
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EntityName) {
		toSerialize["entity_name"] = o.EntityName
	}
	if !IsNil(o.FormationDate) {
		toSerialize["formation_date"] = o.FormationDate
	}
	if !IsNil(o.FormationState) {
		toSerialize["formation_state"] = o.FormationState
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Structure) {
		toSerialize["structure"] = o.Structure
	}
	if !IsNil(o.TradeNames) {
		toSerialize["trade_names"] = o.TradeNames
	}
	if !IsNil(o.VerificationLastRun) {
		toSerialize["verification_last_run"] = o.VerificationLastRun
	}
	if !IsNil(o.VerificationStatus) {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return toSerialize, nil
}

type NullableBusiness1 struct {
	value *Business1
	isSet bool
}

func (v NullableBusiness1) Get() *Business1 {
	return v.value
}

func (v *NullableBusiness1) Set(val *Business1) {
	v.value = val
	v.isSet = true
}

func (v NullableBusiness1) IsSet() bool {
	return v.isSet
}

func (v *NullableBusiness1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusiness1(val *Business1) *NullableBusiness1 {
	return &NullableBusiness1{value: val, isSet: true}
}

func (v NullableBusiness1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusiness1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
