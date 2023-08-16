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

// DocumentStatus An outcome status for this specific document submission. Distinct from the overall `documentary_verification.status` that summarizes the verification outcome from one or more documents.
type DocumentStatus string

var _ = fmt.Printf

// List of DocumentStatus
const (
	DOCUMENTSTATUS_SUCCESS DocumentStatus = "success"
	DOCUMENTSTATUS_FAILED DocumentStatus = "failed"
	DOCUMENTSTATUS_MANUALLY_APPROVED DocumentStatus = "manually_approved"
)

var allowedDocumentStatusEnumValues = []DocumentStatus{
	"success",
	"failed",
	"manually_approved",
}

func (v *DocumentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := DocumentStatus(value)


	*v = enumTypeValue
	return nil
}

// NewDocumentStatusFromValue returns a pointer to a valid DocumentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentStatusFromValue(v string) (*DocumentStatus, error) {
	ev := DocumentStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentStatus) IsValid() bool {
	for _, existing := range allowedDocumentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentStatus value
func (v DocumentStatus) Ptr() *DocumentStatus {
	return &v
}

type NullableDocumentStatus struct {
	value *DocumentStatus
	isSet bool
}

func (v NullableDocumentStatus) Get() *DocumentStatus {
	return v.value
}

func (v *NullableDocumentStatus) Set(val *DocumentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentStatus(val *DocumentStatus) *NullableDocumentStatus {
	return &NullableDocumentStatus{value: val, isSet: true}
}

func (v NullableDocumentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

