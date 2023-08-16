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

// TransferBalanceType The type of balance.  `prefunded_rtp_credits` - Your prefunded RTP credit balance with Plaid `prefunded_ach_credits` - Your prefunded ACH credit balance with Plaid
type TransferBalanceType string

var _ = fmt.Printf

// List of TransferBalanceType
const (
	TRANSFERBALANCETYPE_RTP_CREDITS TransferBalanceType = "prefunded_rtp_credits"
	TRANSFERBALANCETYPE_ACH_CREDITS TransferBalanceType = "prefunded_ach_credits"
)

var allowedTransferBalanceTypeEnumValues = []TransferBalanceType{
	"prefunded_rtp_credits",
	"prefunded_ach_credits",
}

func (v *TransferBalanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferBalanceType(value)


	*v = enumTypeValue
	return nil
}

// NewTransferBalanceTypeFromValue returns a pointer to a valid TransferBalanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferBalanceTypeFromValue(v string) (*TransferBalanceType, error) {
	ev := TransferBalanceType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferBalanceType) IsValid() bool {
	for _, existing := range allowedTransferBalanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferBalanceType value
func (v TransferBalanceType) Ptr() *TransferBalanceType {
	return &v
}

type NullableTransferBalanceType struct {
	value *TransferBalanceType
	isSet bool
}

func (v NullableTransferBalanceType) Get() *TransferBalanceType {
	return v.value
}

func (v *NullableTransferBalanceType) Set(val *TransferBalanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferBalanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferBalanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferBalanceType(val *TransferBalanceType) *NullableTransferBalanceType {
	return &NullableTransferBalanceType{value: val, isSet: true}
}

func (v NullableTransferBalanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferBalanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

