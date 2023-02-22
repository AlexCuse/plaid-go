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

// BankTransferNetwork The network or rails used for the transfer. Valid options are `ach`, `same-day-ach`, or `wire`.
type BankTransferNetwork string

var _ = fmt.Printf

// List of BankTransferNetwork
const (
	BANKTRANSFERNETWORK_ACH BankTransferNetwork = "ach"
	BANKTRANSFERNETWORK_SAME_DAY_ACH BankTransferNetwork = "same-day-ach"
	BANKTRANSFERNETWORK_WIRE BankTransferNetwork = "wire"
)

var allowedBankTransferNetworkEnumValues = []BankTransferNetwork{
	"ach",
	"same-day-ach",
	"wire",
}

func (v *BankTransferNetwork) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := BankTransferNetwork(value)


	*v = enumTypeValue
	return nil
}

// NewBankTransferNetworkFromValue returns a pointer to a valid BankTransferNetwork
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBankTransferNetworkFromValue(v string) (*BankTransferNetwork, error) {
	ev := BankTransferNetwork(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankTransferNetwork) IsValid() bool {
	for _, existing := range allowedBankTransferNetworkEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BankTransferNetwork value
func (v BankTransferNetwork) Ptr() *BankTransferNetwork {
	return &v
}

type NullableBankTransferNetwork struct {
	value *BankTransferNetwork
	isSet bool
}

func (v NullableBankTransferNetwork) Get() *BankTransferNetwork {
	return v.value
}

func (v *NullableBankTransferNetwork) Set(val *BankTransferNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferNetwork(val *BankTransferNetwork) *NullableBankTransferNetwork {
	return &NullableBankTransferNetwork{value: val, isSet: true}
}

func (v NullableBankTransferNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

