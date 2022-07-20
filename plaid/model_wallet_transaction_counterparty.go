/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.146.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// WalletTransactionCounterparty An object representing the e-wallet transaction's counterparty
type WalletTransactionCounterparty struct {
	// The name of the counterparty
	Name string `json:"name"`
	Numbers WalletTransactionCounterpartyNumbers `json:"numbers"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionCounterparty WalletTransactionCounterparty

// NewWalletTransactionCounterparty instantiates a new WalletTransactionCounterparty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionCounterparty(name string, numbers WalletTransactionCounterpartyNumbers) *WalletTransactionCounterparty {
	this := WalletTransactionCounterparty{}
	this.Name = name
	this.Numbers = numbers
	return &this
}

// NewWalletTransactionCounterpartyWithDefaults instantiates a new WalletTransactionCounterparty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionCounterpartyWithDefaults() *WalletTransactionCounterparty {
	this := WalletTransactionCounterparty{}
	return &this
}

// GetName returns the Name field value
func (o *WalletTransactionCounterparty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionCounterparty) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WalletTransactionCounterparty) SetName(v string) {
	o.Name = v
}

// GetNumbers returns the Numbers field value
func (o *WalletTransactionCounterparty) GetNumbers() WalletTransactionCounterpartyNumbers {
	if o == nil {
		var ret WalletTransactionCounterpartyNumbers
		return ret
	}

	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionCounterparty) GetNumbersOk() (*WalletTransactionCounterpartyNumbers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Numbers, true
}

// SetNumbers sets field value
func (o *WalletTransactionCounterparty) SetNumbers(v WalletTransactionCounterpartyNumbers) {
	o.Numbers = v
}

func (o WalletTransactionCounterparty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["numbers"] = o.Numbers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransactionCounterparty) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionCounterparty := _WalletTransactionCounterparty{}

	if err = json.Unmarshal(bytes, &varWalletTransactionCounterparty); err == nil {
		*o = WalletTransactionCounterparty(varWalletTransactionCounterparty)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "numbers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionCounterparty struct {
	value *WalletTransactionCounterparty
	isSet bool
}

func (v NullableWalletTransactionCounterparty) Get() *WalletTransactionCounterparty {
	return v.value
}

func (v *NullableWalletTransactionCounterparty) Set(val *WalletTransactionCounterparty) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionCounterparty) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionCounterparty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionCounterparty(val *WalletTransactionCounterparty) *NullableWalletTransactionCounterparty {
	return &NullableWalletTransactionCounterparty{value: val, isSet: true}
}

func (v NullableWalletTransactionCounterparty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionCounterparty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


