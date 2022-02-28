/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.77.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PaymentConsentPeriodicInterval Payment consent periodic interval.
type PaymentConsentPeriodicInterval string

// List of PaymentConsentPeriodicInterval
const (
	PAYMENTCONSENTPERIODICINTERVAL_DAY PaymentConsentPeriodicInterval = "DAY"
	PAYMENTCONSENTPERIODICINTERVAL_WEEK PaymentConsentPeriodicInterval = "WEEK"
	PAYMENTCONSENTPERIODICINTERVAL_MONTH PaymentConsentPeriodicInterval = "MONTH"
	PAYMENTCONSENTPERIODICINTERVAL_YEAR PaymentConsentPeriodicInterval = "YEAR"
)

var allowedPaymentConsentPeriodicIntervalEnumValues = []PaymentConsentPeriodicInterval{
	"DAY",
	"WEEK",
	"MONTH",
	"YEAR",
}

func (v *PaymentConsentPeriodicInterval) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentConsentPeriodicInterval(value)
	for _, existing := range allowedPaymentConsentPeriodicIntervalEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentConsentPeriodicInterval", value)
}

// NewPaymentConsentPeriodicIntervalFromValue returns a pointer to a valid PaymentConsentPeriodicInterval
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentConsentPeriodicIntervalFromValue(v string) (*PaymentConsentPeriodicInterval, error) {
	ev := PaymentConsentPeriodicInterval(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentConsentPeriodicInterval: valid values are %v", v, allowedPaymentConsentPeriodicIntervalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentConsentPeriodicInterval) IsValid() bool {
	for _, existing := range allowedPaymentConsentPeriodicIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentConsentPeriodicInterval value
func (v PaymentConsentPeriodicInterval) Ptr() *PaymentConsentPeriodicInterval {
	return &v
}

type NullablePaymentConsentPeriodicInterval struct {
	value *PaymentConsentPeriodicInterval
	isSet bool
}

func (v NullablePaymentConsentPeriodicInterval) Get() *PaymentConsentPeriodicInterval {
	return v.value
}

func (v *NullablePaymentConsentPeriodicInterval) Set(val *PaymentConsentPeriodicInterval) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentConsentPeriodicInterval) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentConsentPeriodicInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentConsentPeriodicInterval(val *PaymentConsentPeriodicInterval) *NullablePaymentConsentPeriodicInterval {
	return &NullablePaymentConsentPeriodicInterval{value: val, isSet: true}
}

func (v NullablePaymentConsentPeriodicInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentConsentPeriodicInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

