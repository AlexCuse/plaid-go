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

// IdentityVerificationStepStatus The status of a step in the identity verification process.
type IdentityVerificationStepStatus string

var _ = fmt.Printf

// List of IdentityVerificationStepStatus
const (
	IDENTITYVERIFICATIONSTEPSTATUS_SUCCESS IdentityVerificationStepStatus = "success"
	IDENTITYVERIFICATIONSTEPSTATUS_ACTIVE IdentityVerificationStepStatus = "active"
	IDENTITYVERIFICATIONSTEPSTATUS_FAILED IdentityVerificationStepStatus = "failed"
	IDENTITYVERIFICATIONSTEPSTATUS_WAITING_FOR_PREREQUISITE IdentityVerificationStepStatus = "waiting_for_prerequisite"
	IDENTITYVERIFICATIONSTEPSTATUS_NOT_APPLICABLE IdentityVerificationStepStatus = "not_applicable"
	IDENTITYVERIFICATIONSTEPSTATUS_SKIPPED IdentityVerificationStepStatus = "skipped"
	IDENTITYVERIFICATIONSTEPSTATUS_EXPIRED IdentityVerificationStepStatus = "expired"
	IDENTITYVERIFICATIONSTEPSTATUS_CANCELED IdentityVerificationStepStatus = "canceled"
	IDENTITYVERIFICATIONSTEPSTATUS_PENDING_REVIEW IdentityVerificationStepStatus = "pending_review"
	IDENTITYVERIFICATIONSTEPSTATUS_MANUALLY_APPROVED IdentityVerificationStepStatus = "manually_approved"
	IDENTITYVERIFICATIONSTEPSTATUS_MANUALLY_REJECTED IdentityVerificationStepStatus = "manually_rejected"
)

var allowedIdentityVerificationStepStatusEnumValues = []IdentityVerificationStepStatus{
	"success",
	"active",
	"failed",
	"waiting_for_prerequisite",
	"not_applicable",
	"skipped",
	"expired",
	"canceled",
	"pending_review",
	"manually_approved",
	"manually_rejected",
}

func (v *IdentityVerificationStepStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := IdentityVerificationStepStatus(value)


	*v = enumTypeValue
	return nil
}

// NewIdentityVerificationStepStatusFromValue returns a pointer to a valid IdentityVerificationStepStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdentityVerificationStepStatusFromValue(v string) (*IdentityVerificationStepStatus, error) {
	ev := IdentityVerificationStepStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdentityVerificationStepStatus) IsValid() bool {
	for _, existing := range allowedIdentityVerificationStepStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdentityVerificationStepStatus value
func (v IdentityVerificationStepStatus) Ptr() *IdentityVerificationStepStatus {
	return &v
}

type NullableIdentityVerificationStepStatus struct {
	value *IdentityVerificationStepStatus
	isSet bool
}

func (v NullableIdentityVerificationStepStatus) Get() *IdentityVerificationStepStatus {
	return v.value
}

func (v *NullableIdentityVerificationStepStatus) Set(val *IdentityVerificationStepStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationStepStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationStepStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationStepStatus(val *IdentityVerificationStepStatus) *NullableIdentityVerificationStepStatus {
	return &NullableIdentityVerificationStepStatus{value: val, isSet: true}
}

func (v NullableIdentityVerificationStepStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationStepStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

