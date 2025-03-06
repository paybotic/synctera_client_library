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

// checks if the AccountRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountRelationship{}

// AccountRelationship The relationship of the account and the customer/business. Either customer_id OR business_id must be specified, but not both.
type AccountRelationship struct {
	// Business associated with the current account. One of business_id or customer_id must be specified.
	BusinessId *string `json:"business_id,omitempty"`
	// Date and time when this association was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Personal customer associated with the current account. One of customer_id or business_id must be specified.
	CustomerId *string `json:"customer_id,omitempty"`
	// Date and time when this association was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// ID of account relationship
	Id *string `json:"id,omitempty"`
	// Person associated with the current account. This attribute is deprecated and will be removed in a future API version. Use customer_id instead.
	// Deprecated
	PersonId         *string                 `json:"person_id,omitempty"`
	RelationshipType AccountRelationshipType `json:"relationship_type"`
	// Date and time when this association was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type _AccountRelationship AccountRelationship

// NewAccountRelationship instantiates a new AccountRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRelationship(relationshipType AccountRelationshipType) *AccountRelationship {
	this := AccountRelationship{}
	this.RelationshipType = relationshipType
	return &this
}

// NewAccountRelationshipWithDefaults instantiates a new AccountRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRelationshipWithDefaults() *AccountRelationship {
	this := AccountRelationship{}
	return &this
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *AccountRelationship) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRelationship) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *AccountRelationship) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *AccountRelationship) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccountRelationship) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRelationship) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccountRelationship) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AccountRelationship) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *AccountRelationship) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRelationship) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AccountRelationship) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *AccountRelationship) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *AccountRelationship) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRelationship) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AccountRelationship) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *AccountRelationship) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountRelationship) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRelationship) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountRelationship) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountRelationship) SetId(v string) {
	o.Id = &v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
// Deprecated
func (o *AccountRelationship) GetPersonId() string {
	if o == nil || IsNil(o.PersonId) {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AccountRelationship) GetPersonIdOk() (*string, bool) {
	if o == nil || IsNil(o.PersonId) {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *AccountRelationship) HasPersonId() bool {
	if o != nil && !IsNil(o.PersonId) {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
// Deprecated
func (o *AccountRelationship) SetPersonId(v string) {
	o.PersonId = &v
}

// GetRelationshipType returns the RelationshipType field value
func (o *AccountRelationship) GetRelationshipType() AccountRelationshipType {
	if o == nil {
		var ret AccountRelationshipType
		return ret
	}

	return o.RelationshipType
}

// GetRelationshipTypeOk returns a tuple with the RelationshipType field value
// and a boolean to check if the value has been set.
func (o *AccountRelationship) GetRelationshipTypeOk() (*AccountRelationshipType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationshipType, true
}

// SetRelationshipType sets field value
func (o *AccountRelationship) SetRelationshipType(v AccountRelationshipType) {
	o.RelationshipType = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AccountRelationship) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRelationship) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AccountRelationship) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AccountRelationship) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o AccountRelationship) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PersonId) {
		toSerialize["person_id"] = o.PersonId
	}
	toSerialize["relationship_type"] = o.RelationshipType
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *AccountRelationship) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"relationship_type",
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

	varAccountRelationship := _AccountRelationship{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountRelationship)

	if err != nil {
		return err
	}

	*o = AccountRelationship(varAccountRelationship)

	return err
}

type NullableAccountRelationship struct {
	value *AccountRelationship
	isSet bool
}

func (v NullableAccountRelationship) Get() *AccountRelationship {
	return v.value
}

func (v *NullableAccountRelationship) Set(val *AccountRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRelationship(val *AccountRelationship) *NullableAccountRelationship {
	return &NullableAccountRelationship{value: val, isSet: true}
}

func (v NullableAccountRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
