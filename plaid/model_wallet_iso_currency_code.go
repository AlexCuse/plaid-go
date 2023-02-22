/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// WalletISOCurrencyCode An ISO-4217 currency code, used with e-wallets and transactions.
type WalletISOCurrencyCode string

var _ = fmt.Printf

// List of WalletISOCurrencyCode
const (
	WALLETISOCURRENCYCODE_GBP WalletISOCurrencyCode = "GBP"
	WALLETISOCURRENCYCODE_EUR WalletISOCurrencyCode = "EUR"
)

var allowedWalletISOCurrencyCodeEnumValues = []WalletISOCurrencyCode{
	"GBP",
	"EUR",
}

func (v *WalletISOCurrencyCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WalletISOCurrencyCode(value)


	*v = enumTypeValue
	return nil
}

// NewWalletISOCurrencyCodeFromValue returns a pointer to a valid WalletISOCurrencyCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWalletISOCurrencyCodeFromValue(v string) (*WalletISOCurrencyCode, error) {
	ev := WalletISOCurrencyCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WalletISOCurrencyCode) IsValid() bool {
	for _, existing := range allowedWalletISOCurrencyCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WalletISOCurrencyCode value
func (v WalletISOCurrencyCode) Ptr() *WalletISOCurrencyCode {
	return &v
}

type NullableWalletISOCurrencyCode struct {
	value *WalletISOCurrencyCode
	isSet bool
}

func (v NullableWalletISOCurrencyCode) Get() *WalletISOCurrencyCode {
	return v.value
}

func (v *NullableWalletISOCurrencyCode) Set(val *WalletISOCurrencyCode) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletISOCurrencyCode) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletISOCurrencyCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletISOCurrencyCode(val *WalletISOCurrencyCode) *NullableWalletISOCurrencyCode {
	return &NullableWalletISOCurrencyCode{value: val, isSet: true}
}

func (v NullableWalletISOCurrencyCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletISOCurrencyCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

