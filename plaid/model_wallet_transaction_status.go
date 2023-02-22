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

// WalletTransactionStatus The status of the transaction.  `AUTHORISING`: The transaction is being processed for validation and compliance.  `INITIATED`: The transaction has been initiated and is currently being processed.  `EXECUTED`: The transaction has been successfully executed and is considered complete. This is only applicable for debit transactions.  `SETTLED`: The transaction has settled and funds are available for use. This is only applicable for credit transactions. A transaction will typically settle within seconds to several days, depending on which payment rail is used.  `FAILED`: The transaction failed to process successfully. This is a terminal status.  `BLOCKED`: The transaction has been blocked for violating compliance rules. This is a terminal status.
type WalletTransactionStatus string

var _ = fmt.Printf

// List of WalletTransactionStatus
const (
	WALLETTRANSACTIONSTATUS_AUTHORISING WalletTransactionStatus = "AUTHORISING"
	WALLETTRANSACTIONSTATUS_INITIATED WalletTransactionStatus = "INITIATED"
	WALLETTRANSACTIONSTATUS_EXECUTED WalletTransactionStatus = "EXECUTED"
	WALLETTRANSACTIONSTATUS_SETTLED WalletTransactionStatus = "SETTLED"
	WALLETTRANSACTIONSTATUS_BLOCKED WalletTransactionStatus = "BLOCKED"
	WALLETTRANSACTIONSTATUS_FAILED WalletTransactionStatus = "FAILED"
)

var allowedWalletTransactionStatusEnumValues = []WalletTransactionStatus{
	"AUTHORISING",
	"INITIATED",
	"EXECUTED",
	"SETTLED",
	"BLOCKED",
	"FAILED",
}

func (v *WalletTransactionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WalletTransactionStatus(value)


	*v = enumTypeValue
	return nil
}

// NewWalletTransactionStatusFromValue returns a pointer to a valid WalletTransactionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWalletTransactionStatusFromValue(v string) (*WalletTransactionStatus, error) {
	ev := WalletTransactionStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WalletTransactionStatus) IsValid() bool {
	for _, existing := range allowedWalletTransactionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WalletTransactionStatus value
func (v WalletTransactionStatus) Ptr() *WalletTransactionStatus {
	return &v
}

type NullableWalletTransactionStatus struct {
	value *WalletTransactionStatus
	isSet bool
}

func (v NullableWalletTransactionStatus) Get() *WalletTransactionStatus {
	return v.value
}

func (v *NullableWalletTransactionStatus) Set(val *WalletTransactionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionStatus(val *WalletTransactionStatus) *NullableWalletTransactionStatus {
	return &NullableWalletTransactionStatus{value: val, isSet: true}
}

func (v NullableWalletTransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

