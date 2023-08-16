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

// FDXNotificationType Type of Notification
type FDXNotificationType string

var _ = fmt.Printf

// List of FDXNotificationType
const (
	FDXNOTIFICATIONTYPE_CONSENT_REVOKED FDXNotificationType = "CONSENT_REVOKED"
	FDXNOTIFICATIONTYPE_CONSENT_UPDATED FDXNotificationType = "CONSENT_UPDATED"
	FDXNOTIFICATIONTYPE_CUSTOM FDXNotificationType = "CUSTOM"
	FDXNOTIFICATIONTYPE_SERVICE FDXNotificationType = "SERVICE"
	FDXNOTIFICATIONTYPE_BALANCE FDXNotificationType = "BALANCE"
	FDXNOTIFICATIONTYPE_PLANNED_OUTAGE FDXNotificationType = "PLANNED_OUTAGE"
)

var allowedFDXNotificationTypeEnumValues = []FDXNotificationType{
	"CONSENT_REVOKED",
	"CONSENT_UPDATED",
	"CUSTOM",
	"SERVICE",
	"BALANCE",
	"PLANNED_OUTAGE",
}

func (v *FDXNotificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := FDXNotificationType(value)


	*v = enumTypeValue
	return nil
}

// NewFDXNotificationTypeFromValue returns a pointer to a valid FDXNotificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFDXNotificationTypeFromValue(v string) (*FDXNotificationType, error) {
	ev := FDXNotificationType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FDXNotificationType) IsValid() bool {
	for _, existing := range allowedFDXNotificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FDXNotificationType value
func (v FDXNotificationType) Ptr() *FDXNotificationType {
	return &v
}

type NullableFDXNotificationType struct {
	value *FDXNotificationType
	isSet bool
}

func (v NullableFDXNotificationType) Get() *FDXNotificationType {
	return v.value
}

func (v *NullableFDXNotificationType) Set(val *FDXNotificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXNotificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXNotificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXNotificationType(val *FDXNotificationType) *NullableFDXNotificationType {
	return &NullableFDXNotificationType{value: val, isSet: true}
}

func (v NullableFDXNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXNotificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

