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
)

// KYCCheckAddressSummary Result summary object specifying how the `address` field matched.
type KYCCheckAddressSummary struct {
	Summary MatchSummaryCode `json:"summary"`
	PoBox POBoxStatus `json:"po_box"`
	Type AddressPurposeLabel `json:"type"`
}

// NewKYCCheckAddressSummary instantiates a new KYCCheckAddressSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCCheckAddressSummary(summary MatchSummaryCode, poBox POBoxStatus, type_ AddressPurposeLabel) *KYCCheckAddressSummary {
	this := KYCCheckAddressSummary{}
	this.Summary = summary
	this.PoBox = poBox
	this.Type = type_
	return &this
}

// NewKYCCheckAddressSummaryWithDefaults instantiates a new KYCCheckAddressSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCCheckAddressSummaryWithDefaults() *KYCCheckAddressSummary {
	this := KYCCheckAddressSummary{}
	return &this
}

// GetSummary returns the Summary field value
func (o *KYCCheckAddressSummary) GetSummary() MatchSummaryCode {
	if o == nil {
		var ret MatchSummaryCode
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *KYCCheckAddressSummary) GetSummaryOk() (*MatchSummaryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *KYCCheckAddressSummary) SetSummary(v MatchSummaryCode) {
	o.Summary = v
}

// GetPoBox returns the PoBox field value
func (o *KYCCheckAddressSummary) GetPoBox() POBoxStatus {
	if o == nil {
		var ret POBoxStatus
		return ret
	}

	return o.PoBox
}

// GetPoBoxOk returns a tuple with the PoBox field value
// and a boolean to check if the value has been set.
func (o *KYCCheckAddressSummary) GetPoBoxOk() (*POBoxStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PoBox, true
}

// SetPoBox sets field value
func (o *KYCCheckAddressSummary) SetPoBox(v POBoxStatus) {
	o.PoBox = v
}

// GetType returns the Type field value
func (o *KYCCheckAddressSummary) GetType() AddressPurposeLabel {
	if o == nil {
		var ret AddressPurposeLabel
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KYCCheckAddressSummary) GetTypeOk() (*AddressPurposeLabel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KYCCheckAddressSummary) SetType(v AddressPurposeLabel) {
	o.Type = v
}

func (o KYCCheckAddressSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["summary"] = o.Summary
	}
	if true {
		toSerialize["po_box"] = o.PoBox
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableKYCCheckAddressSummary struct {
	value *KYCCheckAddressSummary
	isSet bool
}

func (v NullableKYCCheckAddressSummary) Get() *KYCCheckAddressSummary {
	return v.value
}

func (v *NullableKYCCheckAddressSummary) Set(val *KYCCheckAddressSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCCheckAddressSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCCheckAddressSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCCheckAddressSummary(val *KYCCheckAddressSummary) *NullableKYCCheckAddressSummary {
	return &NullableKYCCheckAddressSummary{value: val, isSet: true}
}

func (v NullableKYCCheckAddressSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCCheckAddressSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


