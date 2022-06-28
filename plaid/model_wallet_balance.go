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

// WalletBalance An object representing the e-wallet balance
type WalletBalance struct {
	// The ISO-4217 currency code of the balance
	IsoCurrencyCode string `json:"iso_currency_code"`
	// The total amount of funds in the account
	Current float64 `json:"current"`
	AdditionalProperties map[string]interface{}
}

type _WalletBalance WalletBalance

// NewWalletBalance instantiates a new WalletBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletBalance(isoCurrencyCode string, current float64) *WalletBalance {
	this := WalletBalance{}
	this.IsoCurrencyCode = isoCurrencyCode
	this.Current = current
	return &this
}

// NewWalletBalanceWithDefaults instantiates a new WalletBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletBalanceWithDefaults() *WalletBalance {
	this := WalletBalance{}
	return &this
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *WalletBalance) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *WalletBalance) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *WalletBalance) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetCurrent returns the Current field value
func (o *WalletBalance) GetCurrent() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *WalletBalance) GetCurrentOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *WalletBalance) SetCurrent(v float64) {
	o.Current = v
}

func (o WalletBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["current"] = o.Current
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletBalance) UnmarshalJSON(bytes []byte) (err error) {
	varWalletBalance := _WalletBalance{}

	if err = json.Unmarshal(bytes, &varWalletBalance); err == nil {
		*o = WalletBalance(varWalletBalance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "current")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletBalance struct {
	value *WalletBalance
	isSet bool
}

func (v NullableWalletBalance) Get() *WalletBalance {
	return v.value
}

func (v *NullableWalletBalance) Set(val *WalletBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletBalance(val *WalletBalance) *NullableWalletBalance {
	return &NullableWalletBalance{value: val, isSet: true}
}

func (v NullableWalletBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


