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
)

// TransferCapabilitiesGetRTP Contains the supported service types in RTP
type TransferCapabilitiesGetRTP struct {
	// When `true`, the linked Item's institution supports RTP credit transfer.
	Credit *bool `json:"credit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferCapabilitiesGetRTP TransferCapabilitiesGetRTP

// NewTransferCapabilitiesGetRTP instantiates a new TransferCapabilitiesGetRTP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCapabilitiesGetRTP() *TransferCapabilitiesGetRTP {
	this := TransferCapabilitiesGetRTP{}
	var credit bool = false
	this.Credit = &credit
	return &this
}

// NewTransferCapabilitiesGetRTPWithDefaults instantiates a new TransferCapabilitiesGetRTP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCapabilitiesGetRTPWithDefaults() *TransferCapabilitiesGetRTP {
	this := TransferCapabilitiesGetRTP{}
	var credit bool = false
	this.Credit = &credit
	return &this
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *TransferCapabilitiesGetRTP) GetCredit() bool {
	if o == nil || o.Credit == nil {
		var ret bool
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCapabilitiesGetRTP) GetCreditOk() (*bool, bool) {
	if o == nil || o.Credit == nil {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *TransferCapabilitiesGetRTP) HasCredit() bool {
	if o != nil && o.Credit != nil {
		return true
	}

	return false
}

// SetCredit gets a reference to the given bool and assigns it to the Credit field.
func (o *TransferCapabilitiesGetRTP) SetCredit(v bool) {
	o.Credit = &v
}

func (o TransferCapabilitiesGetRTP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credit != nil {
		toSerialize["credit"] = o.Credit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferCapabilitiesGetRTP) UnmarshalJSON(bytes []byte) (err error) {
	varTransferCapabilitiesGetRTP := _TransferCapabilitiesGetRTP{}

	if err = json.Unmarshal(bytes, &varTransferCapabilitiesGetRTP); err == nil {
		*o = TransferCapabilitiesGetRTP(varTransferCapabilitiesGetRTP)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "credit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferCapabilitiesGetRTP struct {
	value *TransferCapabilitiesGetRTP
	isSet bool
}

func (v NullableTransferCapabilitiesGetRTP) Get() *TransferCapabilitiesGetRTP {
	return v.value
}

func (v *NullableTransferCapabilitiesGetRTP) Set(val *TransferCapabilitiesGetRTP) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCapabilitiesGetRTP) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCapabilitiesGetRTP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCapabilitiesGetRTP(val *TransferCapabilitiesGetRTP) *NullableTransferCapabilitiesGetRTP {
	return &NullableTransferCapabilitiesGetRTP{value: val, isSet: true}
}

func (v NullableTransferCapabilitiesGetRTP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCapabilitiesGetRTP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


