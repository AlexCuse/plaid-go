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

// PayStubEarningsTotal An object representing both the current pay period and year to date amount for an earning category.
type PayStubEarningsTotal struct {
	// Total amount of the earnings for this pay period.
	CurrentAmount NullableFloat64 `json:"current_amount"`
	// Total number of hours worked for this pay period.
	Hours NullableFloat32 `json:"hours"`
	// The ISO-4217 currency code of the line item. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the security. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	// The total year-to-date amount of the earnings.
	YtdAmount NullableFloat64 `json:"ytd_amount"`
	AdditionalProperties map[string]interface{}
}

type _PayStubEarningsTotal PayStubEarningsTotal

// NewPayStubEarningsTotal instantiates a new PayStubEarningsTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayStubEarningsTotal(currentAmount NullableFloat64, hours NullableFloat32, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString, ytdAmount NullableFloat64) *PayStubEarningsTotal {
	this := PayStubEarningsTotal{}
	this.CurrentAmount = currentAmount
	this.Hours = hours
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	this.YtdAmount = ytdAmount
	return &this
}

// NewPayStubEarningsTotalWithDefaults instantiates a new PayStubEarningsTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayStubEarningsTotalWithDefaults() *PayStubEarningsTotal {
	this := PayStubEarningsTotal{}
	return &this
}

// GetCurrentAmount returns the CurrentAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubEarningsTotal) GetCurrentAmount() float64 {
	if o == nil || o.CurrentAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.CurrentAmount.Get()
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsTotal) GetCurrentAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentAmount.Get(), o.CurrentAmount.IsSet()
}

// SetCurrentAmount sets field value
func (o *PayStubEarningsTotal) SetCurrentAmount(v float64) {
	o.CurrentAmount.Set(&v)
}

// GetHours returns the Hours field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PayStubEarningsTotal) GetHours() float32 {
	if o == nil || o.Hours.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Hours.Get()
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsTotal) GetHoursOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Hours.Get(), o.Hours.IsSet()
}

// SetHours sets field value
func (o *PayStubEarningsTotal) SetHours(v float32) {
	o.Hours.Set(&v)
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubEarningsTotal) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsTotal) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *PayStubEarningsTotal) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubEarningsTotal) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsTotal) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *PayStubEarningsTotal) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

// GetYtdAmount returns the YtdAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubEarningsTotal) GetYtdAmount() float64 {
	if o == nil || o.YtdAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.YtdAmount.Get()
}

// GetYtdAmountOk returns a tuple with the YtdAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubEarningsTotal) GetYtdAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdAmount.Get(), o.YtdAmount.IsSet()
}

// SetYtdAmount sets field value
func (o *PayStubEarningsTotal) SetYtdAmount(v float64) {
	o.YtdAmount.Set(&v)
}

func (o PayStubEarningsTotal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["current_amount"] = o.CurrentAmount.Get()
	}
	if true {
		toSerialize["hours"] = o.Hours.Get()
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if true {
		toSerialize["ytd_amount"] = o.YtdAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayStubEarningsTotal) UnmarshalJSON(bytes []byte) (err error) {
	varPayStubEarningsTotal := _PayStubEarningsTotal{}

	if err = json.Unmarshal(bytes, &varPayStubEarningsTotal); err == nil {
		*o = PayStubEarningsTotal(varPayStubEarningsTotal)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "hours")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "ytd_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayStubEarningsTotal struct {
	value *PayStubEarningsTotal
	isSet bool
}

func (v NullablePayStubEarningsTotal) Get() *PayStubEarningsTotal {
	return v.value
}

func (v *NullablePayStubEarningsTotal) Set(val *PayStubEarningsTotal) {
	v.value = val
	v.isSet = true
}

func (v NullablePayStubEarningsTotal) IsSet() bool {
	return v.isSet
}

func (v *NullablePayStubEarningsTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayStubEarningsTotal(val *PayStubEarningsTotal) *NullablePayStubEarningsTotal {
	return &NullablePayStubEarningsTotal{value: val, isSet: true}
}

func (v NullablePayStubEarningsTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayStubEarningsTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


