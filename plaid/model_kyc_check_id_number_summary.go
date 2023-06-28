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
)

// KYCCheckIDNumberSummary Result summary object specifying how the `id_number` field matched.
type KYCCheckIDNumberSummary struct {
	Summary MatchSummaryCode `json:"summary"`
	AdditionalProperties map[string]interface{}
}

type _KYCCheckIDNumberSummary KYCCheckIDNumberSummary

// NewKYCCheckIDNumberSummary instantiates a new KYCCheckIDNumberSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCCheckIDNumberSummary(summary MatchSummaryCode) *KYCCheckIDNumberSummary {
	this := KYCCheckIDNumberSummary{}
	this.Summary = summary
	return &this
}

// NewKYCCheckIDNumberSummaryWithDefaults instantiates a new KYCCheckIDNumberSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCCheckIDNumberSummaryWithDefaults() *KYCCheckIDNumberSummary {
	this := KYCCheckIDNumberSummary{}
	return &this
}

// GetSummary returns the Summary field value
func (o *KYCCheckIDNumberSummary) GetSummary() MatchSummaryCode {
	if o == nil {
		var ret MatchSummaryCode
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *KYCCheckIDNumberSummary) GetSummaryOk() (*MatchSummaryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *KYCCheckIDNumberSummary) SetSummary(v MatchSummaryCode) {
	o.Summary = v
}

func (o KYCCheckIDNumberSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KYCCheckIDNumberSummary) UnmarshalJSON(bytes []byte) (err error) {
	varKYCCheckIDNumberSummary := _KYCCheckIDNumberSummary{}

	if err = json.Unmarshal(bytes, &varKYCCheckIDNumberSummary); err == nil {
		*o = KYCCheckIDNumberSummary(varKYCCheckIDNumberSummary)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "summary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKYCCheckIDNumberSummary struct {
	value *KYCCheckIDNumberSummary
	isSet bool
}

func (v NullableKYCCheckIDNumberSummary) Get() *KYCCheckIDNumberSummary {
	return v.value
}

func (v *NullableKYCCheckIDNumberSummary) Set(val *KYCCheckIDNumberSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCCheckIDNumberSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCCheckIDNumberSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCCheckIDNumberSummary(val *KYCCheckIDNumberSummary) *NullableKYCCheckIDNumberSummary {
	return &NullableKYCCheckIDNumberSummary{value: val, isSet: true}
}

func (v NullableKYCCheckIDNumberSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCCheckIDNumberSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


