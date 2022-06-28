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

// PayrollIncomeObject An object representing payroll data.
type PayrollIncomeObject struct {
	// ID of the payroll provider account.
	AccountId NullableString `json:"account_id"`
	// Array of pay stubs for the user.
	PayStubs []CreditPayStub `json:"pay_stubs"`
	// Array of tax form W-2s.
	W2s []CreditW2 `json:"w2s"`
	AdditionalProperties map[string]interface{}
}

type _PayrollIncomeObject PayrollIncomeObject

// NewPayrollIncomeObject instantiates a new PayrollIncomeObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayrollIncomeObject(accountId NullableString, payStubs []CreditPayStub, w2s []CreditW2) *PayrollIncomeObject {
	this := PayrollIncomeObject{}
	this.AccountId = accountId
	this.PayStubs = payStubs
	this.W2s = w2s
	return &this
}

// NewPayrollIncomeObjectWithDefaults instantiates a new PayrollIncomeObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayrollIncomeObjectWithDefaults() *PayrollIncomeObject {
	this := PayrollIncomeObject{}
	return &this
}

// GetAccountId returns the AccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayrollIncomeObject) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayrollIncomeObject) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// SetAccountId sets field value
func (o *PayrollIncomeObject) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// GetPayStubs returns the PayStubs field value
func (o *PayrollIncomeObject) GetPayStubs() []CreditPayStub {
	if o == nil {
		var ret []CreditPayStub
		return ret
	}

	return o.PayStubs
}

// GetPayStubsOk returns a tuple with the PayStubs field value
// and a boolean to check if the value has been set.
func (o *PayrollIncomeObject) GetPayStubsOk() (*[]CreditPayStub, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PayStubs, true
}

// SetPayStubs sets field value
func (o *PayrollIncomeObject) SetPayStubs(v []CreditPayStub) {
	o.PayStubs = v
}

// GetW2s returns the W2s field value
func (o *PayrollIncomeObject) GetW2s() []CreditW2 {
	if o == nil {
		var ret []CreditW2
		return ret
	}

	return o.W2s
}

// GetW2sOk returns a tuple with the W2s field value
// and a boolean to check if the value has been set.
func (o *PayrollIncomeObject) GetW2sOk() (*[]CreditW2, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.W2s, true
}

// SetW2s sets field value
func (o *PayrollIncomeObject) SetW2s(v []CreditW2) {
	o.W2s = v
}

func (o PayrollIncomeObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if true {
		toSerialize["pay_stubs"] = o.PayStubs
	}
	if true {
		toSerialize["w2s"] = o.W2s
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayrollIncomeObject) UnmarshalJSON(bytes []byte) (err error) {
	varPayrollIncomeObject := _PayrollIncomeObject{}

	if err = json.Unmarshal(bytes, &varPayrollIncomeObject); err == nil {
		*o = PayrollIncomeObject(varPayrollIncomeObject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "pay_stubs")
		delete(additionalProperties, "w2s")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayrollIncomeObject struct {
	value *PayrollIncomeObject
	isSet bool
}

func (v NullablePayrollIncomeObject) Get() *PayrollIncomeObject {
	return v.value
}

func (v *NullablePayrollIncomeObject) Set(val *PayrollIncomeObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePayrollIncomeObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePayrollIncomeObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayrollIncomeObject(val *PayrollIncomeObject) *NullablePayrollIncomeObject {
	return &NullablePayrollIncomeObject{value: val, isSet: true}
}

func (v NullablePayrollIncomeObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayrollIncomeObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


