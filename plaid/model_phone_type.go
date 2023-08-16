/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.413.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PhoneType An enum indicating whether a phone number is a phone line or a fax line.
type PhoneType string

var _ = fmt.Printf

// List of PhoneType
const (
	PHONETYPE_PHONE PhoneType = "phone"
	PHONETYPE_FAX PhoneType = "fax"
)

var allowedPhoneTypeEnumValues = []PhoneType{
	"phone",
	"fax",
}

func (v *PhoneType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PhoneType(value)


	*v = enumTypeValue
	return nil
}

// NewPhoneTypeFromValue returns a pointer to a valid PhoneType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPhoneTypeFromValue(v string) (*PhoneType, error) {
	ev := PhoneType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PhoneType) IsValid() bool {
	for _, existing := range allowedPhoneTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PhoneType value
func (v PhoneType) Ptr() *PhoneType {
	return &v
}

type NullablePhoneType struct {
	value *PhoneType
	isSet bool
}

func (v NullablePhoneType) Get() *PhoneType {
	return v.value
}

func (v *NullablePhoneType) Set(val *PhoneType) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneType) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneType(val *PhoneType) *NullablePhoneType {
	return &NullablePhoneType{value: val, isSet: true}
}

func (v NullablePhoneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
