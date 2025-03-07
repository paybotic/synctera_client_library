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

// CreateEddRequest - struct for CreateEddRequest
type CreateEddRequest struct {
	EddAccount     *EddAccount
	EddBusiness    *EddBusiness
	EddCustomer    *EddCustomer
	EddTransaction *EddTransaction
}

// EddAccountAsCreateEddRequest is a convenience function that returns EddAccount wrapped in CreateEddRequest
func EddAccountAsCreateEddRequest(v *EddAccount) CreateEddRequest {
	return CreateEddRequest{
		EddAccount: v,
	}
}

// EddBusinessAsCreateEddRequest is a convenience function that returns EddBusiness wrapped in CreateEddRequest
func EddBusinessAsCreateEddRequest(v *EddBusiness) CreateEddRequest {
	return CreateEddRequest{
		EddBusiness: v,
	}
}

// EddCustomerAsCreateEddRequest is a convenience function that returns EddCustomer wrapped in CreateEddRequest
func EddCustomerAsCreateEddRequest(v *EddCustomer) CreateEddRequest {
	return CreateEddRequest{
		EddCustomer: v,
	}
}

// EddTransactionAsCreateEddRequest is a convenience function that returns EddTransaction wrapped in CreateEddRequest
func EddTransactionAsCreateEddRequest(v *EddTransaction) CreateEddRequest {
	return CreateEddRequest{
		EddTransaction: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateEddRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ACCOUNT'
	if jsonDict["related_resource_type"] == "ACCOUNT" {
		// try to unmarshal JSON data into EddAccount
		err = json.Unmarshal(data, &dst.EddAccount)
		if err == nil {
			return nil // data stored in dst.EddAccount, return on the first match
		} else {
			dst.EddAccount = nil
			return fmt.Errorf("failed to unmarshal CreateEddRequest as EddAccount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BUSINESS'
	if jsonDict["related_resource_type"] == "BUSINESS" {
		// try to unmarshal JSON data into EddBusiness
		err = json.Unmarshal(data, &dst.EddBusiness)
		if err == nil {
			return nil // data stored in dst.EddBusiness, return on the first match
		} else {
			dst.EddBusiness = nil
			return fmt.Errorf("failed to unmarshal CreateEddRequest as EddBusiness: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CUSTOMER'
	if jsonDict["related_resource_type"] == "CUSTOMER" {
		// try to unmarshal JSON data into EddCustomer
		err = json.Unmarshal(data, &dst.EddCustomer)
		if err == nil {
			return nil // data stored in dst.EddCustomer, return on the first match
		} else {
			dst.EddCustomer = nil
			return fmt.Errorf("failed to unmarshal CreateEddRequest as EddCustomer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TRANSACTION'
	if jsonDict["related_resource_type"] == "TRANSACTION" {
		// try to unmarshal JSON data into EddTransaction
		err = json.Unmarshal(data, &dst.EddTransaction)
		if err == nil {
			return nil // data stored in dst.EddTransaction, return on the first match
		} else {
			dst.EddTransaction = nil
			return fmt.Errorf("failed to unmarshal CreateEddRequest as EddTransaction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'edd_account'
	if jsonDict["related_resource_type"] == "edd_account" {
		// try to unmarshal JSON data into EddAccount
		err = json.Unmarshal(data, &dst.EddAccount)
		if err == nil {
			return nil // data stored in dst.EddAccount, return on the first match
		} else {
			dst.EddAccount = nil
			return fmt.Errorf("failed to unmarshal CreateEddRequest as EddAccount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'edd_business'
	if jsonDict["related_resource_type"] == "edd_business" {
		// try to unmarshal JSON data into EddBusiness
		err = json.Unmarshal(data, &dst.EddBusiness)
		if err == nil {
			return nil // data stored in dst.EddBusiness, return on the first match
		} else {
			dst.EddBusiness = nil
			return fmt.Errorf("failed to unmarshal CreateEddRequest as EddBusiness: %s", err.Error())
		}
	}

	// check if the discriminator value is 'edd_customer'
	if jsonDict["related_resource_type"] == "edd_customer" {
		// try to unmarshal JSON data into EddCustomer
		err = json.Unmarshal(data, &dst.EddCustomer)
		if err == nil {
			return nil // data stored in dst.EddCustomer, return on the first match
		} else {
			dst.EddCustomer = nil
			return fmt.Errorf("failed to unmarshal CreateEddRequest as EddCustomer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'edd_transaction'
	if jsonDict["related_resource_type"] == "edd_transaction" {
		// try to unmarshal JSON data into EddTransaction
		err = json.Unmarshal(data, &dst.EddTransaction)
		if err == nil {
			return nil // data stored in dst.EddTransaction, return on the first match
		} else {
			dst.EddTransaction = nil
			return fmt.Errorf("failed to unmarshal CreateEddRequest as EddTransaction: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateEddRequest) MarshalJSON() ([]byte, error) {
	if src.EddAccount != nil {
		return json.Marshal(&src.EddAccount)
	}

	if src.EddBusiness != nil {
		return json.Marshal(&src.EddBusiness)
	}

	if src.EddCustomer != nil {
		return json.Marshal(&src.EddCustomer)
	}

	if src.EddTransaction != nil {
		return json.Marshal(&src.EddTransaction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateEddRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EddAccount != nil {
		return obj.EddAccount
	}

	if obj.EddBusiness != nil {
		return obj.EddBusiness
	}

	if obj.EddCustomer != nil {
		return obj.EddCustomer
	}

	if obj.EddTransaction != nil {
		return obj.EddTransaction
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CreateEddRequest) GetActualInstanceValue() interface{} {
	if obj.EddAccount != nil {
		return *obj.EddAccount
	}

	if obj.EddBusiness != nil {
		return *obj.EddBusiness
	}

	if obj.EddCustomer != nil {
		return *obj.EddCustomer
	}

	if obj.EddTransaction != nil {
		return *obj.EddTransaction
	}

	// all schemas are nil
	return nil
}

type NullableCreateEddRequest struct {
	value *CreateEddRequest
	isSet bool
}

func (v NullableCreateEddRequest) Get() *CreateEddRequest {
	return v.value
}

func (v *NullableCreateEddRequest) Set(val *CreateEddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEddRequest(val *CreateEddRequest) *NullableCreateEddRequest {
	return &NullableCreateEddRequest{value: val, isSet: true}
}

func (v NullableCreateEddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
