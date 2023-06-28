/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.385.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// LinkDeliveryVerificationStatus Indicates an Item's micro-deposit-based verification status.
type LinkDeliveryVerificationStatus string

var _ = fmt.Printf

// List of LinkDeliveryVerificationStatus
const (
	LINKDELIVERYVERIFICATIONSTATUS_AUTOMATICALLY_VERIFIED LinkDeliveryVerificationStatus = "automatically_verified"
	LINKDELIVERYVERIFICATIONSTATUS_PENDING_AUTOMATIC_VERIFICATION LinkDeliveryVerificationStatus = "pending_automatic_verification"
	LINKDELIVERYVERIFICATIONSTATUS_PENDING_MANUAL_VERIFICATION LinkDeliveryVerificationStatus = "pending_manual_verification"
	LINKDELIVERYVERIFICATIONSTATUS_MANUALLY_VERIFIED LinkDeliveryVerificationStatus = "manually_verified"
	LINKDELIVERYVERIFICATIONSTATUS_VERIFICATION_EXPIRED LinkDeliveryVerificationStatus = "verification_expired"
	LINKDELIVERYVERIFICATIONSTATUS_VERIFICATION_FAILED LinkDeliveryVerificationStatus = "verification_failed"
)

var allowedLinkDeliveryVerificationStatusEnumValues = []LinkDeliveryVerificationStatus{
	"automatically_verified",
	"pending_automatic_verification",
	"pending_manual_verification",
	"manually_verified",
	"verification_expired",
	"verification_failed",
}

func (v *LinkDeliveryVerificationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := LinkDeliveryVerificationStatus(value)


	*v = enumTypeValue
	return nil
}

// NewLinkDeliveryVerificationStatusFromValue returns a pointer to a valid LinkDeliveryVerificationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkDeliveryVerificationStatusFromValue(v string) (*LinkDeliveryVerificationStatus, error) {
	ev := LinkDeliveryVerificationStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkDeliveryVerificationStatus) IsValid() bool {
	for _, existing := range allowedLinkDeliveryVerificationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkDeliveryVerificationStatus value
func (v LinkDeliveryVerificationStatus) Ptr() *LinkDeliveryVerificationStatus {
	return &v
}

type NullableLinkDeliveryVerificationStatus struct {
	value *LinkDeliveryVerificationStatus
	isSet bool
}

func (v NullableLinkDeliveryVerificationStatus) Get() *LinkDeliveryVerificationStatus {
	return v.value
}

func (v *NullableLinkDeliveryVerificationStatus) Set(val *LinkDeliveryVerificationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryVerificationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryVerificationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryVerificationStatus(val *LinkDeliveryVerificationStatus) *NullableLinkDeliveryVerificationStatus {
	return &NullableLinkDeliveryVerificationStatus{value: val, isSet: true}
}

func (v NullableLinkDeliveryVerificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryVerificationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

