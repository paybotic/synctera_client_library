/*
Synctera API

This is the official reference documentation for Synctera APIs. If you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact-us' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.150.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the L2l3Model type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &L2l3Model{}

// L2l3Model struct for L2l3Model
type L2l3Model struct {
	EnhancedDataId   *string            `json:"enhanced_data_id,omitempty"`
	Financial        *Financial         `json:"financial,omitempty"`
	FleetEmv         *FleetsEmv         `json:"fleet_emv,omitempty"`
	Fleets           *Fleets            `json:"fleets,omitempty"`
	InventoryDetails []InventoryDetails `json:"inventory_details,omitempty"`
	// Value of the `transaction.token` returned in the simulated clearing response.
	OriginalTransactionId string  `json:"original_transaction_id"`
	PurchaseOrder         *string `json:"purchase_order,omitempty"`
}

type _L2l3Model L2l3Model

// NewL2l3Model instantiates a new L2l3Model object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewL2l3Model(originalTransactionId string) *L2l3Model {
	this := L2l3Model{}
	this.OriginalTransactionId = originalTransactionId
	return &this
}

// NewL2l3ModelWithDefaults instantiates a new L2l3Model object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewL2l3ModelWithDefaults() *L2l3Model {
	this := L2l3Model{}
	return &this
}

// GetEnhancedDataId returns the EnhancedDataId field value if set, zero value otherwise.
func (o *L2l3Model) GetEnhancedDataId() string {
	if o == nil || IsNil(o.EnhancedDataId) {
		var ret string
		return ret
	}
	return *o.EnhancedDataId
}

// GetEnhancedDataIdOk returns a tuple with the EnhancedDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2l3Model) GetEnhancedDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnhancedDataId) {
		return nil, false
	}
	return o.EnhancedDataId, true
}

// HasEnhancedDataId returns a boolean if a field has been set.
func (o *L2l3Model) HasEnhancedDataId() bool {
	if o != nil && !IsNil(o.EnhancedDataId) {
		return true
	}

	return false
}

// SetEnhancedDataId gets a reference to the given string and assigns it to the EnhancedDataId field.
func (o *L2l3Model) SetEnhancedDataId(v string) {
	o.EnhancedDataId = &v
}

// GetFinancial returns the Financial field value if set, zero value otherwise.
func (o *L2l3Model) GetFinancial() Financial {
	if o == nil || IsNil(o.Financial) {
		var ret Financial
		return ret
	}
	return *o.Financial
}

// GetFinancialOk returns a tuple with the Financial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2l3Model) GetFinancialOk() (*Financial, bool) {
	if o == nil || IsNil(o.Financial) {
		return nil, false
	}
	return o.Financial, true
}

// HasFinancial returns a boolean if a field has been set.
func (o *L2l3Model) HasFinancial() bool {
	if o != nil && !IsNil(o.Financial) {
		return true
	}

	return false
}

// SetFinancial gets a reference to the given Financial and assigns it to the Financial field.
func (o *L2l3Model) SetFinancial(v Financial) {
	o.Financial = &v
}

// GetFleetEmv returns the FleetEmv field value if set, zero value otherwise.
func (o *L2l3Model) GetFleetEmv() FleetsEmv {
	if o == nil || IsNil(o.FleetEmv) {
		var ret FleetsEmv
		return ret
	}
	return *o.FleetEmv
}

// GetFleetEmvOk returns a tuple with the FleetEmv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2l3Model) GetFleetEmvOk() (*FleetsEmv, bool) {
	if o == nil || IsNil(o.FleetEmv) {
		return nil, false
	}
	return o.FleetEmv, true
}

// HasFleetEmv returns a boolean if a field has been set.
func (o *L2l3Model) HasFleetEmv() bool {
	if o != nil && !IsNil(o.FleetEmv) {
		return true
	}

	return false
}

// SetFleetEmv gets a reference to the given FleetsEmv and assigns it to the FleetEmv field.
func (o *L2l3Model) SetFleetEmv(v FleetsEmv) {
	o.FleetEmv = &v
}

// GetFleets returns the Fleets field value if set, zero value otherwise.
func (o *L2l3Model) GetFleets() Fleets {
	if o == nil || IsNil(o.Fleets) {
		var ret Fleets
		return ret
	}
	return *o.Fleets
}

// GetFleetsOk returns a tuple with the Fleets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2l3Model) GetFleetsOk() (*Fleets, bool) {
	if o == nil || IsNil(o.Fleets) {
		return nil, false
	}
	return o.Fleets, true
}

// HasFleets returns a boolean if a field has been set.
func (o *L2l3Model) HasFleets() bool {
	if o != nil && !IsNil(o.Fleets) {
		return true
	}

	return false
}

// SetFleets gets a reference to the given Fleets and assigns it to the Fleets field.
func (o *L2l3Model) SetFleets(v Fleets) {
	o.Fleets = &v
}

// GetInventoryDetails returns the InventoryDetails field value if set, zero value otherwise.
func (o *L2l3Model) GetInventoryDetails() []InventoryDetails {
	if o == nil || IsNil(o.InventoryDetails) {
		var ret []InventoryDetails
		return ret
	}
	return o.InventoryDetails
}

// GetInventoryDetailsOk returns a tuple with the InventoryDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2l3Model) GetInventoryDetailsOk() ([]InventoryDetails, bool) {
	if o == nil || IsNil(o.InventoryDetails) {
		return nil, false
	}
	return o.InventoryDetails, true
}

// HasInventoryDetails returns a boolean if a field has been set.
func (o *L2l3Model) HasInventoryDetails() bool {
	if o != nil && !IsNil(o.InventoryDetails) {
		return true
	}

	return false
}

// SetInventoryDetails gets a reference to the given []InventoryDetails and assigns it to the InventoryDetails field.
func (o *L2l3Model) SetInventoryDetails(v []InventoryDetails) {
	o.InventoryDetails = v
}

// GetOriginalTransactionId returns the OriginalTransactionId field value
func (o *L2l3Model) GetOriginalTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalTransactionId
}

// GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field value
// and a boolean to check if the value has been set.
func (o *L2l3Model) GetOriginalTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalTransactionId, true
}

// SetOriginalTransactionId sets field value
func (o *L2l3Model) SetOriginalTransactionId(v string) {
	o.OriginalTransactionId = v
}

// GetPurchaseOrder returns the PurchaseOrder field value if set, zero value otherwise.
func (o *L2l3Model) GetPurchaseOrder() string {
	if o == nil || IsNil(o.PurchaseOrder) {
		var ret string
		return ret
	}
	return *o.PurchaseOrder
}

// GetPurchaseOrderOk returns a tuple with the PurchaseOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *L2l3Model) GetPurchaseOrderOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrder) {
		return nil, false
	}
	return o.PurchaseOrder, true
}

// HasPurchaseOrder returns a boolean if a field has been set.
func (o *L2l3Model) HasPurchaseOrder() bool {
	if o != nil && !IsNil(o.PurchaseOrder) {
		return true
	}

	return false
}

// SetPurchaseOrder gets a reference to the given string and assigns it to the PurchaseOrder field.
func (o *L2l3Model) SetPurchaseOrder(v string) {
	o.PurchaseOrder = &v
}

func (o L2l3Model) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o L2l3Model) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnhancedDataId) {
		toSerialize["enhanced_data_id"] = o.EnhancedDataId
	}
	if !IsNil(o.Financial) {
		toSerialize["financial"] = o.Financial
	}
	if !IsNil(o.FleetEmv) {
		toSerialize["fleet_emv"] = o.FleetEmv
	}
	if !IsNil(o.Fleets) {
		toSerialize["fleets"] = o.Fleets
	}
	if !IsNil(o.InventoryDetails) {
		toSerialize["inventory_details"] = o.InventoryDetails
	}
	toSerialize["original_transaction_id"] = o.OriginalTransactionId
	if !IsNil(o.PurchaseOrder) {
		toSerialize["purchase_order"] = o.PurchaseOrder
	}
	return toSerialize, nil
}

func (o *L2l3Model) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"original_transaction_id",
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

	varL2l3Model := _L2l3Model{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varL2l3Model)

	if err != nil {
		return err
	}

	*o = L2l3Model(varL2l3Model)

	return err
}

type NullableL2l3Model struct {
	value *L2l3Model
	isSet bool
}

func (v NullableL2l3Model) Get() *L2l3Model {
	return v.value
}

func (v *NullableL2l3Model) Set(val *L2l3Model) {
	v.value = val
	v.isSet = true
}

func (v NullableL2l3Model) IsSet() bool {
	return v.isSet
}

func (v *NullableL2l3Model) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableL2l3Model(val *L2l3Model) *NullableL2l3Model {
	return &NullableL2l3Model{value: val, isSet: true}
}

func (v NullableL2l3Model) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableL2l3Model) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
