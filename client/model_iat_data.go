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
)

// checks if the IatData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IatData{}

// IatData Aggregates contents of the IAT addenda records (10-16)
type IatData struct {
	// For inbound IAT payments this field should contain the USD amount or may be blank. Amount is in cents ($100 would be 10000).
	ForeignPaymentAmount int32  `json:"foreign_payment_amount"`
	ForeignTraceNumber   string `json:"foreign_trace_number"`
	// Receiving Company Name/Individual Name
	Name string `json:"name"`
	// For Inbound IATs: This 3 position field contains a 2-character code as approved by the International Organization for Standardization (ISO) used to identify the country in which the branch of the bank that originated the entry is located. Values for other countries can be found on the International Organization for Standardization website: www.iso.org.
	OdfiBranchCountryCode string `json:"odfi_branch_country_code"`
	// For Inbound IATs: The 2-digit code that identifies the numbering scheme used in the Foreign DFI Identification Number field: 01 = National Clearing System 02 = BIC Code 03 = IBAN Code
	OdfiIdNumberQualifier string `json:"odfi_id_number_qualifier"`
	// For Inbound IATs: This field contains the bank ID number of the Foreign Bank providing funding for the payment transaction.
	OdfiIdentification string `json:"odfi_identification"`
	// For Inbound IATs: The name of the foreign bank providing funding for the payment transaction
	OdfiName          string     `json:"odfi_name"`
	OriginatorAddress IatAddress `json:"originator_address"`
	// The originators name
	OriginatorName string `json:"originator_name"`
	// This 3 position field contains a 2-character code as approved by the International Organization for Standardization (ISO) used to identify the country in which the branch of the bank that receives the entry is located. Values for other countries can be found on the International Organization for Standardization website: www.iso.org
	RdfiBranchCountryCode string `json:"rdfi_branch_country_code"`
	// The 2-digit code that identifies the numbering scheme used in the Receiving DFI Identification Number field: 01 = National Clearing System 02 = BIC Code 03 = IBAN Code
	RdfiIdNumberQualifier string `json:"rdfi_id_number_qualifier"`
	// The bank identification number of the DFI at which the Receiver maintains his account.
	RdfiIdentification string `json:"rdfi_identification"`
	// Name of the Receiver's bank
	RdfiName        string     `json:"rdfi_name"`
	ReceiverAddress IatAddress `json:"receiver_address"`
	// The accounting number by which the Originator is known to the Receiver for descriptive purposes.
	ReceiverIdNumber string `json:"receiver_id_number"`
	// Transaction Type Code Describes the type of payment ANN = Annuity, BUS = Business/Commercial, DEP = Deposit, LOA = Loan, MIS = Miscellaneous, MOR = Mortgage PEN = Pension, RLS = Rent/Lease, REM = Remittance2, SAL = Salary/Payroll, TAX = Tax, TEL = Telephone-Initiated Transaction WEB = Internet-Initiated Transaction, ARC = Accounts Receivable Entry, BOC = Back Office Conversion Entry, POP = Point of Purchase Entry, RCK = Re-presented Check Entry
	TransactionTypeCode  string `json:"transaction_type_code"`
	AdditionalProperties map[string]interface{}
}

type _IatData IatData

// NewIatData instantiates a new IatData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIatData(foreignPaymentAmount int32, foreignTraceNumber string, name string, odfiBranchCountryCode string, odfiIdNumberQualifier string, odfiIdentification string, odfiName string, originatorAddress IatAddress, originatorName string, rdfiBranchCountryCode string, rdfiIdNumberQualifier string, rdfiIdentification string, rdfiName string, receiverAddress IatAddress, receiverIdNumber string, transactionTypeCode string) *IatData {
	this := IatData{}
	this.ForeignPaymentAmount = foreignPaymentAmount
	this.ForeignTraceNumber = foreignTraceNumber
	this.Name = name
	this.OdfiBranchCountryCode = odfiBranchCountryCode
	this.OdfiIdNumberQualifier = odfiIdNumberQualifier
	this.OdfiIdentification = odfiIdentification
	this.OdfiName = odfiName
	this.OriginatorAddress = originatorAddress
	this.OriginatorName = originatorName
	this.RdfiBranchCountryCode = rdfiBranchCountryCode
	this.RdfiIdNumberQualifier = rdfiIdNumberQualifier
	this.RdfiIdentification = rdfiIdentification
	this.RdfiName = rdfiName
	this.ReceiverAddress = receiverAddress
	this.ReceiverIdNumber = receiverIdNumber
	this.TransactionTypeCode = transactionTypeCode
	return &this
}

