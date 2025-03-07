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

// checks if the PostedTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostedTransaction{}

// PostedTransaction struct for PostedTransaction
type PostedTransaction struct {
	// The creation date of the transaction
	Created time.Time             `json:"created"`
	Data    PostedTransactionData `json:"data"`
	// The set of disputes related to this transaction. Since a dispute can be for a partial amount of a transaction, a single transaction can be involved in multiple disputes.
	Disputes []TransactionDispute `json:"disputes,omitempty"`
	// The \"effective date\" of a transaction. This may be earlier than posted_date in some cases (for example, a transaction that occurs on a Saturday may not be posted until the following Monday, but would have an effective date of Saturday)
	EffectiveDate time.Time `json:"effective_date"`
	Id            int64     `json:"id"`
	// The idempotency key used when initially creating this transaction.
	Idemkey string `json:"idemkey"`
	// Whether or not this transaction represents a purely informational operation or an actual money movement
	InfoOnly bool `json:"info_only"`
	// Whether or not this transaction was created operating in \"lead ledger\" mode
	LeadMode bool `json:"lead_mode"`
	// The date the transaction was posted. This is the date any money is considered to be added or removed from an account.
	PostedDate time.Time `json:"posted_date"`
	// An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers.
	ReferenceId NullableString `json:"reference_id"`
	// The date the transaction was settled according to Synctera's platform. Generally, this can be interpretted the date the transaction was actually processed and settlement by the payment network.
	SettlementDate *string `json:"settlement_date,omitempty"`
	Status         string  `json:"status"`
	// The specific transaction type. For example, for `ach`, this may be \"outgoing_debit\".
	Subtype string `json:"subtype"`
	// The id of the tenant containing the resource. This is relevant for Fintechs that have multiple workspaces.
	Tenant string `json:"tenant"`
	// The time the transaction occurred.
	TransactionTime time.Time `json:"transaction_time"`
	// The general type of transaction. For example, \"card\" or \"ach\".
	Type string `json:"type"`
	// The date the transaction was last updated
	Updated time.Time `json:"updated"`
	// The unique identifier of the transaction.
	Uuid                 string `json:"uuid"`
	AdditionalProperties map[string]interface{}
}

type _PostedTransaction PostedTransaction

// NewPostedTransaction instantiates a new PostedTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostedTransaction(created time.Time, data PostedTransactionData, effectiveDate time.Time, id int64, idemkey string, infoOnly bool, leadMode bool, postedDate time.Time, referenceId NullableString, status string, subtype string, tenant string, transactionTime time.Time, type_ string, updated time.Time, uuid string) *PostedTransaction {
	this := PostedTransaction{}
	this.Created = created
	this.Data = data
	this.EffectiveDate = effectiveDate
	this.Id = id
	this.Idemkey = idemkey
	this.InfoOnly = infoOnly
	this.LeadMode = leadMode
	this.PostedDate = postedDate
	this.ReferenceId = referenceId
	this.Status = status
	this.Subtype = subtype
	this.Tenant = tenant
	this.TransactionTime = transactionTime
	this.Type = type_
	this.Updated = updated
	this.Uuid = uuid
	return &this
}

// NewPostedTransactionWithDefaults instantiates a new PostedTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostedTransactionWithDefaults() *PostedTransaction {
	this := PostedTransaction{}
	return &this
}

// GetCreated returns the Created field value
func (o *PostedTransaction) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PostedTransaction) SetCreated(v time.Time) {
	o.Created = v
}

// GetData returns the Data field value
func (o *PostedTransaction) GetData() PostedTransactionData {
	if o == nil {
		var ret PostedTransactionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetDataOk() (*PostedTransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostedTransaction) SetData(v PostedTransactionData) {
	o.Data = v
}

// GetDisputes returns the Disputes field value if set, zero value otherwise.
func (o *PostedTransaction) GetDisputes() []TransactionDispute {
	if o == nil || IsNil(o.Disputes) {
		var ret []TransactionDispute
		return ret
	}
	return o.Disputes
}

// GetDisputesOk returns a tuple with the Disputes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetDisputesOk() ([]TransactionDispute, bool) {
	if o == nil || IsNil(o.Disputes) {
		return nil, false
	}
	return o.Disputes, true
}

// HasDisputes returns a boolean if a field has been set.
func (o *PostedTransaction) HasDisputes() bool {
	if o != nil && !IsNil(o.Disputes) {
		return true
	}

	return false
}

// SetDisputes gets a reference to the given []TransactionDispute and assigns it to the Disputes field.
func (o *PostedTransaction) SetDisputes(v []TransactionDispute) {
	o.Disputes = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *PostedTransaction) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *PostedTransaction) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetId returns the Id field value
func (o *PostedTransaction) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostedTransaction) SetId(v int64) {
	o.Id = v
}

// GetIdemkey returns the Idemkey field value
func (o *PostedTransaction) GetIdemkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Idemkey
}

// GetIdemkeyOk returns a tuple with the Idemkey field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetIdemkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idemkey, true
}

// SetIdemkey sets field value
func (o *PostedTransaction) SetIdemkey(v string) {
	o.Idemkey = v
}

// GetInfoOnly returns the InfoOnly field value
func (o *PostedTransaction) GetInfoOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InfoOnly
}

