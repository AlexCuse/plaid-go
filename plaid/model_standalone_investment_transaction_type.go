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

// StandaloneInvestmentTransactionType Valid values for investment transaction types and subtypes. Note that transactions representing inflow of cash will appear as negative amounts, outflow of cash will appear as positive amounts.
type StandaloneInvestmentTransactionType struct {
	// Buying an investment
	Buy string `json:"buy"`
	// Selling an investment
	Sell string `json:"sell"`
	// A cancellation of a pending transaction
	Cancel string `json:"cancel"`
	// Activity that modifies a cash position
	Cash string `json:"cash"`
	// Fees on the account, e.g. commission, bookkeeping, options-related.
	Fee string `json:"fee"`
	// Activity that modifies a position, but not through buy/sell activity e.g. options exercise, portfolio transfer
	Transfer string `json:"transfer"`
	AdditionalProperties map[string]interface{}
}

type _StandaloneInvestmentTransactionType StandaloneInvestmentTransactionType

// NewStandaloneInvestmentTransactionType instantiates a new StandaloneInvestmentTransactionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandaloneInvestmentTransactionType(buy string, sell string, cancel string, cash string, fee string, transfer string) *StandaloneInvestmentTransactionType {
	this := StandaloneInvestmentTransactionType{}
	this.Buy = buy
	this.Sell = sell
	this.Cancel = cancel
	this.Cash = cash
	this.Fee = fee
	this.Transfer = transfer
	return &this
}

// NewStandaloneInvestmentTransactionTypeWithDefaults instantiates a new StandaloneInvestmentTransactionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandaloneInvestmentTransactionTypeWithDefaults() *StandaloneInvestmentTransactionType {
	this := StandaloneInvestmentTransactionType{}
	return &this
}

// GetBuy returns the Buy field value
func (o *StandaloneInvestmentTransactionType) GetBuy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Buy
}

// GetBuyOk returns a tuple with the Buy field value
// and a boolean to check if the value has been set.
func (o *StandaloneInvestmentTransactionType) GetBuyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Buy, true
}

// SetBuy sets field value
func (o *StandaloneInvestmentTransactionType) SetBuy(v string) {
	o.Buy = v
}

// GetSell returns the Sell field value
func (o *StandaloneInvestmentTransactionType) GetSell() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sell
}

// GetSellOk returns a tuple with the Sell field value
// and a boolean to check if the value has been set.
func (o *StandaloneInvestmentTransactionType) GetSellOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sell, true
}

// SetSell sets field value
func (o *StandaloneInvestmentTransactionType) SetSell(v string) {
	o.Sell = v
}

// GetCancel returns the Cancel field value
func (o *StandaloneInvestmentTransactionType) GetCancel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value
// and a boolean to check if the value has been set.
func (o *StandaloneInvestmentTransactionType) GetCancelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cancel, true
}

// SetCancel sets field value
func (o *StandaloneInvestmentTransactionType) SetCancel(v string) {
	o.Cancel = v
}

// GetCash returns the Cash field value
func (o *StandaloneInvestmentTransactionType) GetCash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cash
}

// GetCashOk returns a tuple with the Cash field value
// and a boolean to check if the value has been set.
func (o *StandaloneInvestmentTransactionType) GetCashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cash, true
}

// SetCash sets field value
func (o *StandaloneInvestmentTransactionType) SetCash(v string) {
	o.Cash = v
}

// GetFee returns the Fee field value
func (o *StandaloneInvestmentTransactionType) GetFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *StandaloneInvestmentTransactionType) GetFeeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *StandaloneInvestmentTransactionType) SetFee(v string) {
	o.Fee = v
}

// GetTransfer returns the Transfer field value
func (o *StandaloneInvestmentTransactionType) GetTransfer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value
// and a boolean to check if the value has been set.
func (o *StandaloneInvestmentTransactionType) GetTransferOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transfer, true
}

// SetTransfer sets field value
func (o *StandaloneInvestmentTransactionType) SetTransfer(v string) {
	o.Transfer = v
}

func (o StandaloneInvestmentTransactionType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["buy"] = o.Buy
	}
	if true {
		toSerialize["sell"] = o.Sell
	}
	if true {
		toSerialize["cancel"] = o.Cancel
	}
	if true {
		toSerialize["cash"] = o.Cash
	}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if true {
		toSerialize["transfer"] = o.Transfer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StandaloneInvestmentTransactionType) UnmarshalJSON(bytes []byte) (err error) {
	varStandaloneInvestmentTransactionType := _StandaloneInvestmentTransactionType{}

	if err = json.Unmarshal(bytes, &varStandaloneInvestmentTransactionType); err == nil {
		*o = StandaloneInvestmentTransactionType(varStandaloneInvestmentTransactionType)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "buy")
		delete(additionalProperties, "sell")
		delete(additionalProperties, "cancel")
		delete(additionalProperties, "cash")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "transfer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStandaloneInvestmentTransactionType struct {
	value *StandaloneInvestmentTransactionType
	isSet bool
}

func (v NullableStandaloneInvestmentTransactionType) Get() *StandaloneInvestmentTransactionType {
	return v.value
}

func (v *NullableStandaloneInvestmentTransactionType) Set(val *StandaloneInvestmentTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableStandaloneInvestmentTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableStandaloneInvestmentTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandaloneInvestmentTransactionType(val *StandaloneInvestmentTransactionType) *NullableStandaloneInvestmentTransactionType {
	return &NullableStandaloneInvestmentTransactionType{value: val, isSet: true}
}

func (v NullableStandaloneInvestmentTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandaloneInvestmentTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


