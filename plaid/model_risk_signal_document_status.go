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

// RiskSignalDocumentStatus Status of a document for risk signal analysis
type RiskSignalDocumentStatus string

var _ = fmt.Printf

// List of RiskSignalDocumentStatus
const (
	RISKSIGNALDOCUMENTSTATUS_PROCESSING RiskSignalDocumentStatus = "PROCESSING"
	RISKSIGNALDOCUMENTSTATUS_PROCESSING_COMPLETE RiskSignalDocumentStatus = "PROCESSING_COMPLETE"
	RISKSIGNALDOCUMENTSTATUS_PROCESSING_ERROR RiskSignalDocumentStatus = "PROCESSING_ERROR"
	RISKSIGNALDOCUMENTSTATUS_PASSWORD_PROTECTED RiskSignalDocumentStatus = "PASSWORD_PROTECTED"
	RISKSIGNALDOCUMENTSTATUS_VIRUS_DETECTED RiskSignalDocumentStatus = "VIRUS_DETECTED"
)

var allowedRiskSignalDocumentStatusEnumValues = []RiskSignalDocumentStatus{
	"PROCESSING",
	"PROCESSING_COMPLETE",
	"PROCESSING_ERROR",
	"PASSWORD_PROTECTED",
	"VIRUS_DETECTED",
}

func (v *RiskSignalDocumentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := RiskSignalDocumentStatus(value)


	*v = enumTypeValue
	return nil
}

// NewRiskSignalDocumentStatusFromValue returns a pointer to a valid RiskSignalDocumentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskSignalDocumentStatusFromValue(v string) (*RiskSignalDocumentStatus, error) {
	ev := RiskSignalDocumentStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskSignalDocumentStatus) IsValid() bool {
	for _, existing := range allowedRiskSignalDocumentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RiskSignalDocumentStatus value
func (v RiskSignalDocumentStatus) Ptr() *RiskSignalDocumentStatus {
	return &v
}

type NullableRiskSignalDocumentStatus struct {
	value *RiskSignalDocumentStatus
	isSet bool
}

func (v NullableRiskSignalDocumentStatus) Get() *RiskSignalDocumentStatus {
	return v.value
}

func (v *NullableRiskSignalDocumentStatus) Set(val *RiskSignalDocumentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskSignalDocumentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskSignalDocumentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskSignalDocumentStatus(val *RiskSignalDocumentStatus) *NullableRiskSignalDocumentStatus {
	return &NullableRiskSignalDocumentStatus{value: val, isSet: true}
}

func (v NullableRiskSignalDocumentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskSignalDocumentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
