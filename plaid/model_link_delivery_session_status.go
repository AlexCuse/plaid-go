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

// LinkDeliverySessionStatus The status of the given Hosted Link session.  `CREATED`: The session is created but not yet accessed by the user  `OPENED`: The session is opened by the user but not yet completed  `EXITED`: The session has been exited by the user  `COMPLETED`: The session has been completed by the user  `EXPIRED`: The session has expired
type LinkDeliverySessionStatus string

var _ = fmt.Printf

// List of LinkDeliverySessionStatus
const (
	LINKDELIVERYSESSIONSTATUS_CREATED LinkDeliverySessionStatus = "CREATED"
	LINKDELIVERYSESSIONSTATUS_OPENED LinkDeliverySessionStatus = "OPENED"
	LINKDELIVERYSESSIONSTATUS_EXITED LinkDeliverySessionStatus = "EXITED"
	LINKDELIVERYSESSIONSTATUS_COMPLETED LinkDeliverySessionStatus = "COMPLETED"
	LINKDELIVERYSESSIONSTATUS_EXPIRED LinkDeliverySessionStatus = "EXPIRED"
)

var allowedLinkDeliverySessionStatusEnumValues = []LinkDeliverySessionStatus{
	"CREATED",
	"OPENED",
	"EXITED",
	"COMPLETED",
	"EXPIRED",
}

func (v *LinkDeliverySessionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := LinkDeliverySessionStatus(value)


	*v = enumTypeValue
	return nil
}

// NewLinkDeliverySessionStatusFromValue returns a pointer to a valid LinkDeliverySessionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkDeliverySessionStatusFromValue(v string) (*LinkDeliverySessionStatus, error) {
	ev := LinkDeliverySessionStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkDeliverySessionStatus) IsValid() bool {
	for _, existing := range allowedLinkDeliverySessionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkDeliverySessionStatus value
func (v LinkDeliverySessionStatus) Ptr() *LinkDeliverySessionStatus {
	return &v
}

type NullableLinkDeliverySessionStatus struct {
	value *LinkDeliverySessionStatus
	isSet bool
}

func (v NullableLinkDeliverySessionStatus) Get() *LinkDeliverySessionStatus {
	return v.value
}

func (v *NullableLinkDeliverySessionStatus) Set(val *LinkDeliverySessionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliverySessionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliverySessionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliverySessionStatus(val *LinkDeliverySessionStatus) *NullableLinkDeliverySessionStatus {
	return &NullableLinkDeliverySessionStatus{value: val, isSet: true}
}

func (v NullableLinkDeliverySessionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliverySessionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

