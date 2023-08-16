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

// TransferAuthorizationDecisionRationaleCode A code representing the rationale for approving or declining the proposed transfer. Possible values are:  `MANUALLY_VERIFIED_ITEM` – Item created via same-day micro deposits, limited information available. Plaid will offer `approved` as a transaction decision.  `ITEM_LOGIN_REQUIRED` – Unable to collect the account information due to Item staleness. Can be rectified using Link in update mode. Plaid will offer `approved` as a transaction decision.  `PAYMENT_PROFILE_LOGIN_REQUIRED` - Unable to collect the account information due to invalid login when using Payment Profiles. Can be rectified using update mode for Payment Profile. Plaid will offer `approved` as a transaction decision.  `ERROR` – Unable to collect the account information due to an error. Plaid will offer `approved` as a transaction decision.  `NSF` – Transaction likely to result in a return due to insufficient funds. Plaid will offer `declined` as a transaction decision.  `RISK` - Transaction is high-risk. Plaid will offer `declined` as a transaction decision.  `TRANSFER_LIMIT_REACHED` - One or several transfer limits are reached, e.g. monthly transfer limit. Plaid will offer `declined` as a transaction decision.  `MIGRATED_ACCOUNT_ITEM` - Item created via `/transfer/account_migration` endpoint, limited information available. Plaid will offer `approved` as a transaction decision.
type TransferAuthorizationDecisionRationaleCode string

var _ = fmt.Printf

// List of TransferAuthorizationDecisionRationaleCode
const (
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_NSF TransferAuthorizationDecisionRationaleCode = "NSF"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_RISK TransferAuthorizationDecisionRationaleCode = "RISK"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_TRANSFER_LIMIT_REACHED TransferAuthorizationDecisionRationaleCode = "TRANSFER_LIMIT_REACHED"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_MANUALLY_VERIFIED_ITEM TransferAuthorizationDecisionRationaleCode = "MANUALLY_VERIFIED_ITEM"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_ITEM_LOGIN_REQUIRED TransferAuthorizationDecisionRationaleCode = "ITEM_LOGIN_REQUIRED"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_PAYMENT_PROFILE_LOGIN_REQUIRED TransferAuthorizationDecisionRationaleCode = "PAYMENT_PROFILE_LOGIN_REQUIRED"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_ERROR TransferAuthorizationDecisionRationaleCode = "ERROR"
	TRANSFERAUTHORIZATIONDECISIONRATIONALECODE_MIGRATED_ACCOUNT_ITEM TransferAuthorizationDecisionRationaleCode = "MIGRATED_ACCOUNT_ITEM"
)

var allowedTransferAuthorizationDecisionRationaleCodeEnumValues = []TransferAuthorizationDecisionRationaleCode{
	"NSF",
	"RISK",
	"TRANSFER_LIMIT_REACHED",
	"MANUALLY_VERIFIED_ITEM",
	"ITEM_LOGIN_REQUIRED",
	"PAYMENT_PROFILE_LOGIN_REQUIRED",
	"ERROR",
	"MIGRATED_ACCOUNT_ITEM",
}

func (v *TransferAuthorizationDecisionRationaleCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferAuthorizationDecisionRationaleCode(value)


	*v = enumTypeValue
	return nil
}

// NewTransferAuthorizationDecisionRationaleCodeFromValue returns a pointer to a valid TransferAuthorizationDecisionRationaleCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferAuthorizationDecisionRationaleCodeFromValue(v string) (*TransferAuthorizationDecisionRationaleCode, error) {
	ev := TransferAuthorizationDecisionRationaleCode(v)


	return &ev, nil
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
