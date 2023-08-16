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

// FraudAmount The amount and currency of the fraud or attempted fraud. `fraud_amount` should be omitted to indicate an unknown fraud amount.
type FraudAmount struct {
	IsoCurrencyCode ISOCurrencyCode `json:"iso_currency_code"`
	// The amount value. This value can be 0 to indicate no money was lost. Must not contain more than two digits of precision (e.g., `1.23`).
	Value float64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _FraudAmount FraudAmount

// NewFraudAmount instantiates a new FraudAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFraudAmount(isoCurrencyCode ISOCurrencyCode, value float64) *FraudAmount {
	this := FraudAmount{}
	this.IsoCurrencyCode = isoCurrencyCode
	this.Value = value
	return &this
}

// NewFraudAmountWithDefaults instantiates a new FraudAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFraudAmountWithDefaults() *FraudAmount {
	this := FraudAmount{}
	return &this
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *FraudAmount) GetIsoCurrencyCode() ISOCurrencyCode {
	if o == nil {
		var ret ISOCurrencyCode
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *FraudAmount) GetIsoCurrencyCodeOk() (*ISOCurrencyCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *FraudAmount) SetIsoCurrencyCode(v ISOCurrencyCode) {
	o.IsoCurrencyCode = v
}

// GetValue returns the Value field value
func (o *FraudAmount) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FraudAmount) GetValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FraudAmount) SetValue(v float64) {
	o.Value = v
}

func (o FraudAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FraudAmount) UnmarshalJSON(bytes []byte) (err error) {
	varFraudAmount := _FraudAmount{}

	if err = json.Unmarshal(bytes, &varFraudAmount); err == nil {
		*o = FraudAmount(varFraudAmount)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFraudAmount struct {
	value *FraudAmount
	isSet bool
}

func (v NullableFraudAmount) Get() *FraudAmount {
	return v.value
}

func (v *NullableFraudAmount) Set(val *FraudAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableFraudAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableFraudAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFraudAmount(val *FraudAmount) *NullableFraudAmount {
	return &NullableFraudAmount{value: val, isSet: true}
}

func (v NullableFraudAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFraudAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


