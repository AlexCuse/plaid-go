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

// TransferAuthorizationDecision  A decision regarding the proposed transfer.  `approved` – The proposed transfer has received the end user's consent and has been approved for processing by Plaid. The `decision_rationale` field is set if Plaid was unable to fetch the account information. You may proceed with the transfer, but further review is recommended (i.e., use Link in update to re-authenticate your user when `decision_rationale.code` is `ITEM_LOGIN_REQUIRED`). Refer to the `code` field in the `decision_rationale` object for details.  `declined` – Plaid reviewed the proposed transfer and declined processing. Refer to the `code` field in the `decision_rationale` object for details.
type TransferAuthorizationDecision string

var _ = fmt.Printf

// List of TransferAuthorizationDecision
const (
	TRANSFERAUTHORIZATIONDECISION_APPROVED TransferAuthorizationDecision = "approved"
	TRANSFERAUTHORIZATIONDECISION_DECLINED TransferAuthorizationDecision = "declined"
)

var allowedTransferAuthorizationDecisionEnumValues = []TransferAuthorizationDecision{
	"approved",
	"declined",
}

func (v *TransferAuthorizationDecision) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferAuthorizationDecision(value)


	*v = enumTypeValue
	return nil
}

// NewTransferAuthorizationDecisionFromValue returns a pointer to a valid TransferAuthorizationDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferAuthorizationDecisionFromValue(v string) (*TransferAuthorizationDecision, error) {
	ev := TransferAuthorizationDecision(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferAuthorizationDecision) IsValid() bool {
	for _, existing := range allowedTransferAuthorizationDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferAuthorizationDecision value
func (v TransferAuthorizationDecision) Ptr() *TransferAuthorizationDecision {
	return &v
}

type NullableTransferAuthorizationDecision struct {
	value *TransferAuthorizationDecision
	isSet bool
}

func (v NullableTransferAuthorizationDecision) Get() *TransferAuthorizationDecision {
	return v.value
}

func (v *NullableTransferAuthorizationDecision) Set(val *TransferAuthorizationDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationDecision(val *TransferAuthorizationDecision) *NullableTransferAuthorizationDecision {
	return &NullableTransferAuthorizationDecision{value: val, isSet: true}
}

func (v NullableTransferAuthorizationDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

