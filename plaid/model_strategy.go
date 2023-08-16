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

// Strategy An instruction specifying what steps the new Identity Verification attempt should require the user to complete:   `reset` - Restart the user at the beginning of the session, regardless of whether they successfully completed part of their previous session.  `incomplete` - Start the new session at the step that the user failed in the previous session, skipping steps that have already been successfully completed.  `infer` - If the most recent Identity Verification attempt associated with the given `client_user_id` has a status of `failed` or `expired`, retry using the `incomplete` strategy. Otherwise, use the `reset` strategy.  `custom` - Start the new session with a custom configuration, specified by the value of the `steps` field  Note:  The `incomplete` strategy cannot be applied if the session's failing step is `screening` or `risk_check`.  The `infer` strategy cannot be applied if the session's status is still `active`
type Strategy string

var _ = fmt.Printf

// List of Strategy
const (
	STRATEGY_RESET Strategy = "reset"
	STRATEGY_INCOMPLETE Strategy = "incomplete"
	STRATEGY_INFER Strategy = "infer"
	STRATEGY_CUSTOM Strategy = "custom"
)

var allowedStrategyEnumValues = []Strategy{
	"reset",
	"incomplete",
	"infer",
	"custom",
}

func (v *Strategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := Strategy(value)


	*v = enumTypeValue
	return nil
}

// NewStrategyFromValue returns a pointer to a valid Strategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStrategyFromValue(v string) (*Strategy, error) {
	ev := Strategy(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Strategy) IsValid() bool {
	for _, existing := range allowedStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Strategy value
func (v Strategy) Ptr() *Strategy {
	return &v
}

type NullableStrategy struct {
	value *Strategy
	isSet bool
}

func (v NullableStrategy) Get() *Strategy {
	return v.value
}

func (v *NullableStrategy) Set(val *Strategy) {
	v.value = val
	v.isSet = true
}

func (v NullableStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrategy(val *Strategy) *NullableStrategy {
	return &NullableStrategy{value: val, isSet: true}
}

func (v NullableStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