// GetInfoOnlyOk returns a tuple with the InfoOnly field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetInfoOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InfoOnly, true
}

// SetInfoOnly sets field value
func (o *PostedTransaction) SetInfoOnly(v bool) {
	o.InfoOnly = v
}

// GetLeadMode returns the LeadMode field value
func (o *PostedTransaction) GetLeadMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LeadMode
}

// GetLeadModeOk returns a tuple with the LeadMode field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetLeadModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeadMode, true
}

// SetLeadMode sets field value
func (o *PostedTransaction) SetLeadMode(v bool) {
	o.LeadMode = v
}

// GetPostedDate returns the PostedDate field value
func (o *PostedTransaction) GetPostedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetPostedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostedDate, true
}

// SetPostedDate sets field value
func (o *PostedTransaction) SetPostedDate(v time.Time) {
	o.PostedDate = v
}

// GetReferenceId returns the ReferenceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostedTransaction) GetReferenceId() string {
	if o == nil || o.ReferenceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostedTransaction) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// SetReferenceId sets field value
func (o *PostedTransaction) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}

// GetSettlementDate returns the SettlementDate field value if set, zero value otherwise.
func (o *PostedTransaction) GetSettlementDate() string {
	if o == nil || IsNil(o.SettlementDate) {
		var ret string
		return ret
	}
	return *o.SettlementDate
}

// GetSettlementDateOk returns a tuple with the SettlementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetSettlementDateOk() (*string, bool) {
	if o == nil || IsNil(o.SettlementDate) {
		return nil, false
	}
	return o.SettlementDate, true
}

// HasSettlementDate returns a boolean if a field has been set.
func (o *PostedTransaction) HasSettlementDate() bool {
	if o != nil && !IsNil(o.SettlementDate) {
		return true
	}

	return false
}

// SetSettlementDate gets a reference to the given string and assigns it to the SettlementDate field.
func (o *PostedTransaction) SetSettlementDate(v string) {
	o.SettlementDate = &v
}

// GetStatus returns the Status field value
func (o *PostedTransaction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PostedTransaction) SetStatus(v string) {
	o.Status = v
}

// GetSubtype returns the Subtype field value
func (o *PostedTransaction) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *PostedTransaction) SetSubtype(v string) {
	o.Subtype = v
}

// GetTenant returns the Tenant field value
func (o *PostedTransaction) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *PostedTransaction) SetTenant(v string) {
	o.Tenant = v
}

// GetTransactionTime returns the TransactionTime field value
func (o *PostedTransaction) GetTransactionTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetTransactionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTime, true
}

// SetTransactionTime sets field value
func (o *PostedTransaction) SetTransactionTime(v time.Time) {
	o.TransactionTime = v
}

// GetType returns the Type field value
func (o *PostedTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostedTransaction) SetType(v string) {
	o.Type = v
}

// GetUpdated returns the Updated field value
func (o *PostedTransaction) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *PostedTransaction) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetUuid returns the Uuid field value
func (o *PostedTransaction) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *PostedTransaction) SetUuid(v string) {
	o.Uuid = v
}

func (o PostedTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostedTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["data"] = o.Data
	if !IsNil(o.Disputes) {
		toSerialize["disputes"] = o.Disputes
	}
	toSerialize["effective_date"] = o.EffectiveDate
	toSerialize["id"] = o.Id
	toSerialize["idemkey"] = o.Idemkey
	toSerialize["info_only"] = o.InfoOnly
	toSerialize["lead_mode"] = o.LeadMode
	toSerialize["posted_date"] = o.PostedDate
	toSerialize["reference_id"] = o.ReferenceId.Get()
	if !IsNil(o.SettlementDate) {
		toSerialize["settlement_date"] = o.SettlementDate
	}
	toSerialize["status"] = o.Status
	toSerialize["subtype"] = o.Subtype
	toSerialize["tenant"] = o.Tenant
	toSerialize["transaction_time"] = o.TransactionTime
	toSerialize["type"] = o.Type
	toSerialize["updated"] = o.Updated
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostedTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"data",
		"effective_date",
		"id",
		"idemkey",
		"info_only",
		"lead_mode",
		"posted_date",
		"reference_id",
		"status",
		"subtype",
		"tenant",
		"transaction_time",
		"type",
		"updated",
		"uuid",
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

	varPostedTransaction := _PostedTransaction{}

	err = json.Unmarshal(data, &varPostedTransaction)

	if err != nil {
		return err
	}

	*o = PostedTransaction(varPostedTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "data")
		delete(additionalProperties, "disputes")
		delete(additionalProperties, "effective_date")
		delete(additionalProperties, "id")
		delete(additionalProperties, "idemkey")
		delete(additionalProperties, "info_only")
		delete(additionalProperties, "lead_mode")
		delete(additionalProperties, "posted_date")
		delete(additionalProperties, "reference_id")
		delete(additionalProperties, "settlement_date")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "transaction_time")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostedTransaction struct {
	value *PostedTransaction
	isSet bool
}

func (v NullablePostedTransaction) Get() *PostedTransaction {
	return v.value
}

func (v *NullablePostedTransaction) Set(val *PostedTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePostedTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePostedTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostedTransaction(val *PostedTransaction) *NullablePostedTransaction {
	return &NullablePostedTransaction{value: val, isSet: true}
}

func (v NullablePostedTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostedTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
