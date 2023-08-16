/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.413.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PayrollIncomeAccountData An object containing account level data.
type PayrollIncomeAccountData struct {
	// ID of the payroll provider account.
	AccountId NullableString `json:"account_id"`
	RateOfPay PayrollIncomeRateOfPay `json:"rate_of_pay"`
	// The frequency at which an individual is paid.
	PayFrequency NullableString `json:"pay_frequency"`
	AdditionalProperties map[string]interface{}
}

type _PayrollIncomeAccountData PayrollIncomeAccountData

// NewPayrollIncomeAccountData instantiates a new PayrollIncomeAccountData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayrollIncomeAccountData(accountId NullableString, rateOfPay PayrollIncomeRateOfPay, payFrequency NullableString) *PayrollIncomeAccountData {
	this := PayrollIncomeAccountData{}
	this.AccountId = accountId
	this.RateOfPay = rateOfPay
	this.PayFrequency = payFrequency
	return &this
}

// NewPayrollIncomeAccountDataWithDefaults instantiates a new PayrollIncomeAccountData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayrollIncomeAccountDataWithDefaults() *PayrollIncomeAccountData {
	this := PayrollIncomeAccountData{}
	return &this
}

// GetAccountId returns the AccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayrollIncomeAccountData) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayrollIncomeAccountData) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// SetAccountId sets field value
func (o *PayrollIncomeAccountData) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// GetRateOfPay returns the RateOfPay field value
func (o *PayrollIncomeAccountData) GetRateOfPay() PayrollIncomeRateOfPay {
	if o == nil {
		var ret PayrollIncomeRateOfPay
		return ret
	}

	return o.RateOfPay
}

// GetRateOfPayOk returns a tuple with the RateOfPay field value
// and a boolean to check if the value has been set.
func (o *PayrollIncomeAccountData) GetRateOfPayOk() (*PayrollIncomeRateOfPay, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RateOfPay, true
}

// SetRateOfPay sets field value
func (o *PayrollIncomeAccountData) SetRateOfPay(v PayrollIncomeRateOfPay) {
	o.RateOfPay = v
}

// GetPayFrequency returns the PayFrequency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayrollIncomeAccountData) GetPayFrequency() string {
	if o == nil || o.PayFrequency.Get() == nil {
		var ret string
		return ret
	}

	return *o.PayFrequency.Get()
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayrollIncomeAccountData) GetPayFrequencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayFrequency.Get(), o.PayFrequency.IsSet()
}

// SetPayFrequency sets field value
func (o *PayrollIncomeAccountData) SetPayFrequency(v string) {
	o.PayFrequency.Set(&v)
}

func (o PayrollIncomeAccountData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if true {
		toSerialize["rate_of_pay"] = o.RateOfPay
	}
	if true {
		toSerialize["pay_frequency"] = o.PayFrequency.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayrollIncomeAccountData) UnmarshalJSON(bytes []byte) (err error) {
	varPayrollIncomeAccountData := _PayrollIncomeAccountData{}

	if err = json.Unmarshal(bytes, &varPayrollIncomeAccountData); err == nil {
		*o = PayrollIncomeAccountData(varPayrollIncomeAccountData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "rate_of_pay")
		delete(additionalProperties, "pay_frequency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayrollIncomeAccountData struct {
	value *PayrollIncomeAccountData
	isSet bool
}

func (v NullablePayrollIncomeAccountData) Get() *PayrollIncomeAccountData {
	return v.value
}

func (v *NullablePayrollIncomeAccountData) Set(val *PayrollIncomeAccountData) {
	v.value = val
	v.isSet = true
}

func (v NullablePayrollIncomeAccountData) IsSet() bool {
	return v.isSet
}

func (v *NullablePayrollIncomeAccountData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayrollIncomeAccountData(val *PayrollIncomeAccountData) *NullablePayrollIncomeAccountData {
	return &NullablePayrollIncomeAccountData{value: val, isSet: true}
}

func (v NullablePayrollIncomeAccountData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayrollIncomeAccountData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

