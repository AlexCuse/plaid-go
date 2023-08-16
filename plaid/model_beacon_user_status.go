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

// BeaconUserStatus A status of a Beacon User.  `rejected`: The Beacon User has been rejected for fraud. Users can be automatically or manually rejected.  `pending_review`: The Beacon User has been marked for review.  `cleared`: The Beacon User has been cleared of fraud.
type BeaconUserStatus string

var _ = fmt.Printf

// List of BeaconUserStatus
const (
	BEACONUSERSTATUS_REJECTED BeaconUserStatus = "rejected"
	BEACONUSERSTATUS_PENDING_REVIEW BeaconUserStatus = "pending_review"
	BEACONUSERSTATUS_CLEARED BeaconUserStatus = "cleared"
)

var allowedBeaconUserStatusEnumValues = []BeaconUserStatus{
	"rejected",
	"pending_review",
	"cleared",
}

func (v *BeaconUserStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := BeaconUserStatus(value)


	*v = enumTypeValue
	return nil
}

// NewBeaconUserStatusFromValue returns a pointer to a valid BeaconUserStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBeaconUserStatusFromValue(v string) (*BeaconUserStatus, error) {
	ev := BeaconUserStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BeaconUserStatus) IsValid() bool {
	for _, existing := range allowedBeaconUserStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BeaconUserStatus value
func (v BeaconUserStatus) Ptr() *BeaconUserStatus {
	return &v
}

type NullableBeaconUserStatus struct {
	value *BeaconUserStatus
	isSet bool
}

func (v NullableBeaconUserStatus) Get() *BeaconUserStatus {
	return v.value
}

func (v *NullableBeaconUserStatus) Set(val *BeaconUserStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconUserStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconUserStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconUserStatus(val *BeaconUserStatus) *NullableBeaconUserStatus {
	return &NullableBeaconUserStatus{value: val, isSet: true}
}

func (v NullableBeaconUserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconUserStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

