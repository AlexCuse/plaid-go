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

// BeaconAuditTrailSource A type indicating what caused a resource to be changed or updated.   `dashboard` - The resource was created or updated by a member of your team via the Plaid dashboard.  `api` - The resource was created or updated via the Plaid API.  `system` - The resource was created or updated automatically by a part of the Plaid Beacon system. For example, if another business using Plaid Beacon created a fraud report that matched one of your users, your matching user's status would automatically be updated and the audit trail source would be `system`.  `bulk_import` - The resource was created or updated as part of a bulk import process. For example, if your company provided a CSV of user data as part of your initial onboarding, the audit trail source would be `bulk_import`.
type BeaconAuditTrailSource string

var _ = fmt.Printf

// List of BeaconAuditTrailSource
const (
	BEACONAUDITTRAILSOURCE_DASHBOARD BeaconAuditTrailSource = "dashboard"
	BEACONAUDITTRAILSOURCE_API BeaconAuditTrailSource = "api"
	BEACONAUDITTRAILSOURCE_SYSTEM BeaconAuditTrailSource = "system"
	BEACONAUDITTRAILSOURCE_BULK_IMPORT BeaconAuditTrailSource = "bulk_import"
)

var allowedBeaconAuditTrailSourceEnumValues = []BeaconAuditTrailSource{
	"dashboard",
	"api",
	"system",
	"bulk_import",
}

func (v *BeaconAuditTrailSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := BeaconAuditTrailSource(value)


	*v = enumTypeValue
	return nil
}

// NewBeaconAuditTrailSourceFromValue returns a pointer to a valid BeaconAuditTrailSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBeaconAuditTrailSourceFromValue(v string) (*BeaconAuditTrailSource, error) {
	ev := BeaconAuditTrailSource(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BeaconAuditTrailSource) IsValid() bool {
	for _, existing := range allowedBeaconAuditTrailSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BeaconAuditTrailSource value
func (v BeaconAuditTrailSource) Ptr() *BeaconAuditTrailSource {
	return &v
}

type NullableBeaconAuditTrailSource struct {
	value *BeaconAuditTrailSource
	isSet bool
}

func (v NullableBeaconAuditTrailSource) Get() *BeaconAuditTrailSource {
	return v.value
}

func (v *NullableBeaconAuditTrailSource) Set(val *BeaconAuditTrailSource) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconAuditTrailSource) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconAuditTrailSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconAuditTrailSource(val *BeaconAuditTrailSource) *NullableBeaconAuditTrailSource {
	return &NullableBeaconAuditTrailSource{value: val, isSet: true}
}

func (v NullableBeaconAuditTrailSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconAuditTrailSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

