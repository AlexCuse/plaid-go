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

// CounterpartyType The counterparty type.  `merchant`: a provider of goods or services for purchase `financial_institution`: a financial entity (bank, credit union, BNPL, fintech) `payment_app`: a transfer or P2P app (e.g. Zelle) `marketplace`: a marketplace (e.g DoorDash, Google Play Store) `payment_terminal`: a point-of-sale payment terminal (e.g Square, Toast) `income_source`: the payer in an income transaction (e.g., an employer, client, or government agency)
type CounterpartyType string

var _ = fmt.Printf

// List of CounterpartyType
const (
	COUNTERPARTYTYPE_MERCHANT CounterpartyType = "merchant"
	COUNTERPARTYTYPE_FINANCIAL_INSTITUTION CounterpartyType = "financial_institution"
	COUNTERPARTYTYPE_PAYMENT_APP CounterpartyType = "payment_app"
	COUNTERPARTYTYPE_MARKETPLACE CounterpartyType = "marketplace"
	COUNTERPARTYTYPE_PAYMENT_TERMINAL CounterpartyType = "payment_terminal"
	COUNTERPARTYTYPE_INCOME_SOURCE CounterpartyType = "income_source"
)

var allowedCounterpartyTypeEnumValues = []CounterpartyType{
	"merchant",
	"financial_institution",
	"payment_app",
	"marketplace",
	"payment_terminal",
	"income_source",
}

func (v *CounterpartyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CounterpartyType(value)


	*v = enumTypeValue
	return nil
}

// NewCounterpartyTypeFromValue returns a pointer to a valid CounterpartyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCounterpartyTypeFromValue(v string) (*CounterpartyType, error) {
	ev := CounterpartyType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CounterpartyType) IsValid() bool {
	for _, existing := range allowedCounterpartyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CounterpartyType value
func (v CounterpartyType) Ptr() *CounterpartyType {
	return &v
}

type NullableCounterpartyType struct {
	value *CounterpartyType
	isSet bool
}

func (v NullableCounterpartyType) Get() *CounterpartyType {
	return v.value
}

func (v *NullableCounterpartyType) Set(val *CounterpartyType) {
	v.value = val
	v.isSet = true
}

func (v NullableCounterpartyType) IsSet() bool {
	return v.isSet
}

func (v *NullableCounterpartyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCounterpartyType(val *CounterpartyType) *NullableCounterpartyType {
	return &NullableCounterpartyType{value: val, isSet: true}
}

func (v NullableCounterpartyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCounterpartyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

