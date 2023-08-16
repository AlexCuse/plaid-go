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

// WalletTransactionFailureReason The error code of a failed transaction. Error codes include: `EXTERNAL_SYSTEM`: The transaction was declined by an external system. `EXPIRED`: The transaction request has expired. `CANCELLED`: The transaction request was rescinded. `INVALID`: The transaction did not meet certain criteria, such as an inactive account or no valid counterparty, etc. `UNKNOWN`: The transaction was unsuccessful, but the exact cause is unknown.
type WalletTransactionFailureReason string

var _ = fmt.Printf

// List of WalletTransactionFailureReason
const (
	WALLETTRANSACTIONFAILUREREASON_EXTERNAL_SYSTEM WalletTransactionFailureReason = "EXTERNAL_SYSTEM"
	WALLETTRANSACTIONFAILUREREASON_EXPIRED WalletTransactionFailureReason = "EXPIRED"
	WALLETTRANSACTIONFAILUREREASON_CANCELLED WalletTransactionFailureReason = "CANCELLED"
	WALLETTRANSACTIONFAILUREREASON_INVALID WalletTransactionFailureReason = "INVALID"
	WALLETTRANSACTIONFAILUREREASON_UNKNOWN WalletTransactionFailureReason = "UNKNOWN"
)

var allowedWalletTransactionFailureReasonEnumValues = []WalletTransactionFailureReason{
	"EXTERNAL_SYSTEM",
	"EXPIRED",
	"CANCELLED",
	"INVALID",
	"UNKNOWN",
}

func (v *WalletTransactionFailureReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WalletTransactionFailureReason(value)


	*v = enumTypeValue
	return nil
}

// NewWalletTransactionFailureReasonFromValue returns a pointer to a valid WalletTransactionFailureReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWalletTransactionFailureReasonFromValue(v string) (*WalletTransactionFailureReason, error) {
	ev := WalletTransactionFailureReason(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WalletTransactionFailureReason) IsValid() bool {
	for _, existing := range allowedWalletTransactionFailureReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WalletTransactionFailureReason value
func (v WalletTransactionFailureReason) Ptr() *WalletTransactionFailureReason {
	return &v
}

type NullableWalletTransactionFailureReason struct {
	value *WalletTransactionFailureReason
	isSet bool
}

func (v NullableWalletTransactionFailureReason) Get() *WalletTransactionFailureReason {
	return v.value
}

func (v *NullableWalletTransactionFailureReason) Set(val *WalletTransactionFailureReason) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionFailureReason) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionFailureReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionFailureReason(val *WalletTransactionFailureReason) *NullableWalletTransactionFailureReason {
	return &NullableWalletTransactionFailureReason{value: val, isSet: true}
}

func (v NullableWalletTransactionFailureReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionFailureReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

