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

// ISOCurrencyCode An ISO-4217 currency code.
type ISOCurrencyCode string

var _ = fmt.Printf

// List of ISOCurrencyCode
const (
	ISOCURRENCYCODE_USD ISOCurrencyCode = "USD"
)

var allowedISOCurrencyCodeEnumValues = []ISOCurrencyCode{
	"USD",
}

func (v *ISOCurrencyCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := ISOCurrencyCode(value)


	*v = enumTypeValue
	return nil
}

// NewISOCurrencyCodeFromValue returns a pointer to a valid ISOCurrencyCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewISOCurrencyCodeFromValue(v string) (*ISOCurrencyCode, error) {
	ev := ISOCurrencyCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ISOCurrencyCode) IsValid() bool {
	for _, existing := range allowedISOCurrencyCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ISOCurrencyCode value
func (v ISOCurrencyCode) Ptr() *ISOCurrencyCode {
	return &v
}

type NullableISOCurrencyCode struct {
	value *ISOCurrencyCode
	isSet bool
}

func (v NullableISOCurrencyCode) Get() *ISOCurrencyCode {
	return v.value
}

func (v *NullableISOCurrencyCode) Set(val *ISOCurrencyCode) {
	v.value = val
	v.isSet = true
}

func (v NullableISOCurrencyCode) IsSet() bool {
	return v.isSet
}

func (v *NullableISOCurrencyCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableISOCurrencyCode(val *ISOCurrencyCode) *NullableISOCurrencyCode {
	return &NullableISOCurrencyCode{value: val, isSet: true}
}

func (v NullableISOCurrencyCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableISOCurrencyCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

