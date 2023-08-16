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

// SignalPaymentMethod The payment method to complete the transaction after the risk assessment. It may be different from the default payment method.  `SAME_DAY_ACH`: Same Day ACH by NACHA. The debit transaction is processed and settled on the same day  `NEXT_DAY_ACH`: Next Day ACH settlement for debit transactions, offered by some payment processors  `STANDARD_ACH`: standard ACH by NACHA  `REAL_TIME_PAYMENTS`: real-time payments such as RTP and FedNow  `DEBIT_CARD`: if the default payment is over debit card networks  `MULTIPLE_PAYMENT_METHODS`: if there is no default debit rail or there are multiple payment methods  Possible values: `SAME_DAY_ACH`, `NEXT_DAY_ACH`, `STANDARD_ACH`, `REAL_TIME_PAYMENTS`, `DEBIT_CARD`, `MULTIPLE_PAYMENT_METHODS` 
type SignalPaymentMethod string

var _ = fmt.Printf

// List of SignalPaymentMethod
const (
	SIGNALPAYMENTMETHOD_SAME_DAY_ACH SignalPaymentMethod = "SAME_DAY_ACH"
	SIGNALPAYMENTMETHOD_NEXT_DAY_ACH SignalPaymentMethod = "NEXT_DAY_ACH"
	SIGNALPAYMENTMETHOD_STANDARD_ACH SignalPaymentMethod = "STANDARD_ACH"
	SIGNALPAYMENTMETHOD_REAL_TIME_PAYMENTS SignalPaymentMethod = "REAL_TIME_PAYMENTS"
	SIGNALPAYMENTMETHOD_DEBIT_CARD SignalPaymentMethod = "DEBIT_CARD"
	SIGNALPAYMENTMETHOD_MULTIPLE_PAYMENT_METHODS SignalPaymentMethod = "MULTIPLE_PAYMENT_METHODS"
)

var allowedSignalPaymentMethodEnumValues = []SignalPaymentMethod{
	"SAME_DAY_ACH",
	"NEXT_DAY_ACH",
	"STANDARD_ACH",
	"REAL_TIME_PAYMENTS",
	"DEBIT_CARD",
	"MULTIPLE_PAYMENT_METHODS",
}

func (v *SignalPaymentMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := SignalPaymentMethod(value)


	*v = enumTypeValue
	return nil
}

// NewSignalPaymentMethodFromValue returns a pointer to a valid SignalPaymentMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignalPaymentMethodFromValue(v string) (*SignalPaymentMethod, error) {
	ev := SignalPaymentMethod(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignalPaymentMethod) IsValid() bool {
	for _, existing := range allowedSignalPaymentMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignalPaymentMethod value
func (v SignalPaymentMethod) Ptr() *SignalPaymentMethod {
	return &v
}

type NullableSignalPaymentMethod struct {
	value *SignalPaymentMethod
	isSet bool
}

func (v NullableSignalPaymentMethod) Get() *SignalPaymentMethod {
	return v.value
}

func (v *NullableSignalPaymentMethod) Set(val *SignalPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalPaymentMethod(val *SignalPaymentMethod) *NullableSignalPaymentMethod {
	return &NullableSignalPaymentMethod{value: val, isSet: true}
}

func (v NullableSignalPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

