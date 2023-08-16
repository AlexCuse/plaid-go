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

// PaymentScheme Payment scheme. If not specified - the default in the region will be used (e.g. `SEPA_CREDIT_TRANSFER` for EU). Using unsupported values will result in a failed payment.  `LOCAL_DEFAULT`: The default payment scheme for the selected market and currency will be used.  `LOCAL_INSTANT`: The instant payment scheme for the selected market and currency will be used (if applicable). Fees may be applied by the institution. If the market does not support an Instant Scheme (e.g. Denmark), the default in the region will be used.  `SEPA_CREDIT_TRANSFER`: The standard payment to a beneficiary within the SEPA area.  `SEPA_CREDIT_TRANSFER_INSTANT`: Instant payment within the SEPA area. May involve additional fees and may not be available at some banks.
type PaymentScheme string

var _ = fmt.Printf

// List of PaymentScheme
const (
	PAYMENTSCHEME_NULL PaymentScheme = "null"
	PAYMENTSCHEME_LOCAL_DEFAULT PaymentScheme = "LOCAL_DEFAULT"
	PAYMENTSCHEME_LOCAL_INSTANT PaymentScheme = "LOCAL_INSTANT"
	PAYMENTSCHEME_SEPA_CREDIT_TRANSFER PaymentScheme = "SEPA_CREDIT_TRANSFER"
	PAYMENTSCHEME_SEPA_CREDIT_TRANSFER_INSTANT PaymentScheme = "SEPA_CREDIT_TRANSFER_INSTANT"
)

var allowedPaymentSchemeEnumValues = []PaymentScheme{
	"null",
	"LOCAL_DEFAULT",
	"LOCAL_INSTANT",
	"SEPA_CREDIT_TRANSFER",
	"SEPA_CREDIT_TRANSFER_INSTANT",
}

func (v *PaymentScheme) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PaymentScheme(value)


	*v = enumTypeValue
	return nil
}

// NewPaymentSchemeFromValue returns a pointer to a valid PaymentScheme
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentSchemeFromValue(v string) (*PaymentScheme, error) {
	ev := PaymentScheme(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentScheme) IsValid() bool {
	for _, existing := range allowedPaymentSchemeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentScheme value
func (v PaymentScheme) Ptr() *PaymentScheme {
	return &v
}

type NullablePaymentScheme struct {
	value *PaymentScheme
	isSet bool
}

func (v NullablePaymentScheme) Get() *PaymentScheme {
	return v.value
}

func (v *NullablePaymentScheme) Set(val *PaymentScheme) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentScheme) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentScheme(val *PaymentScheme) *NullablePaymentScheme {
	return &NullablePaymentScheme{value: val, isSet: true}
}

func (v NullablePaymentScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

