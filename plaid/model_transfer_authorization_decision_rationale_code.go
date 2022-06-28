/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.131.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// TransferAuthorizationDecisionRationaleCode A code representing the rationale for approving or declining the proposed transfer. Possible values are:  `MANUALLY_VERIFIED_ITEM` – Item created via same-day micro deposits, limited information available. Plaid will offer `approved` as a transaction decision.  `LOGIN_REQUIRED` – Unable to collect the account information due to Item staleness. Can be rectified using Link in update mode. Plaid will offer `approved` as a transaction decision.  `ERROR` – Unable to collect the account information due to an error. Plaid will offer `approved` as a transaction decision.  `NSF` – Transaction likely to result in a return due to insufficient funds. Plaid will offer `declined` as a transaction decision.  `RISK` - Transaction is high-risk. Plaid will offer `declined` as a transaction decision.  `TRANSFER_LIMIT_REACHED` - One or several transfer limits are reached, e.g. monthly transfer limit. Plaid will offer `declined` as a transaction decision.
type TransferAuthorizationDecisionRationaleCode string

// List of TransferAuthorizationDecisionRationaleCode
const (
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_NSF TransferAuthorizationDecisionRationaleCode = "NSF"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_RISK TransferAuthorizationDecisionRationaleCode = "RISK"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_TRANSFER_LIMIT_REACHED TransferAuthorizationDecisionRationaleCode = "TRANSFER_LIMIT_REACHED"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_MANUALLY_VERIFIED_ITEM TransferAuthorizationDecisionRationaleCode = "MANUALLY_VERIFIED_ITEM"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_LOGIN_REQUIRED TransferAuthorizationDecisionRationaleCode = "LOGIN_REQUIRED"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_ERROR TransferAuthorizationDecisionRationaleCode = "ERROR"
)

var allowedTransferAuthorizationDecisionRationaleCodeEnumValues = []TransferAuthorizationDecisionRationaleCode{
	"NSF",
	"RISK",
	"TRANSFER_LIMIT_REACHED",
	"MANUALLY_VERIFIED_ITEM",
	"LOGIN_REQUIRED",
	"ERROR",
}

func (v *TransferAuthorizationDecisionRationaleCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransferAuthorizationDecisionRationaleCode(value)
	for _, existing := range allowedTransferAuthorizationDecisionRationaleCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransferAuthorizationDecisionRationaleCode", value)
}

// NewTransferAuthorizationDecisionRationaleCodeFromValue returns a pointer to a valid TransferAuthorizationDecisionRationaleCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferAuthorizationDecisionRationaleCodeFromValue(v string) (*TransferAuthorizationDecisionRationaleCode, error) {
	ev := TransferAuthorizationDecisionRationaleCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransferAuthorizationDecisionRationaleCode: valid values are %v", v, allowedTransferAuthorizationDecisionRationaleCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferAuthorizationDecisionRationaleCode) IsValid() bool {
	for _, existing := range allowedTransferAuthorizationDecisionRationaleCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferAuthorizationDecisionRationaleCode value
func (v TransferAuthorizationDecisionRationaleCode) Ptr() *TransferAuthorizationDecisionRationaleCode {
	return &v
}

type NullableTransferAuthorizationDecisionRationaleCode struct {
	value *TransferAuthorizationDecisionRationaleCode
	isSet bool
}

func (v NullableTransferAuthorizationDecisionRationaleCode) Get() *TransferAuthorizationDecisionRationaleCode {
	return v.value
}

func (v *NullableTransferAuthorizationDecisionRationaleCode) Set(val *TransferAuthorizationDecisionRationaleCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationDecisionRationaleCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationDecisionRationaleCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationDecisionRationaleCode(val *TransferAuthorizationDecisionRationaleCode) *NullableTransferAuthorizationDecisionRationaleCode {
	return &NullableTransferAuthorizationDecisionRationaleCode{value: val, isSet: true}
}

func (v NullableTransferAuthorizationDecisionRationaleCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationDecisionRationaleCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

