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

// WalletTransactionAmount The amount and currency of a transaction
type WalletTransactionAmount struct {
	IsoCurrencyCode WalletISOCurrencyCode `json:"iso_currency_code"`
	// The amount of the transaction. Must contain at most two digits of precision e.g. `1.23`.
	Value float64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionAmount WalletTransactionAmount

// NewWalletTransactionAmount instantiates a new WalletTransactionAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionAmount(isoCurrencyCode WalletISOCurrencyCode, value float64) *WalletTransactionAmount {
	this := WalletTransactionAmount{}
	this.IsoCurrencyCode = isoCurrencyCode
	this.Value = value
	return &this
}

// NewWalletTransactionAmountWithDefaults instantiates a new WalletTransactionAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionAmountWithDefaults() *WalletTransactionAmount {
	this := WalletTransactionAmount{}
	return &this
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *WalletTransactionAmount) GetIsoCurrencyCode() WalletISOCurrencyCode {
	if o == nil {
		var ret WalletISOCurrencyCode
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionAmount) GetIsoCurrencyCodeOk() (*WalletISOCurrencyCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *WalletTransactionAmount) SetIsoCurrencyCode(v WalletISOCurrencyCode) {
	o.IsoCurrencyCode = v
}

// GetValue returns the Value field value
func (o *WalletTransactionAmount) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionAmount) GetValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *WalletTransactionAmount) SetValue(v float64) {
	o.Value = v
}

func (o WalletTransactionAmount) MarshalJSON() ([]byte, error) {
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

func (o *WalletTransactionAmount) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionAmount := _WalletTransactionAmount{}

	if err = json.Unmarshal(bytes, &varWalletTransactionAmount); err == nil {
		*o = WalletTransactionAmount(varWalletTransactionAmount)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionAmount struct {
	value *WalletTransactionAmount
	isSet bool
}

func (v NullableWalletTransactionAmount) Get() *WalletTransactionAmount {
	return v.value
}

func (v *NullableWalletTransactionAmount) Set(val *WalletTransactionAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionAmount(val *WalletTransactionAmount) *NullableWalletTransactionAmount {
	return &NullableWalletTransactionAmount{value: val, isSet: true}
}

func (v NullableWalletTransactionAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


