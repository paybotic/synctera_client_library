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

// DocumentEncryptionVersionPost Whether the file will be encrypted by the Synctera platform before storing. All documents containing PII must be encrypted. When creating a new document version, the encryption value must match the encryption value of the previous version. Due to this, we recommend that you omit this property when creating a new document version. This property is deprecated and will be removed in a future release.
type DocumentEncryptionVersionPost string

// List of document_encryption_version_post
const (
	DOCUMENTENCRYPTIONVERSIONPOST_NOT_REQUIRED DocumentEncryptionVersionPost = "NOT_REQUIRED"
	DOCUMENTENCRYPTIONVERSIONPOST_REQUIRED     DocumentEncryptionVersionPost = "REQUIRED"
)

// All allowed values of DocumentEncryptionVersionPost enum
var AllowedDocumentEncryptionVersionPostEnumValues = []DocumentEncryptionVersionPost{
	"NOT_REQUIRED",
	"REQUIRED",
}

func (v *DocumentEncryptionVersionPost) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentEncryptionVersionPost(value)
	for _, existing := range AllowedDocumentEncryptionVersionPostEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentEncryptionVersionPost", value)
}

// NewDocumentEncryptionVersionPostFromValue returns a pointer to a valid DocumentEncryptionVersionPost
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentEncryptionVersionPostFromValue(v string) (*DocumentEncryptionVersionPost, error) {
	ev := DocumentEncryptionVersionPost(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentEncryptionVersionPost: valid values are %v", v, AllowedDocumentEncryptionVersionPostEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentEncryptionVersionPost) IsValid() bool {
	for _, existing := range AllowedDocumentEncryptionVersionPostEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to document_encryption_version_post value
func (v DocumentEncryptionVersionPost) Ptr() *DocumentEncryptionVersionPost {
	return &v
}

type NullableDocumentEncryptionVersionPost struct {
	value *DocumentEncryptionVersionPost
	isSet bool
}

func (v NullableDocumentEncryptionVersionPost) Get() *DocumentEncryptionVersionPost {
	return v.value
}

func (v *NullableDocumentEncryptionVersionPost) Set(val *DocumentEncryptionVersionPost) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentEncryptionVersionPost) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentEncryptionVersionPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentEncryptionVersionPost(val *DocumentEncryptionVersionPost) *NullableDocumentEncryptionVersionPost {
	return &NullableDocumentEncryptionVersionPost{value: val, isSet: true}
}

func (v NullableDocumentEncryptionVersionPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentEncryptionVersionPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
