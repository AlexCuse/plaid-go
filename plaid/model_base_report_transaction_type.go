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

// BaseReportTransactionType `digital:` transactions that took place online.  `place:` transactions that were made at a physical location.  `special:` transactions that relate to banks, e.g. fees or deposits.  `unresolved:` transactions that do not fit into the other three types. 
type BaseReportTransactionType string

var _ = fmt.Printf

// List of BaseReportTransactionType
const (
	BASEREPORTTRANSACTIONTYPE_DIGITAL BaseReportTransactionType = "digital"
	BASEREPORTTRANSACTIONTYPE_PLACE BaseReportTransactionType = "place"
	BASEREPORTTRANSACTIONTYPE_SPECIAL BaseReportTransactionType = "special"
	BASEREPORTTRANSACTIONTYPE_UNRESOLVED BaseReportTransactionType = "unresolved"
)

var allowedBaseReportTransactionTypeEnumValues = []BaseReportTransactionType{
	"digital",
	"place",
	"special",
	"unresolved",
}

func (v *BaseReportTransactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := BaseReportTransactionType(value)


	*v = enumTypeValue
	return nil
}

// NewBaseReportTransactionTypeFromValue returns a pointer to a valid BaseReportTransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBaseReportTransactionTypeFromValue(v string) (*BaseReportTransactionType, error) {
	ev := BaseReportTransactionType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BaseReportTransactionType) IsValid() bool {
	for _, existing := range allowedBaseReportTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BaseReportTransactionType value
func (v BaseReportTransactionType) Ptr() *BaseReportTransactionType {
	return &v
}

type NullableBaseReportTransactionType struct {
	value *BaseReportTransactionType
	isSet bool
}

func (v NullableBaseReportTransactionType) Get() *BaseReportTransactionType {
	return v.value
}

func (v *NullableBaseReportTransactionType) Set(val *BaseReportTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReportTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReportTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReportTransactionType(val *BaseReportTransactionType) *NullableBaseReportTransactionType {
	return &NullableBaseReportTransactionType{value: val, isSet: true}
}

func (v NullableBaseReportTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReportTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

