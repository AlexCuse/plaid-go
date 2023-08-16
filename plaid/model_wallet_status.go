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

// WalletStatus The status of the wallet.  `UNKNOWN`: The wallet status is unknown.  `ACTIVE`: The wallet is active and ready to send money to and receive money from.  `CLOSED`: The wallet is closed. Any transactions made to or from this wallet will error.
type WalletStatus string

var _ = fmt.Printf

// List of WalletStatus
const (
	WALLETSTATUS_UNKNOWN WalletStatus = "UNKNOWN"
	WALLETSTATUS_ACTIVE WalletStatus = "ACTIVE"
	WALLETSTATUS_CLOSED WalletStatus = "CLOSED"
)

var allowedWalletStatusEnumValues = []WalletStatus{
	"UNKNOWN",
	"ACTIVE",
	"CLOSED",
}

func (v *WalletStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WalletStatus(value)


	*v = enumTypeValue
	return nil
}

// NewWalletStatusFromValue returns a pointer to a valid WalletStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWalletStatusFromValue(v string) (*WalletStatus, error) {
	ev := WalletStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WalletStatus) IsValid() bool {
	for _, existing := range allowedWalletStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WalletStatus value
func (v WalletStatus) Ptr() *WalletStatus {
	return &v
}

type NullableWalletStatus struct {
	value *WalletStatus
	isSet bool
}

func (v NullableWalletStatus) Get() *WalletStatus {
	return v.value
}

func (v *NullableWalletStatus) Set(val *WalletStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletStatus(val *WalletStatus) *NullableWalletStatus {
	return &NullableWalletStatus{value: val, isSet: true}
}

func (v NullableWalletStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

