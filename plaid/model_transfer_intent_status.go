/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// TransferIntentStatus The status of the transfer intent.  - `PENDING` – The transfer intent is pending. - `SUCCEEDED` – The transfer intent was successfully created. - `FAILED` – The transfer intent was unable to be created.
type TransferIntentStatus string

// List of TransferIntentStatus
const (
	TRANSFERINTENTSTATUS_PENDING TransferIntentStatus = "PENDING"
	TRANSFERINTENTSTATUS_SUCCEEDED TransferIntentStatus = "SUCCEEDED"
	TRANSFERINTENTSTATUS_FAILED TransferIntentStatus = "FAILED"
)

var allowedTransferIntentStatusEnumValues = []TransferIntentStatus{
	"PENDING",
	"SUCCEEDED",
	"FAILED",
}

func (v *TransferIntentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferIntentStatus(value)
	for _, existing := range allowedTransferIntentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferIntentStatus", value)
}

// NewTransferIntentStatusFromValue returns a pointer to a valid TransferIntentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferIntentStatusFromValue(v string) (*TransferIntentStatus, error) {
	ev := TransferIntentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferIntentStatus: valid values are %v", v, allowedTransferIntentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferIntentStatus) IsValid() bool {
	for _, existing := range allowedTransferIntentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferIntentStatus value
func (v TransferIntentStatus) Ptr() *TransferIntentStatus {
	return &v
}

type NullableTransferIntentStatus struct {
	value *TransferIntentStatus
	isSet bool
}

func (v NullableTransferIntentStatus) Get() *TransferIntentStatus {
	return v.value
}

func (v *NullableTransferIntentStatus) Set(val *TransferIntentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentStatus(val *TransferIntentStatus) *NullableTransferIntentStatus {
	return &NullableTransferIntentStatus{value: val, isSet: true}
}

func (v NullableTransferIntentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