// NewIatDataWithDefaults instantiates a new IatData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIatDataWithDefaults() *IatData {
	this := IatData{}
	return &this
}

// GetForeignPaymentAmount returns the ForeignPaymentAmount field value
func (o *IatData) GetForeignPaymentAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ForeignPaymentAmount
}

// GetForeignPaymentAmountOk returns a tuple with the ForeignPaymentAmount field value
// and a boolean to check if the value has been set.
func (o *IatData) GetForeignPaymentAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForeignPaymentAmount, true
}

// SetForeignPaymentAmount sets field value
func (o *IatData) SetForeignPaymentAmount(v int32) {
	o.ForeignPaymentAmount = v
}

// GetForeignTraceNumber returns the ForeignTraceNumber field value
func (o *IatData) GetForeignTraceNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForeignTraceNumber
}

// GetForeignTraceNumberOk returns a tuple with the ForeignTraceNumber field value
// and a boolean to check if the value has been set.
func (o *IatData) GetForeignTraceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForeignTraceNumber, true
}

// SetForeignTraceNumber sets field value
func (o *IatData) SetForeignTraceNumber(v string) {
	o.ForeignTraceNumber = v
}

// GetName returns the Name field value
func (o *IatData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IatData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IatData) SetName(v string) {
	o.Name = v
}

// GetOdfiBranchCountryCode returns the OdfiBranchCountryCode field value
func (o *IatData) GetOdfiBranchCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdfiBranchCountryCode
}

// GetOdfiBranchCountryCodeOk returns a tuple with the OdfiBranchCountryCode field value
// and a boolean to check if the value has been set.
func (o *IatData) GetOdfiBranchCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdfiBranchCountryCode, true
}

// SetOdfiBranchCountryCode sets field value
func (o *IatData) SetOdfiBranchCountryCode(v string) {
	o.OdfiBranchCountryCode = v
}

// GetOdfiIdNumberQualifier returns the OdfiIdNumberQualifier field value
func (o *IatData) GetOdfiIdNumberQualifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdfiIdNumberQualifier
}

// GetOdfiIdNumberQualifierOk returns a tuple with the OdfiIdNumberQualifier field value
// and a boolean to check if the value has been set.
func (o *IatData) GetOdfiIdNumberQualifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdfiIdNumberQualifier, true
}

// SetOdfiIdNumberQualifier sets field value
func (o *IatData) SetOdfiIdNumberQualifier(v string) {
	o.OdfiIdNumberQualifier = v
}

// GetOdfiIdentification returns the OdfiIdentification field value
func (o *IatData) GetOdfiIdentification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdfiIdentification
}

// GetOdfiIdentificationOk returns a tuple with the OdfiIdentification field value
// and a boolean to check if the value has been set.
func (o *IatData) GetOdfiIdentificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdfiIdentification, true
}

// SetOdfiIdentification sets field value
func (o *IatData) SetOdfiIdentification(v string) {
	o.OdfiIdentification = v
}

// GetOdfiName returns the OdfiName field value
func (o *IatData) GetOdfiName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OdfiName
}

// GetOdfiNameOk returns a tuple with the OdfiName field value
// and a boolean to check if the value has been set.
func (o *IatData) GetOdfiNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OdfiName, true
}

