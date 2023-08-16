/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.410.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// RiskCheckIdentityAbuseSignals Result summary object capturing abuse signals related to `identity abuse`, e.g. stolen and synthetic identity fraud.
type RiskCheckIdentityAbuseSignals struct {
	SyntheticIdentity NullableRiskCheckSyntheticIdentity `json:"synthetic_identity"`
	StolenIdentity NullableRiskCheckStolenIdentity `json:"stolen_identity"`
	AdditionalProperties map[string]interface{}
}

type _RiskCheckIdentityAbuseSignals RiskCheckIdentityAbuseSignals

// NewRiskCheckIdentityAbuseSignals instantiates a new RiskCheckIdentityAbuseSignals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskCheckIdentityAbuseSignals(syntheticIdentity NullableRiskCheckSyntheticIdentity, stolenIdentity NullableRiskCheckStolenIdentity) *RiskCheckIdentityAbuseSignals {
	this := RiskCheckIdentityAbuseSignals{}
	this.SyntheticIdentity = syntheticIdentity
	this.StolenIdentity = stolenIdentity
	return &this
}

// NewRiskCheckIdentityAbuseSignalsWithDefaults instantiates a new RiskCheckIdentityAbuseSignals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskCheckIdentityAbuseSignalsWithDefaults() *RiskCheckIdentityAbuseSignals {
	this := RiskCheckIdentityAbuseSignals{}
	return &this
}

// GetSyntheticIdentity returns the SyntheticIdentity field value
// If the value is explicit nil, the zero value for RiskCheckSyntheticIdentity will be returned
func (o *RiskCheckIdentityAbuseSignals) GetSyntheticIdentity() RiskCheckSyntheticIdentity {
	if o == nil || o.SyntheticIdentity.Get() == nil {
		var ret RiskCheckSyntheticIdentity
		return ret
	}

	return *o.SyntheticIdentity.Get()
}

// GetSyntheticIdentityOk returns a tuple with the SyntheticIdentity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckIdentityAbuseSignals) GetSyntheticIdentityOk() (*RiskCheckSyntheticIdentity, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SyntheticIdentity.Get(), o.SyntheticIdentity.IsSet()
}

// SetSyntheticIdentity sets field value
func (o *RiskCheckIdentityAbuseSignals) SetSyntheticIdentity(v RiskCheckSyntheticIdentity) {
	o.SyntheticIdentity.Set(&v)
}

// GetStolenIdentity returns the StolenIdentity field value
// If the value is explicit nil, the zero value for RiskCheckStolenIdentity will be returned
func (o *RiskCheckIdentityAbuseSignals) GetStolenIdentity() RiskCheckStolenIdentity {
	if o == nil || o.StolenIdentity.Get() == nil {
		var ret RiskCheckStolenIdentity
		return ret
	}

	return *o.StolenIdentity.Get()
}

// GetStolenIdentityOk returns a tuple with the StolenIdentity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckIdentityAbuseSignals) GetStolenIdentityOk() (*RiskCheckStolenIdentity, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StolenIdentity.Get(), o.StolenIdentity.IsSet()
}

// SetStolenIdentity sets field value
func (o *RiskCheckIdentityAbuseSignals) SetStolenIdentity(v RiskCheckStolenIdentity) {
	o.StolenIdentity.Set(&v)
}

func (o RiskCheckIdentityAbuseSignals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["synthetic_identity"] = o.SyntheticIdentity.Get()
	}
	if true {
		toSerialize["stolen_identity"] = o.StolenIdentity.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskCheckIdentityAbuseSignals) UnmarshalJSON(bytes []byte) (err error) {
	varRiskCheckIdentityAbuseSignals := _RiskCheckIdentityAbuseSignals{}

	if err = json.Unmarshal(bytes, &varRiskCheckIdentityAbuseSignals); err == nil {
		*o = RiskCheckIdentityAbuseSignals(varRiskCheckIdentityAbuseSignals)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "synthetic_identity")
		delete(additionalProperties, "stolen_identity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskCheckIdentityAbuseSignals struct {
	value *RiskCheckIdentityAbuseSignals
	isSet bool
}

func (v NullableRiskCheckIdentityAbuseSignals) Get() *RiskCheckIdentityAbuseSignals {
	return v.value
}

func (v *NullableRiskCheckIdentityAbuseSignals) Set(val *RiskCheckIdentityAbuseSignals) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckIdentityAbuseSignals) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckIdentityAbuseSignals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckIdentityAbuseSignals(val *RiskCheckIdentityAbuseSignals) *NullableRiskCheckIdentityAbuseSignals {
	return &NullableRiskCheckIdentityAbuseSignals{value: val, isSet: true}
}

func (v NullableRiskCheckIdentityAbuseSignals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckIdentityAbuseSignals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


