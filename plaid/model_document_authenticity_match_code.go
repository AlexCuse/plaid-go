/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.410.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// DocumentAuthenticityMatchCode High level summary of whether the document in the provided image matches the formatting rules and security checks for the associated jurisdiction.  For example, most identity documents have formatting rules like the following:   The image of the person's face must have a certain contrast in order to highlight skin tone   The subject in the document's image must remove eye glasses and pose in a certain way   The informational fields (name, date of birth, ID number, etc.) must be colored and aligned according to specific rules   Security features like watermarks and background patterns must be present  So a `match` status for this field indicates that the document in the provided image seems to conform to the various formatting and security rules associated with the detected document.
type DocumentAuthenticityMatchCode string

var _ = fmt.Printf

// List of DocumentAuthenticityMatchCode
const (
	DOCUMENTAUTHENTICITYMATCHCODE_MATCH DocumentAuthenticityMatchCode = "match"
	DOCUMENTAUTHENTICITYMATCHCODE_PARTIAL_MATCH DocumentAuthenticityMatchCode = "partial_match"
	DOCUMENTAUTHENTICITYMATCHCODE_NO_MATCH DocumentAuthenticityMatchCode = "no_match"
	DOCUMENTAUTHENTICITYMATCHCODE_NO_DATA DocumentAuthenticityMatchCode = "no_data"
)

var allowedDocumentAuthenticityMatchCodeEnumValues = []DocumentAuthenticityMatchCode{
	"match",
	"partial_match",
	"no_match",
	"no_data",
}

func (v *DocumentAuthenticityMatchCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := DocumentAuthenticityMatchCode(value)


	*v = enumTypeValue
	return nil
}

// NewDocumentAuthenticityMatchCodeFromValue returns a pointer to a valid DocumentAuthenticityMatchCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentAuthenticityMatchCodeFromValue(v string) (*DocumentAuthenticityMatchCode, error) {
	ev := DocumentAuthenticityMatchCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentAuthenticityMatchCode) IsValid() bool {
	for _, existing := range allowedDocumentAuthenticityMatchCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentAuthenticityMatchCode value
func (v DocumentAuthenticityMatchCode) Ptr() *DocumentAuthenticityMatchCode {
	return &v
}

type NullableDocumentAuthenticityMatchCode struct {
	value *DocumentAuthenticityMatchCode
	isSet bool
}

func (v NullableDocumentAuthenticityMatchCode) Get() *DocumentAuthenticityMatchCode {
	return v.value
}

func (v *NullableDocumentAuthenticityMatchCode) Set(val *DocumentAuthenticityMatchCode) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentAuthenticityMatchCode) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentAuthenticityMatchCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentAuthenticityMatchCode(val *DocumentAuthenticityMatchCode) *NullableDocumentAuthenticityMatchCode {
	return &NullableDocumentAuthenticityMatchCode{value: val, isSet: true}
}

func (v NullableDocumentAuthenticityMatchCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentAuthenticityMatchCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

