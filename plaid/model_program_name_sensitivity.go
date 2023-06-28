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

// ProgramNameSensitivity The valid name matching sensitivity configurations for a screening program. Note that while certain matching techniques may be more prevalent on less strict settings, all matching algorithms are enabled for every sensitivity.  `coarse` - See more potential matches. This sensitivity will see more broad phonetic matches across alphabets that make missing a potential hit very unlikely. This setting is noisier and will require more manual review.  `balanced` - A good default for most companies. This sensitivity is balanced to show high quality hits with reduced noise.  `strict` - Aggressive false positive reduction. This sensitivity will require names to be more similar than `coarse` and `balanced` settings, relying less on phonetics, while still accounting for character transpositions, missing tokens, and other common permutations.  `exact` - Matches must be nearly exact. This sensitivity will only show hits with exact or nearly exact name matches with only basic correction such as extraneous symbols and capitalization. This setting is generally not recommended unless you have a very specific use case.
type ProgramNameSensitivity string

var _ = fmt.Printf

// List of ProgramNameSensitivity
const (
	PROGRAMNAMESENSITIVITY_COARSE ProgramNameSensitivity = "coarse"
	PROGRAMNAMESENSITIVITY_BALANCED ProgramNameSensitivity = "balanced"
	PROGRAMNAMESENSITIVITY_STRICT ProgramNameSensitivity = "strict"
	PROGRAMNAMESENSITIVITY_EXACT ProgramNameSensitivity = "exact"
)

var allowedProgramNameSensitivityEnumValues = []ProgramNameSensitivity{
	"coarse",
	"balanced",
	"strict",
	"exact",
}

func (v *ProgramNameSensitivity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := ProgramNameSensitivity(value)


	*v = enumTypeValue
	return nil
}

// NewProgramNameSensitivityFromValue returns a pointer to a valid ProgramNameSensitivity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProgramNameSensitivityFromValue(v string) (*ProgramNameSensitivity, error) {
	ev := ProgramNameSensitivity(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProgramNameSensitivity) IsValid() bool {
	for _, existing := range allowedProgramNameSensitivityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProgramNameSensitivity value
func (v ProgramNameSensitivity) Ptr() *ProgramNameSensitivity {
	return &v
}

type NullableProgramNameSensitivity struct {
	value *ProgramNameSensitivity
	isSet bool
}

func (v NullableProgramNameSensitivity) Get() *ProgramNameSensitivity {
	return v.value
}

func (v *NullableProgramNameSensitivity) Set(val *ProgramNameSensitivity) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramNameSensitivity) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramNameSensitivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramNameSensitivity(val *ProgramNameSensitivity) *NullableProgramNameSensitivity {
	return &NullableProgramNameSensitivity{value: val, isSet: true}
}

func (v NullableProgramNameSensitivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramNameSensitivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

