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

// TransferIntentCreateMode The direction of the flow of transfer funds.  `PAYMENT`: Transfers funds from an end user's account to your business account.  `DISBURSEMENT`: Transfers funds from your business account to an end user's account.
type TransferIntentCreateMode string

var _ = fmt.Printf

// List of TransferIntentCreateMode
const (
	TRANSFERINTENTCREATEMODE_PAYMENT TransferIntentCreateMode = "PAYMENT"
	TRANSFERINTENTCREATEMODE_DISBURSEMENT TransferIntentCreateMode = "DISBURSEMENT"
)

var allowedTransferIntentCreateModeEnumValues = []TransferIntentCreateMode{
	"PAYMENT",
	"DISBURSEMENT",
}

func (v *TransferIntentCreateMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferIntentCreateMode(value)


	*v = enumTypeValue
	return nil
}

// NewTransferIntentCreateModeFromValue returns a pointer to a valid TransferIntentCreateMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferIntentCreateModeFromValue(v string) (*TransferIntentCreateMode, error) {
	ev := TransferIntentCreateMode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferIntentCreateMode) IsValid() bool {
	for _, existing := range allowedTransferIntentCreateModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferIntentCreateMode value
func (v TransferIntentCreateMode) Ptr() *TransferIntentCreateMode {
	return &v
}

type NullableTransferIntentCreateMode struct {
	value *TransferIntentCreateMode
	isSet bool
}

func (v NullableTransferIntentCreateMode) Get() *TransferIntentCreateMode {
	return v.value
}

func (v *NullableTransferIntentCreateMode) Set(val *TransferIntentCreateMode) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentCreateMode) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentCreateMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentCreateMode(val *TransferIntentCreateMode) *NullableTransferIntentCreateMode {
	return &NullableTransferIntentCreateMode{value: val, isSet: true}
}

func (v NullableTransferIntentCreateMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentCreateMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
