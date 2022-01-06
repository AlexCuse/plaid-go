/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.61.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PaystubPayFrequency The frequency at which the employee is paid. Possible values: `MONTHLY`, `BI-WEEKLY`, `WEEKLY`, `SEMI-MONTHLY`.
type PaystubPayFrequency string

// List of PaystubPayFrequency
const (
	PAYSTUBPAYFREQUENCY_MONTHLY PaystubPayFrequency = "MONTHLY"
	PAYSTUBPAYFREQUENCY_BI_WEEKLY PaystubPayFrequency = "BI-WEEKLY"
	PAYSTUBPAYFREQUENCY_WEEKLY PaystubPayFrequency = "WEEKLY"
	PAYSTUBPAYFREQUENCY_SEMI_MONTHLY PaystubPayFrequency = "SEMI-MONTHLY"
	PAYSTUBPAYFREQUENCY_NULL PaystubPayFrequency = "null"
)

var allowedPaystubPayFrequencyEnumValues = []PaystubPayFrequency{
	"MONTHLY",
	"BI-WEEKLY",
	"WEEKLY",
	"SEMI-MONTHLY",
	"null",
}

func (v *PaystubPayFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaystubPayFrequency(value)
	for _, existing := range allowedPaystubPayFrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaystubPayFrequency", value)
}

// NewPaystubPayFrequencyFromValue returns a pointer to a valid PaystubPayFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaystubPayFrequencyFromValue(v string) (*PaystubPayFrequency, error) {
	ev := PaystubPayFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaystubPayFrequency: valid values are %v", v, allowedPaystubPayFrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaystubPayFrequency) IsValid() bool {
	for _, existing := range allowedPaystubPayFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaystubPayFrequency value
func (v PaystubPayFrequency) Ptr() *PaystubPayFrequency {
	return &v
}

type NullablePaystubPayFrequency struct {
	value *PaystubPayFrequency
	isSet bool
}

func (v NullablePaystubPayFrequency) Get() *PaystubPayFrequency {
	return v.value
}

func (v *NullablePaystubPayFrequency) Set(val *PaystubPayFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubPayFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubPayFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubPayFrequency(val *PaystubPayFrequency) *NullablePaystubPayFrequency {
	return &NullablePaystubPayFrequency{value: val, isSet: true}
}

func (v NullablePaystubPayFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubPayFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

