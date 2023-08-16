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

// Pay An object representing a monetary amount.
type Pay struct {
	// A numerical amount of a specific currency.
	Amount NullableFloat64 `json:"amount,omitempty"`
	// Currency code, e.g. USD
	Currency NullableString `json:"currency,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Pay Pay

// NewPay instantiates a new Pay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPay() *Pay {
	this := Pay{}
	return &this
}

// NewPayWithDefaults instantiates a new Pay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayWithDefaults() *Pay {
	this := Pay{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pay) GetAmount() float64 {
	if o == nil || o.Amount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pay) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *Pay) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *Pay) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *Pay) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *Pay) UnsetAmount() {
	o.Amount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Pay) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pay) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *Pay) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *Pay) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *Pay) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *Pay) UnsetCurrency() {
	o.Currency.Unset()
}

func (o Pay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Pay) UnmarshalJSON(bytes []byte) (err error) {
	varPay := _Pay{}

	if err = json.Unmarshal(bytes, &varPay); err == nil {
		*o = Pay(varPay)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePay struct {
	value *Pay
	isSet bool
}

func (v NullablePay) Get() *Pay {
	return v.value
}

func (v *NullablePay) Set(val *Pay) {
	v.value = val
	v.isSet = true
}

func (v NullablePay) IsSet() bool {
	return v.isSet
}

func (v *NullablePay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePay(val *Pay) *NullablePay {
	return &NullablePay{value: val, isSet: true}
}

func (v NullablePay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


