/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.385.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PaymentInitiationPaymentCreateStatus For a payment returned by this endpoint, there is only one possible value:  `PAYMENT_STATUS_INPUT_NEEDED`: The initial phase of the payment
type PaymentInitiationPaymentCreateStatus string

var _ = fmt.Printf

// List of PaymentInitiationPaymentCreateStatus
const (
	PAYMENTINITIATIONPAYMENTCREATESTATUS_PAYMENT_STATUS_INPUT_NEEDED PaymentInitiationPaymentCreateStatus = "PAYMENT_STATUS_INPUT_NEEDED"
)

var allowedPaymentInitiationPaymentCreateStatusEnumValues = []PaymentInitiationPaymentCreateStatus{
	"PAYMENT_STATUS_INPUT_NEEDED",
}

func (v *PaymentInitiationPaymentCreateStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PaymentInitiationPaymentCreateStatus(value)


	*v = enumTypeValue
	return nil
}

// NewPaymentInitiationPaymentCreateStatusFromValue returns a pointer to a valid PaymentInitiationPaymentCreateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentInitiationPaymentCreateStatusFromValue(v string) (*PaymentInitiationPaymentCreateStatus, error) {
	ev := PaymentInitiationPaymentCreateStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentInitiationPaymentCreateStatus) IsValid() bool {
	for _, existing := range allowedPaymentInitiationPaymentCreateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentInitiationPaymentCreateStatus value
func (v PaymentInitiationPaymentCreateStatus) Ptr() *PaymentInitiationPaymentCreateStatus {
	return &v
}

type NullablePaymentInitiationPaymentCreateStatus struct {
	value *PaymentInitiationPaymentCreateStatus
	isSet bool
}

func (v NullablePaymentInitiationPaymentCreateStatus) Get() *PaymentInitiationPaymentCreateStatus {
	return v.value
}

func (v *NullablePaymentInitiationPaymentCreateStatus) Set(val *PaymentInitiationPaymentCreateStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationPaymentCreateStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationPaymentCreateStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationPaymentCreateStatus(val *PaymentInitiationPaymentCreateStatus) *NullablePaymentInitiationPaymentCreateStatus {
	return &NullablePaymentInitiationPaymentCreateStatus{value: val, isSet: true}
}

func (v NullablePaymentInitiationPaymentCreateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationPaymentCreateStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

