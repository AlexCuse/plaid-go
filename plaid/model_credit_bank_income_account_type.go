/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// CreditBankIncomeAccountType The account type. This will always be `depository`.
type CreditBankIncomeAccountType string

// List of CreditBankIncomeAccountType
const (
	CREDITBANKINCOMEACCOUNTTYPE_DEPOSITORY CreditBankIncomeAccountType = "depository"
)

var allowedCreditBankIncomeAccountTypeEnumValues = []CreditBankIncomeAccountType{
	"depository",
}

func (v *CreditBankIncomeAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreditBankIncomeAccountType(value)
	for _, existing := range allowedCreditBankIncomeAccountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreditBankIncomeAccountType", value)
}

// NewCreditBankIncomeAccountTypeFromValue returns a pointer to a valid CreditBankIncomeAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditBankIncomeAccountTypeFromValue(v string) (*CreditBankIncomeAccountType, error) {
	ev := CreditBankIncomeAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreditBankIncomeAccountType: valid values are %v", v, allowedCreditBankIncomeAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditBankIncomeAccountType) IsValid() bool {
	for _, existing := range allowedCreditBankIncomeAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditBankIncomeAccountType value
func (v CreditBankIncomeAccountType) Ptr() *CreditBankIncomeAccountType {
	return &v
}

type NullableCreditBankIncomeAccountType struct {
	value *CreditBankIncomeAccountType
	isSet bool
}

func (v NullableCreditBankIncomeAccountType) Get() *CreditBankIncomeAccountType {
	return v.value
}

func (v *NullableCreditBankIncomeAccountType) Set(val *CreditBankIncomeAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeAccountType(val *CreditBankIncomeAccountType) *NullableCreditBankIncomeAccountType {
	return &NullableCreditBankIncomeAccountType{value: val, isSet: true}
}

func (v NullableCreditBankIncomeAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

