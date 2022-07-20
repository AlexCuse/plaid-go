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

// CreditPayStubNetPay An object representing information about the net pay amount on the pay stub.
type CreditPayStubNetPay struct {
	// Raw amount of the net pay for the pay period.
	CurrentAmount NullableFloat64 `json:"current_amount"`
	// Description of the net pay.
	Description NullableString `json:"description"`
	// The ISO-4217 currency code of the net pay. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the net pay. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	// The year-to-date amount of the net pay.
	YtdAmount NullableFloat64 `json:"ytd_amount"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayStubNetPay CreditPayStubNetPay

// NewCreditPayStubNetPay instantiates a new CreditPayStubNetPay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayStubNetPay(currentAmount NullableFloat64, description NullableString, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString, ytdAmount NullableFloat64) *CreditPayStubNetPay {
	this := CreditPayStubNetPay{}
	this.CurrentAmount = currentAmount
	this.Description = description
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	this.YtdAmount = ytdAmount
	return &this
}

// NewCreditPayStubNetPayWithDefaults instantiates a new CreditPayStubNetPay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayStubNetPayWithDefaults() *CreditPayStubNetPay {
	this := CreditPayStubNetPay{}
	return &this
}

// GetCurrentAmount returns the CurrentAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *CreditPayStubNetPay) GetCurrentAmount() float64 {
	if o == nil || o.CurrentAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.CurrentAmount.Get()
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubNetPay) GetCurrentAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentAmount.Get(), o.CurrentAmount.IsSet()
}

// SetCurrentAmount sets field value
func (o *CreditPayStubNetPay) SetCurrentAmount(v float64) {
	o.CurrentAmount.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStubNetPay) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubNetPay) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *CreditPayStubNetPay) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStubNetPay) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubNetPay) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *CreditPayStubNetPay) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPayStubNetPay) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubNetPay) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *CreditPayStubNetPay) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

// GetYtdAmount returns the YtdAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *CreditPayStubNetPay) GetYtdAmount() float64 {
	if o == nil || o.YtdAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.YtdAmount.Get()
}

// GetYtdAmountOk returns a tuple with the YtdAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPayStubNetPay) GetYtdAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdAmount.Get(), o.YtdAmount.IsSet()
}

// SetYtdAmount sets field value
func (o *CreditPayStubNetPay) SetYtdAmount(v float64) {
	o.YtdAmount.Set(&v)
}

func (o CreditPayStubNetPay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["current_amount"] = o.CurrentAmount.Get()
	}
	if true {
		toSerialize["description"] = o.Description.Get()
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

func (o *CreditPayStubNetPay) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayStubNetPay := _CreditPayStubNetPay{}

	if err = json.Unmarshal(bytes, &varCreditPayStubNetPay); err == nil {
		*o = CreditPayStubNetPay(varCreditPayStubNetPay)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "description")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "ytd_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayStubNetPay struct {
	value *CreditPayStubNetPay
	isSet bool
}

func (v NullableCreditPayStubNetPay) Get() *CreditPayStubNetPay {
	return v.value
}

func (v *NullableCreditPayStubNetPay) Set(val *CreditPayStubNetPay) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayStubNetPay) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayStubNetPay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayStubNetPay(val *CreditPayStubNetPay) *NullableCreditPayStubNetPay {
	return &NullableCreditPayStubNetPay{value: val, isSet: true}
}

func (v NullableCreditPayStubNetPay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayStubNetPay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


