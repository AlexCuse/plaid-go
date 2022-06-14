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

// CreditBankIncomeErrorType A broad categorization of the error. Safe for programmatic use.
type CreditBankIncomeErrorType string

// List of CreditBankIncomeErrorType
const (
	CREDITBANKINCOMEERRORTYPE_INTERNAL_SERVER_ERROR CreditBankIncomeErrorType = "INTERNAL_SERVER_ERROR"
	CREDITBANKINCOMEERRORTYPE_INSUFFICIENT_CREDENTIALS CreditBankIncomeErrorType = "INSUFFICIENT_CREDENTIALS"
	CREDITBANKINCOMEERRORTYPE_ITEM_LOCKED CreditBankIncomeErrorType = "ITEM_LOCKED"
	CREDITBANKINCOMEERRORTYPE_USER_SETUP_REQUIRED CreditBankIncomeErrorType = "USER_SETUP_REQUIRED"
	CREDITBANKINCOMEERRORTYPE_COUNTRY_NOT_SUPPORTED CreditBankIncomeErrorType = "COUNTRY_NOT_SUPPORTED"
	CREDITBANKINCOMEERRORTYPE_INSTITUTION_DOWN CreditBankIncomeErrorType = "INSTITUTION_DOWN"
	CREDITBANKINCOMEERRORTYPE_INSTITUTION_NO_LONGER_SUPPORTED CreditBankIncomeErrorType = "INSTITUTION_NO_LONGER_SUPPORTED"
	CREDITBANKINCOMEERRORTYPE_INSTITUTION_NOT_RESPONDING CreditBankIncomeErrorType = "INSTITUTION_NOT_RESPONDING"
	CREDITBANKINCOMEERRORTYPE_INVALID_CREDENTIALS CreditBankIncomeErrorType = "INVALID_CREDENTIALS"
	CREDITBANKINCOMEERRORTYPE_INVALID_MFA CreditBankIncomeErrorType = "INVALID_MFA"
	CREDITBANKINCOMEERRORTYPE_INVALID_SEND_METHOD CreditBankIncomeErrorType = "INVALID_SEND_METHOD"
	CREDITBANKINCOMEERRORTYPE_ITEM_LOGIN_REQUIRED CreditBankIncomeErrorType = "ITEM_LOGIN_REQUIRED"
	CREDITBANKINCOMEERRORTYPE_MFA_NOT_SUPPORTED CreditBankIncomeErrorType = "MFA_NOT_SUPPORTED"
	CREDITBANKINCOMEERRORTYPE_NO_ACCOUNTS CreditBankIncomeErrorType = "NO_ACCOUNTS"
	CREDITBANKINCOMEERRORTYPE_ITEM_NOT_SUPPORTED CreditBankIncomeErrorType = "ITEM_NOT_SUPPORTED"
	CREDITBANKINCOMEERRORTYPE_ACCESS_NOT_GRANTED CreditBankIncomeErrorType = "ACCESS_NOT_GRANTED"
)

var allowedCreditBankIncomeErrorTypeEnumValues = []CreditBankIncomeErrorType{
	"INTERNAL_SERVER_ERROR",
	"INSUFFICIENT_CREDENTIALS",
	"ITEM_LOCKED",
	"USER_SETUP_REQUIRED",
	"COUNTRY_NOT_SUPPORTED",
	"INSTITUTION_DOWN",
	"INSTITUTION_NO_LONGER_SUPPORTED",
	"INSTITUTION_NOT_RESPONDING",
	"INVALID_CREDENTIALS",
	"INVALID_MFA",
	"INVALID_SEND_METHOD",
	"ITEM_LOGIN_REQUIRED",
	"MFA_NOT_SUPPORTED",
	"NO_ACCOUNTS",
	"ITEM_NOT_SUPPORTED",
	"ACCESS_NOT_GRANTED",
}

func (v *CreditBankIncomeErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreditBankIncomeErrorType(value)
	for _, existing := range allowedCreditBankIncomeErrorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreditBankIncomeErrorType", value)
}

// NewCreditBankIncomeErrorTypeFromValue returns a pointer to a valid CreditBankIncomeErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditBankIncomeErrorTypeFromValue(v string) (*CreditBankIncomeErrorType, error) {
	ev := CreditBankIncomeErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreditBankIncomeErrorType: valid values are %v", v, allowedCreditBankIncomeErrorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditBankIncomeErrorType) IsValid() bool {
	for _, existing := range allowedCreditBankIncomeErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditBankIncomeErrorType value
func (v CreditBankIncomeErrorType) Ptr() *CreditBankIncomeErrorType {
	return &v
}

type NullableCreditBankIncomeErrorType struct {
	value *CreditBankIncomeErrorType
	isSet bool
}

func (v NullableCreditBankIncomeErrorType) Get() *CreditBankIncomeErrorType {
	return v.value
}

func (v *NullableCreditBankIncomeErrorType) Set(val *CreditBankIncomeErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeErrorType(val *CreditBankIncomeErrorType) *NullableCreditBankIncomeErrorType {
	return &NullableCreditBankIncomeErrorType{value: val, isSet: true}
}

func (v NullableCreditBankIncomeErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

