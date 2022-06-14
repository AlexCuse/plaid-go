/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// EntityWatchlistCode Shorthand identifier for a specific screening list for entities.
type EntityWatchlistCode string

// List of EntityWatchlistCode
const (
	ENTITYWATCHLISTCODE_CA_CON EntityWatchlistCode = "CA_CON"
	ENTITYWATCHLISTCODE_EU_CON EntityWatchlistCode = "EU_CON"
	ENTITYWATCHLISTCODE_IZ_SOE EntityWatchlistCode = "IZ_SOE"
	ENTITYWATCHLISTCODE_IZ_UNC EntityWatchlistCode = "IZ_UNC"
	ENTITYWATCHLISTCODE_US_CAP EntityWatchlistCode = "US_CAP"
	ENTITYWATCHLISTCODE_US_FSE EntityWatchlistCode = "US_FSE"
	ENTITYWATCHLISTCODE_US_MBS EntityWatchlistCode = "US_MBS"
	ENTITYWATCHLISTCODE_US_SDN EntityWatchlistCode = "US_SDN"
	ENTITYWATCHLISTCODE_US_SSI EntityWatchlistCode = "US_SSI"
	ENTITYWATCHLISTCODE_US_CMC EntityWatchlistCode = "US_CMC"
	ENTITYWATCHLISTCODE_US_UVL EntityWatchlistCode = "US_UVL"
	ENTITYWATCHLISTCODE_AU_CON EntityWatchlistCode = "AU_CON"
	ENTITYWATCHLISTCODE_UK_HMC EntityWatchlistCode = "UK_HMC"
)

var allowedEntityWatchlistCodeEnumValues = []EntityWatchlistCode{
	"CA_CON",
	"EU_CON",
	"IZ_SOE",
	"IZ_UNC",
	"US_CAP",
	"US_FSE",
	"US_MBS",
	"US_SDN",
	"US_SSI",
	"US_CMC",
	"US_UVL",
	"AU_CON",
	"UK_HMC",
}

func (v *EntityWatchlistCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EntityWatchlistCode(value)
	for _, existing := range allowedEntityWatchlistCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EntityWatchlistCode", value)
}

// NewEntityWatchlistCodeFromValue returns a pointer to a valid EntityWatchlistCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEntityWatchlistCodeFromValue(v string) (*EntityWatchlistCode, error) {
	ev := EntityWatchlistCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EntityWatchlistCode: valid values are %v", v, allowedEntityWatchlistCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EntityWatchlistCode) IsValid() bool {
	for _, existing := range allowedEntityWatchlistCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EntityWatchlistCode value
func (v EntityWatchlistCode) Ptr() *EntityWatchlistCode {
	return &v
}

type NullableEntityWatchlistCode struct {
	value *EntityWatchlistCode
	isSet bool
}

func (v NullableEntityWatchlistCode) Get() *EntityWatchlistCode {
	return v.value
}

func (v *NullableEntityWatchlistCode) Set(val *EntityWatchlistCode) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityWatchlistCode) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityWatchlistCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityWatchlistCode(val *EntityWatchlistCode) *NullableEntityWatchlistCode {
	return &NullableEntityWatchlistCode{value: val, isSet: true}
}

func (v NullableEntityWatchlistCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityWatchlistCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

