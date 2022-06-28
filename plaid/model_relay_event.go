/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.131.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// RelayEvent The webhook code indicating which endpoint was called. It can be one of `GET_CALLED`, `REFRESH_CALLED` or `AUDIT_COPY_CREATE_CALLED`.
type RelayEvent string

// List of RelayEvent
const (
	RELAYEVENT_GET_CALLED RelayEvent = "GET_CALLED"
	RELAYEVENT_REFRESH_CALLED RelayEvent = "REFRESH_CALLED"
	RELAYEVENT_AUDIT_COPY_CREATE_CALLED RelayEvent = "AUDIT_COPY_CREATE_CALLED"
)

var allowedRelayEventEnumValues = []RelayEvent{
	"GET_CALLED",
	"REFRESH_CALLED",
	"AUDIT_COPY_CREATE_CALLED",
}

func (v *RelayEvent) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RelayEvent(value)
	for _, existing := range allowedRelayEventEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RelayEvent", value)
}

// NewRelayEventFromValue returns a pointer to a valid RelayEvent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRelayEventFromValue(v string) (*RelayEvent, error) {
	ev := RelayEvent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RelayEvent: valid values are %v", v, allowedRelayEventEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RelayEvent) IsValid() bool {
	for _, existing := range allowedRelayEventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RelayEvent value
func (v RelayEvent) Ptr() *RelayEvent {
	return &v
}

type NullableRelayEvent struct {
	value *RelayEvent
	isSet bool
}

func (v NullableRelayEvent) Get() *RelayEvent {
	return v.value
}

func (v *NullableRelayEvent) Set(val *RelayEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableRelayEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableRelayEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelayEvent(val *RelayEvent) *NullableRelayEvent {
	return &NullableRelayEvent{value: val, isSet: true}
}

func (v NullableRelayEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelayEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

