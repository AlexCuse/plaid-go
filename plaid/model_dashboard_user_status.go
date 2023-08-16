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

// DashboardUserStatus The current status of the user.
type DashboardUserStatus string

var _ = fmt.Printf

// List of DashboardUserStatus
const (
	DASHBOARDUSERSTATUS_INVITED DashboardUserStatus = "invited"
	DASHBOARDUSERSTATUS_ACTIVE DashboardUserStatus = "active"
	DASHBOARDUSERSTATUS_DEACTIVATED DashboardUserStatus = "deactivated"
)

var allowedDashboardUserStatusEnumValues = []DashboardUserStatus{
	"invited",
	"active",
	"deactivated",
}

func (v *DashboardUserStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := DashboardUserStatus(value)


	*v = enumTypeValue
	return nil
}

// NewDashboardUserStatusFromValue returns a pointer to a valid DashboardUserStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDashboardUserStatusFromValue(v string) (*DashboardUserStatus, error) {
	ev := DashboardUserStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DashboardUserStatus) IsValid() bool {
	for _, existing := range allowedDashboardUserStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardUserStatus value
func (v DashboardUserStatus) Ptr() *DashboardUserStatus {
	return &v
}

type NullableDashboardUserStatus struct {
	value *DashboardUserStatus
	isSet bool
}

func (v NullableDashboardUserStatus) Get() *DashboardUserStatus {
	return v.value
}

func (v *NullableDashboardUserStatus) Set(val *DashboardUserStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardUserStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardUserStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardUserStatus(val *DashboardUserStatus) *NullableDashboardUserStatus {
	return &NullableDashboardUserStatus{value: val, isSet: true}
}

func (v NullableDashboardUserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardUserStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

