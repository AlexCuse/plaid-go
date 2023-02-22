/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.334.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// SignalDecisionOutcome The payment decision from the risk assessment.  `APPROVE`: approve the transaction without requiring further actions from your customers. For example, use this field if you are placing a standard hold for all the approved transactions before making funds available to your customers. You should also use this field if you decide to accelerate the fund availability for your customers.  `REVIEW`: the transaction requires manual review  `REJECT`: reject the transaction  `TAKE_OTHER_RISK_MEASURES`: for example, placing a longer hold on funds than those approved transactions or introducing customer frictions such as step-up verification/authentication  `NOT_EVALUATED`: if only logging the Signal results without using them  Possible values:  `APPROVE`, `REVIEW`, `REJECT`, `TAKE_OTHER_RISK_MEASURES`, `NOT_EVALUATED` 
type SignalDecisionOutcome string

var _ = fmt.Printf

// List of SignalDecisionOutcome
const (
	SIGNALDECISIONOUTCOME_APPROVE SignalDecisionOutcome = "APPROVE"
	SIGNALDECISIONOUTCOME_REVIEW SignalDecisionOutcome = "REVIEW"
	SIGNALDECISIONOUTCOME_REJECT SignalDecisionOutcome = "REJECT"
	SIGNALDECISIONOUTCOME_TAKE_OTHER_RISK_MEASURES SignalDecisionOutcome = "TAKE_OTHER_RISK_MEASURES"
	SIGNALDECISIONOUTCOME_NOT_EVALUATED SignalDecisionOutcome = "NOT_EVALUATED"
)

var allowedSignalDecisionOutcomeEnumValues = []SignalDecisionOutcome{
	"APPROVE",
	"REVIEW",
	"REJECT",
	"TAKE_OTHER_RISK_MEASURES",
	"NOT_EVALUATED",
}

func (v *SignalDecisionOutcome) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := SignalDecisionOutcome(value)


	*v = enumTypeValue
	return nil
}

// NewSignalDecisionOutcomeFromValue returns a pointer to a valid SignalDecisionOutcome
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignalDecisionOutcomeFromValue(v string) (*SignalDecisionOutcome, error) {
	ev := SignalDecisionOutcome(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignalDecisionOutcome) IsValid() bool {
	for _, existing := range allowedSignalDecisionOutcomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignalDecisionOutcome value
func (v SignalDecisionOutcome) Ptr() *SignalDecisionOutcome {
	return &v
}

type NullableSignalDecisionOutcome struct {
	value *SignalDecisionOutcome
	isSet bool
}

func (v NullableSignalDecisionOutcome) Get() *SignalDecisionOutcome {
	return v.value
}

func (v *NullableSignalDecisionOutcome) Set(val *SignalDecisionOutcome) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalDecisionOutcome) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalDecisionOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalDecisionOutcome(val *SignalDecisionOutcome) *NullableSignalDecisionOutcome {
	return &NullableSignalDecisionOutcome{value: val, isSet: true}
}

func (v NullableSignalDecisionOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalDecisionOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

