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

// ActivityType Types of consent activities
type ActivityType string

var _ = fmt.Printf

// List of ActivityType
const (
	ACTIVITYTYPE_UNKNOWN ActivityType = "UNKNOWN"
	ACTIVITYTYPE_ITEM_CREATE ActivityType = "ITEM_CREATE"
	ACTIVITYTYPE_ITEM_IMPORT ActivityType = "ITEM_IMPORT"
	ACTIVITYTYPE_ITEM_UPDATE ActivityType = "ITEM_UPDATE"
	ACTIVITYTYPE_ITEM_UNLINK ActivityType = "ITEM_UNLINK"
	ACTIVITYTYPE_PORTAL_UNLINK ActivityType = "PORTAL_UNLINK"
	ACTIVITYTYPE_PORTAL_ITEMS_DELETE ActivityType = "PORTAL_ITEMS_DELETE"
	ACTIVITYTYPE_ITEM_REMOVE ActivityType = "ITEM_REMOVE"
	ACTIVITYTYPE_INVARIANT_CHECKER_DELETION ActivityType = "INVARIANT_CHECKER_DELETION"
	ACTIVITYTYPE_SCOPES_UPDATE ActivityType = "SCOPES_UPDATE"
)

var allowedActivityTypeEnumValues = []ActivityType{
	"UNKNOWN",
	"ITEM_CREATE",
	"ITEM_IMPORT",
	"ITEM_UPDATE",
	"ITEM_UNLINK",
	"PORTAL_UNLINK",
	"PORTAL_ITEMS_DELETE",
	"ITEM_REMOVE",
	"INVARIANT_CHECKER_DELETION",
	"SCOPES_UPDATE",
}

func (v *ActivityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := ActivityType(value)


	*v = enumTypeValue
	return nil
}

// NewActivityTypeFromValue returns a pointer to a valid ActivityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActivityTypeFromValue(v string) (*ActivityType, error) {
	ev := ActivityType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActivityType) IsValid() bool {
	for _, existing := range allowedActivityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActivityType value
func (v ActivityType) Ptr() *ActivityType {
	return &v
}

type NullableActivityType struct {
	value *ActivityType
	isSet bool
}

func (v NullableActivityType) Get() *ActivityType {
	return v.value
}

func (v *NullableActivityType) Set(val *ActivityType) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityType) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityType(val *ActivityType) *NullableActivityType {
	return &NullableActivityType{value: val, isSet: true}
}

func (v NullableActivityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

