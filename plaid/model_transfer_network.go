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

// TransferNetwork The network or rails used for the transfer.  For transfers submitted as either `ach` or `same-day-ach` the cutoff for same-day is 3:30 PM Eastern Time and the cutoff for next-day transfers is 5:30 PM Eastern Time. It is recommended to submit a transfer at least 15 minutes before the cutoff time in order to ensure that it will be processed before the cutoff. Any transfer that is indicated as `same-day-ach` and that misses the same-day cutoff, but is submitted in time for the next-day cutoff, will be sent over next-day rails and will not incur same-day charges. Note that both legs of the transfer will be downgraded if applicable.
type TransferNetwork string

var _ = fmt.Printf

// List of TransferNetwork
const (
	TRANSFERNETWORK_ACH TransferNetwork = "ach"
	TRANSFERNETWORK_SAME_DAY_ACH TransferNetwork = "same-day-ach"
	TRANSFERNETWORK_RTP TransferNetwork = "rtp"
)

var allowedTransferNetworkEnumValues = []TransferNetwork{
	"ach",
	"same-day-ach",
	"rtp",
}

func (v *TransferNetwork) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferNetwork(value)


	*v = enumTypeValue
	return nil
}

// NewTransferNetworkFromValue returns a pointer to a valid TransferNetwork
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferNetworkFromValue(v string) (*TransferNetwork, error) {
	ev := TransferNetwork(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferNetwork) IsValid() bool {
	for _, existing := range allowedTransferNetworkEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferNetwork value
func (v TransferNetwork) Ptr() *TransferNetwork {
	return &v
}

type NullableTransferNetwork struct {
	value *TransferNetwork
	isSet bool
}

func (v NullableTransferNetwork) Get() *TransferNetwork {
	return v.value
}

func (v *NullableTransferNetwork) Set(val *TransferNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferNetwork(val *TransferNetwork) *NullableTransferNetwork {
	return &NullableTransferNetwork{value: val, isSet: true}
}

func (v NullableTransferNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

