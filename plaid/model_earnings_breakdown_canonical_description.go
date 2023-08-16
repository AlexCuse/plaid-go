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

// EarningsBreakdownCanonicalDescription Commonly used term to describe the earning line item.
type EarningsBreakdownCanonicalDescription string

var _ = fmt.Printf

// List of EarningsBreakdownCanonicalDescription
const (
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_BONUS EarningsBreakdownCanonicalDescription = "BONUS"
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_COMMISSION EarningsBreakdownCanonicalDescription = "COMMISSION"
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_OVERTIME EarningsBreakdownCanonicalDescription = "OVERTIME"
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_PAID_TIME_OFF EarningsBreakdownCanonicalDescription = "PAID TIME OFF"
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_REGULAR_PAY EarningsBreakdownCanonicalDescription = "REGULAR PAY"
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_VACATION EarningsBreakdownCanonicalDescription = "VACATION"
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_BASIC_ALLOWANCE_HOUSING EarningsBreakdownCanonicalDescription = "BASIC ALLOWANCE HOUSING"
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_BASIC_ALLOWANCE_SUBSISTENCE EarningsBreakdownCanonicalDescription = "BASIC ALLOWANCE SUBSISTENCE"
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_OTHER EarningsBreakdownCanonicalDescription = "OTHER"
	EARNINGSBREAKDOWNCANONICALDESCRIPTION_NULL EarningsBreakdownCanonicalDescription = "null"
)

var allowedEarningsBreakdownCanonicalDescriptionEnumValues = []EarningsBreakdownCanonicalDescription{
	"BONUS",
	"COMMISSION",
	"OVERTIME",
	"PAID TIME OFF",
	"REGULAR PAY",
	"VACATION",
	"BASIC ALLOWANCE HOUSING",
	"BASIC ALLOWANCE SUBSISTENCE",
	"OTHER",
	"null",
}

func (v *EarningsBreakdownCanonicalDescription) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := EarningsBreakdownCanonicalDescription(value)


	*v = enumTypeValue
	return nil
}

// NewEarningsBreakdownCanonicalDescriptionFromValue returns a pointer to a valid EarningsBreakdownCanonicalDescription
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEarningsBreakdownCanonicalDescriptionFromValue(v string) (*EarningsBreakdownCanonicalDescription, error) {
	ev := EarningsBreakdownCanonicalDescription(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EarningsBreakdownCanonicalDescription) IsValid() bool {
	for _, existing := range allowedEarningsBreakdownCanonicalDescriptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EarningsBreakdownCanonicalDescription value
func (v EarningsBreakdownCanonicalDescription) Ptr() *EarningsBreakdownCanonicalDescription {
	return &v
}

type NullableEarningsBreakdownCanonicalDescription struct {
	value *EarningsBreakdownCanonicalDescription
	isSet bool
}

func (v NullableEarningsBreakdownCanonicalDescription) Get() *EarningsBreakdownCanonicalDescription {
	return v.value
}

func (v *NullableEarningsBreakdownCanonicalDescription) Set(val *EarningsBreakdownCanonicalDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableEarningsBreakdownCanonicalDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableEarningsBreakdownCanonicalDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarningsBreakdownCanonicalDescription(val *EarningsBreakdownCanonicalDescription) *NullableEarningsBreakdownCanonicalDescription {
	return &NullableEarningsBreakdownCanonicalDescription{value: val, isSet: true}
}

func (v NullableEarningsBreakdownCanonicalDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarningsBreakdownCanonicalDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

