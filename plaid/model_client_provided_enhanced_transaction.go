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

// ClientProvidedEnhancedTransaction A client-provided transaction that Plaid has enhanced.
type ClientProvidedEnhancedTransaction struct {
	// Unique transaction identifier to tie transactions back to clients' systems.
	Id string `json:"id"`
	// The raw description of the transaction.
	Description string `json:"description"`
	// The value of the transaction, denominated in the account's currency, as stated in `iso_currency_code`. Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative.
	Amount float64 `json:"amount"`
	// The ISO-4217 currency code of the transaction.
	IsoCurrencyCode string `json:"iso_currency_code"`
	Enhancements Enhancements `json:"enhancements"`
	AdditionalProperties map[string]interface{}
}

type _ClientProvidedEnhancedTransaction ClientProvidedEnhancedTransaction

// NewClientProvidedEnhancedTransaction instantiates a new ClientProvidedEnhancedTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProvidedEnhancedTransaction(id string, description string, amount float64, isoCurrencyCode string, enhancements Enhancements) *ClientProvidedEnhancedTransaction {
	this := ClientProvidedEnhancedTransaction{}
	this.Id = id
	this.Description = description
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	this.Enhancements = enhancements
	return &this
}

// NewClientProvidedEnhancedTransactionWithDefaults instantiates a new ClientProvidedEnhancedTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProvidedEnhancedTransactionWithDefaults() *ClientProvidedEnhancedTransaction {
	this := ClientProvidedEnhancedTransaction{}
	return &this
}

// GetId returns the Id field value
func (o *ClientProvidedEnhancedTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedEnhancedTransaction) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClientProvidedEnhancedTransaction) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *ClientProvidedEnhancedTransaction) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedEnhancedTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ClientProvidedEnhancedTransaction) SetDescription(v string) {
	o.Description = v
}

// GetAmount returns the Amount field value
func (o *ClientProvidedEnhancedTransaction) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedEnhancedTransaction) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ClientProvidedEnhancedTransaction) SetAmount(v float64) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *ClientProvidedEnhancedTransaction) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedEnhancedTransaction) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *ClientProvidedEnhancedTransaction) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetEnhancements returns the Enhancements field value
func (o *ClientProvidedEnhancedTransaction) GetEnhancements() Enhancements {
	if o == nil {
		var ret Enhancements
		return ret
	}

	return o.Enhancements
}

// GetEnhancementsOk returns a tuple with the Enhancements field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedEnhancedTransaction) GetEnhancementsOk() (*Enhancements, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Enhancements, true
}

// SetEnhancements sets field value
func (o *ClientProvidedEnhancedTransaction) SetEnhancements(v Enhancements) {
	o.Enhancements = v
}

func (o ClientProvidedEnhancedTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if true {
		toSerialize["enhancements"] = o.Enhancements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ClientProvidedEnhancedTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varClientProvidedEnhancedTransaction := _ClientProvidedEnhancedTransaction{}

	if err = json.Unmarshal(bytes, &varClientProvidedEnhancedTransaction); err == nil {
		*o = ClientProvidedEnhancedTransaction(varClientProvidedEnhancedTransaction)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "enhancements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientProvidedEnhancedTransaction struct {
	value *ClientProvidedEnhancedTransaction
	isSet bool
}

func (v NullableClientProvidedEnhancedTransaction) Get() *ClientProvidedEnhancedTransaction {
	return v.value
}

func (v *NullableClientProvidedEnhancedTransaction) Set(val *ClientProvidedEnhancedTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProvidedEnhancedTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProvidedEnhancedTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProvidedEnhancedTransaction(val *ClientProvidedEnhancedTransaction) *NullableClientProvidedEnhancedTransaction {
	return &NullableClientProvidedEnhancedTransaction{value: val, isSet: true}
}

func (v NullableClientProvidedEnhancedTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProvidedEnhancedTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


