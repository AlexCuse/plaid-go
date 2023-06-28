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

// LoanIdentifierType A value from a MISMO prescribed list that specifies the type of loan identifier.
type LoanIdentifierType string

var _ = fmt.Printf

// List of LoanIdentifierType
const (
	LOANIDENTIFIERTYPE_LENDER_LOAN LoanIdentifierType = "LenderLoan"
	LOANIDENTIFIERTYPE_UNIVERSAL_LOAN LoanIdentifierType = "UniversalLoan"
)

var allowedLoanIdentifierTypeEnumValues = []LoanIdentifierType{
	"LenderLoan",
	"UniversalLoan",
}

func (v *LoanIdentifierType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := LoanIdentifierType(value)


	*v = enumTypeValue
	return nil
}

// NewLoanIdentifierTypeFromValue returns a pointer to a valid LoanIdentifierType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoanIdentifierTypeFromValue(v string) (*LoanIdentifierType, error) {
	ev := LoanIdentifierType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoanIdentifierType) IsValid() bool {
	for _, existing := range allowedLoanIdentifierTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LoanIdentifierType value
func (v LoanIdentifierType) Ptr() *LoanIdentifierType {
	return &v
}

type NullableLoanIdentifierType struct {
	value *LoanIdentifierType
	isSet bool
}

func (v NullableLoanIdentifierType) Get() *LoanIdentifierType {
	return v.value
}

func (v *NullableLoanIdentifierType) Set(val *LoanIdentifierType) {
	v.value = val
	v.isSet = true
}

func (v NullableLoanIdentifierType) IsSet() bool {
	return v.isSet
}

func (v *NullableLoanIdentifierType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoanIdentifierType(val *LoanIdentifierType) *NullableLoanIdentifierType {
	return &NullableLoanIdentifierType{value: val, isSet: true}
}

func (v NullableLoanIdentifierType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoanIdentifierType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