// SetOdfiName sets field value
func (o *IatData) SetOdfiName(v string) {
	o.OdfiName = v
}

// GetOriginatorAddress returns the OriginatorAddress field value
func (o *IatData) GetOriginatorAddress() IatAddress {
	if o == nil {
		var ret IatAddress
		return ret
	}

	return o.OriginatorAddress
}

// GetOriginatorAddressOk returns a tuple with the OriginatorAddress field value
// and a boolean to check if the value has been set.
func (o *IatData) GetOriginatorAddressOk() (*IatAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatorAddress, true
}

// SetOriginatorAddress sets field value
func (o *IatData) SetOriginatorAddress(v IatAddress) {
	o.OriginatorAddress = v
}

// GetOriginatorName returns the OriginatorName field value
func (o *IatData) GetOriginatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatorName
}

// GetOriginatorNameOk returns a tuple with the OriginatorName field value
// and a boolean to check if the value has been set.
func (o *IatData) GetOriginatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatorName, true
}

// SetOriginatorName sets field value
func (o *IatData) SetOriginatorName(v string) {
	o.OriginatorName = v
}

// GetRdfiBranchCountryCode returns the RdfiBranchCountryCode field value
func (o *IatData) GetRdfiBranchCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RdfiBranchCountryCode
}

// GetRdfiBranchCountryCodeOk returns a tuple with the RdfiBranchCountryCode field value
// and a boolean to check if the value has been set.
func (o *IatData) GetRdfiBranchCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RdfiBranchCountryCode, true
}

// SetRdfiBranchCountryCode sets field value
func (o *IatData) SetRdfiBranchCountryCode(v string) {
	o.RdfiBranchCountryCode = v
}

// GetRdfiIdNumberQualifier returns the RdfiIdNumberQualifier field value
func (o *IatData) GetRdfiIdNumberQualifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RdfiIdNumberQualifier
}

// GetRdfiIdNumberQualifierOk returns a tuple with the RdfiIdNumberQualifier field value
// and a boolean to check if the value has been set.
func (o *IatData) GetRdfiIdNumberQualifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RdfiIdNumberQualifier, true
}

// SetRdfiIdNumberQualifier sets field value
func (o *IatData) SetRdfiIdNumberQualifier(v string) {
	o.RdfiIdNumberQualifier = v
}

// GetRdfiIdentification returns the RdfiIdentification field value
func (o *IatData) GetRdfiIdentification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RdfiIdentification
}

// GetRdfiIdentificationOk returns a tuple with the RdfiIdentification field value
// and a boolean to check if the value has been set.
func (o *IatData) GetRdfiIdentificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RdfiIdentification, true
}

// SetRdfiIdentification sets field value
func (o *IatData) SetRdfiIdentification(v string) {
	o.RdfiIdentification = v
}

// GetRdfiName returns the RdfiName field value
func (o *IatData) GetRdfiName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RdfiName
}

// GetRdfiNameOk returns a tuple with the RdfiName field value
// and a boolean to check if the value has been set.
func (o *IatData) GetRdfiNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RdfiName, true
}

// SetRdfiName sets field value
func (o *IatData) SetRdfiName(v string) {
	o.RdfiName = v
}

// GetReceiverAddress returns the ReceiverAddress field value
func (o *IatData) GetReceiverAddress() IatAddress {
	if o == nil {
		var ret IatAddress
		return ret
	}

	return o.ReceiverAddress
}

// GetReceiverAddressOk returns a tuple with the ReceiverAddress field value
// and a boolean to check if the value has been set.
func (o *IatData) GetReceiverAddressOk() (*IatAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverAddress, true
}

// SetReceiverAddress sets field value
func (o *IatData) SetReceiverAddress(v IatAddress) {
	o.ReceiverAddress = v
}

// GetReceiverIdNumber returns the ReceiverIdNumber field value
func (o *IatData) GetReceiverIdNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiverIdNumber
}

