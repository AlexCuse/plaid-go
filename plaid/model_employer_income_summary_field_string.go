/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.61.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// EmployerIncomeSummaryFieldString struct for EmployerIncomeSummaryFieldString
type EmployerIncomeSummaryFieldString struct {
	// The value of the field.
	Value string `json:"value"`
	VerificationStatus VerificationStatus `json:"verification_status"`
}

// NewEmployerIncomeSummaryFieldString instantiates a new EmployerIncomeSummaryFieldString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployerIncomeSummaryFieldString(value string, verificationStatus VerificationStatus) *EmployerIncomeSummaryFieldString {
	this := EmployerIncomeSummaryFieldString{}
	this.Value = value
	this.VerificationStatus = verificationStatus
	return &this
}

// NewEmployerIncomeSummaryFieldStringWithDefaults instantiates a new EmployerIncomeSummaryFieldString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployerIncomeSummaryFieldStringWithDefaults() *EmployerIncomeSummaryFieldString {
	this := EmployerIncomeSummaryFieldString{}
	return &this
}

// GetValue returns the Value field value
func (o *EmployerIncomeSummaryFieldString) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EmployerIncomeSummaryFieldString) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EmployerIncomeSummaryFieldString) SetValue(v string) {
	o.Value = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *EmployerIncomeSummaryFieldString) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *EmployerIncomeSummaryFieldString) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *EmployerIncomeSummaryFieldString) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

func (o EmployerIncomeSummaryFieldString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableEmployerIncomeSummaryFieldString struct {
	value *EmployerIncomeSummaryFieldString
	isSet bool
}

func (v NullableEmployerIncomeSummaryFieldString) Get() *EmployerIncomeSummaryFieldString {
	return v.value
}

func (v *NullableEmployerIncomeSummaryFieldString) Set(val *EmployerIncomeSummaryFieldString) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployerIncomeSummaryFieldString) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployerIncomeSummaryFieldString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployerIncomeSummaryFieldString(val *EmployerIncomeSummaryFieldString) *NullableEmployerIncomeSummaryFieldString {
	return &NullableEmployerIncomeSummaryFieldString{value: val, isSet: true}
}

func (v NullableEmployerIncomeSummaryFieldString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployerIncomeSummaryFieldString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