// GetReceiverIdNumberOk returns a tuple with the ReceiverIdNumber field value
// and a boolean to check if the value has been set.
func (o *IatData) GetReceiverIdNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverIdNumber, true
}

// SetReceiverIdNumber sets field value
func (o *IatData) SetReceiverIdNumber(v string) {
	o.ReceiverIdNumber = v
}

// GetTransactionTypeCode returns the TransactionTypeCode field value
func (o *IatData) GetTransactionTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionTypeCode
}

// GetTransactionTypeCodeOk returns a tuple with the TransactionTypeCode field value
// and a boolean to check if the value has been set.
func (o *IatData) GetTransactionTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionTypeCode, true
}

// SetTransactionTypeCode sets field value
func (o *IatData) SetTransactionTypeCode(v string) {
	o.TransactionTypeCode = v
}

func (o IatData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IatData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["foreign_payment_amount"] = o.ForeignPaymentAmount
	toSerialize["foreign_trace_number"] = o.ForeignTraceNumber
	toSerialize["name"] = o.Name
	toSerialize["odfi_branch_country_code"] = o.OdfiBranchCountryCode
	toSerialize["odfi_id_number_qualifier"] = o.OdfiIdNumberQualifier
	toSerialize["odfi_identification"] = o.OdfiIdentification
	toSerialize["odfi_name"] = o.OdfiName
	toSerialize["originator_address"] = o.OriginatorAddress
	toSerialize["originator_name"] = o.OriginatorName
	toSerialize["rdfi_branch_country_code"] = o.RdfiBranchCountryCode
	toSerialize["rdfi_id_number_qualifier"] = o.RdfiIdNumberQualifier
	toSerialize["rdfi_identification"] = o.RdfiIdentification
	toSerialize["rdfi_name"] = o.RdfiName
	toSerialize["receiver_address"] = o.ReceiverAddress
	toSerialize["receiver_id_number"] = o.ReceiverIdNumber
	toSerialize["transaction_type_code"] = o.TransactionTypeCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IatData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"foreign_payment_amount",
		"foreign_trace_number",
		"name",
		"odfi_branch_country_code",
		"odfi_id_number_qualifier",
		"odfi_identification",
		"odfi_name",
		"originator_address",
		"originator_name",
		"rdfi_branch_country_code",
		"rdfi_id_number_qualifier",
		"rdfi_identification",
		"rdfi_name",
		"receiver_address",
		"receiver_id_number",
		"transaction_type_code",
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

	varIatData := _IatData{}

	err = json.Unmarshal(data, &varIatData)

	if err != nil {
		return err
	}

	*o = IatData(varIatData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "foreign_payment_amount")
		delete(additionalProperties, "foreign_trace_number")
		delete(additionalProperties, "name")
		delete(additionalProperties, "odfi_branch_country_code")
		delete(additionalProperties, "odfi_id_number_qualifier")
		delete(additionalProperties, "odfi_identification")
		delete(additionalProperties, "odfi_name")
		delete(additionalProperties, "originator_address")
		delete(additionalProperties, "originator_name")
		delete(additionalProperties, "rdfi_branch_country_code")
		delete(additionalProperties, "rdfi_id_number_qualifier")
		delete(additionalProperties, "rdfi_identification")
		delete(additionalProperties, "rdfi_name")
		delete(additionalProperties, "receiver_address")
		delete(additionalProperties, "receiver_id_number")
		delete(additionalProperties, "transaction_type_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIatData struct {
	value *IatData
	isSet bool
}

func (v NullableIatData) Get() *IatData {
	return v.value
}

func (v *NullableIatData) Set(val *IatData) {
	v.value = val
	v.isSet = true
}

func (v NullableIatData) IsSet() bool {
	return v.isSet
}

func (v *NullableIatData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIatData(val *IatData) *NullableIatData {
	return &NullableIatData{value: val, isSet: true}
}

func (v NullableIatData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIatData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
